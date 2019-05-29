// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: log.proto

/*
	Package sls is a generated protocol buffer package.

	It is generated from these files:
		log.proto

	It has these top-level messages:
		LogContent
		Log
		LogTag
		LogGroup
		SlsLogPackage
		SlsLogPackageList
		LogGroupList
*/
package hook

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LogContent struct {
	Key              *string `protobuf:"bytes,1,req,name=Key" json:"Key,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=Value" json:"Value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LogContent) Reset()                    { *m = LogContent{} }
func (m *LogContent) String() string            { return proto.CompactTextString(m) }
func (*LogContent) ProtoMessage()               {}
func (*LogContent) Descriptor() ([]byte, []int) { return fileDescriptorLog, []int{0} }

func (m *LogContent) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *LogContent) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Log struct {
	Time             *uint32       `protobuf:"varint,1,req,name=Time" json:"Time,omitempty"`
	Contents         []*LogContent `protobuf:"bytes,2,rep,name=Contents" json:"Contents,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptorLog, []int{1} }

func (m *Log) GetTime() uint32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Log) GetContents() []*LogContent {
	if m != nil {
		return m.Contents
	}
	return nil
}

type LogTag struct {
	Key              *string `protobuf:"bytes,1,req,name=Key" json:"Key,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=Value" json:"Value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LogTag) Reset()                    { *m = LogTag{} }
func (m *LogTag) String() string            { return proto.CompactTextString(m) }
func (*LogTag) ProtoMessage()               {}
func (*LogTag) Descriptor() ([]byte, []int) { return fileDescriptorLog, []int{2} }

func (m *LogTag) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *LogTag) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type LogGroup struct {
	Logs             []*Log    `protobuf:"bytes,1,rep,name=Logs" json:"Logs,omitempty"`
	Category         *string   `protobuf:"bytes,2,opt,name=Category" json:"Category,omitempty"`
	Topic            *string   `protobuf:"bytes,3,opt,name=Topic" json:"Topic,omitempty"`
	Source           *string   `protobuf:"bytes,4,opt,name=Source" json:"Source,omitempty"`
	MachineUUID      *string   `protobuf:"bytes,5,opt,name=MachineUUID" json:"MachineUUID,omitempty"`
	LogTags          []*LogTag `protobuf:"bytes,6,rep,name=LogTags" json:"LogTags,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *LogGroup) Reset()                    { *m = LogGroup{} }
func (m *LogGroup) String() string            { return proto.CompactTextString(m) }
func (*LogGroup) ProtoMessage()               {}
func (*LogGroup) Descriptor() ([]byte, []int) { return fileDescriptorLog, []int{3} }

func (m *LogGroup) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *LogGroup) GetCategory() string {
	if m != nil && m.Category != nil {
		return *m.Category
	}
	return ""
}

func (m *LogGroup) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *LogGroup) GetSource() string {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return ""
}

func (m *LogGroup) GetMachineUUID() string {
	if m != nil && m.MachineUUID != nil {
		return *m.MachineUUID
	}
	return ""
}

func (m *LogGroup) GetLogTags() []*LogTag {
	if m != nil {
		return m.LogTags
	}
	return nil
}

