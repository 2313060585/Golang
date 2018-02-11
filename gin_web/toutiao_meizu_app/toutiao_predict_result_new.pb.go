// Code generated by protoc-gen-go. DO NOT EDIT.
// source: toutiao_predict_result_new.proto

package toutiao_meizu_app

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DisplayResult struct {
	Code             *int32           `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Message          *string          `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Value            *DisplayUnitList `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *DisplayResult) Reset()                    { *m = DisplayResult{} }
func (m *DisplayResult) String() string            { return proto.CompactTextString(m) }
func (*DisplayResult) ProtoMessage()               {}
func (*DisplayResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *DisplayResult) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *DisplayResult) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *DisplayResult) GetValue() *DisplayUnitList {
	if m != nil {
		return m.Value
	}
	return nil
}

type DisplayUnitList struct {
	Data             []*DisplayUnit `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	Version          *string        `protobuf:"bytes,2,req,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *DisplayUnitList) Reset()                    { *m = DisplayUnitList{} }
func (m *DisplayUnitList) String() string            { return proto.CompactTextString(m) }
func (*DisplayUnitList) ProtoMessage()               {}
func (*DisplayUnitList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *DisplayUnitList) GetData() []*DisplayUnit {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DisplayUnitList) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

type DisplayUnit struct {
	PositionId       *int32   `protobuf:"varint,1,req,name=position_id" json:"position_id,omitempty"`
	UnitId           *int32   `protobuf:"varint,2,req,name=unit_id" json:"unit_id,omitempty"`
	AppId            *uint32  `protobuf:"varint,3,req,name=app_id" json:"app_id,omitempty"`
	Ctr              *float32 `protobuf:"fixed32,4,req,name=ctr" json:"ctr,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DisplayUnit) Reset()                    { *m = DisplayUnit{} }
func (m *DisplayUnit) String() string            { return proto.CompactTextString(m) }
func (*DisplayUnit) ProtoMessage()               {}
func (*DisplayUnit) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *DisplayUnit) GetPositionId() int32 {
	if m != nil && m.PositionId != nil {
		return *m.PositionId
	}
	return 0
}

func (m *DisplayUnit) GetUnitId() int32 {
	if m != nil && m.UnitId != nil {
		return *m.UnitId
	}
	return 0
}

func (m *DisplayUnit) GetAppId() uint32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

func (m *DisplayUnit) GetCtr() float32 {
	if m != nil && m.Ctr != nil {
		return *m.Ctr
	}
	return 0
}

type SearchResult struct {
	Code             *int32          `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Message          *string         `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Value            *SearchUnitList `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *SearchResult) Reset()                    { *m = SearchResult{} }
func (m *SearchResult) String() string            { return proto.CompactTextString(m) }
func (*SearchResult) ProtoMessage()               {}
func (*SearchResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *SearchResult) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *SearchResult) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *SearchResult) GetValue() *SearchUnitList {
	if m != nil {
		return m.Value
	}
	return nil
}

type SearchUnitList struct {
	Data             []*SearchUnit `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	Version          *string       `protobuf:"bytes,2,req,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *SearchUnitList) Reset()                    { *m = SearchUnitList{} }
func (m *SearchUnitList) String() string            { return proto.CompactTextString(m) }
func (*SearchUnitList) ProtoMessage()               {}
func (*SearchUnitList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *SearchUnitList) GetData() []*SearchUnit {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SearchUnitList) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

type SearchUnit struct {
	UnitId           *int32   `protobuf:"varint,1,req,name=unit_id" json:"unit_id,omitempty"`
	AppId            *uint32  `protobuf:"varint,2,req,name=app_id" json:"app_id,omitempty"`
	Ctr              *float32 `protobuf:"fixed32,3,req,name=ctr" json:"ctr,omitempty"`
	Relevance        *float32 `protobuf:"fixed32,4,req,name=relevance" json:"relevance,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SearchUnit) Reset()                    { *m = SearchUnit{} }
func (m *SearchUnit) String() string            { return proto.CompactTextString(m) }
func (*SearchUnit) ProtoMessage()               {}
func (*SearchUnit) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *SearchUnit) GetUnitId() int32 {
	if m != nil && m.UnitId != nil {
		return *m.UnitId
	}
	return 0
}

func (m *SearchUnit) GetAppId() uint32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

func (m *SearchUnit) GetCtr() float32 {
	if m != nil && m.Ctr != nil {
		return *m.Ctr
	}
	return 0
}

func (m *SearchUnit) GetRelevance() float32 {
	if m != nil && m.Relevance != nil {
		return *m.Relevance
	}
	return 0
}

