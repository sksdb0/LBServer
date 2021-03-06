// Code generated by protoc-gen-go.
// source: LeBangErrandsClassification.proto
// DO NOT EDIT!

package lebangproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MainErrandsClassification struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *MainErrandsClassification) Reset()                    { *m = MainErrandsClassification{} }
func (m *MainErrandsClassification) String() string            { return proto.CompactTextString(m) }
func (*MainErrandsClassification) ProtoMessage()               {}
func (*MainErrandsClassification) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *MainErrandsClassification) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ErrandsClassification struct {
	Classification string `protobuf:"bytes,1,opt,name=classification" json:"classification,omitempty"`
	Labels         string `protobuf:"bytes,2,opt,name=labels" json:"labels,omitempty"`
	Hint           string `protobuf:"bytes,3,opt,name=hint" json:"hint,omitempty"`
}

func (m *ErrandsClassification) Reset()                    { *m = ErrandsClassification{} }
func (m *ErrandsClassification) String() string            { return proto.CompactTextString(m) }
func (*ErrandsClassification) ProtoMessage()               {}
func (*ErrandsClassification) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ErrandsClassification) GetClassification() string {
	if m != nil {
		return m.Classification
	}
	return ""
}

func (m *ErrandsClassification) GetLabels() string {
	if m != nil {
		return m.Labels
	}
	return ""
}

func (m *ErrandsClassification) GetHint() string {
	if m != nil {
		return m.Hint
	}
	return ""
}

type GetErrandsClassificationReq struct {
	Typename string `protobuf:"bytes,1,opt,name=typename" json:"typename,omitempty"`
}

func (m *GetErrandsClassificationReq) Reset()                    { *m = GetErrandsClassificationReq{} }
func (m *GetErrandsClassificationReq) String() string            { return proto.CompactTextString(m) }
func (*GetErrandsClassificationReq) ProtoMessage()               {}
func (*GetErrandsClassificationReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *GetErrandsClassificationReq) GetTypename() string {
	if m != nil {
		return m.Typename
	}
	return ""
}

type GetErrandsClassificationRes struct {
	Classification *ErrandsClassification `protobuf:"bytes,1,opt,name=classification" json:"classification,omitempty"`
	Errorcode      string                 `protobuf:"bytes,2,opt,name=errorcode" json:"errorcode,omitempty"`
}

func (m *GetErrandsClassificationRes) Reset()                    { *m = GetErrandsClassificationRes{} }
func (m *GetErrandsClassificationRes) String() string            { return proto.CompactTextString(m) }
func (*GetErrandsClassificationRes) ProtoMessage()               {}
func (*GetErrandsClassificationRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *GetErrandsClassificationRes) GetClassification() *ErrandsClassification {
	if m != nil {
		return m.Classification
	}
	return nil
}

func (m *GetErrandsClassificationRes) GetErrorcode() string {
	if m != nil {
		return m.Errorcode
	}
	return ""
}

func init() {
	proto.RegisterType((*MainErrandsClassification)(nil), "lebangproto.MainErrandsClassification")
	proto.RegisterType((*ErrandsClassification)(nil), "lebangproto.ErrandsClassification")
	proto.RegisterType((*GetErrandsClassificationReq)(nil), "lebangproto.GetErrandsClassificationReq")
	proto.RegisterType((*GetErrandsClassificationRes)(nil), "lebangproto.GetErrandsClassificationRes")
}

func init() { proto.RegisterFile("LeBangErrandsClassification.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0xf4, 0x49, 0x75, 0x4a,
	0xcc, 0x4b, 0x77, 0x2d, 0x2a, 0x4a, 0xcc, 0x4b, 0x29, 0x76, 0xce, 0x49, 0x2c, 0x2e, 0xce, 0x4c,
	0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce,
	0x49, 0x4d, 0x4a, 0xcc, 0x4b, 0x07, 0x73, 0x94, 0xf4, 0xb9, 0x24, 0x7d, 0x13, 0x33, 0xf3, 0xb0,
	0xaa, 0x17, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x02, 0xb3, 0x95, 0xb2, 0xb9, 0x44, 0xb1, 0x2b, 0x56, 0xe3, 0xe2, 0x4b, 0x46, 0x11, 0x81, 0x6a,
	0x43, 0x13, 0x15, 0x12, 0xe3, 0x62, 0xcb, 0x49, 0x4c, 0x4a, 0xcd, 0x29, 0x96, 0x60, 0x02, 0xcb,
	0x43, 0x79, 0x20, 0xcb, 0x32, 0x32, 0xf3, 0x4a, 0x24, 0x98, 0x21, 0x96, 0x81, 0xd8, 0x4a, 0x96,
	0x5c, 0xd2, 0xee, 0xa9, 0x25, 0x58, 0xed, 0x0b, 0x4a, 0x2d, 0x14, 0x92, 0xe2, 0xe2, 0x28, 0xa9,
	0x2c, 0x48, 0x45, 0x72, 0x23, 0x9c, 0xaf, 0xd4, 0xce, 0x88, 0x4f, 0x6f, 0xb1, 0x90, 0x17, 0x56,
	0xe7, 0x72, 0x1b, 0x29, 0xe9, 0x21, 0x05, 0x8f, 0x1e, 0x76, 0xed, 0xe8, 0x5e, 0x92, 0xe1, 0xe2,
	0x4c, 0x2d, 0x2a, 0xca, 0x2f, 0x4a, 0xce, 0x4f, 0x49, 0x85, 0xfa, 0x0a, 0x21, 0x90, 0xc4, 0x06,
	0x36, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xeb, 0xd6, 0x36, 0x9b, 0x01, 0x00, 0x00,
}
