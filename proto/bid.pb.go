// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bid.proto

/*
Package sonm is a generated protocol buffer package.

It is generated from these files:
	bid.proto
	capabilities.proto
	hub.proto
	insonmnia.proto
	locator.proto
	marketplace.proto
	miner.proto
	node.proto

It has these top-level messages:
	Geo
	Resources
	Slot
	Order
	Capabilities
	CPUDevice
	RAMDevice
	GPUDevice
	ListReply
	TaskRequirements
	HubStartTaskRequest
	HubStartTaskReply
	HubStatusReply
	DealRequest
	GetDevicePropertiesReply
	SetDevicePropertiesRequest
	SlotsReply
	GetAllSlotsReply
	AddSlotRequest
	RemoveSlotRequest
	GetRegistredWorkersReply
	TaskListReply
	CPUDeviceInfo
	GPUDeviceInfo
	DevicesReply
	Empty
	ID
	PingReply
	CPUUsage
	MemoryUsage
	NetworkUsage
	ResourceUsage
	InfoReply
	TaskStatusReply
	StatusMapReply
	ContainerRestartPolicy
	TaskLogsRequest
	TaskLogsChunk
	DiscoverHubRequest
	TaskResourceRequirements
	Timestamp
	AnnounceRequest
	ResolveRequest
	ResolveReply
	GetOrdersRequest
	GetOrdersReply
	MinerHandshakeRequest
	MinerHandshakeReply
	MinerStartRequest
	MinerStartReply
	TaskInfo
	MinerStatusMapRequest
	TaskListRequest
	Deal
	DealListRequest
	DealListReply
*/
package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type OrderType int32

const (
	OrderType_ANY OrderType = 0
	OrderType_BID OrderType = 1
	OrderType_ASK OrderType = 2
)

var OrderType_name = map[int32]string{
	0: "ANY",
	1: "BID",
	2: "ASK",
}
var OrderType_value = map[string]int32{
	"ANY": 0,
	"BID": 1,
	"ASK": 2,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}
func (OrderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type NetworkType int32

const (
	NetworkType_NO_NETWORK NetworkType = 0
	NetworkType_OUTBOUND   NetworkType = 1
	NetworkType_INCOMING   NetworkType = 2
)

var NetworkType_name = map[int32]string{
	0: "NO_NETWORK",
	1: "OUTBOUND",
	2: "INCOMING",
}
var NetworkType_value = map[string]int32{
	"NO_NETWORK": 0,
	"OUTBOUND":   1,
	"INCOMING":   2,
}

func (x NetworkType) String() string {
	return proto.EnumName(NetworkType_name, int32(x))
}
func (NetworkType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GPUCount int32

const (
	GPUCount_NO_GPU       GPUCount = 0
	GPUCount_SINGLE_GPU   GPUCount = 1
	GPUCount_MULTIPLE_GPU GPUCount = 2
)

var GPUCount_name = map[int32]string{
	0: "NO_GPU",
	1: "SINGLE_GPU",
	2: "MULTIPLE_GPU",
}
var GPUCount_value = map[string]int32{
	"NO_GPU":       0,
	"SINGLE_GPU":   1,
	"MULTIPLE_GPU": 2,
}

func (x GPUCount) String() string {
	return proto.EnumName(GPUCount_name, int32(x))
}
func (GPUCount) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Geo represent GeoIP results for node
type Geo struct {
	Country string  `protobuf:"bytes,1,opt,name=country" json:"country,omitempty"`
	City    string  `protobuf:"bytes,2,opt,name=city" json:"city,omitempty"`
	Lat     float32 `protobuf:"fixed32,3,opt,name=lat" json:"lat,omitempty"`
	Lon     float32 `protobuf:"fixed32,4,opt,name=lon" json:"lon,omitempty"`
}

func (m *Geo) Reset()                    { *m = Geo{} }
func (m *Geo) String() string            { return proto.CompactTextString(m) }
func (*Geo) ProtoMessage()               {}
func (*Geo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Geo) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Geo) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Geo) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Geo) GetLon() float32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

