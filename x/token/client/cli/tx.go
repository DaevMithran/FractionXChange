package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/MANTRA-Finance/mantrachain/x/token/types"
	"github.com/cosmos/cosmos-sdk/client"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdCreateNftCollection(),
		CmdMintNfts(),
		CmdBurnNfts(),
		CmdTransferNfts(),
		CmdApproveNfts(),
		CmdApproveAllNfts(),
		CmdMintNft(),
		CmdBurnNft(),
		CmdTransferNft(),
		CmdApproveNft(),
	)
	// this line is used by starport scaffolding # 1

	return cmd
}
