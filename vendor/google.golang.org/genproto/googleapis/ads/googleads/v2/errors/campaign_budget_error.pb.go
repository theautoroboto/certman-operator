// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/campaign_budget_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Enum describing possible campaign budget errors.
type CampaignBudgetErrorEnum_CampaignBudgetError int32

const (
	// Enum unspecified.
	CampaignBudgetErrorEnum_UNSPECIFIED CampaignBudgetErrorEnum_CampaignBudgetError = 0
	// The received error code is not known in this version.
	CampaignBudgetErrorEnum_UNKNOWN CampaignBudgetErrorEnum_CampaignBudgetError = 1
	// The campaign budget cannot be shared.
	CampaignBudgetErrorEnum_CAMPAIGN_BUDGET_CANNOT_BE_SHARED CampaignBudgetErrorEnum_CampaignBudgetError = 17
	// The requested campaign budget no longer exists.
	CampaignBudgetErrorEnum_CAMPAIGN_BUDGET_REMOVED CampaignBudgetErrorEnum_CampaignBudgetError = 2
	// The campaign budget is associated with at least one campaign, and so the
	// campaign budget cannot be removed.
	CampaignBudgetErrorEnum_CAMPAIGN_BUDGET_IN_USE CampaignBudgetErrorEnum_CampaignBudgetError = 3
	// Customer is not whitelisted for this campaign budget period.
	CampaignBudgetErrorEnum_CAMPAIGN_BUDGET_PERIOD_NOT_AVAILABLE CampaignBudgetErrorEnum_CampaignBudgetError = 4
	// This field is not mutable on implicitly shared campaign budgets
	CampaignBudgetErrorEnum_CANNOT_MODIFY_FIELD_OF_IMPLICITLY_SHARED_CAMPAIGN_BUDGET CampaignBudgetErrorEnum_CampaignBudgetError = 6
	// Cannot change explicitly shared campaign budgets back to implicitly
	// shared ones.
	CampaignBudgetErrorEnum_CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_IMPLICITLY_SHARED CampaignBudgetErrorEnum_CampaignBudgetError = 7
	// An implicit campaign budget without a name cannot be changed to
	// explicitly shared campaign budget.
	CampaignBudgetErrorEnum_CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED_WITHOUT_NAME CampaignBudgetErrorEnum_CampaignBudgetError = 8
	// Cannot change an implicitly shared campaign budget to an explicitly
	// shared one.
	CampaignBudgetErrorEnum_CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED CampaignBudgetErrorEnum_CampaignBudgetError = 9
	// Only explicitly shared campaign budgets can be used with multiple
	// campaigns.
	CampaignBudgetErrorEnum_CANNOT_USE_IMPLICITLY_SHARED_CAMPAIGN_BUDGET_WITH_MULTIPLE_CAMPAIGNS CampaignBudgetErrorEnum_CampaignBudgetError = 10
	// A campaign budget with this name already exists.
	CampaignBudgetErrorEnum_DUPLICATE_NAME CampaignBudgetErrorEnum_CampaignBudgetError = 11
	// A money amount was not in the expected currency.
	CampaignBudgetErrorEnum_MONEY_AMOUNT_IN_WRONG_CURRENCY CampaignBudgetErrorEnum_CampaignBudgetError = 12
	// A money amount was less than the minimum CPC for currency.
	CampaignBudgetErrorEnum_MONEY_AMOUNT_LESS_THAN_CURRENCY_MINIMUM_CPC CampaignBudgetErrorEnum_CampaignBudgetError = 13
	// A money amount was greater than the maximum allowed.
	CampaignBudgetErrorEnum_MONEY_AMOUNT_TOO_LARGE CampaignBudgetErrorEnum_CampaignBudgetError = 14
	// A money amount was negative.
	CampaignBudgetErrorEnum_NEGATIVE_MONEY_AMOUNT CampaignBudgetErrorEnum_CampaignBudgetError = 15
	// A money amount was not a multiple of a minimum unit.
	CampaignBudgetErrorEnum_NON_MULTIPLE_OF_MINIMUM_CURRENCY_UNIT CampaignBudgetErrorEnum_CampaignBudgetError = 16
)

