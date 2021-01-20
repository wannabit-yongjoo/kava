package hard

// DO NOT EDIT - generated by aliasgen tool (github.com/rhuairahrighairidh/aliasgen)

import (
	"github.com/kava-labs/kava/x/hard/keeper"
	"github.com/kava-labs/kava/x/hard/types"
)

const (
	AttributeKeyBlockHeight            = types.AttributeKeyBlockHeight
	AttributeKeyClaimAmount            = types.AttributeKeyClaimAmount
	AttributeKeyClaimHolder            = types.AttributeKeyClaimHolder
	AttributeKeyClaimMultiplier        = types.AttributeKeyClaimMultiplier
	AttributeKeyClaimType              = types.AttributeKeyClaimType
	AttributeKeyDeposit                = types.AttributeKeyDeposit
	AttributeKeyDepositDenom           = types.AttributeKeyDepositDenom
	AttributeKeyDepositor              = types.AttributeKeyDepositor
	AttributeKeyRewardsDistribution    = types.AttributeKeyRewardsDistribution
	AttributeValueCategory             = types.AttributeValueCategory
	DefaultParamspace                  = types.DefaultParamspace
	DelegatorAccount                   = types.DelegatorAccount
	EventTypeClaimHardReward           = types.EventTypeClaimHardReward
	EventTypeDeleteHardDeposit         = types.EventTypeDeleteHardDeposit
	EventTypeHardDelegatorDistribution = types.EventTypeHardDelegatorDistribution
	EventTypeHardDeposit               = types.EventTypeHardDeposit
	EventTypeHardLPDistribution        = types.EventTypeHardLPDistribution
	EventTypeHardWithdrawal            = types.EventTypeHardWithdrawal
	LP                                 = types.LP
	LPAccount                          = types.LPAccount
	LiquidatorAccount                  = types.LiquidatorAccount
	Large                              = types.Large
	Medium                             = types.Medium
	ModuleAccountName                  = types.ModuleAccountName
	ModuleName                         = types.ModuleName
	QuerierRoute                       = types.QuerierRoute
	QueryGetClaims                     = types.QueryGetClaims
	QueryGetDeposits                   = types.QueryGetDeposits
	QueryGetModuleAccounts             = types.QueryGetModuleAccounts
	QueryGetParams                     = types.QueryGetParams
	RouterKey                          = types.RouterKey
	Small                              = types.Small
	Stake                              = types.Stake
	StoreKey                           = types.StoreKey
)

var (
	// function aliases
	NewKeeper                     = keeper.NewKeeper
	NewQuerier                    = keeper.NewQuerier
	CalculateUtilizationRatio     = keeper.CalculateUtilizationRatio
	CalculateBorrowRate           = keeper.CalculateBorrowRate
	CalculateBorrowInterestFactor = keeper.CalculateBorrowInterestFactor
	CalculateSupplyInterestFactor = keeper.CalculateSupplyInterestFactor
	APYToSPY                      = keeper.APYToSPY
	DefaultGenesisState           = types.DefaultGenesisState
	DefaultParams                 = types.DefaultParams
	DepositTypeIteratorKey        = types.DepositTypeIteratorKey
	GetTotalVestingPeriodLength   = types.GetTotalVestingPeriodLength
	NewDeposit                    = types.NewDeposit
	NewGenesisState               = types.NewGenesisState
	NewMsgClaimReward             = types.NewMsgClaimReward
	NewMsgDeposit                 = types.NewMsgDeposit
	NewMsgWithdraw                = types.NewMsgWithdraw
	NewMultiplier                 = types.NewMultiplier
	NewParams                     = types.NewParams
	NewPeriod                     = types.NewPeriod
	NewQueryAccountParams         = types.NewQueryAccountParams
	NewQueryClaimParams           = types.NewQueryClaimParams
	ParamKeyTable                 = types.ParamKeyTable
	RegisterCodec                 = types.RegisterCodec

	// variable aliases
	BorrowsKeyPrefix                 = types.BorrowsKeyPrefix
	DefaultActive                    = types.DefaultActive
	DefaultPreviousBlockTime         = types.DefaultPreviousBlockTime
	ClaimTypesClaimQuery             = types.ClaimTypesClaimQuery
	DepositsKeyPrefix                = types.DepositsKeyPrefix
	ErrAccountNotFound               = types.ErrAccountNotFound
	ErrClaimExpired                  = types.ErrClaimExpired
	ErrClaimNotFound                 = types.ErrClaimNotFound
	ErrDepositNotFound               = types.ErrDepositNotFound
	ErrGovScheduleNotFound           = types.ErrGovScheduleNotFound
	ErrInsufficientModAccountBalance = types.ErrInsufficientModAccountBalance
	ErrInvaliWithdrawAmount          = types.ErrInvalidWithdrawAmount
	ErrInvalidAccountType            = types.ErrInvalidAccountType
	ErrInvalidDepositDenom           = types.ErrInvalidDepositDenom
	ErrInvalidClaimType              = types.ErrInvalidClaimType
	ErrInvalidMultiplier             = types.ErrInvalidMultiplier
	ErrLPScheduleNotFound            = types.ErrLPScheduleNotFound
	ErrZeroClaim                     = types.ErrZeroClaim
	GovDenom                         = types.GovDenom
	KeyActive                        = types.KeyActive
	ModuleCdc                        = types.ModuleCdc
	PreviousBlockTimeKey             = types.PreviousBlockTimeKey
)

type (
	Keeper             = keeper.Keeper
	AccountKeeper      = types.AccountKeeper
	Borrow             = types.Borrow
	MoneyMarket        = types.MoneyMarket
	MoneyMarkets       = types.MoneyMarkets
	Deposit            = types.Deposit
	ClaimType          = types.ClaimType
	GenesisState       = types.GenesisState
	MsgClaimReward     = types.MsgClaimReward
	MsgDeposit         = types.MsgDeposit
	MsgWithdraw        = types.MsgWithdraw
	Multiplier         = types.Multiplier
	MultiplierName     = types.MultiplierName
	Multipliers        = types.Multipliers
	Params             = types.Params
	QueryAccountParams = types.QueryAccountParams
	QueryClaimParams   = types.QueryClaimParams
	StakingKeeper      = types.StakingKeeper
	SupplyKeeper       = types.SupplyKeeper
)
