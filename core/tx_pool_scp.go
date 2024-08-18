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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/script"
	"github.com/ethereum/go-ethereum/core/script/coin"
	"github.com/ethereum/go-ethereum/core/script/config"
	"github.com/ethereum/go-ethereum/core/script/erc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

const (
	poolID       string  = "0x42Cf1Af7Fa9c2b50855A47806706D623De73316b"
	node         string  = "http://127.0.0.1:8588"
	nodeWebSite  string  = "wss://ws.wemix.com"
	myaddress    string  = "0x26ea8cd8b613b5eab41682da649e0df39dbaa025"
	contract     string  = "0x80a5A916FB355A8758f0a3e47891dc288DAC2665"
	methodId1    string  = "0x06fd4ac5"
	methodId2    string  = "0x41876647"
	methodId     string  = "38ed1739"
	coin1        string  = "0x8e81fcc2d4a3baa0ee9044e0d7e36f59c9bba9c1"
	coin2        string  = "0x770d9d14c4ae2f78dca810958c1d9b7ea4620289"
	coinone32    string  = "0000000000000000000000008e81fcc2d4a3baa0ee9044e0d7e36f59c9bba9c1"
	cointwe32    string  = "000000000000000000000000770d9d14c4ae2f78dca810958c1d9b7ea4620289"
	priceDefault float64 = 0.745

	MaxGas int64 = 38694000460
)

var (
	once          sync.Once
	coreSPool     *script.CoreSession
	coreSERC20    *erc.CoreSession
	coin1Contract *coin.CoinSession
	coin2Contract *coin.SciptSession
	myAddress     common.Address = common.HexToAddress(myaddress)
	toAddress     common.Address = common.HexToAddress(contract)
	privateKey    *ecdsa.PrivateKey
	chainId       *big.Int
	client        *ethclient.Client
	amountMin     *big.Int
	prikey        string
	// nonceAtomic *atomic.Uint64

	mapAddr []common.Address = []common.Address{

		common.HexToAddress(myaddress),
		common.HexToAddress("0xCd51c15e940a9feB43551C4b8C5c5c0498310137"),
		common.HexToAddress("0xBee95FD1c50099a8FfF5204EfD53C77900ab5052"),
		common.HexToAddress("0x1424e1be1b2299abFd10A7B8DE07CD4810a51B4A"),
		common.HexToAddress("0x12FAe2373c76C1dDB4eA87f852c2E06983fb315b"),
		common.HexToAddress("0xDcD92Aa378Efa9394D8E7bCa8714dedeb37f9dd9"),
		common.HexToAddress("0xfe1E89d8d31717d6a95A19e3ef5dAFf125992e5E"),
		common.HexToAddress("0x878cE6B0e10E05fA77e33bC429e00414e19c408F"),
	}
)

func DOTxScript(tx types.Transaction) {

	once.Do(func() {

		script.InitLog("./logs", "scirpt", "debug")

		cfg := &config.Config{}

		var err error

		err = config.ResolveConfig("../config.yaml", cfg)
		if err != nil {
			logrus.Errorf("ResolveConfig err : %v", err)
			return
		}
		prikey = cfg.PrivateKey

		client, err = ethclient.Dial(node) // 本地节点的默认RPC端口
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
		amountMin = new(big.Int).Mul(new(big.Int).SetInt64(200), big.NewInt(1e18))

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

		coin1C, err := coin.NewCoin(common.HexToAddress(coin1), client)

		if err != nil {
			logrus.Errorf(" coin1C NewCore  err : %v", err)
			return
		}
		coin1Contract = &coin.CoinSession{
			Contract: coin1C,
			CallOpts: callOpts,
		}

		coin2C, err := coin.NewScipt(common.HexToAddress(coin2), client)

		if err != nil {
			logrus.Errorf(" coin2C NewCore  err : %v", err)
			return
		}
		coin2Contract = &coin.SciptSession{
			Contract: coin2C,
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

		// nonceAtomic = &atomic.Uint64{}

		// nonceAtomic.Store(nonce)

	})
	// from := tx.

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

	nonce, err := client.PendingNonceAt(context.Background(), myAddress)
	if err != nil {
		logrus.Errorf("NonceAt  err : %v", err)
		return
	}

	if coin == cointwe32 {
		input, _ := new(big.Int).SetString(txData[9:72], 16)
		output, _ := new(big.Int).SetString(txData[73:136], 16)
		totalCoin1 := coinData.Reserve1.Add(coinData.Reserve1, input)
		totalCoin2 := coinData.Reserve0.Sub(coinData.Reserve0, output)

		ratA := new(big.Rat).SetInt(totalCoin1)
		ratB := new(big.Rat).SetInt(totalCoin2)

		result := new(big.Rat).Quo(ratA, ratB)

		// 获取结果字符串并格式化
		decimalPos := result.FloatString(3)

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
		if amountMin.Cmp(amountIn) == 1 {
			logrus.Infof("amout not less than 200   tx hash is %v", tx.Hash().String())
			return
		}
		balance, err := coin2Contract.BalanceOf(myAddress)
		if err != nil {
			logrus.Infof("BalanceOf get err: %v tx hash is %v", err, tx.Hash().String())
			return
		}

		if balance.Cmp(amountIn) < 1 {
			amountIn = balance
		}

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
			new(big.Int).SetUint64(nonce))

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
		decimalPos := result.FloatString(3)

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
		if amountMin.Cmp(amountIn) == 1 {
			logrus.Infof("amout not less than 200  tx hash is %v", tx.Hash().String())
			return
		}
		balance, err := coin1Contract.BalanceOf(myAddress)
		if err != nil {
			logrus.Infof("BalanceOf get err: %v tx hash is %v", err, tx.Hash().String())
			return
		}

		if balance.Cmp(amountIn) < 1 {
			amountIn = balance
		}

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
			new(big.Int).SetUint64(nonce),
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
	txOpts.GasLimit = 300000
	// txOpts.GasFeeCap = tx.GasFeeCap()
	// txOpts.GasTipCap = tx.GasTipCap()
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

func FilterAddress(addr common.Address) bool {
	for _, v := range mapAddr {
		if strings.EqualFold(v.String(), addr.String()) {
			return false
		}
	}

	return true
}
