package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	contracts "github.com/gnasnik/fvm-contracts/build/v0.8"
	"github.com/gnasnik/fvm-contracts/client"
	"github.com/google/uuid"
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

	vpsContractAddress := common.HexToAddress("0x9652dBa7177F7ea8c3eB14EF0e5c1088f170aDd4")

	orderID := uuid.NewString()
	result, err := c.InvokeContract(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		instance, err := contracts.NewTitanVPS(vpsContractAddress, c.Client)
		if err != nil {
			return nil, err
		}

		buyerAddr := common.HexToAddress("0xe003B2Fb03F3126347afDBba460ED39e57F9588d")
		now := time.Now()
		price := big.NewInt(10000000000000000)
		createdAt := big.NewInt(now.Unix())
		expiration := big.NewInt(now.Add(time.Hour * 24 * 7).Unix())
		return instance.NewOrder(opts, orderID, buyerAddr, price, expiration, createdAt)
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("order id: ", orderID)
	log.Println("create order OK: ", string(result))
	log.Println("querying order...")

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-time.After(2 * time.Minute):
			log.Println("query order timeout!")
			return
		case <-ticker.C:
			instance, err := contracts.NewTitanVPS(vpsContractAddress, c.Client)
			if err != nil {
				log.Println(err)
				continue
			}
			order, err := instance.QueryOrder(&bind.CallOpts{
				Pending: true,
			}, orderID)
			if err != nil {
				continue
			}

			log.Println("Query order OK: ", order)
			return
		}
	}

}
