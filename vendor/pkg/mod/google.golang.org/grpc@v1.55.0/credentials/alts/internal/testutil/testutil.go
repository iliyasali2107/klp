/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package testutil include useful test utilities for the handshaker.
package testutil

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/alts/internal/conn"
	altsgrpc "google.golang.org/grpc/credentials/alts/internal/proto/grpc_gcp"
	altspb "google.golang.org/grpc/credentials/alts/internal/proto/grpc_gcp"
)

// Stats is used to collect statistics about concurrent handshake calls.
type Stats struct {
	mu                 sync.Mutex
	calls              int
	MaxConcurrentCalls int
}

// Update updates the statistics by adding one call.
func (s *Stats) Update() func() {
	s.mu.Lock()
	s.calls++
	if s.calls > s.MaxConcurrentCalls {
		s.MaxConcurrentCalls = s.calls
	}
	s.mu.Unlock()

	return func() {
		s.mu.Lock()
		s.calls--
		s.mu.Unlock()
	}
}

// Reset resets the statistics.
func (s *Stats) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.calls = 0
	s.MaxConcurrentCalls = 0
}

// testConn mimics a net.Conn to the peer.
type testConn struct {
	net.Conn
	in  *bytes.Buffer
	out *bytes.Buffer
}

// NewTestConn creates a new instance of testConn object.
func NewTestConn(in *bytes.Buffer, out *bytes.Buffer) net.Conn {
	return &testConn{
		in:  in,
		out: out,
	}
}

// Read reads from the in buffer.
func (c *testConn) Read(b []byte) (n int, err error) {
	return c.in.Read(b)
}

// Write writes to the out buffer.
func (c *testConn) Write(b []byte) (n int, err error) {
	return c.out.Write(b)
}

// Close closes the testConn object.
func (c *testConn) Close() error {
	return nil
}

// unresponsiveTestConn mimics a net.Conn for an unresponsive peer. It is used
// for testing the PeerNotResponding case.
type unresponsiveTestConn struct {
	net.Conn
}

// NewUnresponsiveTestConn creates a new instance of unresponsiveTestConn object.
func NewUnresponsiveTestConn() net.Conn {
	return &unresponsiveTestConn{}
}

// Read reads from the in buffer.
func (c *unresponsiveTestConn) Read(b []byte) (n int, err error) {
	return 0, io.EOF
}

// Write writes to the out buffer.
func (c *unresponsiveTestConn) Write(b []byte) (n int, err error) {
	return 0, nil
}

// Close closes the TestConn object.
func (c *unresponsiveTestConn) Close() error {
	return nil
}

// MakeFrame creates a handshake frame.
func MakeFrame(pl string) []byte {
	f := make([]byte, len(pl)+conn.MsgLenFieldSize)
	binary.LittleEndian.PutUint32(f, uint32(len(pl)))
	copy(f[conn.MsgLenFieldSize:], []byte(pl))
	return f
}

// FakeHandshaker is a fake implementation of the ALTS handshaker service.
type FakeHandshaker struct {
	altsgrpc.HandshakerServiceServer
}

// DoHandshake performs a fake ALTS handshake.
func (h *FakeHandshaker) DoHandshake(stream altsgrpc.HandshakerService_DoHandshakeServer) error {
	var isAssistingClient bool
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("stream recv failure: %v", err)
		}
		var resp *altspb.HandshakerResp
		switch req := req.ReqOneof.(type) {
		case *altspb.HandshakerReq_ClientStart:
			isAssistingClient = true
			resp, err = h.processStartClient(req.ClientStart)
			if err != nil {
				return fmt.Errorf("processStartClient failure: %v", err)
			}
		case *altspb.HandshakerReq_ServerStart:
			isAssistingClient = false
			resp, err = h.processServerStart(req.ServerStart)
			if err != nil {
				return fmt.Errorf("processServerClient failure: %v", err)
			}
		case *altspb.HandshakerReq_Next:
			resp, err = h.processNext(req.Next, isAssistingClient)
			if err != nil {
				return fmt.Errorf("processNext failure: %v", err)
			}
		default:
			return fmt.Errorf("handshake request has unexpected type: %v", req)
		}

		if err = stream.Send(resp); err != nil {
			return fmt.Errorf("stream send failure: %v", err)
		}
	}
}

