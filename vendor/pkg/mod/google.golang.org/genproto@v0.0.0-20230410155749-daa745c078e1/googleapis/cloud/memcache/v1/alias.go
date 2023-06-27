// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package memcache aliases all exported identifiers in package
// "cloud.google.com/go/memcache/apiv1/memcachepb".
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package memcache

import (
	src "cloud.google.com/go/memcache/apiv1/memcachepb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/memcache/apiv1/memcachepb
const (
	Instance_CREATING                                     = src.Instance_CREATING
	Instance_DELETING                                     = src.Instance_DELETING
	Instance_InstanceMessage_CODE_UNSPECIFIED             = src.Instance_InstanceMessage_CODE_UNSPECIFIED
	Instance_InstanceMessage_ZONE_DISTRIBUTION_UNBALANCED = src.Instance_InstanceMessage_ZONE_DISTRIBUTION_UNBALANCED
	Instance_Node_CREATING                                = src.Instance_Node_CREATING
	Instance_Node_DELETING                                = src.Instance_Node_DELETING
	Instance_Node_READY                                   = src.Instance_Node_READY
	Instance_Node_STATE_UNSPECIFIED                       = src.Instance_Node_STATE_UNSPECIFIED
	Instance_Node_UPDATING                                = src.Instance_Node_UPDATING
	Instance_PERFORMING_MAINTENANCE                       = src.Instance_PERFORMING_MAINTENANCE
	Instance_READY                                        = src.Instance_READY
	Instance_STATE_UNSPECIFIED                            = src.Instance_STATE_UNSPECIFIED
	MemcacheVersion_MEMCACHE_1_5                          = src.MemcacheVersion_MEMCACHE_1_5
	MemcacheVersion_MEMCACHE_VERSION_UNSPECIFIED          = src.MemcacheVersion_MEMCACHE_VERSION_UNSPECIFIED
)

// Deprecated: Please use vars in: cloud.google.com/go/memcache/apiv1/memcachepb
var (
	File_google_cloud_memcache_v1_cloud_memcache_proto = src.File_google_cloud_memcache_v1_cloud_memcache_proto
	Instance_InstanceMessage_Code_name                 = src.Instance_InstanceMessage_Code_name
	Instance_InstanceMessage_Code_value                = src.Instance_InstanceMessage_Code_value
	Instance_Node_State_name                           = src.Instance_Node_State_name
	Instance_Node_State_value                          = src.Instance_Node_State_value
	Instance_State_name                                = src.Instance_State_name
	Instance_State_value                               = src.Instance_State_value
	MemcacheVersion_name                               = src.MemcacheVersion_name
	MemcacheVersion_value                              = src.MemcacheVersion_value
)

// Request for
// [ApplyParameters][google.cloud.memcache.v1.CloudMemcache.ApplyParameters].
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type ApplyParametersRequest = src.ApplyParametersRequest

// CloudMemcacheClient is the client API for CloudMemcache service. For
// semantics around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type CloudMemcacheClient = src.CloudMemcacheClient

// CloudMemcacheServer is the server API for CloudMemcache service.
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type CloudMemcacheServer = src.CloudMemcacheServer

// Request for
// [CreateInstance][google.cloud.memcache.v1.CloudMemcache.CreateInstance].
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type CreateInstanceRequest = src.CreateInstanceRequest

// Request for
// [DeleteInstance][google.cloud.memcache.v1.CloudMemcache.DeleteInstance].
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type DeleteInstanceRequest = src.DeleteInstanceRequest

// Request for
// [GetInstance][google.cloud.memcache.v1.CloudMemcache.GetInstance].
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type GetInstanceRequest = src.GetInstanceRequest
type Instance = src.Instance
type Instance_InstanceMessage = src.Instance_InstanceMessage
type Instance_InstanceMessage_Code = src.Instance_InstanceMessage_Code
type Instance_Node = src.Instance_Node

// Configuration for a Memcached Node.
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type Instance_NodeConfig = src.Instance_NodeConfig

// Different states of a Memcached node.
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type Instance_Node_State = src.Instance_Node_State

// Different states of a Memcached instance.
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type Instance_State = src.Instance_State

// Request for
// [ListInstances][google.cloud.memcache.v1.CloudMemcache.ListInstances].
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type ListInstancesRequest = src.ListInstancesRequest

// Response for
// [ListInstances][google.cloud.memcache.v1.CloudMemcache.ListInstances].
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type ListInstancesResponse = src.ListInstancesResponse
type MemcacheParameters = src.MemcacheParameters

// Memcached versions supported by our service.
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type MemcacheVersion = src.MemcacheVersion

// Represents the metadata of a long-running operation.
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type OperationMetadata = src.OperationMetadata

// UnimplementedCloudMemcacheServer can be embedded to have forward compatible
// implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type UnimplementedCloudMemcacheServer = src.UnimplementedCloudMemcacheServer

// Request for
// [UpdateInstance][google.cloud.memcache.v1.CloudMemcache.UpdateInstance].
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type UpdateInstanceRequest = src.UpdateInstanceRequest

// Request for
// [UpdateParameters][google.cloud.memcache.v1.CloudMemcache.UpdateParameters].
//
// Deprecated: Please use types in: cloud.google.com/go/memcache/apiv1/memcachepb
type UpdateParametersRequest = src.UpdateParametersRequest

// Deprecated: Please use funcs in: cloud.google.com/go/memcache/apiv1/memcachepb
func NewCloudMemcacheClient(cc grpc.ClientConnInterface) CloudMemcacheClient {
	return src.NewCloudMemcacheClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/memcache/apiv1/memcachepb
func RegisterCloudMemcacheServer(s *grpc.Server, srv CloudMemcacheServer) {
	src.RegisterCloudMemcacheServer(s, srv)
}
