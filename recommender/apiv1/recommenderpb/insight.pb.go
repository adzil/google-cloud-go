// Copyright 2023 Google LLC
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
// source: google/cloud/recommender/v1/insight.proto

package recommenderpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Insight category.
type Insight_Category int32

const (
	// Unspecified category.
	Insight_CATEGORY_UNSPECIFIED Insight_Category = 0
	// The insight is related to cost.
	Insight_COST Insight_Category = 1
	// The insight is related to security.
	Insight_SECURITY Insight_Category = 2
	// The insight is related to performance.
	Insight_PERFORMANCE Insight_Category = 3
	// This insight is related to manageability.
	Insight_MANAGEABILITY Insight_Category = 4
	// The insight is related to sustainability.
	Insight_SUSTAINABILITY Insight_Category = 5
	// This insight is related to reliability.
	Insight_RELIABILITY Insight_Category = 6
)

// Enum value maps for Insight_Category.
var (
	Insight_Category_name = map[int32]string{
		0: "CATEGORY_UNSPECIFIED",
		1: "COST",
		2: "SECURITY",
		3: "PERFORMANCE",
		4: "MANAGEABILITY",
		5: "SUSTAINABILITY",
		6: "RELIABILITY",
	}
	Insight_Category_value = map[string]int32{
		"CATEGORY_UNSPECIFIED": 0,
		"COST":                 1,
		"SECURITY":             2,
		"PERFORMANCE":          3,
		"MANAGEABILITY":        4,
		"SUSTAINABILITY":       5,
		"RELIABILITY":          6,
	}
)

func (x Insight_Category) Enum() *Insight_Category {
	p := new(Insight_Category)
	*p = x
	return p
}

func (x Insight_Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Insight_Category) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_recommender_v1_insight_proto_enumTypes[0].Descriptor()
}

func (Insight_Category) Type() protoreflect.EnumType {
	return &file_google_cloud_recommender_v1_insight_proto_enumTypes[0]
}

func (x Insight_Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Insight_Category.Descriptor instead.
func (Insight_Category) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_recommender_v1_insight_proto_rawDescGZIP(), []int{0, 0}
}

// Insight severity levels.
type Insight_Severity int32

const (
	// Insight has unspecified severity.
	Insight_SEVERITY_UNSPECIFIED Insight_Severity = 0
	// Insight has low severity.
	Insight_LOW Insight_Severity = 1
	// Insight has medium severity.
	Insight_MEDIUM Insight_Severity = 2
	// Insight has high severity.
	Insight_HIGH Insight_Severity = 3
	// Insight has critical severity.
	Insight_CRITICAL Insight_Severity = 4
)

// Enum value maps for Insight_Severity.
var (
	Insight_Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "LOW",
		2: "MEDIUM",
		3: "HIGH",
		4: "CRITICAL",
	}
	Insight_Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"LOW":                  1,
		"MEDIUM":               2,
		"HIGH":                 3,
		"CRITICAL":             4,
	}
)

func (x Insight_Severity) Enum() *Insight_Severity {
	p := new(Insight_Severity)
	*p = x
	return p
}

func (x Insight_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Insight_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_recommender_v1_insight_proto_enumTypes[1].Descriptor()
}

func (Insight_Severity) Type() protoreflect.EnumType {
	return &file_google_cloud_recommender_v1_insight_proto_enumTypes[1]
}

func (x Insight_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Insight_Severity.Descriptor instead.
func (Insight_Severity) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_recommender_v1_insight_proto_rawDescGZIP(), []int{0, 1}
}

// Represents insight state.
type InsightStateInfo_State int32

const (
	// Unspecified state.
	InsightStateInfo_STATE_UNSPECIFIED InsightStateInfo_State = 0
	// Insight is active. Content for ACTIVE insights can be updated by Google.
	// ACTIVE insights can be marked DISMISSED OR ACCEPTED.
	InsightStateInfo_ACTIVE InsightStateInfo_State = 1
	// Some action has been taken based on this insight. Insights become
	// accepted when a recommendation derived from the insight has been marked
	// CLAIMED, SUCCEEDED, or FAILED. ACTIVE insights can also be marked
	// ACCEPTED explicitly. Content for ACCEPTED insights is immutable. ACCEPTED
	// insights can only be marked ACCEPTED (which may update state metadata).
	InsightStateInfo_ACCEPTED InsightStateInfo_State = 2
	// Insight is dismissed. Content for DISMISSED insights can be updated by
	// Google. DISMISSED insights can be marked as ACTIVE.
	InsightStateInfo_DISMISSED InsightStateInfo_State = 3
)

