// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

import (
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_modeMintBurn\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ERC20FeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"}],\"name\":\"FeeReceiverChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestTxHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumBridgeTransfer.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueOrTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"handleNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lowerHandleNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"HandleValueTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"KLAYFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"KLAYLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"KLAYUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumBridgeTransfer.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueOrTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"requestNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"RequestValueTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumBridgeTransfer.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueOrTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"requestNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"encodingVer\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedData\",\"type\":\"bytes\"}],\"name\":\"RequestValueTransferEncoded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenUnlocked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_OPERATOR\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chargeWithoutEvent\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"closedValueTransferVotes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configurationNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpartBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"deregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"deregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feeOfERC20\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeOfKLAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceiver\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperatorList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredTokenList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_requestedNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_requestedBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"handleERC20Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_requestedNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_requestedBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"handleERC721Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_requestedNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_requestedBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"handleKLAYTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"handleNoncesToBlockNums\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"handledRequestTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"indexOfTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLockedKLAY\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRunning\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockKLAY\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"lockToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lowerHandleNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modeMintBurn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"onERC20Received\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"operatorThresholds\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoveryBlockNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registeredTokenList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"requestERC20Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"requestERC721Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"requestKLAYTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"setCounterPartBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_requestNonce\",\"type\":\"uint64\"}],\"name\":\"setERC20Fee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeReceiver\",\"type\":\"address\"}],\"name\":\"setFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_requestNonce\",\"type\":\"uint64\"}],\"name\":\"setKLAYFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumBridgeOperator.VoteType\",\"name\":\"_voteType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"}],\"name\":\"setOperatorThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockKLAY\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"unlockToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upperHandleNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// BridgeBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BridgeBinRuntime = ``

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x60806040526000600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060006002556000600f60086101000a81548160ff0219169083151502179055506001600f60096101000a81548160ff0219169083151502179055506001601060086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060405162007180380380620071808339818101604052810190620000d0919062000393565b80600080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000135620001296200028360201b60201c565b6200028b60201b60201c565b60005b6002808111156200014e576200014d620003c5565b5b60ff168160ff161015620001a5576001600e60008360ff1660ff16815260200190815260200160002060006101000a81548160ff021916908360ff16021790555080806200019c9062000430565b91505062000138565b506001600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600d339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600f60086101000a81548160ff02191690831515021790555050506200045e565b600033905090565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b60008115159050919050565b6200036d8162000356565b81146200037957600080fd5b50565b6000815190506200038d8162000362565b92915050565b600060208284031215620003ac57620003ab62000351565b5b6000620003bc848285016200037c565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600060ff82169050919050565b60006200043d8262000423565b915060ff8203620004535762000452620003f4565b5b600182019050919050565b616d12806200046e6000396000f3fe6080604052600436106103035760003560e01c80638a75eee211610190578063c263b5d6116100dc578063ea21eade11610095578063f1656e531161006f578063f1656e5314610baf578063f171996614610bd8578063f2fde38b14610c03578063ffa1ad7414610c2c57610364565b8063ea21eade14610b32578063ee2aec6514610b5d578063efdcd97414610b8657610364565b8063c263b5d614610a45578063c877cf3714610a70578063cb38f40714610a99578063cf0da29014610ad6578063d8cf98ca14610aff578063dd9222d614610b2857610364565b80639f07132911610149578063afb6022311610123578063afb602231461099d578063b2c01030146109c6578063b3f00674146109f1578063bab2af1d14610a1c57610364565b80639f07132914610932578063a066a7ed14610949578063ac6fff0b1461097257610364565b80638a75eee2146107fc5780638c0bd916146108395780638da5cb5b146108765780639832c1d7146108a1578063989ba0d3146108de5780639ef2017b1461090957610364565b8063407e6bae1161024f5780635526f76b11610208578063715018a6116101e2578063715018a61461077557806375ebdc091461078c5780637c1a0302146107a857806387b04c55146107d357610364565b80635526f76b146106d05780635eb7413a1461070d5780636e176ec21461074a57610364565b8063407e6bae146105ae5780634739f7e5146105d7578063488af8711461060057806348a18a6a1461063d5780634b40b8261461067a57806354edad72146106a557610364565b806322604742116102bc5780633682a450116102965780633682a450146104f25780633a3099d11461051b5780633a348533146105465780633e4fe9491461057157610364565b8063226047421461047757806326c23b54146104a05780632f88396c146104c957610364565b806310693fcd1461036957806313a6738a1461039257806313e7c9d8146103cf5780631a2ae53e1461040c5780631ebdca38146104355780632014e5d11461044c57610364565b366103645761036233600254600067ffffffffffffffff81111561032a57610329614dc7565b5b6040519080825280601f01601f19166020018201604052801561035c5781602001600182028036833780820191505090505b50610c57565b005b600080fd5b34801561037557600080fd5b50610390600480360381019061038b9190614e68565b610e45565b005b34801561039e57600080fd5b506103b960048036038101906103b49190614ed5565b6110c4565b6040516103c69190614f11565b60405180910390f35b3480156103db57600080fd5b506103f660048036038101906103f19190614e68565b6110eb565b6040516104039190614f47565b60405180910390f35b34801561041857600080fd5b50610433600480360381019061042e9190614f98565b61110b565b005b34801561044157600080fd5b5061044a6111b3565b005b34801561045857600080fd5b506104616112ce565b60405161046e9190614f47565b60405180910390f35b34801561048357600080fd5b5061049e600480360381019061049991906150ef565b6112e1565b005b3480156104ac57600080fd5b506104c760048036038101906104c29190615172565b611363565b005b3480156104d557600080fd5b506104f060048036038101906104eb9190615209565b611442565b005b3480156104fe57600080fd5b5061051960048036038101906105149190614e68565b6114ec565b005b34801561052757600080fd5b50610530611705565b60405161053d9190614f11565b60405180910390f35b34801561055257600080fd5b5061055b61170a565b604051610568919061526b565b60405180910390f35b34801561057d57600080fd5b5061059860048036038101906105939190615286565b611730565b6040516105a5919061526b565b60405180910390f35b3480156105ba57600080fd5b506105d560048036038101906105d091906152e9565b61176f565b005b3480156105e357600080fd5b506105fe60048036038101906105f991906153bb565b611a5f565b005b34801561060c57600080fd5b5061062760048036038101906106229190614e68565b611d1c565b604051610634919061540a565b60405180910390f35b34801561064957600080fd5b50610664600480360381019061065f9190614e68565b611d34565b604051610671919061540a565b60405180910390f35b34801561068657600080fd5b5061068f611d4c565b60405161069c9190614f11565b60405180910390f35b3480156106b157600080fd5b506106ba611d66565b6040516106c79190614f11565b60405180910390f35b3480156106dc57600080fd5b506106f760048036038101906106f2919061545e565b611d80565b604051610704919061549a565b60405180910390f35b34801561071957600080fd5b50610734600480360381019061072f9190614e68565b611da0565b6040516107419190614f47565b60405180910390f35b34801561075657600080fd5b5061075f611dc0565b60405161076c9190614f47565b60405180910390f35b34801561078157600080fd5b5061078a611dd3565b005b6107a660048036038101906107a19190615515565b611e5b565b005b3480156107b457600080fd5b506107bd611ec1565b6040516107ca9190614f11565b60405180910390f35b3480156107df57600080fd5b506107fa60048036038101906107f59190614e68565b611edb565b005b34801561080857600080fd5b50610823600480360381019061081e9190615589565b611f9b565b6040516108309190614f47565b60405180910390f35b34801561084557600080fd5b50610860600480360381019061085b9190614e68565b611fbb565b60405161086d919061526b565b60405180910390f35b34801561088257600080fd5b5061088b611fee565b604051610898919061526b565b60405180910390f35b3480156108ad57600080fd5b506108c860048036038101906108c39190614ed5565b612018565b6040516108d59190614f47565b60405180910390f35b3480156108ea57600080fd5b506108f3612038565b6040516109009190614f11565b60405180910390f35b34801561091557600080fd5b50610930600480360381019061092b9190614e68565b612052565b005b34801561093e57600080fd5b506109476122c8565b005b34801561095557600080fd5b50610970600480360381019061096b91906155f4565b6123e3565b005b34801561097e57600080fd5b506109876125d5565b6040516109949190614f11565b60405180910390f35b3480156109a957600080fd5b506109c460048036038101906109bf9190615753565b6125ef565b005b3480156109d257600080fd5b506109db612894565b6040516109e89190615913565b60405180910390f35b3480156109fd57600080fd5b50610a06612922565b604051610a139190615944565b60405180910390f35b348015610a2857600080fd5b50610a436004803603810190610a3e9190614e68565b612948565b005b348015610a5157600080fd5b50610a5a612da2565b604051610a67919061540a565b60405180910390f35b348015610a7c57600080fd5b50610a976004803603810190610a92919061598b565b612da8565b005b348015610aa557600080fd5b50610ac06004803603810190610abb9190615286565b612e41565b604051610acd919061526b565b60405180910390f35b348015610ae257600080fd5b50610afd6004803603810190610af891906159b8565b612e80565b005b348015610b0b57600080fd5b50610b266004803603810190610b219190614e68565b612e93565b005b610b30613144565b005b348015610b3e57600080fd5b50610b47613146565b604051610b549190615913565b60405180910390f35b348015610b6957600080fd5b50610b846004803603810190610b7f9190615a60565b6131d4565b005b348015610b9257600080fd5b50610bad6004803603810190610ba89190615aa0565b613329565b005b348015610bbb57600080fd5b50610bd66004803603810190610bd19190615172565b6133b1565b005b348015610be457600080fd5b50610bed6133c6565b604051610bfa9190614f47565b60405180910390f35b348015610c0f57600080fd5b50610c2a6004803603810190610c259190614e68565b6133d9565b005b348015610c3857600080fd5b50610c416134d0565b604051610c4e9190614f11565b60405180910390f35b60001515601260009054906101000a900460ff16151514610cad576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ca490615b2a565b60405180910390fd5b600f60099054906101000a900460ff16610cfc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf390615b96565b60405180910390fd5b813411610d3e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3590615c02565b60405180910390fd5b6000610d49836134d5565b9050600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167feff76c36e53fa5ff52f27acc8a34d5047a8246abb07b77b12f1309f71e337f0960008734610dc09190615c51565b600f600a9054906101000a900467ffffffffffffffff168789604051610dea959493929190615d84565b60405180910390a4600f600a81819054906101000a900467ffffffffffffffff1680929190610e1890615dde565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050505050565b610e4d6136af565b73ffffffffffffffffffffffffffffffffffffffff16610e6b611fee565b73ffffffffffffffffffffffffffffffffffffffff1614610ec1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eb890615e5a565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff16600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610f90576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f8790615ec6565b60405180910390fd5b8160001515600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514611024576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101b90615f32565b60405180910390fd5b6001600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508273ffffffffffffffffffffffffffffffffffffffff167fca1b0a14e18ada4c44846768dd186e35630cdc5cfeaca83c404ae4acaafbecd760405160405180910390a2505050565b60116020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b600c6020528060005260406000206000915054906101000a900460ff1681565b600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611197576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161118e90615f9e565b60405180910390fd5b6111a0816136b7565b156111af576111ae826137b4565b5b5050565b6111bb6136af565b73ffffffffffffffffffffffffffffffffffffffff166111d9611fee565b73ffffffffffffffffffffffffffffffffffffffff161461122f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161122690615e5a565b60405180910390fd5b60011515601260009054906101000a900460ff16151514611285576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161127c9061600a565b60405180910390fd5b6000601260006101000a81548160ff0219169083151502179055507fd20610c9b78a6903ef134539e3deb5d243be461de6ef12d4c29536bb9b54fa1b60405160405180910390a1565b600f60099054906101000a900460ff1681565b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b815260040161131e9392919061602a565b600060405180830381600087803b15801561133857600080fd5b505af115801561134c573d6000803e3d6000fd5b5050505061135d84338585856137eb565b50505050565b8473ffffffffffffffffffffffffffffffffffffffff166323b872dd3330858761138d9190616061565b6040518463ffffffff1660e01b81526004016113ab9392919061602a565b6020604051808303816000875af11580156113ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113ee91906160cc565b61142d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114249061616b565b60405180910390fd5b61143b853386868686613bb8565b5050505050565b600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166114ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114c590615f9e565b60405180910390fd5b6114d7816136b7565b156114e7576114e68383613f30565b5b505050565b6114f46136af565b73ffffffffffffffffffffffffffffffffffffffff16611512611fee565b73ffffffffffffffffffffffffffffffffffffffff1614611568576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161155f90615e5a565b60405180910390fd5b600c67ffffffffffffffff16600d80549050106115ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115b1906161d7565b60405180910390fd5b600c60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611647576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161163e90616243565b60405180910390fd5b6001600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600d819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600c81565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6008818154811061174057600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166117fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117f290615f9e565b60405180910390fd5b61180483613fbc565b61180d8361402c565b15611a555761181b88614125565b81601160008567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061187283614153565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f12b02f226d965a2881e0e8ffed6c421803a22d57ad91f9ef996fe0748ea101758b60018989600f60129054906101000a900467ffffffffffffffff168960405161190796959493929190616272565b60405180910390a4600f60089054906101000a900460ff1615611996578473ffffffffffffffffffffffffffffffffffffffff166340c10f1987866040518363ffffffff1660e01b815260040161195f9291906162da565b600060405180830381600087803b15801561197957600080fd5b505af115801561198d573d6000803e3d6000fd5b50505050611a54565b8473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb87866040518363ffffffff1660e01b81526004016119d19291906162da565b6020604051808303816000875af11580156119f0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a1491906160cc565b611a53576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a4a90616375565b60405180910390fd5b5b5b5050505050505050565b611a676136af565b73ffffffffffffffffffffffffffffffffffffffff16611a85611fee565b73ffffffffffffffffffffffffffffffffffffffff1614611adb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ad290615e5a565b60405180910390fd5b81600073ffffffffffffffffffffffffffffffffffffffff16600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611baa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ba1906163e1565b60405180910390fd5b81600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600880549050600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506008839080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff167f158412daecdc1456d01568828bcdb18464cc7f1ce0215ddbc3f3cfede9d1e63d60405160405180910390a2505050565b60036020528060005260406000206000915090505481565b60076020528060005260406000206000915090505481565b600f60129054906101000a900467ffffffffffffffff1681565b601060009054906101000a900467ffffffffffffffff1681565b600e6020528060005260406000206000915054906101000a900460ff1681565b60096020528060005260406000206000915054906101000a900460ff1681565b600f60089054906101000a900460ff1681565b611ddb6136af565b73ffffffffffffffffffffffffffffffffffffffff16611df9611fee565b73ffffffffffffffffffffffffffffffffffffffff1614611e4f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e4690615e5a565b60405180910390fd5b611e5960006143c6565b565b60008334611e699190615c51565b9050611eba858285858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610c57565b5050505050565b600f600a9054906101000a900467ffffffffffffffff1681565b611ee36136af565b73ffffffffffffffffffffffffffffffffffffffff16611f01611fee565b73ffffffffffffffffffffffffffffffffffffffff1614611f57576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f4e90615e5a565b60405180910390fd5b80600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60006020528060005260406000206000915054906101000a900460ff1681565b60066020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600b6020528060005260406000206000915054906101000a900460ff1681565b601060089054906101000a900467ffffffffffffffff1681565b61205a6136af565b73ffffffffffffffffffffffffffffffffffffffff16612078611fee565b73ffffffffffffffffffffffffffffffffffffffff16146120ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120c590615e5a565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff16600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361219d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161219490615ec6565b60405180910390fd5b8160011515600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514612231576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122289061644d565b60405180910390fd5b600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690558273ffffffffffffffffffffffffffffffffffffffff167f81ec08d3372506e176c49e626d8beb7e091712ef92908a130f4ccc6524fe2eec60405160405180910390a2505050565b6122d06136af565b73ffffffffffffffffffffffffffffffffffffffff166122ee611fee565b73ffffffffffffffffffffffffffffffffffffffff1614612344576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161233b90615e5a565b60405180910390fd5b60001515601260009054906101000a900460ff1615151461239a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161239190615b2a565b60405180910390fd5b6001601260006101000a81548160ff0219169083151502179055507f915f3053cbc6842207cd97b68c0b585109b4f2fe61c5dbeb25d7678bfdfb8dfa60405160405180910390a1565b600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661246f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161246690615f9e565b60405180910390fd5b61247883613fbc565b6124818361402c565b156125cc5761248f87614125565b81601160008567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506124e683614153565b600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f12b02f226d965a2881e0e8ffed6c421803a22d57ad91f9ef996fe0748ea101758a60008989600f60129054906101000a900467ffffffffffffffff168960405161257c96959493929190616272565b60405180910390a48473ffffffffffffffffffffffffffffffffffffffff166108fc859081150290604051600060405180830381858888f193505050501580156125ca573d6000803e3d6000fd5b505b50505050505050565b600f60009054906101000a900467ffffffffffffffff1681565b600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661267b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161267290615f9e565b60405180910390fd5b61268484613fbc565b61268d8461402c565b156128895761269b89614125565b82601160008667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506126f284614153565b8573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167f12b02f226d965a2881e0e8ffed6c421803a22d57ad91f9ef996fe0748ea101758c60028a8a600f60129054906101000a900467ffffffffffffffff168960405161278796959493929190616272565b60405180910390a4600f60089054906101000a900460ff1615612818578573ffffffffffffffffffffffffffffffffffffffff166350bb4e7f8887856040518463ffffffff1660e01b81526004016127e1939291906164b1565b600060405180830381600087803b1580156127fb57600080fd5b505af115801561280f573d6000803e3d6000fd5b50505050612888565b8573ffffffffffffffffffffffffffffffffffffffff166342842e0e3089886040518463ffffffff1660e01b81526004016128559392919061602a565b600060405180830381600087803b15801561286f57600080fd5b505af1158015612883573d6000803e3d6000fd5b505050505b5b505050505050505050565b6060600d80548060200260200160405190810160405280929190818152602001828054801561291857602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116128ce575b5050505050905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6129506136af565b73ffffffffffffffffffffffffffffffffffffffff1661296e611fee565b73ffffffffffffffffffffffffffffffffffffffff16146129c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129bb90615e5a565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff16600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612a93576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a8a90615ec6565b60405180910390fd5b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690556000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600090556001600880549050612bdd9190615c51565b811015612d135760086001600880549050612bf89190615c51565b81548110612c0957612c086164ef565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660088281548110612c4857612c476164ef565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806007600060088481548110612ca957612ca86164ef565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b6008805480612d2557612d2461651e565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590558273ffffffffffffffffffffffffffffffffffffffff167f1d735ca20b63676dde668b718be78606b061d6bd7534ff815a90a121a6c084b660405160405180910390a2505050565b60025481565b612db06136af565b73ffffffffffffffffffffffffffffffffffffffff16612dce611fee565b73ffffffffffffffffffffffffffffffffffffffff1614612e24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e1b90615e5a565b60405180910390fd5b80600f60096101000a81548160ff02191690831515021790555050565b600d8181548110612e5157600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b612e8d33858486856137eb565b50505050565b612e9b6136af565b73ffffffffffffffffffffffffffffffffffffffff16612eb9611fee565b73ffffffffffffffffffffffffffffffffffffffff1614612f0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f0690615e5a565b60405180910390fd5b600c60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16612f6557600080fd5b600c60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff021916905560005b600d80549050811015613140578173ffffffffffffffffffffffffffffffffffffffff16600d8281548110612fef57612fee6164ef565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361312d57600d6001600d805490506130499190615c51565b8154811061305a576130596164ef565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600d8281548110613099576130986164ef565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600d8054806130f3576130f261651e565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055613140565b80806131389061654d565b915050612fb7565b5050565b565b606060088054806020026020016040519081016040528092919081815260200182805480156131ca57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311613180575b5050505050905090565b6131dc6136af565b73ffffffffffffffffffffffffffffffffffffffff166131fa611fee565b73ffffffffffffffffffffffffffffffffffffffff1614613250576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161324790615e5a565b60405180910390fd5b60008160ff1611613296576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161328d906165e1565b60405180910390fd5b8060ff16600d8054905010156132e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016132d89061664d565b60405180910390fd5b80600e60008460028111156132f9576132f8615c85565b5b60ff1660ff16815260200190815260200160002060006101000a81548160ff021916908360ff1602179055505050565b6133316136af565b73ffffffffffffffffffffffffffffffffffffffff1661334f611fee565b73ffffffffffffffffffffffffffffffffffffffff16146133a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161339c90615e5a565b60405180910390fd5b6133ae8161448c565b50565b6133bf338686868686613bb8565b5050505050565b601260009054906101000a900460ff1681565b6133e16136af565b73ffffffffffffffffffffffffffffffffffffffff166133ff611fee565b73ffffffffffffffffffffffffffffffffffffffff1614613455576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161344c90615e5a565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036134c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134bb906166df565b60405180910390fd5b6134cd816143c6565b50565b600281565b6000806002549050600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415801561353c5750600081115b1561365d5780831015613584576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161357b9061674b565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156135ec573d6000803e3d6000fd5b50600081846135fb9190615c51565b1115613654573373ffffffffffffffffffffffffffffffffffffffff166108fc82856136279190615c51565b9081150290604051600060405180830381858888f19350505050158015613652573d6000803e3d6000fd5b505b809150506136aa565b3373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f193505050501580156136a3573d6000803e3d6000fd5b5060009150505b919050565b600033905090565b60008167ffffffffffffffff16600f60009054906101000a900467ffffffffffffffff1667ffffffffffffffff1614613725576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161371c906167b7565b60405180910390fd5b60008036604051613737929190616807565b6040518091039020905061374d60018483614513565b156137a957600f600081819054906101000a900467ffffffffffffffff168092919061377890615dde565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505060019150506137af565b60009150505b919050565b80600281905550807fa7a33d0996347e1aa55ca2206015b61b9534bdd881d59d59aa680e25eefac36560405160405180910390a250565b84600073ffffffffffffffffffffffffffffffffffffffff16600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036138ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016138b190615ec6565b60405180910390fd5b8560001515600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151461394e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161394590615f32565b60405180910390fd5b600f60099054906101000a900460ff1661399d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161399490615b96565b60405180910390fd5b60008773ffffffffffffffffffffffffffffffffffffffff1663c87b56dd866040518263ffffffff1660e01b81526004016139d8919061540a565b600060405180830381865afa1580156139f5573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190613a1e9190616890565b9050600f60089054906101000a900460ff1615613aa1578773ffffffffffffffffffffffffffffffffffffffff166342966c68866040518263ffffffff1660e01b8152600401613a6e919061540a565b600060405180830381600087803b158015613a8857600080fd5b505af1158015613a9c573d6000803e3d6000fd5b505050505b8773ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f17d76053ca34a4dd8c402fe6498deb797fac89bf7ed02f3f5161aa9368cc8c1f600289600f600a9054906101000a900467ffffffffffffffff1660008b60028a604051602001613b3791906168d9565b604051602081830303815290604052604051613b59979695949392919061697b565b60405180910390a4600f600a81819054906101000a900467ffffffffffffffff1680929190613b8790615dde565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050505050505050565b85600073ffffffffffffffffffffffffffffffffffffffff16600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603613c87576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613c7e90615ec6565b60405180910390fd5b8660001515600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514613d1b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613d1290615f32565b60405180910390fd5b600f60099054906101000a900460ff16613d6a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613d6190615b96565b60405180910390fd5b60008511613dad576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613da490616a44565b60405180910390fd5b6000613dba888a876147e2565b9050600f60089054906101000a900460ff1615613e3d578873ffffffffffffffffffffffffffffffffffffffff166342966c68876040518263ffffffff1660e01b8152600401613e0a919061540a565b600060405180830381600087803b158015613e2457600080fd5b505af1158015613e38573d6000803e3d6000fd5b505050505b8873ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167feff76c36e53fa5ff52f27acc8a34d5047a8246abb07b77b12f1309f71e337f0960018a600f600a9054906101000a900467ffffffffffffffff16878b604051613ed0959493929190615d84565b60405180910390a4600f600a81819054906101000a900467ffffffffffffffff1680929190613efe90615dde565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050505050505050565b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550808273ffffffffffffffffffffffffffffffffffffffff167fdb5ad2e76ae20cfa4e7adbc7305d7538442164d85ead9937c98620a1aa4c255b60405160405180910390a35050565b8067ffffffffffffffff16600f60129054906101000a900467ffffffffffffffff1667ffffffffffffffff161115614029576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161402090616ab0565b60405180910390fd5b50565b6000600b60008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156140a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161409a90616b1c565b60405180910390fd5b600080366040516140b5929190616807565b604051809103902090506140cb60008483614513565b1561411a576001600b60008567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001915050614120565b60009150505b919050565b600160008083815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b601060009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff1611156141af5780601060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505b600060c8600f60129054906101000a900467ffffffffffffffff166141d49190616b3c565b9050601060009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff16111561422257601060009054906101000a900467ffffffffffffffff1690505b6000600f60129054906101000a900467ffffffffffffffff1690505b8167ffffffffffffffff168167ffffffffffffffff16111580156142a757506000601160008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1667ffffffffffffffff16115b1561439857601160008267ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff16601060086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550601160008267ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81549067ffffffffffffffff0219169055600b60008267ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81549060ff0219169055808061439090615dde565b91505061423e565b80600f60126101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f647672599d3468abcfa241a13c9e3d34383caadb5cc80fb67c3cdfcd5f78605960405160405180910390a250565b600080600a600086600281111561452d5761452c615c85565b5b60ff1660ff16815260200190815260200160002060008567ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020905060008160010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000801b81036146225781600001339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061466e565b816003016000828152602001908152602001600020600081819054906101000a900460ff168092919061465490616b7a565b91906101000a81548160ff021916908360ff160217905550505b838260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600082600301600086815260200190815260200160002060009054906101000a900460ff1660ff160361470d57816002018490806001815401808255809150506001900390600052602060002001600090919091909150555b816003016000858152602001908152602001600020600081819054906101000a900460ff168092919061473f90616ba3565b91906101000a81548160ff021916908360ff16021790555050600e600087600281111561476f5761476e615c85565b5b60ff1660ff16815260200190815260200160002060009054906101000a900460ff1660ff1682600301600086815260200190815260200160002060009054906101000a900460ff1660ff16106147d4576147c98686614b5d565b6001925050506147db565b6000925050505b9392505050565b600080600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141580156148865750600081115b15614a9357808310156148ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016148c59061674b565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161492b929190616c21565b6020604051808303816000875af115801561494a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061496e91906160cc565b6149ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016149a490616cbc565b60405180910390fd5b600081846149bb9190615c51565b1115614a8a578373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8683866149ea9190615c51565b6040518363ffffffff1660e01b8152600401614a079291906162da565b6020604051808303816000875af1158015614a26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614a4a91906160cc565b614a89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614a8090616cbc565b60405180910390fd5b5b80915050614b56565b8373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86856040518363ffffffff1660e01b8152600401614ace9291906162da565b6020604051808303816000875af1158015614aed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614b1191906160cc565b614b50576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614b4790616cbc565b60405180910390fd5b60009150505b9392505050565b6000600a6000846002811115614b7657614b75615c85565b5b60ff1660ff16815260200190815260200160002060008367ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020905060005b81600001805490508160ff161015614c6157816001016000836000018360ff1681548110614be557614be46164ef565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600090558080614c5990616ba3565b915050614bb4565b5060005b81600201805490508160ff161015614cd257816003016000836002018360ff1681548110614c9657614c956164ef565b5b9060005260206000200154815260200190815260200160002060006101000a81549060ff02191690558080614cca90616ba3565b915050614c65565b50600a6000846002811115614cea57614ce9615c85565b5b60ff1660ff16815260200190815260200160002060008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008082016000614d349190614d4b565b600282016000614d449190614d6c565b5050505050565b5080546000825590600052602060002090810190614d699190614d8d565b50565b5080546000825590600052602060002090810190614d8a9190614daa565b50565b5b80821115614da6576000816000905550600101614d8e565b5090565b5b80821115614dc3576000816000905550600101614dab565b5090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000614e3582614e0a565b9050919050565b614e4581614e2a565b8114614e5057600080fd5b50565b600081359050614e6281614e3c565b92915050565b600060208284031215614e7e57614e7d614e00565b5b6000614e8c84828501614e53565b91505092915050565b600067ffffffffffffffff82169050919050565b614eb281614e95565b8114614ebd57600080fd5b50565b600081359050614ecf81614ea9565b92915050565b600060208284031215614eeb57614eea614e00565b5b6000614ef984828501614ec0565b91505092915050565b614f0b81614e95565b82525050565b6000602082019050614f266000830184614f02565b92915050565b60008115159050919050565b614f4181614f2c565b82525050565b6000602082019050614f5c6000830184614f38565b92915050565b6000819050919050565b614f7581614f62565b8114614f8057600080fd5b50565b600081359050614f9281614f6c565b92915050565b60008060408385031215614faf57614fae614e00565b5b6000614fbd85828601614f83565b9250506020614fce85828601614ec0565b9150509250929050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b614ffc82614fe2565b810181811067ffffffffffffffff8211171561501b5761501a614dc7565b5b80604052505050565b600061502e614df6565b905061503a8282614ff3565b919050565b600067ffffffffffffffff82111561505a57615059614dc7565b5b61506382614fe2565b9050602081019050919050565b82818337600083830152505050565b600061509261508d8461503f565b615024565b9050828152602081018484840111156150ae576150ad614fdd565b5b6150b9848285615070565b509392505050565b600082601f8301126150d6576150d5614fd8565b5b81356150e684826020860161507f565b91505092915050565b6000806000806080858703121561510957615108614e00565b5b600061511787828801614e53565b945050602061512887828801614e53565b935050604061513987828801614f83565b925050606085013567ffffffffffffffff81111561515a57615159614e05565b5b615166878288016150c1565b91505092959194509250565b600080600080600060a0868803121561518e5761518d614e00565b5b600061519c88828901614e53565b95505060206151ad88828901614e53565b94505060406151be88828901614f83565b93505060606151cf88828901614f83565b925050608086013567ffffffffffffffff8111156151f0576151ef614e05565b5b6151fc888289016150c1565b9150509295509295909350565b60008060006060848603121561522257615221614e00565b5b600061523086828701614e53565b935050602061524186828701614f83565b925050604061525286828701614ec0565b9150509250925092565b61526581614e2a565b82525050565b6000602082019050615280600083018461525c565b92915050565b60006020828403121561529c5761529b614e00565b5b60006152aa84828501614f83565b91505092915050565b6000819050919050565b6152c6816152b3565b81146152d157600080fd5b50565b6000813590506152e3816152bd565b92915050565b600080600080600080600080610100898b03121561530a57615309614e00565b5b60006153188b828c016152d4565b98505060206153298b828c01614e53565b975050604061533a8b828c01614e53565b965050606061534b8b828c01614e53565b955050608061535c8b828c01614f83565b94505060a061536d8b828c01614ec0565b93505060c061537e8b828c01614ec0565b92505060e089013567ffffffffffffffff81111561539f5761539e614e05565b5b6153ab8b828c016150c1565b9150509295985092959890939650565b600080604083850312156153d2576153d1614e00565b5b60006153e085828601614e53565b92505060206153f185828601614e53565b9150509250929050565b61540481614f62565b82525050565b600060208201905061541f60008301846153fb565b92915050565b600060ff82169050919050565b61543b81615425565b811461544657600080fd5b50565b60008135905061545881615432565b92915050565b60006020828403121561547457615473614e00565b5b600061548284828501615449565b91505092915050565b61549481615425565b82525050565b60006020820190506154af600083018461548b565b92915050565b600080fd5b600080fd5b60008083601f8401126154d5576154d4614fd8565b5b8235905067ffffffffffffffff8111156154f2576154f16154b5565b5b60208301915083600182028301111561550e5761550d6154ba565b5b9250929050565b6000806000806060858703121561552f5761552e614e00565b5b600061553d87828801614e53565b945050602061554e87828801614f83565b935050604085013567ffffffffffffffff81111561556f5761556e614e05565b5b61557b878288016154bf565b925092505092959194509250565b60006020828403121561559f5761559e614e00565b5b60006155ad848285016152d4565b91505092915050565b60006155c182614e0a565b9050919050565b6155d1816155b6565b81146155dc57600080fd5b50565b6000813590506155ee816155c8565b92915050565b600080600080600080600060e0888a03121561561357615612614e00565b5b60006156218a828b016152d4565b97505060206156328a828b01614e53565b96505060406156438a828b016155df565b95505060606156548a828b01614f83565b94505060806156658a828b01614ec0565b93505060a06156768a828b01614ec0565b92505060c088013567ffffffffffffffff81111561569757615696614e05565b5b6156a38a828b016150c1565b91505092959891949750929550565b600067ffffffffffffffff8211156156cd576156cc614dc7565b5b6156d682614fe2565b9050602081019050919050565b60006156f66156f1846156b2565b615024565b90508281526020810184848401111561571257615711614fdd565b5b61571d848285615070565b509392505050565b600082601f83011261573a57615739614fd8565b5b813561574a8482602086016156e3565b91505092915050565b60008060008060008060008060006101208a8c03121561577657615775614e00565b5b60006157848c828d016152d4565b99505060206157958c828d01614e53565b98505060406157a68c828d01614e53565b97505060606157b78c828d01614e53565b96505060806157c88c828d01614f83565b95505060a06157d98c828d01614ec0565b94505060c06157ea8c828d01614ec0565b93505060e08a013567ffffffffffffffff81111561580b5761580a614e05565b5b6158178c828d01615725565b9250506101008a013567ffffffffffffffff81111561583957615838614e05565b5b6158458c828d016150c1565b9150509295985092959850929598565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61588a81614e2a565b82525050565b600061589c8383615881565b60208301905092915050565b6000602082019050919050565b60006158c082615855565b6158ca8185615860565b93506158d583615871565b8060005b838110156159065781516158ed8882615890565b97506158f8836158a8565b9250506001810190506158d9565b5085935050505092915050565b6000602082019050818103600083015261592d81846158b5565b905092915050565b61593e816155b6565b82525050565b60006020820190506159596000830184615935565b92915050565b61596881614f2c565b811461597357600080fd5b50565b6000813590506159858161595f565b92915050565b6000602082840312156159a1576159a0614e00565b5b60006159af84828501615976565b91505092915050565b600080600080608085870312156159d2576159d1614e00565b5b60006159e087828801614e53565b94505060206159f187828801614f83565b9350506040615a0287828801614e53565b925050606085013567ffffffffffffffff811115615a2357615a22614e05565b5b615a2f878288016150c1565b91505092959194509250565b60038110615a4857600080fd5b50565b600081359050615a5a81615a3b565b92915050565b60008060408385031215615a7757615a76614e00565b5b6000615a8585828601615a4b565b9250506020615a9685828601615449565b9150509250929050565b600060208284031215615ab657615ab5614e00565b5b6000615ac4848285016155df565b91505092915050565b600082825260208201905092915050565b7f6c6f636b65640000000000000000000000000000000000000000000000000000600082015250565b6000615b14600683615acd565b9150615b1f82615ade565b602082019050919050565b60006020820190508181036000830152615b4381615b07565b9050919050565b7f73746f7070656420627269646765000000000000000000000000000000000000600082015250565b6000615b80600e83615acd565b9150615b8b82615b4a565b602082019050919050565b60006020820190508181036000830152615baf81615b73565b9050919050565b7f696e73756666696369656e7420616d6f756e7400000000000000000000000000600082015250565b6000615bec601383615acd565b9150615bf782615bb6565b602082019050919050565b60006020820190508181036000830152615c1b81615bdf565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000615c5c82614f62565b9150615c6783614f62565b925082821015615c7a57615c79615c22565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110615cc557615cc4615c85565b5b50565b6000819050615cd682615cb4565b919050565b6000615ce682615cc8565b9050919050565b615cf681615cdb565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015615d36578082015181840152602081019050615d1b565b83811115615d45576000848401525b50505050565b6000615d5682615cfc565b615d608185615d07565b9350615d70818560208601615d18565b615d7981614fe2565b840191505092915050565b600060a082019050615d996000830188615ced565b615da660208301876153fb565b615db36040830186614f02565b615dc060608301856153fb565b8181036080830152615dd28184615d4b565b90509695505050505050565b6000615de982614e95565b915067ffffffffffffffff8203615e0357615e02615c22565b5b600182019050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000615e44602083615acd565b9150615e4f82615e0e565b602082019050919050565b60006020820190508181036000830152615e7381615e37565b9050919050565b7f6e6f7420616c6c6f77656420746f6b656e000000000000000000000000000000600082015250565b6000615eb0601183615acd565b9150615ebb82615e7a565b602082019050919050565b60006020820190508181036000830152615edf81615ea3565b9050919050565b7f6c6f636b656420746f6b656e0000000000000000000000000000000000000000600082015250565b6000615f1c600c83615acd565b9150615f2782615ee6565b602082019050919050565b60006020820190508181036000830152615f4b81615f0f565b9050919050565b7f6d73672e73656e646572206973206e6f7420616e206f70657261746f72000000600082015250565b6000615f88601d83615acd565b9150615f9382615f52565b602082019050919050565b60006020820190508181036000830152615fb781615f7b565b9050919050565b7f756e6c6f636b6564000000000000000000000000000000000000000000000000600082015250565b6000615ff4600883615acd565b9150615fff82615fbe565b602082019050919050565b6000602082019050818103600083015261602381615fe7565b9050919050565b600060608201905061603f600083018661525c565b61604c602083018561525c565b61605960408301846153fb565b949350505050565b600061606c82614f62565b915061607783614f62565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156160ac576160ab615c22565b5b828201905092915050565b6000815190506160c68161595f565b92915050565b6000602082840312156160e2576160e1614e00565b5b60006160f0848285016160b7565b91505092915050565b7f7265717565737445524332305472616e736665723a207472616e73666572467260008201527f6f6d206661696c65640000000000000000000000000000000000000000000000602082015250565b6000616155602983615acd565b9150616160826160f9565b604082019050919050565b6000602082019050818103600083015261618481616148565b9050919050565b7f6d6178206f70657261746f72206c696d69740000000000000000000000000000600082015250565b60006161c1601283615acd565b91506161cc8261618b565b602082019050919050565b600060208201905081810360008301526161f0816161b4565b9050919050565b7f6578697374206f70657261746f72000000000000000000000000000000000000600082015250565b600061622d600e83615acd565b9150616238826161f7565b602082019050919050565b6000602082019050818103600083015261625c81616220565b9050919050565b61626c816152b3565b82525050565b600060c0820190506162876000830189616263565b6162946020830188615ced565b6162a160408301876153fb565b6162ae6060830186614f02565b6162bb6080830185614f02565b81810360a08301526162cd8184615d4b565b9050979650505050505050565b60006040820190506162ef600083018561525c565b6162fc60208301846153fb565b9392505050565b7f68616e646c6545524332305472616e736665723a207472616e7366657220666160008201527f696c656400000000000000000000000000000000000000000000000000000000602082015250565b600061635f602483615acd565b915061636a82616303565b604082019050919050565b6000602082019050818103600083015261638e81616352565b9050919050565b7f616c6c6f77656420746f6b656e00000000000000000000000000000000000000600082015250565b60006163cb600d83615acd565b91506163d682616395565b602082019050919050565b600060208201905081810360008301526163fa816163be565b9050919050565b7f756e6c6f636b656420746f6b656e000000000000000000000000000000000000600082015250565b6000616437600e83615acd565b915061644282616401565b602082019050919050565b600060208201905081810360008301526164668161642a565b9050919050565b600081519050919050565b60006164838261646d565b61648d8185615acd565b935061649d818560208601615d18565b6164a681614fe2565b840191505092915050565b60006060820190506164c6600083018661525c565b6164d360208301856153fb565b81810360408301526164e58184616478565b9050949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600061655882614f62565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361658a57616589615c22565b5b600182019050919050565b7f7a65726f207468726573686f6c64000000000000000000000000000000000000600082015250565b60006165cb600e83615acd565b91506165d682616595565b602082019050919050565b600060208201905081810360008301526165fa816165be565b9050919050565b7f626967676572207468616e206e756d206f66206f70657261746f727300000000600082015250565b6000616637601c83615acd565b915061664282616601565b602082019050919050565b600060208201905081810360008301526166668161662a565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006166c9602683615acd565b91506166d48261666d565b604082019050919050565b600060208201905081810360008301526166f8816166bc565b9050919050565b7f696e73756666696369656e74206665654c696d69740000000000000000000000600082015250565b6000616735601583615acd565b9150616740826166ff565b602082019050919050565b6000602082019050818103600083015261676481616728565b9050919050565b7f6e6f6e6365206d69736d61746368000000000000000000000000000000000000600082015250565b60006167a1600e83615acd565b91506167ac8261676b565b602082019050919050565b600060208201905081810360008301526167d081616794565b9050919050565b600081905092915050565b60006167ee83856167d7565b93506167fb838584615070565b82840190509392505050565b60006168148284866167e2565b91508190509392505050565b600061683361682e846156b2565b615024565b90508281526020810184848401111561684f5761684e614fdd565b5b61685a848285615d18565b509392505050565b600082601f83011261687757616876614fd8565b5b8151616887848260208601616820565b91505092915050565b6000602082840312156168a6576168a5614e00565b5b600082015167ffffffffffffffff8111156168c4576168c3614e05565b5b6168d084828501616862565b91505092915050565b600060208201905081810360008301526168f38184616478565b905092915050565b6000819050919050565b6000819050919050565b600061692a616925616920846168fb565b616905565b614f62565b9050919050565b61693a8161690f565b82525050565b6000819050919050565b600061696561696061695b84616940565b616905565b615425565b9050919050565b6169758161694a565b82525050565b600060e082019050616990600083018a615ced565b61699d60208301896153fb565b6169aa6040830188614f02565b6169b76060830187616931565b81810360808301526169c98186615d4b565b90506169d860a083018561696c565b81810360c08301526169ea8184615d4b565b905098975050505050505050565b7f7a65726f206d73672e76616c7565000000000000000000000000000000000000600082015250565b6000616a2e600e83615acd565b9150616a39826169f8565b602082019050919050565b60006020820190508181036000830152616a5d81616a21565b9050919050565b7f72656d6f76656420766f74650000000000000000000000000000000000000000600082015250565b6000616a9a600c83615acd565b9150616aa582616a64565b602082019050919050565b60006020820190508181036000830152616ac981616a8d565b9050919050565b7f636c6f73656420766f7465000000000000000000000000000000000000000000600082015250565b6000616b06600b83615acd565b9150616b1182616ad0565b602082019050919050565b60006020820190508181036000830152616b3581616af9565b9050919050565b6000616b4782614e95565b9150616b5283614e95565b92508267ffffffffffffffff03821115616b6f57616b6e615c22565b5b828201905092915050565b6000616b8582615425565b915060008203616b9857616b97615c22565b5b600182039050919050565b6000616bae82615425565b915060ff8203616bc157616bc0615c22565b5b600182019050919050565b6000616be7616be2616bdd84614e0a565b616905565b614e0a565b9050919050565b6000616bf982616bcc565b9050919050565b6000616c0b82616bee565b9050919050565b616c1b81616c00565b82525050565b6000604082019050616c366000830185616c12565b616c4360208301846153fb565b9392505050565b7f5f7061794552433230466565416e64526566756e644368616e67653a2074726160008201527f6e73666572206661696c65640000000000000000000000000000000000000000602082015250565b6000616ca6602c83615acd565b9150616cb182616c4a565b604082019050919050565b60006020820190508181036000830152616cd581616c99565b905091905056fea2646970667358221220b18a7c48ac36f6475ec334388fb13b4f8dfbd070abc0ae133f06ae064ec4b06f64736f6c634300080e0033"

