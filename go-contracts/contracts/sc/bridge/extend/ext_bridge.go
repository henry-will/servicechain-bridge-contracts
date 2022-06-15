// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package extbridge

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

// ExtbridgeABI is the input ABI used to generate the binding from.
const ExtbridgeABI = "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_modeMintBurn\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ERC20FeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"}],\"name\":\"FeeReceiverChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestTxHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumBridgeTransfer.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueOrTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"handleNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lowerHandleNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"HandleValueTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"KLAYFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumBridgeTransfer.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueOrTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"requestNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"RequestValueTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumBridgeTransfer.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueOrTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"requestNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"encodingVer\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedData\",\"type\":\"bytes\"}],\"name\":\"RequestValueTransferEncoded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenUnlocked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_OPERATOR\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callback\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"closedValueTransferVotes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configurationNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"deregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"deregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feeOfERC20\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeOfKLAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceiver\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperatorList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredTokenList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"handleNoncesToBlockNums\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"handledRequestTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"indexOfTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRunning\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"lockToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lowerHandleNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modeMintBurn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"onERC20Received\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"operatorThresholds\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoveryBlockNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registeredTokenList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_requestNonce\",\"type\":\"uint64\"}],\"name\":\"setERC20Fee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeReceiver\",\"type\":\"address\"}],\"name\":\"setFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumBridgeOperator.VoteType\",\"name\":\"_voteType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"}],\"name\":\"setOperatorThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"unlockToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upperHandleNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"requestSellERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"requestERC20Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"requestSellERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"requestERC721Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_requestNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_requestBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"handleERC20Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_requestNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_requestBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"handleERC721Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExtbridgeBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const ExtbridgeBinRuntime = ``

