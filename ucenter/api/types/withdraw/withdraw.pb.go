// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: withdraw.proto

package withdraw

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

type WithdrawReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinId     int64   `protobuf:"varint,1,opt,name=coinId,proto3" json:"coinId,omitempty"`
	Ip         string  `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	UserId     int64   `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Phone      string  `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Unit       string  `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
	Address    string  `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Amount     float64 `protobuf:"fixed64,7,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee        float64 `protobuf:"fixed64,8,opt,name=fee,proto3" json:"fee,omitempty"`
	JyPassword string  `protobuf:"bytes,9,opt,name=jyPassword,proto3" json:"jyPassword,omitempty"`
	Code       string  `protobuf:"bytes,10,opt,name=code,proto3" json:"code,omitempty"`
	Page       int64   `protobuf:"varint,11,opt,name=page,proto3" json:"page,omitempty"`
	PageSize   int64   `protobuf:"varint,12,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *WithdrawReq) Reset() {
	*x = WithdrawReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawReq) ProtoMessage() {}

func (x *WithdrawReq) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawReq.ProtoReflect.Descriptor instead.
func (*WithdrawReq) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{0}
}

func (x *WithdrawReq) GetCoinId() int64 {
	if x != nil {
		return x.CoinId
	}
	return 0
}

func (x *WithdrawReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *WithdrawReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *WithdrawReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *WithdrawReq) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *WithdrawReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *WithdrawReq) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *WithdrawReq) GetFee() float64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *WithdrawReq) GetJyPassword() string {
	if x != nil {
		return x.JyPassword
	}
	return ""
}

func (x *WithdrawReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *WithdrawReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *WithdrawReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type AddressSimple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Remark  string `protobuf:"bytes,1,opt,name=remark,proto3" json:"remark,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *AddressSimple) Reset() {
	*x = AddressSimple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressSimple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressSimple) ProtoMessage() {}

func (x *AddressSimple) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressSimple.ProtoReflect.Descriptor instead.
func (*AddressSimple) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{1}
}

