package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v2/modules/apps/interchain-price/host/types"
)

// IsHostEnabled retrieves the host enabled boolean from the paramstore.
// True is returned if the host submodule is enabled.
func (k Keeper) IsHostEnabled(ctx sdk.Context) bool {
	var res bool
	k.paramSpace.Get(ctx, types.KeyHostEnabled, &res)
	return res
}