// ExtbridgeBin is the compiled bytecode used for deploying new contracts.
var ExtbridgeBin = "0x60806040526000600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060006002556000600e60086101000a81548160ff0219169083151502179055506001600e60096101000a81548160ff0219169083151502179055506001600f60086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000601160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405162006978380380620069788339818101604052810190620001129190620003d5565b80600080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050620001776200016b620002c560201b60201c565b620002cd60201b60201c565b60005b60028081111562000190576200018f62000407565b5b60ff168160ff161015620001e7576001600d60008360ff1660ff16815260200190815260200160002060006101000a81548160ff021916908360ff1602179055508080620001de9062000472565b9150506200017a565b506001600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600c339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600e60086101000a81548160ff0219169083151502179055505050620004a0565b600033905090565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b60008115159050919050565b620003af8162000398565b8114620003bb57600080fd5b50565b600081519050620003cf81620003a4565b92915050565b600060208284031215620003ee57620003ed62000393565b5b6000620003fe84828501620003be565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600060ff82169050919050565b60006200047f8262000465565b915060ff820362000495576200049462000436565b5b600182019050919050565b6164c880620004b06000396000f3fe608060405234801561001057600080fd5b506004361061028a5760003560e01c80638a75eee21161015c578063bab2af1d116100ce578063d8cf98ca11610087578063d8cf98ca146107d3578063ea21eade146107ef578063ee2aec651461080d578063efdcd97414610829578063f1656e5314610845578063f2fde38b146108615761028a565b8063bab2af1d14610715578063c263b5d614610731578063c5e490731461074f578063c877cf371461076b578063cb38f40714610787578063cf0da290146107b75761028a565b8063989ba0d311610120578063989ba0d3146106655780639ef2017b14610683578063ac6fff0b1461069f578063afb60223146106bd578063b2c01030146106d9578063b3f00674146106f75761028a565b80638a75eee21461059b5780638c0bd916146105cb5780638da5cb5b146105fb5780638daa63ac146106195780639832c1d7146106355761028a565b8063407e6bae1161020057806354edad72116101b957806354edad72146104d75780635526f76b146104f55780635eb7413a146105255780636e176ec214610555578063715018a6146105735780637c1a03021461057d5761028a565b8063407e6bae146104055780634739f7e514610421578063488af8711461043d57806348a18a6a1461046d5780634b40b8261461049d5780634c5146f5146104bb5761028a565b80632260474211610252578063226047421461034757806326c23b54146103635780632f88396c1461037f5780633682a4501461039b5780633a3099d1146103b75780633e4fe949146103d55761028a565b8063083b27321461028f57806310693fcd146102ad57806313a6738a146102c957806313e7c9d8146102f95780632014e5d114610329575b600080fd5b61029761087d565b6040516102a491906145bb565b60405180910390f35b6102c760048036038101906102c29190614616565b6108a3565b005b6102e360048036038101906102de9190614683565b610b22565b6040516102f091906146bf565b60405180910390f35b610313600480360381019061030e9190614616565b610b49565b60405161032091906146f5565b60405180910390f35b610331610b69565b60405161033e91906146f5565b60405180910390f35b610361600480360381019061035c919061488c565b610b7c565b005b61037d6004803603810190610378919061490f565b610bb7565b005b610399600480360381019061039491906149a6565b610bf2565b005b6103b560048036038101906103b09190614616565b610c9c565b005b6103bf610eb5565b6040516103cc91906146bf565b60405180910390f35b6103ef60048036038101906103ea91906149f9565b610eba565b6040516103fc91906145bb565b60405180910390f35b61041f600480360381019061041a9190614a5c565b610ef9565b005b61043b60048036038101906104369190614b2e565b6110f9565b005b61045760048036038101906104529190614616565b6113b6565b6040516104649190614b7d565b60405180910390f35b61048760048036038101906104829190614616565b6113ce565b6040516104949190614b7d565b60405180910390f35b6104a56113e6565b6040516104b291906146bf565b60405180910390f35b6104d560048036038101906104d09190614b98565b611400565b005b6104df611431565b6040516104ec91906146bf565b60405180910390f35b61050f600480360381019061050a9190614c38565b61144b565b60405161051c9190614c74565b60405180910390f35b61053f600480360381019061053a9190614616565b61146b565b60405161054c91906146f5565b60405180910390f35b61055d61148b565b60405161056a91906146f5565b60405180910390f35b61057b61149e565b005b610585611526565b60405161059291906146bf565b60405180910390f35b6105b560048036038101906105b09190614c8f565b611540565b6040516105c291906146f5565b60405180910390f35b6105e560048036038101906105e09190614616565b611560565b6040516105f291906145bb565b60405180910390f35b610603611593565b60405161061091906145bb565b60405180910390f35b610633600480360381019061062e9190614616565b6115bd565b005b61064f600480360381019061064a9190614683565b61167d565b60405161065c91906146f5565b60405180910390f35b61066d61169d565b60405161067a91906146bf565b60405180910390f35b61069d60048036038101906106989190614616565b6116b7565b005b6106a761192d565b6040516106b491906146bf565b60405180910390f35b6106d760048036038101906106d29190614d5d565b611947565b005b6106e1611b49565b6040516106ee9190614f1d565b60405180910390f35b6106ff611bd7565b60405161070c9190614f60565b60405180910390f35b61072f600480360381019061072a9190614616565b611bfd565b005b610739612057565b6040516107469190614b7d565b60405180910390f35b61076960048036038101906107649190614f7b565b61205d565b005b61078560048036038101906107809190615022565b612090565b005b6107a1600480360381019061079c91906149f9565b612129565b6040516107ae91906145bb565b60405180910390f35b6107d160048036038101906107cc919061504f565b612168565b005b6107ed60048036038101906107e89190614616565b61217b565b005b6107f761242c565b6040516108049190614f1d565b60405180910390f35b610827600480360381019061082291906150f7565b6124ba565b005b610843600480360381019061083e9190615163565b61260f565b005b61085f600480360381019061085a919061490f565b612697565b005b61087b60048036038101906108769190614616565b6126ac565b005b601160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6108ab6127a3565b73ffffffffffffffffffffffffffffffffffffffff166108c9611593565b73ffffffffffffffffffffffffffffffffffffffff161461091f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610916906151ed565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff16600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036109ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e590615259565b60405180910390fd5b8160001515600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514610a82576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a79906152c5565b60405180910390fd5b6001600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508273ffffffffffffffffffffffffffffffffffffffff167fca1b0a14e18ada4c44846768dd186e35630cdc5cfeaca83c404ae4acaafbecd760405160405180910390a2505050565b60106020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b600b6020528060005260406000206000915054906101000a900460ff1681565b600e60099054906101000a900460ff1681565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bae90615331565b60405180910390fd5b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be990615331565b60405180910390fd5b600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610c7e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c759061539d565b60405180910390fd5b610c87816127ab565b15610c9757610c9683836128a8565b5b505050565b610ca46127a3565b73ffffffffffffffffffffffffffffffffffffffff16610cc2611593565b73ffffffffffffffffffffffffffffffffffffffff1614610d18576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0f906151ed565b60405180910390fd5b600c67ffffffffffffffff16600c8054905010610d6a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d6190615409565b60405180910390fd5b600b60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610df7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dee90615475565b60405180910390fd5b6001600b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600c819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600c81565b60078181548110610eca57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6020815114610f3d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f34906154e1565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16601160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610fce576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fc59061554d565b60405180910390fd5b600081806020019051810190610fe49190615582565b905060008111611029576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611020906155fb565b60405180910390fd5b61105b8989601160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168989898989612934565b601160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635ec467e2888789856040518563ffffffff1660e01b81526004016110bc949392919061561b565b600060405180830381600087803b1580156110d657600080fd5b505af11580156110ea573d6000803e3d6000fd5b50505050505050505050505050565b6111016127a3565b73ffffffffffffffffffffffffffffffffffffffff1661111f611593565b73ffffffffffffffffffffffffffffffffffffffff1614611175576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161116c906151ed565b60405180910390fd5b81600073ffffffffffffffffffffffffffffffffffffffff16600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611244576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161123b906156ac565b60405180910390fd5b81600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600780549050600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506007839080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff167f158412daecdc1456d01568828bcdb18464cc7f1ce0215ddbc3f3cfede9d1e63d60405160405180910390a2505050565b60036020528060005260406000206000915090505481565b60066020528060005260406000206000915090505481565b600e60129054906101000a900467ffffffffffffffff1681565b61142b848484846040516020016114179190614b7d565b604051602081830303815290604052612c24565b50505050565b600f60009054906101000a900467ffffffffffffffff1681565b600d6020528060005260406000206000915054906101000a900460ff1681565b60086020528060005260406000206000915054906101000a900460ff1681565b600e60089054906101000a900460ff1681565b6114a66127a3565b73ffffffffffffffffffffffffffffffffffffffff166114c4611593565b73ffffffffffffffffffffffffffffffffffffffff161461151a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611511906151ed565b60405180910390fd5b6115246000612ca6565b565b600e600a9054906101000a900467ffffffffffffffff1681565b60006020528060005260406000206000915054906101000a900460ff1681565b60056020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6115c56127a3565b73ffffffffffffffffffffffffffffffffffffffff166115e3611593565b73ffffffffffffffffffffffffffffffffffffffff1614611639576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611630906151ed565b60405180910390fd5b80601160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600a6020528060005260406000206000915054906101000a900460ff1681565b600f60089054906101000a900467ffffffffffffffff1681565b6116bf6127a3565b73ffffffffffffffffffffffffffffffffffffffff166116dd611593565b73ffffffffffffffffffffffffffffffffffffffff1614611733576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161172a906151ed565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff16600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611802576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117f990615259565b60405180910390fd5b8160011515600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514611896576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161188d90615718565b60405180910390fd5b600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690558273ffffffffffffffffffffffffffffffffffffffff167f81ec08d3372506e176c49e626d8beb7e091712ef92908a130f4ccc6524fe2eec60405160405180910390a2505050565b600e60009054906101000a900467ffffffffffffffff1681565b602081511461198b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611982906154e1565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16601160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611a1c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a139061554d565b60405180910390fd5b600081806020019051810190611a329190615582565b905060008111611a77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a6e906155fb565b60405180910390fd5b611aaa8a8a601160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168a8a8a8a8a8a612d6c565b601160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635ec467e289888a856040518563ffffffff1660e01b8152600401611b0b949392919061561b565b600060405180830381600087803b158015611b2557600080fd5b505af1158015611b39573d6000803e3d6000fd5b5050505050505050505050505050565b6060600c805480602002602001604051908101604052809291908181526020018280548015611bcd57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611b83575b5050505050905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611c056127a3565b73ffffffffffffffffffffffffffffffffffffffff16611c23611593565b73ffffffffffffffffffffffffffffffffffffffff1614611c79576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c70906151ed565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff16600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611d48576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d3f90615259565b60405180910390fd5b600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690556000600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600090556001600780549050611e929190615767565b811015611fc85760076001600780549050611ead9190615767565b81548110611ebe57611ebd61579b565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660078281548110611efd57611efc61579b565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806006600060078481548110611f5e57611f5d61579b565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b6007805480611fda57611fd96157ca565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590558273ffffffffffffffffffffffffffffffffffffffff167f1d735ca20b63676dde668b718be78606b061d6bd7534ff815a90a121a6c084b660405160405180910390a2505050565b60025481565b61208985858585856040516020016120759190614b7d565b604051602081830303815290604052613011565b5050505050565b6120986127a3565b73ffffffffffffffffffffffffffffffffffffffff166120b6611593565b73ffffffffffffffffffffffffffffffffffffffff161461210c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612103906151ed565b60405180910390fd5b80600e60096101000a81548160ff02191690831515021790555050565b600c818154811061213957600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61217533858486856130f0565b50505050565b6121836127a3565b73ffffffffffffffffffffffffffffffffffffffff166121a1611593565b73ffffffffffffffffffffffffffffffffffffffff16146121f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121ee906151ed565b60405180910390fd5b600b60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661224d57600080fd5b600b60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff021916905560005b600c80549050811015612428578173ffffffffffffffffffffffffffffffffffffffff16600c82815481106122d7576122d661579b565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361241557600c6001600c805490506123319190615767565b815481106123425761234161579b565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600c82815481106123815761238061579b565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600c8054806123db576123da6157ca565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055612428565b8080612420906157f9565b91505061229f565b5050565b606060078054806020026020016040519081016040528092919081815260200182805480156124b057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612466575b5050505050905090565b6124c26127a3565b73ffffffffffffffffffffffffffffffffffffffff166124e0611593565b73ffffffffffffffffffffffffffffffffffffffff1614612536576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161252d906151ed565b60405180910390fd5b60008160ff161161257c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125739061588d565b60405180910390fd5b8060ff16600c8054905010156125c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125be906158f9565b60405180910390fd5b80600d60008460028111156125df576125de615919565b5b60ff1660ff16815260200190815260200160002060006101000a81548160ff021916908360ff1602179055505050565b6126176127a3565b73ffffffffffffffffffffffffffffffffffffffff16612635611593565b73ffffffffffffffffffffffffffffffffffffffff161461268b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612682906151ed565b60405180910390fd5b612694816134bd565b50565b6126a5338686868686613544565b5050505050565b6126b46127a3565b73ffffffffffffffffffffffffffffffffffffffff166126d2611593565b73ffffffffffffffffffffffffffffffffffffffff1614612728576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161271f906151ed565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612797576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161278e906159ba565b60405180910390fd5b6127a081612ca6565b50565b600033905090565b60008167ffffffffffffffff16600e60009054906101000a900467ffffffffffffffff1667ffffffffffffffff1614612819576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161281090615a26565b60405180910390fd5b6000803660405161282b929190615a76565b60405180910390209050612841600184836138bc565b1561289d57600e600081819054906101000a900467ffffffffffffffff168092919061286c90615a8f565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505060019150506128a3565b60009150505b919050565b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550808273ffffffffffffffffffffffffffffffffffffffff167fdb5ad2e76ae20cfa4e7adbc7305d7538442164d85ead9937c98620a1aa4c255b60405160405180910390a35050565b600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166129c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129b79061539d565b60405180910390fd5b6129c983613b8b565b6129d283613bfb565b15612c1a576129e088613cf4565b81601060008567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550612a3783613d22565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f12b02f226d965a2881e0e8ffed6c421803a22d57ad91f9ef996fe0748ea101758b60018989600e60129054906101000a900467ffffffffffffffff1689604051612acc96959493929190615b9e565b60405180910390a4600e60089054906101000a900460ff1615612b5b578473ffffffffffffffffffffffffffffffffffffffff166340c10f1987866040518363ffffffff1660e01b8152600401612b24929190615c06565b600060405180830381600087803b158015612b3e57600080fd5b505af1158015612b52573d6000803e3d6000fd5b50505050612c19565b8473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb87866040518363ffffffff1660e01b8152600401612b96929190615c06565b6020604051808303816000875af1158015612bb5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bd99190615c44565b612c18576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c0f90615ce3565b60405180910390fd5b5b5b5050505050505050565b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b8152600401612c6193929190615d03565b600060405180830381600087803b158015612c7b57600080fd5b505af1158015612c8f573d6000803e3d6000fd5b50505050612ca084338585856130f0565b50505050565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16612df8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612def9061539d565b60405180910390fd5b612e0184613b8b565b612e0a84613bfb565b1561300657612e1889613cf4565b82601060008667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550612e6f84613d22565b8573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167f12b02f226d965a2881e0e8ffed6c421803a22d57ad91f9ef996fe0748ea101758c60028a8a600e60129054906101000a900467ffffffffffffffff1689604051612f0496959493929190615b9e565b60405180910390a4600e60089054906101000a900460ff1615612f95578573ffffffffffffffffffffffffffffffffffffffff166350bb4e7f8887856040518463ffffffff1660e01b8152600401612f5e93929190615d7e565b600060405180830381600087803b158015612f7857600080fd5b505af1158015612f8c573d6000803e3d6000fd5b50505050613005565b8573ffffffffffffffffffffffffffffffffffffffff166342842e0e3089886040518463ffffffff1660e01b8152600401612fd293929190615d03565b600060405180830381600087803b158015612fec57600080fd5b505af1158015613000573d6000803e3d6000fd5b505050505b5b505050505050505050565b8473ffffffffffffffffffffffffffffffffffffffff166323b872dd3330858761303b9190615dbc565b6040518463ffffffff1660e01b815260040161305993929190615d03565b6020604051808303816000875af1158015613078573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061309c9190615c44565b6130db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130d290615e84565b60405180910390fd5b6130e9853386868686613544565b5050505050565b84600073ffffffffffffffffffffffffffffffffffffffff16600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036131bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131b690615259565b60405180910390fd5b8560001515600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514613253576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161324a906152c5565b60405180910390fd5b600e60099054906101000a900460ff166132a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161329990615ef0565b60405180910390fd5b60008773ffffffffffffffffffffffffffffffffffffffff1663c87b56dd866040518263ffffffff1660e01b81526004016132dd9190614b7d565b600060405180830381865afa1580156132fa573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906133239190615f80565b9050600e60089054906101000a900460ff16156133a6578773ffffffffffffffffffffffffffffffffffffffff166342966c68866040518263ffffffff1660e01b81526004016133739190614b7d565b600060405180830381600087803b15801561338d57600080fd5b505af11580156133a1573d6000803e3d6000fd5b505050505b8773ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f17d76053ca34a4dd8c402fe6498deb797fac89bf7ed02f3f5161aa9368cc8c1f600289600e600a9054906101000a900467ffffffffffffffff1660008b60028a60405160200161343c9190615fc9565b60405160208183030381529060405260405161345e979695949392919061606b565b60405180910390a4600e600a81819054906101000a900467ffffffffffffffff168092919061348c90615a8f565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050505050505050565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f647672599d3468abcfa241a13c9e3d34383caadb5cc80fb67c3cdfcd5f78605960405160405180910390a250565b85600073ffffffffffffffffffffffffffffffffffffffff16600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603613613576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161360a90615259565b60405180910390fd5b8660001515600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515146136a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161369e906152c5565b60405180910390fd5b600e60099054906101000a900460ff166136f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136ed90615ef0565b60405180910390fd5b60008511613739576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161373090616134565b60405180910390fd5b6000613746888a87613f95565b9050600e60089054906101000a900460ff16156137c9578873ffffffffffffffffffffffffffffffffffffffff166342966c68876040518263ffffffff1660e01b81526004016137969190614b7d565b600060405180830381600087803b1580156137b057600080fd5b505af11580156137c4573d6000803e3d6000fd5b505050505b8873ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167feff76c36e53fa5ff52f27acc8a34d5047a8246abb07b77b12f1309f71e337f0960018a600e600a9054906101000a900467ffffffffffffffff16878b60405161385c959493929190616154565b60405180910390a4600e600a81819054906101000a900467ffffffffffffffff168092919061388a90615a8f565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050505050505050565b600080600960008660028111156138d6576138d5615919565b5b60ff1660ff16815260200190815260200160002060008567ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020905060008160010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000801b81036139cb5781600001339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550613a17565b816003016000828152602001908152602001600020600081819054906101000a900460ff16809291906139fd906161ae565b91906101000a81548160ff021916908360ff160217905550505b838260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600082600301600086815260200190815260200160002060009054906101000a900460ff1660ff1603613ab657816002018490806001815401808255809150506001900390600052602060002001600090919091909150555b816003016000858152602001908152602001600020600081819054906101000a900460ff1680929190613ae8906161d7565b91906101000a81548160ff021916908360ff16021790555050600d6000876002811115613b1857613b17615919565b5b60ff1660ff16815260200190815260200160002060009054906101000a900460ff1660ff1682600301600086815260200190815260200160002060009054906101000a900460ff1660ff1610613b7d57613b728686614310565b600192505050613b84565b6000925050505b9392505050565b8067ffffffffffffffff16600e60129054906101000a900467ffffffffffffffff1667ffffffffffffffff161115613bf8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613bef9061624c565b60405180910390fd5b50565b6000600a60008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615613c72576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613c69906162b8565b60405180910390fd5b60008036604051613c84929190615a76565b60405180910390209050613c9a600084836138bc565b15613ce9576001600a60008567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001915050613cef565b60009150505b919050565b600160008083815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600f60009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161115613d7e5780600f60006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505b600060c8600e60129054906101000a900467ffffffffffffffff16613da391906162d8565b9050600f60009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161115613df157600f60009054906101000a900467ffffffffffffffff1690505b6000600e60129054906101000a900467ffffffffffffffff1690505b8167ffffffffffffffff168167ffffffffffffffff1611158015613e7657506000601060008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1667ffffffffffffffff16115b15613f6757601060008267ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff16600f60086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550601060008267ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81549067ffffffffffffffff0219169055600a60008267ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690558080613f5f90615a8f565b915050613e0d565b80600e60126101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050565b600080600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141580156140395750600081115b156142465780831015614081576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161407890616362565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b81526004016140de9291906163d7565b6020604051808303816000875af11580156140fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141219190615c44565b614160576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161415790616472565b60405180910390fd5b6000818461416e9190615767565b111561423d578373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86838661419d9190615767565b6040518363ffffffff1660e01b81526004016141ba929190615c06565b6020604051808303816000875af11580156141d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141fd9190615c44565b61423c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161423390616472565b60405180910390fd5b5b80915050614309565b8373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86856040518363ffffffff1660e01b8152600401614281929190615c06565b6020604051808303816000875af11580156142a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906142c49190615c44565b614303576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016142fa90616472565b60405180910390fd5b60009150505b9392505050565b60006009600084600281111561432957614328615919565b5b60ff1660ff16815260200190815260200160002060008367ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020905060005b81600001805490508160ff16101561441457816001016000836000018360ff16815481106143985761439761579b565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009055808061440c906161d7565b915050614367565b5060005b81600201805490508160ff16101561448557816003016000836002018360ff16815481106144495761444861579b565b5b9060005260206000200154815260200190815260200160002060006101000a81549060ff0219169055808061447d906161d7565b915050614418565b506009600084600281111561449d5761449c615919565b5b60ff1660ff16815260200190815260200160002060008367ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600080820160006144e791906144fe565b6002820160006144f7919061451f565b5050505050565b508054600082559060005260206000209081019061451c9190614540565b50565b508054600082559060005260206000209081019061453d919061455d565b50565b5b80821115614559576000816000905550600101614541565b5090565b5b8082111561457657600081600090555060010161455e565b5090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006145a58261457a565b9050919050565b6145b58161459a565b82525050565b60006020820190506145d060008301846145ac565b92915050565b6000604051905090565b600080fd5b600080fd5b6145f38161459a565b81146145fe57600080fd5b50565b600081359050614610816145ea565b92915050565b60006020828403121561462c5761462b6145e0565b5b600061463a84828501614601565b91505092915050565b600067ffffffffffffffff82169050919050565b61466081614643565b811461466b57600080fd5b50565b60008135905061467d81614657565b92915050565b600060208284031215614699576146986145e0565b5b60006146a78482850161466e565b91505092915050565b6146b981614643565b82525050565b60006020820190506146d460008301846146b0565b92915050565b60008115159050919050565b6146ef816146da565b82525050565b600060208201905061470a60008301846146e6565b92915050565b6000819050919050565b61472381614710565b811461472e57600080fd5b50565b6000813590506147408161471a565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61479982614750565b810181811067ffffffffffffffff821117156147b8576147b7614761565b5b80604052505050565b60006147cb6145d6565b90506147d78282614790565b919050565b600067ffffffffffffffff8211156147f7576147f6614761565b5b61480082614750565b9050602081019050919050565b82818337600083830152505050565b600061482f61482a846147dc565b6147c1565b90508281526020810184848401111561484b5761484a61474b565b5b61485684828561480d565b509392505050565b600082601f83011261487357614872614746565b5b813561488384826020860161481c565b91505092915050565b600080600080608085870312156148a6576148a56145e0565b5b60006148b487828801614601565b94505060206148c587828801614601565b93505060406148d687828801614731565b925050606085013567ffffffffffffffff8111156148f7576148f66145e5565b5b6149038782880161485e565b91505092959194509250565b600080600080600060a0868803121561492b5761492a6145e0565b5b600061493988828901614601565b955050602061494a88828901614601565b945050604061495b88828901614731565b935050606061496c88828901614731565b925050608086013567ffffffffffffffff81111561498d5761498c6145e5565b5b6149998882890161485e565b9150509295509295909350565b6000806000606084860312156149bf576149be6145e0565b5b60006149cd86828701614601565b93505060206149de86828701614731565b92505060406149ef8682870161466e565b9150509250925092565b600060208284031215614a0f57614a0e6145e0565b5b6000614a1d84828501614731565b91505092915050565b6000819050919050565b614a3981614a26565b8114614a4457600080fd5b50565b600081359050614a5681614a30565b92915050565b600080600080600080600080610100898b031215614a7d57614a7c6145e0565b5b6000614a8b8b828c01614a47565b9850506020614a9c8b828c01614601565b9750506040614aad8b828c01614601565b9650506060614abe8b828c01614601565b9550506080614acf8b828c01614731565b94505060a0614ae08b828c0161466e565b93505060c0614af18b828c0161466e565b92505060e089013567ffffffffffffffff811115614b1257614b116145e5565b5b614b1e8b828c0161485e565b9150509295985092959890939650565b60008060408385031215614b4557614b446145e0565b5b6000614b5385828601614601565b9250506020614b6485828601614601565b9150509250929050565b614b7781614710565b82525050565b6000602082019050614b926000830184614b6e565b92915050565b60008060008060808587031215614bb257614bb16145e0565b5b6000614bc087828801614601565b9450506020614bd187828801614601565b9350506040614be287828801614731565b9250506060614bf387828801614731565b91505092959194509250565b600060ff82169050919050565b614c1581614bff565b8114614c2057600080fd5b50565b600081359050614c3281614c0c565b92915050565b600060208284031215614c4e57614c4d6145e0565b5b6000614c5c84828501614c23565b91505092915050565b614c6e81614bff565b82525050565b6000602082019050614c896000830184614c65565b92915050565b600060208284031215614ca557614ca46145e0565b5b6000614cb384828501614a47565b91505092915050565b600067ffffffffffffffff821115614cd757614cd6614761565b5b614ce082614750565b9050602081019050919050565b6000614d00614cfb84614cbc565b6147c1565b905082815260208101848484011115614d1c57614d1b61474b565b5b614d2784828561480d565b509392505050565b600082601f830112614d4457614d43614746565b5b8135614d54848260208601614ced565b91505092915050565b60008060008060008060008060006101208a8c031215614d8057614d7f6145e0565b5b6000614d8e8c828d01614a47565b9950506020614d9f8c828d01614601565b9850506040614db08c828d01614601565b9750506060614dc18c828d01614601565b9650506080614dd28c828d01614731565b95505060a0614de38c828d0161466e565b94505060c0614df48c828d0161466e565b93505060e08a013567ffffffffffffffff811115614e1557614e146145e5565b5b614e218c828d01614d2f565b9250506101008a013567ffffffffffffffff811115614e4357614e426145e5565b5b614e4f8c828d0161485e565b9150509295985092959850929598565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b614e948161459a565b82525050565b6000614ea68383614e8b565b60208301905092915050565b6000602082019050919050565b6000614eca82614e5f565b614ed48185614e6a565b9350614edf83614e7b565b8060005b83811015614f10578151614ef78882614e9a565b9750614f0283614eb2565b925050600181019050614ee3565b5085935050505092915050565b60006020820190508181036000830152614f378184614ebf565b905092915050565b6000614f4a8261457a565b9050919050565b614f5a81614f3f565b82525050565b6000602082019050614f756000830184614f51565b92915050565b600080600080600060a08688031215614f9757614f966145e0565b5b6000614fa588828901614601565b9550506020614fb688828901614601565b9450506040614fc788828901614731565b9350506060614fd888828901614731565b9250506080614fe988828901614731565b9150509295509295909350565b614fff816146da565b811461500a57600080fd5b50565b60008135905061501c81614ff6565b92915050565b600060208284031215615038576150376145e0565b5b60006150468482850161500d565b91505092915050565b60008060008060808587031215615069576150686145e0565b5b600061507787828801614601565b945050602061508887828801614731565b935050604061509987828801614601565b925050606085013567ffffffffffffffff8111156150ba576150b96145e5565b5b6150c68782880161485e565b91505092959194509250565b600381106150df57600080fd5b50565b6000813590506150f1816150d2565b92915050565b6000806040838503121561510e5761510d6145e0565b5b600061511c858286016150e2565b925050602061512d85828601614c23565b9150509250929050565b61514081614f3f565b811461514b57600080fd5b50565b60008135905061515d81615137565b92915050565b600060208284031215615179576151786145e0565b5b60006151878482850161514e565b91505092915050565b600082825260208201905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006151d7602083615190565b91506151e2826151a1565b602082019050919050565b60006020820190508181036000830152615206816151ca565b9050919050565b7f6e6f7420616c6c6f77656420746f6b656e000000000000000000000000000000600082015250565b6000615243601183615190565b915061524e8261520d565b602082019050919050565b6000602082019050818103600083015261527281615236565b9050919050565b7f6c6f636b656420746f6b656e0000000000000000000000000000000000000000600082015250565b60006152af600c83615190565b91506152ba82615279565b602082019050919050565b600060208201905081810360008301526152de816152a2565b9050919050565b7f6e6f7420737570706f7274000000000000000000000000000000000000000000600082015250565b600061531b600b83615190565b9150615326826152e5565b602082019050919050565b6000602082019050818103600083015261534a8161530e565b9050919050565b7f6d73672e73656e646572206973206e6f7420616e206f70657261746f72000000600082015250565b6000615387601d83615190565b915061539282615351565b602082019050919050565b600060208201905081810360008301526153b68161537a565b9050919050565b7f6d6178206f70657261746f72206c696d69740000000000000000000000000000600082015250565b60006153f3601283615190565b91506153fe826153bd565b602082019050919050565b60006020820190508181036000830152615422816153e6565b9050919050565b7f6578697374206f70657261746f72000000000000000000000000000000000000600082015250565b600061545f600e83615190565b915061546a82615429565b602082019050919050565b6000602082019050818103600083015261548e81615452565b9050919050565b7f6578747261446174612073697a65206572726f72000000000000000000000000600082015250565b60006154cb601483615190565b91506154d682615495565b602082019050919050565b600060208201905081810360008301526154fa816154be565b9050919050565b7f63616c6c6261636b2061646472657373206572726f7200000000000000000000600082015250565b6000615537601683615190565b915061554282615501565b602082019050919050565b600060208201905081810360008301526155668161552a565b9050919050565b60008151905061557c8161471a565b92915050565b600060208284031215615598576155976145e0565b5b60006155a68482850161556d565b91505092915050565b7f6f666665725072696365206572726f7200000000000000000000000000000000600082015250565b60006155e5601083615190565b91506155f0826155af565b602082019050919050565b60006020820190508181036000830152615614816155d8565b9050919050565b600060808201905061563060008301876145ac565b61563d6020830186614b6e565b61564a60408301856145ac565b6156576060830184614b6e565b95945050505050565b7f616c6c6f77656420746f6b656e00000000000000000000000000000000000000600082015250565b6000615696600d83615190565b91506156a182615660565b602082019050919050565b600060208201905081810360008301526156c581615689565b9050919050565b7f756e6c6f636b656420746f6b656e000000000000000000000000000000000000600082015250565b6000615702600e83615190565b915061570d826156cc565b602082019050919050565b60006020820190508181036000830152615731816156f5565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061577282614710565b915061577d83614710565b9250828210156157905761578f615738565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600061580482614710565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361583657615835615738565b5b600182019050919050565b7f7a65726f207468726573686f6c64000000000000000000000000000000000000600082015250565b6000615877600e83615190565b915061588282615841565b602082019050919050565b600060208201905081810360008301526158a68161586a565b9050919050565b7f626967676572207468616e206e756d206f66206f70657261746f727300000000600082015250565b60006158e3601c83615190565b91506158ee826158ad565b602082019050919050565b60006020820190508181036000830152615912816158d6565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006159a4602683615190565b91506159af82615948565b604082019050919050565b600060208201905081810360008301526159d381615997565b9050919050565b7f6e6f6e6365206d69736d61746368000000000000000000000000000000000000600082015250565b6000615a10600e83615190565b9150615a1b826159da565b602082019050919050565b60006020820190508181036000830152615a3f81615a03565b9050919050565b600081905092915050565b6000615a5d8385615a46565b9350615a6a83858461480d565b82840190509392505050565b6000615a83828486615a51565b91508190509392505050565b6000615a9a82614643565b915067ffffffffffffffff8203615ab457615ab3615738565b5b600182019050919050565b615ac881614a26565b82525050565b60038110615adf57615ade615919565b5b50565b6000819050615af082615ace565b919050565b6000615b0082615ae2565b9050919050565b615b1081615af5565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015615b50578082015181840152602081019050615b35565b83811115615b5f576000848401525b50505050565b6000615b7082615b16565b615b7a8185615b21565b9350615b8a818560208601615b32565b615b9381614750565b840191505092915050565b600060c082019050615bb36000830189615abf565b615bc06020830188615b07565b615bcd6040830187614b6e565b615bda60608301866146b0565b615be760808301856146b0565b81810360a0830152615bf98184615b65565b9050979650505050505050565b6000604082019050615c1b60008301856145ac565b615c286020830184614b6e565b9392505050565b600081519050615c3e81614ff6565b92915050565b600060208284031215615c5a57615c596145e0565b5b6000615c6884828501615c2f565b91505092915050565b7f68616e646c6545524332305472616e736665723a207472616e7366657220666160008201527f696c656400000000000000000000000000000000000000000000000000000000602082015250565b6000615ccd602483615190565b9150615cd882615c71565b604082019050919050565b60006020820190508181036000830152615cfc81615cc0565b9050919050565b6000606082019050615d1860008301866145ac565b615d2560208301856145ac565b615d326040830184614b6e565b949350505050565b600081519050919050565b6000615d5082615d3a565b615d5a8185615190565b9350615d6a818560208601615b32565b615d7381614750565b840191505092915050565b6000606082019050615d9360008301866145ac565b615da06020830185614b6e565b8181036040830152615db28184615d45565b9050949350505050565b6000615dc782614710565b9150615dd283614710565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115615e0757615e06615738565b5b828201905092915050565b7f7265717565737445524332305472616e736665723a207472616e73666572467260008201527f6f6d206661696c65640000000000000000000000000000000000000000000000602082015250565b6000615e6e602983615190565b9150615e7982615e12565b604082019050919050565b60006020820190508181036000830152615e9d81615e61565b9050919050565b7f73746f7070656420627269646765000000000000000000000000000000000000600082015250565b6000615eda600e83615190565b9150615ee582615ea4565b602082019050919050565b60006020820190508181036000830152615f0981615ecd565b9050919050565b6000615f23615f1e84614cbc565b6147c1565b905082815260208101848484011115615f3f57615f3e61474b565b5b615f4a848285615b32565b509392505050565b600082601f830112615f6757615f66614746565b5b8151615f77848260208601615f10565b91505092915050565b600060208284031215615f9657615f956145e0565b5b600082015167ffffffffffffffff811115615fb457615fb36145e5565b5b615fc084828501615f52565b91505092915050565b60006020820190508181036000830152615fe38184615d45565b905092915050565b6000819050919050565b6000819050919050565b600061601a61601561601084615feb565b615ff5565b614710565b9050919050565b61602a81615fff565b82525050565b6000819050919050565b600061605561605061604b84616030565b615ff5565b614bff565b9050919050565b6160658161603a565b82525050565b600060e082019050616080600083018a615b07565b61608d6020830189614b6e565b61609a60408301886146b0565b6160a76060830187616021565b81810360808301526160b98186615b65565b90506160c860a083018561605c565b81810360c08301526160da8184615b65565b905098975050505050505050565b7f7a65726f206d73672e76616c7565000000000000000000000000000000000000600082015250565b600061611e600e83615190565b9150616129826160e8565b602082019050919050565b6000602082019050818103600083015261614d81616111565b9050919050565b600060a0820190506161696000830188615b07565b6161766020830187614b6e565b61618360408301866146b0565b6161906060830185614b6e565b81810360808301526161a28184615b65565b90509695505050505050565b60006161b982614bff565b9150600082036161cc576161cb615738565b5b600182039050919050565b60006161e282614bff565b915060ff82036161f5576161f4615738565b5b600182019050919050565b7f72656d6f76656420766f74650000000000000000000000000000000000000000600082015250565b6000616236600c83615190565b915061624182616200565b602082019050919050565b6000602082019050818103600083015261626581616229565b9050919050565b7f636c6f73656420766f7465000000000000000000000000000000000000000000600082015250565b60006162a2600b83615190565b91506162ad8261626c565b602082019050919050565b600060208201905081810360008301526162d181616295565b9050919050565b60006162e382614643565b91506162ee83614643565b92508267ffffffffffffffff0382111561630b5761630a615738565b5b828201905092915050565b7f696e73756666696369656e74206665654c696d69740000000000000000000000600082015250565b600061634c601583615190565b915061635782616316565b602082019050919050565b6000602082019050818103600083015261637b8161633f565b9050919050565b600061639d6163986163938461457a565b615ff5565b61457a565b9050919050565b60006163af82616382565b9050919050565b60006163c1826163a4565b9050919050565b6163d1816163b6565b82525050565b60006040820190506163ec60008301856163c8565b6163f96020830184614b6e565b9392505050565b7f5f7061794552433230466565416e64526566756e644368616e67653a2074726160008201527f6e73666572206661696c65640000000000000000000000000000000000000000602082015250565b600061645c602c83615190565b915061646782616400565b604082019050919050565b6000602082019050818103600083015261648b8161644f565b905091905056fea2646970667358221220ffca0e07cb23acd039153556caaba3c422fdd798f1c26a3def99e316966bc16364736f6c634300080e0033"

