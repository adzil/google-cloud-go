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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.2
// source: google/cloud/beyondcorp/appconnectors/v1/app_connector_instance_config.proto

package appconnectorspb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AppConnectorInstanceConfig defines the instance config of a AppConnector.
type AppConnectorInstanceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. A monotonically increasing number generated and maintained
	// by the API provider. Every time a config changes in the backend, the
	// sequenceNumber should be bumped up to reflect the change.
	SequenceNumber int64 `protobuf:"varint,1,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	// The SLM instance agent configuration.
	InstanceConfig *anypb.Any `protobuf:"bytes,2,opt,name=instance_config,json=instanceConfig,proto3" json:"instance_config,omitempty"`
	// NotificationConfig defines the notification mechanism that the remote
	// instance should subscribe to in order to receive notification.
	NotificationConfig *NotificationConfig `protobuf:"bytes,3,opt,name=notification_config,json=notificationConfig,proto3" json:"notification_config,omitempty"`
	// ImageConfig defines the GCR images to run for the remote agent's control
	// plane.
	ImageConfig *ImageConfig `protobuf:"bytes,4,opt,name=image_config,json=imageConfig,proto3" json:"image_config,omitempty"`
}

func (x *AppConnectorInstanceConfig) Reset() {
	*x = AppConnectorInstanceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConnectorInstanceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConnectorInstanceConfig) ProtoMessage() {}

func (x *AppConnectorInstanceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConnectorInstanceConfig.ProtoReflect.Descriptor instead.
func (*AppConnectorInstanceConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_rawDescGZIP(), []int{0}
}

func (x *AppConnectorInstanceConfig) GetSequenceNumber() int64 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}

func (x *AppConnectorInstanceConfig) GetInstanceConfig() *anypb.Any {
	if x != nil {
		return x.InstanceConfig
	}
	return nil
}

func (x *AppConnectorInstanceConfig) GetNotificationConfig() *NotificationConfig {
	if x != nil {
		return x.NotificationConfig
	}
	return nil
}

func (x *AppConnectorInstanceConfig) GetImageConfig() *ImageConfig {
	if x != nil {
		return x.ImageConfig
	}
	return nil
}

// NotificationConfig defines the mechanisms to notify instance agent.
type NotificationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Config:
	//
	//	*NotificationConfig_PubsubNotification
	Config isNotificationConfig_Config `protobuf_oneof:"config"`
}

func (x *NotificationConfig) Reset() {
	*x = NotificationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationConfig) ProtoMessage() {}

func (x *NotificationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationConfig.ProtoReflect.Descriptor instead.
func (*NotificationConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_rawDescGZIP(), []int{1}
}

func (m *NotificationConfig) GetConfig() isNotificationConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *NotificationConfig) GetPubsubNotification() *NotificationConfig_CloudPubSubNotificationConfig {
	if x, ok := x.GetConfig().(*NotificationConfig_PubsubNotification); ok {
		return x.PubsubNotification
	}
	return nil
}

type isNotificationConfig_Config interface {
	isNotificationConfig_Config()
}

type NotificationConfig_PubsubNotification struct {
	// Cloud Pub/Sub Configuration to receive notifications.
	PubsubNotification *NotificationConfig_CloudPubSubNotificationConfig `protobuf:"bytes,1,opt,name=pubsub_notification,json=pubsubNotification,proto3,oneof"`
}

func (*NotificationConfig_PubsubNotification) isNotificationConfig_Config() {}

// ImageConfig defines the control plane images to run.
type ImageConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The initial image the remote agent will attempt to run for the control
	// plane.
	TargetImage string `protobuf:"bytes,1,opt,name=target_image,json=targetImage,proto3" json:"target_image,omitempty"`
	// The stable image that the remote agent will fallback to if the target image
	// fails.
	StableImage string `protobuf:"bytes,2,opt,name=stable_image,json=stableImage,proto3" json:"stable_image,omitempty"`
}

func (x *ImageConfig) Reset() {
	*x = ImageConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageConfig) ProtoMessage() {}

func (x *ImageConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageConfig.ProtoReflect.Descriptor instead.
func (*ImageConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_rawDescGZIP(), []int{2}
}

func (x *ImageConfig) GetTargetImage() string {
	if x != nil {
		return x.TargetImage
	}
	return ""
}

func (x *ImageConfig) GetStableImage() string {
	if x != nil {
		return x.StableImage
	}
	return ""
}

// The configuration for Pub/Sub messaging for the AppConnector.
type NotificationConfig_CloudPubSubNotificationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Pub/Sub subscription the AppConnector uses to receive notifications.
	PubsubSubscription string `protobuf:"bytes,1,opt,name=pubsub_subscription,json=pubsubSubscription,proto3" json:"pubsub_subscription,omitempty"`
}

func (x *NotificationConfig_CloudPubSubNotificationConfig) Reset() {
	*x = NotificationConfig_CloudPubSubNotificationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationConfig_CloudPubSubNotificationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationConfig_CloudPubSubNotificationConfig) ProtoMessage() {}

func (x *NotificationConfig_CloudPubSubNotificationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationConfig_CloudPubSubNotificationConfig.ProtoReflect.Descriptor instead.
func (*NotificationConfig_CloudPubSubNotificationConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_rawDescGZIP(), []int{1, 0}
}

func (x *NotificationConfig_CloudPubSubNotificationConfig) GetPubsubSubscription() string {
	if x != nil {
		return x.PubsubSubscription
	}
	return ""
}

var File_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto protoreflect.FileDescriptor

var file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79,
	0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x02, 0x0a, 0x1a, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x3d, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x6d, 0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79,
	0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x12, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x58, 0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x80, 0x02, 0x0a, 0x12, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x8d, 0x01, 0x0a, 0x13, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65,
	0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x12, 0x70, 0x75,
	0x62, 0x73, 0x75, 0x62, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x50, 0x0a, 0x1d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x53, 0x0a, 0x0b,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x42, 0xaa, 0x02, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x1f, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x65, 0x79, 0x6f, 0x6e,
	0x64, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x70, 0x62, 0x3b, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x70, 0x62, 0xaa, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x43,
	0x6f, 0x72, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x42, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x43, 0x6f, 0x72, 0x70, 0x5c, 0x41,
	0x70, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x2c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x42, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x43, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x41, 0x70, 0x70,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_rawDescOnce sync.Once
	file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_rawDescData = file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_rawDesc
)

func file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_rawDescGZIP() []byte {
	file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_rawDescData)
	})
	return file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_rawDescData
}

var file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_goTypes = []interface{}{
	(*AppConnectorInstanceConfig)(nil),                       // 0: google.cloud.beyondcorp.appconnectors.v1.AppConnectorInstanceConfig
	(*NotificationConfig)(nil),                               // 1: google.cloud.beyondcorp.appconnectors.v1.NotificationConfig
	(*ImageConfig)(nil),                                      // 2: google.cloud.beyondcorp.appconnectors.v1.ImageConfig
	(*NotificationConfig_CloudPubSubNotificationConfig)(nil), // 3: google.cloud.beyondcorp.appconnectors.v1.NotificationConfig.CloudPubSubNotificationConfig
	(*anypb.Any)(nil),                                        // 4: google.protobuf.Any
}
var file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_depIdxs = []int32{
	4, // 0: google.cloud.beyondcorp.appconnectors.v1.AppConnectorInstanceConfig.instance_config:type_name -> google.protobuf.Any
	1, // 1: google.cloud.beyondcorp.appconnectors.v1.AppConnectorInstanceConfig.notification_config:type_name -> google.cloud.beyondcorp.appconnectors.v1.NotificationConfig
	2, // 2: google.cloud.beyondcorp.appconnectors.v1.AppConnectorInstanceConfig.image_config:type_name -> google.cloud.beyondcorp.appconnectors.v1.ImageConfig
	3, // 3: google.cloud.beyondcorp.appconnectors.v1.NotificationConfig.pubsub_notification:type_name -> google.cloud.beyondcorp.appconnectors.v1.NotificationConfig.CloudPubSubNotificationConfig
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_init() }
func file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_init() {
	if File_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConnectorInstanceConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationConfig_CloudPubSubNotificationConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*NotificationConfig_PubsubNotification)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_depIdxs,
		MessageInfos:      file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_msgTypes,
	}.Build()
	File_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto = out.File
	file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_rawDesc = nil
	file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_goTypes = nil
	file_google_cloud_beyondcorp_appconnectors_v1_app_connector_instance_config_proto_depIdxs = nil
}