type Resources struct {
	// CPU core count
	CpuCores uint64 `protobuf:"varint,1,opt,name=cpuCores" json:"cpuCores,omitempty"`
	// RAM, in bytes
	RamBytes uint64 `protobuf:"varint,2,opt,name=ramBytes" json:"ramBytes,omitempty"`
	// GPU devices count
	GpuCount GPUCount `protobuf:"varint,3,opt,name=gpuCount,enum=sonm.GPUCount" json:"gpuCount,omitempty"`
	// todo: discuss
	// storage volume, in Megabytes
	Storage uint64 `protobuf:"varint,4,opt,name=storage" json:"storage,omitempty"`
	// Inbound network traffic (the higher value), in bytes
	NetTrafficIn uint64 `protobuf:"varint,5,opt,name=netTrafficIn" json:"netTrafficIn,omitempty"`
	// Outbound network traffic (the higher value), in bytes
	NetTrafficOut uint64 `protobuf:"varint,6,opt,name=netTrafficOut" json:"netTrafficOut,omitempty"`
	// Allowed network connections
	NetworkType NetworkType `protobuf:"varint,7,opt,name=networkType,enum=sonm.NetworkType" json:"networkType,omitempty"`
	// Other properties/benchmarks. The higher means better
	Props map[string]string `protobuf:"bytes,8,rep,name=props" json:"props,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Resources) Reset()                    { *m = Resources{} }
func (m *Resources) String() string            { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()               {}
func (*Resources) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Resources) GetCpuCores() uint64 {
	if m != nil {
		return m.CpuCores
	}
	return 0
}

func (m *Resources) GetRamBytes() uint64 {
	if m != nil {
		return m.RamBytes
	}
	return 0
}

func (m *Resources) GetGpuCount() GPUCount {
	if m != nil {
		return m.GpuCount
	}
	return GPUCount_NO_GPU
}

func (m *Resources) GetStorage() uint64 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *Resources) GetNetTrafficIn() uint64 {
	if m != nil {
		return m.NetTrafficIn
	}
	return 0
}

func (m *Resources) GetNetTrafficOut() uint64 {
	if m != nil {
		return m.NetTrafficOut
	}
	return 0
}

func (m *Resources) GetNetworkType() NetworkType {
	if m != nil {
		return m.NetworkType
	}
	return NetworkType_NO_NETWORK
}

func (m *Resources) GetProps() map[string]string {
	if m != nil {
		return m.Props
	}
	return nil
}

type Slot struct {
	// Buyer’s rating. Got from Buyer’s profile for BID orders rating_supplier.
	BuyerRating int64 `protobuf:"varint,1,opt,name=buyerRating" json:"buyerRating,omitempty"`
	// Supplier’s rating. Got from Supplier’s profile for ASK orders.
	SupplierRating int64 `protobuf:"varint,2,opt,name=supplierRating" json:"supplierRating,omitempty"`
	// Geo represent Worker's position
	Geo *Geo `protobuf:"bytes,3,opt,name=geo" json:"geo,omitempty"`
	// Hardware resources requirements
	Resources *Resources `protobuf:"bytes,4,opt,name=resources" json:"resources,omitempty"`
}

func (m *Slot) Reset()                    { *m = Slot{} }
func (m *Slot) String() string            { return proto.CompactTextString(m) }
func (*Slot) ProtoMessage()               {}
func (*Slot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Slot) GetBuyerRating() int64 {
	if m != nil {
		return m.BuyerRating
	}
	return 0
}

func (m *Slot) GetSupplierRating() int64 {
	if m != nil {
		return m.SupplierRating
	}
	return 0
}

func (m *Slot) GetGeo() *Geo {
	if m != nil {
		return m.Geo
	}
	return nil
}

func (m *Slot) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

type Order struct {
	// Order ID, UUIDv4
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Buyer's EtherumID
	ByuerID string `protobuf:"bytes,2,opt,name=byuerID" json:"byuerID,omitempty"`
	// Supplier's is EtherumID
	SupplierID string `protobuf:"bytes,3,opt,name=supplierID" json:"supplierID,omitempty"`
	// Order price
	Price int64 `protobuf:"varint,4,opt,name=price" json:"price,omitempty"`
	// Order type (Bid or Ask)
	OrderType OrderType `protobuf:"varint,5,opt,name=orderType,enum=sonm.OrderType" json:"orderType,omitempty"`
	// Slot describe resource requiements
	Slot *Slot `protobuf:"bytes,6,opt,name=slot" json:"slot,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetByuerID() string {
	if m != nil {
		return m.ByuerID
	}
	return ""
}

