## Install Dependency

1. Ethereum solidity complier.

```bash
docker pull ethereum/solc:0.6.12 
```

2. Ethereum ABI generator from go-ethereum
```bash
$ go install github.com/ethereum/go-ethereum 
$ cd $GOPATH/src/github.com/ethereum/go-ethereum
$ make abigen
```


## Usage
1. Compile Smart contract

```
make compile-masterchef
```

2. Generate Masterchef ABI
```
make abigen-masterchef
```

ref:
[https://goethereumbook.org/smart-contract-load/](https://goethereumbook.org/smart-contract-load/)
[https://goethereumbook.org/smart-contract-read-erc20/](https://goethereumbook.org/smart-contract-read-erc20/)