var CampaignBudgetErrorEnum_CampaignBudgetError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	17: "CAMPAIGN_BUDGET_CANNOT_BE_SHARED",
	2:  "CAMPAIGN_BUDGET_REMOVED",
	3:  "CAMPAIGN_BUDGET_IN_USE",
	4:  "CAMPAIGN_BUDGET_PERIOD_NOT_AVAILABLE",
	6:  "CANNOT_MODIFY_FIELD_OF_IMPLICITLY_SHARED_CAMPAIGN_BUDGET",
	7:  "CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_IMPLICITLY_SHARED",
	8:  "CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED_WITHOUT_NAME",
	9:  "CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED",
	10: "CANNOT_USE_IMPLICITLY_SHARED_CAMPAIGN_BUDGET_WITH_MULTIPLE_CAMPAIGNS",
	11: "DUPLICATE_NAME",
	12: "MONEY_AMOUNT_IN_WRONG_CURRENCY",
	13: "MONEY_AMOUNT_LESS_THAN_CURRENCY_MINIMUM_CPC",
	14: "MONEY_AMOUNT_TOO_LARGE",
	15: "NEGATIVE_MONEY_AMOUNT",
	16: "NON_MULTIPLE_OF_MINIMUM_CURRENCY_UNIT",
}

var CampaignBudgetErrorEnum_CampaignBudgetError_value = map[string]int32{
	"UNSPECIFIED":                          0,
	"UNKNOWN":                              1,
	"CAMPAIGN_BUDGET_CANNOT_BE_SHARED":     17,
	"CAMPAIGN_BUDGET_REMOVED":              2,
	"CAMPAIGN_BUDGET_IN_USE":               3,
	"CAMPAIGN_BUDGET_PERIOD_NOT_AVAILABLE": 4,
	"CANNOT_MODIFY_FIELD_OF_IMPLICITLY_SHARED_CAMPAIGN_BUDGET":             6,
	"CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_IMPLICITLY_SHARED":                   7,
	"CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED_WITHOUT_NAME":      8,
	"CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED":                   9,
	"CANNOT_USE_IMPLICITLY_SHARED_CAMPAIGN_BUDGET_WITH_MULTIPLE_CAMPAIGNS": 10,
	"DUPLICATE_NAME":                              11,
	"MONEY_AMOUNT_IN_WRONG_CURRENCY":              12,
	"MONEY_AMOUNT_LESS_THAN_CURRENCY_MINIMUM_CPC": 13,
	"MONEY_AMOUNT_TOO_LARGE":                      14,
	"NEGATIVE_MONEY_AMOUNT":                       15,
	"NON_MULTIPLE_OF_MINIMUM_CURRENCY_UNIT":       16,
}

func (x CampaignBudgetErrorEnum_CampaignBudgetError) String() string {
	return proto.EnumName(CampaignBudgetErrorEnum_CampaignBudgetError_name, int32(x))
}

func (CampaignBudgetErrorEnum_CampaignBudgetError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7e40ba8196c3aebd, []int{0, 0}
}

// Container for enum describing possible campaign budget errors.
type CampaignBudgetErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignBudgetErrorEnum) Reset()         { *m = CampaignBudgetErrorEnum{} }
func (m *CampaignBudgetErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignBudgetErrorEnum) ProtoMessage()    {}
func (*CampaignBudgetErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e40ba8196c3aebd, []int{0}
}

func (m *CampaignBudgetErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignBudgetErrorEnum.Unmarshal(m, b)
}
func (m *CampaignBudgetErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignBudgetErrorEnum.Marshal(b, m, deterministic)
}
func (m *CampaignBudgetErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignBudgetErrorEnum.Merge(m, src)
}
func (m *CampaignBudgetErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignBudgetErrorEnum.Size(m)
}
func (m *CampaignBudgetErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignBudgetErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignBudgetErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.CampaignBudgetErrorEnum_CampaignBudgetError", CampaignBudgetErrorEnum_CampaignBudgetError_name, CampaignBudgetErrorEnum_CampaignBudgetError_value)
	proto.RegisterType((*CampaignBudgetErrorEnum)(nil), "google.ads.googleads.v2.errors.CampaignBudgetErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/campaign_budget_error.proto", fileDescriptor_7e40ba8196c3aebd)
}

