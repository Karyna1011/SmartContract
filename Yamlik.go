package main

import (
	"awesomeProject/config"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"io"
	"strings"
	"time")



const erc20ABI = "[{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

func CallFunction() {
	cfg := config.NewConfig(kv.MustFromEnv())
	eth := cfg.EthClient()

	log := logan.New()
	log = cfg.Log()

	parsed, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		log.Fatal("failed to parse contract ABI")
	}

	address := common.HexToAddress("0x118b69e0BE87a87BB30e093F496b1eE989aA15E4")

	var Contract = bind.NewBoundContract(
		address,
		parsed,
		eth,
		eth,
		eth,
	)

	result := make([]interface{}, len(""))
	for i, e := range "" {
		result[i] = e
	}

	err = Contract.Call(&bind.CallOpts{}, &result, "get")
	if err != nil {
		log.WithError(err).Error("error during calling contract")
		return
	}

	fmt.Println("RESULT:", result)
}

func main() {
	d := time.NewTicker(3 * time.Second)

	for {
		select {
		case tm := <-d.C:
			CallFunction()
			fmt.Println("The Current time is: ", tm)
			d.Reset(3 * time.Second)

		}

	}
}
