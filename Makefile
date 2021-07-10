# If you see pwd_unknown showing up, this is why. Re-calibrate your system.
PWD ?= pwd_unknown

compile-masterchef: 
	docker run --rm -v $(PWD):/root ethereum/solc:0.6.12 --abi --bin /root/contracts/MasterChef.sol -o /root/build

abigen-masterchef:
	mkdir -p masterchef && abigen --bin=./build/MasterChef.bin --abi=./build/MasterChef.abi --pkg=masterchef --out=./masterchef/master-chef.go

abigen-bep20:
	mkdir -p bep20 && abigen --bin=./build/BEP20.bin --abi=./build/BEP20.abi --pkg=bep20 --out=./bep20/bep20.go