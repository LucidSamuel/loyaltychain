package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"loyaltychain/app"
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
    "github.com/cosmos/cosmos-sdk/server"
    "github.com/cosmos/cosmos-sdk/server/cmd"
    "loyaltychain/x/loyalty/client/cli"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}

func NewRootCmd() (*cobra.Command, simapp.EncodingConfig) {
    encodingConfig := simapp.MakeEncodingConfig()
    initClientCtx := client.Context{}.
        WithCodec(encodingConfig.Marshaler).
        WithTxConfig(encodingConfig.TxConfig).
        WithInterfaceRegistry(encodingConfig.InterfaceRegistry)

    rootCmd := &cobra.Command{
        Use:   "loyaltychaind",
        Short: "LoyaltyChain App",
        PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
            initClientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
            if err != nil {
                return err
            }

            return client.SetCmdClientContextHandler(initClientCtx, cmd)
        },
    }

    server.AddCommands(rootCmd, simapp.DefaultNodeHome, NewApp, createSimappAndExport)

    // Register the CLI commands
    rootCmd.AddCommand(
        cli.CmdIssuePoints(),
        cli.CmdRedeemReward(),
    )

    return rootCmd, encodingConfig
}