package txpool

import (
	"fmt"
	"testing"
	"time"

	"github.com/0xPolygon/minimal/crypto"
	"github.com/0xPolygon/minimal/network"
	"github.com/0xPolygon/minimal/types"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

func TestTxPool(t *testing.T) {

}

func TestNotifyCh(t *testing.T) {
	pool, err := NewTxPool(hclog.NewNullLogger(), &mockStore{}, nil, nil)
	assert.NoError(t, err)

	pool.EnableDev()
	pool.NotifyCh = make(chan struct{}, 1)

	txn := &types.Transaction{
		From: types.Address{0x1},
	}
	assert.NoError(t, pool.AddTx(txn))

	select {
	case <-time.After(1 * time.Second):
		t.Fatal("bad")
	case <-pool.NotifyCh:
	}
}

func TestMultipleTransactions(t *testing.T) {
	// if we add the same transaction it should only be included once
	pool, err := NewTxPool(hclog.NewNullLogger(), &mockStore{}, nil, nil)
	assert.NoError(t, err)
	pool.EnableDev()

	from1 := types.Address{0x1}

	txn0 := &types.Transaction{
		From:  from1,
		Nonce: 10,
	}
	assert.NoError(t, pool.addImpl("", txn0))
	assert.NoError(t, pool.addImpl("", txn0))

	assert.Len(t, pool.queue[from1].txs, 1)
	assert.Equal(t, pool.Length(), uint64(0))

	from2 := types.Address{0x2}
	txn1 := &types.Transaction{
		From: from2,
	}
	assert.NoError(t, pool.addImpl("", txn1))
	assert.NoError(t, pool.addImpl("", txn1))

	assert.Len(t, pool.queue[from2].txs, 0)
	assert.Equal(t, pool.Length(), uint64(1))
}

func TestBroadcast(t *testing.T) {
	// TODO
	t.Skip()

	// we need a fully encrypted txn with (r, s, v) values so that we can
	// safely encrypt in RLP and broadcast it
	key0, _ := crypto.GenerateKey()
	addr0 := crypto.PubKeyToAddress(&key0.PublicKey)

	fmt.Println("-- addr")
	fmt.Println(addr0)

	signer := &crypto.FrontierSigner{}

	createPool := func() *TxPool {
		pool, err := NewTxPool(hclog.NewNullLogger(), &mockStore{}, nil, network.CreateServer(t, nil))
		assert.NoError(t, err)
		pool.AddSigner(signer)
		return pool
	}

	pool1 := createPool()
	pool2 := createPool()

	network.MultiJoin(t, pool1.network, pool2.network)

	// broadcast txn1 from pool1
	txn1 := &types.Transaction{}
	txn1, err := signer.SignTx(txn1, key0)
	assert.NoError(t, err)

	assert.NoError(t, pool1.AddTx(txn1))
	fmt.Println(pool1.Length())
}

type mockStore struct {
}

func (m *mockStore) GetNonce(root types.Hash, addr types.Address) uint64 {
	return 0
}

func (m *mockStore) GetBlockByHash(types.Hash, bool) (*types.Block, bool) {
	return nil, false
}

func (m *mockStore) Header() *types.Header {
	return &types.Header{}
}