// DeployExtbridge deploys a new Klaytn contract, binding an instance of Extbridge to it.
func DeployExtbridge(auth *bind.TransactOpts, backend bind.ContractBackend, _modeMintBurn bool) (common.Address, *types.Transaction, *Extbridge, error) {
	parsed, err := abi.JSON(strings.NewReader(ExtbridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExtbridgeBin), backend, _modeMintBurn)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Extbridge{ExtbridgeCaller: ExtbridgeCaller{contract: contract}, ExtbridgeTransactor: ExtbridgeTransactor{contract: contract}, ExtbridgeFilterer: ExtbridgeFilterer{contract: contract}}, nil
}

// Extbridge is an auto generated Go binding around a Klaytn contract.
type Extbridge struct {
	ExtbridgeCaller     // Read-only binding to the contract
	ExtbridgeTransactor // Write-only binding to the contract
	ExtbridgeFilterer   // Log filterer for contract events
}

// ExtbridgeCaller is an auto generated read-only Go binding around a Klaytn contract.
type ExtbridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtbridgeTransactor is an auto generated write-only Go binding around a Klaytn contract.
type ExtbridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtbridgeFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type ExtbridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtbridgeSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type ExtbridgeSession struct {
	Contract     *Extbridge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExtbridgeCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type ExtbridgeCallerSession struct {
	Contract *ExtbridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ExtbridgeTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type ExtbridgeTransactorSession struct {
	Contract     *ExtbridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ExtbridgeRaw is an auto generated low-level Go binding around a Klaytn contract.
type ExtbridgeRaw struct {
	Contract *Extbridge // Generic contract binding to access the raw methods on
}

// ExtbridgeCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type ExtbridgeCallerRaw struct {
	Contract *ExtbridgeCaller // Generic read-only contract binding to access the raw methods on
}

// ExtbridgeTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type ExtbridgeTransactorRaw struct {
	Contract *ExtbridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtbridge creates a new instance of Extbridge, bound to a specific deployed contract.
func NewExtbridge(address common.Address, backend bind.ContractBackend) (*Extbridge, error) {
	contract, err := bindExtbridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Extbridge{ExtbridgeCaller: ExtbridgeCaller{contract: contract}, ExtbridgeTransactor: ExtbridgeTransactor{contract: contract}, ExtbridgeFilterer: ExtbridgeFilterer{contract: contract}}, nil
}

// NewExtbridgeCaller creates a new read-only instance of Extbridge, bound to a specific deployed contract.
func NewExtbridgeCaller(address common.Address, caller bind.ContractCaller) (*ExtbridgeCaller, error) {
	contract, err := bindExtbridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtbridgeCaller{contract: contract}, nil
}

// NewExtbridgeTransactor creates a new write-only instance of Extbridge, bound to a specific deployed contract.
func NewExtbridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtbridgeTransactor, error) {
	contract, err := bindExtbridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtbridgeTransactor{contract: contract}, nil
}

// NewExtbridgeFilterer creates a new log filterer instance of Extbridge, bound to a specific deployed contract.
func NewExtbridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtbridgeFilterer, error) {
	contract, err := bindExtbridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtbridgeFilterer{contract: contract}, nil
}