// DeployBridge deploys a new Klaytn contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _modeMintBurn bool) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBin), backend, _modeMintBurn)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around a Klaytn contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around a Klaytn contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around a Klaytn contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around a Klaytn contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// MAXOPERATOR is a free data retrieval call binding the contract method 0x3a3099d1.
//
// Solidity: function MAX_OPERATOR() view returns(uint64)
func (_Bridge *BridgeCaller) MAXOPERATOR(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "MAX_OPERATOR")
	return *ret0, err
}

// MAXOPERATOR is a free data retrieval call binding the contract method 0x3a3099d1.
//
// Solidity: function MAX_OPERATOR() view returns(uint64)
func (_Bridge *BridgeSession) MAXOPERATOR() (uint64, error) {
	return _Bridge.Contract.MAXOPERATOR(&_Bridge.CallOpts)
}

// MAXOPERATOR is a free data retrieval call binding the contract method 0x3a3099d1.
//
// Solidity: function MAX_OPERATOR() view returns(uint64)
func (_Bridge *BridgeCallerSession) MAXOPERATOR() (uint64, error) {
	return _Bridge.Contract.MAXOPERATOR(&_Bridge.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint64)
func (_Bridge *BridgeCaller) VERSION(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint64)
func (_Bridge *BridgeSession) VERSION() (uint64, error) {
	return _Bridge.Contract.VERSION(&_Bridge.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint64)
func (_Bridge *BridgeCallerSession) VERSION() (uint64, error) {
	return _Bridge.Contract.VERSION(&_Bridge.CallOpts)
}

// ClosedValueTransferVotes is a free data retrieval call binding the contract method 0x9832c1d7.
//
// Solidity: function closedValueTransferVotes(uint64 ) view returns(bool)
func (_Bridge *BridgeCaller) ClosedValueTransferVotes(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "closedValueTransferVotes", arg0)
	return *ret0, err
}

// ClosedValueTransferVotes is a free data retrieval call binding the contract method 0x9832c1d7.
//
// Solidity: function closedValueTransferVotes(uint64 ) view returns(bool)
func (_Bridge *BridgeSession) ClosedValueTransferVotes(arg0 uint64) (bool, error) {
	return _Bridge.Contract.ClosedValueTransferVotes(&_Bridge.CallOpts, arg0)
}

// ClosedValueTransferVotes is a free data retrieval call binding the contract method 0x9832c1d7.
//
// Solidity: function closedValueTransferVotes(uint64 ) view returns(bool)
func (_Bridge *BridgeCallerSession) ClosedValueTransferVotes(arg0 uint64) (bool, error) {
	return _Bridge.Contract.ClosedValueTransferVotes(&_Bridge.CallOpts, arg0)
}

// ConfigurationNonce is a free data retrieval call binding the contract method 0xac6fff0b.
//
// Solidity: function configurationNonce() view returns(uint64)
func (_Bridge *BridgeCaller) ConfigurationNonce(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "configurationNonce")
	return *ret0, err
}

