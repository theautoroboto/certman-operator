// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/detail_placement_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A view with metrics aggregated by ad group and URL or YouTube video.
type DetailPlacementView struct {
	// The resource name of the detail placement view.
	// Detail placement view resource names have the form:
	//
	//
	// `customers/{customer_id}/detailPlacementViews/{ad_group_id}~{base64_placement}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The automatic placement string at detail level, e. g. website URL, mobile
	// application ID, or a YouTube video ID.
	Placement *wrappers.StringValue `protobuf:"bytes,2,opt,name=placement,proto3" json:"placement,omitempty"`
	// The display name is URL name for websites, YouTube video name for YouTube
	// videos, and translated mobile app name for mobile apps.
	DisplayName *wrappers.StringValue `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// URL of the group placement, e.g. domain, link to the mobile application in
	// app store, or a YouTube channel URL.
	GroupPlacementTargetUrl *wrappers.StringValue `protobuf:"bytes,4,opt,name=group_placement_target_url,json=groupPlacementTargetUrl,proto3" json:"group_placement_target_url,omitempty"`
	// URL of the placement, e.g. website, link to the mobile application in app
	// store, or a YouTube video URL.
	TargetUrl *wrappers.StringValue `protobuf:"bytes,5,opt,name=target_url,json=targetUrl,proto3" json:"target_url,omitempty"`
	// Type of the placement, e.g. Website, YouTube Video, and Mobile Application.
	PlacementType        enums.PlacementTypeEnum_PlacementType `protobuf:"varint,6,opt,name=placement_type,json=placementType,proto3,enum=google.ads.googleads.v1.enums.PlacementTypeEnum_PlacementType" json:"placement_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *DetailPlacementView) Reset()         { *m = DetailPlacementView{} }
func (m *DetailPlacementView) String() string { return proto.CompactTextString(m) }
func (*DetailPlacementView) ProtoMessage()    {}
func (*DetailPlacementView) Descriptor() ([]byte, []int) {
	return fileDescriptor_9703f91bccf33f03, []int{0}
}

func (m *DetailPlacementView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailPlacementView.Unmarshal(m, b)
}
func (m *DetailPlacementView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailPlacementView.Marshal(b, m, deterministic)
}
func (m *DetailPlacementView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailPlacementView.Merge(m, src)
}
func (m *DetailPlacementView) XXX_Size() int {
	return xxx_messageInfo_DetailPlacementView.Size(m)
}
func (m *DetailPlacementView) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailPlacementView.DiscardUnknown(m)
}

var xxx_messageInfo_DetailPlacementView proto.InternalMessageInfo

func (m *DetailPlacementView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *DetailPlacementView) GetPlacement() *wrappers.StringValue {
	if m != nil {
		return m.Placement
	}
	return nil
}

func (m *DetailPlacementView) GetDisplayName() *wrappers.StringValue {
	if m != nil {
		return m.DisplayName
	}
	return nil
}

func (m *DetailPlacementView) GetGroupPlacementTargetUrl() *wrappers.StringValue {
	if m != nil {
		return m.GroupPlacementTargetUrl
	}
	return nil
}

func (m *DetailPlacementView) GetTargetUrl() *wrappers.StringValue {
	if m != nil {
		return m.TargetUrl
	}
	return nil
}

func (m *DetailPlacementView) GetPlacementType() enums.PlacementTypeEnum_PlacementType {
	if m != nil {
		return m.PlacementType
	}
	return enums.PlacementTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*DetailPlacementView)(nil), "google.ads.googleads.v1.resources.DetailPlacementView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/detail_placement_view.proto", fileDescriptor_9703f91bccf33f03)
}

