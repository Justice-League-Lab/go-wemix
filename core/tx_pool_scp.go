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
	channel "github.com/ethereum/go-ethereum/chan"
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

type OptType uint8

const (
	BuyType OptType = iota + 1
	SellType
)

const (
	poolID      string = "0x42Cf1Af7Fa9c2b50855A47806706D623De73316b"
	node        string = "http://127.0.0.1:8588"
	nodeWebSite string = "wss://ws.wemix.com"
	nodeHttp    string = "https://api.wemix.com"
	myaddress   string = "0x26ea8cd8b613b5eab41682da649e0df39dbaa025"
	contract    string = "0x80a5A916FB355A8758f0a3e47891dc288DAC2665"
	methodId    string = "38ed1739"
	methodId1   string = "06fd4ac5"
	methodId2   string = "41876647"
	methodId3   string = "09c5eabe"
	methodId4   string = "d97495c9"
	methodId5   string = "592db2b9"
	methodId6   string = "baa2abde"

	coin1           string  = "0x8e81fcc2d4a3baa0ee9044e0d7e36f59c9bba9c1"
	coin2           string  = "0x770d9d14c4ae2f78dca810958c1d9b7ea4620289"
	coinCmpSell     string  = `000000000000000000000000770d9d14c4ae2f78dca810958c1d9b7ea46202890000000000000000000000008e81fcc2d4a3baa0ee9044e0d7e36f59c9bba9c1`
	coinCmpBuy      string  = `0000000000000000000000008e81fcc2d4a3baa0ee9044e0d7e36f59c9bba9c1000000000000000000000000770d9d14c4ae2f78dca810958c1d9b7ea4620289`
	coinone32       string  = "0000000000000000000000008e81fcc2d4a3baa0ee9044e0d7e36f59c9bba9c1"
	cointwe32       string  = "000000000000000000000000770d9d14c4ae2f78dca810958c1d9b7ea4620289"
	priceDefaultDel float64 = 0.745
	priceDefaultAdd float64 = 0.75

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

	contractList []common.Address = []common.Address{
		common.HexToAddress(contract),
		common.HexToAddress("0x20fdb3f41371ec0834a11dafcdb9acf5157236c4"),
		common.HexToAddress("0x398D227685614Aaeb2e4711b048626b0551Bc0Ee"),
	}

	methodIdList []string = []string{
		methodId,
		methodId1,
		methodId2,
		methodId3,
		methodId4,
		methodId5,
		methodId6,
	}

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

func DOTxScript(tx types.Transaction, pool *TxPool, optType string) {

	defer func() {
		if e := recover(); e != nil {
			logrus.Errorf("painc err : %v", e)
		}
	}()

	once.Do(func() {

		script.InitLog("./logs", "scirpt", "debug")

		cfg := &config.Config{}

		var err error

		err = config.ResolveConfig("/opt/gwemix/bin/config.yaml", cfg)
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

		channel.Init()

		// nonceAtomic = &atomic.Uint64{}

		// nonceAtomic.Store(nonce)

	})
	// from := tx.

	addr := tx.To()

	for i, v := range contractList {
		if strings.EqualFold(addr.String(), v.String()) {
			break
		}
		if i == len(contractList)-1 {
			logrus.Infof("tx is invalid tx id %s, type is %v", tx.Hash(), optType)
			return
		}

	}

	txData := hex.EncodeToString(tx.Data())

	// txData := string(tx.Data())

	// logrus.Infof("tx data: %s", txData)

	for i, v := range methodIdList {
		if txData[0:8] == v {
			break
		}
		if i == len(methodIdList)-1 {
			logrus.Infof("method id not find tx id %s", tx.Hash())
			return
		}
	}

	coinData, err := coreSPool.GetReserves()
	if err != nil {
		logrus.Errorf("GetReserves  err : %v", err)
		return
	}
	nonce, err := client.PendingNonceAt(context.Background(), myAddress)
	if err != nil {
		logrus.Errorf("NonceAt  err : %v", err)
		return
	}
	// logrus.Infof("crow pool balance is %v  , wemix pool balance is %v", coinData.Reserve0, coinData.Reserve1)

	{
		Do0x06fd4ac5(txData, nonce, coinData.Reserve0, coinData.Reserve1, tx, coreSERC20, pool)
		Do0x09c5eabe(txData, nonce, coinData.Reserve0, coinData.Reserve1, tx, coreSERC20, pool)
		Do0x38ed1739(txData, nonce, coinData.Reserve0, coinData.Reserve1, tx, coreSERC20, pool)
		Do0x41876647(txData, nonce, coinData.Reserve0, coinData.Reserve1, tx, coreSERC20, pool)
		Do0x592db2b9(txData, nonce, coinData.Reserve0, coinData.Reserve1, tx, coreSERC20, pool)
		Do0xbaa2abde(txData, nonce, coinData.Reserve0, coinData.Reserve1, tx, coreSERC20, pool)
		Do0xd97495c9(txData, nonce, coinData.Reserve0, coinData.Reserve1, tx, coreSERC20, pool)
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
	nonce *big.Int,
	pool *TxPool) (*types.Transaction, error) {

	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		err := fmt.Errorf("NewKeyedTransactorWithChainID is err %v", err)
		return nil, err
	}

	txOpts.From = myAddress
	txOpts.Nonce = nonce
	txOpts.GasLimit = 330000

	txOpts.GasFeeCap = tx.GasFeeCap()
	txOpts.GasTipCap = tx.GasTipCap()
	// txOpts.GasPrice = tx.GasPrice()

	// gas, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	err := fmt.Errorf("SuggestGasPrice is err %v", err)
	// 	return "", err
	// }

	// if gas.Cmp(tx.GasPrice()) != 1 {
	// 	txOpts.GasPrice = gas
	// }

	coreSERC20.TransactOpts = *txOpts

	dealline, _ := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)

	txNew, err := coreSERC20.SwapExactTokensForTokens(
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
		return nil, err
	}

	pool.AddLocal(txNew)

	channel.Txchan <- struct{}{}

	return txNew, nil
}

