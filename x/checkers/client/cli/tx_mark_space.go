package cli

import (
	"strconv"

	"github.com/alice/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMarkSpace() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-space [id-value] [x] [y]",
		Short: "Broadcast message markSpace",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argIdValue, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argX, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argY, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMarkSpace(
				clientCtx.GetFromAddress().String(),
				argIdValue,
				argX,
				argY,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
