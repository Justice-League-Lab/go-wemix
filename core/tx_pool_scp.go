package core

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/script"
	"github.com/ethereum/go-ethereum/core/script/erc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

const (
	poolID       string  = "0x42Cf1Af7Fa9c2b50855A47806706D623De73316b"
	node         string  = "127.0.0.1:8545"
	nodeWebSite  string  = "https://api.wemix.com"
	myaddress    string  = "0xE8DB41c5e9ef0f09a1e65f8dC8e9fEF1879250A9"
	contract     string  = "0x80a5A916FB355A8758f0a3e47891dc288DAC2665"
	methodId1    string  = "0x06fd4ac5"
	methodId2    string  = "0x41876647"
	methodId     string  = "38ed1739"
	coin1        string  = "0x8e81fcc2d4a3baa0ee9044e0d7e36f59c9bba9c1"
	coin2        string  = "0x770d9d14c4ae2f78dca810958c1d9b7ea4620289"
	coinone32    string  = "0000000000000000000000008e81fcc2d4a3baa0ee9044e0d7e36f59c9bba9c1"
	cointwe32    string  = "000000000000000000000000770d9d14c4ae2f78dca810958c1d9b7ea4620289"
	priceDefault float64 = 0.75

	prikey string = "4e041f6f473bb6250db7688e2fba855b787708448840f271abadb3214944fec4"

	MaxGas int64 = 38694000460
)

var (
	once        sync.Once
	coreSPool   *script.CoreSession
	coreSERC20  *erc.CoreSession
	myAddress   common.Address = common.HexToAddress(myaddress)
	toAddress   common.Address = common.HexToAddress(contract)
	privateKey  *ecdsa.PrivateKey
	chainId     *big.Int
	client      *ethclient.Client
	nonceAtomic *atomic.Uint64
)

