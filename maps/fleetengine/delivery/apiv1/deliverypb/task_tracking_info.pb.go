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
// source: google/maps/fleetengine/delivery/v1/task_tracking_info.proto

package deliverypb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The `TaskTrackingInfo` message. The message contains task tracking
// information which will be used for display. If a tracking ID is associated
// with multiple Tasks, Fleet Engine uses a heuristic to decide which Task's
// TaskTrackingInfo to select.
type TaskTrackingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Must be in the format `providers/{provider}/taskTrackingInfo/{tracking}`,
	// where `tracking` represents the tracking ID.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Immutable. The tracking ID of a Task.
	// * Must be a valid Unicode string.
	// * Limited to a maximum length of 64 characters.
	// * Normalized according to [Unicode Normalization Form C]
	// (http://www.unicode.org/reports/tr15/).
	// * May not contain any of the following ASCII characters: '/', ':', '?',
	// ',', or '#'.
	TrackingId string `protobuf:"bytes,2,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
	// The vehicle's last location.
	VehicleLocation *DeliveryVehicleLocation `protobuf:"bytes,3,opt,name=vehicle_location,json=vehicleLocation,proto3" json:"vehicle_location,omitempty"`
	// A list of points which when connected forms a polyline of the vehicle's
	// expected route to the location of this task.
	RoutePolylinePoints []*latlng.LatLng `protobuf:"bytes,4,rep,name=route_polyline_points,json=routePolylinePoints,proto3" json:"route_polyline_points,omitempty"`
	// Indicates the number of stops the vehicle remaining until the task stop is
	// reached, including the task stop. For example, if the vehicle's next stop
	// is the task stop, the value will be 1.
	RemainingStopCount *wrapperspb.Int32Value `protobuf:"bytes,5,opt,name=remaining_stop_count,json=remainingStopCount,proto3" json:"remaining_stop_count,omitempty"`
	// The total remaining distance in meters to the `VehicleStop` of interest.
	RemainingDrivingDistanceMeters *wrapperspb.Int32Value `protobuf:"bytes,6,opt,name=remaining_driving_distance_meters,json=remainingDrivingDistanceMeters,proto3" json:"remaining_driving_distance_meters,omitempty"`
	// The timestamp that indicates the estimated arrival time to the stop
	// location.
	EstimatedArrivalTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=estimated_arrival_time,json=estimatedArrivalTime,proto3" json:"estimated_arrival_time,omitempty"`
	// The timestamp that indicates the estimated completion time of a Task.
	EstimatedTaskCompletionTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=estimated_task_completion_time,json=estimatedTaskCompletionTime,proto3" json:"estimated_task_completion_time,omitempty"`
	// The current execution state of the Task.
	State Task_State `protobuf:"varint,11,opt,name=state,proto3,enum=maps.fleetengine.delivery.v1.Task_State" json:"state,omitempty"`
	// The outcome of attempting to execute a Task.
	TaskOutcome Task_TaskOutcome `protobuf:"varint,9,opt,name=task_outcome,json=taskOutcome,proto3,enum=maps.fleetengine.delivery.v1.Task_TaskOutcome" json:"task_outcome,omitempty"`
	// The timestamp that indicates when the Task's outcome was set by the
	// provider.
	TaskOutcomeTime *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=task_outcome_time,json=taskOutcomeTime,proto3" json:"task_outcome_time,omitempty"`
	// Immutable. The location where the Task will be completed.
	PlannedLocation *LocationInfo `protobuf:"bytes,10,opt,name=planned_location,json=plannedLocation,proto3" json:"planned_location,omitempty"`
	// The time window during which the task should be completed.
	TargetTimeWindow *TimeWindow `protobuf:"bytes,13,opt,name=target_time_window,json=targetTimeWindow,proto3" json:"target_time_window,omitempty"`
	// The custom attributes set on the task.
	Attributes []*TaskAttribute `protobuf:"bytes,14,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *TaskTrackingInfo) Reset() {
	*x = TaskTrackingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskTrackingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTrackingInfo) ProtoMessage() {}

func (x *TaskTrackingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskTrackingInfo.ProtoReflect.Descriptor instead.
func (*TaskTrackingInfo) Descriptor() ([]byte, []int) {
	return file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_rawDescGZIP(), []int{0}
}

func (x *TaskTrackingInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskTrackingInfo) GetTrackingId() string {
	if x != nil {
		return x.TrackingId
	}
	return ""
}

func (x *TaskTrackingInfo) GetVehicleLocation() *DeliveryVehicleLocation {
	if x != nil {
		return x.VehicleLocation
	}
	return nil
}

func (x *TaskTrackingInfo) GetRoutePolylinePoints() []*latlng.LatLng {
	if x != nil {
		return x.RoutePolylinePoints
	}
	return nil
}

func (x *TaskTrackingInfo) GetRemainingStopCount() *wrapperspb.Int32Value {
	if x != nil {
		return x.RemainingStopCount
	}
	return nil
}

func (x *TaskTrackingInfo) GetRemainingDrivingDistanceMeters() *wrapperspb.Int32Value {
	if x != nil {
		return x.RemainingDrivingDistanceMeters
	}
	return nil
}

func (x *TaskTrackingInfo) GetEstimatedArrivalTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EstimatedArrivalTime
	}
	return nil
}

func (x *TaskTrackingInfo) GetEstimatedTaskCompletionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EstimatedTaskCompletionTime
	}
	return nil
}

func (x *TaskTrackingInfo) GetState() Task_State {
	if x != nil {
		return x.State
	}
	return Task_STATE_UNSPECIFIED
}

func (x *TaskTrackingInfo) GetTaskOutcome() Task_TaskOutcome {
	if x != nil {
		return x.TaskOutcome
	}
	return Task_TASK_OUTCOME_UNSPECIFIED
}

func (x *TaskTrackingInfo) GetTaskOutcomeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.TaskOutcomeTime
	}
	return nil
}

func (x *TaskTrackingInfo) GetPlannedLocation() *LocationInfo {
	if x != nil {
		return x.PlannedLocation
	}
	return nil
}

func (x *TaskTrackingInfo) GetTargetTimeWindow() *TimeWindow {
	if x != nil {
		return x.TargetTimeWindow
	}
	return nil
}

func (x *TaskTrackingInfo) GetAttributes() []*TaskAttribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

var File_google_maps_fleetengine_delivery_v1_task_tracking_info_proto protoreflect.FileDescriptor

var file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x6d, 0x61, 0x70, 0x73, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x09, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x12, 0x60, 0x0a, 0x10, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x15, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x6f,
	0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x13, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x4d, 0x0a,
	0x14, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x6f, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x66, 0x0a, 0x21,
	0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e,
	0x67, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x1e, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44,
	0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x50, 0x0a, 0x16, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x14, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x72, 0x69, 0x76,
	0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5f, 0x0a, 0x1e, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x1b, 0x65, 0x73, 0x74, 0x69,
	0x6d, 0x61, 0x74, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x0b, 0x74,
	0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x11, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x5a, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0f, 0x70,
	0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x56,
	0x0a, 0x12, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d, 0x61, 0x70,
	0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x52, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x4b, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6d, 0x61, 0x70,
	0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x3a, 0x62, 0xea, 0x41, 0x5f, 0x0a, 0x2b, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x7d, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x2f, 0x7b, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x7d, 0x42, 0xb6, 0x01, 0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x42,
	0x15, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61,
	0x70, 0x73, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x70, 0x62, 0x3b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x70, 0x62, 0xa2, 0x02, 0x04, 0x43, 0x46, 0x45, 0x44, 0xaa, 0x02, 0x23, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_rawDescOnce sync.Once
	file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_rawDescData = file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_rawDesc
)

func file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_rawDescGZIP() []byte {
	file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_rawDescOnce.Do(func() {
		file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_rawDescData)
	})
	return file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_rawDescData
}

var file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_goTypes = []interface{}{
	(*TaskTrackingInfo)(nil),        // 0: maps.fleetengine.delivery.v1.TaskTrackingInfo
	(*DeliveryVehicleLocation)(nil), // 1: maps.fleetengine.delivery.v1.DeliveryVehicleLocation
	(*latlng.LatLng)(nil),           // 2: google.type.LatLng
	(*wrapperspb.Int32Value)(nil),   // 3: google.protobuf.Int32Value
	(*timestamppb.Timestamp)(nil),   // 4: google.protobuf.Timestamp
	(Task_State)(0),                 // 5: maps.fleetengine.delivery.v1.Task.State
	(Task_TaskOutcome)(0),           // 6: maps.fleetengine.delivery.v1.Task.TaskOutcome
	(*LocationInfo)(nil),            // 7: maps.fleetengine.delivery.v1.LocationInfo
	(*TimeWindow)(nil),              // 8: maps.fleetengine.delivery.v1.TimeWindow
	(*TaskAttribute)(nil),           // 9: maps.fleetengine.delivery.v1.TaskAttribute
}
var file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_depIdxs = []int32{
	1,  // 0: maps.fleetengine.delivery.v1.TaskTrackingInfo.vehicle_location:type_name -> maps.fleetengine.delivery.v1.DeliveryVehicleLocation
	2,  // 1: maps.fleetengine.delivery.v1.TaskTrackingInfo.route_polyline_points:type_name -> google.type.LatLng
	3,  // 2: maps.fleetengine.delivery.v1.TaskTrackingInfo.remaining_stop_count:type_name -> google.protobuf.Int32Value
	3,  // 3: maps.fleetengine.delivery.v1.TaskTrackingInfo.remaining_driving_distance_meters:type_name -> google.protobuf.Int32Value
	4,  // 4: maps.fleetengine.delivery.v1.TaskTrackingInfo.estimated_arrival_time:type_name -> google.protobuf.Timestamp
	4,  // 5: maps.fleetengine.delivery.v1.TaskTrackingInfo.estimated_task_completion_time:type_name -> google.protobuf.Timestamp
	5,  // 6: maps.fleetengine.delivery.v1.TaskTrackingInfo.state:type_name -> maps.fleetengine.delivery.v1.Task.State
	6,  // 7: maps.fleetengine.delivery.v1.TaskTrackingInfo.task_outcome:type_name -> maps.fleetengine.delivery.v1.Task.TaskOutcome
	4,  // 8: maps.fleetengine.delivery.v1.TaskTrackingInfo.task_outcome_time:type_name -> google.protobuf.Timestamp
	7,  // 9: maps.fleetengine.delivery.v1.TaskTrackingInfo.planned_location:type_name -> maps.fleetengine.delivery.v1.LocationInfo
	8,  // 10: maps.fleetengine.delivery.v1.TaskTrackingInfo.target_time_window:type_name -> maps.fleetengine.delivery.v1.TimeWindow
	9,  // 11: maps.fleetengine.delivery.v1.TaskTrackingInfo.attributes:type_name -> maps.fleetengine.delivery.v1.TaskAttribute
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_init() }
func file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_init() {
	if File_google_maps_fleetengine_delivery_v1_task_tracking_info_proto != nil {
		return
	}
	file_google_maps_fleetengine_delivery_v1_common_proto_init()
	file_google_maps_fleetengine_delivery_v1_delivery_vehicles_proto_init()
	file_google_maps_fleetengine_delivery_v1_tasks_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskTrackingInfo); i {
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
			RawDescriptor: file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_goTypes,
		DependencyIndexes: file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_depIdxs,
		MessageInfos:      file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_msgTypes,
	}.Build()
	File_google_maps_fleetengine_delivery_v1_task_tracking_info_proto = out.File
	file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_rawDesc = nil
	file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_goTypes = nil
	file_google_maps_fleetengine_delivery_v1_task_tracking_info_proto_depIdxs = nil
}
