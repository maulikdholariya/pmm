// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent/agent.proto

package agent

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Type represents Agent type.
type Type int32

const (
	Type_TYPE_INVALID    Type = 0
	Type_NODE_EXPORTER   Type = 1
	Type_MYSQLD_EXPORTER Type = 2
)

var Type_name = map[int32]string{
	0: "TYPE_INVALID",
	1: "NODE_EXPORTER",
	2: "MYSQLD_EXPORTER",
}
var Type_value = map[string]int32{
	"TYPE_INVALID":    0,
	"NODE_EXPORTER":   1,
	"MYSQLD_EXPORTER": 2,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_agent_bdab11a77c62c240, []int{0}
}

// QANDataRequest is an AgentMessage for sending QAN data.
type QANDataRequest struct {
	Data                 *any.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QANDataRequest) Reset()         { *m = QANDataRequest{} }
func (m *QANDataRequest) String() string { return proto.CompactTextString(m) }
func (*QANDataRequest) ProtoMessage()    {}
func (*QANDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_bdab11a77c62c240, []int{0}
}
func (m *QANDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QANDataRequest.Unmarshal(m, b)
}
func (m *QANDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QANDataRequest.Marshal(b, m, deterministic)
}
func (dst *QANDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QANDataRequest.Merge(dst, src)
}
func (m *QANDataRequest) XXX_Size() int {
	return xxx_messageInfo_QANDataRequest.Size(m)
}
func (m *QANDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QANDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QANDataRequest proto.InternalMessageInfo

func (m *QANDataRequest) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

// QANDataResponse is a ServerMessage for QAN data acceptance.
type QANDataResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QANDataResponse) Reset()         { *m = QANDataResponse{} }
func (m *QANDataResponse) String() string { return proto.CompactTextString(m) }
func (*QANDataResponse) ProtoMessage()    {}
func (*QANDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_bdab11a77c62c240, []int{1}
}
func (m *QANDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QANDataResponse.Unmarshal(m, b)
}
func (m *QANDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QANDataResponse.Marshal(b, m, deterministic)
}
func (dst *QANDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QANDataResponse.Merge(dst, src)
}
func (m *QANDataResponse) XXX_Size() int {
	return xxx_messageInfo_QANDataResponse.Size(m)
}
func (m *QANDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QANDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QANDataResponse proto.InternalMessageInfo

// PingRequest is a ServerMessage for checking connectivity, latency and clock drift.
type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_bdab11a77c62c240, []int{2}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (dst *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(dst, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

// PingResponse is an AgentMessage with current time for measuring clock drift.
type PingResponse struct {
	CurrentTime          *timestamp.Timestamp `protobuf:"bytes,1,opt,name=current_time,json=currentTime,proto3" json:"current_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_bdab11a77c62c240, []int{3}
}
func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (dst *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(dst, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetCurrentTime() *timestamp.Timestamp {
	if m != nil {
		return m.CurrentTime
	}
	return nil
}

// SetStateRequest is a ServerMessage asking pmm-agent to run agents according to desired state.
type SetStateRequest struct {
	AgentProcesses       []*SetStateRequest_AgentProcess `protobuf:"bytes,1,rep,name=agent_processes,json=agentProcesses,proto3" json:"agent_processes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *SetStateRequest) Reset()         { *m = SetStateRequest{} }
func (m *SetStateRequest) String() string { return proto.CompactTextString(m) }
func (*SetStateRequest) ProtoMessage()    {}
func (*SetStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_bdab11a77c62c240, []int{4}
}
func (m *SetStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetStateRequest.Unmarshal(m, b)
}
func (m *SetStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetStateRequest.Marshal(b, m, deterministic)
}
func (dst *SetStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetStateRequest.Merge(dst, src)
}
func (m *SetStateRequest) XXX_Size() int {
	return xxx_messageInfo_SetStateRequest.Size(m)
}
func (m *SetStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetStateRequest proto.InternalMessageInfo

func (m *SetStateRequest) GetAgentProcesses() []*SetStateRequest_AgentProcess {
	if m != nil {
		return m.AgentProcesses
	}
	return nil
}

// AgentProcess describes desired configuration of a single agent.
type SetStateRequest_AgentProcess struct {
	AgentId              uint32            `protobuf:"varint,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Type                 Type              `protobuf:"varint,2,opt,name=type,proto3,enum=agent.Type" json:"type,omitempty"`
	Args                 []string          `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	Env                  []string          `protobuf:"bytes,4,rep,name=env,proto3" json:"env,omitempty"`
	Configs              map[string]string `protobuf:"bytes,5,rep,name=configs,proto3" json:"configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SetStateRequest_AgentProcess) Reset()         { *m = SetStateRequest_AgentProcess{} }
func (m *SetStateRequest_AgentProcess) String() string { return proto.CompactTextString(m) }
func (*SetStateRequest_AgentProcess) ProtoMessage()    {}
func (*SetStateRequest_AgentProcess) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_bdab11a77c62c240, []int{4, 0}
}
func (m *SetStateRequest_AgentProcess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetStateRequest_AgentProcess.Unmarshal(m, b)
}
func (m *SetStateRequest_AgentProcess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetStateRequest_AgentProcess.Marshal(b, m, deterministic)
}
func (dst *SetStateRequest_AgentProcess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetStateRequest_AgentProcess.Merge(dst, src)
}
func (m *SetStateRequest_AgentProcess) XXX_Size() int {
	return xxx_messageInfo_SetStateRequest_AgentProcess.Size(m)
}
func (m *SetStateRequest_AgentProcess) XXX_DiscardUnknown() {
	xxx_messageInfo_SetStateRequest_AgentProcess.DiscardUnknown(m)
}

var xxx_messageInfo_SetStateRequest_AgentProcess proto.InternalMessageInfo

func (m *SetStateRequest_AgentProcess) GetAgentId() uint32 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *SetStateRequest_AgentProcess) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_TYPE_INVALID
}

func (m *SetStateRequest_AgentProcess) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *SetStateRequest_AgentProcess) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *SetStateRequest_AgentProcess) GetConfigs() map[string]string {
	if m != nil {
		return m.Configs
	}
	return nil
}

// SetStateResponse is an AgentMessage describing agent statuses.
type SetStateResponse struct {
	AgentProcesses       []*SetStateResponse_AgentProcess `protobuf:"bytes,1,rep,name=agent_processes,json=agentProcesses,proto3" json:"agent_processes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *SetStateResponse) Reset()         { *m = SetStateResponse{} }
func (m *SetStateResponse) String() string { return proto.CompactTextString(m) }
func (*SetStateResponse) ProtoMessage()    {}
func (*SetStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_bdab11a77c62c240, []int{5}
}
func (m *SetStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetStateResponse.Unmarshal(m, b)
}
func (m *SetStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetStateResponse.Marshal(b, m, deterministic)
}
func (dst *SetStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetStateResponse.Merge(dst, src)
}
func (m *SetStateResponse) XXX_Size() int {
	return xxx_messageInfo_SetStateResponse.Size(m)
}
func (m *SetStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetStateResponse proto.InternalMessageInfo

func (m *SetStateResponse) GetAgentProcesses() []*SetStateResponse_AgentProcess {
	if m != nil {
		return m.AgentProcesses
	}
	return nil
}

// AgentProcess describes status of a single agent.
type SetStateResponse_AgentProcess struct {
	AgentId              uint32   `protobuf:"varint,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Disabled             bool     `protobuf:"varint,2,opt,name=disabled,proto3" json:"disabled,omitempty"`
	ListenPort           uint32   `protobuf:"varint,3,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetStateResponse_AgentProcess) Reset()         { *m = SetStateResponse_AgentProcess{} }
func (m *SetStateResponse_AgentProcess) String() string { return proto.CompactTextString(m) }
func (*SetStateResponse_AgentProcess) ProtoMessage()    {}
func (*SetStateResponse_AgentProcess) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_bdab11a77c62c240, []int{5, 0}
}
func (m *SetStateResponse_AgentProcess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetStateResponse_AgentProcess.Unmarshal(m, b)
}
func (m *SetStateResponse_AgentProcess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetStateResponse_AgentProcess.Marshal(b, m, deterministic)
}
func (dst *SetStateResponse_AgentProcess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetStateResponse_AgentProcess.Merge(dst, src)
}
func (m *SetStateResponse_AgentProcess) XXX_Size() int {
	return xxx_messageInfo_SetStateResponse_AgentProcess.Size(m)
}
func (m *SetStateResponse_AgentProcess) XXX_DiscardUnknown() {
	xxx_messageInfo_SetStateResponse_AgentProcess.DiscardUnknown(m)
}

var xxx_messageInfo_SetStateResponse_AgentProcess proto.InternalMessageInfo

func (m *SetStateResponse_AgentProcess) GetAgentId() uint32 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *SetStateResponse_AgentProcess) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *SetStateResponse_AgentProcess) GetListenPort() uint32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

type AgentMessage struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*AgentMessage_QanData
	//	*AgentMessage_Ping
	//	*AgentMessage_State
	Payload              isAgentMessage_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AgentMessage) Reset()         { *m = AgentMessage{} }
func (m *AgentMessage) String() string { return proto.CompactTextString(m) }
func (*AgentMessage) ProtoMessage()    {}
func (*AgentMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_bdab11a77c62c240, []int{6}
}
func (m *AgentMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentMessage.Unmarshal(m, b)
}
func (m *AgentMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentMessage.Marshal(b, m, deterministic)
}
func (dst *AgentMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentMessage.Merge(dst, src)
}
func (m *AgentMessage) XXX_Size() int {
	return xxx_messageInfo_AgentMessage.Size(m)
}
func (m *AgentMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AgentMessage proto.InternalMessageInfo

func (m *AgentMessage) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type isAgentMessage_Payload interface {
	isAgentMessage_Payload()
}

type AgentMessage_QanData struct {
	QanData *QANDataRequest `protobuf:"bytes,2,opt,name=qan_data,json=qanData,proto3,oneof"`
}

type AgentMessage_Ping struct {
	Ping *PingResponse `protobuf:"bytes,8,opt,name=ping,proto3,oneof"`
}

type AgentMessage_State struct {
	State *SetStateResponse `protobuf:"bytes,9,opt,name=state,proto3,oneof"`
}

func (*AgentMessage_QanData) isAgentMessage_Payload() {}

func (*AgentMessage_Ping) isAgentMessage_Payload() {}

func (*AgentMessage_State) isAgentMessage_Payload() {}

func (m *AgentMessage) GetPayload() isAgentMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *AgentMessage) GetQanData() *QANDataRequest {
	if x, ok := m.GetPayload().(*AgentMessage_QanData); ok {
		return x.QanData
	}
	return nil
}

func (m *AgentMessage) GetPing() *PingResponse {
	if x, ok := m.GetPayload().(*AgentMessage_Ping); ok {
		return x.Ping
	}
	return nil
}

func (m *AgentMessage) GetState() *SetStateResponse {
	if x, ok := m.GetPayload().(*AgentMessage_State); ok {
		return x.State
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AgentMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AgentMessage_OneofMarshaler, _AgentMessage_OneofUnmarshaler, _AgentMessage_OneofSizer, []interface{}{
		(*AgentMessage_QanData)(nil),
		(*AgentMessage_Ping)(nil),
		(*AgentMessage_State)(nil),
	}
}

func _AgentMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AgentMessage)
	// payload
	switch x := m.Payload.(type) {
	case *AgentMessage_QanData:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.QanData); err != nil {
			return err
		}
	case *AgentMessage_Ping:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ping); err != nil {
			return err
		}
	case *AgentMessage_State:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.State); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AgentMessage.Payload has unexpected type %T", x)
	}
	return nil
}