func (h *FakeHandshaker) processStartClient(req *altspb.StartClientHandshakeReq) (*altspb.HandshakerResp, error) {
	if req.HandshakeSecurityProtocol != altspb.HandshakeProtocol_ALTS {
		return nil, fmt.Errorf("unexpected handshake security protocol: %v", req.HandshakeSecurityProtocol)
	}
	if len(req.ApplicationProtocols) != 1 || req.ApplicationProtocols[0] != "grpc" {
		return nil, fmt.Errorf("unexpected application protocols: %v", req.ApplicationProtocols)
	}
	if len(req.RecordProtocols) != 1 || req.RecordProtocols[0] != "ALTSRP_GCM_AES128_REKEY" {
		return nil, fmt.Errorf("unexpected record protocols: %v", req.RecordProtocols)
	}
	return &altspb.HandshakerResp{
		OutFrames:     []byte("ClientInit"),
		BytesConsumed: 0,
		Status: &altspb.HandshakerStatus{
			Code: uint32(codes.OK),
		},
	}, nil
}

func (h *FakeHandshaker) processServerStart(req *altspb.StartServerHandshakeReq) (*altspb.HandshakerResp, error) {
	if len(req.ApplicationProtocols) != 1 || req.ApplicationProtocols[0] != "grpc" {
		return nil, fmt.Errorf("unexpected application protocols: %v", req.ApplicationProtocols)
	}
	parameters, ok := req.GetHandshakeParameters()[int32(altspb.HandshakeProtocol_ALTS)]
	if !ok {
		return nil, fmt.Errorf("missing ALTS handshake parameters")
	}
	if len(parameters.RecordProtocols) != 1 || parameters.RecordProtocols[0] != "ALTSRP_GCM_AES128_REKEY" {
		return nil, fmt.Errorf("unexpected record protocols: %v", parameters.RecordProtocols)
	}
	if string(req.InBytes) != "ClientInit" {
		return nil, fmt.Errorf("unexpected in bytes: %v", req.InBytes)
	}
	return &altspb.HandshakerResp{
		OutFrames:     []byte("ServerInitServerFinished"),
		BytesConsumed: 10,
		Status: &altspb.HandshakerStatus{
			Code: uint32(codes.OK),
		},
	}, nil
}

func (h *FakeHandshaker) processNext(req *altspb.NextHandshakeMessageReq, isAssistingClient bool) (*altspb.HandshakerResp, error) {
	if isAssistingClient {
		if !bytes.Equal(req.InBytes, []byte("ServerInitServerFinished")) {
			return nil, fmt.Errorf("unexpected in bytes: got: %v, want: %v", req.InBytes, []byte("ServerInitServerFinished"))
		}
		return &altspb.HandshakerResp{
			OutFrames:     []byte("ClientFinished"),
			BytesConsumed: 24,
			Result: &altspb.HandshakerResult{
				ApplicationProtocol: "grpc",
				RecordProtocol:      "ALTSRP_GCM_AES128_REKEY",
				KeyData:             []byte("negotiated-key-data-for-altsrp-gcm-aes128-rekey"),
				PeerIdentity: &altspb.Identity{
					IdentityOneof: &altspb.Identity_ServiceAccount{
						ServiceAccount: "server@bar.com",
					},
				},
				PeerRpcVersions: &altspb.RpcProtocolVersions{
					MaxRpcVersion: &altspb.RpcProtocolVersions_Version{
						Minor: 1,
						Major: 2,
					},
					MinRpcVersion: &altspb.RpcProtocolVersions_Version{
						Minor: 1,
						Major: 2,
					},
				},
			},
			Status: &altspb.HandshakerStatus{
				Code: uint32(codes.OK),
			},
		}, nil
	}
	if !bytes.Equal(req.InBytes, []byte("ClientFinished")) {
		return nil, fmt.Errorf("unexpected in bytes: got: %v, want: %v", req.InBytes, []byte("ClientFinished"))
	}
	return &altspb.HandshakerResp{
		BytesConsumed: 14,
		Result: &altspb.HandshakerResult{
			ApplicationProtocol: "grpc",
			RecordProtocol:      "ALTSRP_GCM_AES128_REKEY",
			KeyData:             []byte("negotiated-key-data-for-altsrp-gcm-aes128-rekey"),
			PeerIdentity: &altspb.Identity{
				IdentityOneof: &altspb.Identity_ServiceAccount{
					ServiceAccount: "client@baz.com",
				},
			},
			PeerRpcVersions: &altspb.RpcProtocolVersions{
				MaxRpcVersion: &altspb.RpcProtocolVersions_Version{
					Minor: 1,
					Major: 2,
				},
				MinRpcVersion: &altspb.RpcProtocolVersions_Version{
					Minor: 1,
					Major: 2,
				},
			},
		},
		Status: &altspb.HandshakerStatus{
			Code: uint32(codes.OK),
		},
	}, nil
}
