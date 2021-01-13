package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName name that will be used throughout the module
	ModuleName = "hard"

	// LPAccount LP distribution module account
	LPAccount = "hard_lp_distribution"

	// DelegatorAccount delegator distribution module account
	DelegatorAccount = "hard_delegator_distribution"

	// LiquidatorAccount module account for liquidator
	LiquidatorAccount = "hard_liquidator"

	// ModuleAccountName name of module account used to hold deposits
	ModuleAccountName = "hard"

	// StoreKey Top level store key where all module items will be stored
	StoreKey = ModuleName

	// RouterKey Top level router key
	RouterKey = ModuleName

	// QuerierRoute Top level query string
	QuerierRoute = ModuleName

	// DefaultParamspace default name for parameter store
	DefaultParamspace = ModuleName
)

var (
	PreviousBlockTimeKey      = []byte{0x01}
	DepositsKeyPrefix         = []byte{0x02}
	BorrowsKeyPrefix          = []byte{0x03}
	BorrowedCoinsPrefix       = []byte{0x04}
	MoneyMarketsPrefix        = []byte{0x05}
	PreviousAccrualTimePrefix = []byte{0x06} // denom -> time
	TotalReservesPrefix       = []byte{0x07} // denom -> sdk.Coin
	InterestFactorPrefix      = []byte{0x8}  // denom -> sdk.Dec
	LtvIndexPrefix            = []byte{0x9}
	sep                       = []byte(":")
)

// DepositTypeIteratorKey returns an interator prefix for interating over deposits by deposit denom
func DepositTypeIteratorKey(denom string) []byte {
	return createKey([]byte(denom))
}

// GetBorrowByLtvKey is used by the LTV index
func GetBorrowByLtvKey(ltv sdk.Dec, borrower sdk.AccAddress) []byte {
	return append(ltv.Bytes(), borrower...)
}

func createKey(bytes ...[]byte) (r []byte) {
	for _, b := range bytes {
		r = append(r, b...)
	}
	return
}