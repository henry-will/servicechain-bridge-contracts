# This is a ServiceChain bridge
It is for making servicechain bridge and for generating go files for klaytn from bridge contract.
### Versions used to build 
- npm 6.14.16
- node v12.22.10
- truffle 5.5.18
- solc 0.8.14+commit.80d49f37.Emscripten.clang 
- klaytn abigen v1.8.4-464dcb16
- klaytn 1.9.0
- go 1.18.2

## The `Solidity` Tool

## Solc is available as a Homebrew package for macOS.
```shell
brew update
brew tap ethereum/ethereum
brew install solidity
```

## Install node modules
- yarn add truffle
- yarn add @chainsafe/truffle-plugin-abigen
- yarn add @klaytn/contracts@1.0.2
```shell
yarn install
```

## The `Truffle` Tool
### Compile the contract
```
$ yarn truffle compile
```

### Generate the ABI and BIN files (stored in abigenBindings/)
```
$ yarn truffle run abigen
```

## Create Go code from solidity bytecode
### use klaytn abigen
Build abigen
```shell
$ git clone https://github.com/klaytn/klaytn.git
$ make abigen
$ cp ${klaytn git clone home}/build/bin/abigen ${servicechain-bridge-contracts git clone home}/bin/abigen-${klaytn version}
ex) cp ./build/bin/abigen ../servicechain-bridge-contracts/bin/abi-v1.9.0
```
### Create the Go binding with bytecode
```
$ node go-bindings.js
```

### package for klaytn projects
- go.mod in klaytn
```
require ( 
  github.com/henry-will/servicechain-bridge-contracts v0.9.7
)

```
import  servicechain-bridge-contracts
```go
import "github.com/klaytn/servicechain-bridge-contracts/scnft"
```
