package cli

import (
	"fmt"

	"github.com/MANTRA-Finance/mantrachain/x/coinfactory/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group coinfactory queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryDenomAuthorityMetadata())
	cmd.AddCommand(CmdQueryDenomAuthorityMetadata2())
	cmd.AddCommand(CmdQueryDenomsFromCreator())
	cmd.AddCommand(CmdQueryBalance())
	cmd.AddCommand(CmdQueryDenomMetadata())
	cmd.AddCommand(CmdQuerySupplyOf())
	// this line is used by starport scaffolding # 1

	return cmd
}