func (m *Order) GetSupplierID() string {
	if m != nil {
		return m.SupplierID
	}
	return ""
}

func (m *Order) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetOrderType() OrderType {
	if m != nil {
		return m.OrderType
	}
	return OrderType_ANY
}

func (m *Order) GetSlot() *Slot {
	if m != nil {
		return m.Slot
	}
	return nil
}

func init() {
	proto.RegisterType((*Geo)(nil), "sonm.Geo")
	proto.RegisterType((*Resources)(nil), "sonm.Resources")
	proto.RegisterType((*Slot)(nil), "sonm.Slot")
	proto.RegisterType((*Order)(nil), "sonm.Order")
	proto.RegisterEnum("sonm.OrderType", OrderType_name, OrderType_value)
	proto.RegisterEnum("sonm.NetworkType", NetworkType_name, NetworkType_value)
	proto.RegisterEnum("sonm.GPUCount", GPUCount_name, GPUCount_value)
}

func init() { proto.RegisterFile("bid.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x53, 0xdb, 0x6e, 0xda, 0x40,
	0x10, 0x8d, 0x2f, 0x24, 0xf6, 0x98, 0x52, 0x77, 0xd4, 0x07, 0x8b, 0x4a, 0x11, 0x42, 0x55, 0x84,
	0x90, 0x8a, 0x2a, 0xf2, 0x42, 0xfb, 0x56, 0x02, 0x42, 0x56, 0x12, 0x1b, 0x2d, 0xa0, 0xa8, 0x4f,
	0x91, 0x31, 0x1b, 0x64, 0x85, 0x78, 0xad, 0xb5, 0xdd, 0xca, 0xff, 0xd2, 0xef, 0xe8, 0x97, 0xf5,
	0x03, 0xaa, 0x5d, 0x5f, 0x20, 0x79, 0x9b, 0x73, 0xe6, 0x2c, 0x73, 0xe6, 0x8c, 0x01, 0x73, 0x1b,
	0xed, 0x46, 0x09, 0x67, 0x19, 0x43, 0x3d, 0x65, 0xf1, 0x4b, 0xff, 0x01, 0xb4, 0x05, 0x65, 0xe8,
	0xc0, 0x45, 0xc8, 0xf2, 0x38, 0xe3, 0x85, 0xa3, 0xf4, 0x94, 0x81, 0x49, 0x6a, 0x88, 0x08, 0x7a,
	0x18, 0x65, 0x85, 0xa3, 0x4a, 0x5a, 0xd6, 0x68, 0x83, 0x76, 0x08, 0x32, 0x47, 0xeb, 0x29, 0x03,
	0x95, 0x88, 0x52, 0x32, 0x2c, 0x76, 0xf4, 0x8a, 0x61, 0x71, 0xff, 0x9f, 0x0a, 0x26, 0xa1, 0x29,
	0xcb, 0x79, 0x48, 0x53, 0xec, 0x82, 0x11, 0x26, 0xf9, 0x0d, 0xe3, 0x34, 0x95, 0x03, 0x74, 0xd2,
	0x60, 0xd1, 0xe3, 0xc1, 0xcb, 0xb4, 0xc8, 0x68, 0x2a, 0xa7, 0xe8, 0xa4, 0xc1, 0x38, 0x04, 0x63,
	0x2f, 0x74, 0x79, 0x5c, 0x8e, 0xeb, 0x8c, 0x3b, 0x23, 0xe1, 0x7b, 0xb4, 0x58, 0x6e, 0x24, 0x4b,
	0x9a, 0xbe, 0xd8, 0x21, 0xcd, 0x18, 0x0f, 0xf6, 0x54, 0xfa, 0xd0, 0x49, 0x0d, 0xb1, 0x0f, 0xed,
	0x98, 0x66, 0x6b, 0x1e, 0x3c, 0x3d, 0x45, 0xa1, 0x1b, 0x3b, 0x2d, 0xd9, 0x7e, 0xc5, 0xe1, 0x67,
	0x78, 0x77, 0xc4, 0x7e, 0x9e, 0x39, 0xe7, 0x52, 0xf4, 0x9a, 0xc4, 0x6b, 0xb0, 0x62, 0x9a, 0xfd,
	0x66, 0xfc, 0x79, 0x5d, 0x24, 0xd4, 0xb9, 0x90, 0x96, 0x3e, 0x94, 0x96, 0xbc, 0x63, 0x83, 0x9c,
	0xaa, 0xf0, 0x2b, 0xb4, 0x12, 0xce, 0x92, 0xd4, 0x31, 0x7a, 0xda, 0xc0, 0x1a, 0x77, 0x4b, 0x79,
	0x13, 0xce, 0x68, 0x29, 0x9a, 0x73, 0x91, 0x36, 0x29, 0x85, 0xdd, 0x09, 0xc0, 0x91, 0x14, 0xe1,
	0x3e, 0xd3, 0xfa, 0x30, 0xa2, 0xc4, 0x8f, 0xd0, 0xfa, 0x15, 0x1c, 0x72, 0x5a, 0x5d, 0xa5, 0x04,
	0xdf, 0xd5, 0x89, 0xd2, 0xff, 0xa3, 0x80, 0xbe, 0x3a, 0xb0, 0x0c, 0x7b, 0x60, 0x6d, 0xf3, 0x82,
	0x72, 0x12, 0x64, 0x51, 0xbc, 0x97, 0x8f, 0x35, 0x72, 0x4a, 0xe1, 0x15, 0x74, 0xd2, 0x3c, 0x49,
	0x0e, 0x51, 0x23, 0x52, 0xa5, 0xe8, 0x0d, 0x8b, 0x9f, 0x40, 0xdb, 0x53, 0x26, 0xe3, 0xb7, 0xc6,
	0x66, 0x15, 0x3f, 0x65, 0x44, 0xb0, 0xf8, 0x05, 0x4c, 0x5e, 0x2f, 0x22, 0x63, 0xb7, 0xc6, 0xef,
	0xdf, 0xec, 0x47, 0x8e, 0x8a, 0xfe, 0x5f, 0x05, 0x5a, 0x3e, 0xdf, 0x51, 0x8e, 0x1d, 0x50, 0xa3,
	0x5d, 0xb5, 0x93, 0x1a, 0xed, 0xc4, 0xf5, 0xb6, 0x45, 0x4e, 0xb9, 0x3b, 0xab, 0x96, 0xaa, 0x21,
	0x5e, 0x02, 0xd4, 0x8e, 0xdc, 0x99, 0xb4, 0x61, 0x92, 0x13, 0x46, 0x84, 0x91, 0xf0, 0x28, 0x2c,
	0xaf, 0xae, 0x91, 0x12, 0x08, 0x63, 0x4c, 0x0c, 0x92, 0x77, 0x6a, 0xc9, 0x3b, 0x55, 0xc6, 0xfc,
	0x9a, 0x26, 0x47, 0x05, 0x5e, 0x82, 0x9e, 0x1e, 0x58, 0x79, 0x75, 0x6b, 0x0c, 0xa5, 0x52, 0x04,
	0x49, 0x24, 0x3f, 0xbc, 0x02, 0xb3, 0x79, 0x87, 0x17, 0xa0, 0xfd, 0xf0, 0x7e, 0xda, 0x67, 0xa2,
	0x98, 0xba, 0x33, 0x5b, 0x91, 0xcc, 0xea, 0xd6, 0x56, 0x87, 0xdf, 0xc0, 0x3a, 0xf9, 0x0e, 0xb0,
	0x03, 0xe0, 0xf9, 0x8f, 0xde, 0x7c, 0xfd, 0xe0, 0x93, 0x5b, 0xfb, 0x0c, 0xdb, 0x60, 0xf8, 0x9b,
	0xf5, 0xd4, 0xdf, 0x78, 0xe2, 0x55, 0x1b, 0x0c, 0xd7, 0xbb, 0xf1, 0xef, 0x5d, 0x6f, 0x61, 0xab,
	0xc3, 0x09, 0x18, 0xf5, 0x57, 0x8d, 0x00, 0xe7, 0x9e, 0xff, 0xb8, 0x58, 0x6e, 0xec, 0x33, 0xf1,
	0x1b, 0x2b, 0xd7, 0x5b, 0xdc, 0xcd, 0x25, 0x56, 0xd0, 0x86, 0xf6, 0xfd, 0xe6, 0x6e, 0xed, 0x2e,
	0x2b, 0x46, 0xdd, 0x9e, 0xcb, 0x7f, 0xf4, 0xf5, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xb9,
	0x1d, 0x09, 0xde, 0x03, 0x00, 0x00,
}
