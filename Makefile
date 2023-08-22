build: build_contracts

build_contracts: build_vps

build_vps:
	./bin/solc contracts/v0.8/TitanVPS.sol --output-dir ./build/v0.8 --overwrite --bin --hashes --opcodes --abi
	./bin/abigen --abi build/v0.8/TitanVPS.abi --pkg contracts --type TitanVPS --out build/v0.8/TitanVPS.go

install_solc_linux:
	wget https://binaries.soliditylang.org/linux-amd64/solc-linux-amd64-v0.8.15+commit.e14f2714
	mv solc-linux-amd64-v0.8.15+commit.e14f2714 bin/solc
	chmod +x bin/solc

install_solc_win:
	@echo "No Windows. Only Linux."

install_solc_mac:
	wget https://github.com/ethereum/solidity/releases/download/v0.8.15/solc-macos
	mv solc-macos bin/solc
	chmod +x bin/solc

install_abigen:
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	cp $(GOPATH)/bin/abigen bin/abigen
	chmod +x bin/abigen