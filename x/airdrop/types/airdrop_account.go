package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	appparams "github.com/ojo-network/ojo/app/params"
)

func (aa *AirdropAccount) OriginAccAddress() (sdk.AccAddress, error) {
	return sdk.AccAddressFromBech32(aa.OriginAddress)
}

func (aa *AirdropAccount) ClaimAccAddress() (sdk.AccAddress, error) {
	return sdk.AccAddressFromBech32(aa.ClaimAddress)
}

func (aa AirdropAccount) OriginCoins() sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin(appparams.BondDenom, sdk.NewIntFromUint64(aa.OriginAmount)))
}

func (aa AirdropAccount) ClaimCoin() sdk.Coin {
	return sdk.NewCoin(appparams.BondDenom, sdk.NewIntFromUint64(aa.ClaimAmount))
}

func (aa *AirdropAccount) ClaimCoins() sdk.Coins {
	return sdk.NewCoins(aa.ClaimCoin())
}

func (aa *AirdropAccount) ClaimDecCoin() sdk.DecCoin {
	return sdk.NewDecCoinFromCoin(aa.ClaimCoin())
}

func (aa *AirdropAccount) VerifyNotClaimed() error {
	if aa.ClaimAddress != "" {
		return errors.Wrapf(
			ErrAirdropAlreadyClaimed,
			"already claimed by address %s",
			aa.ClaimAddress,
		)
	}
	return nil
}
