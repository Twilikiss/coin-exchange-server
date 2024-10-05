// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: asset.proto

package asset

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

type AssetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinName  string `protobuf:"bytes,1,opt,name=coinName,proto3" json:"coinName,omitempty"`
	Ip        string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	UserId    int64  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	StartTime string `protobuf:"bytes,4,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime   string `protobuf:"bytes,5,opt,name=endTime,proto3" json:"endTime,omitempty"`
	PageNo    int64  `protobuf:"varint,6,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize  int64  `protobuf:"varint,7,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Type      string `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	Symbol    string `protobuf:"bytes,9,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *AssetReq) Reset() {
	*x = AssetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetReq) ProtoMessage() {}

func (x *AssetReq) ProtoReflect() protoreflect.Message {
	mi := &file_asset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetReq.ProtoReflect.Descriptor instead.
func (*AssetReq) Descriptor() ([]byte, []int) {
	return file_asset_proto_rawDescGZIP(), []int{0}
}

func (x *AssetReq) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *AssetReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *AssetReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AssetReq) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *AssetReq) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *AssetReq) GetPageNo() int64 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *AssetReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AssetReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AssetReq) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
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
		mi := &file_asset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coin) ProtoMessage() {}

func (x *Coin) ProtoReflect() protoreflect.Message {
	mi := &file_asset_proto_msgTypes[1]
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
	return file_asset_proto_rawDescGZIP(), []int{1}
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

type MemberWallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address        string  `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Balance        float64 `protobuf:"fixed64,3,opt,name=balance,proto3" json:"balance,omitempty"`
	FrozenBalance  float64 `protobuf:"fixed64,4,opt,name=frozenBalance,proto3" json:"frozenBalance,omitempty"`
	ReleaseBalance float64 `protobuf:"fixed64,5,opt,name=releaseBalance,proto3" json:"releaseBalance,omitempty"`
	IsLock         int32   `protobuf:"varint,6,opt,name=isLock,proto3" json:"isLock,omitempty"`
	MemberId       int64   `protobuf:"varint,7,opt,name=memberId,proto3" json:"memberId,omitempty"`
	Version        int32   `protobuf:"varint,8,opt,name=version,proto3" json:"version,omitempty"`
	Coin           *Coin   `protobuf:"bytes,9,opt,name=coin,proto3" json:"coin,omitempty"`
	ToReleased     float64 `protobuf:"fixed64,10,opt,name=toReleased,proto3" json:"toReleased,omitempty"`
}

func (x *MemberWallet) Reset() {
	*x = MemberWallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberWallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberWallet) ProtoMessage() {}

func (x *MemberWallet) ProtoReflect() protoreflect.Message {
	mi := &file_asset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberWallet.ProtoReflect.Descriptor instead.
func (*MemberWallet) Descriptor() ([]byte, []int) {
	return file_asset_proto_rawDescGZIP(), []int{2}
}

func (x *MemberWallet) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberWallet) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MemberWallet) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *MemberWallet) GetFrozenBalance() float64 {
	if x != nil {
		return x.FrozenBalance
	}
	return 0
}

func (x *MemberWallet) GetReleaseBalance() float64 {
	if x != nil {
		return x.ReleaseBalance
	}
	return 0
}

func (x *MemberWallet) GetIsLock() int32 {
	if x != nil {
		return x.IsLock
	}
	return 0
}

func (x *MemberWallet) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *MemberWallet) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *MemberWallet) GetCoin() *Coin {
	if x != nil {
		return x.Coin
	}
	return nil
}

func (x *MemberWallet) GetToReleased() float64 {
	if x != nil {
		return x.ToReleased
	}
	return 0
}

type MemberWalletList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*MemberWallet `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *MemberWalletList) Reset() {
	*x = MemberWalletList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberWalletList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberWalletList) ProtoMessage() {}

