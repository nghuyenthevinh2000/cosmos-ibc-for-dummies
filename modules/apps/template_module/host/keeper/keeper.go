package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	icptypes "github.com/cosmos/ibc-go/v2/modules/apps/interchain-price/types"

	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
)

// Keeper defines the IBC price keeper
type Keeper struct {
	paramSpace paramtypes.Subspace

	channelKeeper icptypes.ChannelKeeper
	portKeeper    icptypes.PortKeeper

	scopedKeeper capabilitykeeper.ScopedKeeper
}

// NewKeeper creates a new IBC price Keeper instance
func NewKeeper(
	paramSpace paramtypes.Subspace,
	channelKeeper icptypes.ChannelKeeper,
	portKeeper icptypes.PortKeeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
) Keeper {
	return Keeper{
		paramSpace:    paramSpace,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		scopedKeeper:  scopedKeeper,
	}
}

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function
func (k Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	return k.scopedKeeper.AuthenticateCapability(ctx, cap, name)
}

// ClaimCapability wraps the scopedKeeper's ClaimCapability function
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}
