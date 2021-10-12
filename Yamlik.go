package main

import (
	"awesomeProject/config"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/go/xdrbuild"
	"io"

	//"log"
	"strings"
	"time"

	/*"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient" */

)
type Transformer interface {
	Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error)
	Reset()
}

/*type SpanningTransformer interface {
	Transformer
	Span(src []byte, atEOF bool) (n int, err error)
}*/

type NopResetter struct{}

// Reset implements the Reset method of the Transformer interface.
func (NopResetter) Reset() {}

// Reader wraps another io.Reader by transforming the bytes read.
type Reader struct {
	r   io.Reader
	t   Transformer
	err error
	dst        []byte
	dst0, dst1 int
	src        []byte
	src0, src1 int
	transformComplete bool
}

const defaultBufSize = 4096
func NewReader(r io.Reader, t Transformer) *Reader {
	t.Reset()
	return &Reader{
		r:   r,
		t:   t,
		dst: make([]byte, defaultBufSize),
		src: make([]byte, defaultBufSize),
	}
}


type Opts struct {
	Client ethclient.Client
	Builder   xdrbuild.Builder
	Config    config.Config

}
const erc20ABI ="[{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	func MachenEtwas(){



	cfg := config.NewConfig(kv.MustFromEnv())
	//eth:=cfg.EthClient()


	log:=logan.New()
	log= cfg.Log()


	var opts Opts
	parsed, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		log.Fatal("failed to parse contract ABI")
	}
	address:=common.HexToAddress("0x118b69e0BE87a87BB30e093F496b1eE989aA15E4")

	var Contract = bind.NewBoundContract(
		address,
		parsed,
		&opts.Client,
		&opts.Client,
		&opts.Client,
		)

	vs := make([]interface{}, len(""))
	for i, e := range "" {
		vs[i] = e
	}



		fmt.Println(Contract.Call(&bind.CallOpts{},&vs,"get"))

}

func main() {
	d := time.NewTicker(3 * time.Second)
	MyChannel := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second)
		MyChannel <- true
	}()

	for {
		select {
		case <-MyChannel:
			fmt.Println("Completed!")
			return
		case tm := <-d.C:
			//MakeTransaction()
			MachenEtwas()
			fmt.Println("The Current time is: ", tm)
		}
	}

}
