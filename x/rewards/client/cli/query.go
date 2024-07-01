package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/MANTRA-Finance/mantrachain/x/rewards/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group rewards queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListSnapshot())
	cmd.AddCommand(CmdShowSnapshot())
	cmd.AddCommand(CmdListProvider())
	cmd.AddCommand(CmdProviderPairs())
	cmd.AddCommand(CmdProvider())
	cmd.AddCommand(CmdRewards())
	// this line is used by starport scaffolding # 1

	return cmd
}