var fileDescriptor_7e40ba8196c3aebd = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6a, 0xd4, 0x40,
	0x14, 0x86, 0xed, 0xd6, 0xb6, 0x3a, 0xd5, 0x76, 0x1c, 0xd1, 0x6a, 0x95, 0x22, 0x4b, 0x05, 0x45,
	0x48, 0x20, 0x82, 0x48, 0x14, 0x64, 0x92, 0xcc, 0x66, 0x07, 0x93, 0x99, 0x90, 0xcc, 0xa4, 0xae,
	0x2c, 0x0c, 0x69, 0xb3, 0x84, 0x85, 0x36, 0x59, 0x36, 0xdb, 0x3e, 0x90, 0x97, 0x3e, 0x87, 0x57,
	0x3e, 0x85, 0xd7, 0xde, 0xfa, 0x02, 0x92, 0x64, 0x77, 0x6b, 0xd3, 0x6a, 0xbd, 0xca, 0x61, 0xce,
	0xff, 0xfd, 0xff, 0x09, 0x33, 0x07, 0x98, 0x59, 0x51, 0x64, 0xc7, 0x23, 0x3d, 0x49, 0x4b, 0xbd,
	0x29, 0xab, 0xea, 0xcc, 0xd0, 0x47, 0xd3, 0x69, 0x31, 0x2d, 0xf5, 0xa3, 0xe4, 0x64, 0x92, 0x8c,
	0xb3, 0x5c, 0x1d, 0x9e, 0xa6, 0xd9, 0x68, 0xa6, 0xea, 0x63, 0x6d, 0x32, 0x2d, 0x66, 0x05, 0xda,
	0x6b, 0x00, 0x2d, 0x49, 0x4b, 0x6d, 0xc9, 0x6a, 0x67, 0x86, 0xd6, 0xb0, 0xbb, 0x4f, 0x17, 0xde,
	0x93, 0xb1, 0x9e, 0xe4, 0x79, 0x31, 0x4b, 0x66, 0xe3, 0x22, 0x2f, 0x1b, 0xba, 0xfb, 0x63, 0x0d,
	0xec, 0xd8, 0x73, 0x77, 0xab, 0x36, 0x27, 0x15, 0x46, 0xf2, 0xd3, 0x93, 0xee, 0xb7, 0x35, 0x70,
	0xff, 0x8a, 0x1e, 0xda, 0x06, 0x9b, 0x92, 0x45, 0x01, 0xb1, 0x69, 0x8f, 0x12, 0x07, 0xde, 0x40,
	0x9b, 0x60, 0x43, 0xb2, 0x8f, 0x8c, 0x1f, 0x30, 0xb8, 0x82, 0xf6, 0xc1, 0x33, 0x1b, 0xfb, 0x01,
	0xa6, 0x2e, 0x53, 0x96, 0x74, 0x5c, 0x22, 0x94, 0x8d, 0x19, 0xe3, 0x42, 0x59, 0x44, 0x45, 0x7d,
	0x1c, 0x12, 0x07, 0xde, 0x43, 0x4f, 0xc0, 0x4e, 0x5b, 0x15, 0x12, 0x9f, 0xc7, 0xc4, 0x81, 0x1d,
	0xb4, 0x0b, 0x1e, 0xb6, 0x9b, 0x94, 0x29, 0x19, 0x11, 0xb8, 0x8a, 0x5e, 0x80, 0xfd, 0x76, 0x2f,
	0x20, 0x21, 0xe5, 0x8e, 0xaa, 0x22, 0x70, 0x8c, 0xa9, 0x87, 0x2d, 0x8f, 0xc0, 0x9b, 0xe8, 0x3d,
	0x78, 0x3b, 0x0f, 0xf6, 0xb9, 0x43, 0x7b, 0x03, 0xd5, 0xa3, 0xc4, 0x73, 0x14, 0xef, 0x29, 0xea,
	0x07, 0x1e, 0xb5, 0xa9, 0xf0, 0x06, 0xf3, 0x81, 0x54, 0xcb, 0x12, 0xae, 0xa3, 0x37, 0xc0, 0x98,
	0xd3, 0x32, 0x70, 0xb0, 0x20, 0x6d, 0x89, 0x12, 0xfc, 0xb2, 0x0f, 0xdc, 0x40, 0x36, 0xf8, 0x70,
	0x2d, 0x47, 0x3e, 0xb5, 0xf3, 0x0f, 0xa8, 0xe8, 0x73, 0x29, 0x14, 0xc3, 0x3e, 0x81, 0xb7, 0xfe,
	0x2b, 0xfc, 0x92, 0x09, 0xbc, 0x8d, 0xfa, 0xc0, 0x59, 0x70, 0x11, 0xb9, 0xfe, 0x37, 0xeb, 0x58,
	0xe5, 0x4b, 0x4f, 0xd0, 0xc0, 0x3b, 0x0f, 0x89, 0x20, 0x40, 0x08, 0x6c, 0x39, 0xb2, 0xc2, 0xab,
	0xf4, 0x7a, 0xaa, 0x4d, 0xd4, 0x05, 0x7b, 0x3e, 0x67, 0x64, 0xa0, 0xb0, 0xcf, 0x25, 0xab, 0xef,
	0xe4, 0x20, 0xe4, 0xcc, 0x55, 0xb6, 0x0c, 0x43, 0xc2, 0xec, 0x01, 0xbc, 0x83, 0x74, 0xf0, 0xea,
	0x82, 0xc6, 0x23, 0x51, 0xa4, 0x44, 0x1f, 0xb3, 0xa5, 0x48, 0xf9, 0x94, 0x51, 0x5f, 0xfa, 0xca,
	0x0e, 0x6c, 0x78, 0xb7, 0xba, 0xeb, 0x0b, 0x80, 0xe0, 0x5c, 0x79, 0x38, 0x74, 0x09, 0xdc, 0x42,
	0x8f, 0xc1, 0x03, 0x46, 0x5c, 0x2c, 0x68, 0x4c, 0xd4, 0x9f, 0x22, 0xb8, 0x8d, 0x5e, 0x82, 0xe7,
	0x8c, 0xb3, 0xf3, 0xd9, 0x79, 0xef, 0xdc, 0x77, 0x11, 0x24, 0x19, 0x15, 0x10, 0x5a, 0xbf, 0x56,
	0x40, 0xf7, 0xa8, 0x38, 0xd1, 0xfe, 0xbd, 0x27, 0xd6, 0xa3, 0x2b, 0x9e, 0x7a, 0x50, 0xed, 0x48,
	0xb0, 0xf2, 0xd9, 0x99, 0xb3, 0x59, 0x71, 0x9c, 0xe4, 0x99, 0x56, 0x4c, 0x33, 0x3d, 0x1b, 0xe5,
	0xf5, 0x06, 0x2d, 0xf6, 0x75, 0x32, 0x2e, 0xff, 0xb6, 0xbe, 0xef, 0x9a, 0xcf, 0x97, 0xce, 0xaa,
	0x8b, 0xf1, 0xd7, 0xce, 0x9e, 0xdb, 0x98, 0xe1, 0xb4, 0xd4, 0x9a, 0xb2, 0xaa, 0x62, 0x43, 0xab,
	0x23, 0xcb, 0xef, 0x0b, 0xc1, 0x10, 0xa7, 0xe5, 0x70, 0x29, 0x18, 0xc6, 0xc6, 0xb0, 0x11, 0xfc,
	0xec, 0x74, 0x9b, 0x53, 0xd3, 0xc4, 0x69, 0x69, 0x9a, 0x4b, 0x89, 0x69, 0xc6, 0x86, 0x69, 0x36,
	0xa2, 0xc3, 0xf5, 0x7a, 0xba, 0xd7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x58, 0x39, 0x9a,
	0x5b, 0x04, 0x00, 0x00,
}