func (x *AddressSimple) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *AddressSimple) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type AddressSimpleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*AddressSimple `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *AddressSimpleList) Reset() {
	*x = AddressSimpleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressSimpleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressSimpleList) ProtoMessage() {}

func (x *AddressSimpleList) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressSimpleList.ProtoReflect.Descriptor instead.
func (*AddressSimpleList) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{2}
}

func (x *AddressSimpleList) GetList() []*AddressSimple {
	if x != nil {
		return x.List
	}
	return nil
}

type NoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoRes) Reset() {
	*x = NoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoRes) ProtoMessage() {}

func (x *NoRes) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoRes.ProtoReflect.Descriptor instead.
func (*NoRes) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{3}
}

type Coin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CanAutoWithdraw   int32   `protobuf:"varint,3,opt,name=canAutoWithdraw,proto3" json:"canAutoWithdraw,omitempty"`
	CanRecharge       int32   `protobuf:"varint,4,opt,name=canRecharge,proto3" json:"canRecharge,omitempty"`
	CanTransfer       int32   `protobuf:"varint,5,opt,name=canTransfer,proto3" json:"canTransfer,omitempty"`
	CanWithdraw       int32   `protobuf:"varint,6,opt,name=canWithdraw,proto3" json:"canWithdraw,omitempty"`
	CnyRate           float64 `protobuf:"fixed64,7,opt,name=cnyRate,proto3" json:"cnyRate,omitempty"`
	EnableRpc         int32   `protobuf:"varint,8,opt,name=enableRpc,proto3" json:"enableRpc,omitempty"`
	IsPlatformCoin    int32   `protobuf:"varint,9,opt,name=isPlatformCoin,proto3" json:"isPlatformCoin,omitempty"`
	MaxTxFee          float64 `protobuf:"fixed64,10,opt,name=maxTxFee,proto3" json:"maxTxFee,omitempty"`
	MaxWithdrawAmount float64 `protobuf:"fixed64,11,opt,name=maxWithdrawAmount,proto3" json:"maxWithdrawAmount,omitempty"`
	MinTxFee          float64 `protobuf:"fixed64,12,opt,name=minTxFee,proto3" json:"minTxFee,omitempty"`
	MinWithdrawAmount float64 `protobuf:"fixed64,13,opt,name=minWithdrawAmount,proto3" json:"minWithdrawAmount,omitempty"`
	NameCn            string  `protobuf:"bytes,14,opt,name=nameCn,proto3" json:"nameCn,omitempty"`
	Sort              int32   `protobuf:"varint,15,opt,name=sort,proto3" json:"sort,omitempty"`
	Status            int32   `protobuf:"varint,16,opt,name=status,proto3" json:"status,omitempty"`
	Unit              string  `protobuf:"bytes,17,opt,name=unit,proto3" json:"unit,omitempty"`
	UsdRate           float64 `protobuf:"fixed64,18,opt,name=usdRate,proto3" json:"usdRate,omitempty"`
	WithdrawThreshold float64 `protobuf:"fixed64,19,opt,name=withdrawThreshold,proto3" json:"withdrawThreshold,omitempty"`
	HasLegal          int32   `protobuf:"varint,20,opt,name=hasLegal,proto3" json:"hasLegal,omitempty"`
	ColdWalletAddress string  `protobuf:"bytes,21,opt,name=coldWalletAddress,proto3" json:"coldWalletAddress,omitempty"`
	MinerFee          float64 `protobuf:"fixed64,22,opt,name=minerFee,proto3" json:"minerFee,omitempty"`
	WithdrawScale     int32   `protobuf:"varint,23,opt,name=withdrawScale,proto3" json:"withdrawScale,omitempty"`
	AccountType       int32   `protobuf:"varint,24,opt,name=accountType,proto3" json:"accountType,omitempty"`
	DepositAddress    string  `protobuf:"bytes,25,opt,name=depositAddress,proto3" json:"depositAddress,omitempty"`
	Infolink          string  `protobuf:"bytes,26,opt,name=infolink,proto3" json:"infolink,omitempty"`
	Information       string  `protobuf:"bytes,27,opt,name=information,proto3" json:"information,omitempty"`
	MinRechargeAmount float64 `protobuf:"fixed64,28,opt,name=minRechargeAmount,proto3" json:"minRechargeAmount,omitempty"`
}

func (x *Coin) Reset() {
	*x = Coin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coin) ProtoMessage() {}

func (x *Coin) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coin.ProtoReflect.Descriptor instead.
func (*Coin) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{4}
}

func (x *Coin) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Coin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Coin) GetCanAutoWithdraw() int32 {
	if x != nil {
		return x.CanAutoWithdraw
	}
	return 0
}

func (x *Coin) GetCanRecharge() int32 {
	if x != nil {
		return x.CanRecharge
	}
	return 0
}

func (x *Coin) GetCanTransfer() int32 {
	if x != nil {
		return x.CanTransfer
	}
	return 0
}

func (x *Coin) GetCanWithdraw() int32 {
	if x != nil {
		return x.CanWithdraw
	}
	return 0
}

func (x *Coin) GetCnyRate() float64 {
	if x != nil {
		return x.CnyRate
	}
	return 0
}

func (x *Coin) GetEnableRpc() int32 {
	if x != nil {
		return x.EnableRpc
	}
	return 0
}

func (x *Coin) GetIsPlatformCoin() int32 {
	if x != nil {
		return x.IsPlatformCoin
	}
	return 0
}

func (x *Coin) GetMaxTxFee() float64 {
	if x != nil {
		return x.MaxTxFee
	}
	return 0
}

func (x *Coin) GetMaxWithdrawAmount() float64 {
	if x != nil {
		return x.MaxWithdrawAmount
	}
	return 0
}

func (x *Coin) GetMinTxFee() float64 {
	if x != nil {
		return x.MinTxFee
	}
	return 0
}

func (x *Coin) GetMinWithdrawAmount() float64 {
	if x != nil {
		return x.MinWithdrawAmount
	}
	return 0
}

func (x *Coin) GetNameCn() string {
	if x != nil {
		return x.NameCn
	}
	return ""
}

func (x *Coin) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Coin) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Coin) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Coin) GetUsdRate() float64 {
	if x != nil {
		return x.UsdRate
	}
	return 0
}

func (x *Coin) GetWithdrawThreshold() float64 {
	if x != nil {
		return x.WithdrawThreshold
	}
	return 0
}

func (x *Coin) GetHasLegal() int32 {
	if x != nil {
		return x.HasLegal
	}
	return 0
}

func (x *Coin) GetColdWalletAddress() string {
	if x != nil {
		return x.ColdWalletAddress
	}
	return ""
}

func (x *Coin) GetMinerFee() float64 {
	if x != nil {
		return x.MinerFee
	}
	return 0
}

func (x *Coin) GetWithdrawScale() int32 {
	if x != nil {
		return x.WithdrawScale
	}
	return 0
}

func (x *Coin) GetAccountType() int32 {
	if x != nil {
		return x.AccountType
	}
	return 0
}

func (x *Coin) GetDepositAddress() string {
	if x != nil {
		return x.DepositAddress
	}
	return ""
}

func (x *Coin) GetInfolink() string {
	if x != nil {
		return x.Infolink
	}
	return ""
}

func (x *Coin) GetInformation() string {
	if x != nil {
		return x.Information
	}
	return ""
}

func (x *Coin) GetMinRechargeAmount() float64 {
	if x != nil {
		return x.MinRechargeAmount
	}
	return 0
}

type WithdrawRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MemberId          int64   `protobuf:"varint,2,opt,name=memberId,proto3" json:"memberId,omitempty"`
	Coin              *Coin   `protobuf:"bytes,3,opt,name=coin,proto3" json:"coin,omitempty"`
	TotalAmount       float64 `protobuf:"fixed64,4,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
	Fee               float64 `protobuf:"fixed64,5,opt,name=fee,proto3" json:"fee,omitempty"`
	ArrivedAmount     float64 `protobuf:"fixed64,6,opt,name=arrivedAmount,proto3" json:"arrivedAmount,omitempty"`
	Address           string  `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Remark            string  `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	TransactionNumber string  `protobuf:"bytes,9,opt,name=transactionNumber,proto3" json:"transactionNumber,omitempty"`
	CanAutoWithdraw   int32   `protobuf:"varint,10,opt,name=canAutoWithdraw,proto3" json:"canAutoWithdraw,omitempty"`
	IsAuto            int32   `protobuf:"varint,11,opt,name=isAuto,proto3" json:"isAuto,omitempty"`
	Status            int32   `protobuf:"varint,12,opt,name=status,proto3" json:"status,omitempty"`
	CreateTime        string  `protobuf:"bytes,13,opt,name=createTime,proto3" json:"createTime,omitempty"`
	DealTime          string  `protobuf:"bytes,14,opt,name=dealTime,proto3" json:"dealTime,omitempty"`
}

