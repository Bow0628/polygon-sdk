package system

import (
	"math/big"

	"github.com/0xPolygon/minimal/types"
)

// unstakingHandler implements the unstaking logic for the System runtime
type unstakingHandler struct {
	s *System
}

// gas returns the fixed gas price of the unstaking operation
func (uh *unstakingHandler) gas(_ []byte) uint64 {
	return 40000
}

// run executes the system contract unstaking method
// Unstaking is for the FULL amount staked, no partial support is currently present
func (uh *unstakingHandler) run(state *systemState) ([]byte, error) {
	// Grab the unstaking address
	unstakingAddress := types.StringToAddress(GetOperationsMap()["unstaking"])

	// Grab the address calling the staking method
	staker := state.contract.Caller

	// Grab the staked balance for the account
	stakedBalance := state.host.GetStakedBalance(staker)

	// Grab the balance on the unstaking address
	unstakingAccountBalance := state.host.GetBalance(unstakingAddress)

	// Sanity check
	// Can't unstake if the balance isn't previously present on the unstaking account
	// Can't unstake if the value is different from 0
	if unstakingAccountBalance.Cmp(stakedBalance) < 0 || state.contract.Value.Cmp(big.NewInt(0)) != 0 {
		return nil, nil
	}

	// Decrease the staked balance on the unstaking address
	state.host.SubBalance(unstakingAddress, stakedBalance)

	// Decrease the staked amount from the account's staked balance
	state.host.SubStakedBalance(staker, stakedBalance)

	// Increase the account's actual balance
	state.host.AddBalance(staker, stakedBalance)

	// TODO Remove the staker from the validator set after this point + checks

	return nil, nil
}