func FilterAddress(addr common.Address) bool {
	for _, v := range mapAddr {
		if strings.EqualFold(v.String(), addr.String()) {
			return false
		}
	}

	return true
}

func dealwithAmout(dividend, divisor float64, tx types.Transaction, optType OptType) (*big.Int, *big.Int) {

	var balance, amountIn, amountOut *big.Int
	if optType == SellType {
		amountIn = new(big.Int).Mul(new(big.Int).SetInt64(int64(((dividend/divisor - 1) * 240000))), big.NewInt(1e18))

		var err error
		balance, err = coin2Contract.BalanceOf(myAddress)
		if err != nil {
			logrus.Errorf("BalanceOf get err: %v tx hash is %v", err, tx.Hash().String())
			return nil, nil
		}
		if balance.Cmp(amountIn) < 1 {
			amountIn = balance
		}

		if amountMin.Cmp(amountIn) == 1 {
			logrus.Infof("amout not less than 200  tx hash is %v", tx.Hash().String())
			return nil, nil
		}

		amountOut = new(big.Int).Div(new(big.Int).Mul(amountIn, big.NewInt(74812)), big.NewInt(100000))

	}

	if optType == BuyType {
		amountOut = new(big.Int).Mul(new(big.Int).SetInt64(int64(((dividend/divisor - 1) * 240000))), big.NewInt(1e18))

		if amountMin.Cmp(amountOut) == 1 {
			logrus.Infof("amout not less than 200  tx hash is %v", tx.Hash().String())
			return nil, nil
		}
		amountIn = new(big.Int).Div(new(big.Int).Mul(amountOut, big.NewInt(74812)), big.NewInt(100000))
		var err error
		balance, err = coin1Contract.BalanceOf(myAddress)
		if err != nil {
			logrus.Errorf("BalanceOf get err: %v tx hash is %v", err, tx.Hash().String())
			return nil, nil
		}
		if balance.Cmp(amountIn) < 1 {
			amountIn = balance
			if amountMin.Cmp(amountIn) == 1 {
				logrus.Infof("amout not less than 200  tx hash is %v", tx.Hash().String())
				return nil, nil
			}
			amountOut = new(big.Int).Div(new(big.Int).Mul(amountIn, big.NewInt(100000)), big.NewInt(74812))
		}

	}

	return amountIn, amountOut

}

