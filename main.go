package main

import (
	"fmt"
	"log"
	"masterchef-go/bep20"
	"masterchef-go/masterchef"
	"masterchef-go/pancake"
	pancakepair "masterchef-go/pancake-pair"

	"math/big"

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
	moonMasterContractAddress := common.HexToAddress("0xbe739A112eF6278cEb374Bcad977252Bc3918cA9")
	moonInstance, err := masterchef.NewMasterchef(moonMasterContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	for i := int64(0); i < 9; i++ {
		poolInfo, err := moonInstance.PoolInfo(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("moon poolinfo %+v\n", poolInfo)
	}

	chefFoodCourtAddress := common.HexToAddress("0xe43b7c5c4c2df51306cceb7cbc4b2fcc038874f1")
	foodCourtInstance, err := masterchef.NewMasterchef(chefFoodCourtAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	for i := int64(0); i < 25; i++ {
		poolInfo, err := foodCourtInstance.PoolInfo(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("food court poolinfo %+v\n", poolInfo)
	}

	for i := int64(0); i < 25; i++ {
		userAddress := common.HexToAddress("0x330d2c6C22E82fd613830B3f9a97728dFC9B32FF")
		userInfo, err := foodCourtInstance.UserInfo(&bind.CallOpts{}, big.NewInt(i), userAddress)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("food court userinfo %+v\n", userInfo)
	}

	r := checkSCZLP(0.18, "0x330d2c6C22E82fd613830B3f9a97728dFC9B32FF", "0xb3d2C0cb104CBfA2167Af0D82A12475946B22386")
	fmt.Printf("scz lp %+v\n", r)

	// // BEP20
	bep20ContractAddress := common.HexToAddress("0x39f1014a88c8ec087cedf1bfc7064d24f507b894")
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

	userAddress := common.HexToAddress("0x330d2c6C22E82fd613830B3f9a97728dFC9B32FF")
	bal, err := bep20contract.BalanceOf(&bind.CallOpts{}, userAddress)
	if err != nil {

		log.Fatal(err)
	}
	fmt.Printf("scz bal : %+v\n", bal)

}

const SCZ_BUSD = "0xb3d2C0cb104CBfA2167Af0D82A12475946B22386"
const SCZ_DBM = "0xF31455AE22DfE3637b87D62371F6fc945F9ea301"
const chefMoon = "0xbe739a112ef6278ceb374bcad977252bc3918ca9"

const SCZ_BNB = "0xad468bcbcb33037e061ed5ec905ddea06cafc67f"
const chefFoodCourt = "0xe43b7c5c4c2df51306cceb7cbc4b2fcc038874f1"

func checkSCZLP(rate float64, userAddress, lpTokenAddress string) float64 {
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(lpTokenAddress)
	pairInstance, err := pancakepair.NewPancakepair(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	ua := common.HexToAddress(userAddress)
	_balance, err := pairInstance.BalanceOf(&bind.CallOpts{}, ua)
	if err != nil {
		log.Fatal(err)
	}
	reserve, err := pairInstance.GetReserves(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	_totalSupply, err := pairInstance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	token0, err := pairInstance.Token0(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	const SCZ = "0x39f1014a88c8ec087cedf1bfc7064d24f507b894"
	result := big.NewInt(0)
	if token0 == common.HexToAddress(SCZ) {
		return float64(result.Mul(_balance, reserve.Reserve0).Div(result, _totalSupply).Int64()) * rate
	}
	return float64(result.Mul(_balance, reserve.Reserve1).Div(result, _totalSupply).Int64()) * rate
}