func DOTxScript(tx types.Transaction) {

	once.Do(func() {

		script.InitLog("./logs", "scirpt", "debug")

		var err error

		client, err = ethclient.Dial(nodeWebSite) // 本地节点的默认RPC端口
		if err != nil {
			logrus.Errorf("Dial client err : %v", err)
			return
		}

		privateKey, err = crypto.HexToECDSA(prikey)
		if err != nil {
			logrus.Errorf("hexToECDSA prikey err : %v", err)
			return
		}

		chainId = new(big.Int).SetInt64(1111)

		core, err := script.NewCore(common.HexToAddress(poolID), client)

		if err != nil {
			logrus.Errorf("NewCore  err : %v", err)
			return
		}

		callOpts := bind.CallOpts{
			Pending: true,                 // 是否查询待处理的区块
			From:    myAddress,            // 可选，指定调用者地址
			Context: context.Background(), // 可选，指定上下文
		}

		coreSPool = &script.CoreSession{
			Contract: core,
			CallOpts: callOpts,
		}

		coreERC20, err := erc.NewCore(toAddress, client)

		if err != nil {
			logrus.Errorf("erc NewCore  err : %v", err)
			return
		}
		coreSERC20 = &erc.CoreSession{
			Contract: coreERC20,
			CallOpts: callOpts,
		}

		nonce, err := client.PendingNonceAt(context.Background(), myAddress)
		if err != nil {
			logrus.Errorf("NonceAt  err : %v", err)
			return
		}

		nonceAtomic = &atomic.Uint64{}

		nonceAtomic.Store(nonce)

	})

	addr := tx.To()
	if !strings.EqualFold(strings.ToLower(addr.String()), strings.ToLower(contract)) {
		logrus.Infof("tx is invalid tx id %s", tx.Hash())
		return
	}

	txData := hex.EncodeToString(tx.Data())

	// txData := string(tx.Data())

	logrus.Infof("tx data: %s", txData)

	if txData[0:8] != methodId {
		logrus.Infof("tx methodId  err tx id %s", tx.Hash())
		return
	}

	coinData, err := coreSPool.GetReserves()
	if err != nil {
		logrus.Errorf("GetReserves  err : %v", err)
		return
	}

	logrus.Infof("crow pool balance is %v  , wemix pool balance is %v", coinData.Reserve0, coinData.Reserve1)

	coin := txData[456:520]

	if coin == cointwe32 {
		input, _ := new(big.Int).SetString(txData[9:72], 16)
		output, _ := new(big.Int).SetString(txData[73:136], 16)
		totalCoin1 := coinData.Reserve1.Add(coinData.Reserve1, input)
		totalCoin2 := coinData.Reserve0.Sub(coinData.Reserve0, output)

		ratA := new(big.Rat).SetInt(totalCoin1)
		ratB := new(big.Rat).SetInt(totalCoin2)

		result := new(big.Rat).Quo(ratA, ratB)

		// 获取结果字符串并格式化
		decimalPos := result.FloatString(2)

		price, err := strconv.ParseFloat(decimalPos, 64)
		if err != nil {
			logrus.Errorf("ParseFloat  err : %v", err)
			return
		}
		if price <= priceDefault {
			return
		}

		//TODO：send tx

		decimalPos2 := result.FloatString(5)

		priceCalc, err := strconv.ParseFloat(decimalPos2, 64)
		if err != nil {
			logrus.Errorf("ParseFloat  err : %v", err)
			return
		}

		amountIn := new(big.Int).Mul(new(big.Int).SetInt64(int64(((priceCalc/priceDefault - 1) * 200000))), big.NewInt(1e18))
		amountOut := new(big.Int).Div(new(big.Int).Mul(amountIn, big.NewInt(75)), big.NewInt(100))

		logrus.Infof("crow input is %v  , wemix output is %v", amountIn.String(), amountOut.String())

		txHash, err := SendTx(
			client,
			coreSERC20,
			tx,
			privateKey,
			coin2,
			coin1,
			chainId,
			amountIn,
			amountOut,
			new(big.Int).SetUint64(nonceAtomic.Add(1)))

		if err != nil {
			logrus.Errorf("SendTx  err : %v  tx hash is %v", err, txHash)
			return
		}

		logrus.Infof(" tx succuess hash is %v", txHash)

	}

	if coin == coinone32 {
		input, _ := new(big.Int).SetString(txData[9:72], 16)
		output, _ := new(big.Int).SetString(txData[73:136], 16)
		totalCoin1 := coinData.Reserve1.Sub(coinData.Reserve1, input)
		totalCoin2 := coinData.Reserve0.Add(coinData.Reserve0, output)

		ratA := new(big.Rat).SetInt(totalCoin1)
		ratB := new(big.Rat).SetInt(totalCoin2)

		result := new(big.Rat).Quo(ratA, ratB)

		// 获取结果字符串并格式化
		decimalPos := result.FloatString(2)

		price, err := strconv.ParseFloat(decimalPos, 64)
		if err != nil {
			logrus.Errorf("ParseFloat  err : %v", err)
			return
		}
		if price >= priceDefault {
			return
		}

		//TODO：send tx

		decimalPos2 := result.FloatString(5)

		priceCalc, err := strconv.ParseFloat(decimalPos2, 64)
		if err != nil {
			logrus.Errorf("ParseFloat  err : %v", err)
			return
		}

		amountIn := new(big.Int).Mul(new(big.Int).SetInt64(int64(((priceDefault/priceCalc - 1) * 200000))), big.NewInt(1e18))
		amountOut := new(big.Int).Div(new(big.Int).Mul(amountIn, big.NewInt(100)), big.NewInt(75))

		logrus.Infof("wemix input is %v  ,crow  output is %v", amountIn.String(), amountOut.String())

		txHash, err := SendTx(
			client,
			coreSERC20,
			tx,
			privateKey,
			coin1,
			coin2,
			chainId,
			amountIn,
			amountOut,
			new(big.Int).SetUint64(nonceAtomic.Add(1)),
		)
		if err != nil {
			logrus.Errorf("SendTx  err : %v  tx hash is %v", err, txHash)
			return
		}

		logrus.Infof(" tx succuess hash is %v", txHash)

	}

}

func SendTx(
	client *ethclient.Client,
	coreSERC20 *erc.CoreSession,
	tx types.Transaction,
	privateKey *ecdsa.PrivateKey,
	coin1, coin2 string,
	chainId *big.Int,
	amountIn *big.Int,
	amountOut *big.Int,
	nonce *big.Int) (string, error) {

	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		err := fmt.Errorf("NewKeyedTransactorWithChainID is err %v", err)
		return "", err
	}

	txOpts.From = myAddress
	txOpts.Nonce = nonce
	txOpts.GasLimit = 250000
	txOpts.GasFeeCap = tx.GasFeeCap()
	txOpts.GasTipCap = tx.GasTipCap()
	txOpts.GasPrice = tx.GasPrice()

	gas, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		err := fmt.Errorf("SuggestGasPrice is err %v", err)
		return "", err
	}
	if gas.Cmp(tx.GasPrice()) != 1 {
		txOpts.GasPrice = gas
	}

	coreSERC20.TransactOpts = *txOpts

	dealline, _ := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)

	tx1, err := coreSERC20.SwapExactTokensForTokens(
		amountIn,
		amountOut,
		[]common.Address{
			common.HexToAddress(coin1),
			common.HexToAddress(coin2),
		},
		myAddress,
		dealline)
	if err != nil {
		err := fmt.Errorf("SwapExactTokensForTokens is err %v", err)
		return tx.Hash().String(), err
	}

	return tx1.Hash().String(), nil
}
