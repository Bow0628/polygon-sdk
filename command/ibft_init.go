package command

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/0xPolygon/minimal/types"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/0xPolygon/minimal/consensus/ibft"
	"github.com/0xPolygon/minimal/crypto"
	"github.com/0xPolygon/minimal/minimal"
	"github.com/0xPolygon/minimal/network"
)

// IbftInit is the command to query the snapshot
type IbftInit struct {
	Meta
}

// GetHelperText returns a simple description of the command
func (p *IbftInit) GetHelperText() string {
	return "Initializes IBFT for the Polygon SDK, in the specified directory"
}

// Help implements the cli.IbftInit interface
func (p *IbftInit) Help() string {
	usage := fmt.Sprintf("%s DATA_DIRECTORY", p.GetBaseCommand())

	return types.GenerateHelp(p.Synopsis(), usage, p.flagMap)
}

// Synopsis implements the cli.IbftInit interface
func (p *IbftInit) Synopsis() string {
	return p.GetHelperText()
}

func (p *IbftInit) GetBaseCommand() string {
	return "ibft-init"
}

// Run implements the cli.IbftInit interface
func (p *IbftInit) Run(args []string) int {
	flags := flag.NewFlagSet(p.GetBaseCommand(), flag.ContinueOnError)
	if err := flags.Parse(args); err != nil {
		p.UI.Error(err.Error())
		return 1
	}

	args = flags.Args()
	if len(args) != 1 {
		p.UI.Error("required argument (data directory) not passed in")
		return 1
	}

	pathName := args[0]
	if err := minimal.SetupDataDir(pathName, []string{"consensus", "libp2p"}); err != nil {
		p.UI.Error(err.Error())
		return 1
	}

	// try to write the ibft private key
	key, err := crypto.ReadPrivKey(filepath.Join(pathName, "consensus", ibft.IbftKeyName))
	if err != nil {
		p.UI.Error(err.Error())
		return 1
	}

	// try to create also a libp2p address
	libp2pKey, err := network.ReadLibp2pKey(filepath.Join(pathName, "libp2p"))
	if err != nil {
		p.UI.Error(err.Error())
		return 1
	}

	nodeId, err := peer.IDFromPrivateKey(libp2pKey)
	if err != nil {
		p.UI.Error(err.Error())
		return 1
	}

	output := "\n[IBFT INIT]\n"

	output += formatKV([]string{
		fmt.Sprintf("Public key (address)|%s", crypto.PubKeyToAddress(&key.PublicKey)),
		fmt.Sprintf("Node ID|%s", nodeId.String()),
	})

	output += "\n"

	p.UI.Output(output)

	return 0
}