type SlsLogPackage struct {
	Data             []byte `protobuf:"bytes,1,req,name=data" json:"data,omitempty"`
	UncompressSize   *int32 `protobuf:"varint,2,opt,name=uncompress_size,json=uncompressSize" json:"uncompress_size,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SlsLogPackage) Reset()                    { *m = SlsLogPackage{} }
func (m *SlsLogPackage) String() string            { return proto.CompactTextString(m) }
func (*SlsLogPackage) ProtoMessage()               {}
func (*SlsLogPackage) Descriptor() ([]byte, []int) { return fileDescriptorLog, []int{4} }

func (m *SlsLogPackage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SlsLogPackage) GetUncompressSize() int32 {
	if m != nil && m.UncompressSize != nil {
		return *m.UncompressSize
	}
	return 0
}

type SlsLogPackageList struct {
	Packages         []*SlsLogPackage `protobuf:"bytes,1,rep,name=packages" json:"packages,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *SlsLogPackageList) Reset()                    { *m = SlsLogPackageList{} }
func (m *SlsLogPackageList) String() string            { return proto.CompactTextString(m) }
func (*SlsLogPackageList) ProtoMessage()               {}
func (*SlsLogPackageList) Descriptor() ([]byte, []int) { return fileDescriptorLog, []int{5} }

func (m *SlsLogPackageList) GetPackages() []*SlsLogPackage {
	if m != nil {
		return m.Packages
	}
	return nil
}

type LogGroupList struct {
	LogGroups        []*LogGroup `protobuf:"bytes,1,rep,name=LogGroups" json:"LogGroups,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *LogGroupList) Reset()                    { *m = LogGroupList{} }
func (m *LogGroupList) String() string            { return proto.CompactTextString(m) }
func (*LogGroupList) ProtoMessage()               {}
func (*LogGroupList) Descriptor() ([]byte, []int) { return fileDescriptorLog, []int{6} }

func (m *LogGroupList) GetLogGroups() []*LogGroup {
	if m != nil {
		return m.LogGroups
	}
	return nil
}

func init() {
	proto.RegisterType((*LogContent)(nil), "sls.LogContent")
	proto.RegisterType((*Log)(nil), "sls.Log")
	proto.RegisterType((*LogTag)(nil), "sls.LogTag")
	proto.RegisterType((*LogGroup)(nil), "sls.LogGroup")
	proto.RegisterType((*SlsLogPackage)(nil), "sls.SlsLogPackage")
	proto.RegisterType((*SlsLogPackageList)(nil), "sls.SlsLogPackageList")
	proto.RegisterType((*LogGroupList)(nil), "sls.LogGroupList")
}
func (m *LogContent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogContent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Key == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(*m.Key)))
		i += copy(dAtA[i:], *m.Key)
	}
	if m.Value == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(*m.Value)))
		i += copy(dAtA[i:], *m.Value)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Log) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Log) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Time == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLog(dAtA, i, uint64(*m.Time))
	}
	if len(m.Contents) > 0 {
		for _, msg := range m.Contents {
			dAtA[i] = 0x12
			i++
			i = encodeVarintLog(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *LogTag) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogTag) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Key == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(*m.Key)))
		i += copy(dAtA[i:], *m.Key)
	}
	if m.Value == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(*m.Value)))
		i += copy(dAtA[i:], *m.Value)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *LogGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogGroup) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Logs) > 0 {
		for _, msg := range m.Logs {
			dAtA[i] = 0xa
			i++
			i = encodeVarintLog(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Category != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(*m.Category)))
		i += copy(dAtA[i:], *m.Category)
	}
	if m.Topic != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(*m.Topic)))
		i += copy(dAtA[i:], *m.Topic)
	}
	if m.Source != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(*m.Source)))
		i += copy(dAtA[i:], *m.Source)
	}
	if m.MachineUUID != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(*m.MachineUUID)))
		i += copy(dAtA[i:], *m.MachineUUID)
	}
	if len(m.LogTags) > 0 {
		for _, msg := range m.LogTags {
			dAtA[i] = 0x32
			i++
			i = encodeVarintLog(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SlsLogPackage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SlsLogPackage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.UncompressSize != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLog(dAtA, i, uint64(*m.UncompressSize))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SlsLogPackageList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SlsLogPackageList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Packages) > 0 {
		for _, msg := range m.Packages {
			dAtA[i] = 0xa
			i++
			i = encodeVarintLog(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *LogGroupList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogGroupList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.LogGroups) > 0 {
		for _, msg := range m.LogGroups {
			dAtA[i] = 0xa
			i++
			i = encodeVarintLog(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintLog(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LogContent) Size() (n int) {
	var l int
	_ = l
	if m.Key != nil {
		l = len(*m.Key)
		n += 1 + l + sovLog(uint64(l))
	}
	if m.Value != nil {
		l = len(*m.Value)
		n += 1 + l + sovLog(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Log) Size() (n int) {
	var l int
	_ = l
	if m.Time != nil {
		n += 1 + sovLog(uint64(*m.Time))
	}
	if len(m.Contents) > 0 {
		for _, e := range m.Contents {
			l = e.Size()
			n += 1 + l + sovLog(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LogTag) Size() (n int) {
	var l int
	_ = l
	if m.Key != nil {
		l = len(*m.Key)
		n += 1 + l + sovLog(uint64(l))
	}
	if m.Value != nil {
		l = len(*m.Value)
		n += 1 + l + sovLog(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LogGroup) Size() (n int) {
	var l int
	_ = l
	if len(m.Logs) > 0 {
		for _, e := range m.Logs {
			l = e.Size()
			n += 1 + l + sovLog(uint64(l))
		}
	}
	if m.Category != nil {
		l = len(*m.Category)
		n += 1 + l + sovLog(uint64(l))
	}
	if m.Topic != nil {
		l = len(*m.Topic)
		n += 1 + l + sovLog(uint64(l))
	}
	if m.Source != nil {
		l = len(*m.Source)
		n += 1 + l + sovLog(uint64(l))
	}
	if m.MachineUUID != nil {
		l = len(*m.MachineUUID)
		n += 1 + l + sovLog(uint64(l))
	}
	if len(m.LogTags) > 0 {
		for _, e := range m.LogTags {
			l = e.Size()
			n += 1 + l + sovLog(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SlsLogPackage) Size() (n int) {
	var l int
	_ = l
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovLog(uint64(l))
	}
	if m.UncompressSize != nil {
		n += 1 + sovLog(uint64(*m.UncompressSize))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SlsLogPackageList) Size() (n int) {
	var l int
	_ = l
	if len(m.Packages) > 0 {
		for _, e := range m.Packages {
			l = e.Size()
			n += 1 + l + sovLog(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LogGroupList) Size() (n int) {
	var l int
	_ = l
	if len(m.LogGroups) > 0 {
		for _, e := range m.LogGroups {
			l = e.Size()
			n += 1 + l + sovLog(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLog(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLog(x uint64) (n int) {
	return sovLog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LogContent) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogContent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogContent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Key = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Value = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Log) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Log: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Log: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Time = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contents = append(m.Contents, &LogContent{})
			if err := m.Contents[len(m.Contents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LogTag) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogTag: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogTag: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Key = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Value = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LogGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Logs = append(m.Logs, &Log{})
			if err := m.Logs[len(m.Logs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Category", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Category = &s
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Topic = &s
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Source = &s
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MachineUUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.MachineUUID = &s
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogTags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogTags = append(m.LogTags, &LogTag{})
			if err := m.LogTags[len(m.LogTags)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SlsLogPackage) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SlsLogPackage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SlsLogPackage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UncompressSize", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UncompressSize = &v
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SlsLogPackageList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SlsLogPackageList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SlsLogPackageList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Packages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Packages = append(m.Packages, &SlsLogPackage{})
			if err := m.Packages[len(m.Packages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LogGroupList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogGroupList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogGroupList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogGroups", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogGroups = append(m.LogGroups, &LogGroup{})
			if err := m.LogGroups[len(m.LogGroups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLog(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLog
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLog
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLog
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthLog
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLog
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLog(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLog = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLog   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("log.proto", fileDescriptorLog) }

var fileDescriptorLog = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x4d, 0xd2, 0xd6, 0xe4, 0x74, 0xeb, 0xae, 0x83, 0xc8, 0xb0, 0x48, 0x08, 0x01, 0xb1,
	0xb0, 0x98, 0x15, 0xf1, 0xce, 0x3b, 0x2b, 0x8a, 0x38, 0x82, 0x4c, 0xb3, 0xde, 0xca, 0x6c, 0x1c,
	0x67, 0x83, 0x69, 0x4e, 0xc8, 0x24, 0x17, 0xdb, 0x27, 0xf1, 0x5d, 0x7c, 0x01, 0x2f, 0x7d, 0x04,
	0xa9, 0x2f, 0x22, 0x3d, 0xe9, 0xb4, 0xf6, 0xce, 0xbb, 0xf3, 0xfb, 0xbe, 0xf3, 0x1f, 0xa2, 0x0a,
	0x4d, 0xd6, 0xb4, 0xd8, 0x21, 0x0b, 0x6c, 0x65, 0xcf, 0x9f, 0x9a, 0xb2, 0xbb, 0xe9, 0xaf, 0xb3,
	0x02, 0x57, 0x97, 0x06, 0x0d, 0x5e, 0x92, 0x77, 0xdd, 0x7f, 0x25, 0x22, 0xa0, 0x68, 0xa8, 0x49,
	0x5f, 0x00, 0x08, 0x34, 0x0b, 0xac, 0x3b, 0x5d, 0x77, 0xec, 0x0c, 0x82, 0xf7, 0xfa, 0x96, 0x7b,
	0x89, 0x3f, 0x8f, 0xe4, 0x36, 0x64, 0x0f, 0x60, 0xfc, 0x49, 0x55, 0xbd, 0xe6, 0x3e, 0x69, 0x03,
	0xa4, 0x6f, 0x20, 0x10, 0x68, 0x18, 0x83, 0x51, 0x5e, 0xae, 0x34, 0xe5, 0xcf, 0x24, 0xc5, 0xec,
	0x02, 0xc2, 0x5d, 0x37, 0xcb, 0xfd, 0x24, 0x98, 0x4f, 0x9f, 0x9f, 0x66, 0xb6, 0xb2, 0xd9, 0x61,
	0x8a, 0xdc, 0x27, 0xa4, 0xcf, 0x60, 0x22, 0xd0, 0xe4, 0xca, 0xfc, 0xf7, 0xe4, 0x1f, 0x1e, 0x84,
	0x02, 0xcd, 0xdb, 0x16, 0xfb, 0x86, 0x3d, 0x82, 0x91, 0x40, 0x63, 0xb9, 0x47, 0x73, 0x42, 0x37,
	0x47, 0x92, 0xca, 0xce, 0x21, 0x5c, 0xa8, 0x4e, 0x1b, 0x6c, 0x6f, 0xb9, 0x9f, 0x78, 0xf3, 0x48,
	0xee, 0x79, 0xdb, 0x3c, 0xc7, 0xa6, 0x2c, 0x78, 0x40, 0xc6, 0x00, 0xec, 0x21, 0x4c, 0x96, 0xd8,
	0xb7, 0x85, 0xe6, 0x23, 0x92, 0x77, 0xc4, 0x12, 0x98, 0x7e, 0x50, 0xc5, 0x4d, 0x59, 0xeb, 0xab,
	0xab, 0x77, 0xaf, 0xf9, 0x98, 0xcc, 0x7f, 0x25, 0xf6, 0x18, 0xee, 0x0e, 0x87, 0x58, 0x3e, 0xa1,
	0x65, 0xa6, 0x6e, 0x99, 0x5c, 0x19, 0xe9, 0xbc, 0x54, 0xc0, 0x6c, 0x59, 0x59, 0x81, 0xe6, 0xa3,
	0x2a, 0xbe, 0x29, 0xa3, 0xb7, 0x1f, 0xfc, 0xa2, 0x3a, 0x45, 0x77, 0x9f, 0x48, 0x8a, 0xd9, 0x13,
	0x38, 0xed, 0xeb, 0x02, 0x57, 0x4d, 0xab, 0xad, 0xfd, 0x6c, 0xcb, 0xb5, 0xa6, 0xf5, 0xc7, 0xf2,
	0xde, 0x41, 0x5e, 0x96, 0x6b, 0x9d, 0x2e, 0xe0, 0xfe, 0x51, 0x37, 0x51, 0xda, 0x8e, 0x65, 0x10,
	0x36, 0x03, 0xba, 0xbf, 0x30, 0x5a, 0xe5, 0x28, 0x53, 0xee, 0x73, 0xd2, 0x97, 0x70, 0xe2, 0xfe,
	0x49, 0xf5, 0x17, 0x10, 0x39, 0x76, 0x0d, 0x66, 0xee, 0x16, 0x52, 0xe5, 0xc1, 0x7f, 0x75, 0xf6,
	0x73, 0x13, 0x7b, 0xbf, 0x36, 0xb1, 0xf7, 0x7b, 0x13, 0x7b, 0xdf, 0xff, 0xc4, 0x77, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xfa, 0x12, 0xf5, 0x31, 0x8f, 0x02, 0x00, 0x00,
}