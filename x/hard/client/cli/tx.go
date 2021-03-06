package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/kava-labs/kava/x/hard/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	hardTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	hardTxCmd.AddCommand(flags.PostCommands(
		getCmdDeposit(cdc),
		getCmdWithdraw(cdc),
		getCmdBorrow(cdc),
		addOptionalFlag(getCmdRepay(cdc), flagOwner, "", "original borrower's address whose loan will be repaid"),
		getCmdLiquidate(cdc),
	)...)

	return hardTxCmd
}

// addFlag adds a cobra flag and binds it using viper
func addOptionalFlag(cmd *cobra.Command, flagName, flagValue, flagUsage string) *cobra.Command {
	cmd.Flags().String(flagName, flagValue, flagUsage)
	viper.BindPFlag(flagName, cmd.Flags().Lookup(flagName))
	return cmd
}

func getCmdDeposit(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "deposit [amount]",
		Short: "deposit coins to hard",
		Example: fmt.Sprintf(
			`%s tx %s deposit 10000000bnb --from <key>`, version.ClientName, types.ModuleName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			amount, err := sdk.ParseCoins(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgDeposit(cliCtx.GetFromAddress(), amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func getCmdWithdraw(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "withdraw [amount]",
		Short: "withdraw coins from hard",
		Args:  cobra.ExactArgs(1),
		Example: fmt.Sprintf(
			`%s tx %s withdraw 10000000bnb --from <key>`, version.ClientName, types.ModuleName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			amount, err := sdk.ParseCoins(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgWithdraw(cliCtx.GetFromAddress(), amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func getCmdBorrow(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "borrow [amount]",
		Short: "borrow tokens from the hard protocol",
		Long:  strings.TrimSpace(`borrows tokens from the hard protocol`),
		Args:  cobra.ExactArgs(1),
		Example: fmt.Sprintf(
			`%s tx %s borrow 1000000000ukava --from <key>`, version.ClientName, types.ModuleName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			coins, err := sdk.ParseCoins(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgBorrow(cliCtx.GetFromAddress(), coins)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func getCmdRepay(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "repay [amount]",
		Short: "repay tokens to the hard protocol",
		Long:  strings.TrimSpace(`repay tokens to the hard protocol with optional --owner param to repay another account's loan`),
		Args:  cobra.ExactArgs(1),
		Example: strings.TrimSpace(`
kvcli tx hard repay 1000000000ukava --from <key>
kvcli tx hard repay 1000000000ukava,25000000000bnb --from <key>
kvcli tx hard repay 1000000000ukava,25000000000bnb --owner <owner-address> --from <key>
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			var owner sdk.AccAddress
			ownerStr := viper.GetString(flagOwner)

			// Parse optional owner argument or default to sender
			if len(ownerStr) > 0 {
				ownerAddr, err := sdk.AccAddressFromBech32(ownerStr)
				if err != nil {
					return err
				}
				owner = ownerAddr
			} else {
				owner = cliCtx.GetFromAddress()
			}

			coins, err := sdk.ParseCoins(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgRepay(cliCtx.GetFromAddress(), owner, coins)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func getCmdLiquidate(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "liquidate [borrower-addr]",
		Short: "liquidate a borrower that's over their loan-to-value ratio",
		Long:  strings.TrimSpace(`liquidate a borrower that's over their loan-to-value ratio`),
		Args:  cobra.ExactArgs(1),
		Example: fmt.Sprintf(
			`%s tx %s liquidate kava1hgcfsuwc889wtdmt8pjy7qffua9dd2tralu64j --from <key>`, version.ClientName, types.ModuleName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			borrower, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgLiquidate(cliCtx.GetFromAddress(), borrower)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