// bindExtbridge binds a generic wrapper to an already deployed contract.
func bindExtbridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExtbridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Extbridge *ExtbridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Extbridge.Contract.ExtbridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Extbridge *ExtbridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extbridge.Contract.ExtbridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Extbridge *ExtbridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Extbridge.Contract.ExtbridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Extbridge *ExtbridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Extbridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Extbridge *ExtbridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extbridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Extbridge *ExtbridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Extbridge.Contract.contract.Transact(opts, method, params...)
}

// MAXOPERATOR is a free data retrieval call binding the contract method 0x3a3099d1.
//
// Solidity: function MAX_OPERATOR() view returns(uint64)
func (_Extbridge *ExtbridgeCaller) MAXOPERATOR(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "MAX_OPERATOR")
	return *ret0, err
}

// MAXOPERATOR is a free data retrieval call binding the contract method 0x3a3099d1.
//
// Solidity: function MAX_OPERATOR() view returns(uint64)
func (_Extbridge *ExtbridgeSession) MAXOPERATOR() (uint64, error) {
	return _Extbridge.Contract.MAXOPERATOR(&_Extbridge.CallOpts)
}

// MAXOPERATOR is a free data retrieval call binding the contract method 0x3a3099d1.
//
// Solidity: function MAX_OPERATOR() view returns(uint64)
func (_Extbridge *ExtbridgeCallerSession) MAXOPERATOR() (uint64, error) {
	return _Extbridge.Contract.MAXOPERATOR(&_Extbridge.CallOpts)
}

// Callback is a free data retrieval call binding the contract method 0x083b2732.
//
// Solidity: function callback() view returns(address)
func (_Extbridge *ExtbridgeCaller) Callback(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "callback")
	return *ret0, err
}

// Callback is a free data retrieval call binding the contract method 0x083b2732.
//
// Solidity: function callback() view returns(address)
func (_Extbridge *ExtbridgeSession) Callback() (common.Address, error) {
	return _Extbridge.Contract.Callback(&_Extbridge.CallOpts)
}

// Callback is a free data retrieval call binding the contract method 0x083b2732.
//
// Solidity: function callback() view returns(address)
func (_Extbridge *ExtbridgeCallerSession) Callback() (common.Address, error) {
	return _Extbridge.Contract.Callback(&_Extbridge.CallOpts)
}

// ClosedValueTransferVotes is a free data retrieval call binding the contract method 0x9832c1d7.
//
// Solidity: function closedValueTransferVotes(uint64 ) view returns(bool)
func (_Extbridge *ExtbridgeCaller) ClosedValueTransferVotes(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "closedValueTransferVotes", arg0)
	return *ret0, err
}

// ClosedValueTransferVotes is a free data retrieval call binding the contract method 0x9832c1d7.
//
// Solidity: function closedValueTransferVotes(uint64 ) view returns(bool)
func (_Extbridge *ExtbridgeSession) ClosedValueTransferVotes(arg0 uint64) (bool, error) {
	return _Extbridge.Contract.ClosedValueTransferVotes(&_Extbridge.CallOpts, arg0)
}

// ClosedValueTransferVotes is a free data retrieval call binding the contract method 0x9832c1d7.
//
// Solidity: function closedValueTransferVotes(uint64 ) view returns(bool)
func (_Extbridge *ExtbridgeCallerSession) ClosedValueTransferVotes(arg0 uint64) (bool, error) {
	return _Extbridge.Contract.ClosedValueTransferVotes(&_Extbridge.CallOpts, arg0)
}

// ConfigurationNonce is a free data retrieval call binding the contract method 0xac6fff0b.
//
// Solidity: function configurationNonce() view returns(uint64)
func (_Extbridge *ExtbridgeCaller) ConfigurationNonce(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "configurationNonce")
	return *ret0, err
}

// ConfigurationNonce is a free data retrieval call binding the contract method 0xac6fff0b.
//
// Solidity: function configurationNonce() view returns(uint64)
func (_Extbridge *ExtbridgeSession) ConfigurationNonce() (uint64, error) {
	return _Extbridge.Contract.ConfigurationNonce(&_Extbridge.CallOpts)
}

// ConfigurationNonce is a free data retrieval call binding the contract method 0xac6fff0b.
//
// Solidity: function configurationNonce() view returns(uint64)
func (_Extbridge *ExtbridgeCallerSession) ConfigurationNonce() (uint64, error) {
	return _Extbridge.Contract.ConfigurationNonce(&_Extbridge.CallOpts)
}

// FeeOfERC20 is a free data retrieval call binding the contract method 0x488af871.
//
// Solidity: function feeOfERC20(address ) view returns(uint256)
func (_Extbridge *ExtbridgeCaller) FeeOfERC20(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "feeOfERC20", arg0)
	return *ret0, err
}

// FeeOfERC20 is a free data retrieval call binding the contract method 0x488af871.
//
// Solidity: function feeOfERC20(address ) view returns(uint256)
func (_Extbridge *ExtbridgeSession) FeeOfERC20(arg0 common.Address) (*big.Int, error) {
	return _Extbridge.Contract.FeeOfERC20(&_Extbridge.CallOpts, arg0)
}

// FeeOfERC20 is a free data retrieval call binding the contract method 0x488af871.
//
// Solidity: function feeOfERC20(address ) view returns(uint256)
func (_Extbridge *ExtbridgeCallerSession) FeeOfERC20(arg0 common.Address) (*big.Int, error) {
	return _Extbridge.Contract.FeeOfERC20(&_Extbridge.CallOpts, arg0)
}

// FeeOfKLAY is a free data retrieval call binding the contract method 0xc263b5d6.
//
// Solidity: function feeOfKLAY() view returns(uint256)
func (_Extbridge *ExtbridgeCaller) FeeOfKLAY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "feeOfKLAY")
	return *ret0, err
}

// FeeOfKLAY is a free data retrieval call binding the contract method 0xc263b5d6.
//
// Solidity: function feeOfKLAY() view returns(uint256)
func (_Extbridge *ExtbridgeSession) FeeOfKLAY() (*big.Int, error) {
	return _Extbridge.Contract.FeeOfKLAY(&_Extbridge.CallOpts)
}

// FeeOfKLAY is a free data retrieval call binding the contract method 0xc263b5d6.
//
// Solidity: function feeOfKLAY() view returns(uint256)
func (_Extbridge *ExtbridgeCallerSession) FeeOfKLAY() (*big.Int, error) {
	return _Extbridge.Contract.FeeOfKLAY(&_Extbridge.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Extbridge *ExtbridgeCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "feeReceiver")
	return *ret0, err
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Extbridge *ExtbridgeSession) FeeReceiver() (common.Address, error) {
	return _Extbridge.Contract.FeeReceiver(&_Extbridge.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Extbridge *ExtbridgeCallerSession) FeeReceiver() (common.Address, error) {
	return _Extbridge.Contract.FeeReceiver(&_Extbridge.CallOpts)
}

// GetOperatorList is a free data retrieval call binding the contract method 0xb2c01030.
//
// Solidity: function getOperatorList() view returns(address[])
func (_Extbridge *ExtbridgeCaller) GetOperatorList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "getOperatorList")
	return *ret0, err
}

// GetOperatorList is a free data retrieval call binding the contract method 0xb2c01030.
//
// Solidity: function getOperatorList() view returns(address[])
func (_Extbridge *ExtbridgeSession) GetOperatorList() ([]common.Address, error) {
	return _Extbridge.Contract.GetOperatorList(&_Extbridge.CallOpts)
}

// GetOperatorList is a free data retrieval call binding the contract method 0xb2c01030.
//
// Solidity: function getOperatorList() view returns(address[])
func (_Extbridge *ExtbridgeCallerSession) GetOperatorList() ([]common.Address, error) {
	return _Extbridge.Contract.GetOperatorList(&_Extbridge.CallOpts)
}

// GetRegisteredTokenList is a free data retrieval call binding the contract method 0xea21eade.
//
// Solidity: function getRegisteredTokenList() view returns(address[])
func (_Extbridge *ExtbridgeCaller) GetRegisteredTokenList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "getRegisteredTokenList")
	return *ret0, err
}

// GetRegisteredTokenList is a free data retrieval call binding the contract method 0xea21eade.
//
// Solidity: function getRegisteredTokenList() view returns(address[])
func (_Extbridge *ExtbridgeSession) GetRegisteredTokenList() ([]common.Address, error) {
	return _Extbridge.Contract.GetRegisteredTokenList(&_Extbridge.CallOpts)
}

// GetRegisteredTokenList is a free data retrieval call binding the contract method 0xea21eade.
//
// Solidity: function getRegisteredTokenList() view returns(address[])
func (_Extbridge *ExtbridgeCallerSession) GetRegisteredTokenList() ([]common.Address, error) {
	return _Extbridge.Contract.GetRegisteredTokenList(&_Extbridge.CallOpts)
}

// HandleNoncesToBlockNums is a free data retrieval call binding the contract method 0x13a6738a.
//
// Solidity: function handleNoncesToBlockNums(uint64 ) view returns(uint64)
func (_Extbridge *ExtbridgeCaller) HandleNoncesToBlockNums(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "handleNoncesToBlockNums", arg0)
	return *ret0, err
}

// HandleNoncesToBlockNums is a free data retrieval call binding the contract method 0x13a6738a.
//
// Solidity: function handleNoncesToBlockNums(uint64 ) view returns(uint64)
func (_Extbridge *ExtbridgeSession) HandleNoncesToBlockNums(arg0 uint64) (uint64, error) {
	return _Extbridge.Contract.HandleNoncesToBlockNums(&_Extbridge.CallOpts, arg0)
}

// HandleNoncesToBlockNums is a free data retrieval call binding the contract method 0x13a6738a.
//
// Solidity: function handleNoncesToBlockNums(uint64 ) view returns(uint64)
func (_Extbridge *ExtbridgeCallerSession) HandleNoncesToBlockNums(arg0 uint64) (uint64, error) {
	return _Extbridge.Contract.HandleNoncesToBlockNums(&_Extbridge.CallOpts, arg0)
}

// HandledRequestTx is a free data retrieval call binding the contract method 0x8a75eee2.
//
// Solidity: function handledRequestTx(bytes32 ) view returns(bool)
func (_Extbridge *ExtbridgeCaller) HandledRequestTx(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "handledRequestTx", arg0)
	return *ret0, err
}

// HandledRequestTx is a free data retrieval call binding the contract method 0x8a75eee2.
//
// Solidity: function handledRequestTx(bytes32 ) view returns(bool)
func (_Extbridge *ExtbridgeSession) HandledRequestTx(arg0 [32]byte) (bool, error) {
	return _Extbridge.Contract.HandledRequestTx(&_Extbridge.CallOpts, arg0)
}

// HandledRequestTx is a free data retrieval call binding the contract method 0x8a75eee2.
//
// Solidity: function handledRequestTx(bytes32 ) view returns(bool)
func (_Extbridge *ExtbridgeCallerSession) HandledRequestTx(arg0 [32]byte) (bool, error) {
	return _Extbridge.Contract.HandledRequestTx(&_Extbridge.CallOpts, arg0)
}

// IndexOfTokens is a free data retrieval call binding the contract method 0x48a18a6a.
//
// Solidity: function indexOfTokens(address ) view returns(uint256)
func (_Extbridge *ExtbridgeCaller) IndexOfTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "indexOfTokens", arg0)
	return *ret0, err
}

// IndexOfTokens is a free data retrieval call binding the contract method 0x48a18a6a.
//
// Solidity: function indexOfTokens(address ) view returns(uint256)
func (_Extbridge *ExtbridgeSession) IndexOfTokens(arg0 common.Address) (*big.Int, error) {
	return _Extbridge.Contract.IndexOfTokens(&_Extbridge.CallOpts, arg0)
}

// IndexOfTokens is a free data retrieval call binding the contract method 0x48a18a6a.
//
// Solidity: function indexOfTokens(address ) view returns(uint256)
func (_Extbridge *ExtbridgeCallerSession) IndexOfTokens(arg0 common.Address) (*big.Int, error) {
	return _Extbridge.Contract.IndexOfTokens(&_Extbridge.CallOpts, arg0)
}

// IsRunning is a free data retrieval call binding the contract method 0x2014e5d1.
//
// Solidity: function isRunning() view returns(bool)
func (_Extbridge *ExtbridgeCaller) IsRunning(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "isRunning")
	return *ret0, err
}

// IsRunning is a free data retrieval call binding the contract method 0x2014e5d1.
//
// Solidity: function isRunning() view returns(bool)
func (_Extbridge *ExtbridgeSession) IsRunning() (bool, error) {
	return _Extbridge.Contract.IsRunning(&_Extbridge.CallOpts)
}

// IsRunning is a free data retrieval call binding the contract method 0x2014e5d1.
//
// Solidity: function isRunning() view returns(bool)
func (_Extbridge *ExtbridgeCallerSession) IsRunning() (bool, error) {
	return _Extbridge.Contract.IsRunning(&_Extbridge.CallOpts)
}

// LockedTokens is a free data retrieval call binding the contract method 0x5eb7413a.
//
// Solidity: function lockedTokens(address ) view returns(bool)
func (_Extbridge *ExtbridgeCaller) LockedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "lockedTokens", arg0)
	return *ret0, err
}

// LockedTokens is a free data retrieval call binding the contract method 0x5eb7413a.
//
// Solidity: function lockedTokens(address ) view returns(bool)
func (_Extbridge *ExtbridgeSession) LockedTokens(arg0 common.Address) (bool, error) {
	return _Extbridge.Contract.LockedTokens(&_Extbridge.CallOpts, arg0)
}

// LockedTokens is a free data retrieval call binding the contract method 0x5eb7413a.
//
// Solidity: function lockedTokens(address ) view returns(bool)
func (_Extbridge *ExtbridgeCallerSession) LockedTokens(arg0 common.Address) (bool, error) {
	return _Extbridge.Contract.LockedTokens(&_Extbridge.CallOpts, arg0)
}

// LowerHandleNonce is a free data retrieval call binding the contract method 0x4b40b826.
//
// Solidity: function lowerHandleNonce() view returns(uint64)
func (_Extbridge *ExtbridgeCaller) LowerHandleNonce(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "lowerHandleNonce")
	return *ret0, err
}

