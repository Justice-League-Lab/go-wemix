package core

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
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

	for txhash := range txch {
		go func(txhash common.Hash, client *ethclient.Client) {
			tx, _, err := client.TransactionByHash(context.Background(), txhash)
			if err != nil {
				logrus.Errorf("failed to TransactionByHash: %v", err)
				return
			}
			DOTxScript(*tx, "http")
		}(txhash, backend)

	}

}
