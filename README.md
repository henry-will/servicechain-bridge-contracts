## The `Solidity` Tool

## Solc is available as a Homebrew package for macOS.
```shell
brew update
brew tap ethereum/ethereum
brew install solidity
```

## The `GOPATH` environment 
### For convenience, add the workspace's bin subdirectory to your PATH:
```shell
$ export PATH=$PATH:$(go env GOPATH)/bin
```

## The `abigen` Tool
#### Download go-ethereum, build and install the devtools (which includes abigen)
```
git clone https://github.com/ethereum/go-ethereum.git
cd go-ethereum
make devtools
```

### Run abigen and print the version
```
abigen -version
abigen -help
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
### Create the Go binding with bytecode
```
$ node generate-go-contracts.js
```

### package for klaytn projects