var fileDescriptor_9703f91bccf33f03 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x49, 0xaa, 0x0b, 0x3b, 0xfb, 0xe7, 0x22, 0x5e, 0x18, 0xca, 0x22, 0x5d, 0x65, 0xa1,
	0x57, 0x13, 0x52, 0xef, 0x66, 0x51, 0xc9, 0xa2, 0x2c, 0x78, 0x21, 0xa5, 0x6a, 0x40, 0x29, 0x84,
	0xd9, 0xe6, 0x38, 0x04, 0x92, 0x99, 0x61, 0x66, 0xd2, 0xd2, 0x07, 0xf0, 0x45, 0xbc, 0xf4, 0x01,
	0x7c, 0x08, 0x1f, 0xc5, 0xa7, 0x90, 0x4e, 0x32, 0xd3, 0x2e, 0x58, 0xb7, 0x77, 0x27, 0x73, 0xbe,
	0xdf, 0x77, 0xce, 0xc9, 0x39, 0xe8, 0x15, 0x13, 0x82, 0xd5, 0x90, 0xd0, 0x52, 0x27, 0x5d, 0xb8,
	0x89, 0x96, 0x69, 0xa2, 0x40, 0x8b, 0x56, 0x2d, 0x40, 0x27, 0x25, 0x18, 0x5a, 0xd5, 0x85, 0xac,
	0xe9, 0x02, 0x1a, 0xe0, 0xa6, 0x58, 0x56, 0xb0, 0xc2, 0x52, 0x09, 0x23, 0xa2, 0xcb, 0x8e, 0xc1,
	0xb4, 0xd4, 0xd8, 0xe3, 0x78, 0x99, 0x62, 0x8f, 0x0f, 0x27, 0xfb, 0x2a, 0x00, 0x6f, 0x1b, 0x9d,
	0x6c, 0x6d, 0xcd, 0x5a, 0x42, 0x67, 0x3b, 0xbc, 0x70, 0x8c, 0xac, 0x12, 0xca, 0xb9, 0x30, 0xd4,
	0x54, 0x82, 0xeb, 0x3e, 0xfb, 0xac, 0xcf, 0xda, 0xaf, 0xbb, 0xf6, 0x5b, 0xb2, 0x52, 0x54, 0x4a,
	0x50, 0x7d, 0xfe, 0xf9, 0xaf, 0x01, 0x7a, 0xf2, 0xd6, 0x36, 0x3d, 0x75, 0xe6, 0x79, 0x05, 0xab,
	0xe8, 0x05, 0x3a, 0x73, 0x6d, 0x15, 0x9c, 0x36, 0x10, 0x07, 0xa3, 0x60, 0x7c, 0x3c, 0x3b, 0x75,
	0x8f, 0x1f, 0x68, 0x03, 0x11, 0x41, 0xc7, 0xbe, 0xa5, 0x38, 0x1c, 0x05, 0xe3, 0x93, 0xc9, 0x45,
	0x3f, 0x1a, 0x76, 0x05, 0xf1, 0x47, 0xa3, 0x2a, 0xce, 0x72, 0x5a, 0xb7, 0x30, 0xdb, 0xca, 0xa3,
	0x37, 0xe8, 0xb4, 0xac, 0xb4, 0xac, 0xe9, 0xba, 0xf3, 0x1f, 0x1c, 0x80, 0x9f, 0xf4, 0x84, 0x2d,
	0xfe, 0x05, 0x0d, 0x99, 0x12, 0xad, 0xdc, 0xf9, 0xd9, 0x86, 0x2a, 0x06, 0xa6, 0x68, 0x55, 0x1d,
	0x3f, 0x3a, 0xc0, 0xee, 0xa9, 0xe5, 0xfd, 0xdc, 0x9f, 0x2c, 0xfd, 0x59, 0xd5, 0xd1, 0x35, 0x42,
	0x3b, 0x56, 0x8f, 0x0f, 0x19, 0xcc, 0x78, 0x18, 0xd0, 0xf9, 0xfd, 0x3d, 0xc5, 0x47, 0xa3, 0x60,
	0x7c, 0x3e, 0x79, 0x8d, 0xf7, 0xed, 0xdf, 0x2e, 0x17, 0x6f, 0xfb, 0x58, 0x4b, 0x78, 0xc7, 0xdb,
	0xe6, 0xfe, 0xcb, 0xec, 0x4c, 0xee, 0x7e, 0xde, 0x7c, 0x0f, 0xd1, 0xd5, 0x42, 0x34, 0xf8, 0xc1,
	0xa3, 0xba, 0x89, 0xff, 0xb1, 0xdf, 0xe9, 0x66, 0x88, 0x69, 0xf0, 0xf5, 0x7d, 0x8f, 0x33, 0x51,
	0x53, 0xce, 0xb0, 0x50, 0x2c, 0x61, 0xc0, 0xed, 0x88, 0xee, 0x00, 0x65, 0xa5, 0xff, 0x73, 0xf1,
	0xd7, 0x3e, 0xfa, 0x11, 0x0e, 0x6e, 0xb3, 0xec, 0x67, 0x78, 0x79, 0xdb, 0x59, 0x66, 0xa5, 0xc6,
	0x5d, 0xb8, 0x89, 0xf2, 0x14, 0xcf, 0x9c, 0xf2, 0xb7, 0xd3, 0xcc, 0xb3, 0x52, 0xcf, 0xbd, 0x66,
	0x9e, 0xa7, 0x73, 0xaf, 0xf9, 0x13, 0x5e, 0x75, 0x09, 0x42, 0xb2, 0x52, 0x13, 0xe2, 0x55, 0x84,
	0xe4, 0x29, 0x21, 0x5e, 0x77, 0x77, 0x64, 0x9b, 0x7d, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x4c,
	0x9e, 0x22, 0xe9, 0x9d, 0x03, 0x00, 0x00,
}