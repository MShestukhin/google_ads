// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: services/google_ads/pb/google_ads/google_ads.proto

package google_ads

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SyncAdGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdAccountId string `protobuf:"bytes,1,opt,name=ad_account_id,json=adAccountId,proto3" json:"ad_account_id,omitempty"`
	BrandId     int32  `protobuf:"varint,2,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
}

func (x *SyncAdGroupRequest) Reset() {
	*x = SyncAdGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncAdGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncAdGroupRequest) ProtoMessage() {}

func (x *SyncAdGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncAdGroupRequest.ProtoReflect.Descriptor instead.
func (*SyncAdGroupRequest) Descriptor() ([]byte, []int) {
	return file_services_google_ads_pb_google_ads_google_ads_proto_rawDescGZIP(), []int{0}
}

func (x *SyncAdGroupRequest) GetAdAccountId() string {
	if x != nil {
		return x.AdAccountId
	}
	return ""
}

func (x *SyncAdGroupRequest) GetBrandId() int32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

type SyncCampaignCriteriaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdAccountId string `protobuf:"bytes,1,opt,name=ad_account_id,json=adAccountId,proto3" json:"ad_account_id,omitempty"`
	BrandId     int32  `protobuf:"varint,2,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	CampaignId  int64  `protobuf:"varint,3,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
}

func (x *SyncCampaignCriteriaRequest) Reset() {
	*x = SyncCampaignCriteriaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncCampaignCriteriaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncCampaignCriteriaRequest) ProtoMessage() {}

func (x *SyncCampaignCriteriaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncCampaignCriteriaRequest.ProtoReflect.Descriptor instead.
func (*SyncCampaignCriteriaRequest) Descriptor() ([]byte, []int) {
	return file_services_google_ads_pb_google_ads_google_ads_proto_rawDescGZIP(), []int{1}
}

func (x *SyncCampaignCriteriaRequest) GetAdAccountId() string {
	if x != nil {
		return x.AdAccountId
	}
	return ""
}

func (x *SyncCampaignCriteriaRequest) GetBrandId() int32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *SyncCampaignCriteriaRequest) GetCampaignId() int64 {
	if x != nil {
		return x.CampaignId
	}
	return 0
}

type UpdateTokenForBrandIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BrandId int32 `protobuf:"varint,1,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
}

func (x *UpdateTokenForBrandIdRequest) Reset() {
	*x = UpdateTokenForBrandIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTokenForBrandIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTokenForBrandIdRequest) ProtoMessage() {}

func (x *UpdateTokenForBrandIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTokenForBrandIdRequest.ProtoReflect.Descriptor instead.
func (*UpdateTokenForBrandIdRequest) Descriptor() ([]byte, []int) {
	return file_services_google_ads_pb_google_ads_google_ads_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTokenForBrandIdRequest) GetBrandId() int32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

type UpdateTokenForBrandIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UpdateTokenForBrandIdResponse) Reset() {
	*x = UpdateTokenForBrandIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTokenForBrandIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTokenForBrandIdResponse) ProtoMessage() {}

func (x *UpdateTokenForBrandIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTokenForBrandIdResponse.ProtoReflect.Descriptor instead.
func (*UpdateTokenForBrandIdResponse) Descriptor() ([]byte, []int) {
	return file_services_google_ads_pb_google_ads_google_ads_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTokenForBrandIdResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type SyncCampaignsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdAccountId  string `protobuf:"bytes,1,opt,name=ad_account_id,json=adAccountId,proto3" json:"ad_account_id,omitempty"`
	BrandId      int32  `protobuf:"varint,2,opt,name=brandId,proto3" json:"brandId,omitempty"`
	IsIcremental bool   `protobuf:"varint,3,opt,name=is_icremental,json=isIcremental,proto3" json:"is_icremental,omitempty"`
}

func (x *SyncCampaignsRequest) Reset() {
	*x = SyncCampaignsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncCampaignsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncCampaignsRequest) ProtoMessage() {}

func (x *SyncCampaignsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncCampaignsRequest.ProtoReflect.Descriptor instead.
func (*SyncCampaignsRequest) Descriptor() ([]byte, []int) {
	return file_services_google_ads_pb_google_ads_google_ads_proto_rawDescGZIP(), []int{4}
}

func (x *SyncCampaignsRequest) GetAdAccountId() string {
	if x != nil {
		return x.AdAccountId
	}
	return ""
}

func (x *SyncCampaignsRequest) GetBrandId() int32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *SyncCampaignsRequest) GetIsIcremental() bool {
	if x != nil {
		return x.IsIcremental
	}
	return false
}

type TaskIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TaskIdentifier) Reset() {
	*x = TaskIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskIdentifier) ProtoMessage() {}

func (x *TaskIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskIdentifier.ProtoReflect.Descriptor instead.
func (*TaskIdentifier) Descriptor() ([]byte, []int) {
	return file_services_google_ads_pb_google_ads_google_ads_proto_rawDescGZIP(), []int{5}
}

func (x *TaskIdentifier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_services_google_ads_pb_google_ads_google_ads_proto protoreflect.FileDescriptor

var file_services_google_ads_pb_google_ads_google_ads_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5f, 0x61, 0x64, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x61, 0x69, 0x2e, 0x73, 0x6f, 0x6d, 0x69, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64,
	0x73, 0x22, 0x53, 0x0a, 0x12, 0x53, 0x79, 0x6e, 0x63, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x64, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x7d, 0x0a, 0x1b, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x64,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64,
	0x22, 0x39, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x46,
	0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x79, 0x0a, 0x14, 0x53,
	0x79, 0x6e, 0x63, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x64, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x69, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x49, 0x63, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x22, 0x20, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x89, 0x04, 0x0a, 0x09, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x12, 0x71, 0x0a, 0x0d, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x73, 0x12, 0x31, 0x2e, 0x61, 0x69, 0x2e, 0x73, 0x6f, 0x6d,
	0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5f, 0x61, 0x64, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x69, 0x2e,
	0x73, 0x6f, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x00, 0x12, 0x90, 0x01, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x49, 0x64, 0x12, 0x39, 0x2e, 0x61, 0x69, 0x2e, 0x73, 0x6f, 0x6d, 0x69, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x6f, 0x72,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a,
	0x2e, 0x61, 0x69, 0x2e, 0x73, 0x6f, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x14,
	0x53, 0x79, 0x6e, 0x63, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x12, 0x38, 0x2e, 0x61, 0x69, 0x2e, 0x73, 0x6f, 0x6d, 0x69, 0x6e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61,
	0x64, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x61, 0x69, 0x2e, 0x73, 0x6f, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x00, 0x12, 0x75, 0x0a,
	0x13, 0x53, 0x79, 0x6e, 0x63, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x12, 0x2f, 0x2e, 0x61, 0x69, 0x2e, 0x73, 0x6f, 0x6d, 0x69, 0x6e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61,
	0x64, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x69, 0x2e, 0x73, 0x6f, 0x6d, 0x69, 0x6e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f,
	0x61, 0x64, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x22, 0x00, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x73,
	0x6f, 0x6d, 0x69, 0x6e, 0x2e, 0x61, 0x69, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x2f, 0x70,
	0x62, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_google_ads_pb_google_ads_google_ads_proto_rawDescOnce sync.Once
	file_services_google_ads_pb_google_ads_google_ads_proto_rawDescData = file_services_google_ads_pb_google_ads_google_ads_proto_rawDesc
)

func file_services_google_ads_pb_google_ads_google_ads_proto_rawDescGZIP() []byte {
	file_services_google_ads_pb_google_ads_google_ads_proto_rawDescOnce.Do(func() {
		file_services_google_ads_pb_google_ads_google_ads_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_google_ads_pb_google_ads_google_ads_proto_rawDescData)
	})
	return file_services_google_ads_pb_google_ads_google_ads_proto_rawDescData
}

var file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_services_google_ads_pb_google_ads_google_ads_proto_goTypes = []interface{}{
	(*SyncAdGroupRequest)(nil),            // 0: ai.somin.service.google_ads.SyncAdGroupRequest
	(*SyncCampaignCriteriaRequest)(nil),   // 1: ai.somin.service.google_ads.SyncCampaignCriteriaRequest
	(*UpdateTokenForBrandIdRequest)(nil),  // 2: ai.somin.service.google_ads.UpdateTokenForBrandIdRequest
	(*UpdateTokenForBrandIdResponse)(nil), // 3: ai.somin.service.google_ads.UpdateTokenForBrandIdResponse
	(*SyncCampaignsRequest)(nil),          // 4: ai.somin.service.google_ads.SyncCampaignsRequest
	(*TaskIdentifier)(nil),                // 5: ai.somin.service.google_ads.TaskIdentifier
}
var file_services_google_ads_pb_google_ads_google_ads_proto_depIdxs = []int32{
	4, // 0: ai.somin.service.google_ads.GoogleAds.SyncCampaigns:input_type -> ai.somin.service.google_ads.SyncCampaignsRequest
	2, // 1: ai.somin.service.google_ads.GoogleAds.UpdateTokenForBrandId:input_type -> ai.somin.service.google_ads.UpdateTokenForBrandIdRequest
	1, // 2: ai.somin.service.google_ads.GoogleAds.SyncCampaignCriteria:input_type -> ai.somin.service.google_ads.SyncCampaignCriteriaRequest
	0, // 3: ai.somin.service.google_ads.GoogleAds.SyncAdGroupCriteria:input_type -> ai.somin.service.google_ads.SyncAdGroupRequest
	5, // 4: ai.somin.service.google_ads.GoogleAds.SyncCampaigns:output_type -> ai.somin.service.google_ads.TaskIdentifier
	3, // 5: ai.somin.service.google_ads.GoogleAds.UpdateTokenForBrandId:output_type -> ai.somin.service.google_ads.UpdateTokenForBrandIdResponse
	5, // 6: ai.somin.service.google_ads.GoogleAds.SyncCampaignCriteria:output_type -> ai.somin.service.google_ads.TaskIdentifier
	5, // 7: ai.somin.service.google_ads.GoogleAds.SyncAdGroupCriteria:output_type -> ai.somin.service.google_ads.TaskIdentifier
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_google_ads_pb_google_ads_google_ads_proto_init() }
func file_services_google_ads_pb_google_ads_google_ads_proto_init() {
	if File_services_google_ads_pb_google_ads_google_ads_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncAdGroupRequest); i {
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
		file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncCampaignCriteriaRequest); i {
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
		file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTokenForBrandIdRequest); i {
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
		file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTokenForBrandIdResponse); i {
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
		file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncCampaignsRequest); i {
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
		file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskIdentifier); i {
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
			RawDescriptor: file_services_google_ads_pb_google_ads_google_ads_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_google_ads_pb_google_ads_google_ads_proto_goTypes,
		DependencyIndexes: file_services_google_ads_pb_google_ads_google_ads_proto_depIdxs,
		MessageInfos:      file_services_google_ads_pb_google_ads_google_ads_proto_msgTypes,
	}.Build()
	File_services_google_ads_pb_google_ads_google_ads_proto = out.File
	file_services_google_ads_pb_google_ads_google_ads_proto_rawDesc = nil
	file_services_google_ads_pb_google_ads_google_ads_proto_goTypes = nil
	file_services_google_ads_pb_google_ads_google_ads_proto_depIdxs = nil
}