// LowerHandleNonce is a free data retrieval call binding the contract method 0x4b40b826.
//
// Solidity: function lowerHandleNonce() view returns(uint64)
func (_Extbridge *ExtbridgeSession) LowerHandleNonce() (uint64, error) {
	return _Extbridge.Contract.LowerHandleNonce(&_Extbridge.CallOpts)
}

// LowerHandleNonce is a free data retrieval call binding the contract method 0x4b40b826.
//
// Solidity: function lowerHandleNonce() view returns(uint64)
func (_Extbridge *ExtbridgeCallerSession) LowerHandleNonce() (uint64, error) {
	return _Extbridge.Contract.LowerHandleNonce(&_Extbridge.CallOpts)
}

// ModeMintBurn is a free data retrieval call binding the contract method 0x6e176ec2.
//
// Solidity: function modeMintBurn() view returns(bool)
func (_Extbridge *ExtbridgeCaller) ModeMintBurn(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "modeMintBurn")
	return *ret0, err
}

// ModeMintBurn is a free data retrieval call binding the contract method 0x6e176ec2.
//
// Solidity: function modeMintBurn() view returns(bool)
func (_Extbridge *ExtbridgeSession) ModeMintBurn() (bool, error) {
	return _Extbridge.Contract.ModeMintBurn(&_Extbridge.CallOpts)
}

// ModeMintBurn is a free data retrieval call binding the contract method 0x6e176ec2.
//
// Solidity: function modeMintBurn() view returns(bool)
func (_Extbridge *ExtbridgeCallerSession) ModeMintBurn() (bool, error) {
	return _Extbridge.Contract.ModeMintBurn(&_Extbridge.CallOpts)
}

// OperatorList is a free data retrieval call binding the contract method 0xcb38f407.
//
// Solidity: function operatorList(uint256 ) view returns(address)
func (_Extbridge *ExtbridgeCaller) OperatorList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "operatorList", arg0)
	return *ret0, err
}

// OperatorList is a free data retrieval call binding the contract method 0xcb38f407.
//
// Solidity: function operatorList(uint256 ) view returns(address)
func (_Extbridge *ExtbridgeSession) OperatorList(arg0 *big.Int) (common.Address, error) {
	return _Extbridge.Contract.OperatorList(&_Extbridge.CallOpts, arg0)
}

// OperatorList is a free data retrieval call binding the contract method 0xcb38f407.
//
// Solidity: function operatorList(uint256 ) view returns(address)
func (_Extbridge *ExtbridgeCallerSession) OperatorList(arg0 *big.Int) (common.Address, error) {
	return _Extbridge.Contract.OperatorList(&_Extbridge.CallOpts, arg0)
}

// OperatorThresholds is a free data retrieval call binding the contract method 0x5526f76b.
//
// Solidity: function operatorThresholds(uint8 ) view returns(uint8)
func (_Extbridge *ExtbridgeCaller) OperatorThresholds(opts *bind.CallOpts, arg0 uint8) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "operatorThresholds", arg0)
	return *ret0, err
}

// OperatorThresholds is a free data retrieval call binding the contract method 0x5526f76b.
//
// Solidity: function operatorThresholds(uint8 ) view returns(uint8)
func (_Extbridge *ExtbridgeSession) OperatorThresholds(arg0 uint8) (uint8, error) {
	return _Extbridge.Contract.OperatorThresholds(&_Extbridge.CallOpts, arg0)
}

// OperatorThresholds is a free data retrieval call binding the contract method 0x5526f76b.
//
// Solidity: function operatorThresholds(uint8 ) view returns(uint8)
func (_Extbridge *ExtbridgeCallerSession) OperatorThresholds(arg0 uint8) (uint8, error) {
	return _Extbridge.Contract.OperatorThresholds(&_Extbridge.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Extbridge *ExtbridgeCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "operators", arg0)
	return *ret0, err
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Extbridge *ExtbridgeSession) Operators(arg0 common.Address) (bool, error) {
	return _Extbridge.Contract.Operators(&_Extbridge.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Extbridge *ExtbridgeCallerSession) Operators(arg0 common.Address) (bool, error) {
	return _Extbridge.Contract.Operators(&_Extbridge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Extbridge *ExtbridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Extbridge *ExtbridgeSession) Owner() (common.Address, error) {
	return _Extbridge.Contract.Owner(&_Extbridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Extbridge *ExtbridgeCallerSession) Owner() (common.Address, error) {
	return _Extbridge.Contract.Owner(&_Extbridge.CallOpts)
}

// RecoveryBlockNumber is a free data retrieval call binding the contract method 0x989ba0d3.
//
// Solidity: function recoveryBlockNumber() view returns(uint64)
func (_Extbridge *ExtbridgeCaller) RecoveryBlockNumber(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "recoveryBlockNumber")
	return *ret0, err
}

// RecoveryBlockNumber is a free data retrieval call binding the contract method 0x989ba0d3.
//
// Solidity: function recoveryBlockNumber() view returns(uint64)
func (_Extbridge *ExtbridgeSession) RecoveryBlockNumber() (uint64, error) {
	return _Extbridge.Contract.RecoveryBlockNumber(&_Extbridge.CallOpts)
}

// RecoveryBlockNumber is a free data retrieval call binding the contract method 0x989ba0d3.
//
// Solidity: function recoveryBlockNumber() view returns(uint64)
func (_Extbridge *ExtbridgeCallerSession) RecoveryBlockNumber() (uint64, error) {
	return _Extbridge.Contract.RecoveryBlockNumber(&_Extbridge.CallOpts)
}

// RegisteredTokenList is a free data retrieval call binding the contract method 0x3e4fe949.
//
// Solidity: function registeredTokenList(uint256 ) view returns(address)
func (_Extbridge *ExtbridgeCaller) RegisteredTokenList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "registeredTokenList", arg0)
	return *ret0, err
}

// RegisteredTokenList is a free data retrieval call binding the contract method 0x3e4fe949.
//
// Solidity: function registeredTokenList(uint256 ) view returns(address)
func (_Extbridge *ExtbridgeSession) RegisteredTokenList(arg0 *big.Int) (common.Address, error) {
	return _Extbridge.Contract.RegisteredTokenList(&_Extbridge.CallOpts, arg0)
}

// RegisteredTokenList is a free data retrieval call binding the contract method 0x3e4fe949.
//
// Solidity: function registeredTokenList(uint256 ) view returns(address)
func (_Extbridge *ExtbridgeCallerSession) RegisteredTokenList(arg0 *big.Int) (common.Address, error) {
	return _Extbridge.Contract.RegisteredTokenList(&_Extbridge.CallOpts, arg0)
}

// RegisteredTokens is a free data retrieval call binding the contract method 0x8c0bd916.
//
// Solidity: function registeredTokens(address ) view returns(address)
func (_Extbridge *ExtbridgeCaller) RegisteredTokens(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "registeredTokens", arg0)
	return *ret0, err
}

// RegisteredTokens is a free data retrieval call binding the contract method 0x8c0bd916.
//
// Solidity: function registeredTokens(address ) view returns(address)
func (_Extbridge *ExtbridgeSession) RegisteredTokens(arg0 common.Address) (common.Address, error) {
	return _Extbridge.Contract.RegisteredTokens(&_Extbridge.CallOpts, arg0)
}

// RegisteredTokens is a free data retrieval call binding the contract method 0x8c0bd916.
//
// Solidity: function registeredTokens(address ) view returns(address)
func (_Extbridge *ExtbridgeCallerSession) RegisteredTokens(arg0 common.Address) (common.Address, error) {
	return _Extbridge.Contract.RegisteredTokens(&_Extbridge.CallOpts, arg0)
}

// RequestNonce is a free data retrieval call binding the contract method 0x7c1a0302.
//
// Solidity: function requestNonce() view returns(uint64)
func (_Extbridge *ExtbridgeCaller) RequestNonce(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "requestNonce")
	return *ret0, err
}

// RequestNonce is a free data retrieval call binding the contract method 0x7c1a0302.
//
// Solidity: function requestNonce() view returns(uint64)
func (_Extbridge *ExtbridgeSession) RequestNonce() (uint64, error) {
	return _Extbridge.Contract.RequestNonce(&_Extbridge.CallOpts)
}

// RequestNonce is a free data retrieval call binding the contract method 0x7c1a0302.
//
// Solidity: function requestNonce() view returns(uint64)
func (_Extbridge *ExtbridgeCallerSession) RequestNonce() (uint64, error) {
	return _Extbridge.Contract.RequestNonce(&_Extbridge.CallOpts)
}

// UpperHandleNonce is a free data retrieval call binding the contract method 0x54edad72.
//
// Solidity: function upperHandleNonce() view returns(uint64)
func (_Extbridge *ExtbridgeCaller) UpperHandleNonce(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Extbridge.contract.Call(opts, out, "upperHandleNonce")
	return *ret0, err
}

// UpperHandleNonce is a free data retrieval call binding the contract method 0x54edad72.
//
// Solidity: function upperHandleNonce() view returns(uint64)
func (_Extbridge *ExtbridgeSession) UpperHandleNonce() (uint64, error) {
	return _Extbridge.Contract.UpperHandleNonce(&_Extbridge.CallOpts)
}

// UpperHandleNonce is a free data retrieval call binding the contract method 0x54edad72.
//
// Solidity: function upperHandleNonce() view returns(uint64)
func (_Extbridge *ExtbridgeCallerSession) UpperHandleNonce() (uint64, error) {
	return _Extbridge.Contract.UpperHandleNonce(&_Extbridge.CallOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xd8cf98ca.
//
// Solidity: function deregisterOperator(address _operator) returns()
func (_Extbridge *ExtbridgeTransactor) DeregisterOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "deregisterOperator", _operator)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xd8cf98ca.
//
// Solidity: function deregisterOperator(address _operator) returns()
func (_Extbridge *ExtbridgeSession) DeregisterOperator(_operator common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.DeregisterOperator(&_Extbridge.TransactOpts, _operator)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xd8cf98ca.
//
// Solidity: function deregisterOperator(address _operator) returns()
func (_Extbridge *ExtbridgeTransactorSession) DeregisterOperator(_operator common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.DeregisterOperator(&_Extbridge.TransactOpts, _operator)
}

// DeregisterToken is a paid mutator transaction binding the contract method 0xbab2af1d.
//
// Solidity: function deregisterToken(address _token) returns()
func (_Extbridge *ExtbridgeTransactor) DeregisterToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "deregisterToken", _token)
}

// DeregisterToken is a paid mutator transaction binding the contract method 0xbab2af1d.
//
// Solidity: function deregisterToken(address _token) returns()
func (_Extbridge *ExtbridgeSession) DeregisterToken(_token common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.DeregisterToken(&_Extbridge.TransactOpts, _token)
}

// DeregisterToken is a paid mutator transaction binding the contract method 0xbab2af1d.
//
// Solidity: function deregisterToken(address _token) returns()
func (_Extbridge *ExtbridgeTransactorSession) DeregisterToken(_token common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.DeregisterToken(&_Extbridge.TransactOpts, _token)
}

// HandleERC20Transfer is a paid mutator transaction binding the contract method 0x407e6bae.
//
// Solidity: function handleERC20Transfer(bytes32 _requestTxHash, address _from, address _to, address _tokenAddress, uint256 _value, uint64 _requestNonce, uint64 _requestBlockNumber, bytes _extraData) returns()
func (_Extbridge *ExtbridgeTransactor) HandleERC20Transfer(opts *bind.TransactOpts, _requestTxHash [32]byte, _from common.Address, _to common.Address, _tokenAddress common.Address, _value *big.Int, _requestNonce uint64, _requestBlockNumber uint64, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "handleERC20Transfer", _requestTxHash, _from, _to, _tokenAddress, _value, _requestNonce, _requestBlockNumber, _extraData)
}

// HandleERC20Transfer is a paid mutator transaction binding the contract method 0x407e6bae.
//
// Solidity: function handleERC20Transfer(bytes32 _requestTxHash, address _from, address _to, address _tokenAddress, uint256 _value, uint64 _requestNonce, uint64 _requestBlockNumber, bytes _extraData) returns()
func (_Extbridge *ExtbridgeSession) HandleERC20Transfer(_requestTxHash [32]byte, _from common.Address, _to common.Address, _tokenAddress common.Address, _value *big.Int, _requestNonce uint64, _requestBlockNumber uint64, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.Contract.HandleERC20Transfer(&_Extbridge.TransactOpts, _requestTxHash, _from, _to, _tokenAddress, _value, _requestNonce, _requestBlockNumber, _extraData)
}

// HandleERC20Transfer is a paid mutator transaction binding the contract method 0x407e6bae.
//
// Solidity: function handleERC20Transfer(bytes32 _requestTxHash, address _from, address _to, address _tokenAddress, uint256 _value, uint64 _requestNonce, uint64 _requestBlockNumber, bytes _extraData) returns()
func (_Extbridge *ExtbridgeTransactorSession) HandleERC20Transfer(_requestTxHash [32]byte, _from common.Address, _to common.Address, _tokenAddress common.Address, _value *big.Int, _requestNonce uint64, _requestBlockNumber uint64, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.Contract.HandleERC20Transfer(&_Extbridge.TransactOpts, _requestTxHash, _from, _to, _tokenAddress, _value, _requestNonce, _requestBlockNumber, _extraData)
}

// HandleERC721Transfer is a paid mutator transaction binding the contract method 0xafb60223.
//
// Solidity: function handleERC721Transfer(bytes32 _requestTxHash, address _from, address _to, address _tokenAddress, uint256 _tokenId, uint64 _requestNonce, uint64 _requestBlockNumber, string _tokenURI, bytes _extraData) returns()
func (_Extbridge *ExtbridgeTransactor) HandleERC721Transfer(opts *bind.TransactOpts, _requestTxHash [32]byte, _from common.Address, _to common.Address, _tokenAddress common.Address, _tokenId *big.Int, _requestNonce uint64, _requestBlockNumber uint64, _tokenURI string, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "handleERC721Transfer", _requestTxHash, _from, _to, _tokenAddress, _tokenId, _requestNonce, _requestBlockNumber, _tokenURI, _extraData)
}