func dealWithcoinprice(totalCoin1, totalCoin2 *big.Int) (*big.Rat, bool) {

	ratA := new(big.Rat).SetInt(totalCoin1)
	ratB := new(big.Rat).SetInt(totalCoin2)

	result := new(big.Rat).Quo(ratA, ratB)

	// 获取结果字符串并格式化
	decimalPos := result.FloatString(3)

	price, err := strconv.ParseFloat(decimalPos, 64)
	if err != nil {
		logrus.Errorf("ParseFloat  err : %v", err)
		return result, false
	}

	if price > priceDefaultAdd || price < priceDefaultDel {
		return result, true
	}

	return result, false
}

func Do0x06fd4ac5(
	txData string,
	nonce uint64,
	reserve0 *big.Int,
	reserve1 *big.Int,
	tx types.Transaction,
	coreSERC20 *erc.CoreSession,
	pool *TxPool) {
	if txData[:8] != methodId1 {
		return
	}
	output, _ := new(big.Int).SetString(txData[9:72], 16)
	input := new(big.Int).Div(new(big.Int).Mul(output, new(big.Int).SetInt64(75)), new(big.Int).SetInt64(100))

	totalCoin1 := new(big.Int).Add(reserve1, input)
	totalCoin2 := new(big.Int).Sub(reserve0, output)

	result, ok := dealWithcoinprice(totalCoin1, totalCoin2)

	if !ok {
		return
	}

	decimalPos2 := result.FloatString(5)

	priceCalc, err := strconv.ParseFloat(decimalPos2, 64)
	if err != nil {
		logrus.Errorf("ParseFloat  err : %v", err)
		return
	}
	amountIn, amountOut := dealwithAmout(priceCalc, priceDefaultAdd, tx, SellType)
	if amountIn == nil || amountOut == nil {
		return
	}

	logrus.Infof("crow input is %v  , wemix output is %v", amountIn.String(), amountOut.String())

	txNew, err := SendTx(
		client,
		coreSERC20,
		tx,
		privateKey,
		coin2,
		coin1,
		chainId,
		amountIn,
		amountOut,
		new(big.Int).SetUint64(nonce),
		pool,
	)
	if err != nil {
		logrus.Errorf("SendTx  err : %v  tx hash is %v", err, txNew.Hash())
		return
	}

	logrus.Infof(" tx succuess hash is %v", txNew.Hash())

}

func Do0x41876647(
	txData string,
	nonce uint64,
	reserve0 *big.Int,
	reserve1 *big.Int,
	tx types.Transaction,
	coreSERC20 *erc.CoreSession,
	pool *TxPool) {
	if txData[:8] != methodId2 {
		return
	}
	var v1, v2 string
	var optType OptType

	v1 = txData[9:72]
	v2 = txData[73:136]

	if strings.Contains(txData, coinCmpBuy) {
		optType = BuyType
	}
	if strings.Contains(txData, coinCmpSell) {
		optType = SellType
	}
	if optType == BuyType {
		dealWithBuyData(v1, v2, nonce, reserve0, reserve1, tx, pool)
	}

	if optType == SellType {
		dealWithSellData(v1, v2, nonce, reserve0, reserve1, tx, pool)
	}

}