// ConfigurationNonce is a free data retrieval call binding the contract method 0xac6fff0b.
//
// Solidity: function configurationNonce() view returns(uint64)
func (_Bridge *BridgeSession) ConfigurationNonce() (uint64, error) {
	return _Bridge.Contract.ConfigurationNonce(&_Bridge.CallOpts)
}

// ConfigurationNonce is a free data retrieval call binding the contract method 0xac6fff0b.
//
// Solidity: function configurationNonce() view returns(uint64)
func (_Bridge *BridgeCallerSession) ConfigurationNonce() (uint64, error) {
	return _Bridge.Contract.ConfigurationNonce(&_Bridge.CallOpts)
}

// CounterpartBridge is a free data retrieval call binding the contract method 0x3a348533.
//
// Solidity: function counterpartBridge() view returns(address)
func (_Bridge *BridgeCaller) CounterpartBridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "counterpartBridge")
	return *ret0, err
}

// CounterpartBridge is a free data retrieval call binding the contract method 0x3a348533.
//
// Solidity: function counterpartBridge() view returns(address)
func (_Bridge *BridgeSession) CounterpartBridge() (common.Address, error) {
	return _Bridge.Contract.CounterpartBridge(&_Bridge.CallOpts)
}

// CounterpartBridge is a free data retrieval call binding the contract method 0x3a348533.
//
// Solidity: function counterpartBridge() view returns(address)
func (_Bridge *BridgeCallerSession) CounterpartBridge() (common.Address, error) {
	return _Bridge.Contract.CounterpartBridge(&_Bridge.CallOpts)
}