func (x *WithdrawRecord) Reset() {
	*x = WithdrawRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawRecord) ProtoMessage() {}

func (x *WithdrawRecord) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawRecord.ProtoReflect.Descriptor instead.
func (*WithdrawRecord) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{5}
}

func (x *WithdrawRecord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WithdrawRecord) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *WithdrawRecord) GetCoin() *Coin {
	if x != nil {
		return x.Coin
	}
	return nil
}

func (x *WithdrawRecord) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *WithdrawRecord) GetFee() float64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *WithdrawRecord) GetArrivedAmount() float64 {
	if x != nil {
		return x.ArrivedAmount
	}
	return 0
}

func (x *WithdrawRecord) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *WithdrawRecord) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *WithdrawRecord) GetTransactionNumber() string {
	if x != nil {
		return x.TransactionNumber
	}
	return ""
}

func (x *WithdrawRecord) GetCanAutoWithdraw() int32 {
	if x != nil {
		return x.CanAutoWithdraw
	}
	return 0
}

func (x *WithdrawRecord) GetIsAuto() int32 {
	if x != nil {
		return x.IsAuto
	}
	return 0
}

func (x *WithdrawRecord) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *WithdrawRecord) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *WithdrawRecord) GetDealTime() string {
	if x != nil {
		return x.DealTime
	}
	return ""
}

type RecordList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*WithdrawRecord `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int64             `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *RecordList) Reset() {
	*x = RecordList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordList) ProtoMessage() {}

func (x *RecordList) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordList.ProtoReflect.Descriptor instead.
func (*RecordList) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{6}
}

func (x *RecordList) GetList() []*WithdrawRecord {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *RecordList) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_withdraw_proto protoreflect.FileDescriptor

