## The `Solidity` Tool

## Solc is available as a Homebrew package for macOS.
```shell
brew update
brew tap ethereum/ethereum
brew install solidity
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
