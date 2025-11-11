package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/testchain/evmos/x/blockreward/types"
)

type Keeper struct {
	bankKeeper  types.BankKeeper
	distrKeeper types.DistrKeeper
}

func NewKeeper(bankKeeper types.BankKeeper, distrKeeper types.DistrKeeper) Keeper {
	return Keeper{
		bankKeeper:  bankKeeper,
		distrKeeper: distrKeeper,
	}
}

// MintAndDistribute mints ~9,510 tokens per block until total supply reaches 300 Cr.
func (k Keeper) MintAndDistribute(ctx sdk.Context) error {
	// Constants (in whole tokens × 1e18 for base units)
	maxSupply := sdk.NewInt(3_000_000_000).MulRaw(1e18) // 300 Cr
	yearlyMint := sdk.NewInt(6_000_0000).MulRaw(1e18)   // 6 Cr per year
	blocksPerYear := sdk.NewInt(6_307_200)
	perBlockMint := yearlyMint.Quo(blocksPerYear)        // ≈ 9,510 tokens

	totalSupply := k.bankKeeper.GetSupply(ctx, "aevmos").Amount
	if totalSupply.GTE(maxSupply) {
		return nil // Stop minting once 300 Cr reached
	}

	rewardCoin := sdk.NewCoin("aevmos", perBlockMint)

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(rewardCoin)); err != nil {
		return err
	}

	// Send to distribution module for validator rewards
	if err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, "distribution", sdk.NewCoins(rewardCoin)); err != nil {
		return err
	}

	return nil
}