// HandleERC721Transfer is a paid mutator transaction binding the contract method 0xafb60223.
//
// Solidity: function handleERC721Transfer(bytes32 _requestTxHash, address _from, address _to, address _tokenAddress, uint256 _tokenId, uint64 _requestNonce, uint64 _requestBlockNumber, string _tokenURI, bytes _extraData) returns()
func (_Extbridge *ExtbridgeSession) HandleERC721Transfer(_requestTxHash [32]byte, _from common.Address, _to common.Address, _tokenAddress common.Address, _tokenId *big.Int, _requestNonce uint64, _requestBlockNumber uint64, _tokenURI string, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.Contract.HandleERC721Transfer(&_Extbridge.TransactOpts, _requestTxHash, _from, _to, _tokenAddress, _tokenId, _requestNonce, _requestBlockNumber, _tokenURI, _extraData)
}

// HandleERC721Transfer is a paid mutator transaction binding the contract method 0xafb60223.
//
// Solidity: function handleERC721Transfer(bytes32 _requestTxHash, address _from, address _to, address _tokenAddress, uint256 _tokenId, uint64 _requestNonce, uint64 _requestBlockNumber, string _tokenURI, bytes _extraData) returns()
func (_Extbridge *ExtbridgeTransactorSession) HandleERC721Transfer(_requestTxHash [32]byte, _from common.Address, _to common.Address, _tokenAddress common.Address, _tokenId *big.Int, _requestNonce uint64, _requestBlockNumber uint64, _tokenURI string, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.Contract.HandleERC721Transfer(&_Extbridge.TransactOpts, _requestTxHash, _from, _to, _tokenAddress, _tokenId, _requestNonce, _requestBlockNumber, _tokenURI, _extraData)
}

// LockToken is a paid mutator transaction binding the contract method 0x10693fcd.
//
// Solidity: function lockToken(address _token) returns()
func (_Extbridge *ExtbridgeTransactor) LockToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "lockToken", _token)
}

// LockToken is a paid mutator transaction binding the contract method 0x10693fcd.
//
// Solidity: function lockToken(address _token) returns()
func (_Extbridge *ExtbridgeSession) LockToken(_token common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.LockToken(&_Extbridge.TransactOpts, _token)
}

// LockToken is a paid mutator transaction binding the contract method 0x10693fcd.
//
// Solidity: function lockToken(address _token) returns()
func (_Extbridge *ExtbridgeTransactorSession) LockToken(_token common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.LockToken(&_Extbridge.TransactOpts, _token)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0xf1656e53.
//
// Solidity: function onERC20Received(address _from, address _to, uint256 _value, uint256 _feeLimit, bytes _extraData) returns()
func (_Extbridge *ExtbridgeTransactor) OnERC20Received(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "onERC20Received", _from, _to, _value, _feeLimit, _extraData)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0xf1656e53.
//
// Solidity: function onERC20Received(address _from, address _to, uint256 _value, uint256 _feeLimit, bytes _extraData) returns()
func (_Extbridge *ExtbridgeSession) OnERC20Received(_from common.Address, _to common.Address, _value *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.Contract.OnERC20Received(&_Extbridge.TransactOpts, _from, _to, _value, _feeLimit, _extraData)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0xf1656e53.
//
// Solidity: function onERC20Received(address _from, address _to, uint256 _value, uint256 _feeLimit, bytes _extraData) returns()
func (_Extbridge *ExtbridgeTransactorSession) OnERC20Received(_from common.Address, _to common.Address, _value *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.Contract.OnERC20Received(&_Extbridge.TransactOpts, _from, _to, _value, _feeLimit, _extraData)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0xcf0da290.
//
// Solidity: function onERC721Received(address _from, uint256 _tokenId, address _to, bytes _extraData) returns()
func (_Extbridge *ExtbridgeTransactor) OnERC721Received(opts *bind.TransactOpts, _from common.Address, _tokenId *big.Int, _to common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "onERC721Received", _from, _tokenId, _to, _extraData)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0xcf0da290.
//
// Solidity: function onERC721Received(address _from, uint256 _tokenId, address _to, bytes _extraData) returns()
func (_Extbridge *ExtbridgeSession) OnERC721Received(_from common.Address, _tokenId *big.Int, _to common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.Contract.OnERC721Received(&_Extbridge.TransactOpts, _from, _tokenId, _to, _extraData)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0xcf0da290.
//
// Solidity: function onERC721Received(address _from, uint256 _tokenId, address _to, bytes _extraData) returns()
func (_Extbridge *ExtbridgeTransactorSession) OnERC721Received(_from common.Address, _tokenId *big.Int, _to common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.Contract.OnERC721Received(&_Extbridge.TransactOpts, _from, _tokenId, _to, _extraData)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address _operator) returns()
func (_Extbridge *ExtbridgeTransactor) RegisterOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "registerOperator", _operator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address _operator) returns()
func (_Extbridge *ExtbridgeSession) RegisterOperator(_operator common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.RegisterOperator(&_Extbridge.TransactOpts, _operator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address _operator) returns()
func (_Extbridge *ExtbridgeTransactorSession) RegisterOperator(_operator common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.RegisterOperator(&_Extbridge.TransactOpts, _operator)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x4739f7e5.
//
// Solidity: function registerToken(address _token, address _cToken) returns()
func (_Extbridge *ExtbridgeTransactor) RegisterToken(opts *bind.TransactOpts, _token common.Address, _cToken common.Address) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "registerToken", _token, _cToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x4739f7e5.
//
// Solidity: function registerToken(address _token, address _cToken) returns()
func (_Extbridge *ExtbridgeSession) RegisterToken(_token common.Address, _cToken common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.RegisterToken(&_Extbridge.TransactOpts, _token, _cToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x4739f7e5.
//
// Solidity: function registerToken(address _token, address _cToken) returns()
func (_Extbridge *ExtbridgeTransactorSession) RegisterToken(_token common.Address, _cToken common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.RegisterToken(&_Extbridge.TransactOpts, _token, _cToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Extbridge *ExtbridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Extbridge *ExtbridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Extbridge.Contract.RenounceOwnership(&_Extbridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Extbridge *ExtbridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Extbridge.Contract.RenounceOwnership(&_Extbridge.TransactOpts)
}

// RequestERC20Transfer is a paid mutator transaction binding the contract method 0x26c23b54.
//
// Solidity: function requestERC20Transfer(address _tokenAddress, address _to, uint256 _value, uint256 _feeLimit, bytes _extraData) returns()
func (_Extbridge *ExtbridgeTransactor) RequestERC20Transfer(opts *bind.TransactOpts, _tokenAddress common.Address, _to common.Address, _value *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "requestERC20Transfer", _tokenAddress, _to, _value, _feeLimit, _extraData)
}

// RequestERC20Transfer is a paid mutator transaction binding the contract method 0x26c23b54.
//
// Solidity: function requestERC20Transfer(address _tokenAddress, address _to, uint256 _value, uint256 _feeLimit, bytes _extraData) returns()
func (_Extbridge *ExtbridgeSession) RequestERC20Transfer(_tokenAddress common.Address, _to common.Address, _value *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.Contract.RequestERC20Transfer(&_Extbridge.TransactOpts, _tokenAddress, _to, _value, _feeLimit, _extraData)
}

// RequestERC20Transfer is a paid mutator transaction binding the contract method 0x26c23b54.
//
// Solidity: function requestERC20Transfer(address _tokenAddress, address _to, uint256 _value, uint256 _feeLimit, bytes _extraData) returns()
func (_Extbridge *ExtbridgeTransactorSession) RequestERC20Transfer(_tokenAddress common.Address, _to common.Address, _value *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.Contract.RequestERC20Transfer(&_Extbridge.TransactOpts, _tokenAddress, _to, _value, _feeLimit, _extraData)
}

// RequestERC721Transfer is a paid mutator transaction binding the contract method 0x22604742.
//
// Solidity: function requestERC721Transfer(address _tokenAddress, address _to, uint256 _tokenId, bytes _extraData) returns()
func (_Extbridge *ExtbridgeTransactor) RequestERC721Transfer(opts *bind.TransactOpts, _tokenAddress common.Address, _to common.Address, _tokenId *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "requestERC721Transfer", _tokenAddress, _to, _tokenId, _extraData)
}

// RequestERC721Transfer is a paid mutator transaction binding the contract method 0x22604742.
//
// Solidity: function requestERC721Transfer(address _tokenAddress, address _to, uint256 _tokenId, bytes _extraData) returns()
func (_Extbridge *ExtbridgeSession) RequestERC721Transfer(_tokenAddress common.Address, _to common.Address, _tokenId *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.Contract.RequestERC721Transfer(&_Extbridge.TransactOpts, _tokenAddress, _to, _tokenId, _extraData)
}

// RequestERC721Transfer is a paid mutator transaction binding the contract method 0x22604742.
//
// Solidity: function requestERC721Transfer(address _tokenAddress, address _to, uint256 _tokenId, bytes _extraData) returns()
func (_Extbridge *ExtbridgeTransactorSession) RequestERC721Transfer(_tokenAddress common.Address, _to common.Address, _tokenId *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Extbridge.Contract.RequestERC721Transfer(&_Extbridge.TransactOpts, _tokenAddress, _to, _tokenId, _extraData)
}

// RequestSellERC20 is a paid mutator transaction binding the contract method 0xc5e49073.
//
// Solidity: function requestSellERC20(address _tokenAddress, address _to, uint256 _value, uint256 _feeLimit, uint256 _price) returns()
func (_Extbridge *ExtbridgeTransactor) RequestSellERC20(opts *bind.TransactOpts, _tokenAddress common.Address, _to common.Address, _value *big.Int, _feeLimit *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "requestSellERC20", _tokenAddress, _to, _value, _feeLimit, _price)
}

// RequestSellERC20 is a paid mutator transaction binding the contract method 0xc5e49073.
//
// Solidity: function requestSellERC20(address _tokenAddress, address _to, uint256 _value, uint256 _feeLimit, uint256 _price) returns()
func (_Extbridge *ExtbridgeSession) RequestSellERC20(_tokenAddress common.Address, _to common.Address, _value *big.Int, _feeLimit *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Extbridge.Contract.RequestSellERC20(&_Extbridge.TransactOpts, _tokenAddress, _to, _value, _feeLimit, _price)
}

// RequestSellERC20 is a paid mutator transaction binding the contract method 0xc5e49073.
//
// Solidity: function requestSellERC20(address _tokenAddress, address _to, uint256 _value, uint256 _feeLimit, uint256 _price) returns()
func (_Extbridge *ExtbridgeTransactorSession) RequestSellERC20(_tokenAddress common.Address, _to common.Address, _value *big.Int, _feeLimit *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Extbridge.Contract.RequestSellERC20(&_Extbridge.TransactOpts, _tokenAddress, _to, _value, _feeLimit, _price)
}

// RequestSellERC721 is a paid mutator transaction binding the contract method 0x4c5146f5.
//
// Solidity: function requestSellERC721(address _tokenAddress, address _to, uint256 _tokenId, uint256 _price) returns()
func (_Extbridge *ExtbridgeTransactor) RequestSellERC721(opts *bind.TransactOpts, _tokenAddress common.Address, _to common.Address, _tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "requestSellERC721", _tokenAddress, _to, _tokenId, _price)
}

// RequestSellERC721 is a paid mutator transaction binding the contract method 0x4c5146f5.
//
// Solidity: function requestSellERC721(address _tokenAddress, address _to, uint256 _tokenId, uint256 _price) returns()
func (_Extbridge *ExtbridgeSession) RequestSellERC721(_tokenAddress common.Address, _to common.Address, _tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Extbridge.Contract.RequestSellERC721(&_Extbridge.TransactOpts, _tokenAddress, _to, _tokenId, _price)
}

// RequestSellERC721 is a paid mutator transaction binding the contract method 0x4c5146f5.
//
// Solidity: function requestSellERC721(address _tokenAddress, address _to, uint256 _tokenId, uint256 _price) returns()
func (_Extbridge *ExtbridgeTransactorSession) RequestSellERC721(_tokenAddress common.Address, _to common.Address, _tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Extbridge.Contract.RequestSellERC721(&_Extbridge.TransactOpts, _tokenAddress, _to, _tokenId, _price)
}

// SetCallback is a paid mutator transaction binding the contract method 0x8daa63ac.
//
// Solidity: function setCallback(address _addr) returns()
func (_Extbridge *ExtbridgeTransactor) SetCallback(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "setCallback", _addr)
}

// SetCallback is a paid mutator transaction binding the contract method 0x8daa63ac.
//
// Solidity: function setCallback(address _addr) returns()
func (_Extbridge *ExtbridgeSession) SetCallback(_addr common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.SetCallback(&_Extbridge.TransactOpts, _addr)
}

// SetCallback is a paid mutator transaction binding the contract method 0x8daa63ac.
//
// Solidity: function setCallback(address _addr) returns()
func (_Extbridge *ExtbridgeTransactorSession) SetCallback(_addr common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.SetCallback(&_Extbridge.TransactOpts, _addr)
}

// SetERC20Fee is a paid mutator transaction binding the contract method 0x2f88396c.
//
// Solidity: function setERC20Fee(address _token, uint256 _fee, uint64 _requestNonce) returns()
func (_Extbridge *ExtbridgeTransactor) SetERC20Fee(opts *bind.TransactOpts, _token common.Address, _fee *big.Int, _requestNonce uint64) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "setERC20Fee", _token, _fee, _requestNonce)
}

// SetERC20Fee is a paid mutator transaction binding the contract method 0x2f88396c.
//
// Solidity: function setERC20Fee(address _token, uint256 _fee, uint64 _requestNonce) returns()
func (_Extbridge *ExtbridgeSession) SetERC20Fee(_token common.Address, _fee *big.Int, _requestNonce uint64) (*types.Transaction, error) {
	return _Extbridge.Contract.SetERC20Fee(&_Extbridge.TransactOpts, _token, _fee, _requestNonce)
}

// SetERC20Fee is a paid mutator transaction binding the contract method 0x2f88396c.
//
// Solidity: function setERC20Fee(address _token, uint256 _fee, uint64 _requestNonce) returns()
func (_Extbridge *ExtbridgeTransactorSession) SetERC20Fee(_token common.Address, _fee *big.Int, _requestNonce uint64) (*types.Transaction, error) {
	return _Extbridge.Contract.SetERC20Fee(&_Extbridge.TransactOpts, _token, _fee, _requestNonce)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiver) returns()
func (_Extbridge *ExtbridgeTransactor) SetFeeReceiver(opts *bind.TransactOpts, _feeReceiver common.Address) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "setFeeReceiver", _feeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiver) returns()
func (_Extbridge *ExtbridgeSession) SetFeeReceiver(_feeReceiver common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.SetFeeReceiver(&_Extbridge.TransactOpts, _feeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiver) returns()
func (_Extbridge *ExtbridgeTransactorSession) SetFeeReceiver(_feeReceiver common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.SetFeeReceiver(&_Extbridge.TransactOpts, _feeReceiver)
}

// SetOperatorThreshold is a paid mutator transaction binding the contract method 0xee2aec65.
//
// Solidity: function setOperatorThreshold(uint8 _voteType, uint8 _threshold) returns()
func (_Extbridge *ExtbridgeTransactor) SetOperatorThreshold(opts *bind.TransactOpts, _voteType uint8, _threshold uint8) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "setOperatorThreshold", _voteType, _threshold)
}

