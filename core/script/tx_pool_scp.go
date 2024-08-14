package script

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/script/erc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	poolID       string  = "0x42Cf1Af7Fa9c2b50855A47806706D623De73316b"
	node         string  = "127.0.0.1:8545"
	nodeWebSite  string  = "https://api.wemix.com"
	myaddress    string  = "0xE8DB41c5e9ef0f09a1e65f8dC8e9fEF1879250A9"
	contract     string  = "0x8E81fCc2d4A3bAa0eE9044E0D7E36F59C9BbA9c1"
	methodId1    string  = "0x06fd4ac5"
	methodId2    string  = "0x41876647"
	methodId     string  = "0x38ed1739"
	coin1        string  = "0000000000000000000000008e81fcc2d4a3baa0ee9044e0d7e36f59c9bba9c1"
	coin2        string  = "000000000000000000000000770d9d14c4ae2f78dca810958c1d9b7ea4620289"
	priceDefault float64 = 0.75

	prikey string = "0x4e041f6f473bb6250db7688e2fba855b787708448840f271abadb3214944fec4"

	MaxGas int64 = 38694000460
)

var (
	once       sync.Once
	coreSPool  *CoreSession
	coreSERC20 *erc.CoreSession
	myAddress  common.Address = common.HexToAddress("0xE8DB41c5e9ef0f09a1e65f8dC8e9fEF1879250A9")
	privateKey *ecdsa.PrivateKey
	chainId    *big.Int
)

func DOTxScript(tx *types.Transaction, pool *core.TxPool) {

	once.Do(func() {
		client, err := ethclient.Dial(nodeWebSite) // 本地节点的默认RPC端口
		if err != nil {
			panic(err)
		}

		privateKey, err = crypto.HexToECDSA(prikey)
		if err != nil {
			panic(err)
		}

		chainId = new(big.Int).SetInt64(1111)

		core, err := NewCore(common.HexToAddress(poolID), client)

		if err != nil {
			panic(err)
		}

		callOpts := bind.CallOpts{
			Pending: true,                 // 是否查询待处理的区块
			From:    myAddress,            // 可选，指定调用者地址
			Context: context.Background(), // 可选，指定上下文
		}

		coreSPool = &CoreSession{
			Contract: core,
			CallOpts: callOpts,
		}

		coreERC20, err := erc.NewCore(common.HexToAddress(contract), client)

		if err != nil {
			panic(err)
		}
		coreSERC20 = &erc.CoreSession{
			Contract: coreERC20,
			CallOpts: callOpts,
		}

	})

	addr := tx.To()
	if !strings.EqualFold(strings.ToLower(addr.String()), strings.ToLower(contract)) {
		return
	}

	txData := string(tx.Data())

	if txData[0:10] != methodId {
		return
	}

	coinData, err := coreSPool.GetReserves()
	if err != nil {
		panic(err)
	}

	coin := txData[458:522]

	if coin == coin2 {
		input, _ := new(big.Int).SetString(txData[11:74], 16)
		output, _ := new(big.Int).SetString(txData[75:138], 16)
		totalCoin1 := coinData.Reserve1.Add(coinData.Reserve1, input)
		totalCoin2 := coinData.Reserve0.Sub(coinData.Reserve0, output)

		ratA := new(big.Rat).SetInt(totalCoin1)
		ratB := new(big.Rat).SetInt(totalCoin2)

		result := new(big.Rat).Quo(ratA, ratB)

		// 获取结果字符串并格式化
		decimalPos := result.FloatString(2)

		price, err := strconv.ParseFloat(decimalPos, 64)
		if err != nil {
			panic(err)
		}
		if price <= priceDefault {
			return
		}

		//TODO：send tx

		// decimalPos2 := result.FloatString(5)

		// priceCalc, err := strconv.ParseFloat(decimalPos2, 64)

		txOpts := bind.TransactOpts{
			From:      myAddress,                                                          // 必填，指定转账者地址
			Nonce:     new(big.Int).SetUint64(pool.GetCurrentState().GetNonce(myAddress)), // 可选，指定转账序号
			GasLimit:  3000000,
			GasTipCap: tx.GasTipCap(), // 可选，指定转账金额
			GasFeeCap: tx.GasFeeCap(),
			GasPrice:  tx.GasPrice(), // 必填，指定转账 gas 价格
		}

		types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)

		coreSERC20.TransactOpts = txOpts
		// amount :=  ((priceCalc/priceDefault - 1) * 308888) * 1e18

		// coreSERC20.SwapExactTokensForTokens(new(big.Int).)

		// coreS.

	}

	if coin == coin1 {
		input, _ := new(big.Int).SetString(txData[11:74], 16)
		output, _ := new(big.Int).SetString(txData[75:138], 16)
		totalCoin1 := coinData.Reserve1.Sub(coinData.Reserve1, input)
		totalCoin2 := coinData.Reserve0.Add(coinData.Reserve0, output)

		ratA := new(big.Rat).SetInt(totalCoin1)
		ratB := new(big.Rat).SetInt(totalCoin2)

		result := new(big.Rat).Quo(ratA, ratB)

		// 获取结果字符串并格式化
		decimalPos := result.FloatString(2)

		price, err := strconv.ParseFloat(decimalPos, 64)
		if err != nil {
			panic(err)
		}
		if price >= priceDefault {
			return
		}

		//TODO：send tx

		// amount := (priceDefault/price - 1) * 308888

	}

}
