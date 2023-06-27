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

// Package securitycenter aliases all exported identifiers in package
// "cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb".
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package securitycenter

import (
	src "cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
const (
	Finding_ACTIVE                                                       = src.Finding_ACTIVE
	Finding_CRITICAL                                                     = src.Finding_CRITICAL
	Finding_HIGH                                                         = src.Finding_HIGH
	Finding_INACTIVE                                                     = src.Finding_INACTIVE
	Finding_LOW                                                          = src.Finding_LOW
	Finding_MEDIUM                                                       = src.Finding_MEDIUM
	Finding_SEVERITY_UNSPECIFIED                                         = src.Finding_SEVERITY_UNSPECIFIED
	Finding_STATE_UNSPECIFIED                                            = src.Finding_STATE_UNSPECIFIED
	ListAssetsResponse_ListAssetsResult_ACTIVE                           = src.ListAssetsResponse_ListAssetsResult_ACTIVE
	ListAssetsResponse_ListAssetsResult_ADDED                            = src.ListAssetsResponse_ListAssetsResult_ADDED
	ListAssetsResponse_ListAssetsResult_REMOVED                          = src.ListAssetsResponse_ListAssetsResult_REMOVED
	ListAssetsResponse_ListAssetsResult_UNUSED                           = src.ListAssetsResponse_ListAssetsResult_UNUSED
	ListFindingsResponse_ListFindingsResult_ADDED                        = src.ListFindingsResponse_ListFindingsResult_ADDED
	ListFindingsResponse_ListFindingsResult_CHANGED                      = src.ListFindingsResponse_ListFindingsResult_CHANGED
	ListFindingsResponse_ListFindingsResult_REMOVED                      = src.ListFindingsResponse_ListFindingsResult_REMOVED
	ListFindingsResponse_ListFindingsResult_UNCHANGED                    = src.ListFindingsResponse_ListFindingsResult_UNCHANGED
	ListFindingsResponse_ListFindingsResult_UNUSED                       = src.ListFindingsResponse_ListFindingsResult_UNUSED
	NotificationConfig_EVENT_TYPE_UNSPECIFIED                            = src.NotificationConfig_EVENT_TYPE_UNSPECIFIED
	NotificationConfig_FINDING                                           = src.NotificationConfig_FINDING
	OrganizationSettings_AssetDiscoveryConfig_EXCLUDE                    = src.OrganizationSettings_AssetDiscoveryConfig_EXCLUDE
	OrganizationSettings_AssetDiscoveryConfig_INCLUDE_ONLY               = src.OrganizationSettings_AssetDiscoveryConfig_INCLUDE_ONLY
	OrganizationSettings_AssetDiscoveryConfig_INCLUSION_MODE_UNSPECIFIED = src.OrganizationSettings_AssetDiscoveryConfig_INCLUSION_MODE_UNSPECIFIED
	RunAssetDiscoveryResponse_COMPLETED                                  = src.RunAssetDiscoveryResponse_COMPLETED
	RunAssetDiscoveryResponse_STATE_UNSPECIFIED                          = src.RunAssetDiscoveryResponse_STATE_UNSPECIFIED
	RunAssetDiscoveryResponse_SUPERSEDED                                 = src.RunAssetDiscoveryResponse_SUPERSEDED
	RunAssetDiscoveryResponse_TERMINATED                                 = src.RunAssetDiscoveryResponse_TERMINATED
)

// Deprecated: Please use vars in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
var (
	File_google_cloud_securitycenter_v1p1beta1_asset_proto                        = src.File_google_cloud_securitycenter_v1p1beta1_asset_proto
	File_google_cloud_securitycenter_v1p1beta1_finding_proto                      = src.File_google_cloud_securitycenter_v1p1beta1_finding_proto
	File_google_cloud_securitycenter_v1p1beta1_folder_proto                       = src.File_google_cloud_securitycenter_v1p1beta1_folder_proto
	File_google_cloud_securitycenter_v1p1beta1_notification_config_proto          = src.File_google_cloud_securitycenter_v1p1beta1_notification_config_proto
	File_google_cloud_securitycenter_v1p1beta1_notification_message_proto         = src.File_google_cloud_securitycenter_v1p1beta1_notification_message_proto
	File_google_cloud_securitycenter_v1p1beta1_organization_settings_proto        = src.File_google_cloud_securitycenter_v1p1beta1_organization_settings_proto
	File_google_cloud_securitycenter_v1p1beta1_resource_proto                     = src.File_google_cloud_securitycenter_v1p1beta1_resource_proto
	File_google_cloud_securitycenter_v1p1beta1_run_asset_discovery_response_proto = src.File_google_cloud_securitycenter_v1p1beta1_run_asset_discovery_response_proto
	File_google_cloud_securitycenter_v1p1beta1_security_marks_proto               = src.File_google_cloud_securitycenter_v1p1beta1_security_marks_proto
	File_google_cloud_securitycenter_v1p1beta1_securitycenter_service_proto       = src.File_google_cloud_securitycenter_v1p1beta1_securitycenter_service_proto
	File_google_cloud_securitycenter_v1p1beta1_source_proto                       = src.File_google_cloud_securitycenter_v1p1beta1_source_proto
	Finding_Severity_name                                                         = src.Finding_Severity_name
	Finding_Severity_value                                                        = src.Finding_Severity_value
	Finding_State_name                                                            = src.Finding_State_name
	Finding_State_value                                                           = src.Finding_State_value
	ListAssetsResponse_ListAssetsResult_StateChange_name                          = src.ListAssetsResponse_ListAssetsResult_StateChange_name
	ListAssetsResponse_ListAssetsResult_StateChange_value                         = src.ListAssetsResponse_ListAssetsResult_StateChange_value
	ListFindingsResponse_ListFindingsResult_StateChange_name                      = src.ListFindingsResponse_ListFindingsResult_StateChange_name
	ListFindingsResponse_ListFindingsResult_StateChange_value                     = src.ListFindingsResponse_ListFindingsResult_StateChange_value
	NotificationConfig_EventType_name                                             = src.NotificationConfig_EventType_name
	NotificationConfig_EventType_value                                            = src.NotificationConfig_EventType_value
	OrganizationSettings_AssetDiscoveryConfig_InclusionMode_name                  = src.OrganizationSettings_AssetDiscoveryConfig_InclusionMode_name
	OrganizationSettings_AssetDiscoveryConfig_InclusionMode_value                 = src.OrganizationSettings_AssetDiscoveryConfig_InclusionMode_value
	RunAssetDiscoveryResponse_State_name                                          = src.RunAssetDiscoveryResponse_State_name
	RunAssetDiscoveryResponse_State_value                                         = src.RunAssetDiscoveryResponse_State_value
)

// Security Command Center representation of a Google Cloud resource. The
// Asset is a Security Command Center resource that captures information about
// a single Google Cloud resource. All modifications to an Asset are only
// within the context of Security Command Center and don't affect the
// referenced Google Cloud resource.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type Asset = src.Asset

// Cloud IAM Policy information associated with the Google Cloud resource
// described by the Security Command Center asset. This information is managed
// and defined by the Google Cloud resource and cannot be modified by the user.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type Asset_IamPolicy = src.Asset_IamPolicy

// Security Command Center managed properties. These properties are managed by
// Security Command Center and cannot be modified by the user.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type Asset_SecurityCenterProperties = src.Asset_SecurityCenterProperties

// Request message for creating a finding.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type CreateFindingRequest = src.CreateFindingRequest

// Request message for creating a notification config.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type CreateNotificationConfigRequest = src.CreateNotificationConfigRequest

// Request message for creating a source.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type CreateSourceRequest = src.CreateSourceRequest

// Request message for deleting a notification config.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type DeleteNotificationConfigRequest = src.DeleteNotificationConfigRequest

// Security Command Center finding. A finding is a record of assessment data
// (security, risk, health or privacy) ingested into Security Command Center
// for presentation, notification, analysis, policy testing, and enforcement.
// For example, an XSS vulnerability in an App Engine application is a finding.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type Finding = src.Finding

// The severity of the finding. This field is managed by the source that
// writes the finding.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type Finding_Severity = src.Finding_Severity

// The state of the finding.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type Finding_State = src.Finding_State

// Message that contains the resource name and display name of a folder
// resource.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type Folder = src.Folder

// Request message for getting a notification config.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type GetNotificationConfigRequest = src.GetNotificationConfigRequest

// Request message for getting organization settings.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type GetOrganizationSettingsRequest = src.GetOrganizationSettingsRequest

// Request message for getting a source.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type GetSourceRequest = src.GetSourceRequest

// Request message for grouping by assets.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type GroupAssetsRequest = src.GroupAssetsRequest

// Response message for grouping by assets.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type GroupAssetsResponse = src.GroupAssetsResponse

// Request message for grouping by findings.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type GroupFindingsRequest = src.GroupFindingsRequest

// Response message for group by findings.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type GroupFindingsResponse = src.GroupFindingsResponse

// Result containing the properties and count of a groupBy request.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type GroupResult = src.GroupResult

// Request message for listing assets.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type ListAssetsRequest = src.ListAssetsRequest

// Response message for listing assets.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type ListAssetsResponse = src.ListAssetsResponse

// Result containing the Asset and its State.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type ListAssetsResponse_ListAssetsResult = src.ListAssetsResponse_ListAssetsResult

// The change in state of the asset. When querying across two points in time
// this describes the change between the two points: ADDED, REMOVED, or ACTIVE.
// If there was no compare_duration supplied in the request the state change
// will be: UNUSED
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type ListAssetsResponse_ListAssetsResult_StateChange = src.ListAssetsResponse_ListAssetsResult_StateChange

// Request message for listing findings.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type ListFindingsRequest = src.ListFindingsRequest

// Response message for listing findings.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type ListFindingsResponse = src.ListFindingsResponse

// Result containing the Finding and its StateChange.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type ListFindingsResponse_ListFindingsResult = src.ListFindingsResponse_ListFindingsResult

// Information related to the Google Cloud resource that is associated with
// this finding.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type ListFindingsResponse_ListFindingsResult_Resource = src.ListFindingsResponse_ListFindingsResult_Resource

// The change in state of the finding. When querying across two points in time
// this describes the change in the finding between the two points: CHANGED,
// UNCHANGED, ADDED, or REMOVED. Findings can not be deleted, so REMOVED
// implies that the finding at timestamp does not match the filter specified,
// but it did at timestamp - compare_duration. If there was no compare_duration
// supplied in the request the state change will be: UNUSED
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type ListFindingsResponse_ListFindingsResult_StateChange = src.ListFindingsResponse_ListFindingsResult_StateChange

// Request message for listing notification configs.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type ListNotificationConfigsRequest = src.ListNotificationConfigsRequest

// Response message for listing notification configs.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type ListNotificationConfigsResponse = src.ListNotificationConfigsResponse

// Request message for listing sources.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type ListSourcesRequest = src.ListSourcesRequest

// Response message for listing sources.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type ListSourcesResponse = src.ListSourcesResponse

// Security Command Center notification configs. A notification config is a
// Security Command Center resource that contains the configuration to send
// notifications for create/update events of findings, assets and etc.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type NotificationConfig = src.NotificationConfig

// The type of events.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type NotificationConfig_EventType = src.NotificationConfig_EventType

// The config for streaming-based notifications, which send each event as soon
// as it is detected.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type NotificationConfig_StreamingConfig = src.NotificationConfig_StreamingConfig
type NotificationConfig_StreamingConfig_ = src.NotificationConfig_StreamingConfig_

// Security Command Center's Notification
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type NotificationMessage = src.NotificationMessage
type NotificationMessage_Finding = src.NotificationMessage_Finding

// User specified settings that are attached to the Security Command Center
// organization.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type OrganizationSettings = src.OrganizationSettings

// The configuration used for Asset Discovery runs.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type OrganizationSettings_AssetDiscoveryConfig = src.OrganizationSettings_AssetDiscoveryConfig

// The mode of inclusion when running Asset Discovery. Asset discovery can be
// limited by explicitly identifying projects to be included or excluded. If
// INCLUDE_ONLY is set, then only those projects within the organization and
// their children are discovered during asset discovery. If EXCLUDE is set,
// then projects that don't match those projects are discovered during asset
// discovery. If neither are set, then all projects within the organization are
// discovered during asset discovery.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type OrganizationSettings_AssetDiscoveryConfig_InclusionMode = src.OrganizationSettings_AssetDiscoveryConfig_InclusionMode

// Information related to the Google Cloud resource.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type Resource = src.Resource

// Request message for running asset discovery for an organization.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type RunAssetDiscoveryRequest = src.RunAssetDiscoveryRequest

// Response of asset discovery run
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type RunAssetDiscoveryResponse = src.RunAssetDiscoveryResponse

// The state of an asset discovery run.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type RunAssetDiscoveryResponse_State = src.RunAssetDiscoveryResponse_State

// SecurityCenterClient is the client API for SecurityCenter service. For
// semantics around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type SecurityCenterClient = src.SecurityCenterClient

// SecurityCenterServer is the server API for SecurityCenter service.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type SecurityCenterServer = src.SecurityCenterServer

// User specified security marks that are attached to the parent Security
// Command Center resource. Security marks are scoped within a Security Command
// Center organization -- they can be modified and viewed by all users who have
// proper permissions on the organization.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type SecurityMarks = src.SecurityMarks

// Request message for updating a finding's state.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type SetFindingStateRequest = src.SetFindingStateRequest

// Security Command Center finding source. A finding source is an entity or a
// mechanism that can produce a finding. A source is like a container of
// findings that come from the same scanner, logger, monitor, etc.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type Source = src.Source

// UnimplementedSecurityCenterServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type UnimplementedSecurityCenterServer = src.UnimplementedSecurityCenterServer

// Request message for updating or creating a finding.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type UpdateFindingRequest = src.UpdateFindingRequest

// Request message for updating a notification config.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type UpdateNotificationConfigRequest = src.UpdateNotificationConfigRequest

// Request message for updating an organization's settings.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type UpdateOrganizationSettingsRequest = src.UpdateOrganizationSettingsRequest

// Request message for updating a SecurityMarks resource.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type UpdateSecurityMarksRequest = src.UpdateSecurityMarksRequest

// Request message for updating a source.
//
// Deprecated: Please use types in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
type UpdateSourceRequest = src.UpdateSourceRequest

// Deprecated: Please use funcs in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
func NewSecurityCenterClient(cc grpc.ClientConnInterface) SecurityCenterClient {
	return src.NewSecurityCenterClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb
func RegisterSecurityCenterServer(s *grpc.Server, srv SecurityCenterServer) {
	src.RegisterSecurityCenterServer(s, srv)
}