func _AgentMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AgentMessage)
	switch tag {
	case 2: // payload.qan_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(QANDataRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &AgentMessage_QanData{msg}
		return true, err
	case 8: // payload.ping
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PingResponse)
		err := b.DecodeMessage(msg)
		m.Payload = &AgentMessage_Ping{msg}
		return true, err
	case 9: // payload.state
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SetStateResponse)
		err := b.DecodeMessage(msg)
		m.Payload = &AgentMessage_State{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AgentMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AgentMessage)
	// payload
	switch x := m.Payload.(type) {
	case *AgentMessage_QanData:
		s := proto.Size(x.QanData)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AgentMessage_Ping:
		s := proto.Size(x.Ping)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AgentMessage_State:
		s := proto.Size(x.State)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ServerMessage struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*ServerMessage_QanData
	//	*ServerMessage_Ping
	//	*ServerMessage_State
	Payload              isServerMessage_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ServerMessage) Reset()         { *m = ServerMessage{} }
func (m *ServerMessage) String() string { return proto.CompactTextString(m) }
func (*ServerMessage) ProtoMessage()    {}
func (*ServerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_bdab11a77c62c240, []int{7}
}
func (m *ServerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerMessage.Unmarshal(m, b)
}
func (m *ServerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerMessage.Marshal(b, m, deterministic)
}
func (dst *ServerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerMessage.Merge(dst, src)
}
func (m *ServerMessage) XXX_Size() int {
	return xxx_messageInfo_ServerMessage.Size(m)
}
func (m *ServerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ServerMessage proto.InternalMessageInfo

func (m *ServerMessage) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type isServerMessage_Payload interface {
	isServerMessage_Payload()
}

type ServerMessage_QanData struct {
	QanData *QANDataResponse `protobuf:"bytes,2,opt,name=qan_data,json=qanData,proto3,oneof"`
}

type ServerMessage_Ping struct {
	Ping *PingRequest `protobuf:"bytes,8,opt,name=ping,proto3,oneof"`
}

type ServerMessage_State struct {
	State *SetStateRequest `protobuf:"bytes,9,opt,name=state,proto3,oneof"`
}

func (*ServerMessage_QanData) isServerMessage_Payload() {}

func (*ServerMessage_Ping) isServerMessage_Payload() {}

func (*ServerMessage_State) isServerMessage_Payload() {}

func (m *ServerMessage) GetPayload() isServerMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ServerMessage) GetQanData() *QANDataResponse {
	if x, ok := m.GetPayload().(*ServerMessage_QanData); ok {
		return x.QanData
	}
	return nil
}

