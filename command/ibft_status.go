package command

import (
	"context"
	"fmt"

	ibftOp "github.com/0xPolygon/minimal/consensus/ibft/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

// IbftStatus is the command to query the snapshot
type IbftStatus struct {
	Meta
}

// GetHelperText returns a simple description of the command
func (p *IbftStatus) GetHelperText() string {
	return "Returns the current validator key of the IBFT client"
}

// Help implements the cli.IbftStatus interface
func (p *IbftStatus) Help() string {
	p.Meta.DefineFlags()

	usage := "ibft status"

	return p.GenerateHelp(p.Synopsis(), usage)
}

// Synopsis implements the cli.IbftStatus interface
func (p *IbftStatus) Synopsis() string {
	return p.GetHelperText()
}

// Run implements the cli.IbftStatus interface
func (p *IbftStatus) Run(args []string) int {
	flags := p.FlagSet("ibft status")

	if err := flags.Parse(args); err != nil {
		p.UI.Error(err.Error())
		return 1
	}

	conn, err := p.Conn()
	if err != nil {
		p.UI.Error(err.Error())
		return 1
	}

	clt := ibftOp.NewIbftOperatorClient(conn)
	resp, err := clt.Status(context.Background(), &empty.Empty{})
	if err != nil {
		p.UI.Error(err.Error())
		return 1
	}

	var output = "\n[VALIDATOR STATUS]\n"
	output += formatKV([]string{
		fmt.Sprintf("Vaidator key|%d", resp.Key),
	})

	p.UI.Output(output)

	return 0
}