func Do0x38ed1739(
	txData string,
	nonce uint64,
	reserve0 *big.Int,
	reserve1 *big.Int,
	tx types.Transaction,
	coreSERC20 *erc.CoreSession,
	pool *TxPool) {
	if txData[:8] != methodId {
		return
	}
	var v1, v2 string
	var optType OptType

	v1 = txData[9:72]
	v2 = txData[73:136]

	if strings.Contains(txData, coinCmpBuy) {
		optType = BuyType
	}
	if strings.Contains(txData, coinCmpSell) {
		optType = SellType
	}
	if optType == BuyType {
		dealWithBuyData(v1, v2, nonce, reserve0, reserve1, tx, pool)
	}

	if optType == SellType {
		dealWithSellData(v1, v2, nonce, reserve0, reserve1, tx, pool)
	}

}

func Do0x09c5eabe(
	txData string,
	nonce uint64,
	reserve0 *big.Int,
	reserve1 *big.Int,
	tx types.Transaction,
	coreSERC20 *erc.CoreSession,
	pool *TxPool) {
	if txData[:8] != methodId3 {
		return
	}
	var v1, v2 string
	var optType OptType

	switch len(txData) {
	case 2056:
		v1 = txData[657 : 657+63]
		v2 = txData[721 : 721+63]
	case 2184:
		v1 = txData[1745 : 1745+63]
		v2 = txData[1809 : 1809+63]
	case 1992:
		v1 = txData[657 : 657+63]
		v2 = txData[721 : 721+63]
	case 2440:
		v1 = txData[657 : 657+63]
		v2 = txData[721 : 721+63]
	case 2120:
		v1 = txData[1681 : 1681+63]
		v2 = txData[1745 : 1745+63]
	case 2568:
		v1 = txData[1681 : 1681+63]
		v2 = txData[1745 : 1745+63]
	default:
		return
	}

	if strings.Contains(txData, coinCmpBuy) {
		optType = BuyType
	}
	if strings.Contains(txData, coinCmpSell) {
		optType = SellType
	}
	if optType == BuyType {
		dealWithBuyData(v1, v2, nonce, reserve0, reserve1, tx, pool)
	}

	if optType == SellType {
		dealWithSellData(v1, v2, nonce, reserve0, reserve1, tx, pool)
	}
}

func Do0xd97495c9(
	txData string,
	nonce uint64,
	reserve0 *big.Int,
	reserve1 *big.Int,
	tx types.Transaction,
	coreSERC20 *erc.CoreSession,
	pool *TxPool) {
	if txData[:8] != methodId4 {
		return
	}
	var v1, v2 string
	var optType OptType

	switch len(txData) {
	case 1416:
		v1 = txData[713 : 713+63]
		v2 = txData[777 : 777+63]
	case 1352:
		v1 = txData[713 : 713+63]
		v2 = txData[777 : 777+63]
	default:
		return
	}

	if strings.Contains(txData, coinCmpBuy) {
		optType = BuyType
	}
	if strings.Contains(txData, coinCmpSell) {
		optType = SellType
	}
	if optType == BuyType {
		dealWithBuyData(v1, v2, nonce, reserve0, reserve1, tx, pool)
	}

	if optType == SellType {
		dealWithSellData(v1, v2, nonce, reserve0, reserve1, tx, pool)
	}
}

func Do0x592db2b9(
	txData string,
	nonce uint64,
	reserve0 *big.Int,
	reserve1 *big.Int,
	tx types.Transaction,
	coreSERC20 *erc.CoreSession,
	pool *TxPool) {
	if txData[:8] != methodId5 {
		return
	}
	var v1, v2 string
	var optType OptType

	switch len(txData) {
	case 1416:
		v1 = txData[713 : 713+63]
		v2 = txData[777 : 777+63]
	case 1352:
		v1 = txData[713 : 713+63]
		v2 = txData[777 : 777+63]
	default:
		return
	}

	if strings.Contains(txData, coinCmpBuy) {
		optType = BuyType
	}
	if strings.Contains(txData, coinCmpSell) {
		optType = SellType
	}
	if optType == BuyType {
		dealWithBuyData(v1, v2, nonce, reserve0, reserve1, tx, pool)
	}

	if optType == SellType {
		dealWithSellData(v1, v2, nonce, reserve0, reserve1, tx, pool)
	}
}