// FeeOfERC20 is a free data retrieval call binding the contract method 0x488af871.
//
// Solidity: function feeOfERC20(address ) view returns(uint256)
func (_Bridge *BridgeCaller) FeeOfERC20(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "feeOfERC20", arg0)
	return *ret0, err
}

// FeeOfERC20 is a free data retrieval call binding the contract method 0x488af871.
//
// Solidity: function feeOfERC20(address ) view returns(uint256)
func (_Bridge *BridgeSession) FeeOfERC20(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.FeeOfERC20(&_Bridge.CallOpts, arg0)
}

// FeeOfERC20 is a free data retrieval call binding the contract method 0x488af871.
//
// Solidity: function feeOfERC20(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) FeeOfERC20(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.FeeOfERC20(&_Bridge.CallOpts, arg0)
}

// FeeOfKLAY is a free data retrieval call binding the contract method 0xc263b5d6.
//
// Solidity: function feeOfKLAY() view returns(uint256)
func (_Bridge *BridgeCaller) FeeOfKLAY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "feeOfKLAY")
	return *ret0, err
}

// FeeOfKLAY is a free data retrieval call binding the contract method 0xc263b5d6.
//
// Solidity: function feeOfKLAY() view returns(uint256)
func (_Bridge *BridgeSession) FeeOfKLAY() (*big.Int, error) {
	return _Bridge.Contract.FeeOfKLAY(&_Bridge.CallOpts)
}

// FeeOfKLAY is a free data retrieval call binding the contract method 0xc263b5d6.
//
// Solidity: function feeOfKLAY() view returns(uint256)
func (_Bridge *BridgeCallerSession) FeeOfKLAY() (*big.Int, error) {
	return _Bridge.Contract.FeeOfKLAY(&_Bridge.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Bridge *BridgeCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "feeReceiver")
	return *ret0, err
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Bridge *BridgeSession) FeeReceiver() (common.Address, error) {
	return _Bridge.Contract.FeeReceiver(&_Bridge.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Bridge *BridgeCallerSession) FeeReceiver() (common.Address, error) {
	return _Bridge.Contract.FeeReceiver(&_Bridge.CallOpts)
}

// GetOperatorList is a free data retrieval call binding the contract method 0xb2c01030.
//
// Solidity: function getOperatorList() view returns(address[])
func (_Bridge *BridgeCaller) GetOperatorList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "getOperatorList")
	return *ret0, err
}

// GetOperatorList is a free data retrieval call binding the contract method 0xb2c01030.
//
// Solidity: function getOperatorList() view returns(address[])
func (_Bridge *BridgeSession) GetOperatorList() ([]common.Address, error) {
	return _Bridge.Contract.GetOperatorList(&_Bridge.CallOpts)
}

// GetOperatorList is a free data retrieval call binding the contract method 0xb2c01030.
//
// Solidity: function getOperatorList() view returns(address[])
func (_Bridge *BridgeCallerSession) GetOperatorList() ([]common.Address, error) {
	return _Bridge.Contract.GetOperatorList(&_Bridge.CallOpts)
}

// GetRegisteredTokenList is a free data retrieval call binding the contract method 0xea21eade.
//
// Solidity: function getRegisteredTokenList() view returns(address[])
func (_Bridge *BridgeCaller) GetRegisteredTokenList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "getRegisteredTokenList")
	return *ret0, err
}

// GetRegisteredTokenList is a free data retrieval call binding the contract method 0xea21eade.
//
// Solidity: function getRegisteredTokenList() view returns(address[])
func (_Bridge *BridgeSession) GetRegisteredTokenList() ([]common.Address, error) {
	return _Bridge.Contract.GetRegisteredTokenList(&_Bridge.CallOpts)
}

// GetRegisteredTokenList is a free data retrieval call binding the contract method 0xea21eade.
//
// Solidity: function getRegisteredTokenList() view returns(address[])
func (_Bridge *BridgeCallerSession) GetRegisteredTokenList() ([]common.Address, error) {
	return _Bridge.Contract.GetRegisteredTokenList(&_Bridge.CallOpts)
}

// HandleNoncesToBlockNums is a free data retrieval call binding the contract method 0x13a6738a.
//
// Solidity: function handleNoncesToBlockNums(uint64 ) view returns(uint64)
func (_Bridge *BridgeCaller) HandleNoncesToBlockNums(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "handleNoncesToBlockNums", arg0)
	return *ret0, err
}

// HandleNoncesToBlockNums is a free data retrieval call binding the contract method 0x13a6738a.
//
// Solidity: function handleNoncesToBlockNums(uint64 ) view returns(uint64)
func (_Bridge *BridgeSession) HandleNoncesToBlockNums(arg0 uint64) (uint64, error) {
	return _Bridge.Contract.HandleNoncesToBlockNums(&_Bridge.CallOpts, arg0)
}

// HandleNoncesToBlockNums is a free data retrieval call binding the contract method 0x13a6738a.
//
// Solidity: function handleNoncesToBlockNums(uint64 ) view returns(uint64)
func (_Bridge *BridgeCallerSession) HandleNoncesToBlockNums(arg0 uint64) (uint64, error) {
	return _Bridge.Contract.HandleNoncesToBlockNums(&_Bridge.CallOpts, arg0)
}

// HandledRequestTx is a free data retrieval call binding the contract method 0x8a75eee2.
//
// Solidity: function handledRequestTx(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) HandledRequestTx(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "handledRequestTx", arg0)
	return *ret0, err
}

// HandledRequestTx is a free data retrieval call binding the contract method 0x8a75eee2.
//
// Solidity: function handledRequestTx(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) HandledRequestTx(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.HandledRequestTx(&_Bridge.CallOpts, arg0)
}

// HandledRequestTx is a free data retrieval call binding the contract method 0x8a75eee2.
//
// Solidity: function handledRequestTx(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) HandledRequestTx(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.HandledRequestTx(&_Bridge.CallOpts, arg0)
}

// IndexOfTokens is a free data retrieval call binding the contract method 0x48a18a6a.
//
// Solidity: function indexOfTokens(address ) view returns(uint256)
func (_Bridge *BridgeCaller) IndexOfTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "indexOfTokens", arg0)
	return *ret0, err
}

// IndexOfTokens is a free data retrieval call binding the contract method 0x48a18a6a.
//
// Solidity: function indexOfTokens(address ) view returns(uint256)
func (_Bridge *BridgeSession) IndexOfTokens(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.IndexOfTokens(&_Bridge.CallOpts, arg0)
}

// IndexOfTokens is a free data retrieval call binding the contract method 0x48a18a6a.
//
// Solidity: function indexOfTokens(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) IndexOfTokens(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.IndexOfTokens(&_Bridge.CallOpts, arg0)
}

// IsLockedKLAY is a free data retrieval call binding the contract method 0xf1719966.
//
// Solidity: function isLockedKLAY() view returns(bool)
func (_Bridge *BridgeCaller) IsLockedKLAY(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "isLockedKLAY")
	return *ret0, err
}

// IsLockedKLAY is a free data retrieval call binding the contract method 0xf1719966.
//
// Solidity: function isLockedKLAY() view returns(bool)
func (_Bridge *BridgeSession) IsLockedKLAY() (bool, error) {
	return _Bridge.Contract.IsLockedKLAY(&_Bridge.CallOpts)
}

// IsLockedKLAY is a free data retrieval call binding the contract method 0xf1719966.
//
// Solidity: function isLockedKLAY() view returns(bool)
func (_Bridge *BridgeCallerSession) IsLockedKLAY() (bool, error) {
	return _Bridge.Contract.IsLockedKLAY(&_Bridge.CallOpts)
}

// IsRunning is a free data retrieval call binding the contract method 0x2014e5d1.
//
// Solidity: function isRunning() view returns(bool)
func (_Bridge *BridgeCaller) IsRunning(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "isRunning")
	return *ret0, err
}

// IsRunning is a free data retrieval call binding the contract method 0x2014e5d1.
//
// Solidity: function isRunning() view returns(bool)
func (_Bridge *BridgeSession) IsRunning() (bool, error) {
	return _Bridge.Contract.IsRunning(&_Bridge.CallOpts)
}

// IsRunning is a free data retrieval call binding the contract method 0x2014e5d1.
//
// Solidity: function isRunning() view returns(bool)
func (_Bridge *BridgeCallerSession) IsRunning() (bool, error) {
	return _Bridge.Contract.IsRunning(&_Bridge.CallOpts)
}

// LockedTokens is a free data retrieval call binding the contract method 0x5eb7413a.
//
// Solidity: function lockedTokens(address ) view returns(bool)
func (_Bridge *BridgeCaller) LockedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "lockedTokens", arg0)
	return *ret0, err
}

// LockedTokens is a free data retrieval call binding the contract method 0x5eb7413a.
//
// Solidity: function lockedTokens(address ) view returns(bool)
func (_Bridge *BridgeSession) LockedTokens(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.LockedTokens(&_Bridge.CallOpts, arg0)
}

// LockedTokens is a free data retrieval call binding the contract method 0x5eb7413a.
//
// Solidity: function lockedTokens(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) LockedTokens(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.LockedTokens(&_Bridge.CallOpts, arg0)
}

// LowerHandleNonce is a free data retrieval call binding the contract method 0x4b40b826.
//
// Solidity: function lowerHandleNonce() view returns(uint64)
func (_Bridge *BridgeCaller) LowerHandleNonce(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "lowerHandleNonce")
	return *ret0, err
}

// LowerHandleNonce is a free data retrieval call binding the contract method 0x4b40b826.
//
// Solidity: function lowerHandleNonce() view returns(uint64)
func (_Bridge *BridgeSession) LowerHandleNonce() (uint64, error) {
	return _Bridge.Contract.LowerHandleNonce(&_Bridge.CallOpts)
}

// LowerHandleNonce is a free data retrieval call binding the contract method 0x4b40b826.
//
// Solidity: function lowerHandleNonce() view returns(uint64)
func (_Bridge *BridgeCallerSession) LowerHandleNonce() (uint64, error) {
	return _Bridge.Contract.LowerHandleNonce(&_Bridge.CallOpts)
}

// ModeMintBurn is a free data retrieval call binding the contract method 0x6e176ec2.
//
// Solidity: function modeMintBurn() view returns(bool)
func (_Bridge *BridgeCaller) ModeMintBurn(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "modeMintBurn")
	return *ret0, err
}

// ModeMintBurn is a free data retrieval call binding the contract method 0x6e176ec2.
//
// Solidity: function modeMintBurn() view returns(bool)
func (_Bridge *BridgeSession) ModeMintBurn() (bool, error) {
	return _Bridge.Contract.ModeMintBurn(&_Bridge.CallOpts)
}

// ModeMintBurn is a free data retrieval call binding the contract method 0x6e176ec2.
//
// Solidity: function modeMintBurn() view returns(bool)
func (_Bridge *BridgeCallerSession) ModeMintBurn() (bool, error) {
	return _Bridge.Contract.ModeMintBurn(&_Bridge.CallOpts)
}

// OperatorList is a free data retrieval call binding the contract method 0xcb38f407.
//
// Solidity: function operatorList(uint256 ) view returns(address)
func (_Bridge *BridgeCaller) OperatorList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "operatorList", arg0)
	return *ret0, err
}

// OperatorList is a free data retrieval call binding the contract method 0xcb38f407.
//
// Solidity: function operatorList(uint256 ) view returns(address)
func (_Bridge *BridgeSession) OperatorList(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.OperatorList(&_Bridge.CallOpts, arg0)
}

// OperatorList is a free data retrieval call binding the contract method 0xcb38f407.
//
// Solidity: function operatorList(uint256 ) view returns(address)
func (_Bridge *BridgeCallerSession) OperatorList(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.OperatorList(&_Bridge.CallOpts, arg0)
}

// OperatorThresholds is a free data retrieval call binding the contract method 0x5526f76b.
//
// Solidity: function operatorThresholds(uint8 ) view returns(uint8)
func (_Bridge *BridgeCaller) OperatorThresholds(opts *bind.CallOpts, arg0 uint8) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "operatorThresholds", arg0)
	return *ret0, err
}

