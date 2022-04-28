package cli

import (
	"fmt"
	"time"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/VelaChain/orion/x/chios/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdCreatePairPool(),
		CmdJoinPairPool(),
		CmdExitPairPool(),
		CmdSwapPair(),
		CmdAddLiquidityPair(),
		CmdRemoveLiquidityPair(),
	)
	// this line is used by starport scaffolding # 1

	return cmd
}

// TODO
func CmdCreatePairPool() *cobra.Command {
	cmd := &cobra.Command {
		Use:	"create-pair-pool [denom-a] [amount-a] [denom-b] [amount-b] [shares]",
		Short:	"Broadcast message create-pair-pool",
		Args:	cobra.ExactArgs(6),
		RunE:	func(cmd *cobra.Command, args []string) (err error) {
			// handle each arg
			argDenomA := args[0]
			
			amountA, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			argAmountA := sdk.NewIntFromUint64(amountA)

			argDenomB := args[2]

			amountB, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			argAmountB := sdk.NewIntFromUint64(amountB)

			shares, err := cast.ToUint64E(args[4])
			if err != nil{
				return err
			}

			argSharesOut := sdk.NewIntFromUint64(shares) 

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// create msg
			msg := types.MsgCreatePairPool{
				Creator:	clientCtx.GetFromAddress().String(),
				DenomA: 	argDenomA,
				AmountA:	argAmountA,
				DenomB:		argDenomB,
				AmountB:	argAmountB,
				SharesOut:	argSharesOut,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// TODO
func CmdJoinPairPool() *cobra.Command {
	cmd := &cobra.Command {
		Use:	"join-pair-pool [denom-a] [amount-a] [denom-b] [amount-b] [shares]",
		Short:	"Broadcast message join-pair-pool",
		Args:	cobra.ExactArgs(6),
		RunE:	func(cmd *cobra.Command, args []string) (err error) {
			// handle each arg
			argDenomA := args[0]
			
			amountA, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			argAmountA := sdk.NewIntFromUint64(amountA)

			argDenomB := args[2]

			amountB, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			argAmountB := sdk.NewIntFromUint64(amountB)

			shares, err := cast.ToUint64E(args[4])
			if err != nil{
				return err
			}

			argSharesOut := sdk.NewIntFromUint64(shares) 

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
					return err
			}
			// create msg
			msg := types.MsgJoinPairPool{
				Creator:	clientCtx.GetFromAddress().String(),
				DenomA:		argDenomA,
				AmountA:	argAmountA,
				DenomB:		argDenomB,
				AmountB:	argAmountB,
				SharesOut:	argSharesOut,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// TODO
func CmdExitPairPool() *cobra.Command {
	cmd := &cobra.Command {
		Use:	"exit-pair-pool [pool-id]",
		Short:	"Broadcast message exit-pair-pool",
		Args:	cobra.ExactArgs(2),
		RunE:	func(cmd *cobra.Command, args []string) (err error) {
			// handle each arg
			argPoolId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
					return err
			}
			// create msg
			msg := types.MsgExitPairPool{
				Creator:	clientCtx.GetFromAddress().String(),
				PoolId:		argPoolId,
			}	
			if err := msg.ValidateBasic(); err != nill {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// TODO
func CmdSwapPair() *cobra.Command {
	cmd := &cobra.Command {
		Use:	"swap-pair [denom-in] [amount-in] [denom-out] [min-amount-out]",
		Short:	"",
		Args:	cobra.ExactArgs(5),
		RunE:	func(cmd *cobra.Command, args []string) (err error) {
			// handle each arg
			argDenomIn := args[0]
			
			amountIn, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			argAmountIn := sdk.NewIntFromUint64(amountIn)

			argDenomOut := args[2]

			amountOut, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			argMinAmountOut := sdk.NewIntFromUint64(amountOut)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
					return err
			}
			// create msg
			msg := types.MsgSwapPair{
				Creator:		clientCtx.GetFromAddress().String(),
				DenomIn:		argDenomIn,
				AmountIn:		argAmountIn,
				DenomOut:		argDenomOut,
				MinAmountOut:	argMinAmountOut,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// TODO
func CmdAddLiquidityPair() *cobra.Command {
	cmd := &cobra.Command {
		Use:	"add-liquidity-pair [denom-a] [amount-a] [denom-b] [amount-b] [shares]",
		Short:	"Broadcast message add-liquidity-pair",
		Args:	cobra.ExactArgs(6),
		RunE:	func(cmd *cobra.Command, args []string) (err error) {
			// handle each arg
			argDenomA := args[0]
			
			amountA, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			argAmountA := sdk.NewIntFromUint64(amountA)

			argDenomB := args[2]

			amountB, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			argAmountB := sdk.NewIntFromUint64(amountB)

			shares, err := cast.ToUint64E(args[4])
			if err != nil{
				return err
			}

			argSharesOut := sdk.NewIntFromUint64(shares) 

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
					return err
			}
			// create msg
			msg := types.MsgAddLiquidityPair{
				Creator:		clientCtx.GetFromAddress().String(),
				DenomA:			argDenomA,
				AmountA:		argAmountA,
				DenomB:			argDenomB,
				AmountB:		argAmountB,
				SharesOut:		argSharesOut,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// TODO
func CmdRemoveLiquidityPair() *cobra.Command {
	cmd := &cobra.Command {
		Use:	"remove-liquidity-pair [shares-denom] [shares-amount]",
		Short:	"Broadcast message remove-liquidity-pair",
		Args:	cobra.ExactArgs(3),
		RunE:	func(cmd *cobra.Command, args []string) (err error) {
			// handle each arg
			argSharesDenom := args[0]

			sharesAmount, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			argSharesAmount := sdk.NewIntFromUint64(sharesAmount)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
					return err
			}
			// create msg
			msg := types.MsgRemoveLiquidityPair{
				Creator:		clientCtx.GetFromAddress().String(),
				SharesDenom:	argShareDenom, 
				SharesAmount:	argShareAmount,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}	