// SetOperatorThreshold is a paid mutator transaction binding the contract method 0xee2aec65.
//
// Solidity: function setOperatorThreshold(uint8 _voteType, uint8 _threshold) returns()
func (_Extbridge *ExtbridgeSession) SetOperatorThreshold(_voteType uint8, _threshold uint8) (*types.Transaction, error) {
	return _Extbridge.Contract.SetOperatorThreshold(&_Extbridge.TransactOpts, _voteType, _threshold)
}

// SetOperatorThreshold is a paid mutator transaction binding the contract method 0xee2aec65.
//
// Solidity: function setOperatorThreshold(uint8 _voteType, uint8 _threshold) returns()
func (_Extbridge *ExtbridgeTransactorSession) SetOperatorThreshold(_voteType uint8, _threshold uint8) (*types.Transaction, error) {
	return _Extbridge.Contract.SetOperatorThreshold(&_Extbridge.TransactOpts, _voteType, _threshold)
}

// Start is a paid mutator transaction binding the contract method 0xc877cf37.
//
// Solidity: function start(bool _status) returns()
func (_Extbridge *ExtbridgeTransactor) Start(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "start", _status)
}

// Start is a paid mutator transaction binding the contract method 0xc877cf37.
//
// Solidity: function start(bool _status) returns()
func (_Extbridge *ExtbridgeSession) Start(_status bool) (*types.Transaction, error) {
	return _Extbridge.Contract.Start(&_Extbridge.TransactOpts, _status)
}