// OperatorThresholds is a free data retrieval call binding the contract method 0x5526f76b.
//
// Solidity: function operatorThresholds(uint8 ) view returns(uint8)
func (_Bridge *BridgeSession) OperatorThresholds(arg0 uint8) (uint8, error) {
	return _Bridge.Contract.OperatorThresholds(&_Bridge.CallOpts, arg0)
}

// OperatorThresholds is a free data retrieval call binding the contract method 0x5526f76b.
//
// Solidity: function operatorThresholds(uint8 ) view returns(uint8)
func (_Bridge *BridgeCallerSession) OperatorThresholds(arg0 uint8) (uint8, error) {
	return _Bridge.Contract.OperatorThresholds(&_Bridge.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Bridge *BridgeCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "operators", arg0)
	return *ret0, err
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Bridge *BridgeSession) Operators(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Operators(&_Bridge.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) Operators(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Operators(&_Bridge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// RecoveryBlockNumber is a free data retrieval call binding the contract method 0x989ba0d3.
//
// Solidity: function recoveryBlockNumber() view returns(uint64)
func (_Bridge *BridgeCaller) RecoveryBlockNumber(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "recoveryBlockNumber")
	return *ret0, err
}

// RecoveryBlockNumber is a free data retrieval call binding the contract method 0x989ba0d3.
//
// Solidity: function recoveryBlockNumber() view returns(uint64)
func (_Bridge *BridgeSession) RecoveryBlockNumber() (uint64, error) {
	return _Bridge.Contract.RecoveryBlockNumber(&_Bridge.CallOpts)
}

// RecoveryBlockNumber is a free data retrieval call binding the contract method 0x989ba0d3.
//
// Solidity: function recoveryBlockNumber() view returns(uint64)
func (_Bridge *BridgeCallerSession) RecoveryBlockNumber() (uint64, error) {
	return _Bridge.Contract.RecoveryBlockNumber(&_Bridge.CallOpts)
}

// RegisteredTokenList is a free data retrieval call binding the contract method 0x3e4fe949.
//
// Solidity: function registeredTokenList(uint256 ) view returns(address)
func (_Bridge *BridgeCaller) RegisteredTokenList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "registeredTokenList", arg0)
	return *ret0, err
}

// RegisteredTokenList is a free data retrieval call binding the contract method 0x3e4fe949.
//
// Solidity: function registeredTokenList(uint256 ) view returns(address)
func (_Bridge *BridgeSession) RegisteredTokenList(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.RegisteredTokenList(&_Bridge.CallOpts, arg0)
}

// RegisteredTokenList is a free data retrieval call binding the contract method 0x3e4fe949.
//
// Solidity: function registeredTokenList(uint256 ) view returns(address)
func (_Bridge *BridgeCallerSession) RegisteredTokenList(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.RegisteredTokenList(&_Bridge.CallOpts, arg0)
}

// RegisteredTokens is a free data retrieval call binding the contract method 0x8c0bd916.
//
// Solidity: function registeredTokens(address ) view returns(address)
func (_Bridge *BridgeCaller) RegisteredTokens(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "registeredTokens", arg0)
	return *ret0, err
}

// RegisteredTokens is a free data retrieval call binding the contract method 0x8c0bd916.
//
// Solidity: function registeredTokens(address ) view returns(address)
func (_Bridge *BridgeSession) RegisteredTokens(arg0 common.Address) (common.Address, error) {
	return _Bridge.Contract.RegisteredTokens(&_Bridge.CallOpts, arg0)
}

// RegisteredTokens is a free data retrieval call binding the contract method 0x8c0bd916.
//
// Solidity: function registeredTokens(address ) view returns(address)
func (_Bridge *BridgeCallerSession) RegisteredTokens(arg0 common.Address) (common.Address, error) {
	return _Bridge.Contract.RegisteredTokens(&_Bridge.CallOpts, arg0)
}

// RequestNonce is a free data retrieval call binding the contract method 0x7c1a0302.
//
// Solidity: function requestNonce() view returns(uint64)
func (_Bridge *BridgeCaller) RequestNonce(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "requestNonce")
	return *ret0, err
}

// RequestNonce is a free data retrieval call binding the contract method 0x7c1a0302.
//
// Solidity: function requestNonce() view returns(uint64)
func (_Bridge *BridgeSession) RequestNonce() (uint64, error) {
	return _Bridge.Contract.RequestNonce(&_Bridge.CallOpts)
}

// RequestNonce is a free data retrieval call binding the contract method 0x7c1a0302.
//
// Solidity: function requestNonce() view returns(uint64)
func (_Bridge *BridgeCallerSession) RequestNonce() (uint64, error) {
	return _Bridge.Contract.RequestNonce(&_Bridge.CallOpts)
}

// UpperHandleNonce is a free data retrieval call binding the contract method 0x54edad72.
//
// Solidity: function upperHandleNonce() view returns(uint64)
func (_Bridge *BridgeCaller) UpperHandleNonce(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "upperHandleNonce")
	return *ret0, err
}

// UpperHandleNonce is a free data retrieval call binding the contract method 0x54edad72.
//
// Solidity: function upperHandleNonce() view returns(uint64)
func (_Bridge *BridgeSession) UpperHandleNonce() (uint64, error) {
	return _Bridge.Contract.UpperHandleNonce(&_Bridge.CallOpts)
}

// UpperHandleNonce is a free data retrieval call binding the contract method 0x54edad72.
//
// Solidity: function upperHandleNonce() view returns(uint64)
func (_Bridge *BridgeCallerSession) UpperHandleNonce() (uint64, error) {
	return _Bridge.Contract.UpperHandleNonce(&_Bridge.CallOpts)
}

// ChargeWithoutEvent is a paid mutator transaction binding the contract method 0xdd9222d6.
//
// Solidity: function chargeWithoutEvent() payable returns()
func (_Bridge *BridgeTransactor) ChargeWithoutEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "chargeWithoutEvent")
}

// ChargeWithoutEvent is a paid mutator transaction binding the contract method 0xdd9222d6.
//
// Solidity: function chargeWithoutEvent() payable returns()
func (_Bridge *BridgeSession) ChargeWithoutEvent() (*types.Transaction, error) {
	return _Bridge.Contract.ChargeWithoutEvent(&_Bridge.TransactOpts)
}

// ChargeWithoutEvent is a paid mutator transaction binding the contract method 0xdd9222d6.
//
// Solidity: function chargeWithoutEvent() payable returns()
func (_Bridge *BridgeTransactorSession) ChargeWithoutEvent() (*types.Transaction, error) {
	return _Bridge.Contract.ChargeWithoutEvent(&_Bridge.TransactOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xd8cf98ca.
//
// Solidity: function deregisterOperator(address _operator) returns()
func (_Bridge *BridgeTransactor) DeregisterOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deregisterOperator", _operator)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xd8cf98ca.
//
// Solidity: function deregisterOperator(address _operator) returns()
func (_Bridge *BridgeSession) DeregisterOperator(_operator common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.DeregisterOperator(&_Bridge.TransactOpts, _operator)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xd8cf98ca.
//
// Solidity: function deregisterOperator(address _operator) returns()
func (_Bridge *BridgeTransactorSession) DeregisterOperator(_operator common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.DeregisterOperator(&_Bridge.TransactOpts, _operator)
}

// DeregisterToken is a paid mutator transaction binding the contract method 0xbab2af1d.
//
// Solidity: function deregisterToken(address _token) returns()
func (_Bridge *BridgeTransactor) DeregisterToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deregisterToken", _token)
}

// DeregisterToken is a paid mutator transaction binding the contract method 0xbab2af1d.
//
// Solidity: function deregisterToken(address _token) returns()
func (_Bridge *BridgeSession) DeregisterToken(_token common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.DeregisterToken(&_Bridge.TransactOpts, _token)
}

// DeregisterToken is a paid mutator transaction binding the contract method 0xbab2af1d.
//
// Solidity: function deregisterToken(address _token) returns()
func (_Bridge *BridgeTransactorSession) DeregisterToken(_token common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.DeregisterToken(&_Bridge.TransactOpts, _token)
}

// HandleERC20Transfer is a paid mutator transaction binding the contract method 0x407e6bae.
//
// Solidity: function handleERC20Transfer(bytes32 _requestTxHash, address _from, address _to, address _tokenAddress, uint256 _value, uint64 _requestedNonce, uint64 _requestedBlockNumber, bytes _extraData) returns()
func (_Bridge *BridgeTransactor) HandleERC20Transfer(opts *bind.TransactOpts, _requestTxHash [32]byte, _from common.Address, _to common.Address, _tokenAddress common.Address, _value *big.Int, _requestedNonce uint64, _requestedBlockNumber uint64, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "handleERC20Transfer", _requestTxHash, _from, _to, _tokenAddress, _value, _requestedNonce, _requestedBlockNumber, _extraData)
}

// HandleERC20Transfer is a paid mutator transaction binding the contract method 0x407e6bae.
//
// Solidity: function handleERC20Transfer(bytes32 _requestTxHash, address _from, address _to, address _tokenAddress, uint256 _value, uint64 _requestedNonce, uint64 _requestedBlockNumber, bytes _extraData) returns()
func (_Bridge *BridgeSession) HandleERC20Transfer(_requestTxHash [32]byte, _from common.Address, _to common.Address, _tokenAddress common.Address, _value *big.Int, _requestedNonce uint64, _requestedBlockNumber uint64, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.HandleERC20Transfer(&_Bridge.TransactOpts, _requestTxHash, _from, _to, _tokenAddress, _value, _requestedNonce, _requestedBlockNumber, _extraData)
}

// HandleERC20Transfer is a paid mutator transaction binding the contract method 0x407e6bae.
//
// Solidity: function handleERC20Transfer(bytes32 _requestTxHash, address _from, address _to, address _tokenAddress, uint256 _value, uint64 _requestedNonce, uint64 _requestedBlockNumber, bytes _extraData) returns()
func (_Bridge *BridgeTransactorSession) HandleERC20Transfer(_requestTxHash [32]byte, _from common.Address, _to common.Address, _tokenAddress common.Address, _value *big.Int, _requestedNonce uint64, _requestedBlockNumber uint64, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.HandleERC20Transfer(&_Bridge.TransactOpts, _requestTxHash, _from, _to, _tokenAddress, _value, _requestedNonce, _requestedBlockNumber, _extraData)
}

// HandleERC721Transfer is a paid mutator transaction binding the contract method 0xafb60223.
//
// Solidity: function handleERC721Transfer(bytes32 _requestTxHash, address _from, address _to, address _tokenAddress, uint256 _tokenId, uint64 _requestedNonce, uint64 _requestedBlockNumber, string _tokenURI, bytes _extraData) returns()
func (_Bridge *BridgeTransactor) HandleERC721Transfer(opts *bind.TransactOpts, _requestTxHash [32]byte, _from common.Address, _to common.Address, _tokenAddress common.Address, _tokenId *big.Int, _requestedNonce uint64, _requestedBlockNumber uint64, _tokenURI string, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "handleERC721Transfer", _requestTxHash, _from, _to, _tokenAddress, _tokenId, _requestedNonce, _requestedBlockNumber, _tokenURI, _extraData)
}

// HandleERC721Transfer is a paid mutator transaction binding the contract method 0xafb60223.
//
// Solidity: function handleERC721Transfer(bytes32 _requestTxHash, address _from, address _to, address _tokenAddress, uint256 _tokenId, uint64 _requestedNonce, uint64 _requestedBlockNumber, string _tokenURI, bytes _extraData) returns()
func (_Bridge *BridgeSession) HandleERC721Transfer(_requestTxHash [32]byte, _from common.Address, _to common.Address, _tokenAddress common.Address, _tokenId *big.Int, _requestedNonce uint64, _requestedBlockNumber uint64, _tokenURI string, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.HandleERC721Transfer(&_Bridge.TransactOpts, _requestTxHash, _from, _to, _tokenAddress, _tokenId, _requestedNonce, _requestedBlockNumber, _tokenURI, _extraData)
}

// HandleERC721Transfer is a paid mutator transaction binding the contract method 0xafb60223.
//
// Solidity: function handleERC721Transfer(bytes32 _requestTxHash, address _from, address _to, address _tokenAddress, uint256 _tokenId, uint64 _requestedNonce, uint64 _requestedBlockNumber, string _tokenURI, bytes _extraData) returns()
func (_Bridge *BridgeTransactorSession) HandleERC721Transfer(_requestTxHash [32]byte, _from common.Address, _to common.Address, _tokenAddress common.Address, _tokenId *big.Int, _requestedNonce uint64, _requestedBlockNumber uint64, _tokenURI string, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.HandleERC721Transfer(&_Bridge.TransactOpts, _requestTxHash, _from, _to, _tokenAddress, _tokenId, _requestedNonce, _requestedBlockNumber, _tokenURI, _extraData)
}

// HandleKLAYTransfer is a paid mutator transaction binding the contract method 0xa066a7ed.
//
// Solidity: function handleKLAYTransfer(bytes32 _requestTxHash, address _from, address _to, uint256 _value, uint64 _requestedNonce, uint64 _requestedBlockNumber, bytes _extraData) returns()
func (_Bridge *BridgeTransactor) HandleKLAYTransfer(opts *bind.TransactOpts, _requestTxHash [32]byte, _from common.Address, _to common.Address, _value *big.Int, _requestedNonce uint64, _requestedBlockNumber uint64, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "handleKLAYTransfer", _requestTxHash, _from, _to, _value, _requestedNonce, _requestedBlockNumber, _extraData)
}

// HandleKLAYTransfer is a paid mutator transaction binding the contract method 0xa066a7ed.
//
// Solidity: function handleKLAYTransfer(bytes32 _requestTxHash, address _from, address _to, uint256 _value, uint64 _requestedNonce, uint64 _requestedBlockNumber, bytes _extraData) returns()
func (_Bridge *BridgeSession) HandleKLAYTransfer(_requestTxHash [32]byte, _from common.Address, _to common.Address, _value *big.Int, _requestedNonce uint64, _requestedBlockNumber uint64, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.HandleKLAYTransfer(&_Bridge.TransactOpts, _requestTxHash, _from, _to, _value, _requestedNonce, _requestedBlockNumber, _extraData)
}

// HandleKLAYTransfer is a paid mutator transaction binding the contract method 0xa066a7ed.
//
// Solidity: function handleKLAYTransfer(bytes32 _requestTxHash, address _from, address _to, uint256 _value, uint64 _requestedNonce, uint64 _requestedBlockNumber, bytes _extraData) returns()
func (_Bridge *BridgeTransactorSession) HandleKLAYTransfer(_requestTxHash [32]byte, _from common.Address, _to common.Address, _value *big.Int, _requestedNonce uint64, _requestedBlockNumber uint64, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.HandleKLAYTransfer(&_Bridge.TransactOpts, _requestTxHash, _from, _to, _value, _requestedNonce, _requestedBlockNumber, _extraData)
}

// LockKLAY is a paid mutator transaction binding the contract method 0x9f071329.
//
// Solidity: function lockKLAY() returns()
func (_Bridge *BridgeTransactor) LockKLAY(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "lockKLAY")
}

