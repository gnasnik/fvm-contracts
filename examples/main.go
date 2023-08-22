package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	contracts "github.com/gnasnik/fvm-contracts/build/v0.8"
	"github.com/gnasnik/fvm-contracts/client"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {
	c, err := client.New(
		client.PrivateKeyOption(os.Getenv("PRIVATE_KEY")),
	)
	if err != nil {
		log.Fatal(err)
	}

	result, err := c.InvokeContract(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		vpsContractAddress := common.HexToAddress("0xfddde2D6A3Ae34f3d4b9a19f49eA963A20Eb8527")
		instance, err := contracts.NewTitanVPS(vpsContractAddress, c.Client)
		if err != nil {
			return nil, err
		}

		buyerAddr := common.HexToAddress("0xe003B2Fb03F3126347afDBba460ED39e57F9588d")
		now := time.Now()
		price := big.NewInt(10000000000000000)
		createdAt := big.NewInt(now.Unix())
		expiration := big.NewInt(now.Add(time.Hour * 24 * 7).Unix())
		return instance.NewOrder(opts, "test2", buyerAddr, price, expiration, createdAt)
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("OK: ", string(result))
}