func Do0xbaa2abde(
	txData string,
	nonce uint64,
	reserve0 *big.Int,
	reserve1 *big.Int,
	tx types.Transaction,
	coreSERC20 *erc.CoreSession,
	pool *TxPool) {
	if txData[:8] != methodId6 {
		return
	}
	input, _ := new(big.Int).SetString(txData[137:137+63], 16)
	output := new(big.Int).Div(new(big.Int).Mul(input, new(big.Int).SetInt64(75)), new(big.Int).SetInt64(100))
	totalCoin1 := reserve1.Sub(reserve1, input)
	totalCoin2 := reserve1.Add(reserve0, output)

	result, ok := dealWithcoinprice(totalCoin1, totalCoin2)

	if !ok {
		return
	}

	//TODO：send tx

	decimalPos2 := result.FloatString(5)

	priceCalc, err := strconv.ParseFloat(decimalPos2, 64)
	if err != nil {
		logrus.Errorf("ParseFloat  err : %v", err)
		return
	}

	amountIn, amountOut := dealwithAmout(priceDefaultDel, priceCalc, tx, BuyType)
	if amountIn == nil || amountOut == nil {
		return
	}

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
		pool,
	)
	if err != nil {
		logrus.Errorf("SendTx  err : %v  tx hash is %v", err, txHash.Hash())
		return
	}

	logrus.Infof(" tx succuess hash is %v", txHash.Hash())

	return
}

func dealWithSellData(
	v1, v2 string,
	nonce uint64,
	reserve0 *big.Int,
	reserve1 *big.Int,
	tx types.Transaction,
	pool *TxPool) {
	input, _ := new(big.Int).SetString(v1, 16)
	output, _ := new(big.Int).SetString(v2, 16)
	totalCoin1 := new(big.Int).Sub(reserve1, input)
	totalCoin2 := new(big.Int).Add(reserve0, output)

	result, ok := dealWithcoinprice(totalCoin1, totalCoin2)

	if !ok {
		return
	}

	//TODO：send tx

	decimalPos2 := result.FloatString(5)

	priceCalc, err := strconv.ParseFloat(decimalPos2, 64)
	if err != nil {
		logrus.Errorf("ParseFloat  err : %v", err)
		return
	}

	amountIn, amountOut := dealwithAmout(priceDefaultDel, priceCalc, tx, BuyType)

	if amountIn == nil || amountOut == nil {
		return
	}

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
		pool,
	)
	if err != nil {
		logrus.Errorf("SendTx  err : %v  tx hash is %v", err, txHash.Hash())
		return
	}

	logrus.Infof(" tx succuess hash is %v", txHash.Hash())
}

func dealWithBuyData(
	v1, v2 string,
	nonce uint64,
	reserve0 *big.Int,
	reserve1 *big.Int,
	tx types.Transaction,
	pool *TxPool) {
	input, _ := new(big.Int).SetString(v1, 16)
	output, _ := new(big.Int).SetString(v2, 16)
	totalCoin1 := new(big.Int).Add(reserve1, input)
	totalCoin2 := new(big.Int).Sub(reserve0, output)

	result, ok := dealWithcoinprice(totalCoin1, totalCoin2)

	if !ok {
		return
	}

	//TODO：send tx

	decimalPos2 := result.FloatString(5)

	priceCalc, err := strconv.ParseFloat(decimalPos2, 64)
	if err != nil {
		logrus.Errorf("ParseFloat  err : %v", err)
		return
	}

	amountIn, amountOut := dealwithAmout(priceCalc, priceDefaultAdd, tx, SellType)

	if amountIn == nil || amountOut == nil {
		return
	}

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
		new(big.Int).SetUint64(nonce),
		pool)

	if err != nil {
		logrus.Errorf("SendTx  err : %v  tx hash is %v", err, txHash.Hash())
		return
	}

	logrus.Infof(" tx succuess hash is %v", txHash.Hash())

}