// Enum value maps for InsightStateInfo_State.
var (
	InsightStateInfo_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ACTIVE",
		2: "ACCEPTED",
		3: "DISMISSED",
	}
	InsightStateInfo_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ACTIVE":            1,
		"ACCEPTED":          2,
		"DISMISSED":         3,
	}
)

func (x InsightStateInfo_State) Enum() *InsightStateInfo_State {
	p := new(InsightStateInfo_State)
	*p = x
	return p
}

func (x InsightStateInfo_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InsightStateInfo_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_recommender_v1_insight_proto_enumTypes[2].Descriptor()
}

func (InsightStateInfo_State) Type() protoreflect.EnumType {
	return &file_google_cloud_recommender_v1_insight_proto_enumTypes[2]
}

func (x InsightStateInfo_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InsightStateInfo_State.Descriptor instead.
func (InsightStateInfo_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_recommender_v1_insight_proto_rawDescGZIP(), []int{1, 0}
}

// An insight along with the information used to derive the insight. The insight
// may have associated recommendations as well.
type Insight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the insight.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Free-form human readable summary in English. The maximum length is 500
	// characters.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Fully qualified resource names that this insight is targeting.
	TargetResources []string `protobuf:"bytes,9,rep,name=target_resources,json=targetResources,proto3" json:"target_resources,omitempty"`
	// Insight subtype. Insight content schema will be stable for a given subtype.
	InsightSubtype string `protobuf:"bytes,10,opt,name=insight_subtype,json=insightSubtype,proto3" json:"insight_subtype,omitempty"`
	// A struct of custom fields to explain the insight.
	// Example: "grantedPermissionsCount": "1000"
	Content *structpb.Struct `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// Timestamp of the latest data used to generate the insight.
	LastRefreshTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=last_refresh_time,json=lastRefreshTime,proto3" json:"last_refresh_time,omitempty"`
	// Observation period that led to the insight. The source data used to
	// generate the insight ends at last_refresh_time and begins at
	// (last_refresh_time - observation_period).
	ObservationPeriod *durationpb.Duration `protobuf:"bytes,5,opt,name=observation_period,json=observationPeriod,proto3" json:"observation_period,omitempty"`
	// Information state and metadata.
	StateInfo *InsightStateInfo `protobuf:"bytes,6,opt,name=state_info,json=stateInfo,proto3" json:"state_info,omitempty"`
	// Category being targeted by the insight.
	Category Insight_Category `protobuf:"varint,7,opt,name=category,proto3,enum=google.cloud.recommender.v1.Insight_Category" json:"category,omitempty"`
	// Insight's severity.
	Severity Insight_Severity `protobuf:"varint,15,opt,name=severity,proto3,enum=google.cloud.recommender.v1.Insight_Severity" json:"severity,omitempty"`
	// Fingerprint of the Insight. Provides optimistic locking when updating
	// states.
	Etag string `protobuf:"bytes,11,opt,name=etag,proto3" json:"etag,omitempty"`
	// Recommendations derived from this insight.
	AssociatedRecommendations []*Insight_RecommendationReference `protobuf:"bytes,8,rep,name=associated_recommendations,json=associatedRecommendations,proto3" json:"associated_recommendations,omitempty"`
}

func (x *Insight) Reset() {
	*x = Insight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_recommender_v1_insight_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Insight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Insight) ProtoMessage() {}

func (x *Insight) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_recommender_v1_insight_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Insight.ProtoReflect.Descriptor instead.
func (*Insight) Descriptor() ([]byte, []int) {
	return file_google_cloud_recommender_v1_insight_proto_rawDescGZIP(), []int{0}
}

func (x *Insight) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Insight) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Insight) GetTargetResources() []string {
	if x != nil {
		return x.TargetResources
	}
	return nil
}

func (x *Insight) GetInsightSubtype() string {
	if x != nil {
		return x.InsightSubtype
	}
	return ""
}

func (x *Insight) GetContent() *structpb.Struct {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Insight) GetLastRefreshTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastRefreshTime
	}
	return nil
}

func (x *Insight) GetObservationPeriod() *durationpb.Duration {
	if x != nil {
		return x.ObservationPeriod
	}
	return nil
}