func (x *MemberWalletList) ProtoReflect() protoreflect.Message {
	mi := &file_asset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberWalletList.ProtoReflect.Descriptor instead.
func (*MemberWalletList) Descriptor() ([]byte, []int) {
	return file_asset_proto_rawDescGZIP(), []int{3}
}

func (x *MemberWalletList) GetList() []*MemberWallet {
	if x != nil {
		return x.List
	}
	return nil
}

type MemberTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address     string  `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Amount      float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	CreateTime  string  `protobuf:"bytes,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Fee         float64 `protobuf:"fixed64,5,opt,name=fee,proto3" json:"fee,omitempty"`
	Flag        int32   `protobuf:"varint,6,opt,name=flag,proto3" json:"flag,omitempty"`
	MemberId    int64   `protobuf:"varint,7,opt,name=memberId,proto3" json:"memberId,omitempty"`
	Symbol      string  `protobuf:"bytes,8,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Type        string  `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`
	DiscountFee string  `protobuf:"bytes,10,opt,name=discountFee,proto3" json:"discountFee,omitempty"`
	RealFee     string  `protobuf:"bytes,11,opt,name=realFee,proto3" json:"realFee,omitempty"`
}

func (x *MemberTransaction) Reset() {
	*x = MemberTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asset_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberTransaction) ProtoMessage() {}

func (x *MemberTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_asset_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberTransaction.ProtoReflect.Descriptor instead.
func (*MemberTransaction) Descriptor() ([]byte, []int) {
	return file_asset_proto_rawDescGZIP(), []int{4}
}

func (x *MemberTransaction) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberTransaction) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MemberTransaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *MemberTransaction) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *MemberTransaction) GetFee() float64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *MemberTransaction) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *MemberTransaction) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *MemberTransaction) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *MemberTransaction) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MemberTransaction) GetDiscountFee() string {
	if x != nil {
		return x.DiscountFee
	}
	return ""
}

func (x *MemberTransaction) GetRealFee() string {
	if x != nil {
		return x.RealFee
	}
	return ""
}

type MemberTransactionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*MemberTransaction `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int64                `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *MemberTransactionList) Reset() {
	*x = MemberTransactionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asset_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberTransactionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberTransactionList) ProtoMessage() {}

func (x *MemberTransactionList) ProtoReflect() protoreflect.Message {
	mi := &file_asset_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberTransactionList.ProtoReflect.Descriptor instead.
func (*MemberTransactionList) Descriptor() ([]byte, []int) {
	return file_asset_proto_rawDescGZIP(), []int{5}
}

func (x *MemberTransactionList) GetList() []*MemberTransaction {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *MemberTransactionList) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type RestAddrResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RestAddrResp) Reset() {
	*x = RestAddrResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asset_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestAddrResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestAddrResp) ProtoMessage() {}

func (x *RestAddrResp) ProtoReflect() protoreflect.Message {
	mi := &file_asset_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestAddrResp.ProtoReflect.Descriptor instead.
func (*RestAddrResp) Descriptor() ([]byte, []int) {
	return file_asset_proto_rawDescGZIP(), []int{6}
}

type AddressList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *AddressList) Reset() {
	*x = AddressList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asset_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressList) ProtoMessage() {}

func (x *AddressList) ProtoReflect() protoreflect.Message {
	mi := &file_asset_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressList.ProtoReflect.Descriptor instead.
func (*AddressList) Descriptor() ([]byte, []int) {
	return file_asset_proto_rawDescGZIP(), []int{7}
}

func (x *AddressList) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

var File_asset_proto protoreflect.FileDescriptor

var file_asset_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x22, 0xe6, 0x01, 0x0a, 0x08, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70,
	0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x90, 0x07,
	0x0a, 0x04, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x61,
	0x6e, 0x41, 0x75, 0x74, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x61, 0x6e, 0x41, 0x75, 0x74, 0x6f, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x52, 0x65,
	0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x61, 0x6e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63,
	0x61, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6e,
	0x79, 0x52, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x63, 0x6e, 0x79,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x70,
	0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x70, 0x63, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x73, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x43, 0x6f, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x69, 0x73, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61,
	0x78, 0x54, 0x78, 0x46, 0x65, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x61,
	0x78, 0x54, 0x78, 0x46, 0x65, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x54, 0x78, 0x46, 0x65, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x54, 0x78, 0x46, 0x65, 0x65,
	0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x6d, 0x69, 0x6e,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x64, 0x52, 0x61, 0x74,
	0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x75, 0x73, 0x64, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x2c, 0x0a, 0x11, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x68, 0x61, 0x73, 0x4c, 0x65, 0x67, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x68, 0x61, 0x73, 0x4c, 0x65, 0x67, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f,
	0x6c, 0x64, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x6c, 0x64, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x65,
	0x72, 0x46, 0x65, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x65,
	0x72, 0x46, 0x65, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x53, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x6c, 0x69, 0x6e, 0x6b,
	0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x6c, 0x69, 0x6e, 0x6b,
	0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67,
	0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0xaf, 0x02, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x66, 0x72,
	0x6f, 0x7a, 0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x4c, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x73, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x04, 0x63, 0x6f,
	0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x64, 0x22, 0x3b, 0x0a, 0x10, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22,
	0x9f, 0x02, 0x0a, 0x11, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61,
	0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x46, 0x65, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x46, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x46,
	0x65, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x46, 0x65,
	0x65, 0x22, 0x5b, 0x0a, 0x15, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x0e,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x21,
	0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x32, 0xa9, 0x02, 0x0a, 0x05, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x3a, 0x0a, 0x12, 0x66,
	0x69, 0x6e, 0x64, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x79, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x12, 0x0f, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x64, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x34, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x0f, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x13, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x43, 0x0a, 0x12, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0a, 0x67, 0x65,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0f, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_asset_proto_rawDescOnce sync.Once
	file_asset_proto_rawDescData = file_asset_proto_rawDesc
)

func file_asset_proto_rawDescGZIP() []byte {
	file_asset_proto_rawDescOnce.Do(func() {
		file_asset_proto_rawDescData = protoimpl.X.CompressGZIP(file_asset_proto_rawDescData)
	})
	return file_asset_proto_rawDescData
}

var file_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_asset_proto_goTypes = []interface{}{
	(*AssetReq)(nil),              // 0: asset.AssetReq
	(*Coin)(nil),                  // 1: asset.Coin
	(*MemberWallet)(nil),          // 2: asset.MemberWallet
	(*MemberWalletList)(nil),      // 3: asset.MemberWalletList
	(*MemberTransaction)(nil),     // 4: asset.MemberTransaction
	(*MemberTransactionList)(nil), // 5: asset.MemberTransactionList
	(*RestAddrResp)(nil),          // 6: asset.RestAddrResp
	(*AddressList)(nil),           // 7: asset.AddressList
}
var file_asset_proto_depIdxs = []int32{
	1, // 0: asset.MemberWallet.coin:type_name -> asset.Coin
	2, // 1: asset.MemberWalletList.list:type_name -> asset.MemberWallet
	4, // 2: asset.MemberTransactionList.list:type_name -> asset.MemberTransaction
	0, // 3: asset.Asset.findWalletBySymbol:input_type -> asset.AssetReq
	0, // 4: asset.Asset.findWallet:input_type -> asset.AssetReq
	0, // 5: asset.Asset.resetAddress:input_type -> asset.AssetReq
	0, // 6: asset.Asset.findAllTransaction:input_type -> asset.AssetReq
	0, // 7: asset.Asset.getAddress:input_type -> asset.AssetReq
	2, // 8: asset.Asset.findWalletBySymbol:output_type -> asset.MemberWallet
	3, // 9: asset.Asset.findWallet:output_type -> asset.MemberWalletList
	6, // 10: asset.Asset.resetAddress:output_type -> asset.RestAddrResp
	5, // 11: asset.Asset.findAllTransaction:output_type -> asset.MemberTransactionList
	7, // 12: asset.Asset.getAddress:output_type -> asset.AddressList
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_asset_proto_init() }
func file_asset_proto_init() {
	if File_asset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_asset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetReq); i {
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
		file_asset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_asset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberWallet); i {
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
		file_asset_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberWalletList); i {
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
		file_asset_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberTransaction); i {
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
		file_asset_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberTransactionList); i {
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
		file_asset_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestAddrResp); i {
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
		file_asset_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressList); i {
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
			RawDescriptor: file_asset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_asset_proto_goTypes,
		DependencyIndexes: file_asset_proto_depIdxs,
		MessageInfos:      file_asset_proto_msgTypes,
	}.Build()
	File_asset_proto = out.File
	file_asset_proto_rawDesc = nil
	file_asset_proto_goTypes = nil
	file_asset_proto_depIdxs = nil
}
