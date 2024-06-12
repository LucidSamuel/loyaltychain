package cli

import (
    "strconv"
    "github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
    "github.com/cosmos/cosmos-sdk/client/flags"
    "github.com/cosmos/cosmos-sdk/client/tx"
    "loyaltychain/x/loyalty/types"
)

func CmdIssuePoints() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "issue-points [owner] [amount]",
        Short: "Issue loyalty points to an account",
        Args:  cobra.ExactArgs(2),
        RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx, err := client.GetClientTxContext(cmd)
            if err != nil {
                return err
            }

            owner := args[0]
            amount, err := strconv.Atoi(args[1])
            if err != nil {
                return err
            }

            msg := types.NewMsgIssuePoints(clientCtx.GetFromAddress().String(), owner, int32(amount))
            if err := msg.ValidateBasic(); err != nil {
                return err
            }

            return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
        },
    }

    flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdRedeemReward() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "redeem-reward [owner] [reward-item]",
        Short: "Redeem reward with loyalty points",
        Args:  cobra.ExactArgs(2),
        RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx, err := client.GetClientTxContext(cmd)
            if err != nil {
                return err
            }

            owner := args[0]
            rewardItem := args[1]

            msg := types.NewMsgRedeemReward(clientCtx.GetFromAddress().String(), owner, rewardItem)
            if err := msg.ValidateBasic(); err != nil {
                return err
            }

            return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
        },
    }

    flags.AddTxFlagsToCmd(cmd)

    return cmd
}