// Start is a paid mutator transaction binding the contract method 0xc877cf37.
//
// Solidity: function start(bool _status) returns()
func (_Extbridge *ExtbridgeTransactorSession) Start(_status bool) (*types.Transaction, error) {
	return _Extbridge.Contract.Start(&_Extbridge.TransactOpts, _status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Extbridge *ExtbridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Extbridge *ExtbridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.TransferOwnership(&_Extbridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Extbridge *ExtbridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.TransferOwnership(&_Extbridge.TransactOpts, newOwner)
}

// UnlockToken is a paid mutator transaction binding the contract method 0x9ef2017b.
//
// Solidity: function unlockToken(address _token) returns()
func (_Extbridge *ExtbridgeTransactor) UnlockToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Extbridge.contract.Transact(opts, "unlockToken", _token)
}

// UnlockToken is a paid mutator transaction binding the contract method 0x9ef2017b.
//
// Solidity: function unlockToken(address _token) returns()
func (_Extbridge *ExtbridgeSession) UnlockToken(_token common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.UnlockToken(&_Extbridge.TransactOpts, _token)
}

// UnlockToken is a paid mutator transaction binding the contract method 0x9ef2017b.
//
// Solidity: function unlockToken(address _token) returns()
func (_Extbridge *ExtbridgeTransactorSession) UnlockToken(_token common.Address) (*types.Transaction, error) {
	return _Extbridge.Contract.UnlockToken(&_Extbridge.TransactOpts, _token)
}

// ExtbridgeERC20FeeChangedIterator is returned from FilterERC20FeeChanged and is used to iterate over the raw logs and unpacked data for ERC20FeeChanged events raised by the Extbridge contract.
type ExtbridgeERC20FeeChangedIterator struct {
	Event *ExtbridgeERC20FeeChanged // Event containing the contract specifics and raw log

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
func (it *ExtbridgeERC20FeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtbridgeERC20FeeChanged)
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
		it.Event = new(ExtbridgeERC20FeeChanged)
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
func (it *ExtbridgeERC20FeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtbridgeERC20FeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtbridgeERC20FeeChanged represents a ERC20FeeChanged event raised by the Extbridge contract.
type ExtbridgeERC20FeeChanged struct {
	Token common.Address
	Fee   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterERC20FeeChanged is a free log retrieval operation binding the contract event 0xdb5ad2e76ae20cfa4e7adbc7305d7538442164d85ead9937c98620a1aa4c255b.
//
// Solidity: event ERC20FeeChanged(address indexed token, uint256 indexed fee)
func (_Extbridge *ExtbridgeFilterer) FilterERC20FeeChanged(opts *bind.FilterOpts, token []common.Address, fee []*big.Int) (*ExtbridgeERC20FeeChangedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Extbridge.contract.FilterLogs(opts, "ERC20FeeChanged", tokenRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &ExtbridgeERC20FeeChangedIterator{contract: _Extbridge.contract, event: "ERC20FeeChanged", logs: logs, sub: sub}, nil
}

// WatchERC20FeeChanged is a free log subscription operation binding the contract event 0xdb5ad2e76ae20cfa4e7adbc7305d7538442164d85ead9937c98620a1aa4c255b.
//
// Solidity: event ERC20FeeChanged(address indexed token, uint256 indexed fee)
func (_Extbridge *ExtbridgeFilterer) WatchERC20FeeChanged(opts *bind.WatchOpts, sink chan<- *ExtbridgeERC20FeeChanged, token []common.Address, fee []*big.Int) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Extbridge.contract.WatchLogs(opts, "ERC20FeeChanged", tokenRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtbridgeERC20FeeChanged)
				if err := _Extbridge.contract.UnpackLog(event, "ERC20FeeChanged", log); err != nil {
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
func (_Extbridge *ExtbridgeFilterer) ParseERC20FeeChanged(log types.Log) (*ExtbridgeERC20FeeChanged, error) {
	event := new(ExtbridgeERC20FeeChanged)
	if err := _Extbridge.contract.UnpackLog(event, "ERC20FeeChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExtbridgeFeeReceiverChangedIterator is returned from FilterFeeReceiverChanged and is used to iterate over the raw logs and unpacked data for FeeReceiverChanged events raised by the Extbridge contract.
type ExtbridgeFeeReceiverChangedIterator struct {
	Event *ExtbridgeFeeReceiverChanged // Event containing the contract specifics and raw log

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
func (it *ExtbridgeFeeReceiverChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtbridgeFeeReceiverChanged)
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
		it.Event = new(ExtbridgeFeeReceiverChanged)
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
func (it *ExtbridgeFeeReceiverChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtbridgeFeeReceiverChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtbridgeFeeReceiverChanged represents a FeeReceiverChanged event raised by the Extbridge contract.
type ExtbridgeFeeReceiverChanged struct {
	FeeReceiver common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeReceiverChanged is a free log retrieval operation binding the contract event 0x647672599d3468abcfa241a13c9e3d34383caadb5cc80fb67c3cdfcd5f786059.
//
// Solidity: event FeeReceiverChanged(address indexed feeReceiver)
func (_Extbridge *ExtbridgeFilterer) FilterFeeReceiverChanged(opts *bind.FilterOpts, feeReceiver []common.Address) (*ExtbridgeFeeReceiverChangedIterator, error) {

	var feeReceiverRule []interface{}
	for _, feeReceiverItem := range feeReceiver {
		feeReceiverRule = append(feeReceiverRule, feeReceiverItem)
	}

	logs, sub, err := _Extbridge.contract.FilterLogs(opts, "FeeReceiverChanged", feeReceiverRule)
	if err != nil {
		return nil, err
	}
	return &ExtbridgeFeeReceiverChangedIterator{contract: _Extbridge.contract, event: "FeeReceiverChanged", logs: logs, sub: sub}, nil
}

// WatchFeeReceiverChanged is a free log subscription operation binding the contract event 0x647672599d3468abcfa241a13c9e3d34383caadb5cc80fb67c3cdfcd5f786059.
//
// Solidity: event FeeReceiverChanged(address indexed feeReceiver)
func (_Extbridge *ExtbridgeFilterer) WatchFeeReceiverChanged(opts *bind.WatchOpts, sink chan<- *ExtbridgeFeeReceiverChanged, feeReceiver []common.Address) (event.Subscription, error) {

	var feeReceiverRule []interface{}
	for _, feeReceiverItem := range feeReceiver {
		feeReceiverRule = append(feeReceiverRule, feeReceiverItem)
	}

	logs, sub, err := _Extbridge.contract.WatchLogs(opts, "FeeReceiverChanged", feeReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtbridgeFeeReceiverChanged)
				if err := _Extbridge.contract.UnpackLog(event, "FeeReceiverChanged", log); err != nil {
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
func (_Extbridge *ExtbridgeFilterer) ParseFeeReceiverChanged(log types.Log) (*ExtbridgeFeeReceiverChanged, error) {
	event := new(ExtbridgeFeeReceiverChanged)
	if err := _Extbridge.contract.UnpackLog(event, "FeeReceiverChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExtbridgeHandleValueTransferIterator is returned from FilterHandleValueTransfer and is used to iterate over the raw logs and unpacked data for HandleValueTransfer events raised by the Extbridge contract.
type ExtbridgeHandleValueTransferIterator struct {
	Event *ExtbridgeHandleValueTransfer // Event containing the contract specifics and raw log

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
func (it *ExtbridgeHandleValueTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtbridgeHandleValueTransfer)
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
		it.Event = new(ExtbridgeHandleValueTransfer)
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
func (it *ExtbridgeHandleValueTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtbridgeHandleValueTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtbridgeHandleValueTransfer represents a HandleValueTransfer event raised by the Extbridge contract.
type ExtbridgeHandleValueTransfer struct {
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
func (_Extbridge *ExtbridgeFilterer) FilterHandleValueTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenAddress []common.Address) (*ExtbridgeHandleValueTransferIterator, error) {

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

	logs, sub, err := _Extbridge.contract.FilterLogs(opts, "HandleValueTransfer", fromRule, toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ExtbridgeHandleValueTransferIterator{contract: _Extbridge.contract, event: "HandleValueTransfer", logs: logs, sub: sub}, nil
}

// WatchHandleValueTransfer is a free log subscription operation binding the contract event 0x12b02f226d965a2881e0e8ffed6c421803a22d57ad91f9ef996fe0748ea10175.
//
// Solidity: event HandleValueTransfer(bytes32 requestTxHash, uint8 tokenType, address indexed from, address indexed to, address indexed tokenAddress, uint256 valueOrTokenId, uint64 handleNonce, uint64 lowerHandleNonce, bytes extraData)
func (_Extbridge *ExtbridgeFilterer) WatchHandleValueTransfer(opts *bind.WatchOpts, sink chan<- *ExtbridgeHandleValueTransfer, from []common.Address, to []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Extbridge.contract.WatchLogs(opts, "HandleValueTransfer", fromRule, toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtbridgeHandleValueTransfer)
				if err := _Extbridge.contract.UnpackLog(event, "HandleValueTransfer", log); err != nil {
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
func (_Extbridge *ExtbridgeFilterer) ParseHandleValueTransfer(log types.Log) (*ExtbridgeHandleValueTransfer, error) {
	event := new(ExtbridgeHandleValueTransfer)
	if err := _Extbridge.contract.UnpackLog(event, "HandleValueTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExtbridgeKLAYFeeChangedIterator is returned from FilterKLAYFeeChanged and is used to iterate over the raw logs and unpacked data for KLAYFeeChanged events raised by the Extbridge contract.
type ExtbridgeKLAYFeeChangedIterator struct {
	Event *ExtbridgeKLAYFeeChanged // Event containing the contract specifics and raw log

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
func (it *ExtbridgeKLAYFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtbridgeKLAYFeeChanged)
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
		it.Event = new(ExtbridgeKLAYFeeChanged)
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
func (it *ExtbridgeKLAYFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtbridgeKLAYFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtbridgeKLAYFeeChanged represents a KLAYFeeChanged event raised by the Extbridge contract.
type ExtbridgeKLAYFeeChanged struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterKLAYFeeChanged is a free log retrieval operation binding the contract event 0xa7a33d0996347e1aa55ca2206015b61b9534bdd881d59d59aa680e25eefac365.
//
// Solidity: event KLAYFeeChanged(uint256 indexed fee)
func (_Extbridge *ExtbridgeFilterer) FilterKLAYFeeChanged(opts *bind.FilterOpts, fee []*big.Int) (*ExtbridgeKLAYFeeChangedIterator, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Extbridge.contract.FilterLogs(opts, "KLAYFeeChanged", feeRule)
	if err != nil {
		return nil, err
	}
	return &ExtbridgeKLAYFeeChangedIterator{contract: _Extbridge.contract, event: "KLAYFeeChanged", logs: logs, sub: sub}, nil
}

// WatchKLAYFeeChanged is a free log subscription operation binding the contract event 0xa7a33d0996347e1aa55ca2206015b61b9534bdd881d59d59aa680e25eefac365.
//
// Solidity: event KLAYFeeChanged(uint256 indexed fee)
func (_Extbridge *ExtbridgeFilterer) WatchKLAYFeeChanged(opts *bind.WatchOpts, sink chan<- *ExtbridgeKLAYFeeChanged, fee []*big.Int) (event.Subscription, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Extbridge.contract.WatchLogs(opts, "KLAYFeeChanged", feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtbridgeKLAYFeeChanged)
				if err := _Extbridge.contract.UnpackLog(event, "KLAYFeeChanged", log); err != nil {
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
func (_Extbridge *ExtbridgeFilterer) ParseKLAYFeeChanged(log types.Log) (*ExtbridgeKLAYFeeChanged, error) {
	event := new(ExtbridgeKLAYFeeChanged)
	if err := _Extbridge.contract.UnpackLog(event, "KLAYFeeChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExtbridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Extbridge contract.
type ExtbridgeOwnershipTransferredIterator struct {
	Event *ExtbridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExtbridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtbridgeOwnershipTransferred)
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
		it.Event = new(ExtbridgeOwnershipTransferred)
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
func (it *ExtbridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtbridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtbridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Extbridge contract.
type ExtbridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Extbridge *ExtbridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExtbridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Extbridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExtbridgeOwnershipTransferredIterator{contract: _Extbridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Extbridge *ExtbridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExtbridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Extbridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtbridgeOwnershipTransferred)
				if err := _Extbridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Extbridge *ExtbridgeFilterer) ParseOwnershipTransferred(log types.Log) (*ExtbridgeOwnershipTransferred, error) {
	event := new(ExtbridgeOwnershipTransferred)
	if err := _Extbridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExtbridgeRequestValueTransferIterator is returned from FilterRequestValueTransfer and is used to iterate over the raw logs and unpacked data for RequestValueTransfer events raised by the Extbridge contract.
type ExtbridgeRequestValueTransferIterator struct {
	Event *ExtbridgeRequestValueTransfer // Event containing the contract specifics and raw log

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
func (it *ExtbridgeRequestValueTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtbridgeRequestValueTransfer)
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
		it.Event = new(ExtbridgeRequestValueTransfer)
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
func (it *ExtbridgeRequestValueTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtbridgeRequestValueTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtbridgeRequestValueTransfer represents a RequestValueTransfer event raised by the Extbridge contract.
type ExtbridgeRequestValueTransfer struct {
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
func (_Extbridge *ExtbridgeFilterer) FilterRequestValueTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenAddress []common.Address) (*ExtbridgeRequestValueTransferIterator, error) {

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

	logs, sub, err := _Extbridge.contract.FilterLogs(opts, "RequestValueTransfer", fromRule, toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ExtbridgeRequestValueTransferIterator{contract: _Extbridge.contract, event: "RequestValueTransfer", logs: logs, sub: sub}, nil
}

// WatchRequestValueTransfer is a free log subscription operation binding the contract event 0xeff76c36e53fa5ff52f27acc8a34d5047a8246abb07b77b12f1309f71e337f09.
//
// Solidity: event RequestValueTransfer(uint8 tokenType, address indexed from, address indexed to, address indexed tokenAddress, uint256 valueOrTokenId, uint64 requestNonce, uint256 fee, bytes extraData)
func (_Extbridge *ExtbridgeFilterer) WatchRequestValueTransfer(opts *bind.WatchOpts, sink chan<- *ExtbridgeRequestValueTransfer, from []common.Address, to []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Extbridge.contract.WatchLogs(opts, "RequestValueTransfer", fromRule, toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtbridgeRequestValueTransfer)
				if err := _Extbridge.contract.UnpackLog(event, "RequestValueTransfer", log); err != nil {
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
func (_Extbridge *ExtbridgeFilterer) ParseRequestValueTransfer(log types.Log) (*ExtbridgeRequestValueTransfer, error) {
	event := new(ExtbridgeRequestValueTransfer)
	if err := _Extbridge.contract.UnpackLog(event, "RequestValueTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExtbridgeRequestValueTransferEncodedIterator is returned from FilterRequestValueTransferEncoded and is used to iterate over the raw logs and unpacked data for RequestValueTransferEncoded events raised by the Extbridge contract.
type ExtbridgeRequestValueTransferEncodedIterator struct {
	Event *ExtbridgeRequestValueTransferEncoded // Event containing the contract specifics and raw log

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
func (it *ExtbridgeRequestValueTransferEncodedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtbridgeRequestValueTransferEncoded)
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
		it.Event = new(ExtbridgeRequestValueTransferEncoded)
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
func (it *ExtbridgeRequestValueTransferEncodedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtbridgeRequestValueTransferEncodedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtbridgeRequestValueTransferEncoded represents a RequestValueTransferEncoded event raised by the Extbridge contract.
type ExtbridgeRequestValueTransferEncoded struct {
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
func (_Extbridge *ExtbridgeFilterer) FilterRequestValueTransferEncoded(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenAddress []common.Address) (*ExtbridgeRequestValueTransferEncodedIterator, error) {

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

	logs, sub, err := _Extbridge.contract.FilterLogs(opts, "RequestValueTransferEncoded", fromRule, toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ExtbridgeRequestValueTransferEncodedIterator{contract: _Extbridge.contract, event: "RequestValueTransferEncoded", logs: logs, sub: sub}, nil
}

// WatchRequestValueTransferEncoded is a free log subscription operation binding the contract event 0x17d76053ca34a4dd8c402fe6498deb797fac89bf7ed02f3f5161aa9368cc8c1f.
//
// Solidity: event RequestValueTransferEncoded(uint8 tokenType, address indexed from, address indexed to, address indexed tokenAddress, uint256 valueOrTokenId, uint64 requestNonce, uint256 fee, bytes extraData, uint8 encodingVer, bytes encodedData)
func (_Extbridge *ExtbridgeFilterer) WatchRequestValueTransferEncoded(opts *bind.WatchOpts, sink chan<- *ExtbridgeRequestValueTransferEncoded, from []common.Address, to []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Extbridge.contract.WatchLogs(opts, "RequestValueTransferEncoded", fromRule, toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtbridgeRequestValueTransferEncoded)
				if err := _Extbridge.contract.UnpackLog(event, "RequestValueTransferEncoded", log); err != nil {
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
func (_Extbridge *ExtbridgeFilterer) ParseRequestValueTransferEncoded(log types.Log) (*ExtbridgeRequestValueTransferEncoded, error) {
	event := new(ExtbridgeRequestValueTransferEncoded)
	if err := _Extbridge.contract.UnpackLog(event, "RequestValueTransferEncoded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExtbridgeTokenDeregisteredIterator is returned from FilterTokenDeregistered and is used to iterate over the raw logs and unpacked data for TokenDeregistered events raised by the Extbridge contract.
type ExtbridgeTokenDeregisteredIterator struct {
	Event *ExtbridgeTokenDeregistered // Event containing the contract specifics and raw log

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
func (it *ExtbridgeTokenDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtbridgeTokenDeregistered)
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
		it.Event = new(ExtbridgeTokenDeregistered)
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
func (it *ExtbridgeTokenDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtbridgeTokenDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtbridgeTokenDeregistered represents a TokenDeregistered event raised by the Extbridge contract.
type ExtbridgeTokenDeregistered struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenDeregistered is a free log retrieval operation binding the contract event 0x1d735ca20b63676dde668b718be78606b061d6bd7534ff815a90a121a6c084b6.
//
// Solidity: event TokenDeregistered(address indexed token)
func (_Extbridge *ExtbridgeFilterer) FilterTokenDeregistered(opts *bind.FilterOpts, token []common.Address) (*ExtbridgeTokenDeregisteredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Extbridge.contract.FilterLogs(opts, "TokenDeregistered", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ExtbridgeTokenDeregisteredIterator{contract: _Extbridge.contract, event: "TokenDeregistered", logs: logs, sub: sub}, nil
}

// WatchTokenDeregistered is a free log subscription operation binding the contract event 0x1d735ca20b63676dde668b718be78606b061d6bd7534ff815a90a121a6c084b6.
//
// Solidity: event TokenDeregistered(address indexed token)
func (_Extbridge *ExtbridgeFilterer) WatchTokenDeregistered(opts *bind.WatchOpts, sink chan<- *ExtbridgeTokenDeregistered, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Extbridge.contract.WatchLogs(opts, "TokenDeregistered", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtbridgeTokenDeregistered)
				if err := _Extbridge.contract.UnpackLog(event, "TokenDeregistered", log); err != nil {
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
func (_Extbridge *ExtbridgeFilterer) ParseTokenDeregistered(log types.Log) (*ExtbridgeTokenDeregistered, error) {
	event := new(ExtbridgeTokenDeregistered)
	if err := _Extbridge.contract.UnpackLog(event, "TokenDeregistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExtbridgeTokenLockedIterator is returned from FilterTokenLocked and is used to iterate over the raw logs and unpacked data for TokenLocked events raised by the Extbridge contract.
type ExtbridgeTokenLockedIterator struct {
	Event *ExtbridgeTokenLocked // Event containing the contract specifics and raw log

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
func (it *ExtbridgeTokenLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtbridgeTokenLocked)
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
		it.Event = new(ExtbridgeTokenLocked)
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
func (it *ExtbridgeTokenLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtbridgeTokenLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtbridgeTokenLocked represents a TokenLocked event raised by the Extbridge contract.
type ExtbridgeTokenLocked struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenLocked is a free log retrieval operation binding the contract event 0xca1b0a14e18ada4c44846768dd186e35630cdc5cfeaca83c404ae4acaafbecd7.
//
// Solidity: event TokenLocked(address indexed token)
func (_Extbridge *ExtbridgeFilterer) FilterTokenLocked(opts *bind.FilterOpts, token []common.Address) (*ExtbridgeTokenLockedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Extbridge.contract.FilterLogs(opts, "TokenLocked", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ExtbridgeTokenLockedIterator{contract: _Extbridge.contract, event: "TokenLocked", logs: logs, sub: sub}, nil
}

// WatchTokenLocked is a free log subscription operation binding the contract event 0xca1b0a14e18ada4c44846768dd186e35630cdc5cfeaca83c404ae4acaafbecd7.
//
// Solidity: event TokenLocked(address indexed token)
func (_Extbridge *ExtbridgeFilterer) WatchTokenLocked(opts *bind.WatchOpts, sink chan<- *ExtbridgeTokenLocked, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Extbridge.contract.WatchLogs(opts, "TokenLocked", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtbridgeTokenLocked)
				if err := _Extbridge.contract.UnpackLog(event, "TokenLocked", log); err != nil {
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
func (_Extbridge *ExtbridgeFilterer) ParseTokenLocked(log types.Log) (*ExtbridgeTokenLocked, error) {
	event := new(ExtbridgeTokenLocked)
	if err := _Extbridge.contract.UnpackLog(event, "TokenLocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExtbridgeTokenRegisteredIterator is returned from FilterTokenRegistered and is used to iterate over the raw logs and unpacked data for TokenRegistered events raised by the Extbridge contract.
type ExtbridgeTokenRegisteredIterator struct {
	Event *ExtbridgeTokenRegistered // Event containing the contract specifics and raw log

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
func (it *ExtbridgeTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtbridgeTokenRegistered)
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
		it.Event = new(ExtbridgeTokenRegistered)
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
func (it *ExtbridgeTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtbridgeTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtbridgeTokenRegistered represents a TokenRegistered event raised by the Extbridge contract.
type ExtbridgeTokenRegistered struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistered is a free log retrieval operation binding the contract event 0x158412daecdc1456d01568828bcdb18464cc7f1ce0215ddbc3f3cfede9d1e63d.
//
// Solidity: event TokenRegistered(address indexed token)
func (_Extbridge *ExtbridgeFilterer) FilterTokenRegistered(opts *bind.FilterOpts, token []common.Address) (*ExtbridgeTokenRegisteredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Extbridge.contract.FilterLogs(opts, "TokenRegistered", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ExtbridgeTokenRegisteredIterator{contract: _Extbridge.contract, event: "TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenRegistered is a free log subscription operation binding the contract event 0x158412daecdc1456d01568828bcdb18464cc7f1ce0215ddbc3f3cfede9d1e63d.
//
// Solidity: event TokenRegistered(address indexed token)
func (_Extbridge *ExtbridgeFilterer) WatchTokenRegistered(opts *bind.WatchOpts, sink chan<- *ExtbridgeTokenRegistered, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Extbridge.contract.WatchLogs(opts, "TokenRegistered", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtbridgeTokenRegistered)
				if err := _Extbridge.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
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
func (_Extbridge *ExtbridgeFilterer) ParseTokenRegistered(log types.Log) (*ExtbridgeTokenRegistered, error) {
	event := new(ExtbridgeTokenRegistered)
	if err := _Extbridge.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExtbridgeTokenUnlockedIterator is returned from FilterTokenUnlocked and is used to iterate over the raw logs and unpacked data for TokenUnlocked events raised by the Extbridge contract.
type ExtbridgeTokenUnlockedIterator struct {
	Event *ExtbridgeTokenUnlocked // Event containing the contract specifics and raw log

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
func (it *ExtbridgeTokenUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtbridgeTokenUnlocked)
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
		it.Event = new(ExtbridgeTokenUnlocked)
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
func (it *ExtbridgeTokenUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtbridgeTokenUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtbridgeTokenUnlocked represents a TokenUnlocked event raised by the Extbridge contract.
type ExtbridgeTokenUnlocked struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenUnlocked is a free log retrieval operation binding the contract event 0x81ec08d3372506e176c49e626d8beb7e091712ef92908a130f4ccc6524fe2eec.
//
// Solidity: event TokenUnlocked(address indexed token)
func (_Extbridge *ExtbridgeFilterer) FilterTokenUnlocked(opts *bind.FilterOpts, token []common.Address) (*ExtbridgeTokenUnlockedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Extbridge.contract.FilterLogs(opts, "TokenUnlocked", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ExtbridgeTokenUnlockedIterator{contract: _Extbridge.contract, event: "TokenUnlocked", logs: logs, sub: sub}, nil
}

// WatchTokenUnlocked is a free log subscription operation binding the contract event 0x81ec08d3372506e176c49e626d8beb7e091712ef92908a130f4ccc6524fe2eec.
//
// Solidity: event TokenUnlocked(address indexed token)
func (_Extbridge *ExtbridgeFilterer) WatchTokenUnlocked(opts *bind.WatchOpts, sink chan<- *ExtbridgeTokenUnlocked, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Extbridge.contract.WatchLogs(opts, "TokenUnlocked", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtbridgeTokenUnlocked)
				if err := _Extbridge.contract.UnpackLog(event, "TokenUnlocked", log); err != nil {
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
func (_Extbridge *ExtbridgeFilterer) ParseTokenUnlocked(log types.Log) (*ExtbridgeTokenUnlocked, error) {
	event := new(ExtbridgeTokenUnlocked)
	if err := _Extbridge.contract.UnpackLog(event, "TokenUnlocked", log); err != nil {
		return nil, err
	}
	return event, nil
}