// LockKLAY is a paid mutator transaction binding the contract method 0x9f071329.
//
// Solidity: function lockKLAY() returns()
func (_Bridge *BridgeSession) LockKLAY() (*types.Transaction, error) {
	return _Bridge.Contract.LockKLAY(&_Bridge.TransactOpts)
}

// LockKLAY is a paid mutator transaction binding the contract method 0x9f071329.
//
// Solidity: function lockKLAY() returns()
func (_Bridge *BridgeTransactorSession) LockKLAY() (*types.Transaction, error) {
	return _Bridge.Contract.LockKLAY(&_Bridge.TransactOpts)
}

// LockToken is a paid mutator transaction binding the contract method 0x10693fcd.
//
// Solidity: function lockToken(address _token) returns()
func (_Bridge *BridgeTransactor) LockToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "lockToken", _token)
}

// LockToken is a paid mutator transaction binding the contract method 0x10693fcd.
//
// Solidity: function lockToken(address _token) returns()
func (_Bridge *BridgeSession) LockToken(_token common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.LockToken(&_Bridge.TransactOpts, _token)
}

// LockToken is a paid mutator transaction binding the contract method 0x10693fcd.
//
// Solidity: function lockToken(address _token) returns()
func (_Bridge *BridgeTransactorSession) LockToken(_token common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.LockToken(&_Bridge.TransactOpts, _token)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0xf1656e53.
//
// Solidity: function onERC20Received(address _from, address _to, uint256 _value, uint256 _feeLimit, bytes _extraData) returns()
func (_Bridge *BridgeTransactor) OnERC20Received(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "onERC20Received", _from, _to, _value, _feeLimit, _extraData)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0xf1656e53.
//
// Solidity: function onERC20Received(address _from, address _to, uint256 _value, uint256 _feeLimit, bytes _extraData) returns()
func (_Bridge *BridgeSession) OnERC20Received(_from common.Address, _to common.Address, _value *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC20Received(&_Bridge.TransactOpts, _from, _to, _value, _feeLimit, _extraData)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0xf1656e53.
//
// Solidity: function onERC20Received(address _from, address _to, uint256 _value, uint256 _feeLimit, bytes _extraData) returns()
func (_Bridge *BridgeTransactorSession) OnERC20Received(_from common.Address, _to common.Address, _value *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC20Received(&_Bridge.TransactOpts, _from, _to, _value, _feeLimit, _extraData)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0xcf0da290.
//
// Solidity: function onERC721Received(address _from, uint256 _tokenId, address _to, bytes _extraData) returns()
func (_Bridge *BridgeTransactor) OnERC721Received(opts *bind.TransactOpts, _from common.Address, _tokenId *big.Int, _to common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "onERC721Received", _from, _tokenId, _to, _extraData)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0xcf0da290.
//
// Solidity: function onERC721Received(address _from, uint256 _tokenId, address _to, bytes _extraData) returns()
func (_Bridge *BridgeSession) OnERC721Received(_from common.Address, _tokenId *big.Int, _to common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC721Received(&_Bridge.TransactOpts, _from, _tokenId, _to, _extraData)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0xcf0da290.
//
// Solidity: function onERC721Received(address _from, uint256 _tokenId, address _to, bytes _extraData) returns()
func (_Bridge *BridgeTransactorSession) OnERC721Received(_from common.Address, _tokenId *big.Int, _to common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC721Received(&_Bridge.TransactOpts, _from, _tokenId, _to, _extraData)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address _operator) returns()
func (_Bridge *BridgeTransactor) RegisterOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "registerOperator", _operator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address _operator) returns()
func (_Bridge *BridgeSession) RegisterOperator(_operator common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RegisterOperator(&_Bridge.TransactOpts, _operator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address _operator) returns()
func (_Bridge *BridgeTransactorSession) RegisterOperator(_operator common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RegisterOperator(&_Bridge.TransactOpts, _operator)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x4739f7e5.
//
// Solidity: function registerToken(address _token, address _cToken) returns()
func (_Bridge *BridgeTransactor) RegisterToken(opts *bind.TransactOpts, _token common.Address, _cToken common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "registerToken", _token, _cToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x4739f7e5.
//
// Solidity: function registerToken(address _token, address _cToken) returns()
func (_Bridge *BridgeSession) RegisterToken(_token common.Address, _cToken common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RegisterToken(&_Bridge.TransactOpts, _token, _cToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x4739f7e5.
//
// Solidity: function registerToken(address _token, address _cToken) returns()
func (_Bridge *BridgeTransactorSession) RegisterToken(_token common.Address, _cToken common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RegisterToken(&_Bridge.TransactOpts, _token, _cToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RequestERC20Transfer is a paid mutator transaction binding the contract method 0x26c23b54.
//
// Solidity: function requestERC20Transfer(address _tokenAddress, address _to, uint256 _value, uint256 _feeLimit, bytes _extraData) returns()
func (_Bridge *BridgeTransactor) RequestERC20Transfer(opts *bind.TransactOpts, _tokenAddress common.Address, _to common.Address, _value *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "requestERC20Transfer", _tokenAddress, _to, _value, _feeLimit, _extraData)
}

// RequestERC20Transfer is a paid mutator transaction binding the contract method 0x26c23b54.
//
// Solidity: function requestERC20Transfer(address _tokenAddress, address _to, uint256 _value, uint256 _feeLimit, bytes _extraData) returns()
func (_Bridge *BridgeSession) RequestERC20Transfer(_tokenAddress common.Address, _to common.Address, _value *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.RequestERC20Transfer(&_Bridge.TransactOpts, _tokenAddress, _to, _value, _feeLimit, _extraData)
}

// RequestERC20Transfer is a paid mutator transaction binding the contract method 0x26c23b54.
//
// Solidity: function requestERC20Transfer(address _tokenAddress, address _to, uint256 _value, uint256 _feeLimit, bytes _extraData) returns()
func (_Bridge *BridgeTransactorSession) RequestERC20Transfer(_tokenAddress common.Address, _to common.Address, _value *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.RequestERC20Transfer(&_Bridge.TransactOpts, _tokenAddress, _to, _value, _feeLimit, _extraData)
}

// RequestERC721Transfer is a paid mutator transaction binding the contract method 0x22604742.
//
// Solidity: function requestERC721Transfer(address _tokenAddress, address _to, uint256 _tokenId, bytes _extraData) returns()
func (_Bridge *BridgeTransactor) RequestERC721Transfer(opts *bind.TransactOpts, _tokenAddress common.Address, _to common.Address, _tokenId *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "requestERC721Transfer", _tokenAddress, _to, _tokenId, _extraData)
}

// RequestERC721Transfer is a paid mutator transaction binding the contract method 0x22604742.
//
// Solidity: function requestERC721Transfer(address _tokenAddress, address _to, uint256 _tokenId, bytes _extraData) returns()
func (_Bridge *BridgeSession) RequestERC721Transfer(_tokenAddress common.Address, _to common.Address, _tokenId *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.RequestERC721Transfer(&_Bridge.TransactOpts, _tokenAddress, _to, _tokenId, _extraData)
}

// RequestERC721Transfer is a paid mutator transaction binding the contract method 0x22604742.
//
// Solidity: function requestERC721Transfer(address _tokenAddress, address _to, uint256 _tokenId, bytes _extraData) returns()
func (_Bridge *BridgeTransactorSession) RequestERC721Transfer(_tokenAddress common.Address, _to common.Address, _tokenId *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.RequestERC721Transfer(&_Bridge.TransactOpts, _tokenAddress, _to, _tokenId, _extraData)
}

// RequestKLAYTransfer is a paid mutator transaction binding the contract method 0x75ebdc09.
//
// Solidity: function requestKLAYTransfer(address _to, uint256 _value, bytes _extraData) payable returns()
func (_Bridge *BridgeTransactor) RequestKLAYTransfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "requestKLAYTransfer", _to, _value, _extraData)
}

// RequestKLAYTransfer is a paid mutator transaction binding the contract method 0x75ebdc09.
//
// Solidity: function requestKLAYTransfer(address _to, uint256 _value, bytes _extraData) payable returns()
func (_Bridge *BridgeSession) RequestKLAYTransfer(_to common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.RequestKLAYTransfer(&_Bridge.TransactOpts, _to, _value, _extraData)
}

// RequestKLAYTransfer is a paid mutator transaction binding the contract method 0x75ebdc09.
//
// Solidity: function requestKLAYTransfer(address _to, uint256 _value, bytes _extraData) payable returns()
func (_Bridge *BridgeTransactorSession) RequestKLAYTransfer(_to common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Bridge.Contract.RequestKLAYTransfer(&_Bridge.TransactOpts, _to, _value, _extraData)
}

// SetCounterPartBridge is a paid mutator transaction binding the contract method 0x87b04c55.
//
// Solidity: function setCounterPartBridge(address _bridge) returns()
func (_Bridge *BridgeTransactor) SetCounterPartBridge(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setCounterPartBridge", _bridge)
}

// SetCounterPartBridge is a paid mutator transaction binding the contract method 0x87b04c55.
//
// Solidity: function setCounterPartBridge(address _bridge) returns()
func (_Bridge *BridgeSession) SetCounterPartBridge(_bridge common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetCounterPartBridge(&_Bridge.TransactOpts, _bridge)
}

// SetCounterPartBridge is a paid mutator transaction binding the contract method 0x87b04c55.
//
// Solidity: function setCounterPartBridge(address _bridge) returns()
func (_Bridge *BridgeTransactorSession) SetCounterPartBridge(_bridge common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetCounterPartBridge(&_Bridge.TransactOpts, _bridge)
}

// SetERC20Fee is a paid mutator transaction binding the contract method 0x2f88396c.
//
// Solidity: function setERC20Fee(address _token, uint256 _fee, uint64 _requestNonce) returns()
func (_Bridge *BridgeTransactor) SetERC20Fee(opts *bind.TransactOpts, _token common.Address, _fee *big.Int, _requestNonce uint64) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setERC20Fee", _token, _fee, _requestNonce)
}

// SetERC20Fee is a paid mutator transaction binding the contract method 0x2f88396c.
//
// Solidity: function setERC20Fee(address _token, uint256 _fee, uint64 _requestNonce) returns()
func (_Bridge *BridgeSession) SetERC20Fee(_token common.Address, _fee *big.Int, _requestNonce uint64) (*types.Transaction, error) {
	return _Bridge.Contract.SetERC20Fee(&_Bridge.TransactOpts, _token, _fee, _requestNonce)
}

// SetERC20Fee is a paid mutator transaction binding the contract method 0x2f88396c.
//
// Solidity: function setERC20Fee(address _token, uint256 _fee, uint64 _requestNonce) returns()
func (_Bridge *BridgeTransactorSession) SetERC20Fee(_token common.Address, _fee *big.Int, _requestNonce uint64) (*types.Transaction, error) {
	return _Bridge.Contract.SetERC20Fee(&_Bridge.TransactOpts, _token, _fee, _requestNonce)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiver) returns()
func (_Bridge *BridgeTransactor) SetFeeReceiver(opts *bind.TransactOpts, _feeReceiver common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setFeeReceiver", _feeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiver) returns()
func (_Bridge *BridgeSession) SetFeeReceiver(_feeReceiver common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetFeeReceiver(&_Bridge.TransactOpts, _feeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiver) returns()
func (_Bridge *BridgeTransactorSession) SetFeeReceiver(_feeReceiver common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetFeeReceiver(&_Bridge.TransactOpts, _feeReceiver)
}

// SetKLAYFee is a paid mutator transaction binding the contract method 0x1a2ae53e.
//
// Solidity: function setKLAYFee(uint256 _fee, uint64 _requestNonce) returns()
func (_Bridge *BridgeTransactor) SetKLAYFee(opts *bind.TransactOpts, _fee *big.Int, _requestNonce uint64) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setKLAYFee", _fee, _requestNonce)
}

// SetKLAYFee is a paid mutator transaction binding the contract method 0x1a2ae53e.
//
// Solidity: function setKLAYFee(uint256 _fee, uint64 _requestNonce) returns()
func (_Bridge *BridgeSession) SetKLAYFee(_fee *big.Int, _requestNonce uint64) (*types.Transaction, error) {
	return _Bridge.Contract.SetKLAYFee(&_Bridge.TransactOpts, _fee, _requestNonce)
}

// SetKLAYFee is a paid mutator transaction binding the contract method 0x1a2ae53e.
//
// Solidity: function setKLAYFee(uint256 _fee, uint64 _requestNonce) returns()
func (_Bridge *BridgeTransactorSession) SetKLAYFee(_fee *big.Int, _requestNonce uint64) (*types.Transaction, error) {
	return _Bridge.Contract.SetKLAYFee(&_Bridge.TransactOpts, _fee, _requestNonce)
}

// SetOperatorThreshold is a paid mutator transaction binding the contract method 0xee2aec65.
//
// Solidity: function setOperatorThreshold(uint8 _voteType, uint8 _threshold) returns()
func (_Bridge *BridgeTransactor) SetOperatorThreshold(opts *bind.TransactOpts, _voteType uint8, _threshold uint8) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setOperatorThreshold", _voteType, _threshold)
}

// SetOperatorThreshold is a paid mutator transaction binding the contract method 0xee2aec65.
//
// Solidity: function setOperatorThreshold(uint8 _voteType, uint8 _threshold) returns()
func (_Bridge *BridgeSession) SetOperatorThreshold(_voteType uint8, _threshold uint8) (*types.Transaction, error) {
	return _Bridge.Contract.SetOperatorThreshold(&_Bridge.TransactOpts, _voteType, _threshold)
}

// SetOperatorThreshold is a paid mutator transaction binding the contract method 0xee2aec65.
//
// Solidity: function setOperatorThreshold(uint8 _voteType, uint8 _threshold) returns()
func (_Bridge *BridgeTransactorSession) SetOperatorThreshold(_voteType uint8, _threshold uint8) (*types.Transaction, error) {
	return _Bridge.Contract.SetOperatorThreshold(&_Bridge.TransactOpts, _voteType, _threshold)
}

// Start is a paid mutator transaction binding the contract method 0xc877cf37.
//
// Solidity: function start(bool _status) returns()
func (_Bridge *BridgeTransactor) Start(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "start", _status)
}

// Start is a paid mutator transaction binding the contract method 0xc877cf37.
//
// Solidity: function start(bool _status) returns()
func (_Bridge *BridgeSession) Start(_status bool) (*types.Transaction, error) {
	return _Bridge.Contract.Start(&_Bridge.TransactOpts, _status)
}

// Start is a paid mutator transaction binding the contract method 0xc877cf37.
//
// Solidity: function start(bool _status) returns()
func (_Bridge *BridgeTransactorSession) Start(_status bool) (*types.Transaction, error) {
	return _Bridge.Contract.Start(&_Bridge.TransactOpts, _status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// UnlockKLAY is a paid mutator transaction binding the contract method 0x1ebdca38.
//
// Solidity: function unlockKLAY() returns()
func (_Bridge *BridgeTransactor) UnlockKLAY(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "unlockKLAY")
}

// UnlockKLAY is a paid mutator transaction binding the contract method 0x1ebdca38.
//
// Solidity: function unlockKLAY() returns()
func (_Bridge *BridgeSession) UnlockKLAY() (*types.Transaction, error) {
	return _Bridge.Contract.UnlockKLAY(&_Bridge.TransactOpts)
}

// UnlockKLAY is a paid mutator transaction binding the contract method 0x1ebdca38.
//
// Solidity: function unlockKLAY() returns()
func (_Bridge *BridgeTransactorSession) UnlockKLAY() (*types.Transaction, error) {
	return _Bridge.Contract.UnlockKLAY(&_Bridge.TransactOpts)
}

// UnlockToken is a paid mutator transaction binding the contract method 0x9ef2017b.
//
// Solidity: function unlockToken(address _token) returns()
func (_Bridge *BridgeTransactor) UnlockToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "unlockToken", _token)
}

// UnlockToken is a paid mutator transaction binding the contract method 0x9ef2017b.
//
// Solidity: function unlockToken(address _token) returns()
func (_Bridge *BridgeSession) UnlockToken(_token common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.UnlockToken(&_Bridge.TransactOpts, _token)
}

// UnlockToken is a paid mutator transaction binding the contract method 0x9ef2017b.
//
// Solidity: function unlockToken(address _token) returns()
func (_Bridge *BridgeTransactorSession) UnlockToken(_token common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.UnlockToken(&_Bridge.TransactOpts, _token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// BridgeERC20FeeChangedIterator is returned from FilterERC20FeeChanged and is used to iterate over the raw logs and unpacked data for ERC20FeeChanged events raised by the Bridge contract.
type BridgeERC20FeeChangedIterator struct {
	Event *BridgeERC20FeeChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeERC20FeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC20FeeChanged)
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
		it.Event = new(BridgeERC20FeeChanged)
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
func (it *BridgeERC20FeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC20FeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC20FeeChanged represents a ERC20FeeChanged event raised by the Bridge contract.
type BridgeERC20FeeChanged struct {
	Token common.Address
	Fee   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterERC20FeeChanged is a free log retrieval operation binding the contract event 0xdb5ad2e76ae20cfa4e7adbc7305d7538442164d85ead9937c98620a1aa4c255b.
//
// Solidity: event ERC20FeeChanged(address indexed token, uint256 indexed fee)
func (_Bridge *BridgeFilterer) FilterERC20FeeChanged(opts *bind.FilterOpts, token []common.Address, fee []*big.Int) (*BridgeERC20FeeChangedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ERC20FeeChanged", tokenRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20FeeChangedIterator{contract: _Bridge.contract, event: "ERC20FeeChanged", logs: logs, sub: sub}, nil
}

// WatchERC20FeeChanged is a free log subscription operation binding the contract event 0xdb5ad2e76ae20cfa4e7adbc7305d7538442164d85ead9937c98620a1aa4c255b.
//
// Solidity: event ERC20FeeChanged(address indexed token, uint256 indexed fee)
func (_Bridge *BridgeFilterer) WatchERC20FeeChanged(opts *bind.WatchOpts, sink chan<- *BridgeERC20FeeChanged, token []common.Address, fee []*big.Int) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ERC20FeeChanged", tokenRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC20FeeChanged)
				if err := _Bridge.contract.UnpackLog(event, "ERC20FeeChanged", log); err != nil {
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

// ParseERC20FeeChanged is a log parse operation binding the contract event 0xdb5ad2e76ae20cfa4e7adbc7305d7538442164d85ead9937c98620a1aa4c255b.
//
// Solidity: event ERC20FeeChanged(address indexed token, uint256 indexed fee)
func (_Bridge *BridgeFilterer) ParseERC20FeeChanged(log types.Log) (*BridgeERC20FeeChanged, error) {
	event := new(BridgeERC20FeeChanged)
	if err := _Bridge.contract.UnpackLog(event, "ERC20FeeChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeFeeReceiverChangedIterator is returned from FilterFeeReceiverChanged and is used to iterate over the raw logs and unpacked data for FeeReceiverChanged events raised by the Bridge contract.
type BridgeFeeReceiverChangedIterator struct {
	Event *BridgeFeeReceiverChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeFeeReceiverChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeReceiverChanged)
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
		it.Event = new(BridgeFeeReceiverChanged)
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
func (it *BridgeFeeReceiverChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeReceiverChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeReceiverChanged represents a FeeReceiverChanged event raised by the Bridge contract.
type BridgeFeeReceiverChanged struct {
	FeeReceiver common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeReceiverChanged is a free log retrieval operation binding the contract event 0x647672599d3468abcfa241a13c9e3d34383caadb5cc80fb67c3cdfcd5f786059.
//
// Solidity: event FeeReceiverChanged(address indexed feeReceiver)
func (_Bridge *BridgeFilterer) FilterFeeReceiverChanged(opts *bind.FilterOpts, feeReceiver []common.Address) (*BridgeFeeReceiverChangedIterator, error) {

	var feeReceiverRule []interface{}
	for _, feeReceiverItem := range feeReceiver {
		feeReceiverRule = append(feeReceiverRule, feeReceiverItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "FeeReceiverChanged", feeReceiverRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeReceiverChangedIterator{contract: _Bridge.contract, event: "FeeReceiverChanged", logs: logs, sub: sub}, nil
}

// WatchFeeReceiverChanged is a free log subscription operation binding the contract event 0x647672599d3468abcfa241a13c9e3d34383caadb5cc80fb67c3cdfcd5f786059.
//
// Solidity: event FeeReceiverChanged(address indexed feeReceiver)
func (_Bridge *BridgeFilterer) WatchFeeReceiverChanged(opts *bind.WatchOpts, sink chan<- *BridgeFeeReceiverChanged, feeReceiver []common.Address) (event.Subscription, error) {

	var feeReceiverRule []interface{}
	for _, feeReceiverItem := range feeReceiver {
		feeReceiverRule = append(feeReceiverRule, feeReceiverItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "FeeReceiverChanged", feeReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeReceiverChanged)
				if err := _Bridge.contract.UnpackLog(event, "FeeReceiverChanged", log); err != nil {
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

// ParseFeeReceiverChanged is a log parse operation binding the contract event 0x647672599d3468abcfa241a13c9e3d34383caadb5cc80fb67c3cdfcd5f786059.
//
// Solidity: event FeeReceiverChanged(address indexed feeReceiver)
func (_Bridge *BridgeFilterer) ParseFeeReceiverChanged(log types.Log) (*BridgeFeeReceiverChanged, error) {
	event := new(BridgeFeeReceiverChanged)
	if err := _Bridge.contract.UnpackLog(event, "FeeReceiverChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeHandleValueTransferIterator is returned from FilterHandleValueTransfer and is used to iterate over the raw logs and unpacked data for HandleValueTransfer events raised by the Bridge contract.
type BridgeHandleValueTransferIterator struct {
	Event *BridgeHandleValueTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHandleValueTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHandleValueTransfer)
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
		it.Event = new(BridgeHandleValueTransfer)
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
func (it *BridgeHandleValueTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHandleValueTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHandleValueTransfer represents a HandleValueTransfer event raised by the Bridge contract.
type BridgeHandleValueTransfer struct {
	RequestTxHash    [32]byte
	TokenType        uint8
	From             common.Address
	To               common.Address
	TokenAddress     common.Address
	ValueOrTokenId   *big.Int
	HandleNonce      uint64
	LowerHandleNonce uint64
	ExtraData        []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterHandleValueTransfer is a free log retrieval operation binding the contract event 0x12b02f226d965a2881e0e8ffed6c421803a22d57ad91f9ef996fe0748ea10175.
//
// Solidity: event HandleValueTransfer(bytes32 requestTxHash, uint8 tokenType, address indexed from, address indexed to, address indexed tokenAddress, uint256 valueOrTokenId, uint64 handleNonce, uint64 lowerHandleNonce, bytes extraData)
func (_Bridge *BridgeFilterer) FilterHandleValueTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenAddress []common.Address) (*BridgeHandleValueTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "HandleValueTransfer", fromRule, toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHandleValueTransferIterator{contract: _Bridge.contract, event: "HandleValueTransfer", logs: logs, sub: sub}, nil
}

// WatchHandleValueTransfer is a free log subscription operation binding the contract event 0x12b02f226d965a2881e0e8ffed6c421803a22d57ad91f9ef996fe0748ea10175.
//
// Solidity: event HandleValueTransfer(bytes32 requestTxHash, uint8 tokenType, address indexed from, address indexed to, address indexed tokenAddress, uint256 valueOrTokenId, uint64 handleNonce, uint64 lowerHandleNonce, bytes extraData)
func (_Bridge *BridgeFilterer) WatchHandleValueTransfer(opts *bind.WatchOpts, sink chan<- *BridgeHandleValueTransfer, from []common.Address, to []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "HandleValueTransfer", fromRule, toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHandleValueTransfer)
				if err := _Bridge.contract.UnpackLog(event, "HandleValueTransfer", log); err != nil {
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

// ParseHandleValueTransfer is a log parse operation binding the contract event 0x12b02f226d965a2881e0e8ffed6c421803a22d57ad91f9ef996fe0748ea10175.
//
// Solidity: event HandleValueTransfer(bytes32 requestTxHash, uint8 tokenType, address indexed from, address indexed to, address indexed tokenAddress, uint256 valueOrTokenId, uint64 handleNonce, uint64 lowerHandleNonce, bytes extraData)
func (_Bridge *BridgeFilterer) ParseHandleValueTransfer(log types.Log) (*BridgeHandleValueTransfer, error) {
	event := new(BridgeHandleValueTransfer)
	if err := _Bridge.contract.UnpackLog(event, "HandleValueTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeKLAYFeeChangedIterator is returned from FilterKLAYFeeChanged and is used to iterate over the raw logs and unpacked data for KLAYFeeChanged events raised by the Bridge contract.
type BridgeKLAYFeeChangedIterator struct {
	Event *BridgeKLAYFeeChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeKLAYFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeKLAYFeeChanged)
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
		it.Event = new(BridgeKLAYFeeChanged)
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
func (it *BridgeKLAYFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeKLAYFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeKLAYFeeChanged represents a KLAYFeeChanged event raised by the Bridge contract.
type BridgeKLAYFeeChanged struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterKLAYFeeChanged is a free log retrieval operation binding the contract event 0xa7a33d0996347e1aa55ca2206015b61b9534bdd881d59d59aa680e25eefac365.
//
// Solidity: event KLAYFeeChanged(uint256 indexed fee)
func (_Bridge *BridgeFilterer) FilterKLAYFeeChanged(opts *bind.FilterOpts, fee []*big.Int) (*BridgeKLAYFeeChangedIterator, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "KLAYFeeChanged", feeRule)
	if err != nil {
		return nil, err
	}
	return &BridgeKLAYFeeChangedIterator{contract: _Bridge.contract, event: "KLAYFeeChanged", logs: logs, sub: sub}, nil
}

// WatchKLAYFeeChanged is a free log subscription operation binding the contract event 0xa7a33d0996347e1aa55ca2206015b61b9534bdd881d59d59aa680e25eefac365.
//
// Solidity: event KLAYFeeChanged(uint256 indexed fee)
func (_Bridge *BridgeFilterer) WatchKLAYFeeChanged(opts *bind.WatchOpts, sink chan<- *BridgeKLAYFeeChanged, fee []*big.Int) (event.Subscription, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "KLAYFeeChanged", feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeKLAYFeeChanged)
				if err := _Bridge.contract.UnpackLog(event, "KLAYFeeChanged", log); err != nil {
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

// ParseKLAYFeeChanged is a log parse operation binding the contract event 0xa7a33d0996347e1aa55ca2206015b61b9534bdd881d59d59aa680e25eefac365.
//
// Solidity: event KLAYFeeChanged(uint256 indexed fee)
func (_Bridge *BridgeFilterer) ParseKLAYFeeChanged(log types.Log) (*BridgeKLAYFeeChanged, error) {
	event := new(BridgeKLAYFeeChanged)
	if err := _Bridge.contract.UnpackLog(event, "KLAYFeeChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeKLAYLockedIterator is returned from FilterKLAYLocked and is used to iterate over the raw logs and unpacked data for KLAYLocked events raised by the Bridge contract.
type BridgeKLAYLockedIterator struct {
	Event *BridgeKLAYLocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeKLAYLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeKLAYLocked)
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
		it.Event = new(BridgeKLAYLocked)
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
func (it *BridgeKLAYLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeKLAYLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeKLAYLocked represents a KLAYLocked event raised by the Bridge contract.
type BridgeKLAYLocked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterKLAYLocked is a free log retrieval operation binding the contract event 0x915f3053cbc6842207cd97b68c0b585109b4f2fe61c5dbeb25d7678bfdfb8dfa.
//
// Solidity: event KLAYLocked()
func (_Bridge *BridgeFilterer) FilterKLAYLocked(opts *bind.FilterOpts) (*BridgeKLAYLockedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "KLAYLocked")
	if err != nil {
		return nil, err
	}
	return &BridgeKLAYLockedIterator{contract: _Bridge.contract, event: "KLAYLocked", logs: logs, sub: sub}, nil
}

// WatchKLAYLocked is a free log subscription operation binding the contract event 0x915f3053cbc6842207cd97b68c0b585109b4f2fe61c5dbeb25d7678bfdfb8dfa.
//
// Solidity: event KLAYLocked()
func (_Bridge *BridgeFilterer) WatchKLAYLocked(opts *bind.WatchOpts, sink chan<- *BridgeKLAYLocked) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "KLAYLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeKLAYLocked)
				if err := _Bridge.contract.UnpackLog(event, "KLAYLocked", log); err != nil {
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

// ParseKLAYLocked is a log parse operation binding the contract event 0x915f3053cbc6842207cd97b68c0b585109b4f2fe61c5dbeb25d7678bfdfb8dfa.
//
// Solidity: event KLAYLocked()
func (_Bridge *BridgeFilterer) ParseKLAYLocked(log types.Log) (*BridgeKLAYLocked, error) {
	event := new(BridgeKLAYLocked)
	if err := _Bridge.contract.UnpackLog(event, "KLAYLocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeKLAYUnlockedIterator is returned from FilterKLAYUnlocked and is used to iterate over the raw logs and unpacked data for KLAYUnlocked events raised by the Bridge contract.
type BridgeKLAYUnlockedIterator struct {
	Event *BridgeKLAYUnlocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeKLAYUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeKLAYUnlocked)
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
		it.Event = new(BridgeKLAYUnlocked)
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
func (it *BridgeKLAYUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeKLAYUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeKLAYUnlocked represents a KLAYUnlocked event raised by the Bridge contract.
type BridgeKLAYUnlocked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterKLAYUnlocked is a free log retrieval operation binding the contract event 0xd20610c9b78a6903ef134539e3deb5d243be461de6ef12d4c29536bb9b54fa1b.
//
// Solidity: event KLAYUnlocked()
func (_Bridge *BridgeFilterer) FilterKLAYUnlocked(opts *bind.FilterOpts) (*BridgeKLAYUnlockedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "KLAYUnlocked")
	if err != nil {
		return nil, err
	}
	return &BridgeKLAYUnlockedIterator{contract: _Bridge.contract, event: "KLAYUnlocked", logs: logs, sub: sub}, nil
}

// WatchKLAYUnlocked is a free log subscription operation binding the contract event 0xd20610c9b78a6903ef134539e3deb5d243be461de6ef12d4c29536bb9b54fa1b.
//
// Solidity: event KLAYUnlocked()
func (_Bridge *BridgeFilterer) WatchKLAYUnlocked(opts *bind.WatchOpts, sink chan<- *BridgeKLAYUnlocked) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "KLAYUnlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeKLAYUnlocked)
				if err := _Bridge.contract.UnpackLog(event, "KLAYUnlocked", log); err != nil {
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

// ParseKLAYUnlocked is a log parse operation binding the contract event 0xd20610c9b78a6903ef134539e3deb5d243be461de6ef12d4c29536bb9b54fa1b.
//
// Solidity: event KLAYUnlocked()
func (_Bridge *BridgeFilterer) ParseKLAYUnlocked(log types.Log) (*BridgeKLAYUnlocked, error) {
	event := new(BridgeKLAYUnlocked)
	if err := _Bridge.contract.UnpackLog(event, "KLAYUnlocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
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
		it.Event = new(BridgeOwnershipTransferred)
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
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeRequestValueTransferIterator is returned from FilterRequestValueTransfer and is used to iterate over the raw logs and unpacked data for RequestValueTransfer events raised by the Bridge contract.
type BridgeRequestValueTransferIterator struct {
	Event *BridgeRequestValueTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRequestValueTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRequestValueTransfer)
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
		it.Event = new(BridgeRequestValueTransfer)
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
func (it *BridgeRequestValueTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRequestValueTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRequestValueTransfer represents a RequestValueTransfer event raised by the Bridge contract.
type BridgeRequestValueTransfer struct {
	TokenType      uint8
	From           common.Address
	To             common.Address
	TokenAddress   common.Address
	ValueOrTokenId *big.Int
	RequestNonce   uint64
	Fee            *big.Int
	ExtraData      []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRequestValueTransfer is a free log retrieval operation binding the contract event 0xeff76c36e53fa5ff52f27acc8a34d5047a8246abb07b77b12f1309f71e337f09.
//
// Solidity: event RequestValueTransfer(uint8 tokenType, address indexed from, address indexed to, address indexed tokenAddress, uint256 valueOrTokenId, uint64 requestNonce, uint256 fee, bytes extraData)
func (_Bridge *BridgeFilterer) FilterRequestValueTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenAddress []common.Address) (*BridgeRequestValueTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RequestValueTransfer", fromRule, toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRequestValueTransferIterator{contract: _Bridge.contract, event: "RequestValueTransfer", logs: logs, sub: sub}, nil
}

// WatchRequestValueTransfer is a free log subscription operation binding the contract event 0xeff76c36e53fa5ff52f27acc8a34d5047a8246abb07b77b12f1309f71e337f09.
//
// Solidity: event RequestValueTransfer(uint8 tokenType, address indexed from, address indexed to, address indexed tokenAddress, uint256 valueOrTokenId, uint64 requestNonce, uint256 fee, bytes extraData)
func (_Bridge *BridgeFilterer) WatchRequestValueTransfer(opts *bind.WatchOpts, sink chan<- *BridgeRequestValueTransfer, from []common.Address, to []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RequestValueTransfer", fromRule, toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRequestValueTransfer)
				if err := _Bridge.contract.UnpackLog(event, "RequestValueTransfer", log); err != nil {
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

// ParseRequestValueTransfer is a log parse operation binding the contract event 0xeff76c36e53fa5ff52f27acc8a34d5047a8246abb07b77b12f1309f71e337f09.
//
// Solidity: event RequestValueTransfer(uint8 tokenType, address indexed from, address indexed to, address indexed tokenAddress, uint256 valueOrTokenId, uint64 requestNonce, uint256 fee, bytes extraData)
func (_Bridge *BridgeFilterer) ParseRequestValueTransfer(log types.Log) (*BridgeRequestValueTransfer, error) {
	event := new(BridgeRequestValueTransfer)
	if err := _Bridge.contract.UnpackLog(event, "RequestValueTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeRequestValueTransferEncodedIterator is returned from FilterRequestValueTransferEncoded and is used to iterate over the raw logs and unpacked data for RequestValueTransferEncoded events raised by the Bridge contract.
type BridgeRequestValueTransferEncodedIterator struct {
	Event *BridgeRequestValueTransferEncoded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRequestValueTransferEncodedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRequestValueTransferEncoded)
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
		it.Event = new(BridgeRequestValueTransferEncoded)
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
func (it *BridgeRequestValueTransferEncodedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRequestValueTransferEncodedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRequestValueTransferEncoded represents a RequestValueTransferEncoded event raised by the Bridge contract.
type BridgeRequestValueTransferEncoded struct {
	TokenType      uint8
	From           common.Address
	To             common.Address
	TokenAddress   common.Address
	ValueOrTokenId *big.Int
	RequestNonce   uint64
	Fee            *big.Int
	ExtraData      []byte
	EncodingVer    uint8
	EncodedData    []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRequestValueTransferEncoded is a free log retrieval operation binding the contract event 0x17d76053ca34a4dd8c402fe6498deb797fac89bf7ed02f3f5161aa9368cc8c1f.
//
// Solidity: event RequestValueTransferEncoded(uint8 tokenType, address indexed from, address indexed to, address indexed tokenAddress, uint256 valueOrTokenId, uint64 requestNonce, uint256 fee, bytes extraData, uint8 encodingVer, bytes encodedData)
func (_Bridge *BridgeFilterer) FilterRequestValueTransferEncoded(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenAddress []common.Address) (*BridgeRequestValueTransferEncodedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RequestValueTransferEncoded", fromRule, toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRequestValueTransferEncodedIterator{contract: _Bridge.contract, event: "RequestValueTransferEncoded", logs: logs, sub: sub}, nil
}

// WatchRequestValueTransferEncoded is a free log subscription operation binding the contract event 0x17d76053ca34a4dd8c402fe6498deb797fac89bf7ed02f3f5161aa9368cc8c1f.
//
// Solidity: event RequestValueTransferEncoded(uint8 tokenType, address indexed from, address indexed to, address indexed tokenAddress, uint256 valueOrTokenId, uint64 requestNonce, uint256 fee, bytes extraData, uint8 encodingVer, bytes encodedData)
func (_Bridge *BridgeFilterer) WatchRequestValueTransferEncoded(opts *bind.WatchOpts, sink chan<- *BridgeRequestValueTransferEncoded, from []common.Address, to []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RequestValueTransferEncoded", fromRule, toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRequestValueTransferEncoded)
				if err := _Bridge.contract.UnpackLog(event, "RequestValueTransferEncoded", log); err != nil {
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

// ParseRequestValueTransferEncoded is a log parse operation binding the contract event 0x17d76053ca34a4dd8c402fe6498deb797fac89bf7ed02f3f5161aa9368cc8c1f.
//
// Solidity: event RequestValueTransferEncoded(uint8 tokenType, address indexed from, address indexed to, address indexed tokenAddress, uint256 valueOrTokenId, uint64 requestNonce, uint256 fee, bytes extraData, uint8 encodingVer, bytes encodedData)
func (_Bridge *BridgeFilterer) ParseRequestValueTransferEncoded(log types.Log) (*BridgeRequestValueTransferEncoded, error) {
	event := new(BridgeRequestValueTransferEncoded)
	if err := _Bridge.contract.UnpackLog(event, "RequestValueTransferEncoded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeTokenDeregisteredIterator is returned from FilterTokenDeregistered and is used to iterate over the raw logs and unpacked data for TokenDeregistered events raised by the Bridge contract.
type BridgeTokenDeregisteredIterator struct {
	Event *BridgeTokenDeregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeTokenDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenDeregistered)
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
		it.Event = new(BridgeTokenDeregistered)
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
func (it *BridgeTokenDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenDeregistered represents a TokenDeregistered event raised by the Bridge contract.
type BridgeTokenDeregistered struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenDeregistered is a free log retrieval operation binding the contract event 0x1d735ca20b63676dde668b718be78606b061d6bd7534ff815a90a121a6c084b6.
//
// Solidity: event TokenDeregistered(address indexed token)
func (_Bridge *BridgeFilterer) FilterTokenDeregistered(opts *bind.FilterOpts, token []common.Address) (*BridgeTokenDeregisteredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenDeregistered", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenDeregisteredIterator{contract: _Bridge.contract, event: "TokenDeregistered", logs: logs, sub: sub}, nil
}

// WatchTokenDeregistered is a free log subscription operation binding the contract event 0x1d735ca20b63676dde668b718be78606b061d6bd7534ff815a90a121a6c084b6.
//
// Solidity: event TokenDeregistered(address indexed token)
func (_Bridge *BridgeFilterer) WatchTokenDeregistered(opts *bind.WatchOpts, sink chan<- *BridgeTokenDeregistered, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenDeregistered", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenDeregistered)
				if err := _Bridge.contract.UnpackLog(event, "TokenDeregistered", log); err != nil {
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

// ParseTokenDeregistered is a log parse operation binding the contract event 0x1d735ca20b63676dde668b718be78606b061d6bd7534ff815a90a121a6c084b6.
//
// Solidity: event TokenDeregistered(address indexed token)
func (_Bridge *BridgeFilterer) ParseTokenDeregistered(log types.Log) (*BridgeTokenDeregistered, error) {
	event := new(BridgeTokenDeregistered)
	if err := _Bridge.contract.UnpackLog(event, "TokenDeregistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeTokenLockedIterator is returned from FilterTokenLocked and is used to iterate over the raw logs and unpacked data for TokenLocked events raised by the Bridge contract.
type BridgeTokenLockedIterator struct {
	Event *BridgeTokenLocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeTokenLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenLocked)
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
		it.Event = new(BridgeTokenLocked)
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
func (it *BridgeTokenLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenLocked represents a TokenLocked event raised by the Bridge contract.
type BridgeTokenLocked struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenLocked is a free log retrieval operation binding the contract event 0xca1b0a14e18ada4c44846768dd186e35630cdc5cfeaca83c404ae4acaafbecd7.
//
// Solidity: event TokenLocked(address indexed token)
func (_Bridge *BridgeFilterer) FilterTokenLocked(opts *bind.FilterOpts, token []common.Address) (*BridgeTokenLockedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenLocked", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenLockedIterator{contract: _Bridge.contract, event: "TokenLocked", logs: logs, sub: sub}, nil
}

// WatchTokenLocked is a free log subscription operation binding the contract event 0xca1b0a14e18ada4c44846768dd186e35630cdc5cfeaca83c404ae4acaafbecd7.
//
// Solidity: event TokenLocked(address indexed token)
func (_Bridge *BridgeFilterer) WatchTokenLocked(opts *bind.WatchOpts, sink chan<- *BridgeTokenLocked, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenLocked", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenLocked)
				if err := _Bridge.contract.UnpackLog(event, "TokenLocked", log); err != nil {
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

// ParseTokenLocked is a log parse operation binding the contract event 0xca1b0a14e18ada4c44846768dd186e35630cdc5cfeaca83c404ae4acaafbecd7.
//
// Solidity: event TokenLocked(address indexed token)
func (_Bridge *BridgeFilterer) ParseTokenLocked(log types.Log) (*BridgeTokenLocked, error) {
	event := new(BridgeTokenLocked)
	if err := _Bridge.contract.UnpackLog(event, "TokenLocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeTokenRegisteredIterator is returned from FilterTokenRegistered and is used to iterate over the raw logs and unpacked data for TokenRegistered events raised by the Bridge contract.
type BridgeTokenRegisteredIterator struct {
	Event *BridgeTokenRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenRegistered)
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
		it.Event = new(BridgeTokenRegistered)
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
func (it *BridgeTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenRegistered represents a TokenRegistered event raised by the Bridge contract.
type BridgeTokenRegistered struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistered is a free log retrieval operation binding the contract event 0x158412daecdc1456d01568828bcdb18464cc7f1ce0215ddbc3f3cfede9d1e63d.
//
// Solidity: event TokenRegistered(address indexed token)
func (_Bridge *BridgeFilterer) FilterTokenRegistered(opts *bind.FilterOpts, token []common.Address) (*BridgeTokenRegisteredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenRegistered", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenRegisteredIterator{contract: _Bridge.contract, event: "TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenRegistered is a free log subscription operation binding the contract event 0x158412daecdc1456d01568828bcdb18464cc7f1ce0215ddbc3f3cfede9d1e63d.
//
// Solidity: event TokenRegistered(address indexed token)
func (_Bridge *BridgeFilterer) WatchTokenRegistered(opts *bind.WatchOpts, sink chan<- *BridgeTokenRegistered, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenRegistered", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenRegistered)
				if err := _Bridge.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
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

// ParseTokenRegistered is a log parse operation binding the contract event 0x158412daecdc1456d01568828bcdb18464cc7f1ce0215ddbc3f3cfede9d1e63d.
//
// Solidity: event TokenRegistered(address indexed token)
func (_Bridge *BridgeFilterer) ParseTokenRegistered(log types.Log) (*BridgeTokenRegistered, error) {
	event := new(BridgeTokenRegistered)
	if err := _Bridge.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeTokenUnlockedIterator is returned from FilterTokenUnlocked and is used to iterate over the raw logs and unpacked data for TokenUnlocked events raised by the Bridge contract.
type BridgeTokenUnlockedIterator struct {
	Event *BridgeTokenUnlocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeTokenUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenUnlocked)
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
		it.Event = new(BridgeTokenUnlocked)
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
func (it *BridgeTokenUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenUnlocked represents a TokenUnlocked event raised by the Bridge contract.
type BridgeTokenUnlocked struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenUnlocked is a free log retrieval operation binding the contract event 0x81ec08d3372506e176c49e626d8beb7e091712ef92908a130f4ccc6524fe2eec.
//
// Solidity: event TokenUnlocked(address indexed token)
func (_Bridge *BridgeFilterer) FilterTokenUnlocked(opts *bind.FilterOpts, token []common.Address) (*BridgeTokenUnlockedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenUnlocked", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenUnlockedIterator{contract: _Bridge.contract, event: "TokenUnlocked", logs: logs, sub: sub}, nil
}

// WatchTokenUnlocked is a free log subscription operation binding the contract event 0x81ec08d3372506e176c49e626d8beb7e091712ef92908a130f4ccc6524fe2eec.
//
// Solidity: event TokenUnlocked(address indexed token)
func (_Bridge *BridgeFilterer) WatchTokenUnlocked(opts *bind.WatchOpts, sink chan<- *BridgeTokenUnlocked, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenUnlocked", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenUnlocked)
				if err := _Bridge.contract.UnpackLog(event, "TokenUnlocked", log); err != nil {
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

// ParseTokenUnlocked is a log parse operation binding the contract event 0x81ec08d3372506e176c49e626d8beb7e091712ef92908a130f4ccc6524fe2eec.
//
// Solidity: event TokenUnlocked(address indexed token)
func (_Bridge *BridgeFilterer) ParseTokenUnlocked(log types.Log) (*BridgeTokenUnlocked, error) {
	event := new(BridgeTokenUnlocked)
	if err := _Bridge.contract.UnpackLog(event, "TokenUnlocked", log); err != nil {
		return nil, err
	}
	return event, nil
}
