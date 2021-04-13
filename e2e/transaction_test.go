package e2e

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/umbracle/go-web3"
	"github.com/umbracle/go-web3/jsonrpc"

	"github.com/0xPolygon/minimal/e2e/framework"
	"github.com/0xPolygon/minimal/types"
)

func TestTransaction(t *testing.T) {

	clt, err := jsonrpc.NewClient("http://127.0.0.1:8545")
	if err != nil {
		t.Fatal(err)
	}

	eth := clt.Eth()
	fmt.Println(eth.BlockNumber())

	root, err := eth.SendTransaction(&web3.Transaction{
		From:     web3.HexToAddress("0x9bd03347a977e4deb0a0ad685f8385f264524b0b"),
		To:       web3.HexToAddress("0x9bd03347a977e4deb0a0ad685f8385f264524b0a").String(),
		GasPrice: 10000,
		Gas:      10000000000,
		Value:    big.NewInt(10000),
	})
	assert.NoError(t, err)
	fmt.Println(root)
}

func TestPreminedBalance(t *testing.T) {
	validAccounts := []struct {
		address types.Address
		balance *big.Int
	}{
		{types.StringToAddress("1"), big.NewInt(0)},
		{types.StringToAddress("2"), big.NewInt(20)},
	}

	testTable := []struct {
		name       string
		address    types.Address
		balance    *big.Int
		shouldFail bool
	}{
		{
			"Account with 0 balance",
			validAccounts[0].address,
			validAccounts[0].balance,
			false,
		},
		{
			"Account with valid balance",
			validAccounts[1].address,
			validAccounts[1].balance,
			false,
		},
		{
			"Account not in genesis",
			types.StringToAddress("3"),
			big.NewInt(0),
			true,
		},
	}

	srv := framework.NewTestServer(t, func(config *framework.TestServerConfig) {
		for _, acc := range validAccounts {
			config.Premine(acc.address, acc.balance)
		}
	})
	defer srv.Stop()

	rpcClient := srv.JSONRPC()

	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T) {

			balance, err := rpcClient.Eth().GetBalance(web3.Address(testCase.address), web3.Latest)

			if err != nil && !testCase.shouldFail {
				assert.Failf(t, "Uncaught error", err.Error())
			}

			if !testCase.shouldFail {
				assert.Equalf(t, testCase.balance, balance, "Balances don't match")
			}
		})
	}
}

func ethToWei(ethValue int64) *big.Int {
	return new(big.Int).Mul(
		big.NewInt(ethValue),
		new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
}

func TestEthTransfer(t *testing.T) {
	validAccounts := []struct {
		address types.Address
		balance *big.Int
	}{
		// Valid account #1
		{
			types.StringToAddress("1"),
			ethToWei(50), // 50 ETH
		},
		// Empty account
		{
			types.StringToAddress("2"),
			big.NewInt(0)},
		// Valid account #2
		{
			types.StringToAddress("3"),
			ethToWei(10), // 10 ETH
		},
	}

	testTable := []struct {
		name       string
		sender     types.Address
		recipient  types.Address
		amount     *big.Int
		shouldFail bool
	}{
		{
			// ACC #1 -> ACC #3
			"Valid ETH transfer #1",
			validAccounts[0].address,
			validAccounts[2].address,
			ethToWei(10), // 10 ETH
			false,
		},
		{
			// ACC #2 -> ACC #3
			"Invalid ETH transfer",
			validAccounts[1].address,
			validAccounts[2].address,
			ethToWei(100),
			true,
		},
		{
			// ACC #2 -> ACC #1
			"Valid ETH transfer #2",
			validAccounts[2].address,
			validAccounts[1].address,
			ethToWei(5), // 5 ETH
			false,
		},
	}

	srv := framework.NewTestServer(t, func(config *framework.TestServerConfig) {
		config.SetDev(true)
		config.SetSeal(true)

		for _, acc := range validAccounts {
			config.Premine(acc.address, acc.balance)
		}
	})
	defer srv.Stop()

	checkSenderReceiver := func(errSender error, errReceiver error, t *testing.T) {
		if errSender != nil || errReceiver != nil {
			if errSender != nil {
				assert.Failf(t, "Uncaught error", errSender.Error())
			} else {
				assert.Failf(t, "Uncaught error", errReceiver.Error())
			}
		}
	}

	rpcClient := srv.JSONRPC()

	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T) {

			preSendData := struct {
				previousSenderBalance   *big.Int
				previousReceiverBalance *big.Int
			}{
				previousSenderBalance:   big.NewInt(0),
				previousReceiverBalance: big.NewInt(0),
			}

			// Fetch the balances before sending
			balanceSender, errSender := rpcClient.Eth().GetBalance(
				web3.Address(testCase.sender),
				web3.Latest,
			)

			balanceReceiver, errReceiver := rpcClient.Eth().GetBalance(
				web3.Address(testCase.recipient),
				web3.Latest,
			)

			checkSenderReceiver(errSender, errReceiver, t)

			// Set the preSend balances
			preSendData.previousSenderBalance = balanceSender
			preSendData.previousReceiverBalance = balanceReceiver

			// Create the transaction
			txnObject := &web3.Transaction{
				From:     web3.Address(testCase.sender),
				To:       testCase.recipient.String(),
				GasPrice: 100,
				Gas:      100,
				Value:    testCase.amount,
			}

			// Do the transfer
			txnHash, err := rpcClient.Eth().SendTransaction(txnObject)

			// Error checking
			if err != nil && !testCase.shouldFail {
				assert.Failf(t, "Uncaught error", err.Error())
			}

			// Wait until the transaction goes through
			time.Sleep(10 * time.Second)

			// Fetch the balances after sending
			balanceSender, errSender = rpcClient.Eth().GetBalance(
				web3.Address(testCase.sender),
				web3.Latest,
			)

			balanceReceiver, errReceiver = rpcClient.Eth().GetBalance(
				web3.Address(testCase.recipient),
				web3.Latest,
			)

			assert.IsTypef(t, web3.Hash{}, txnHash, "Return type mismatch")

			checkSenderReceiver(errSender, errReceiver, t)

			if !testCase.shouldFail {
				// Check the sender balance
				assert.Equalf(t,
					new(big.Int).Sub(preSendData.previousSenderBalance, testCase.amount),
					balanceSender,
					"Sender balance incorrect")

				// Check the receiver balance
				assert.Equalf(t,
					new(big.Int).Add(preSendData.previousReceiverBalance, testCase.amount),
					balanceReceiver,
					"Receiver balance incorrect")
			}
		})
	}
}