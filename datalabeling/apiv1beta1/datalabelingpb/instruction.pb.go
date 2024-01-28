// Copyright 2019 Google LLC.
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
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.2
// source: google/cloud/datalabeling/v1beta1/instruction.proto

package datalabelingpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Instruction of how to perform the labeling task for human operators.
// Currently only PDF instruction is supported.
type Instruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Instruction resource name, format:
	// projects/{project_id}/instructions/{instruction_id}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name of the instruction. Maximum of 64 characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. User-provided description of the instruction.
	// The description can be up to 10000 characters long.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. Creation time of instruction.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Last update time of instruction.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Required. The data type of this instruction.
	DataType DataType `protobuf:"varint,6,opt,name=data_type,json=dataType,proto3,enum=google.cloud.datalabeling.v1beta1.DataType" json:"data_type,omitempty"`
	// Deprecated: this instruction format is not supported any more.
	// Instruction from a CSV file, such as for classification task.
	// The CSV file should have exact two columns, in the following format:
	//
	// * The first column is labeled data, such as an image reference, text.
	// * The second column is comma separated labels associated with data.
	//
	// Deprecated: Marked as deprecated in google/cloud/datalabeling/v1beta1/instruction.proto.
	CsvInstruction *CsvInstruction `protobuf:"bytes,7,opt,name=csv_instruction,json=csvInstruction,proto3" json:"csv_instruction,omitempty"`
	// Instruction from a PDF document. The PDF should be in a Cloud Storage
	// bucket.
	PdfInstruction *PdfInstruction `protobuf:"bytes,9,opt,name=pdf_instruction,json=pdfInstruction,proto3" json:"pdf_instruction,omitempty"`
	// Output only. The names of any related resources that are blocking changes
	// to the instruction.
	BlockingResources []string `protobuf:"bytes,10,rep,name=blocking_resources,json=blockingResources,proto3" json:"blocking_resources,omitempty"`
}

func (x *Instruction) Reset() {
	*x = Instruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datalabeling_v1beta1_instruction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instruction) ProtoMessage() {}

func (x *Instruction) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datalabeling_v1beta1_instruction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instruction.ProtoReflect.Descriptor instead.
func (*Instruction) Descriptor() ([]byte, []int) {
	return file_google_cloud_datalabeling_v1beta1_instruction_proto_rawDescGZIP(), []int{0}
}

func (x *Instruction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Instruction) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Instruction) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Instruction) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Instruction) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Instruction) GetDataType() DataType {
	if x != nil {
		return x.DataType
	}
	return DataType_DATA_TYPE_UNSPECIFIED
}

// Deprecated: Marked as deprecated in google/cloud/datalabeling/v1beta1/instruction.proto.
func (x *Instruction) GetCsvInstruction() *CsvInstruction {
	if x != nil {
		return x.CsvInstruction
	}
	return nil
}

func (x *Instruction) GetPdfInstruction() *PdfInstruction {
	if x != nil {
		return x.PdfInstruction
	}
	return nil
}

func (x *Instruction) GetBlockingResources() []string {
	if x != nil {
		return x.BlockingResources
	}
	return nil
}

// Deprecated: this instruction format is not supported any more.
// Instruction from a CSV file.
type CsvInstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CSV file for the instruction. Only gcs path is allowed.
	GcsFileUri string `protobuf:"bytes,1,opt,name=gcs_file_uri,json=gcsFileUri,proto3" json:"gcs_file_uri,omitempty"`
}

func (x *CsvInstruction) Reset() {
	*x = CsvInstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datalabeling_v1beta1_instruction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CsvInstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CsvInstruction) ProtoMessage() {}

func (x *CsvInstruction) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datalabeling_v1beta1_instruction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CsvInstruction.ProtoReflect.Descriptor instead.
func (*CsvInstruction) Descriptor() ([]byte, []int) {
	return file_google_cloud_datalabeling_v1beta1_instruction_proto_rawDescGZIP(), []int{1}
}

func (x *CsvInstruction) GetGcsFileUri() string {
	if x != nil {
		return x.GcsFileUri
	}
	return ""
}

// Instruction from a PDF file.
type PdfInstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PDF file for the instruction. Only gcs path is allowed.
	GcsFileUri string `protobuf:"bytes,1,opt,name=gcs_file_uri,json=gcsFileUri,proto3" json:"gcs_file_uri,omitempty"`
}

func (x *PdfInstruction) Reset() {
	*x = PdfInstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datalabeling_v1beta1_instruction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PdfInstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PdfInstruction) ProtoMessage() {}

func (x *PdfInstruction) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datalabeling_v1beta1_instruction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PdfInstruction.ProtoReflect.Descriptor instead.
func (*PdfInstruction) Descriptor() ([]byte, []int) {
	return file_google_cloud_datalabeling_v1beta1_instruction_proto_rawDescGZIP(), []int{2}
}

