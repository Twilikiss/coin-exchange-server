// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: order.proto

package order

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

type OrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip           string  `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Symbol       string  `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Page         int64   `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	PageSize     int64   `protobuf:"varint,5,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	UserId       int64   `protobuf:"varint,6,opt,name=userId,proto3" json:"userId,omitempty"`
	Price        float64 `protobuf:"fixed64,7,opt,name=price,proto3" json:"price,omitempty"`
	Amount       float64 `protobuf:"fixed64,8,opt,name=amount,proto3" json:"amount,omitempty"`
	Direction    string  `protobuf:"bytes,9,opt,name=direction,proto3" json:"direction,omitempty"`
	Type         string  `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	UseDiscount  int32   `protobuf:"varint,11,opt,name=useDiscount,proto3" json:"useDiscount,omitempty"`
	OrderId      string  `protobuf:"bytes,12,opt,name=orderId,proto3" json:"orderId,omitempty"`
	UpdateStatus int32   `protobuf:"varint,13,opt,name=updateStatus,proto3" json:"updateStatus,omitempty"`
}

func (x *OrderReq) Reset() {
	*x = OrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderReq) ProtoMessage() {}

func (x *OrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderReq.ProtoReflect.Descriptor instead.
func (*OrderReq) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *OrderReq) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *OrderReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *OrderReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OrderReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderReq) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderReq) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *OrderReq) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *OrderReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OrderReq) GetUseDiscount() int32 {
	if x != nil {
		return x.UseDiscount
	}
	return 0
}

func (x *OrderReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderReq) GetUpdateStatus() int32 {
	if x != nil {
		return x.UpdateStatus
	}
	return 0
}

type OrderRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*ExchangeOrder `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int64            `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *OrderRes) Reset() {
	*x = OrderRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRes) ProtoMessage() {}

func (x *OrderRes) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRes.ProtoReflect.Descriptor instead.
func (*OrderRes) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderRes) GetList() []*ExchangeOrder {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *OrderRes) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ExchangeOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId       string  `protobuf:"bytes,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Amount        float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	BaseSymbol    string  `protobuf:"bytes,4,opt,name=baseSymbol,proto3" json:"baseSymbol,omitempty"`
	CanceledTime  int64   `protobuf:"varint,5,opt,name=canceledTime,proto3" json:"canceledTime,omitempty"`
	CoinSymbol    string  `protobuf:"bytes,6,opt,name=coinSymbol,proto3" json:"coinSymbol,omitempty"`
	CompletedTime int64   `protobuf:"varint,7,opt,name=completedTime,proto3" json:"completedTime,omitempty"`
	Direction     string  `protobuf:"bytes,8,opt,name=direction,proto3" json:"direction,omitempty"`
	MemberId      int64   `protobuf:"varint,11,opt,name=memberId,proto3" json:"memberId,omitempty"`
	Price         float64 `protobuf:"fixed64,12,opt,name=price,proto3" json:"price,omitempty"`
	Status        string  `protobuf:"bytes,13,opt,name=status,proto3" json:"status,omitempty"`
	Symbol        string  `protobuf:"bytes,14,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Time          int64   `protobuf:"varint,15,opt,name=time,proto3" json:"time,omitempty"`
	TradedAmount  float64 `protobuf:"fixed64,16,opt,name=tradedAmount,proto3" json:"tradedAmount,omitempty"`
	Turnover      float64 `protobuf:"fixed64,17,opt,name=turnover,proto3" json:"turnover,omitempty"`
	Type          string  `protobuf:"bytes,18,opt,name=type,proto3" json:"type,omitempty"`
	UseDiscount   string  `protobuf:"bytes,21,opt,name=useDiscount,proto3" json:"useDiscount,omitempty"`
}

func (x *ExchangeOrder) Reset() {
	*x = ExchangeOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeOrder) ProtoMessage() {}

func (x *ExchangeOrder) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeOrder.ProtoReflect.Descriptor instead.
func (*ExchangeOrder) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *ExchangeOrder) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExchangeOrder) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *ExchangeOrder) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ExchangeOrder) GetBaseSymbol() string {
	if x != nil {
		return x.BaseSymbol
	}
	return ""
}

func (x *ExchangeOrder) GetCanceledTime() int64 {
	if x != nil {
		return x.CanceledTime
	}
	return 0
}

func (x *ExchangeOrder) GetCoinSymbol() string {
	if x != nil {
		return x.CoinSymbol
	}
	return ""
}

func (x *ExchangeOrder) GetCompletedTime() int64 {
	if x != nil {
		return x.CompletedTime
	}
	return 0
}

func (x *ExchangeOrder) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *ExchangeOrder) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *ExchangeOrder) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ExchangeOrder) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ExchangeOrder) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *ExchangeOrder) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *ExchangeOrder) GetTradedAmount() float64 {
	if x != nil {
		return x.TradedAmount
	}
	return 0
}

func (x *ExchangeOrder) GetTurnover() float64 {
	if x != nil {
		return x.Turnover
	}
	return 0
}

func (x *ExchangeOrder) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ExchangeOrder) GetUseDiscount() string {
	if x != nil {
		return x.UseDiscount
	}
	return ""
}

type AddOrderRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *AddOrderRes) Reset() {
	*x = AddOrderRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrderRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrderRes) ProtoMessage() {}

func (x *AddOrderRes) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrderRes.ProtoReflect.Descriptor instead.
func (*AddOrderRes) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *AddOrderRes) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type CancelOrderRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *CancelOrderRes) Reset() {
	*x = CancelOrderRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderRes) ProtoMessage() {}

func (x *CancelOrderRes) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderRes.ProtoReflect.Descriptor instead.
func (*CancelOrderRes) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *CancelOrderRes) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type ExchangeOrderOrigin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId       string  `protobuf:"bytes,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Amount        float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	BaseSymbol    string  `protobuf:"bytes,4,opt,name=baseSymbol,proto3" json:"baseSymbol,omitempty"`
	CanceledTime  int64   `protobuf:"varint,5,opt,name=canceledTime,proto3" json:"canceledTime,omitempty"`
	CoinSymbol    string  `protobuf:"bytes,6,opt,name=coinSymbol,proto3" json:"coinSymbol,omitempty"`
	CompletedTime int64   `protobuf:"varint,7,opt,name=completedTime,proto3" json:"completedTime,omitempty"`
	Direction     int32   `protobuf:"varint,8,opt,name=direction,proto3" json:"direction,omitempty"`
	MemberId      int64   `protobuf:"varint,11,opt,name=memberId,proto3" json:"memberId,omitempty"`
	Price         float64 `protobuf:"fixed64,12,opt,name=price,proto3" json:"price,omitempty"`
	Status        int32   `protobuf:"varint,13,opt,name=status,proto3" json:"status,omitempty"`
	Symbol        string  `protobuf:"bytes,14,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Time          int64   `protobuf:"varint,15,opt,name=time,proto3" json:"time,omitempty"`
	TradedAmount  float64 `protobuf:"fixed64,16,opt,name=tradedAmount,proto3" json:"tradedAmount,omitempty"`
	Turnover      float64 `protobuf:"fixed64,17,opt,name=turnover,proto3" json:"turnover,omitempty"`
	Type          int32   `protobuf:"varint,18,opt,name=type,proto3" json:"type,omitempty"`
	UseDiscount   string  `protobuf:"bytes,21,opt,name=useDiscount,proto3" json:"useDiscount,omitempty"`
}

func (x *ExchangeOrderOrigin) Reset() {
	*x = ExchangeOrderOrigin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeOrderOrigin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeOrderOrigin) ProtoMessage() {}

func (x *ExchangeOrderOrigin) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeOrderOrigin.ProtoReflect.Descriptor instead.
func (*ExchangeOrderOrigin) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *ExchangeOrderOrigin) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExchangeOrderOrigin) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *ExchangeOrderOrigin) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ExchangeOrderOrigin) GetBaseSymbol() string {
	if x != nil {
		return x.BaseSymbol
	}
	return ""
}

func (x *ExchangeOrderOrigin) GetCanceledTime() int64 {
	if x != nil {
		return x.CanceledTime
	}
	return 0
}

func (x *ExchangeOrderOrigin) GetCoinSymbol() string {
	if x != nil {
		return x.CoinSymbol
	}
	return ""
}

func (x *ExchangeOrderOrigin) GetCompletedTime() int64 {
	if x != nil {
		return x.CompletedTime
	}
	return 0
}

func (x *ExchangeOrderOrigin) GetDirection() int32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *ExchangeOrderOrigin) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *ExchangeOrderOrigin) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ExchangeOrderOrigin) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ExchangeOrderOrigin) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *ExchangeOrderOrigin) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *ExchangeOrderOrigin) GetTradedAmount() float64 {
	if x != nil {
		return x.TradedAmount
	}
	return 0
}

func (x *ExchangeOrderOrigin) GetTurnover() float64 {
	if x != nil {
		return x.Turnover
	}
	return 0
}

func (x *ExchangeOrderOrigin) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ExchangeOrderOrigin) GetUseDiscount() string {
	if x != nil {
		return x.UseDiscount
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0xba, 0x02, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x4a, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xe5, 0x03,
	0x0a, 0x0d, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x69, 0x6e, 0x53, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x69, 0x6e, 0x53,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x65, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x75, 0x72, 0x6e, 0x6f, 0x76, 0x65, 0x72, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x74, 0x75, 0x72, 0x6e, 0x6f, 0x76, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x27, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2a,
	0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0xeb, 0x03, 0x0a, 0x13, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x53, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x69, 0x6e,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f,
	0x69, 0x6e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x65, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x75, 0x72, 0x6e, 0x6f, 0x76,
	0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x74, 0x75, 0x72, 0x6e, 0x6f, 0x76,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x94, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x34, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x2a,
	0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x41,
	0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0d, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x35, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_order_proto_goTypes = []interface{}{
	(*OrderReq)(nil),            // 0: order.OrderReq
	(*OrderRes)(nil),            // 1: order.OrderRes
	(*ExchangeOrder)(nil),       // 2: order.ExchangeOrder
	(*AddOrderRes)(nil),         // 3: order.AddOrderRes
	(*CancelOrderRes)(nil),      // 4: order.CancelOrderRes
	(*ExchangeOrderOrigin)(nil), // 5: order.ExchangeOrderOrigin
}
var file_order_proto_depIdxs = []int32{
	2, // 0: order.OrderRes.list:type_name -> order.ExchangeOrder
	0, // 1: order.Order.FindOrderHistory:input_type -> order.OrderReq
	0, // 2: order.Order.FindOrderCurrent:input_type -> order.OrderReq
	0, // 3: order.Order.Add:input_type -> order.OrderReq
	0, // 4: order.Order.FindByOrderId:input_type -> order.OrderReq
	0, // 5: order.Order.CancelOrder:input_type -> order.OrderReq
	1, // 6: order.Order.FindOrderHistory:output_type -> order.OrderRes
	1, // 7: order.Order.FindOrderCurrent:output_type -> order.OrderRes
	3, // 8: order.Order.Add:output_type -> order.AddOrderRes
	5, // 9: order.Order.FindByOrderId:output_type -> order.ExchangeOrderOrigin
	4, // 10: order.Order.CancelOrder:output_type -> order.CancelOrderRes
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderReq); i {
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
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRes); i {
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
		file_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeOrder); i {
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
		file_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrderRes); i {
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
		file_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOrderRes); i {
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
		file_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeOrderOrigin); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