type IdeaCTRList struct {
	Data             []*IdeaCTR `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	Version          *string    `protobuf:"bytes,2,req,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *IdeaCTRList) Reset()                    { *m = IdeaCTRList{} }
func (m *IdeaCTRList) String() string            { return proto.CompactTextString(m) }
func (*IdeaCTRList) ProtoMessage()               {}
func (*IdeaCTRList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *IdeaCTRList) GetData() []*IdeaCTR {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *IdeaCTRList) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

type IdeaCTR struct {
	IdeaId           *int32   `protobuf:"varint,1,req,name=idea_id" json:"idea_id,omitempty"`
	Ctr              *float32 `protobuf:"fixed32,2,req,name=ctr" json:"ctr,omitempty"`
	Cvr              *float32 `protobuf:"fixed32,3,opt,name=cvr" json:"cvr,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *IdeaCTR) Reset()                    { *m = IdeaCTR{} }
func (m *IdeaCTR) String() string            { return proto.CompactTextString(m) }
func (*IdeaCTR) ProtoMessage()               {}
func (*IdeaCTR) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *IdeaCTR) GetIdeaId() int32 {
	if m != nil && m.IdeaId != nil {
		return *m.IdeaId
	}
	return 0
}

func (m *IdeaCTR) GetCtr() float32 {
	if m != nil && m.Ctr != nil {
		return *m.Ctr
	}
	return 0
}

func (m *IdeaCTR) GetCvr() float32 {
	if m != nil && m.Cvr != nil {
		return *m.Cvr
	}
	return 0
}

func init() {
	proto.RegisterType((*DisplayResult)(nil), "toutiao.meizu.app.DisplayResult")
	proto.RegisterType((*DisplayUnitList)(nil), "toutiao.meizu.app.DisplayUnitList")
	proto.RegisterType((*DisplayUnit)(nil), "toutiao.meizu.app.DisplayUnit")
	proto.RegisterType((*SearchResult)(nil), "toutiao.meizu.app.SearchResult")
	proto.RegisterType((*SearchUnitList)(nil), "toutiao.meizu.app.SearchUnitList")
	proto.RegisterType((*SearchUnit)(nil), "toutiao.meizu.app.SearchUnit")
	proto.RegisterType((*IdeaCTRList)(nil), "toutiao.meizu.app.IdeaCTRList")
	proto.RegisterType((*IdeaCTR)(nil), "toutiao.meizu.app.IdeaCTR")
}

func init() { proto.RegisterFile("toutiao_predict_result_new.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xdf, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0x69, 0xba, 0x7d, 0xc7, 0x6e, 0xf7, 0x83, 0xe5, 0xfb, 0x52, 0x04, 0x25, 0xf6, 0x29,
	0xa0, 0x14, 0x1d, 0xfe, 0x07, 0xfa, 0xa0, 0xa0, 0x32, 0xa6, 0x3e, 0x97, 0xd0, 0x5e, 0x34, 0xd0,
	0x35, 0x21, 0x49, 0x2b, 0xfa, 0xd7, 0x4b, 0x6b, 0x75, 0xeb, 0x36, 0x86, 0x8f, 0xf7, 0xdc, 0xc3,
	0xe1, 0x73, 0x0e, 0x30, 0xa7, 0x4a, 0x27, 0x85, 0x4a, 0xb4, 0xc1, 0x4c, 0xa6, 0x2e, 0x31, 0x68,
	0xcb, 0xdc, 0x25, 0x05, 0xbe, 0xc7, 0xda, 0x28, 0xa7, 0xe8, 0xac, 0x75, 0xc4, 0x2b, 0x94, 0x9f,
	0x65, 0x2c, 0xb4, 0x8e, 0x04, 0x8c, 0x6f, 0xa4, 0xd5, 0xb9, 0xf8, 0x58, 0x36, 0x6e, 0x3a, 0x82,
	0x5e, 0xaa, 0x32, 0x0c, 0x3d, 0x46, 0x78, 0x9f, 0x4e, 0x61, 0xb0, 0x42, 0x6b, 0xc5, 0x2b, 0x86,
	0x84, 0x79, 0x7c, 0x48, 0x2f, 0xa1, 0x5f, 0x89, 0xbc, 0xc4, 0xd0, 0x67, 0x1e, 0x0f, 0xe6, 0x51,
	0xbc, 0x13, 0x19, 0xb7, 0x79, 0x2f, 0x85, 0x74, 0xf7, 0xd2, 0xba, 0x68, 0x01, 0xd3, 0x2d, 0x89,
	0x9e, 0x43, 0x2f, 0x13, 0x4e, 0x84, 0x1e, 0xf3, 0x79, 0x30, 0x3f, 0x39, 0x1c, 0x52, 0x43, 0x54,
	0x68, 0xac, 0x54, 0x45, 0x48, 0x18, 0xe1, 0xc3, 0x68, 0x01, 0xc1, 0xe6, 0xff, 0x3f, 0x04, 0x5a,
	0x59, 0xe9, 0xa4, 0x2a, 0x12, 0x99, 0xad, 0xc9, 0xcb, 0x42, 0xba, 0x5a, 0x20, 0x8d, 0x30, 0x81,
	0x7f, 0x42, 0xeb, 0xfa, 0xf6, 0x19, 0xe1, 0x63, 0x1a, 0x80, 0x9f, 0x3a, 0x13, 0xf6, 0x18, 0xe1,
	0x24, 0x4a, 0x60, 0xf4, 0x84, 0xc2, 0xa4, 0x6f, 0x7f, 0x5b, 0xe1, 0xa2, 0xbb, 0xc2, 0xe9, 0x9e,
	0x02, 0xdf, 0x71, 0xbf, 0x23, 0x3c, 0xc2, 0xa4, 0xab, 0xd0, 0xb3, 0xce, 0x06, 0xc7, 0x07, 0x23,
	0x76, 0x27, 0x78, 0x00, 0xe8, 0xbe, 0x7f, 0xca, 0x7a, 0x5b, 0x65, 0xc9, 0x66, 0xd9, 0xba, 0x39,
	0xa1, 0x33, 0x18, 0x1a, 0xcc, 0xb1, 0x12, 0x45, 0x8a, 0x6d, 0xff, 0x5b, 0x08, 0xee, 0x32, 0x14,
	0xd7, 0xcf, 0xcb, 0x86, 0x8d, 0x77, 0xd8, 0x8e, 0xf6, 0xb0, 0xb5, 0xee, 0x5d, 0xb0, 0x2b, 0x18,
	0x6c, 0xfc, 0x64, 0x86, 0x62, 0x4d, 0xd5, 0x52, 0x90, 0x86, 0xa2, 0x3e, 0x2a, 0xd3, 0x2c, 0x48,
	0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xda, 0xed, 0x7b, 0x97, 0xbc, 0x02, 0x00, 0x00,
}