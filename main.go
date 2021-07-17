package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/atchapcyp/masterchef-go/pancake"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E")
	instance, err := pancake.NewPancake(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	path := make([]common.Address, 0)
	path = append(path, common.HexToAddress("0x55d398326f99059ff775485246999027b3197955"))
	path = append(path, common.HexToAddress("0x39f1014a88c8ec087cedf1bfc7064d24f507b894"))

	response, err := instance.GetAmountsOut(&bind.CallOpts{}, new(big.Int).SetUint64(100000000), path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response : %+v\n", response)

	//MASTER CHEF

	// contractAddress := common.HexToAddress("0x73feaa1eE314F8c655E354234017bE2193C9E24E")
	// instance, err := masterchef.NewMasterchef(contractAddress, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// devAddr, err := instance.Devaddr(&bind.CallOpts{})
	// if err != nil {

	// 	log.Fatal(err)
	// }

	// fmt.Printf("devAddr %+v\n", devAddr)

	// poolInfo, err := instance.PoolInfo(&bind.CallOpts{}, big.NewInt(100))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("poolinfo %+v\n", poolInfo)

	// // BEP20
	// bep20ContractAddress := common.HexToAddress("0x55d398326f99059fF775485246999027B3197955")
	// bep20contract, err := bep20.NewBep20(bep20ContractAddress, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// tokenName, err := bep20contract.Name(&bind.CallOpts{})
	// if err != nil {

	// 	log.Fatal(err)
	// }

	// fmt.Printf("tokenName : %+v\n", tokenName)

	// tts, err := bep20contract.TotalSupply(&bind.CallOpts{})
	// if err != nil {

	// 	log.Fatal(err)
	// }
	// fmt.Printf("totalSupply : %+v\n", tts)

}