var file_withdraw_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x22, 0x9f, 0x02, 0x0a, 0x0b, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f,
	0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x6f, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x79, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6a, 0x79,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x41, 0x0a, 0x0d,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x40, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0x07, 0x0a, 0x05, 0x4e, 0x6f, 0x52, 0x65, 0x73, 0x22, 0x90, 0x07, 0x0a, 0x04, 0x43,
	0x6f, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x41, 0x75,
	0x74, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x63, 0x61, 0x6e, 0x41, 0x75, 0x74, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6e, 0x79, 0x52, 0x61,
	0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x63, 0x6e, 0x79, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x70, 0x63, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x70, 0x63, 0x12,
	0x26, 0x0a, 0x0e, 0x69, 0x73, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x69,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x69, 0x73, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x54, 0x78,
	0x46, 0x65, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x54, 0x78,
	0x46, 0x65, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11,
	0x6d, 0x61, 0x78, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x54, 0x78, 0x46, 0x65, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x54, 0x78, 0x46, 0x65, 0x65, 0x12, 0x2c, 0x0a,
	0x11, 0x6d, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x6d, 0x69, 0x6e, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x61, 0x6d, 0x65, 0x43, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d,
	0x65, 0x43, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x64, 0x52, 0x61, 0x74, 0x65, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x75, 0x73, 0x64, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a,
	0x11, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x61, 0x73, 0x4c, 0x65, 0x67, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x68,
	0x61, 0x73, 0x4c, 0x65, 0x67, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f, 0x6c, 0x64, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x6c, 0x64, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x46, 0x65,
	0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x46, 0x65,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x63, 0x61,
	0x6c, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x20, 0x0a,
	0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x1b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x0a, 0x11, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xb0, 0x03,
	0x0a, 0x0e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x04,
	0x63, 0x6f, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x04, 0x63, 0x6f, 0x69, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x66, 0x65, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x72, 0x72, 0x69, 0x76, 0x65, 0x64, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x61, 0x72, 0x72,
	0x69, 0x76, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x2c, 0x0a, 0x11,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x61,
	0x6e, 0x41, 0x75, 0x74, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x61, 0x6e, 0x41, 0x75, 0x74, 0x6f, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x41, 0x75, 0x74, 0x6f, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x73, 0x41, 0x75, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x50, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x77,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x32, 0x80, 0x02, 0x0a, 0x08, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12,
	0x49, 0x0a, 0x13, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x79,
	0x43, 0x6f, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e,
	0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x53, 0x65,
	0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e,
	0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x4e, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x36,
	0x0a, 0x0c, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x15,
	0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x2e, 0x4e, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x1a,
	0x14, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_withdraw_proto_rawDescOnce sync.Once
	file_withdraw_proto_rawDescData = file_withdraw_proto_rawDesc
)

func file_withdraw_proto_rawDescGZIP() []byte {
	file_withdraw_proto_rawDescOnce.Do(func() {
		file_withdraw_proto_rawDescData = protoimpl.X.CompressGZIP(file_withdraw_proto_rawDescData)
	})
	return file_withdraw_proto_rawDescData
}

var file_withdraw_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_withdraw_proto_goTypes = []interface{}{
	(*WithdrawReq)(nil),       // 0: withdraw.WithdrawReq
	(*AddressSimple)(nil),     // 1: withdraw.AddressSimple
	(*AddressSimpleList)(nil), // 2: withdraw.AddressSimpleList
	(*NoRes)(nil),             // 3: withdraw.NoRes
	(*Coin)(nil),              // 4: withdraw.Coin
	(*WithdrawRecord)(nil),    // 5: withdraw.WithdrawRecord
	(*RecordList)(nil),        // 6: withdraw.RecordList
}
var file_withdraw_proto_depIdxs = []int32{
	1, // 0: withdraw.AddressSimpleList.list:type_name -> withdraw.AddressSimple
	4, // 1: withdraw.WithdrawRecord.coin:type_name -> withdraw.Coin
	5, // 2: withdraw.RecordList.list:type_name -> withdraw.WithdrawRecord
	0, // 3: withdraw.Withdraw.findAddressByCoinId:input_type -> withdraw.WithdrawReq
	0, // 4: withdraw.Withdraw.SendCode:input_type -> withdraw.WithdrawReq
	0, // 5: withdraw.Withdraw.WithdrawCode:input_type -> withdraw.WithdrawReq
	0, // 6: withdraw.Withdraw.WithdrawRecord:input_type -> withdraw.WithdrawReq
	2, // 7: withdraw.Withdraw.findAddressByCoinId:output_type -> withdraw.AddressSimpleList
	3, // 8: withdraw.Withdraw.SendCode:output_type -> withdraw.NoRes
	3, // 9: withdraw.Withdraw.WithdrawCode:output_type -> withdraw.NoRes
	6, // 10: withdraw.Withdraw.WithdrawRecord:output_type -> withdraw.RecordList
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_withdraw_proto_init() }
func file_withdraw_proto_init() {
	if File_withdraw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_withdraw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawReq); i {
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
		file_withdraw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressSimple); i {
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
		file_withdraw_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressSimpleList); i {
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
		file_withdraw_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoRes); i {
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
		file_withdraw_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coin); i {
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
		file_withdraw_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawRecord); i {
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
		file_withdraw_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordList); i {
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
			RawDescriptor: file_withdraw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_withdraw_proto_goTypes,
		DependencyIndexes: file_withdraw_proto_depIdxs,
		MessageInfos:      file_withdraw_proto_msgTypes,
	}.Build()
	File_withdraw_proto = out.File
	file_withdraw_proto_rawDesc = nil
	file_withdraw_proto_goTypes = nil
	file_withdraw_proto_depIdxs = nil
}
