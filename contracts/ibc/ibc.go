// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibc

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Height is an auto generated low-level Go binding around an user-defined struct.
type Height struct {
	RevisionNumber *big.Int
	RevisionHeight *big.Int
}

// IIBCMsgRecvPacket is an auto generated low-level Go binding around an user-defined struct.
type IIBCMsgRecvPacket struct {
	Packet          Packet
	ProofCommitment []byte
	ProofHeight     Height
	Signer          string
}

// Packet is an auto generated low-level Go binding around an user-defined struct.
type Packet struct {
	Sequence           *big.Int
	SourcePort         string
	SourceChannel      string
	DestinationPort    string
	DestinationChannel string
	Data               []byte
	TimeoutHeight      Height
	TimeoutTimestamp   *big.Int
}

// IBCMetaData contains all meta data concerning the IBC contract.
var IBCMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"timeoutHeight\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destPort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int32\",\"name\":\"channelOrdering\",\"type\":\"int32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"connectionID\",\"type\":\"string\"}],\"name\":\"AckPacket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"timeoutHeight\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destPort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"AcknowledgementError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyChannelId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyPortID\",\"type\":\"string\"}],\"name\":\"ChannelCloseConfirm\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyChannelId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyPortID\",\"type\":\"string\"}],\"name\":\"ChannelCloseInit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyChannelId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyPortID\",\"type\":\"string\"}],\"name\":\"ChannelOpenAck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyChannelId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyPortID\",\"type\":\"string\"}],\"name\":\"ChannelOpenConfirm\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyPortID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"ChannelOpenInit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyChannelId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyPortID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"ChannelOpenTry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"consensusHeight\",\"type\":\"string\"}],\"name\":\"ClientCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"consensusHeight\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"clientMessage\",\"type\":\"bytes\"}],\"name\":\"ClientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"consensusHeight\",\"type\":\"string\"}],\"name\":\"ClientUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyClientID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyConnectionId\",\"type\":\"string\"}],\"name\":\"ConnectionOpenAck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyClientID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyConnectionId\",\"type\":\"string\"}],\"name\":\"ConnectionOpenConfirm\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyClientID\",\"type\":\"string\"}],\"name\":\"ConnectionOpenInit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyClientID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"counterpartyConnectionId\",\"type\":\"string\"}],\"name\":\"ConnectionOpenTry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"timeoutHeight\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destPort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int32\",\"name\":\"channelOrdering\",\"type\":\"int32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"connection\",\"type\":\"string\"}],\"name\":\"PacketRecv\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"timeoutHeight\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destPort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int32\",\"name\":\"channelOrdering\",\"type\":\"int32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"connection\",\"type\":\"string\"}],\"name\":\"PacketSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"timeoutHeight\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destPort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int32\",\"name\":\"channelOrdering\",\"type\":\"int32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"connectionID\",\"type\":\"string\"}],\"name\":\"TimeoutPacket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destPort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ConnectionID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ChannelOrdering\",\"type\":\"string\"}],\"name\":\"TypeChannelClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"clientID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"clientType\",\"type\":\"string\"}],\"name\":\"TypeSubmitMisbehaviour\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"WarpIBCMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"timeoutHeight\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destPort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ack\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"connectionID\",\"type\":\"string\"}],\"name\":\"WriteAck\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"ack\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"OnAcknowledgementPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"Relayer\",\"type\":\"bytes\"}],\"name\":\"OnRecvPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"Relayer\",\"type\":\"bytes\"}],\"name\":\"OnTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"Relayer\",\"type\":\"bytes\"}],\"name\":\"OnTimeoutOnClose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofAcked\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"proofHeight\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"signer\",\"type\":\"string\"}],\"name\":\"acknowledgement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portID\",\"type\":\"string\"}],\"name\":\"bindPort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"channel\",\"type\":\"bytes\"}],\"name\":\"chanOpenInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"channel\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"counterpartyVersion\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofInit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofHeight\",\"type\":\"bytes\"}],\"name\":\"chanOpenTry\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"channelID\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofInit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofHeight\",\"type\":\"bytes\"}],\"name\":\"channelCloseConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelID\",\"type\":\"string\"}],\"name\":\"channelCloseInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"counterpartyChannelID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"counterpartyVersion\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofTry\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofHeight\",\"type\":\"bytes\"}],\"name\":\"channelOpenAck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofAck\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofHeight\",\"type\":\"bytes\"}],\"name\":\"channelOpenConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"version\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"counterpartyConnectionID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofTry\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofClient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofConsensus\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofHeight\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusHeight\",\"type\":\"bytes\"}],\"name\":\"connOpenAck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofAck\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofHeight\",\"type\":\"bytes\"}],\"name\":\"connOpenConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"counterparty\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"version\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"delayPeriod\",\"type\":\"uint32\"}],\"name\":\"connOpenInit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"connectionID\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"counterparty\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"delayPeriod\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"clientID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"counterpartyVersions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofInit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofClient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofConsensus\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofHeight\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusHeight\",\"type\":\"bytes\"}],\"name\":\"connOpenTry\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"connectionID\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientType\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusState\",\"type\":\"bytes\"}],\"name\":\"createClient\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"clientID\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelID\",\"type\":\"string\"}],\"name\":\"queryChannel\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"channel\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryChannelAll\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"channelAll\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"queryClientState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"clientState\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryClientStateAll\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"clientStateAll\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionID\",\"type\":\"string\"}],\"name\":\"queryConnection\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"connection\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryConnectionAll\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"connectionAll\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"queryConsensusState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"consensusState\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryConsensusStateAll\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"consensusStateAll\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"}],\"name\":\"queryPacketAcknowledgement\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"packetAcknowledgement\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelID\",\"type\":\"string\"}],\"name\":\"queryPacketAcknowledgements\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"packetAcknowledgements\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"}],\"name\":\"queryPacketCommitment\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"packetCommitment\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelID\",\"type\":\"string\"}],\"name\":\"queryPacketCommitments\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"packetCommitments\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelID\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"sequences\",\"type\":\"uint256[]\"}],\"name\":\"queryUnreceivedPackets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"seqs\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"height\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proofCommitment\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"proofHeight\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"signer\",\"type\":\"string\"}],\"internalType\":\"structIIBC.MsgRecvPacket\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"recvPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"channelCapability\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proofUnreceived\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"proofHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"nextSequenceRecv\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signer\",\"type\":\"string\"}],\"name\":\"timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proofUnreceived\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofClose\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"proofHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"nextSequenceRecv\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signer\",\"type\":\"string\"}],\"name\":\"timeoutOnClose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientMessage\",\"type\":\"bytes\"}],\"name\":\"updateClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"upgradedClien\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"upgradedConsState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofUpgradeClient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofUpgradeConsState\",\"type\":\"bytes\"}],\"name\":\"upgradeClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBCABI is the input ABI used to generate the binding from.
// Deprecated: Use IBCMetaData.ABI instead.
var IBCABI = IBCMetaData.ABI

// IBC is an auto generated Go binding around an Ethereum contract.
type IBC struct {
	IBCCaller     // Read-only binding to the contract
	IBCTransactor // Write-only binding to the contract
	IBCFilterer   // Log filterer for contract events
}

// IBCCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBCSession struct {
	Contract     *IBC              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBCCallerSession struct {
	Contract *IBCCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IBCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBCTransactorSession struct {
	Contract     *IBCTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBCRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBCRaw struct {
	Contract *IBC // Generic contract binding to access the raw methods on
}

// IBCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBCCallerRaw struct {
	Contract *IBCCaller // Generic read-only contract binding to access the raw methods on
}

// IBCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBCTransactorRaw struct {
	Contract *IBCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBC creates a new instance of IBC, bound to a specific deployed contract.
func NewIBC(address common.Address, backend bind.ContractBackend) (*IBC, error) {
	contract, err := bindIBC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBC{IBCCaller: IBCCaller{contract: contract}, IBCTransactor: IBCTransactor{contract: contract}, IBCFilterer: IBCFilterer{contract: contract}}, nil
}

// NewIBCCaller creates a new read-only instance of IBC, bound to a specific deployed contract.
func NewIBCCaller(address common.Address, caller bind.ContractCaller) (*IBCCaller, error) {
	contract, err := bindIBC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBCCaller{contract: contract}, nil
}

// NewIBCTransactor creates a new write-only instance of IBC, bound to a specific deployed contract.
func NewIBCTransactor(address common.Address, transactor bind.ContractTransactor) (*IBCTransactor, error) {
	contract, err := bindIBC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBCTransactor{contract: contract}, nil
}

// NewIBCFilterer creates a new log filterer instance of IBC, bound to a specific deployed contract.
func NewIBCFilterer(address common.Address, filterer bind.ContractFilterer) (*IBCFilterer, error) {
	contract, err := bindIBC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBCFilterer{contract: contract}, nil
}

// bindIBC binds a generic wrapper to an already deployed contract.
func bindIBC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBC *IBCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBC.Contract.IBCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBC *IBCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBC.Contract.IBCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBC *IBCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBC.Contract.IBCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBC *IBCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBC *IBCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBC *IBCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBC.Contract.contract.Transact(opts, method, params...)
}