func (x *Insight) GetStateInfo() *InsightStateInfo {
	if x != nil {
		return x.StateInfo
	}
	return nil
}

func (x *Insight) GetCategory() Insight_Category {
	if x != nil {
		return x.Category
	}
	return Insight_CATEGORY_UNSPECIFIED
}

func (x *Insight) GetSeverity() Insight_Severity {
	if x != nil {
		return x.Severity
	}
	return Insight_SEVERITY_UNSPECIFIED
}

func (x *Insight) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *Insight) GetAssociatedRecommendations() []*Insight_RecommendationReference {
	if x != nil {
		return x.AssociatedRecommendations
	}
	return nil
}

// Information related to insight state.
type InsightStateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Insight state.
	State InsightStateInfo_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.recommender.v1.InsightStateInfo_State" json:"state,omitempty"`
	// A map of metadata for the state, provided by user or automations systems.
	StateMetadata map[string]string `protobuf:"bytes,2,rep,name=state_metadata,json=stateMetadata,proto3" json:"state_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *InsightStateInfo) Reset() {
	*x = InsightStateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_recommender_v1_insight_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsightStateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsightStateInfo) ProtoMessage() {}

func (x *InsightStateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_recommender_v1_insight_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsightStateInfo.ProtoReflect.Descriptor instead.
func (*InsightStateInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_recommender_v1_insight_proto_rawDescGZIP(), []int{1}
}

func (x *InsightStateInfo) GetState() InsightStateInfo_State {
	if x != nil {
		return x.State
	}
	return InsightStateInfo_STATE_UNSPECIFIED
}

func (x *InsightStateInfo) GetStateMetadata() map[string]string {
	if x != nil {
		return x.StateMetadata
	}
	return nil
}

// Reference to an associated recommendation.
type Insight_RecommendationReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Recommendation resource name, e.g.
	// projects/[PROJECT_NUMBER]/locations/[LOCATION]/recommenders/[RECOMMENDER_ID]/recommendations/[RECOMMENDATION_ID]
	Recommendation string `protobuf:"bytes,1,opt,name=recommendation,proto3" json:"recommendation,omitempty"`
}

func (x *Insight_RecommendationReference) Reset() {
	*x = Insight_RecommendationReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_recommender_v1_insight_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Insight_RecommendationReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Insight_RecommendationReference) ProtoMessage() {}

func (x *Insight_RecommendationReference) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_recommender_v1_insight_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Insight_RecommendationReference.ProtoReflect.Descriptor instead.
func (*Insight_RecommendationReference) Descriptor() ([]byte, []int) {
	return file_google_cloud_recommender_v1_insight_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Insight_RecommendationReference) GetRecommendation() string {
	if x != nil {
		return x.Recommendation
	}
	return ""
}

var File_google_cloud_recommender_v1_insight_proto protoreflect.FileDescriptor

var file_google_cloud_recommender_v1_insight_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8d, 0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68,
	0x74, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x11, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x4c, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x49, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x69,
	0x67, 0x68, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x49, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x53,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x7b, 0x0a, 0x1a, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x19, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x74, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x41, 0x0a, 0x17, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x85, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x43, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49,
	0x54, 0x59, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x45, 0x52, 0x46, 0x4f, 0x52, 0x4d, 0x41,
	0x4e, 0x43, 0x45, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x55, 0x53, 0x54,
	0x41, 0x49, 0x4e, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b,
	0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x06, 0x22, 0x51, 0x0a,
	0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x56,
	0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x49, 0x47, 0x48,
	0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x04,
	0x3a, 0x9f, 0x03, 0xea, 0x41, 0x9b, 0x03, 0x0a, 0x22, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x12, 0x56, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d,
	0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x69, 0x67,
	0x68, 0x74, 0x7d, 0x12, 0x65, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x69, 0x6e, 0x73, 0x69,
	0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x7d, 0x12, 0x54, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x7d, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x7b,
	0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x2f, 0x69, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x7d,
	0x12, 0x60, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x2f,
	0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68,
	0x74, 0x7d, 0x22, 0xd1, 0x02, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x49, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x67, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x40, 0x0a, 0x12, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x47, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43,
	0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x49, 0x53, 0x4d, 0x49,
	0x53, 0x53, 0x45, 0x44, 0x10, 0x03, 0x42, 0xf0, 0x03, 0xea, 0x41, 0xd3, 0x02, 0x0a, 0x26, 0x72,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x43, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x69,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x12, 0x52, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x7b, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x12, 0x41,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x7d,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x7d, 0x12, 0x4d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d,
	0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x0c, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x41, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x70, 0x62, 0xa2, 0x02, 0x04, 0x43, 0x52, 0x45, 0x43, 0xaa, 0x02, 0x1b, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_recommender_v1_insight_proto_rawDescOnce sync.Once
	file_google_cloud_recommender_v1_insight_proto_rawDescData = file_google_cloud_recommender_v1_insight_proto_rawDesc
)

func file_google_cloud_recommender_v1_insight_proto_rawDescGZIP() []byte {
	file_google_cloud_recommender_v1_insight_proto_rawDescOnce.Do(func() {
		file_google_cloud_recommender_v1_insight_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_recommender_v1_insight_proto_rawDescData)
	})
	return file_google_cloud_recommender_v1_insight_proto_rawDescData
}

var file_google_cloud_recommender_v1_insight_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_google_cloud_recommender_v1_insight_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_recommender_v1_insight_proto_goTypes = []interface{}{
	(Insight_Category)(0),                   // 0: google.cloud.recommender.v1.Insight.Category
	(Insight_Severity)(0),                   // 1: google.cloud.recommender.v1.Insight.Severity
	(InsightStateInfo_State)(0),             // 2: google.cloud.recommender.v1.InsightStateInfo.State
	(*Insight)(nil),                         // 3: google.cloud.recommender.v1.Insight
	(*InsightStateInfo)(nil),                // 4: google.cloud.recommender.v1.InsightStateInfo
	(*Insight_RecommendationReference)(nil), // 5: google.cloud.recommender.v1.Insight.RecommendationReference
	nil,                                     // 6: google.cloud.recommender.v1.InsightStateInfo.StateMetadataEntry
	(*structpb.Struct)(nil),                 // 7: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil),           // 8: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),             // 9: google.protobuf.Duration
}
var file_google_cloud_recommender_v1_insight_proto_depIdxs = []int32{
	7, // 0: google.cloud.recommender.v1.Insight.content:type_name -> google.protobuf.Struct
	8, // 1: google.cloud.recommender.v1.Insight.last_refresh_time:type_name -> google.protobuf.Timestamp
	9, // 2: google.cloud.recommender.v1.Insight.observation_period:type_name -> google.protobuf.Duration
	4, // 3: google.cloud.recommender.v1.Insight.state_info:type_name -> google.cloud.recommender.v1.InsightStateInfo
	0, // 4: google.cloud.recommender.v1.Insight.category:type_name -> google.cloud.recommender.v1.Insight.Category
	1, // 5: google.cloud.recommender.v1.Insight.severity:type_name -> google.cloud.recommender.v1.Insight.Severity
	5, // 6: google.cloud.recommender.v1.Insight.associated_recommendations:type_name -> google.cloud.recommender.v1.Insight.RecommendationReference
	2, // 7: google.cloud.recommender.v1.InsightStateInfo.state:type_name -> google.cloud.recommender.v1.InsightStateInfo.State
	6, // 8: google.cloud.recommender.v1.InsightStateInfo.state_metadata:type_name -> google.cloud.recommender.v1.InsightStateInfo.StateMetadataEntry
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_cloud_recommender_v1_insight_proto_init() }
func file_google_cloud_recommender_v1_insight_proto_init() {
	if File_google_cloud_recommender_v1_insight_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_recommender_v1_insight_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Insight); i {
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
		file_google_cloud_recommender_v1_insight_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsightStateInfo); i {
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
		file_google_cloud_recommender_v1_insight_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Insight_RecommendationReference); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_recommender_v1_insight_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_recommender_v1_insight_proto_goTypes,
		DependencyIndexes: file_google_cloud_recommender_v1_insight_proto_depIdxs,
		EnumInfos:         file_google_cloud_recommender_v1_insight_proto_enumTypes,
		MessageInfos:      file_google_cloud_recommender_v1_insight_proto_msgTypes,
	}.Build()
	File_google_cloud_recommender_v1_insight_proto = out.File
	file_google_cloud_recommender_v1_insight_proto_rawDesc = nil
	file_google_cloud_recommender_v1_insight_proto_goTypes = nil
	file_google_cloud_recommender_v1_insight_proto_depIdxs = nil
}