func (m *ServerMessage) GetPing() *PingRequest {
	if x, ok := m.GetPayload().(*ServerMessage_Ping); ok {
		return x.Ping
	}
	return nil
}

func (m *ServerMessage) GetState() *SetStateRequest {
	if x, ok := m.GetPayload().(*ServerMessage_State); ok {
		return x.State
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ServerMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ServerMessage_OneofMarshaler, _ServerMessage_OneofUnmarshaler, _ServerMessage_OneofSizer, []interface{}{
		(*ServerMessage_QanData)(nil),
		(*ServerMessage_Ping)(nil),
		(*ServerMessage_State)(nil),
	}
}

func _ServerMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ServerMessage)
	// payload
	switch x := m.Payload.(type) {
	case *ServerMessage_QanData:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.QanData); err != nil {
			return err
		}
	case *ServerMessage_Ping:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ping); err != nil {
			return err
		}
	case *ServerMessage_State:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.State); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ServerMessage.Payload has unexpected type %T", x)
	}
	return nil
}

func _ServerMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ServerMessage)
	switch tag {
	case 2: // payload.qan_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(QANDataResponse)
		err := b.DecodeMessage(msg)
		m.Payload = &ServerMessage_QanData{msg}
		return true, err
	case 8: // payload.ping
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PingRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &ServerMessage_Ping{msg}
		return true, err
	case 9: // payload.state
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SetStateRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &ServerMessage_State{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ServerMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ServerMessage)
	// payload
	switch x := m.Payload.(type) {
	case *ServerMessage_QanData:
		s := proto.Size(x.QanData)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ServerMessage_Ping:
		s := proto.Size(x.Ping)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ServerMessage_State:
		s := proto.Size(x.State)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*QANDataRequest)(nil), "agent.QANDataRequest")
	proto.RegisterType((*QANDataResponse)(nil), "agent.QANDataResponse")
	proto.RegisterType((*PingRequest)(nil), "agent.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "agent.PingResponse")
	proto.RegisterType((*SetStateRequest)(nil), "agent.SetStateRequest")
	proto.RegisterType((*SetStateRequest_AgentProcess)(nil), "agent.SetStateRequest.AgentProcess")
	proto.RegisterMapType((map[string]string)(nil), "agent.SetStateRequest.AgentProcess.ConfigsEntry")
	proto.RegisterType((*SetStateResponse)(nil), "agent.SetStateResponse")
	proto.RegisterType((*SetStateResponse_AgentProcess)(nil), "agent.SetStateResponse.AgentProcess")
	proto.RegisterType((*AgentMessage)(nil), "agent.AgentMessage")
	proto.RegisterType((*ServerMessage)(nil), "agent.ServerMessage")
	proto.RegisterEnum("agent.Type", Type_name, Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentClient interface {
	// Connect establishes two-way communication channel between pmm-agent and pmm-managed.
	Connect(ctx context.Context, opts ...grpc.CallOption) (Agent_ConnectClient, error)
}

type agentClient struct {
	cc *grpc.ClientConn
}

func NewAgentClient(cc *grpc.ClientConn) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Connect(ctx context.Context, opts ...grpc.CallOption) (Agent_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Agent_serviceDesc.Streams[0], "/agent.Agent/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentConnectClient{stream}
	return x, nil
}

type Agent_ConnectClient interface {
	Send(*AgentMessage) error
	Recv() (*ServerMessage, error)
	grpc.ClientStream
}

type agentConnectClient struct {
	grpc.ClientStream
}

func (x *agentConnectClient) Send(m *AgentMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentConnectClient) Recv() (*ServerMessage, error) {
	m := new(ServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentServer is the server API for Agent service.
type AgentServer interface {
	// Connect establishes two-way communication channel between pmm-agent and pmm-managed.
	Connect(Agent_ConnectServer) error
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).Connect(&agentConnectServer{stream})
}

type Agent_ConnectServer interface {
	Send(*ServerMessage) error
	Recv() (*AgentMessage, error)
	grpc.ServerStream
}

type agentConnectServer struct {
	grpc.ServerStream
}

func (x *agentConnectServer) Send(m *ServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentConnectServer) Recv() (*AgentMessage, error) {
	m := new(AgentMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agent.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Agent_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "agent/agent.proto",
}

func init() { proto.RegisterFile("agent/agent.proto", fileDescriptor_agent_bdab11a77c62c240) }

var fileDescriptor_agent_bdab11a77c62c240 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x4e, 0xdb, 0x40,
	0x10, 0xc6, 0x63, 0x27, 0x69, 0x92, 0xc9, 0x5f, 0x16, 0x4a, 0x8d, 0x2f, 0x44, 0x6e, 0x0f, 0x6e,
	0x0f, 0x0e, 0x0a, 0x17, 0x84, 0x54, 0xa9, 0x29, 0x89, 0x04, 0x15, 0x81, 0xb0, 0x89, 0xaa, 0x72,
	0x8a, 0x96, 0x78, 0xb1, 0xac, 0x86, 0xb5, 0xf1, 0x6e, 0x90, 0xfc, 0x68, 0x7d, 0x84, 0xbe, 0x42,
	0x9f, 0xa4, 0xc7, 0xca, 0xbb, 0x76, 0x48, 0x80, 0x54, 0x5c, 0xd0, 0xce, 0xec, 0x37, 0x33, 0x3f,
	0x7f, 0x3b, 0x04, 0xb6, 0x88, 0x47, 0x99, 0xe8, 0xc8, 0xbf, 0x4e, 0x18, 0x05, 0x22, 0x40, 0x45,
	0x19, 0x98, 0x7b, 0x5e, 0x10, 0x78, 0x73, 0xda, 0x91, 0xc9, 0x9b, 0xc5, 0x6d, 0x87, 0xb0, 0x58,
	0x29, 0xcc, 0xfd, 0xa7, 0x57, 0xc2, 0xbf, 0xa3, 0x5c, 0x90, 0xbb, 0x50, 0x09, 0xac, 0x63, 0x68,
	0x5c, 0xf5, 0x2e, 0xfa, 0x44, 0x10, 0x4c, 0xef, 0x17, 0x94, 0x0b, 0x64, 0x43, 0xc1, 0x25, 0x82,
	0x18, 0x5a, 0x5b, 0xb3, 0xab, 0xdd, 0x1d, 0x47, 0x75, 0x70, 0xb2, 0x0e, 0x4e, 0x8f, 0xc5, 0x58,
	0x2a, 0xac, 0x2d, 0x68, 0x2e, 0x6b, 0x79, 0x18, 0x30, 0x4e, 0xad, 0x3a, 0x54, 0x47, 0x3e, 0xf3,
	0xd2, 0x5e, 0xd6, 0x10, 0x6a, 0x2a, 0x54, 0xd7, 0xe8, 0x33, 0xd4, 0x66, 0x8b, 0x28, 0xa2, 0x4c,
	0x4c, 0x13, 0x90, 0x74, 0x86, 0xf9, 0x6c, 0xc6, 0x24, 0xa3, 0xc4, 0xd5, 0x54, 0x9f, 0x64, 0xac,
	0x3f, 0x3a, 0x34, 0xc7, 0x54, 0x8c, 0x05, 0x11, 0x34, 0xc3, 0x3d, 0x87, 0xa6, 0x74, 0x61, 0x1a,
	0x46, 0xc1, 0x8c, 0x72, 0x4e, 0xb9, 0xa1, 0xb5, 0xf3, 0x76, 0xb5, 0xfb, 0xde, 0x51, 0x56, 0x3d,
	0x29, 0x70, 0x7a, 0x49, 0x76, 0xa4, 0xc4, 0xb8, 0x41, 0x56, 0x22, 0xca, 0xcd, 0xbf, 0x1a, 0xd4,
	0x56, 0x05, 0x68, 0x0f, 0xca, 0xaa, 0xbd, 0xef, 0x4a, 0xda, 0x3a, 0x2e, 0xc9, 0xf8, 0xcc, 0x45,
	0xfb, 0x50, 0x10, 0x71, 0x48, 0x0d, 0xbd, 0xad, 0xd9, 0x8d, 0x6e, 0x35, 0x1d, 0x37, 0x89, 0x43,
	0x8a, 0xe5, 0x05, 0x42, 0x50, 0x20, 0x91, 0xc7, 0x8d, 0x7c, 0x3b, 0x6f, 0x57, 0xb0, 0x3c, 0xa3,
	0x16, 0xe4, 0x29, 0x7b, 0x30, 0x0a, 0x32, 0x95, 0x1c, 0xd1, 0x37, 0x28, 0xcd, 0x02, 0x76, 0xeb,
	0x7b, 0xdc, 0x28, 0x4a, 0xf0, 0x83, 0x57, 0x80, 0x3b, 0x27, 0xaa, 0x64, 0xc0, 0x44, 0x14, 0xe3,
	0xac, 0x81, 0x79, 0x0c, 0xb5, 0xd5, 0x8b, 0x64, 0xda, 0x4f, 0x1a, 0x4b, 0xf0, 0x0a, 0x4e, 0x8e,
	0x68, 0x07, 0x8a, 0x0f, 0x64, 0xbe, 0x50, 0xd4, 0x15, 0xac, 0x82, 0x63, 0xfd, 0x48, 0xb3, 0x7e,
	0x6b, 0xd0, 0x7a, 0x1c, 0x99, 0x3e, 0xd8, 0x70, 0x93, 0xbb, 0x1f, 0x9e, 0x41, 0xaa, 0x8a, 0xff,
	0xdb, 0x7b, 0xfb, 0x7a, 0x77, 0x4d, 0x28, 0xbb, 0x3e, 0x27, 0x37, 0x73, 0xea, 0x4a, 0xd6, 0x32,
	0x5e, 0xc6, 0x68, 0x1f, 0xaa, 0x73, 0x9f, 0x0b, 0xca, 0xa6, 0x61, 0x10, 0x09, 0x23, 0x2f, 0x2b,
	0x41, 0xa5, 0x46, 0x41, 0x24, 0xac, 0x5f, 0xd9, 0x33, 0x0e, 0x29, 0xe7, 0xc4, 0xa3, 0xa8, 0x01,
	0xfa, 0x72, 0x84, 0xee, 0xbb, 0xa8, 0x0b, 0xe5, 0x7b, 0xc2, 0xa6, 0x72, 0xd1, 0x75, 0xb9, 0x84,
	0x6f, 0xd3, 0x0f, 0x5a, 0xff, 0x6f, 0x38, 0xcd, 0xe1, 0xd2, 0x3d, 0x61, 0x49, 0x06, 0x7d, 0x84,
	0x42, 0xe8, 0x33, 0xcf, 0x28, 0x4b, 0xfd, 0x76, 0xaa, 0x5f, 0xdd, 0xef, 0xd3, 0x1c, 0x96, 0x12,
	0xd4, 0x81, 0x22, 0x4f, 0x5c, 0x31, 0x2a, 0x52, 0xfb, 0x6e, 0x83, 0x59, 0xa7, 0x39, 0xac, 0x74,
	0x5f, 0x2b, 0x50, 0x0a, 0x49, 0x3c, 0x0f, 0x88, 0x9b, 0xb0, 0xd7, 0xc7, 0x34, 0x7a, 0xa0, 0xd1,
	0x26, 0xf8, 0xc3, 0x67, 0xf0, 0xbb, 0x4f, 0xe1, 0x97, 0xfd, 0x97, 0xf4, 0xf6, 0x1a, 0x3d, 0x5a,
	0xa3, 0xcf, 0x3e, 0x55, 0xc1, 0x3b, 0xeb, 0xf0, 0xbb, 0x2f, 0xaf, 0xe3, 0x4b, 0xec, 0x9f, 0xbe,
	0x40, 0x21, 0xd9, 0x7f, 0xd4, 0x82, 0xda, 0xe4, 0x7a, 0x34, 0x98, 0x9e, 0x5d, 0x7c, 0xef, 0x9d,
	0x9f, 0xf5, 0x5b, 0x39, 0xb4, 0x05, 0xf5, 0x8b, 0xcb, 0xfe, 0x60, 0x3a, 0xf8, 0x31, 0xba, 0xc4,
	0x93, 0x01, 0x6e, 0x69, 0x68, 0x1b, 0x9a, 0xc3, 0xeb, 0xf1, 0xd5, 0x79, 0xff, 0x31, 0xa9, 0x77,
	0x7b, 0x50, 0x94, 0x0f, 0x87, 0x8e, 0xa0, 0x74, 0x12, 0x30, 0x46, 0x67, 0x02, 0x65, 0x56, 0xaf,
	0xbe, 0xa8, 0xb9, 0xb3, 0xc4, 0x5a, 0xb1, 0xca, 0xd6, 0x0e, 0xb4, 0x9b, 0x37, 0xf2, 0x67, 0xe4,
	0xf0, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xc0, 0x71, 0xe1, 0x31, 0x05, 0x00, 0x00,
}