// QueryChannel is a free data retrieval call binding the contract method 0x421a47dd.
//
// Solidity: function queryChannel(string portID, string channelID) view returns(bytes channel)
func (_IBC *IBCCaller) QueryChannel(opts *bind.CallOpts, portID string, channelID string) ([]byte, error) {
	var out []interface{}
	err := _IBC.contract.Call(opts, &out, "queryChannel", portID, channelID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QueryChannel is a free data retrieval call binding the contract method 0x421a47dd.
//
// Solidity: function queryChannel(string portID, string channelID) view returns(bytes channel)
func (_IBC *IBCSession) QueryChannel(portID string, channelID string) ([]byte, error) {
	return _IBC.Contract.QueryChannel(&_IBC.CallOpts, portID, channelID)
}

// QueryChannel is a free data retrieval call binding the contract method 0x421a47dd.
//
// Solidity: function queryChannel(string portID, string channelID) view returns(bytes channel)
func (_IBC *IBCCallerSession) QueryChannel(portID string, channelID string) ([]byte, error) {
	return _IBC.Contract.QueryChannel(&_IBC.CallOpts, portID, channelID)
}

// QueryChannelAll is a free data retrieval call binding the contract method 0x6c86a60d.
//
// Solidity: function queryChannelAll() view returns(bytes channelAll)
func (_IBC *IBCCaller) QueryChannelAll(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _IBC.contract.Call(opts, &out, "queryChannelAll")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QueryChannelAll is a free data retrieval call binding the contract method 0x6c86a60d.
//
// Solidity: function queryChannelAll() view returns(bytes channelAll)
func (_IBC *IBCSession) QueryChannelAll() ([]byte, error) {
	return _IBC.Contract.QueryChannelAll(&_IBC.CallOpts)
}

// QueryChannelAll is a free data retrieval call binding the contract method 0x6c86a60d.
//
// Solidity: function queryChannelAll() view returns(bytes channelAll)
func (_IBC *IBCCallerSession) QueryChannelAll() ([]byte, error) {
	return _IBC.Contract.QueryChannelAll(&_IBC.CallOpts)
}

// QueryClientState is a free data retrieval call binding the contract method 0x6b8d8117.
//
// Solidity: function queryClientState(string clientId) view returns(bytes clientState)
func (_IBC *IBCCaller) QueryClientState(opts *bind.CallOpts, clientId string) ([]byte, error) {
	var out []interface{}
	err := _IBC.contract.Call(opts, &out, "queryClientState", clientId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QueryClientState is a free data retrieval call binding the contract method 0x6b8d8117.
//
// Solidity: function queryClientState(string clientId) view returns(bytes clientState)
func (_IBC *IBCSession) QueryClientState(clientId string) ([]byte, error) {
	return _IBC.Contract.QueryClientState(&_IBC.CallOpts, clientId)
}

// QueryClientState is a free data retrieval call binding the contract method 0x6b8d8117.
//
// Solidity: function queryClientState(string clientId) view returns(bytes clientState)
func (_IBC *IBCCallerSession) QueryClientState(clientId string) ([]byte, error) {
	return _IBC.Contract.QueryClientState(&_IBC.CallOpts, clientId)
}

// QueryClientStateAll is a free data retrieval call binding the contract method 0xc2783f51.
//
// Solidity: function queryClientStateAll() view returns(bytes clientStateAll)
func (_IBC *IBCCaller) QueryClientStateAll(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _IBC.contract.Call(opts, &out, "queryClientStateAll")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QueryClientStateAll is a free data retrieval call binding the contract method 0xc2783f51.
//
// Solidity: function queryClientStateAll() view returns(bytes clientStateAll)
func (_IBC *IBCSession) QueryClientStateAll() ([]byte, error) {
	return _IBC.Contract.QueryClientStateAll(&_IBC.CallOpts)
}

// QueryClientStateAll is a free data retrieval call binding the contract method 0xc2783f51.
//
// Solidity: function queryClientStateAll() view returns(bytes clientStateAll)
func (_IBC *IBCCallerSession) QueryClientStateAll() ([]byte, error) {
	return _IBC.Contract.QueryClientStateAll(&_IBC.CallOpts)
}

// QueryConnection is a free data retrieval call binding the contract method 0x25456993.
//
// Solidity: function queryConnection(string connectionID) view returns(bytes connection)
func (_IBC *IBCCaller) QueryConnection(opts *bind.CallOpts, connectionID string) ([]byte, error) {
	var out []interface{}
	err := _IBC.contract.Call(opts, &out, "queryConnection", connectionID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QueryConnection is a free data retrieval call binding the contract method 0x25456993.
//
// Solidity: function queryConnection(string connectionID) view returns(bytes connection)
func (_IBC *IBCSession) QueryConnection(connectionID string) ([]byte, error) {
	return _IBC.Contract.QueryConnection(&_IBC.CallOpts, connectionID)
}

// QueryConnection is a free data retrieval call binding the contract method 0x25456993.
//
// Solidity: function queryConnection(string connectionID) view returns(bytes connection)
func (_IBC *IBCCallerSession) QueryConnection(connectionID string) ([]byte, error) {
	return _IBC.Contract.QueryConnection(&_IBC.CallOpts, connectionID)
}

// QueryConnectionAll is a free data retrieval call binding the contract method 0x391f521d.
//
// Solidity: function queryConnectionAll() view returns(bytes connectionAll)
func (_IBC *IBCCaller) QueryConnectionAll(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _IBC.contract.Call(opts, &out, "queryConnectionAll")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QueryConnectionAll is a free data retrieval call binding the contract method 0x391f521d.
//
// Solidity: function queryConnectionAll() view returns(bytes connectionAll)
func (_IBC *IBCSession) QueryConnectionAll() ([]byte, error) {
	return _IBC.Contract.QueryConnectionAll(&_IBC.CallOpts)
}

// QueryConnectionAll is a free data retrieval call binding the contract method 0x391f521d.
//
// Solidity: function queryConnectionAll() view returns(bytes connectionAll)
func (_IBC *IBCCallerSession) QueryConnectionAll() ([]byte, error) {
	return _IBC.Contract.QueryConnectionAll(&_IBC.CallOpts)
}

// QueryConsensusState is a free data retrieval call binding the contract method 0x936ee71f.
//
// Solidity: function queryConsensusState(string clientId) view returns(bytes consensusState)
func (_IBC *IBCCaller) QueryConsensusState(opts *bind.CallOpts, clientId string) ([]byte, error) {
	var out []interface{}
	err := _IBC.contract.Call(opts, &out, "queryConsensusState", clientId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QueryConsensusState is a free data retrieval call binding the contract method 0x936ee71f.
//
// Solidity: function queryConsensusState(string clientId) view returns(bytes consensusState)
func (_IBC *IBCSession) QueryConsensusState(clientId string) ([]byte, error) {
	return _IBC.Contract.QueryConsensusState(&_IBC.CallOpts, clientId)
}

// QueryConsensusState is a free data retrieval call binding the contract method 0x936ee71f.
//
// Solidity: function queryConsensusState(string clientId) view returns(bytes consensusState)
func (_IBC *IBCCallerSession) QueryConsensusState(clientId string) ([]byte, error) {
	return _IBC.Contract.QueryConsensusState(&_IBC.CallOpts, clientId)
}

// QueryConsensusStateAll is a free data retrieval call binding the contract method 0x72231742.
//
// Solidity: function queryConsensusStateAll() view returns(bytes consensusStateAll)
func (_IBC *IBCCaller) QueryConsensusStateAll(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _IBC.contract.Call(opts, &out, "queryConsensusStateAll")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QueryConsensusStateAll is a free data retrieval call binding the contract method 0x72231742.
//
// Solidity: function queryConsensusStateAll() view returns(bytes consensusStateAll)
func (_IBC *IBCSession) QueryConsensusStateAll() ([]byte, error) {
	return _IBC.Contract.QueryConsensusStateAll(&_IBC.CallOpts)
}

// QueryConsensusStateAll is a free data retrieval call binding the contract method 0x72231742.
//
// Solidity: function queryConsensusStateAll() view returns(bytes consensusStateAll)
func (_IBC *IBCCallerSession) QueryConsensusStateAll() ([]byte, error) {
	return _IBC.Contract.QueryConsensusStateAll(&_IBC.CallOpts)
}

// QueryPacketAcknowledgement is a free data retrieval call binding the contract method 0x37d69da6.
//
// Solidity: function queryPacketAcknowledgement(string portID, string channelID, uint256 sequence) view returns(bytes packetAcknowledgement)
func (_IBC *IBCCaller) QueryPacketAcknowledgement(opts *bind.CallOpts, portID string, channelID string, sequence *big.Int) ([]byte, error) {
	var out []interface{}
	err := _IBC.contract.Call(opts, &out, "queryPacketAcknowledgement", portID, channelID, sequence)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QueryPacketAcknowledgement is a free data retrieval call binding the contract method 0x37d69da6.
//
// Solidity: function queryPacketAcknowledgement(string portID, string channelID, uint256 sequence) view returns(bytes packetAcknowledgement)
func (_IBC *IBCSession) QueryPacketAcknowledgement(portID string, channelID string, sequence *big.Int) ([]byte, error) {
	return _IBC.Contract.QueryPacketAcknowledgement(&_IBC.CallOpts, portID, channelID, sequence)
}

// QueryPacketAcknowledgement is a free data retrieval call binding the contract method 0x37d69da6.
//
// Solidity: function queryPacketAcknowledgement(string portID, string channelID, uint256 sequence) view returns(bytes packetAcknowledgement)
func (_IBC *IBCCallerSession) QueryPacketAcknowledgement(portID string, channelID string, sequence *big.Int) ([]byte, error) {
	return _IBC.Contract.QueryPacketAcknowledgement(&_IBC.CallOpts, portID, channelID, sequence)
}

// QueryPacketAcknowledgements is a free data retrieval call binding the contract method 0xe287bfb3.
//
// Solidity: function queryPacketAcknowledgements(string portID, string channelID) view returns(bytes packetAcknowledgements)
func (_IBC *IBCCaller) QueryPacketAcknowledgements(opts *bind.CallOpts, portID string, channelID string) ([]byte, error) {
	var out []interface{}
	err := _IBC.contract.Call(opts, &out, "queryPacketAcknowledgements", portID, channelID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QueryPacketAcknowledgements is a free data retrieval call binding the contract method 0xe287bfb3.
//
// Solidity: function queryPacketAcknowledgements(string portID, string channelID) view returns(bytes packetAcknowledgements)
func (_IBC *IBCSession) QueryPacketAcknowledgements(portID string, channelID string) ([]byte, error) {
	return _IBC.Contract.QueryPacketAcknowledgements(&_IBC.CallOpts, portID, channelID)
}

// QueryPacketAcknowledgements is a free data retrieval call binding the contract method 0xe287bfb3.
//
// Solidity: function queryPacketAcknowledgements(string portID, string channelID) view returns(bytes packetAcknowledgements)
func (_IBC *IBCCallerSession) QueryPacketAcknowledgements(portID string, channelID string) ([]byte, error) {
	return _IBC.Contract.QueryPacketAcknowledgements(&_IBC.CallOpts, portID, channelID)
}

// QueryPacketCommitment is a free data retrieval call binding the contract method 0xc9b08535.
//
// Solidity: function queryPacketCommitment(string portID, string channelID, uint256 sequence) view returns(bytes packetCommitment)
func (_IBC *IBCCaller) QueryPacketCommitment(opts *bind.CallOpts, portID string, channelID string, sequence *big.Int) ([]byte, error) {
	var out []interface{}
	err := _IBC.contract.Call(opts, &out, "queryPacketCommitment", portID, channelID, sequence)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QueryPacketCommitment is a free data retrieval call binding the contract method 0xc9b08535.
//
// Solidity: function queryPacketCommitment(string portID, string channelID, uint256 sequence) view returns(bytes packetCommitment)
func (_IBC *IBCSession) QueryPacketCommitment(portID string, channelID string, sequence *big.Int) ([]byte, error) {
	return _IBC.Contract.QueryPacketCommitment(&_IBC.CallOpts, portID, channelID, sequence)
}

// QueryPacketCommitment is a free data retrieval call binding the contract method 0xc9b08535.
//
// Solidity: function queryPacketCommitment(string portID, string channelID, uint256 sequence) view returns(bytes packetCommitment)
func (_IBC *IBCCallerSession) QueryPacketCommitment(portID string, channelID string, sequence *big.Int) ([]byte, error) {
	return _IBC.Contract.QueryPacketCommitment(&_IBC.CallOpts, portID, channelID, sequence)
}

// QueryPacketCommitments is a free data retrieval call binding the contract method 0x4146b599.
//
// Solidity: function queryPacketCommitments(string portID, string channelID) view returns(bytes packetCommitments)
func (_IBC *IBCCaller) QueryPacketCommitments(opts *bind.CallOpts, portID string, channelID string) ([]byte, error) {
	var out []interface{}
	err := _IBC.contract.Call(opts, &out, "queryPacketCommitments", portID, channelID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QueryPacketCommitments is a free data retrieval call binding the contract method 0x4146b599.
//
// Solidity: function queryPacketCommitments(string portID, string channelID) view returns(bytes packetCommitments)
func (_IBC *IBCSession) QueryPacketCommitments(portID string, channelID string) ([]byte, error) {
	return _IBC.Contract.QueryPacketCommitments(&_IBC.CallOpts, portID, channelID)
}

// QueryPacketCommitments is a free data retrieval call binding the contract method 0x4146b599.
//
// Solidity: function queryPacketCommitments(string portID, string channelID) view returns(bytes packetCommitments)
func (_IBC *IBCCallerSession) QueryPacketCommitments(portID string, channelID string) ([]byte, error) {
	return _IBC.Contract.QueryPacketCommitments(&_IBC.CallOpts, portID, channelID)
}

// QueryUnreceivedPackets is a free data retrieval call binding the contract method 0xd7ade985.
//
// Solidity: function queryUnreceivedPackets(string portID, string channelID, uint256[] sequences) view returns(uint256[] seqs, (uint256,uint256) height)
func (_IBC *IBCCaller) QueryUnreceivedPackets(opts *bind.CallOpts, portID string, channelID string, sequences []*big.Int) (struct {
	Seqs   []*big.Int
	Height Height
}, error) {
	var out []interface{}
	err := _IBC.contract.Call(opts, &out, "queryUnreceivedPackets", portID, channelID, sequences)

	outstruct := new(struct {
		Seqs   []*big.Int
		Height Height
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seqs = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Height = *abi.ConvertType(out[1], new(Height)).(*Height)

	return *outstruct, err

}

// QueryUnreceivedPackets is a free data retrieval call binding the contract method 0xd7ade985.
//
// Solidity: function queryUnreceivedPackets(string portID, string channelID, uint256[] sequences) view returns(uint256[] seqs, (uint256,uint256) height)
func (_IBC *IBCSession) QueryUnreceivedPackets(portID string, channelID string, sequences []*big.Int) (struct {
	Seqs   []*big.Int
	Height Height
}, error) {
	return _IBC.Contract.QueryUnreceivedPackets(&_IBC.CallOpts, portID, channelID, sequences)
}

// QueryUnreceivedPackets is a free data retrieval call binding the contract method 0xd7ade985.
//
// Solidity: function queryUnreceivedPackets(string portID, string channelID, uint256[] sequences) view returns(uint256[] seqs, (uint256,uint256) height)
func (_IBC *IBCCallerSession) QueryUnreceivedPackets(portID string, channelID string, sequences []*big.Int) (struct {
	Seqs   []*big.Int
	Height Height
}, error) {
	return _IBC.Contract.QueryUnreceivedPackets(&_IBC.CallOpts, portID, channelID, sequences)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x2c49a978.
//
// Solidity: function OnAcknowledgementPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ack, bytes ) returns()
func (_IBC *IBCTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet Packet, ack []byte, arg2 []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "OnAcknowledgementPacket", packet, ack, arg2)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x2c49a978.
//
// Solidity: function OnAcknowledgementPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ack, bytes ) returns()
func (_IBC *IBCSession) OnAcknowledgementPacket(packet Packet, ack []byte, arg2 []byte) (*types.Transaction, error) {
	return _IBC.Contract.OnAcknowledgementPacket(&_IBC.TransactOpts, packet, ack, arg2)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x2c49a978.
//
// Solidity: function OnAcknowledgementPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ack, bytes ) returns()
func (_IBC *IBCTransactorSession) OnAcknowledgementPacket(packet Packet, ack []byte, arg2 []byte) (*types.Transaction, error) {
	return _IBC.Contract.OnAcknowledgementPacket(&_IBC.TransactOpts, packet, ack, arg2)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x85f7175c.
//
// Solidity: function OnRecvPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes Relayer) returns()
func (_IBC *IBCTransactor) OnRecvPacket(opts *bind.TransactOpts, packet Packet, Relayer []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "OnRecvPacket", packet, Relayer)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x85f7175c.
//
// Solidity: function OnRecvPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes Relayer) returns()
func (_IBC *IBCSession) OnRecvPacket(packet Packet, Relayer []byte) (*types.Transaction, error) {
	return _IBC.Contract.OnRecvPacket(&_IBC.TransactOpts, packet, Relayer)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x85f7175c.
//
// Solidity: function OnRecvPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes Relayer) returns()
func (_IBC *IBCTransactorSession) OnRecvPacket(packet Packet, Relayer []byte) (*types.Transaction, error) {
	return _IBC.Contract.OnRecvPacket(&_IBC.TransactOpts, packet, Relayer)
}

// OnTimeout is a paid mutator transaction binding the contract method 0x36b8b913.
//
// Solidity: function OnTimeout((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes Relayer) returns()
func (_IBC *IBCTransactor) OnTimeout(opts *bind.TransactOpts, packet Packet, Relayer []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "OnTimeout", packet, Relayer)
}

// OnTimeout is a paid mutator transaction binding the contract method 0x36b8b913.
//
// Solidity: function OnTimeout((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes Relayer) returns()
func (_IBC *IBCSession) OnTimeout(packet Packet, Relayer []byte) (*types.Transaction, error) {
	return _IBC.Contract.OnTimeout(&_IBC.TransactOpts, packet, Relayer)
}

// OnTimeout is a paid mutator transaction binding the contract method 0x36b8b913.
//
// Solidity: function OnTimeout((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes Relayer) returns()
func (_IBC *IBCTransactorSession) OnTimeout(packet Packet, Relayer []byte) (*types.Transaction, error) {
	return _IBC.Contract.OnTimeout(&_IBC.TransactOpts, packet, Relayer)
}

// OnTimeoutOnClose is a paid mutator transaction binding the contract method 0x128ff066.
//
// Solidity: function OnTimeoutOnClose((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes Relayer) returns()
func (_IBC *IBCTransactor) OnTimeoutOnClose(opts *bind.TransactOpts, packet Packet, Relayer []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "OnTimeoutOnClose", packet, Relayer)
}

// OnTimeoutOnClose is a paid mutator transaction binding the contract method 0x128ff066.
//
// Solidity: function OnTimeoutOnClose((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes Relayer) returns()
func (_IBC *IBCSession) OnTimeoutOnClose(packet Packet, Relayer []byte) (*types.Transaction, error) {
	return _IBC.Contract.OnTimeoutOnClose(&_IBC.TransactOpts, packet, Relayer)
}

// OnTimeoutOnClose is a paid mutator transaction binding the contract method 0x128ff066.
//
// Solidity: function OnTimeoutOnClose((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes Relayer) returns()
func (_IBC *IBCTransactorSession) OnTimeoutOnClose(packet Packet, Relayer []byte) (*types.Transaction, error) {
	return _IBC.Contract.OnTimeoutOnClose(&_IBC.TransactOpts, packet, Relayer)
}

// Acknowledgement is a paid mutator transaction binding the contract method 0xf8831420.
//
// Solidity: function acknowledgement((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes acknowledgement, bytes proofAcked, (uint256,uint256) proofHeight, string signer) returns()
func (_IBC *IBCTransactor) Acknowledgement(opts *bind.TransactOpts, packet Packet, acknowledgement []byte, proofAcked []byte, proofHeight Height, signer string) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "acknowledgement", packet, acknowledgement, proofAcked, proofHeight, signer)
}

// Acknowledgement is a paid mutator transaction binding the contract method 0xf8831420.
//
// Solidity: function acknowledgement((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes acknowledgement, bytes proofAcked, (uint256,uint256) proofHeight, string signer) returns()
func (_IBC *IBCSession) Acknowledgement(packet Packet, acknowledgement []byte, proofAcked []byte, proofHeight Height, signer string) (*types.Transaction, error) {
	return _IBC.Contract.Acknowledgement(&_IBC.TransactOpts, packet, acknowledgement, proofAcked, proofHeight, signer)
}

// Acknowledgement is a paid mutator transaction binding the contract method 0xf8831420.
//
// Solidity: function acknowledgement((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes acknowledgement, bytes proofAcked, (uint256,uint256) proofHeight, string signer) returns()
func (_IBC *IBCTransactorSession) Acknowledgement(packet Packet, acknowledgement []byte, proofAcked []byte, proofHeight Height, signer string) (*types.Transaction, error) {
	return _IBC.Contract.Acknowledgement(&_IBC.TransactOpts, packet, acknowledgement, proofAcked, proofHeight, signer)
}

// BindPort is a paid mutator transaction binding the contract method 0xc13b184f.
//
// Solidity: function bindPort(string portID) returns()
func (_IBC *IBCTransactor) BindPort(opts *bind.TransactOpts, portID string) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "bindPort", portID)
}

// BindPort is a paid mutator transaction binding the contract method 0xc13b184f.
//
// Solidity: function bindPort(string portID) returns()
func (_IBC *IBCSession) BindPort(portID string) (*types.Transaction, error) {
	return _IBC.Contract.BindPort(&_IBC.TransactOpts, portID)
}

// BindPort is a paid mutator transaction binding the contract method 0xc13b184f.
//
// Solidity: function bindPort(string portID) returns()
func (_IBC *IBCTransactorSession) BindPort(portID string) (*types.Transaction, error) {
	return _IBC.Contract.BindPort(&_IBC.TransactOpts, portID)
}

// ChanOpenInit is a paid mutator transaction binding the contract method 0xa718c941.
//
// Solidity: function chanOpenInit(string portID, bytes channel) returns()
func (_IBC *IBCTransactor) ChanOpenInit(opts *bind.TransactOpts, portID string, channel []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "chanOpenInit", portID, channel)
}

// ChanOpenInit is a paid mutator transaction binding the contract method 0xa718c941.
//
// Solidity: function chanOpenInit(string portID, bytes channel) returns()
func (_IBC *IBCSession) ChanOpenInit(portID string, channel []byte) (*types.Transaction, error) {
	return _IBC.Contract.ChanOpenInit(&_IBC.TransactOpts, portID, channel)
}

// ChanOpenInit is a paid mutator transaction binding the contract method 0xa718c941.
//
// Solidity: function chanOpenInit(string portID, bytes channel) returns()
func (_IBC *IBCTransactorSession) ChanOpenInit(portID string, channel []byte) (*types.Transaction, error) {
	return _IBC.Contract.ChanOpenInit(&_IBC.TransactOpts, portID, channel)
}

// ChanOpenTry is a paid mutator transaction binding the contract method 0x0ce2b1f6.
//
// Solidity: function chanOpenTry(string portID, bytes channel, string counterpartyVersion, bytes proofInit, bytes proofHeight) returns(string channelID)
func (_IBC *IBCTransactor) ChanOpenTry(opts *bind.TransactOpts, portID string, channel []byte, counterpartyVersion string, proofInit []byte, proofHeight []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "chanOpenTry", portID, channel, counterpartyVersion, proofInit, proofHeight)
}

// ChanOpenTry is a paid mutator transaction binding the contract method 0x0ce2b1f6.
//
// Solidity: function chanOpenTry(string portID, bytes channel, string counterpartyVersion, bytes proofInit, bytes proofHeight) returns(string channelID)
func (_IBC *IBCSession) ChanOpenTry(portID string, channel []byte, counterpartyVersion string, proofInit []byte, proofHeight []byte) (*types.Transaction, error) {
	return _IBC.Contract.ChanOpenTry(&_IBC.TransactOpts, portID, channel, counterpartyVersion, proofInit, proofHeight)
}

// ChanOpenTry is a paid mutator transaction binding the contract method 0x0ce2b1f6.
//
// Solidity: function chanOpenTry(string portID, bytes channel, string counterpartyVersion, bytes proofInit, bytes proofHeight) returns(string channelID)
func (_IBC *IBCTransactorSession) ChanOpenTry(portID string, channel []byte, counterpartyVersion string, proofInit []byte, proofHeight []byte) (*types.Transaction, error) {
	return _IBC.Contract.ChanOpenTry(&_IBC.TransactOpts, portID, channel, counterpartyVersion, proofInit, proofHeight)
}

// ChannelCloseConfirm is a paid mutator transaction binding the contract method 0x460d68fa.
//
// Solidity: function channelCloseConfirm(string portID, string channelID, bytes proofInit, bytes proofHeight) returns()
func (_IBC *IBCTransactor) ChannelCloseConfirm(opts *bind.TransactOpts, portID string, channelID string, proofInit []byte, proofHeight []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "channelCloseConfirm", portID, channelID, proofInit, proofHeight)
}

// ChannelCloseConfirm is a paid mutator transaction binding the contract method 0x460d68fa.
//
// Solidity: function channelCloseConfirm(string portID, string channelID, bytes proofInit, bytes proofHeight) returns()
func (_IBC *IBCSession) ChannelCloseConfirm(portID string, channelID string, proofInit []byte, proofHeight []byte) (*types.Transaction, error) {
	return _IBC.Contract.ChannelCloseConfirm(&_IBC.TransactOpts, portID, channelID, proofInit, proofHeight)
}

// ChannelCloseConfirm is a paid mutator transaction binding the contract method 0x460d68fa.
//
// Solidity: function channelCloseConfirm(string portID, string channelID, bytes proofInit, bytes proofHeight) returns()
func (_IBC *IBCTransactorSession) ChannelCloseConfirm(portID string, channelID string, proofInit []byte, proofHeight []byte) (*types.Transaction, error) {
	return _IBC.Contract.ChannelCloseConfirm(&_IBC.TransactOpts, portID, channelID, proofInit, proofHeight)
}

// ChannelCloseInit is a paid mutator transaction binding the contract method 0x7eb320da.
//
// Solidity: function channelCloseInit(string portID, string channelID) returns()
func (_IBC *IBCTransactor) ChannelCloseInit(opts *bind.TransactOpts, portID string, channelID string) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "channelCloseInit", portID, channelID)
}

// ChannelCloseInit is a paid mutator transaction binding the contract method 0x7eb320da.
//
// Solidity: function channelCloseInit(string portID, string channelID) returns()
func (_IBC *IBCSession) ChannelCloseInit(portID string, channelID string) (*types.Transaction, error) {
	return _IBC.Contract.ChannelCloseInit(&_IBC.TransactOpts, portID, channelID)
}

// ChannelCloseInit is a paid mutator transaction binding the contract method 0x7eb320da.
//
// Solidity: function channelCloseInit(string portID, string channelID) returns()
func (_IBC *IBCTransactorSession) ChannelCloseInit(portID string, channelID string) (*types.Transaction, error) {
	return _IBC.Contract.ChannelCloseInit(&_IBC.TransactOpts, portID, channelID)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0xbd6f4bde.
//
// Solidity: function channelOpenAck(string portID, string channelID, string counterpartyChannelID, string counterpartyVersion, bytes proofTry, bytes proofHeight) returns()
func (_IBC *IBCTransactor) ChannelOpenAck(opts *bind.TransactOpts, portID string, channelID string, counterpartyChannelID string, counterpartyVersion string, proofTry []byte, proofHeight []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "channelOpenAck", portID, channelID, counterpartyChannelID, counterpartyVersion, proofTry, proofHeight)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0xbd6f4bde.
//
// Solidity: function channelOpenAck(string portID, string channelID, string counterpartyChannelID, string counterpartyVersion, bytes proofTry, bytes proofHeight) returns()
func (_IBC *IBCSession) ChannelOpenAck(portID string, channelID string, counterpartyChannelID string, counterpartyVersion string, proofTry []byte, proofHeight []byte) (*types.Transaction, error) {
	return _IBC.Contract.ChannelOpenAck(&_IBC.TransactOpts, portID, channelID, counterpartyChannelID, counterpartyVersion, proofTry, proofHeight)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0xbd6f4bde.
//
// Solidity: function channelOpenAck(string portID, string channelID, string counterpartyChannelID, string counterpartyVersion, bytes proofTry, bytes proofHeight) returns()
func (_IBC *IBCTransactorSession) ChannelOpenAck(portID string, channelID string, counterpartyChannelID string, counterpartyVersion string, proofTry []byte, proofHeight []byte) (*types.Transaction, error) {
	return _IBC.Contract.ChannelOpenAck(&_IBC.TransactOpts, portID, channelID, counterpartyChannelID, counterpartyVersion, proofTry, proofHeight)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x9c110621.
//
// Solidity: function channelOpenConfirm(string portID, string channelID, bytes proofAck, bytes proofHeight) returns()
func (_IBC *IBCTransactor) ChannelOpenConfirm(opts *bind.TransactOpts, portID string, channelID string, proofAck []byte, proofHeight []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "channelOpenConfirm", portID, channelID, proofAck, proofHeight)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x9c110621.
//
// Solidity: function channelOpenConfirm(string portID, string channelID, bytes proofAck, bytes proofHeight) returns()
func (_IBC *IBCSession) ChannelOpenConfirm(portID string, channelID string, proofAck []byte, proofHeight []byte) (*types.Transaction, error) {
	return _IBC.Contract.ChannelOpenConfirm(&_IBC.TransactOpts, portID, channelID, proofAck, proofHeight)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x9c110621.
//
// Solidity: function channelOpenConfirm(string portID, string channelID, bytes proofAck, bytes proofHeight) returns()
func (_IBC *IBCTransactorSession) ChannelOpenConfirm(portID string, channelID string, proofAck []byte, proofHeight []byte) (*types.Transaction, error) {
	return _IBC.Contract.ChannelOpenConfirm(&_IBC.TransactOpts, portID, channelID, proofAck, proofHeight)
}

// ConnOpenAck is a paid mutator transaction binding the contract method 0x528d5cd3.
//
// Solidity: function connOpenAck(string connectionID, bytes clientState, bytes version, bytes counterpartyConnectionID, bytes proofTry, bytes proofClient, bytes proofConsensus, bytes proofHeight, bytes consensusHeight) returns()
func (_IBC *IBCTransactor) ConnOpenAck(opts *bind.TransactOpts, connectionID string, clientState []byte, version []byte, counterpartyConnectionID []byte, proofTry []byte, proofClient []byte, proofConsensus []byte, proofHeight []byte, consensusHeight []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "connOpenAck", connectionID, clientState, version, counterpartyConnectionID, proofTry, proofClient, proofConsensus, proofHeight, consensusHeight)
}

// ConnOpenAck is a paid mutator transaction binding the contract method 0x528d5cd3.
//
// Solidity: function connOpenAck(string connectionID, bytes clientState, bytes version, bytes counterpartyConnectionID, bytes proofTry, bytes proofClient, bytes proofConsensus, bytes proofHeight, bytes consensusHeight) returns()
func (_IBC *IBCSession) ConnOpenAck(connectionID string, clientState []byte, version []byte, counterpartyConnectionID []byte, proofTry []byte, proofClient []byte, proofConsensus []byte, proofHeight []byte, consensusHeight []byte) (*types.Transaction, error) {
	return _IBC.Contract.ConnOpenAck(&_IBC.TransactOpts, connectionID, clientState, version, counterpartyConnectionID, proofTry, proofClient, proofConsensus, proofHeight, consensusHeight)
}

// ConnOpenAck is a paid mutator transaction binding the contract method 0x528d5cd3.
//
// Solidity: function connOpenAck(string connectionID, bytes clientState, bytes version, bytes counterpartyConnectionID, bytes proofTry, bytes proofClient, bytes proofConsensus, bytes proofHeight, bytes consensusHeight) returns()
func (_IBC *IBCTransactorSession) ConnOpenAck(connectionID string, clientState []byte, version []byte, counterpartyConnectionID []byte, proofTry []byte, proofClient []byte, proofConsensus []byte, proofHeight []byte, consensusHeight []byte) (*types.Transaction, error) {
	return _IBC.Contract.ConnOpenAck(&_IBC.TransactOpts, connectionID, clientState, version, counterpartyConnectionID, proofTry, proofClient, proofConsensus, proofHeight, consensusHeight)
}

// ConnOpenConfirm is a paid mutator transaction binding the contract method 0x45870d5e.
//
// Solidity: function connOpenConfirm(string connectionID, bytes proofAck, bytes proofHeight) returns()
func (_IBC *IBCTransactor) ConnOpenConfirm(opts *bind.TransactOpts, connectionID string, proofAck []byte, proofHeight []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "connOpenConfirm", connectionID, proofAck, proofHeight)
}

// ConnOpenConfirm is a paid mutator transaction binding the contract method 0x45870d5e.
//
// Solidity: function connOpenConfirm(string connectionID, bytes proofAck, bytes proofHeight) returns()
func (_IBC *IBCSession) ConnOpenConfirm(connectionID string, proofAck []byte, proofHeight []byte) (*types.Transaction, error) {
	return _IBC.Contract.ConnOpenConfirm(&_IBC.TransactOpts, connectionID, proofAck, proofHeight)
}

// ConnOpenConfirm is a paid mutator transaction binding the contract method 0x45870d5e.
//
// Solidity: function connOpenConfirm(string connectionID, bytes proofAck, bytes proofHeight) returns()
func (_IBC *IBCTransactorSession) ConnOpenConfirm(connectionID string, proofAck []byte, proofHeight []byte) (*types.Transaction, error) {
	return _IBC.Contract.ConnOpenConfirm(&_IBC.TransactOpts, connectionID, proofAck, proofHeight)
}

// ConnOpenInit is a paid mutator transaction binding the contract method 0xd198062b.
//
// Solidity: function connOpenInit(string clientID, bytes counterparty, bytes version, uint32 delayPeriod) returns(string connectionID)
func (_IBC *IBCTransactor) ConnOpenInit(opts *bind.TransactOpts, clientID string, counterparty []byte, version []byte, delayPeriod uint32) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "connOpenInit", clientID, counterparty, version, delayPeriod)
}

// ConnOpenInit is a paid mutator transaction binding the contract method 0xd198062b.
//
// Solidity: function connOpenInit(string clientID, bytes counterparty, bytes version, uint32 delayPeriod) returns(string connectionID)
func (_IBC *IBCSession) ConnOpenInit(clientID string, counterparty []byte, version []byte, delayPeriod uint32) (*types.Transaction, error) {
	return _IBC.Contract.ConnOpenInit(&_IBC.TransactOpts, clientID, counterparty, version, delayPeriod)
}

// ConnOpenInit is a paid mutator transaction binding the contract method 0xd198062b.
//
// Solidity: function connOpenInit(string clientID, bytes counterparty, bytes version, uint32 delayPeriod) returns(string connectionID)
func (_IBC *IBCTransactorSession) ConnOpenInit(clientID string, counterparty []byte, version []byte, delayPeriod uint32) (*types.Transaction, error) {
	return _IBC.Contract.ConnOpenInit(&_IBC.TransactOpts, clientID, counterparty, version, delayPeriod)
}

// ConnOpenTry is a paid mutator transaction binding the contract method 0x6954535e.
//
// Solidity: function connOpenTry(bytes counterparty, uint32 delayPeriod, string clientID, bytes clientState, bytes counterpartyVersions, bytes proofInit, bytes proofClient, bytes proofConsensus, bytes proofHeight, bytes consensusHeight) returns(string connectionID)
func (_IBC *IBCTransactor) ConnOpenTry(opts *bind.TransactOpts, counterparty []byte, delayPeriod uint32, clientID string, clientState []byte, counterpartyVersions []byte, proofInit []byte, proofClient []byte, proofConsensus []byte, proofHeight []byte, consensusHeight []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "connOpenTry", counterparty, delayPeriod, clientID, clientState, counterpartyVersions, proofInit, proofClient, proofConsensus, proofHeight, consensusHeight)
}

// ConnOpenTry is a paid mutator transaction binding the contract method 0x6954535e.
//
// Solidity: function connOpenTry(bytes counterparty, uint32 delayPeriod, string clientID, bytes clientState, bytes counterpartyVersions, bytes proofInit, bytes proofClient, bytes proofConsensus, bytes proofHeight, bytes consensusHeight) returns(string connectionID)
func (_IBC *IBCSession) ConnOpenTry(counterparty []byte, delayPeriod uint32, clientID string, clientState []byte, counterpartyVersions []byte, proofInit []byte, proofClient []byte, proofConsensus []byte, proofHeight []byte, consensusHeight []byte) (*types.Transaction, error) {
	return _IBC.Contract.ConnOpenTry(&_IBC.TransactOpts, counterparty, delayPeriod, clientID, clientState, counterpartyVersions, proofInit, proofClient, proofConsensus, proofHeight, consensusHeight)
}

// ConnOpenTry is a paid mutator transaction binding the contract method 0x6954535e.
//
// Solidity: function connOpenTry(bytes counterparty, uint32 delayPeriod, string clientID, bytes clientState, bytes counterpartyVersions, bytes proofInit, bytes proofClient, bytes proofConsensus, bytes proofHeight, bytes consensusHeight) returns(string connectionID)
func (_IBC *IBCTransactorSession) ConnOpenTry(counterparty []byte, delayPeriod uint32, clientID string, clientState []byte, counterpartyVersions []byte, proofInit []byte, proofClient []byte, proofConsensus []byte, proofHeight []byte, consensusHeight []byte) (*types.Transaction, error) {
	return _IBC.Contract.ConnOpenTry(&_IBC.TransactOpts, counterparty, delayPeriod, clientID, clientState, counterpartyVersions, proofInit, proofClient, proofConsensus, proofHeight, consensusHeight)
}

// CreateClient is a paid mutator transaction binding the contract method 0x2629636b.
//
// Solidity: function createClient(string clientType, bytes clientState, bytes consensusState) returns(string clientID)
func (_IBC *IBCTransactor) CreateClient(opts *bind.TransactOpts, clientType string, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "createClient", clientType, clientState, consensusState)
}

// CreateClient is a paid mutator transaction binding the contract method 0x2629636b.
//
// Solidity: function createClient(string clientType, bytes clientState, bytes consensusState) returns(string clientID)
func (_IBC *IBCSession) CreateClient(clientType string, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _IBC.Contract.CreateClient(&_IBC.TransactOpts, clientType, clientState, consensusState)
}

// CreateClient is a paid mutator transaction binding the contract method 0x2629636b.
//
// Solidity: function createClient(string clientType, bytes clientState, bytes consensusState) returns(string clientID)
func (_IBC *IBCTransactorSession) CreateClient(clientType string, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _IBC.Contract.CreateClient(&_IBC.TransactOpts, clientType, clientState, consensusState)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x15d1edac.
//
// Solidity: function recvPacket(((uint256,string,string,string,string,bytes,(uint256,uint256),uint256),bytes,(uint256,uint256),string) message) returns()
func (_IBC *IBCTransactor) RecvPacket(opts *bind.TransactOpts, message IIBCMsgRecvPacket) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "recvPacket", message)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x15d1edac.
//
// Solidity: function recvPacket(((uint256,string,string,string,string,bytes,(uint256,uint256),uint256),bytes,(uint256,uint256),string) message) returns()
func (_IBC *IBCSession) RecvPacket(message IIBCMsgRecvPacket) (*types.Transaction, error) {
	return _IBC.Contract.RecvPacket(&_IBC.TransactOpts, message)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x15d1edac.
//
// Solidity: function recvPacket(((uint256,string,string,string,string,bytes,(uint256,uint256),uint256),bytes,(uint256,uint256),string) message) returns()
func (_IBC *IBCTransactorSession) RecvPacket(message IIBCMsgRecvPacket) (*types.Transaction, error) {
	return _IBC.Contract.RecvPacket(&_IBC.TransactOpts, message)
}

// SendPacket is a paid mutator transaction binding the contract method 0x6052bf6f.
//
// Solidity: function sendPacket(uint256 channelCapability, string sourcePort, string sourceChannel, (uint256,uint256) timeoutHeight, uint256 timeoutTimestamp, bytes data) returns()
func (_IBC *IBCTransactor) SendPacket(opts *bind.TransactOpts, channelCapability *big.Int, sourcePort string, sourceChannel string, timeoutHeight Height, timeoutTimestamp *big.Int, data []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "sendPacket", channelCapability, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

// SendPacket is a paid mutator transaction binding the contract method 0x6052bf6f.
//
// Solidity: function sendPacket(uint256 channelCapability, string sourcePort, string sourceChannel, (uint256,uint256) timeoutHeight, uint256 timeoutTimestamp, bytes data) returns()
func (_IBC *IBCSession) SendPacket(channelCapability *big.Int, sourcePort string, sourceChannel string, timeoutHeight Height, timeoutTimestamp *big.Int, data []byte) (*types.Transaction, error) {
	return _IBC.Contract.SendPacket(&_IBC.TransactOpts, channelCapability, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

// SendPacket is a paid mutator transaction binding the contract method 0x6052bf6f.
//
// Solidity: function sendPacket(uint256 channelCapability, string sourcePort, string sourceChannel, (uint256,uint256) timeoutHeight, uint256 timeoutTimestamp, bytes data) returns()
func (_IBC *IBCTransactorSession) SendPacket(channelCapability *big.Int, sourcePort string, sourceChannel string, timeoutHeight Height, timeoutTimestamp *big.Int, data []byte) (*types.Transaction, error) {
	return _IBC.Contract.SendPacket(&_IBC.TransactOpts, channelCapability, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

// Timeout is a paid mutator transaction binding the contract method 0x8883ff39.
//
// Solidity: function timeout((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes proofUnreceived, (uint256,uint256) proofHeight, uint256 nextSequenceRecv, string signer) returns()
func (_IBC *IBCTransactor) Timeout(opts *bind.TransactOpts, packet Packet, proofUnreceived []byte, proofHeight Height, nextSequenceRecv *big.Int, signer string) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "timeout", packet, proofUnreceived, proofHeight, nextSequenceRecv, signer)
}

// Timeout is a paid mutator transaction binding the contract method 0x8883ff39.
//
// Solidity: function timeout((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes proofUnreceived, (uint256,uint256) proofHeight, uint256 nextSequenceRecv, string signer) returns()
func (_IBC *IBCSession) Timeout(packet Packet, proofUnreceived []byte, proofHeight Height, nextSequenceRecv *big.Int, signer string) (*types.Transaction, error) {
	return _IBC.Contract.Timeout(&_IBC.TransactOpts, packet, proofUnreceived, proofHeight, nextSequenceRecv, signer)
}

// Timeout is a paid mutator transaction binding the contract method 0x8883ff39.
//
// Solidity: function timeout((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes proofUnreceived, (uint256,uint256) proofHeight, uint256 nextSequenceRecv, string signer) returns()
func (_IBC *IBCTransactorSession) Timeout(packet Packet, proofUnreceived []byte, proofHeight Height, nextSequenceRecv *big.Int, signer string) (*types.Transaction, error) {
	return _IBC.Contract.Timeout(&_IBC.TransactOpts, packet, proofUnreceived, proofHeight, nextSequenceRecv, signer)
}

// TimeoutOnClose is a paid mutator transaction binding the contract method 0xc519baa9.
//
// Solidity: function timeoutOnClose((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes proofUnreceived, bytes proofClose, (uint256,uint256) proofHeight, uint256 nextSequenceRecv, string signer) returns()
func (_IBC *IBCTransactor) TimeoutOnClose(opts *bind.TransactOpts, packet Packet, proofUnreceived []byte, proofClose []byte, proofHeight Height, nextSequenceRecv *big.Int, signer string) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "timeoutOnClose", packet, proofUnreceived, proofClose, proofHeight, nextSequenceRecv, signer)
}

// TimeoutOnClose is a paid mutator transaction binding the contract method 0xc519baa9.
//
// Solidity: function timeoutOnClose((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes proofUnreceived, bytes proofClose, (uint256,uint256) proofHeight, uint256 nextSequenceRecv, string signer) returns()
func (_IBC *IBCSession) TimeoutOnClose(packet Packet, proofUnreceived []byte, proofClose []byte, proofHeight Height, nextSequenceRecv *big.Int, signer string) (*types.Transaction, error) {
	return _IBC.Contract.TimeoutOnClose(&_IBC.TransactOpts, packet, proofUnreceived, proofClose, proofHeight, nextSequenceRecv, signer)
}

// TimeoutOnClose is a paid mutator transaction binding the contract method 0xc519baa9.
//
// Solidity: function timeoutOnClose((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes proofUnreceived, bytes proofClose, (uint256,uint256) proofHeight, uint256 nextSequenceRecv, string signer) returns()
func (_IBC *IBCTransactorSession) TimeoutOnClose(packet Packet, proofUnreceived []byte, proofClose []byte, proofHeight Height, nextSequenceRecv *big.Int, signer string) (*types.Transaction, error) {
	return _IBC.Contract.TimeoutOnClose(&_IBC.TransactOpts, packet, proofUnreceived, proofClose, proofHeight, nextSequenceRecv, signer)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x6fbf8079.
//
// Solidity: function updateClient(string clientID, bytes clientMessage) returns()
func (_IBC *IBCTransactor) UpdateClient(opts *bind.TransactOpts, clientID string, clientMessage []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "updateClient", clientID, clientMessage)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x6fbf8079.
//
// Solidity: function updateClient(string clientID, bytes clientMessage) returns()
func (_IBC *IBCSession) UpdateClient(clientID string, clientMessage []byte) (*types.Transaction, error) {
	return _IBC.Contract.UpdateClient(&_IBC.TransactOpts, clientID, clientMessage)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x6fbf8079.
//
// Solidity: function updateClient(string clientID, bytes clientMessage) returns()
func (_IBC *IBCTransactorSession) UpdateClient(clientID string, clientMessage []byte) (*types.Transaction, error) {
	return _IBC.Contract.UpdateClient(&_IBC.TransactOpts, clientID, clientMessage)
}

// UpgradeClient is a paid mutator transaction binding the contract method 0xed063c16.
//
// Solidity: function upgradeClient(string clientID, bytes upgradedClien, bytes upgradedConsState, bytes proofUpgradeClient, bytes proofUpgradeConsState) returns()
func (_IBC *IBCTransactor) UpgradeClient(opts *bind.TransactOpts, clientID string, upgradedClien []byte, upgradedConsState []byte, proofUpgradeClient []byte, proofUpgradeConsState []byte) (*types.Transaction, error) {
	return _IBC.contract.Transact(opts, "upgradeClient", clientID, upgradedClien, upgradedConsState, proofUpgradeClient, proofUpgradeConsState)
}

// UpgradeClient is a paid mutator transaction binding the contract method 0xed063c16.
//
// Solidity: function upgradeClient(string clientID, bytes upgradedClien, bytes upgradedConsState, bytes proofUpgradeClient, bytes proofUpgradeConsState) returns()
func (_IBC *IBCSession) UpgradeClient(clientID string, upgradedClien []byte, upgradedConsState []byte, proofUpgradeClient []byte, proofUpgradeConsState []byte) (*types.Transaction, error) {
	return _IBC.Contract.UpgradeClient(&_IBC.TransactOpts, clientID, upgradedClien, upgradedConsState, proofUpgradeClient, proofUpgradeConsState)
}

// UpgradeClient is a paid mutator transaction binding the contract method 0xed063c16.
//
// Solidity: function upgradeClient(string clientID, bytes upgradedClien, bytes upgradedConsState, bytes proofUpgradeClient, bytes proofUpgradeConsState) returns()
func (_IBC *IBCTransactorSession) UpgradeClient(clientID string, upgradedClien []byte, upgradedConsState []byte, proofUpgradeClient []byte, proofUpgradeConsState []byte) (*types.Transaction, error) {
	return _IBC.Contract.UpgradeClient(&_IBC.TransactOpts, clientID, upgradedClien, upgradedConsState, proofUpgradeClient, proofUpgradeConsState)
}

// IBCAckPacketIterator is returned from FilterAckPacket and is used to iterate over the raw logs and unpacked data for AckPacket events raised by the IBC contract.
type IBCAckPacketIterator struct {
	Event *IBCAckPacket // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCAckPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCAckPacket)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCAckPacket)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCAckPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCAckPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCAckPacket represents a AckPacket event raised by the IBC contract.
type IBCAckPacket struct {
	TimeoutHeight    string
	TimeoutTimestamp *big.Int
	Sequence         *big.Int
	SourcePort       string
	SourceChannel    string
	DestPort         string
	DestChannel      string
	ChannelOrdering  int32
	ConnectionID     string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAckPacket is a free log retrieval operation binding the contract event 0x12c19a09144e24c66aaad5927966d7d605e4d79fae215cd7ab1880c9b1a440b9.
//
// Solidity: event AckPacket(string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, int32 channelOrdering, string connectionID)
func (_IBC *IBCFilterer) FilterAckPacket(opts *bind.FilterOpts) (*IBCAckPacketIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "AckPacket")
	if err != nil {
		return nil, err
	}
	return &IBCAckPacketIterator{contract: _IBC.contract, event: "AckPacket", logs: logs, sub: sub}, nil
}

// WatchAckPacket is a free log subscription operation binding the contract event 0x12c19a09144e24c66aaad5927966d7d605e4d79fae215cd7ab1880c9b1a440b9.
//
// Solidity: event AckPacket(string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, int32 channelOrdering, string connectionID)
func (_IBC *IBCFilterer) WatchAckPacket(opts *bind.WatchOpts, sink chan<- *IBCAckPacket) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "AckPacket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCAckPacket)
				if err := _IBC.contract.UnpackLog(event, "AckPacket", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAckPacket is a log parse operation binding the contract event 0x12c19a09144e24c66aaad5927966d7d605e4d79fae215cd7ab1880c9b1a440b9.
//
// Solidity: event AckPacket(string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, int32 channelOrdering, string connectionID)
func (_IBC *IBCFilterer) ParseAckPacket(log types.Log) (*IBCAckPacket, error) {
	event := new(IBCAckPacket)
	if err := _IBC.contract.UnpackLog(event, "AckPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCAcknowledgementErrorIterator is returned from FilterAcknowledgementError and is used to iterate over the raw logs and unpacked data for AcknowledgementError events raised by the IBC contract.
type IBCAcknowledgementErrorIterator struct {
	Event *IBCAcknowledgementError // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCAcknowledgementErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCAcknowledgementError)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCAcknowledgementError)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCAcknowledgementErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCAcknowledgementErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCAcknowledgementError represents a AcknowledgementError event raised by the IBC contract.
type IBCAcknowledgementError struct {
	Data             []byte
	TimeoutHeight    string
	TimeoutTimestamp *big.Int
	Sequence         *big.Int
	SourcePort       string
	SourceChannel    string
	DestPort         string
	DestChannel      string
	Error            string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAcknowledgementError is a free log retrieval operation binding the contract event 0x59fac6d3fb2b5eb86e41d879f3d57a504ce59231004dfa51063069b12ec58715.
//
// Solidity: event AcknowledgementError(bytes data, string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, string error)
func (_IBC *IBCFilterer) FilterAcknowledgementError(opts *bind.FilterOpts) (*IBCAcknowledgementErrorIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "AcknowledgementError")
	if err != nil {
		return nil, err
	}
	return &IBCAcknowledgementErrorIterator{contract: _IBC.contract, event: "AcknowledgementError", logs: logs, sub: sub}, nil
}

// WatchAcknowledgementError is a free log subscription operation binding the contract event 0x59fac6d3fb2b5eb86e41d879f3d57a504ce59231004dfa51063069b12ec58715.
//
// Solidity: event AcknowledgementError(bytes data, string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, string error)
func (_IBC *IBCFilterer) WatchAcknowledgementError(opts *bind.WatchOpts, sink chan<- *IBCAcknowledgementError) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "AcknowledgementError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCAcknowledgementError)
				if err := _IBC.contract.UnpackLog(event, "AcknowledgementError", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAcknowledgementError is a log parse operation binding the contract event 0x59fac6d3fb2b5eb86e41d879f3d57a504ce59231004dfa51063069b12ec58715.
//
// Solidity: event AcknowledgementError(bytes data, string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, string error)
func (_IBC *IBCFilterer) ParseAcknowledgementError(log types.Log) (*IBCAcknowledgementError, error) {
	event := new(IBCAcknowledgementError)
	if err := _IBC.contract.UnpackLog(event, "AcknowledgementError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCChannelCloseConfirmIterator is returned from FilterChannelCloseConfirm and is used to iterate over the raw logs and unpacked data for ChannelCloseConfirm events raised by the IBC contract.
type IBCChannelCloseConfirmIterator struct {
	Event *IBCChannelCloseConfirm // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCChannelCloseConfirmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCChannelCloseConfirm)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCChannelCloseConfirm)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCChannelCloseConfirmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCChannelCloseConfirmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCChannelCloseConfirm represents a ChannelCloseConfirm event raised by the IBC contract.
type IBCChannelCloseConfirm struct {
	ConnectionId          string
	ChannelId             string
	PortId                string
	CounterpartyChannelId string
	CounterpartyPortID    string
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterChannelCloseConfirm is a free log retrieval operation binding the contract event 0x1d27827947f32db531c2d0a11a83e392e9391cf32071f1716bc53c3df605b637.
//
// Solidity: event ChannelCloseConfirm(string connectionId, string channelId, string portId, string counterpartyChannelId, string counterpartyPortID)
func (_IBC *IBCFilterer) FilterChannelCloseConfirm(opts *bind.FilterOpts) (*IBCChannelCloseConfirmIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "ChannelCloseConfirm")
	if err != nil {
		return nil, err
	}
	return &IBCChannelCloseConfirmIterator{contract: _IBC.contract, event: "ChannelCloseConfirm", logs: logs, sub: sub}, nil
}

// WatchChannelCloseConfirm is a free log subscription operation binding the contract event 0x1d27827947f32db531c2d0a11a83e392e9391cf32071f1716bc53c3df605b637.
//
// Solidity: event ChannelCloseConfirm(string connectionId, string channelId, string portId, string counterpartyChannelId, string counterpartyPortID)
func (_IBC *IBCFilterer) WatchChannelCloseConfirm(opts *bind.WatchOpts, sink chan<- *IBCChannelCloseConfirm) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "ChannelCloseConfirm")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCChannelCloseConfirm)
				if err := _IBC.contract.UnpackLog(event, "ChannelCloseConfirm", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChannelCloseConfirm is a log parse operation binding the contract event 0x1d27827947f32db531c2d0a11a83e392e9391cf32071f1716bc53c3df605b637.
//
// Solidity: event ChannelCloseConfirm(string connectionId, string channelId, string portId, string counterpartyChannelId, string counterpartyPortID)
func (_IBC *IBCFilterer) ParseChannelCloseConfirm(log types.Log) (*IBCChannelCloseConfirm, error) {
	event := new(IBCChannelCloseConfirm)
	if err := _IBC.contract.UnpackLog(event, "ChannelCloseConfirm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCChannelCloseInitIterator is returned from FilterChannelCloseInit and is used to iterate over the raw logs and unpacked data for ChannelCloseInit events raised by the IBC contract.
type IBCChannelCloseInitIterator struct {
	Event *IBCChannelCloseInit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCChannelCloseInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCChannelCloseInit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCChannelCloseInit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCChannelCloseInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCChannelCloseInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCChannelCloseInit represents a ChannelCloseInit event raised by the IBC contract.
type IBCChannelCloseInit struct {
	ConnectionId          string
	ChannelId             string
	PortId                string
	CounterpartyChannelId string
	CounterpartyPortID    string
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterChannelCloseInit is a free log retrieval operation binding the contract event 0x645976448b76cf17132a7f0f96d505a70aa349bc7753973035352feb57901375.
//
// Solidity: event ChannelCloseInit(string connectionId, string channelId, string portId, string counterpartyChannelId, string counterpartyPortID)
func (_IBC *IBCFilterer) FilterChannelCloseInit(opts *bind.FilterOpts) (*IBCChannelCloseInitIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "ChannelCloseInit")
	if err != nil {
		return nil, err
	}
	return &IBCChannelCloseInitIterator{contract: _IBC.contract, event: "ChannelCloseInit", logs: logs, sub: sub}, nil
}

// WatchChannelCloseInit is a free log subscription operation binding the contract event 0x645976448b76cf17132a7f0f96d505a70aa349bc7753973035352feb57901375.
//
// Solidity: event ChannelCloseInit(string connectionId, string channelId, string portId, string counterpartyChannelId, string counterpartyPortID)
func (_IBC *IBCFilterer) WatchChannelCloseInit(opts *bind.WatchOpts, sink chan<- *IBCChannelCloseInit) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "ChannelCloseInit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCChannelCloseInit)
				if err := _IBC.contract.UnpackLog(event, "ChannelCloseInit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChannelCloseInit is a log parse operation binding the contract event 0x645976448b76cf17132a7f0f96d505a70aa349bc7753973035352feb57901375.
//
// Solidity: event ChannelCloseInit(string connectionId, string channelId, string portId, string counterpartyChannelId, string counterpartyPortID)
func (_IBC *IBCFilterer) ParseChannelCloseInit(log types.Log) (*IBCChannelCloseInit, error) {
	event := new(IBCChannelCloseInit)
	if err := _IBC.contract.UnpackLog(event, "ChannelCloseInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCChannelOpenAckIterator is returned from FilterChannelOpenAck and is used to iterate over the raw logs and unpacked data for ChannelOpenAck events raised by the IBC contract.
type IBCChannelOpenAckIterator struct {
	Event *IBCChannelOpenAck // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCChannelOpenAckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCChannelOpenAck)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCChannelOpenAck)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCChannelOpenAckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCChannelOpenAckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCChannelOpenAck represents a ChannelOpenAck event raised by the IBC contract.
type IBCChannelOpenAck struct {
	ConnectionId          string
	ChannelId             string
	PortId                string
	CounterpartyChannelId string
	CounterpartyPortID    string
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenAck is a free log retrieval operation binding the contract event 0xe9342577bf02f748ba783edd9094f8e93b2ef7face9bc7478d7b30358ddeef6f.
//
// Solidity: event ChannelOpenAck(string connectionId, string channelId, string portId, string counterpartyChannelId, string counterpartyPortID)
func (_IBC *IBCFilterer) FilterChannelOpenAck(opts *bind.FilterOpts) (*IBCChannelOpenAckIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "ChannelOpenAck")
	if err != nil {
		return nil, err
	}
	return &IBCChannelOpenAckIterator{contract: _IBC.contract, event: "ChannelOpenAck", logs: logs, sub: sub}, nil
}

// WatchChannelOpenAck is a free log subscription operation binding the contract event 0xe9342577bf02f748ba783edd9094f8e93b2ef7face9bc7478d7b30358ddeef6f.
//
// Solidity: event ChannelOpenAck(string connectionId, string channelId, string portId, string counterpartyChannelId, string counterpartyPortID)
func (_IBC *IBCFilterer) WatchChannelOpenAck(opts *bind.WatchOpts, sink chan<- *IBCChannelOpenAck) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "ChannelOpenAck")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCChannelOpenAck)
				if err := _IBC.contract.UnpackLog(event, "ChannelOpenAck", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChannelOpenAck is a log parse operation binding the contract event 0xe9342577bf02f748ba783edd9094f8e93b2ef7face9bc7478d7b30358ddeef6f.
//
// Solidity: event ChannelOpenAck(string connectionId, string channelId, string portId, string counterpartyChannelId, string counterpartyPortID)
func (_IBC *IBCFilterer) ParseChannelOpenAck(log types.Log) (*IBCChannelOpenAck, error) {
	event := new(IBCChannelOpenAck)
	if err := _IBC.contract.UnpackLog(event, "ChannelOpenAck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCChannelOpenConfirmIterator is returned from FilterChannelOpenConfirm and is used to iterate over the raw logs and unpacked data for ChannelOpenConfirm events raised by the IBC contract.
type IBCChannelOpenConfirmIterator struct {
	Event *IBCChannelOpenConfirm // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCChannelOpenConfirmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCChannelOpenConfirm)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCChannelOpenConfirm)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCChannelOpenConfirmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCChannelOpenConfirmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCChannelOpenConfirm represents a ChannelOpenConfirm event raised by the IBC contract.
type IBCChannelOpenConfirm struct {
	ConnectionId          string
	ChannelId             string
	PortId                string
	CounterpartyChannelId string
	CounterpartyPortID    string
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenConfirm is a free log retrieval operation binding the contract event 0xcccb79544f2a910ecd04c3bd96f870be6f5c74e0d00c18443c25ecf7b9800918.
//
// Solidity: event ChannelOpenConfirm(string connectionId, string channelId, string portId, string counterpartyChannelId, string counterpartyPortID)
func (_IBC *IBCFilterer) FilterChannelOpenConfirm(opts *bind.FilterOpts) (*IBCChannelOpenConfirmIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "ChannelOpenConfirm")
	if err != nil {
		return nil, err
	}
	return &IBCChannelOpenConfirmIterator{contract: _IBC.contract, event: "ChannelOpenConfirm", logs: logs, sub: sub}, nil
}

// WatchChannelOpenConfirm is a free log subscription operation binding the contract event 0xcccb79544f2a910ecd04c3bd96f870be6f5c74e0d00c18443c25ecf7b9800918.
//
// Solidity: event ChannelOpenConfirm(string connectionId, string channelId, string portId, string counterpartyChannelId, string counterpartyPortID)
func (_IBC *IBCFilterer) WatchChannelOpenConfirm(opts *bind.WatchOpts, sink chan<- *IBCChannelOpenConfirm) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "ChannelOpenConfirm")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCChannelOpenConfirm)
				if err := _IBC.contract.UnpackLog(event, "ChannelOpenConfirm", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChannelOpenConfirm is a log parse operation binding the contract event 0xcccb79544f2a910ecd04c3bd96f870be6f5c74e0d00c18443c25ecf7b9800918.
//
// Solidity: event ChannelOpenConfirm(string connectionId, string channelId, string portId, string counterpartyChannelId, string counterpartyPortID)
func (_IBC *IBCFilterer) ParseChannelOpenConfirm(log types.Log) (*IBCChannelOpenConfirm, error) {
	event := new(IBCChannelOpenConfirm)
	if err := _IBC.contract.UnpackLog(event, "ChannelOpenConfirm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCChannelOpenInitIterator is returned from FilterChannelOpenInit and is used to iterate over the raw logs and unpacked data for ChannelOpenInit events raised by the IBC contract.
type IBCChannelOpenInitIterator struct {
	Event *IBCChannelOpenInit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCChannelOpenInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCChannelOpenInit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCChannelOpenInit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCChannelOpenInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCChannelOpenInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCChannelOpenInit represents a ChannelOpenInit event raised by the IBC contract.
type IBCChannelOpenInit struct {
	ConnectionId       string
	ChannelId          string
	PortId             string
	CounterpartyPortID string
	Version            string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenInit is a free log retrieval operation binding the contract event 0x96be605fd502b1513c6e6e38963c9f28dc030f5e18c7a48ee2d4775b387b6fe0.
//
// Solidity: event ChannelOpenInit(string connectionId, string channelId, string portId, string counterpartyPortID, string version)
func (_IBC *IBCFilterer) FilterChannelOpenInit(opts *bind.FilterOpts) (*IBCChannelOpenInitIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "ChannelOpenInit")
	if err != nil {
		return nil, err
	}
	return &IBCChannelOpenInitIterator{contract: _IBC.contract, event: "ChannelOpenInit", logs: logs, sub: sub}, nil
}

// WatchChannelOpenInit is a free log subscription operation binding the contract event 0x96be605fd502b1513c6e6e38963c9f28dc030f5e18c7a48ee2d4775b387b6fe0.
//
// Solidity: event ChannelOpenInit(string connectionId, string channelId, string portId, string counterpartyPortID, string version)
func (_IBC *IBCFilterer) WatchChannelOpenInit(opts *bind.WatchOpts, sink chan<- *IBCChannelOpenInit) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "ChannelOpenInit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCChannelOpenInit)
				if err := _IBC.contract.UnpackLog(event, "ChannelOpenInit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChannelOpenInit is a log parse operation binding the contract event 0x96be605fd502b1513c6e6e38963c9f28dc030f5e18c7a48ee2d4775b387b6fe0.
//
// Solidity: event ChannelOpenInit(string connectionId, string channelId, string portId, string counterpartyPortID, string version)
func (_IBC *IBCFilterer) ParseChannelOpenInit(log types.Log) (*IBCChannelOpenInit, error) {
	event := new(IBCChannelOpenInit)
	if err := _IBC.contract.UnpackLog(event, "ChannelOpenInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCChannelOpenTryIterator is returned from FilterChannelOpenTry and is used to iterate over the raw logs and unpacked data for ChannelOpenTry events raised by the IBC contract.
type IBCChannelOpenTryIterator struct {
	Event *IBCChannelOpenTry // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCChannelOpenTryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCChannelOpenTry)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCChannelOpenTry)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCChannelOpenTryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCChannelOpenTryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCChannelOpenTry represents a ChannelOpenTry event raised by the IBC contract.
type IBCChannelOpenTry struct {
	ConnectionId          string
	ChannelId             string
	PortId                string
	CounterpartyChannelId string
	CounterpartyPortID    string
	Version               string
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenTry is a free log retrieval operation binding the contract event 0x9c5a76e8bddb2e5c238e35b7ce7a850ad22a776479bfc8b4af5e88e073fa9c70.
//
// Solidity: event ChannelOpenTry(string connectionId, string channelId, string portId, string counterpartyChannelId, string counterpartyPortID, string version)
func (_IBC *IBCFilterer) FilterChannelOpenTry(opts *bind.FilterOpts) (*IBCChannelOpenTryIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "ChannelOpenTry")
	if err != nil {
		return nil, err
	}
	return &IBCChannelOpenTryIterator{contract: _IBC.contract, event: "ChannelOpenTry", logs: logs, sub: sub}, nil
}

// WatchChannelOpenTry is a free log subscription operation binding the contract event 0x9c5a76e8bddb2e5c238e35b7ce7a850ad22a776479bfc8b4af5e88e073fa9c70.
//
// Solidity: event ChannelOpenTry(string connectionId, string channelId, string portId, string counterpartyChannelId, string counterpartyPortID, string version)
func (_IBC *IBCFilterer) WatchChannelOpenTry(opts *bind.WatchOpts, sink chan<- *IBCChannelOpenTry) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "ChannelOpenTry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCChannelOpenTry)
				if err := _IBC.contract.UnpackLog(event, "ChannelOpenTry", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChannelOpenTry is a log parse operation binding the contract event 0x9c5a76e8bddb2e5c238e35b7ce7a850ad22a776479bfc8b4af5e88e073fa9c70.
//
// Solidity: event ChannelOpenTry(string connectionId, string channelId, string portId, string counterpartyChannelId, string counterpartyPortID, string version)
func (_IBC *IBCFilterer) ParseChannelOpenTry(log types.Log) (*IBCChannelOpenTry, error) {
	event := new(IBCChannelOpenTry)
	if err := _IBC.contract.UnpackLog(event, "ChannelOpenTry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCClientCreatedIterator is returned from FilterClientCreated and is used to iterate over the raw logs and unpacked data for ClientCreated events raised by the IBC contract.
type IBCClientCreatedIterator struct {
	Event *IBCClientCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCClientCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCClientCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCClientCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCClientCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCClientCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCClientCreated represents a ClientCreated event raised by the IBC contract.
type IBCClientCreated struct {
	ClientId        string
	ConsensusHeight string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClientCreated is a free log retrieval operation binding the contract event 0xe8ff12a906de082573ebed03b1fcc6155a0313ddeb5adaa039ef09b1f6f27d85.
//
// Solidity: event ClientCreated(string clientId, string consensusHeight)
func (_IBC *IBCFilterer) FilterClientCreated(opts *bind.FilterOpts) (*IBCClientCreatedIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "ClientCreated")
	if err != nil {
		return nil, err
	}
	return &IBCClientCreatedIterator{contract: _IBC.contract, event: "ClientCreated", logs: logs, sub: sub}, nil
}

// WatchClientCreated is a free log subscription operation binding the contract event 0xe8ff12a906de082573ebed03b1fcc6155a0313ddeb5adaa039ef09b1f6f27d85.
//
// Solidity: event ClientCreated(string clientId, string consensusHeight)
func (_IBC *IBCFilterer) WatchClientCreated(opts *bind.WatchOpts, sink chan<- *IBCClientCreated) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "ClientCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCClientCreated)
				if err := _IBC.contract.UnpackLog(event, "ClientCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClientCreated is a log parse operation binding the contract event 0xe8ff12a906de082573ebed03b1fcc6155a0313ddeb5adaa039ef09b1f6f27d85.
//
// Solidity: event ClientCreated(string clientId, string consensusHeight)
func (_IBC *IBCFilterer) ParseClientCreated(log types.Log) (*IBCClientCreated, error) {
	event := new(IBCClientCreated)
	if err := _IBC.contract.UnpackLog(event, "ClientCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCClientUpdatedIterator is returned from FilterClientUpdated and is used to iterate over the raw logs and unpacked data for ClientUpdated events raised by the IBC contract.
type IBCClientUpdatedIterator struct {
	Event *IBCClientUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCClientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCClientUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCClientUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCClientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCClientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCClientUpdated represents a ClientUpdated event raised by the IBC contract.
type IBCClientUpdated struct {
	ClientId        string
	ConsensusHeight string
	ClientMessage   []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClientUpdated is a free log retrieval operation binding the contract event 0xb91f626ae6ca83625067361537862584dece7797f6fcb31844ca9b358801c38e.
//
// Solidity: event ClientUpdated(string clientId, string consensusHeight, bytes clientMessage)
func (_IBC *IBCFilterer) FilterClientUpdated(opts *bind.FilterOpts) (*IBCClientUpdatedIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "ClientUpdated")
	if err != nil {
		return nil, err
	}
	return &IBCClientUpdatedIterator{contract: _IBC.contract, event: "ClientUpdated", logs: logs, sub: sub}, nil
}

// WatchClientUpdated is a free log subscription operation binding the contract event 0xb91f626ae6ca83625067361537862584dece7797f6fcb31844ca9b358801c38e.
//
// Solidity: event ClientUpdated(string clientId, string consensusHeight, bytes clientMessage)
func (_IBC *IBCFilterer) WatchClientUpdated(opts *bind.WatchOpts, sink chan<- *IBCClientUpdated) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "ClientUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCClientUpdated)
				if err := _IBC.contract.UnpackLog(event, "ClientUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClientUpdated is a log parse operation binding the contract event 0xb91f626ae6ca83625067361537862584dece7797f6fcb31844ca9b358801c38e.
//
// Solidity: event ClientUpdated(string clientId, string consensusHeight, bytes clientMessage)
func (_IBC *IBCFilterer) ParseClientUpdated(log types.Log) (*IBCClientUpdated, error) {
	event := new(IBCClientUpdated)
	if err := _IBC.contract.UnpackLog(event, "ClientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCClientUpgradedIterator is returned from FilterClientUpgraded and is used to iterate over the raw logs and unpacked data for ClientUpgraded events raised by the IBC contract.
type IBCClientUpgradedIterator struct {
	Event *IBCClientUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCClientUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCClientUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCClientUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCClientUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCClientUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCClientUpgraded represents a ClientUpgraded event raised by the IBC contract.
type IBCClientUpgraded struct {
	ClientId        string
	ConsensusHeight string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClientUpgraded is a free log retrieval operation binding the contract event 0xd46c86e19730acb2564dcf28f4cb9520768766aecaf0a999d38dac12173a962f.
//
// Solidity: event ClientUpgraded(string clientId, string consensusHeight)
func (_IBC *IBCFilterer) FilterClientUpgraded(opts *bind.FilterOpts) (*IBCClientUpgradedIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "ClientUpgraded")
	if err != nil {
		return nil, err
	}
	return &IBCClientUpgradedIterator{contract: _IBC.contract, event: "ClientUpgraded", logs: logs, sub: sub}, nil
}

// WatchClientUpgraded is a free log subscription operation binding the contract event 0xd46c86e19730acb2564dcf28f4cb9520768766aecaf0a999d38dac12173a962f.
//
// Solidity: event ClientUpgraded(string clientId, string consensusHeight)
func (_IBC *IBCFilterer) WatchClientUpgraded(opts *bind.WatchOpts, sink chan<- *IBCClientUpgraded) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "ClientUpgraded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCClientUpgraded)
				if err := _IBC.contract.UnpackLog(event, "ClientUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClientUpgraded is a log parse operation binding the contract event 0xd46c86e19730acb2564dcf28f4cb9520768766aecaf0a999d38dac12173a962f.
//
// Solidity: event ClientUpgraded(string clientId, string consensusHeight)
func (_IBC *IBCFilterer) ParseClientUpgraded(log types.Log) (*IBCClientUpgraded, error) {
	event := new(IBCClientUpgraded)
	if err := _IBC.contract.UnpackLog(event, "ClientUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCConnectionOpenAckIterator is returned from FilterConnectionOpenAck and is used to iterate over the raw logs and unpacked data for ConnectionOpenAck events raised by the IBC contract.
type IBCConnectionOpenAckIterator struct {
	Event *IBCConnectionOpenAck // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCConnectionOpenAckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCConnectionOpenAck)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCConnectionOpenAck)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCConnectionOpenAckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCConnectionOpenAckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCConnectionOpenAck represents a ConnectionOpenAck event raised by the IBC contract.
type IBCConnectionOpenAck struct {
	ClientId                 string
	ConnectionId             string
	CounterpartyClientID     string
	CounterpartyConnectionId string
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterConnectionOpenAck is a free log retrieval operation binding the contract event 0xe7615b4ebffcb930061f901cc07ee67b4d32c8f9052141eb8bce2dec3f577fe1.
//
// Solidity: event ConnectionOpenAck(string clientId, string connectionId, string counterpartyClientID, string counterpartyConnectionId)
func (_IBC *IBCFilterer) FilterConnectionOpenAck(opts *bind.FilterOpts) (*IBCConnectionOpenAckIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "ConnectionOpenAck")
	if err != nil {
		return nil, err
	}
	return &IBCConnectionOpenAckIterator{contract: _IBC.contract, event: "ConnectionOpenAck", logs: logs, sub: sub}, nil
}

// WatchConnectionOpenAck is a free log subscription operation binding the contract event 0xe7615b4ebffcb930061f901cc07ee67b4d32c8f9052141eb8bce2dec3f577fe1.
//
// Solidity: event ConnectionOpenAck(string clientId, string connectionId, string counterpartyClientID, string counterpartyConnectionId)
func (_IBC *IBCFilterer) WatchConnectionOpenAck(opts *bind.WatchOpts, sink chan<- *IBCConnectionOpenAck) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "ConnectionOpenAck")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCConnectionOpenAck)
				if err := _IBC.contract.UnpackLog(event, "ConnectionOpenAck", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConnectionOpenAck is a log parse operation binding the contract event 0xe7615b4ebffcb930061f901cc07ee67b4d32c8f9052141eb8bce2dec3f577fe1.
//
// Solidity: event ConnectionOpenAck(string clientId, string connectionId, string counterpartyClientID, string counterpartyConnectionId)
func (_IBC *IBCFilterer) ParseConnectionOpenAck(log types.Log) (*IBCConnectionOpenAck, error) {
	event := new(IBCConnectionOpenAck)
	if err := _IBC.contract.UnpackLog(event, "ConnectionOpenAck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCConnectionOpenConfirmIterator is returned from FilterConnectionOpenConfirm and is used to iterate over the raw logs and unpacked data for ConnectionOpenConfirm events raised by the IBC contract.
type IBCConnectionOpenConfirmIterator struct {
	Event *IBCConnectionOpenConfirm // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCConnectionOpenConfirmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCConnectionOpenConfirm)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCConnectionOpenConfirm)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCConnectionOpenConfirmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCConnectionOpenConfirmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCConnectionOpenConfirm represents a ConnectionOpenConfirm event raised by the IBC contract.
type IBCConnectionOpenConfirm struct {
	ClientId                 string
	ConnectionId             string
	CounterpartyClientID     string
	CounterpartyConnectionId string
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterConnectionOpenConfirm is a free log retrieval operation binding the contract event 0x063c0e9664347d8013d3575d502050fd936d3b51035f056696a639523feaed6d.
//
// Solidity: event ConnectionOpenConfirm(string clientId, string connectionId, string counterpartyClientID, string counterpartyConnectionId)
func (_IBC *IBCFilterer) FilterConnectionOpenConfirm(opts *bind.FilterOpts) (*IBCConnectionOpenConfirmIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "ConnectionOpenConfirm")
	if err != nil {
		return nil, err
	}
	return &IBCConnectionOpenConfirmIterator{contract: _IBC.contract, event: "ConnectionOpenConfirm", logs: logs, sub: sub}, nil
}

// WatchConnectionOpenConfirm is a free log subscription operation binding the contract event 0x063c0e9664347d8013d3575d502050fd936d3b51035f056696a639523feaed6d.
//
// Solidity: event ConnectionOpenConfirm(string clientId, string connectionId, string counterpartyClientID, string counterpartyConnectionId)
func (_IBC *IBCFilterer) WatchConnectionOpenConfirm(opts *bind.WatchOpts, sink chan<- *IBCConnectionOpenConfirm) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "ConnectionOpenConfirm")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCConnectionOpenConfirm)
				if err := _IBC.contract.UnpackLog(event, "ConnectionOpenConfirm", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConnectionOpenConfirm is a log parse operation binding the contract event 0x063c0e9664347d8013d3575d502050fd936d3b51035f056696a639523feaed6d.
//
// Solidity: event ConnectionOpenConfirm(string clientId, string connectionId, string counterpartyClientID, string counterpartyConnectionId)
func (_IBC *IBCFilterer) ParseConnectionOpenConfirm(log types.Log) (*IBCConnectionOpenConfirm, error) {
	event := new(IBCConnectionOpenConfirm)
	if err := _IBC.contract.UnpackLog(event, "ConnectionOpenConfirm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCConnectionOpenInitIterator is returned from FilterConnectionOpenInit and is used to iterate over the raw logs and unpacked data for ConnectionOpenInit events raised by the IBC contract.
type IBCConnectionOpenInitIterator struct {
	Event *IBCConnectionOpenInit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCConnectionOpenInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCConnectionOpenInit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCConnectionOpenInit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCConnectionOpenInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCConnectionOpenInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCConnectionOpenInit represents a ConnectionOpenInit event raised by the IBC contract.
type IBCConnectionOpenInit struct {
	ClientId             string
	ConnectionId         string
	CounterpartyClientID string
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterConnectionOpenInit is a free log retrieval operation binding the contract event 0x9f1f1ea41ae20b9e071603ac41a1783f3d7fcbaf413365fe97cfd6b1c155247c.
//
// Solidity: event ConnectionOpenInit(string clientId, string connectionId, string counterpartyClientID)
func (_IBC *IBCFilterer) FilterConnectionOpenInit(opts *bind.FilterOpts) (*IBCConnectionOpenInitIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "ConnectionOpenInit")
	if err != nil {
		return nil, err
	}
	return &IBCConnectionOpenInitIterator{contract: _IBC.contract, event: "ConnectionOpenInit", logs: logs, sub: sub}, nil
}

// WatchConnectionOpenInit is a free log subscription operation binding the contract event 0x9f1f1ea41ae20b9e071603ac41a1783f3d7fcbaf413365fe97cfd6b1c155247c.
//
// Solidity: event ConnectionOpenInit(string clientId, string connectionId, string counterpartyClientID)
func (_IBC *IBCFilterer) WatchConnectionOpenInit(opts *bind.WatchOpts, sink chan<- *IBCConnectionOpenInit) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "ConnectionOpenInit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCConnectionOpenInit)
				if err := _IBC.contract.UnpackLog(event, "ConnectionOpenInit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConnectionOpenInit is a log parse operation binding the contract event 0x9f1f1ea41ae20b9e071603ac41a1783f3d7fcbaf413365fe97cfd6b1c155247c.
//
// Solidity: event ConnectionOpenInit(string clientId, string connectionId, string counterpartyClientID)
func (_IBC *IBCFilterer) ParseConnectionOpenInit(log types.Log) (*IBCConnectionOpenInit, error) {
	event := new(IBCConnectionOpenInit)
	if err := _IBC.contract.UnpackLog(event, "ConnectionOpenInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCConnectionOpenTryIterator is returned from FilterConnectionOpenTry and is used to iterate over the raw logs and unpacked data for ConnectionOpenTry events raised by the IBC contract.
type IBCConnectionOpenTryIterator struct {
	Event *IBCConnectionOpenTry // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCConnectionOpenTryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCConnectionOpenTry)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCConnectionOpenTry)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCConnectionOpenTryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCConnectionOpenTryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCConnectionOpenTry represents a ConnectionOpenTry event raised by the IBC contract.
type IBCConnectionOpenTry struct {
	ClientId                 string
	ConnectionId             string
	CounterpartyClientID     string
	CounterpartyConnectionId string
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterConnectionOpenTry is a free log retrieval operation binding the contract event 0xa616a9aa2c65e935abbd15b07a9b5ff6c9c48b06b460a39b0b8cfda2a985869f.
//
// Solidity: event ConnectionOpenTry(string clientId, string connectionId, string counterpartyClientID, string counterpartyConnectionId)
func (_IBC *IBCFilterer) FilterConnectionOpenTry(opts *bind.FilterOpts) (*IBCConnectionOpenTryIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "ConnectionOpenTry")
	if err != nil {
		return nil, err
	}
	return &IBCConnectionOpenTryIterator{contract: _IBC.contract, event: "ConnectionOpenTry", logs: logs, sub: sub}, nil
}

// WatchConnectionOpenTry is a free log subscription operation binding the contract event 0xa616a9aa2c65e935abbd15b07a9b5ff6c9c48b06b460a39b0b8cfda2a985869f.
//
// Solidity: event ConnectionOpenTry(string clientId, string connectionId, string counterpartyClientID, string counterpartyConnectionId)
func (_IBC *IBCFilterer) WatchConnectionOpenTry(opts *bind.WatchOpts, sink chan<- *IBCConnectionOpenTry) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "ConnectionOpenTry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCConnectionOpenTry)
				if err := _IBC.contract.UnpackLog(event, "ConnectionOpenTry", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConnectionOpenTry is a log parse operation binding the contract event 0xa616a9aa2c65e935abbd15b07a9b5ff6c9c48b06b460a39b0b8cfda2a985869f.
//
// Solidity: event ConnectionOpenTry(string clientId, string connectionId, string counterpartyClientID, string counterpartyConnectionId)
func (_IBC *IBCFilterer) ParseConnectionOpenTry(log types.Log) (*IBCConnectionOpenTry, error) {
	event := new(IBCConnectionOpenTry)
	if err := _IBC.contract.UnpackLog(event, "ConnectionOpenTry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCPacketRecvIterator is returned from FilterPacketRecv and is used to iterate over the raw logs and unpacked data for PacketRecv events raised by the IBC contract.
type IBCPacketRecvIterator struct {
	Event *IBCPacketRecv // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCPacketRecvIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCPacketRecv)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCPacketRecv)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCPacketRecvIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCPacketRecvIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCPacketRecv represents a PacketRecv event raised by the IBC contract.
type IBCPacketRecv struct {
	Data             []byte
	TimeoutHeight    string
	TimeoutTimestamp *big.Int
	Sequence         *big.Int
	SourcePort       string
	SourceChannel    string
	DestPort         string
	DestChannel      string
	ChannelOrdering  int32
	Connection       string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPacketRecv is a free log retrieval operation binding the contract event 0x0573ccb55881bfc7de97538f9e8b9079570f05c34e38d6b21f5434777dc1bae7.
//
// Solidity: event PacketRecv(bytes data, string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, int32 channelOrdering, string connection)
func (_IBC *IBCFilterer) FilterPacketRecv(opts *bind.FilterOpts) (*IBCPacketRecvIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "PacketRecv")
	if err != nil {
		return nil, err
	}
	return &IBCPacketRecvIterator{contract: _IBC.contract, event: "PacketRecv", logs: logs, sub: sub}, nil
}

// WatchPacketRecv is a free log subscription operation binding the contract event 0x0573ccb55881bfc7de97538f9e8b9079570f05c34e38d6b21f5434777dc1bae7.
//
// Solidity: event PacketRecv(bytes data, string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, int32 channelOrdering, string connection)
func (_IBC *IBCFilterer) WatchPacketRecv(opts *bind.WatchOpts, sink chan<- *IBCPacketRecv) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "PacketRecv")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCPacketRecv)
				if err := _IBC.contract.UnpackLog(event, "PacketRecv", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketRecv is a log parse operation binding the contract event 0x0573ccb55881bfc7de97538f9e8b9079570f05c34e38d6b21f5434777dc1bae7.
//
// Solidity: event PacketRecv(bytes data, string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, int32 channelOrdering, string connection)
func (_IBC *IBCFilterer) ParsePacketRecv(log types.Log) (*IBCPacketRecv, error) {
	event := new(IBCPacketRecv)
	if err := _IBC.contract.UnpackLog(event, "PacketRecv", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCPacketSentIterator is returned from FilterPacketSent and is used to iterate over the raw logs and unpacked data for PacketSent events raised by the IBC contract.
type IBCPacketSentIterator struct {
	Event *IBCPacketSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCPacketSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCPacketSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCPacketSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCPacketSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCPacketSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCPacketSent represents a PacketSent event raised by the IBC contract.
type IBCPacketSent struct {
	Data             []byte
	TimeoutHeight    string
	TimeoutTimestamp *big.Int
	Sequence         *big.Int
	SourcePort       string
	SourceChannel    string
	DestPort         string
	DestChannel      string
	ChannelOrdering  int32
	Connection       string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPacketSent is a free log retrieval operation binding the contract event 0x374a665e4553b182bc8023165f699e510a7ead44fbf297239d100bfefcdb32bf.
//
// Solidity: event PacketSent(bytes data, string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, int32 channelOrdering, string connection)
func (_IBC *IBCFilterer) FilterPacketSent(opts *bind.FilterOpts) (*IBCPacketSentIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return &IBCPacketSentIterator{contract: _IBC.contract, event: "PacketSent", logs: logs, sub: sub}, nil
}

// WatchPacketSent is a free log subscription operation binding the contract event 0x374a665e4553b182bc8023165f699e510a7ead44fbf297239d100bfefcdb32bf.
//
// Solidity: event PacketSent(bytes data, string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, int32 channelOrdering, string connection)
func (_IBC *IBCFilterer) WatchPacketSent(opts *bind.WatchOpts, sink chan<- *IBCPacketSent) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCPacketSent)
				if err := _IBC.contract.UnpackLog(event, "PacketSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketSent is a log parse operation binding the contract event 0x374a665e4553b182bc8023165f699e510a7ead44fbf297239d100bfefcdb32bf.
//
// Solidity: event PacketSent(bytes data, string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, int32 channelOrdering, string connection)
func (_IBC *IBCFilterer) ParsePacketSent(log types.Log) (*IBCPacketSent, error) {
	event := new(IBCPacketSent)
	if err := _IBC.contract.UnpackLog(event, "PacketSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCTimeoutPacketIterator is returned from FilterTimeoutPacket and is used to iterate over the raw logs and unpacked data for TimeoutPacket events raised by the IBC contract.
type IBCTimeoutPacketIterator struct {
	Event *IBCTimeoutPacket // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCTimeoutPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCTimeoutPacket)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCTimeoutPacket)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCTimeoutPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCTimeoutPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCTimeoutPacket represents a TimeoutPacket event raised by the IBC contract.
type IBCTimeoutPacket struct {
	TimeoutHeight    string
	TimeoutTimestamp *big.Int
	Sequence         *big.Int
	SourcePort       string
	SourceChannel    string
	DestPort         string
	DestChannel      string
	ChannelOrdering  int32
	ConnectionID     string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTimeoutPacket is a free log retrieval operation binding the contract event 0x280b5c88e7ecdaacc40ca0de13e47206493bdee68e9656ef49e359cb36aa4c12.
//
// Solidity: event TimeoutPacket(string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, int32 channelOrdering, string connectionID)
func (_IBC *IBCFilterer) FilterTimeoutPacket(opts *bind.FilterOpts) (*IBCTimeoutPacketIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "TimeoutPacket")
	if err != nil {
		return nil, err
	}
	return &IBCTimeoutPacketIterator{contract: _IBC.contract, event: "TimeoutPacket", logs: logs, sub: sub}, nil
}

// WatchTimeoutPacket is a free log subscription operation binding the contract event 0x280b5c88e7ecdaacc40ca0de13e47206493bdee68e9656ef49e359cb36aa4c12.
//
// Solidity: event TimeoutPacket(string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, int32 channelOrdering, string connectionID)
func (_IBC *IBCFilterer) WatchTimeoutPacket(opts *bind.WatchOpts, sink chan<- *IBCTimeoutPacket) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "TimeoutPacket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCTimeoutPacket)
				if err := _IBC.contract.UnpackLog(event, "TimeoutPacket", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimeoutPacket is a log parse operation binding the contract event 0x280b5c88e7ecdaacc40ca0de13e47206493bdee68e9656ef49e359cb36aa4c12.
//
// Solidity: event TimeoutPacket(string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, int32 channelOrdering, string connectionID)
func (_IBC *IBCFilterer) ParseTimeoutPacket(log types.Log) (*IBCTimeoutPacket, error) {
	event := new(IBCTimeoutPacket)
	if err := _IBC.contract.UnpackLog(event, "TimeoutPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCTypeChannelClosedIterator is returned from FilterTypeChannelClosed and is used to iterate over the raw logs and unpacked data for TypeChannelClosed events raised by the IBC contract.
type IBCTypeChannelClosedIterator struct {
	Event *IBCTypeChannelClosed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCTypeChannelClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCTypeChannelClosed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCTypeChannelClosed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCTypeChannelClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCTypeChannelClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCTypeChannelClosed represents a TypeChannelClosed event raised by the IBC contract.
type IBCTypeChannelClosed struct {
	SourcePort      string
	SourceChannel   string
	DestPort        string
	DestChannel     string
	ConnectionID    string
	ChannelOrdering string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTypeChannelClosed is a free log retrieval operation binding the contract event 0xf746dad35ccffe1b392a9d8643ff590f9a8faffcfd919afef8d68010d88eba6f.
//
// Solidity: event TypeChannelClosed(string sourcePort, string sourceChannel, string destPort, string destChannel, string ConnectionID, string ChannelOrdering)
func (_IBC *IBCFilterer) FilterTypeChannelClosed(opts *bind.FilterOpts) (*IBCTypeChannelClosedIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "TypeChannelClosed")
	if err != nil {
		return nil, err
	}
	return &IBCTypeChannelClosedIterator{contract: _IBC.contract, event: "TypeChannelClosed", logs: logs, sub: sub}, nil
}

// WatchTypeChannelClosed is a free log subscription operation binding the contract event 0xf746dad35ccffe1b392a9d8643ff590f9a8faffcfd919afef8d68010d88eba6f.
//
// Solidity: event TypeChannelClosed(string sourcePort, string sourceChannel, string destPort, string destChannel, string ConnectionID, string ChannelOrdering)
func (_IBC *IBCFilterer) WatchTypeChannelClosed(opts *bind.WatchOpts, sink chan<- *IBCTypeChannelClosed) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "TypeChannelClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCTypeChannelClosed)
				if err := _IBC.contract.UnpackLog(event, "TypeChannelClosed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTypeChannelClosed is a log parse operation binding the contract event 0xf746dad35ccffe1b392a9d8643ff590f9a8faffcfd919afef8d68010d88eba6f.
//
// Solidity: event TypeChannelClosed(string sourcePort, string sourceChannel, string destPort, string destChannel, string ConnectionID, string ChannelOrdering)
func (_IBC *IBCFilterer) ParseTypeChannelClosed(log types.Log) (*IBCTypeChannelClosed, error) {
	event := new(IBCTypeChannelClosed)
	if err := _IBC.contract.UnpackLog(event, "TypeChannelClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCTypeSubmitMisbehaviourIterator is returned from FilterTypeSubmitMisbehaviour and is used to iterate over the raw logs and unpacked data for TypeSubmitMisbehaviour events raised by the IBC contract.
type IBCTypeSubmitMisbehaviourIterator struct {
	Event *IBCTypeSubmitMisbehaviour // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCTypeSubmitMisbehaviourIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCTypeSubmitMisbehaviour)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCTypeSubmitMisbehaviour)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCTypeSubmitMisbehaviourIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCTypeSubmitMisbehaviourIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCTypeSubmitMisbehaviour represents a TypeSubmitMisbehaviour event raised by the IBC contract.
type IBCTypeSubmitMisbehaviour struct {
	ClientID   string
	ClientType string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTypeSubmitMisbehaviour is a free log retrieval operation binding the contract event 0x0169c10587109b18b08996374f7ca7ea7e818fec9a0557b2c6c79970fba0927f.
//
// Solidity: event TypeSubmitMisbehaviour(string clientID, string clientType)
func (_IBC *IBCFilterer) FilterTypeSubmitMisbehaviour(opts *bind.FilterOpts) (*IBCTypeSubmitMisbehaviourIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "TypeSubmitMisbehaviour")
	if err != nil {
		return nil, err
	}
	return &IBCTypeSubmitMisbehaviourIterator{contract: _IBC.contract, event: "TypeSubmitMisbehaviour", logs: logs, sub: sub}, nil
}

// WatchTypeSubmitMisbehaviour is a free log subscription operation binding the contract event 0x0169c10587109b18b08996374f7ca7ea7e818fec9a0557b2c6c79970fba0927f.
//
// Solidity: event TypeSubmitMisbehaviour(string clientID, string clientType)
func (_IBC *IBCFilterer) WatchTypeSubmitMisbehaviour(opts *bind.WatchOpts, sink chan<- *IBCTypeSubmitMisbehaviour) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "TypeSubmitMisbehaviour")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCTypeSubmitMisbehaviour)
				if err := _IBC.contract.UnpackLog(event, "TypeSubmitMisbehaviour", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTypeSubmitMisbehaviour is a log parse operation binding the contract event 0x0169c10587109b18b08996374f7ca7ea7e818fec9a0557b2c6c79970fba0927f.
//
// Solidity: event TypeSubmitMisbehaviour(string clientID, string clientType)
func (_IBC *IBCFilterer) ParseTypeSubmitMisbehaviour(log types.Log) (*IBCTypeSubmitMisbehaviour, error) {
	event := new(IBCTypeSubmitMisbehaviour)
	if err := _IBC.contract.UnpackLog(event, "TypeSubmitMisbehaviour", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCWarpIBCMessageIterator is returned from FilterWarpIBCMessage and is used to iterate over the raw logs and unpacked data for WarpIBCMessage events raised by the IBC contract.
type IBCWarpIBCMessageIterator struct {
	Event *IBCWarpIBCMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCWarpIBCMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCWarpIBCMessage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCWarpIBCMessage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCWarpIBCMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCWarpIBCMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCWarpIBCMessage represents a WarpIBCMessage event raised by the IBC contract.
type IBCWarpIBCMessage struct {
	Sender    common.Address
	MessageID [32]byte
	Message   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWarpIBCMessage is a free log retrieval operation binding the contract event 0x4545d1449b6bfe63eb4db31f933e745f07e76b3a1f295a66222dc33a4a597ea4.
//
// Solidity: event WarpIBCMessage(address indexed sender, bytes32 indexed messageID, bytes message)
func (_IBC *IBCFilterer) FilterWarpIBCMessage(opts *bind.FilterOpts, sender []common.Address, messageID [][32]byte) (*IBCWarpIBCMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _IBC.contract.FilterLogs(opts, "WarpIBCMessage", senderRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &IBCWarpIBCMessageIterator{contract: _IBC.contract, event: "WarpIBCMessage", logs: logs, sub: sub}, nil
}

// WatchWarpIBCMessage is a free log subscription operation binding the contract event 0x4545d1449b6bfe63eb4db31f933e745f07e76b3a1f295a66222dc33a4a597ea4.
//
// Solidity: event WarpIBCMessage(address indexed sender, bytes32 indexed messageID, bytes message)
func (_IBC *IBCFilterer) WatchWarpIBCMessage(opts *bind.WatchOpts, sink chan<- *IBCWarpIBCMessage, sender []common.Address, messageID [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _IBC.contract.WatchLogs(opts, "WarpIBCMessage", senderRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCWarpIBCMessage)
				if err := _IBC.contract.UnpackLog(event, "WarpIBCMessage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWarpIBCMessage is a log parse operation binding the contract event 0x4545d1449b6bfe63eb4db31f933e745f07e76b3a1f295a66222dc33a4a597ea4.
//
// Solidity: event WarpIBCMessage(address indexed sender, bytes32 indexed messageID, bytes message)
func (_IBC *IBCFilterer) ParseWarpIBCMessage(log types.Log) (*IBCWarpIBCMessage, error) {
	event := new(IBCWarpIBCMessage)
	if err := _IBC.contract.UnpackLog(event, "WarpIBCMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCWriteAckIterator is returned from FilterWriteAck and is used to iterate over the raw logs and unpacked data for WriteAck events raised by the IBC contract.
type IBCWriteAckIterator struct {
	Event *IBCWriteAck // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCWriteAckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCWriteAck)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCWriteAck)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCWriteAckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCWriteAckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCWriteAck represents a WriteAck event raised by the IBC contract.
type IBCWriteAck struct {
	Data             []byte
	TimeoutHeight    string
	TimeoutTimestamp *big.Int
	Sequence         *big.Int
	SourcePort       string
	SourceChannel    string
	DestPort         string
	DestChannel      string
	Ack              []byte
	ConnectionID     string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWriteAck is a free log retrieval operation binding the contract event 0x10c1bf521269fc605b4eb88afd17258ec9fc67515ed2e874e23d9f63bf931fc2.
//
// Solidity: event WriteAck(bytes data, string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, bytes ack, string connectionID)
func (_IBC *IBCFilterer) FilterWriteAck(opts *bind.FilterOpts) (*IBCWriteAckIterator, error) {

	logs, sub, err := _IBC.contract.FilterLogs(opts, "WriteAck")
	if err != nil {
		return nil, err
	}
	return &IBCWriteAckIterator{contract: _IBC.contract, event: "WriteAck", logs: logs, sub: sub}, nil
}

// WatchWriteAck is a free log subscription operation binding the contract event 0x10c1bf521269fc605b4eb88afd17258ec9fc67515ed2e874e23d9f63bf931fc2.
//
// Solidity: event WriteAck(bytes data, string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, bytes ack, string connectionID)
func (_IBC *IBCFilterer) WatchWriteAck(opts *bind.WatchOpts, sink chan<- *IBCWriteAck) (event.Subscription, error) {

	logs, sub, err := _IBC.contract.WatchLogs(opts, "WriteAck")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCWriteAck)
				if err := _IBC.contract.UnpackLog(event, "WriteAck", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWriteAck is a log parse operation binding the contract event 0x10c1bf521269fc605b4eb88afd17258ec9fc67515ed2e874e23d9f63bf931fc2.
//
// Solidity: event WriteAck(bytes data, string timeoutHeight, uint256 timeoutTimestamp, uint256 sequence, string sourcePort, string sourceChannel, string destPort, string destChannel, bytes ack, string connectionID)
func (_IBC *IBCFilterer) ParseWriteAck(log types.Log) (*IBCWriteAck, error) {
	event := new(IBCWriteAck)
	if err := _IBC.contract.UnpackLog(event, "WriteAck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
