# If you see pwd_unknown showing up, this is why. Re-calibrate your system.
PWD ?= pwd_unknown

compile-masterchef: 
	docker run --rm -v $(PWD):/root ethereum/solc:0.6.12 --abi --bin /root/contracts/MasterChef.sol -o /root/build

compile-twindex-swap-router: 
	docker run --rm -v $(PWD):/root ethereum/solc:0.6.12 --abi --bin /root/contracts/TwindexSwapRouter.sol -o /root/build-twindex

compile-pancake-swap-router: 
	docker run --rm -v $(PWD):/root ethereum/solc:0.6.6 --abi --bin /root/contracts/PancakeRouter.sol -o /root/build-pancake-router

abigen-masterchef:
	mkdir -p masterchef && abigen --bin=./build/MasterChef.bin --abi=./build/MasterChef.abi --pkg=masterchef --out=./masterchef/master-chef.go

abigen-bep20:
	mkdir -p bep20 && abigen --bin=./build/BEP20.bin --abi=./build/BEP20.abi --pkg=bep20 --out=./bep20/bep20.go

abigen-chainlink:
	mkdir -p chainlink && abigen --bin=./build-chainlink/EACAggregatorProxy.bin --abi=./build-chainlink/EACAggregatorProxy.abi --pkg=chainlink --out=./chainlink/chainlink.go

abigen-twindex-swap-router:
	mkdir -p twindex && abigen --bin=./build-twindex/TwindexSwapRouter.bin --abi=./build-twindex/TwindexSwapRouter.abi --pkg=twindex --out=./twindex/twindex.go

abigen-pancake-swap-router:
	mkdir -p pancake && abigen --bin=./build-pancake-router/PancakeRouter.bin --abi=./build-pancake-router/PancakeRouter.abi --pkg=pancake --out=./pancake/pancake-router.go

abigen-pancake-pair:
	mkdir -p pancake-pair && abigen --bin=./build-pancake-router/IPancakePair.bin --abi=./build-pancake-router/IPancakePair.abi --pkg=pancakepair --out=./pancake-pair/pancake-pair.go