func (x *PdfInstruction) GetGcsFileUri() string {
	if x != nil {
		return x.GcsFileUri
	}
	return ""
}

var File_google_cloud_datalabeling_v1beta1_instruction_proto protoreflect.FileDescriptor

var file_google_cloud_datalabeling_v1beta1_instruction_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x04, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x5e, 0x0a, 0x0f, 0x63, 0x73, 0x76, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x43, 0x73, 0x76, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x0e, 0x63, 0x73, 0x76, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x5a, 0x0a, 0x0f, 0x70, 0x64, 0x66, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x50, 0x64, 0x66, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e,
	0x70, 0x64, 0x66, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d,
	0x0a, 0x12, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3a, 0x5b, 0xea,
	0x41, 0x58, 0x0a, 0x27, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f,
	0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x22, 0x32, 0x0a, 0x0e, 0x43, 0x73,
	0x76, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0c,
	0x67, 0x63, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x67, 0x63, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x69, 0x22, 0x32,
	0x0a, 0x0e, 0x50, 0x64, 0x66, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x0c, 0x67, 0x63, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x63, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x55,
	0x72, 0x69, 0x42, 0xe3, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x49,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x70, 0x62, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x21,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x61, 0x74,
	0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_datalabeling_v1beta1_instruction_proto_rawDescOnce sync.Once
	file_google_cloud_datalabeling_v1beta1_instruction_proto_rawDescData = file_google_cloud_datalabeling_v1beta1_instruction_proto_rawDesc
)

func file_google_cloud_datalabeling_v1beta1_instruction_proto_rawDescGZIP() []byte {
	file_google_cloud_datalabeling_v1beta1_instruction_proto_rawDescOnce.Do(func() {
		file_google_cloud_datalabeling_v1beta1_instruction_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_datalabeling_v1beta1_instruction_proto_rawDescData)
	})
	return file_google_cloud_datalabeling_v1beta1_instruction_proto_rawDescData
}

var file_google_cloud_datalabeling_v1beta1_instruction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_datalabeling_v1beta1_instruction_proto_goTypes = []interface{}{
	(*Instruction)(nil),           // 0: google.cloud.datalabeling.v1beta1.Instruction
	(*CsvInstruction)(nil),        // 1: google.cloud.datalabeling.v1beta1.CsvInstruction
	(*PdfInstruction)(nil),        // 2: google.cloud.datalabeling.v1beta1.PdfInstruction
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(DataType)(0),                 // 4: google.cloud.datalabeling.v1beta1.DataType
}
var file_google_cloud_datalabeling_v1beta1_instruction_proto_depIdxs = []int32{
	3, // 0: google.cloud.datalabeling.v1beta1.Instruction.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: google.cloud.datalabeling.v1beta1.Instruction.update_time:type_name -> google.protobuf.Timestamp
	4, // 2: google.cloud.datalabeling.v1beta1.Instruction.data_type:type_name -> google.cloud.datalabeling.v1beta1.DataType
	1, // 3: google.cloud.datalabeling.v1beta1.Instruction.csv_instruction:type_name -> google.cloud.datalabeling.v1beta1.CsvInstruction
	2, // 4: google.cloud.datalabeling.v1beta1.Instruction.pdf_instruction:type_name -> google.cloud.datalabeling.v1beta1.PdfInstruction
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_datalabeling_v1beta1_instruction_proto_init() }
func file_google_cloud_datalabeling_v1beta1_instruction_proto_init() {
	if File_google_cloud_datalabeling_v1beta1_instruction_proto != nil {
		return
	}
	file_google_cloud_datalabeling_v1beta1_dataset_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_datalabeling_v1beta1_instruction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instruction); i {
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
		file_google_cloud_datalabeling_v1beta1_instruction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CsvInstruction); i {
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
		file_google_cloud_datalabeling_v1beta1_instruction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PdfInstruction); i {
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
			RawDescriptor: file_google_cloud_datalabeling_v1beta1_instruction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_datalabeling_v1beta1_instruction_proto_goTypes,
		DependencyIndexes: file_google_cloud_datalabeling_v1beta1_instruction_proto_depIdxs,
		MessageInfos:      file_google_cloud_datalabeling_v1beta1_instruction_proto_msgTypes,
	}.Build()
	File_google_cloud_datalabeling_v1beta1_instruction_proto = out.File
	file_google_cloud_datalabeling_v1beta1_instruction_proto_rawDesc = nil
	file_google_cloud_datalabeling_v1beta1_instruction_proto_goTypes = nil
	file_google_cloud_datalabeling_v1beta1_instruction_proto_depIdxs = nil
}
