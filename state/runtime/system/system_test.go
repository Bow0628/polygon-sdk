package system

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/0xPolygon/minimal/state/runtime"
	"github.com/0xPolygon/minimal/types"
)

type contractSetupParams struct {
	depth  int
	origin types.Address
	from   types.Address
	to     types.Address
	value  *big.Int
	gas    uint64
	code   []byte
}

func setupContract(params contractSetupParams) *runtime.Contract {
	return runtime.NewContract(
		params.depth,
		params.origin,
		params.from,
		params.to,
		params.value,
		params.gas,
		params.code,
	)
}

func TestSystem_CanRun(t *testing.T) {
	testTable := []struct {
		name     string
		contract *runtime.Contract
		canRun   bool
	}{
		{
			"Valid System runtime address",
			setupContract(
				contractSetupParams{
					depth:  0,
					origin: types.StringToAddress("0"),
					from:   types.StringToAddress("0"),
					to:     types.StringToAddress("1001"), // Staking handler
					value:  nil,
					gas:    0,
					code:   nil,
				},
			),
			true,
		},
		{
			"Invalid System runtime address",
			setupContract(
				contractSetupParams{
					depth:  0,
					origin: types.StringToAddress("0"),
					from:   types.StringToAddress("0"),
					to:     types.StringToAddress("9999"), // Invalid handler
					value:  nil,
					gas:    0,
					code:   nil,
				},
			),
			false,
		},
	}

	systemRuntime := NewSystem()

	for _, testCase := range testTable {
		canRun := systemRuntime.CanRun(testCase.contract, nil, nil)
		if !canRun && testCase.canRun {
			t.Fatal(fmt.Sprintf("[%s] Runtime doesn't recognize address %s", testCase.name, testCase.contract.CodeAddress))
		}
	}
}