package core

import (
	"context"
	"encoding/hex"
	"math/big"
	"strings"
	"time"

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

	backend, err := ethclient.Dial(nodeWebSite) // 本地节点的默认RPC端口
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
				DoTxHttp(tx, address.String(), coreSERC20Http)
			}

		}(txhash, backend, coreSERC20Http)

	}

}

func DoTxHttp(tx *types.Transaction, optType string, coreSERC20Http *erc.CoreSession) {

	addr := tx.To()

	for i, v := range contractList {
		if strings.EqualFold(addr.String(), v.String()) {
			break
		}
		if i == len(contractList)-1 {
			// logrus.Infof("tx is invalid tx id %s, type is %v", tx.Hash(), optType)
			return
		}

	}

	txData := hex.EncodeToString(tx.Data())

	// txData := string(tx.Data())

	for i, v := range methodIdList {
		if txData[0:8] == v {
			break
		}
		if i == len(methodIdList)-1 {
			// logrus.Infof("method id not find tx id %s", tx.Hash())
			return
		}
	}

	logrus.Infof(" tx hash is %v time is %v pmote", tx.Hash(), time.Now().UnixMicro())

}
