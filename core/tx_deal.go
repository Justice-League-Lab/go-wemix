package core

import (
	"context"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/script/erc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sirupsen/logrus"
)

func DealWithTx() {

	backend, err := ethclient.Dial(nodeHttp) // 本地节点的默认RPC端口
	if err != nil {
		logrus.Errorf("Dial client err : %v", err)
		return
	}

	rpcCli, err := rpc.Dial(nodeWebSite)
	if err != nil {
		logrus.Errorf("failed to dial: %v", err)
		return
	}

	gcli := gethclient.New(rpcCli)

	txch := make(chan common.Hash, 100)

	_, err = gcli.SubscribePendingTransactions(context.Background(), txch)
	if err != nil {
		logrus.Errorf("failed to SubscribePendingTransactions: %v", err)
		return
	}

	coreERC20, err := erc.NewCore(toAddress, backend)

	if err != nil {
		logrus.Errorf("erc NewCore  err : %v", err)
		return
	}

	callOpts := bind.CallOpts{
		Pending: true,                 // 是否查询待处理的区块
		From:    myAddress,            // 可选，指定调用者地址
		Context: context.Background(), // 可选，指定上下文
	}
	coreSERC20Http := &erc.CoreSession{
		Contract: coreERC20,
		CallOpts: callOpts,
	}

	signer := types.NewLondonSigner(new(big.Int).SetInt64(1111))
	for txhash := range txch {
		go func(txhash common.Hash, client *ethclient.Client, coreSERC20Http *erc.CoreSession) {
			tx, _, err := client.TransactionByHash(context.Background(), txhash)
			if err != nil {
				logrus.Errorf("failed to TransactionByHash: %v", err)
				return
			}

			address, err := types.Sender(signer, tx)
			if err != nil {
				logrus.Errorf("failed to Sender: %v", err)
				return
			}

			if FilterAddress(address) {
				DoTxHttp(*tx, "http", coreSERC20Http)
			}

		}(txhash, backend, coreSERC20Http)

	}

}

func DoTxHttp(tx types.Transaction, optType string, coreSERC20Http *erc.CoreSession) {
	defer func() {
		if e := recover(); e != nil {
			logrus.Errorf("painc err : %v", e)
		}
	}()

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
		Do0x06fd4ac5(txData, nonce, coinData.Reserve0, coinData.Reserve1, tx, coreSERC20Http, nil)
		Do0x09c5eabe(txData, nonce, coinData.Reserve0, coinData.Reserve1, tx, coreSERC20Http, nil)
		Do0x38ed1739(txData, nonce, coinData.Reserve0, coinData.Reserve1, tx, coreSERC20Http, nil)
		Do0x41876647(txData, nonce, coinData.Reserve0, coinData.Reserve1, tx, coreSERC20Http, nil)
		Do0x592db2b9(txData, nonce, coinData.Reserve0, coinData.Reserve1, tx, coreSERC20Http, nil)
		Do0xbaa2abde(txData, nonce, coinData.Reserve0, coinData.Reserve1, tx, coreSERC20Http, nil)
		Do0xd97495c9(txData, nonce, coinData.Reserve0, coinData.Reserve1, tx, coreSERC20Http, nil)
	}
}
