package main

import (
	"fmt"
	"log"
	"math/big"

	bep20 "github.com/atchapcyp/masterchef-go/bep20"
	masterchef "github.com/atchapcyp/masterchef-go/masterchef"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		log.Fatal(err)
	}

	//MASTER CHEF

	contractAddress := common.HexToAddress("0x73feaa1eE314F8c655E354234017bE2193C9E24E")
	instance, err := masterchef.NewMasterchef(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	devAddr, err := instance.Devaddr(&bind.CallOpts{})
	if err != nil {

		log.Fatal(err)
	}

	fmt.Printf("devAddr %+v\n", devAddr)

	poolInfo, err := instance.PoolInfo(&bind.CallOpts{}, big.NewInt(100))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("poolinfo %+v\n", poolInfo)

	// BEP20
	bep20ContractAddress := common.HexToAddress("0x55d398326f99059fF775485246999027B3197955")
	bep20contract, err := bep20.NewBep20(bep20ContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tokenName, err := bep20contract.Name(&bind.CallOpts{})
	if err != nil {

		log.Fatal(err)
	}

	fmt.Printf("tokenName : %+v\n", tokenName)

	tts, err := bep20contract.TotalSupply(&bind.CallOpts{})
	if err != nil {

		log.Fatal(err)
	}
	fmt.Printf("totalSupply : %+v\n", tts)

}
