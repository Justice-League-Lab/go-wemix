package core

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/script/erc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestDOTxScriptBuy(t *testing.T) {

	data := `38ed173900000000000000000000000000000000000000000000043C33C19375648000000000000000000000000000000000000000000000000005A59A5510F2C2D1000000000000000000000000000000000000000000000000000000000000000002a0000000000000000000000000bee95fd1c50099a8fff5204efd53c77900ab5052ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000000000000000000000000000000000000000000000000000000000020000000000000000000000008e81fcc2d4a3baa0ee9044e0d7e36f59c9bba9c1000000000000000000000000770d9d14c4ae2f78dca810958c1d9b7ea4620289`
	amount, _ := new(big.Int).SetString("20000749999999997500000", 10)
	tx := types.NewTransaction(1, common.HexToAddress("0x80a5A916FB355A8758f0a3e47891dc288DAC2665"), amount, 12345, new(big.Int), []byte(data))
	DOTxScript(*tx)
}

func TestDOTxScriptSolt(t *testing.T) {

	data := `38ed17390000000000000000000000000000000000000000000000001bc16d674ec80000000000000000000000000000000000000000000000000000148d29f70dc6491a00000000000000000000000000000000000000000000000000000000000000a00000000000000000000000003a3b580ac38ae2b937eb2d8848dec11fda317f82ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000770d9d14c4ae2f78dca810958c1d9b7ea46202890000000000000000000000008e81fcc2d4a3baa0ee9044e0d7e36f59c9bba9c1`
	amount, _ := new(big.Int).SetString("20000749999999997500000", 10)
	tx := types.NewTransaction(1, common.HexToAddress("0x80a5A916FB355A8758f0a3e47891dc288DAC2665"), amount, 300000, new(big.Int).SetInt64(100000000001), common.Hex2Bytes(data))
	DOTxScript(*tx)
}

func TestSendTx(t *testing.T) {

	client, err := ethclient.Dial(node) // 本地节点的默认RPC端口
	if err != nil {
		panic(err)
	}

	gas, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	data := `0x38ed1739
	0000000000000000000000000000000000000000000000003782dace9d900000
	00000000000000000000000000000000000000000000000094079cd1a42aa800
	00000000000000000000000000000000000000000000000000000000000000a0
	000000000000000000000000e8db41c5e9ef0f09a1e65f8dc8e9fef1879250a9
	ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff
	0000000000000000000000000000000000000000000000000000000000000002
	000000000000000000000000770d9d14c4ae2f78dca810958c1d9b7ea4620289
	0000000000000000000000008e81fcc2d4a3baa0ee9044e0d7e36f59c9bba9c1`
	amount, _ := new(big.Int).SetString("20000749999999997500000", 10)
	tx := types.NewTransaction(1, common.HexToAddress("0x80a5a916fb355a8758f0a3e47891dc288dac2665"), amount, 30000, gas, []byte(data))
	coreERC20, err := erc.NewCore(common.HexToAddress(contract), client)

	if err != nil {
		panic(err)
	}

	coreSERC20 = &erc.CoreSession{
		Contract: coreERC20,
	}

	privateKey, err = crypto.HexToECDSA("4e041f6f473bb6250db7688e2fba855b787708448840f271abadb3214944fec4")
	if err != nil {
		panic(err)
	}
	hash, err := SendTx(
		client,
		coreSERC20,
		*tx,
		privateKey,
		coin1,
		coin2,
		new(big.Int).SetInt64(1111), new(big.Int).SetInt64(2*1e18), new(big.Int).SetInt64(2.6666*1e18), new(big.Int).SetInt64(696))
	if err != nil {
		panic(err)
	}
	t.Logf("hash %v", hash)
}

func TestDOTxScript(t *testing.T) {

	data := `38ed17390000000000000000000000000000000000000000000000a2a15d09519be000000000000000000000000000000000000000000000000000782c526d30fcca5e0900000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000dbf7a48837687e349c833bd0d75e91a4bf86f697ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000770d9d14c4ae2f78dca810958c1d9b7ea46202890000000000000000000000008e81fcc2d4a3baa0ee9044e0d7e36f59c9bba9c1`
	amount, _ := new(big.Int).SetString("20000749999999997500000", 10)
	tx := types.NewTransaction(1, common.HexToAddress("0x80a5A916FB355A8758f0a3e47891dc288DAC2665"), amount, 12345, new(big.Int), common.Hex2Bytes(data))
	DOTxScript(*tx)
}
