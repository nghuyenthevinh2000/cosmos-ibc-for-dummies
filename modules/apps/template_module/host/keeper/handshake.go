package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	icptypes "github.com/cosmos/ibc-go/v2/modules/apps/interchain-price/types"
	connectiontypes "github.com/cosmos/ibc-go/v2/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v2/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v2/modules/core/24-host"
)

// OnChanOpenTry performs basic validation of the ICP channel
// and registers a new interchain account (if it doesn't exist).
func (k Keeper) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version,
	counterpartyVersion string,
) error {
	//============ CHECKING GROUP ============
	//1. check order of how packets are sent
	if order != channeltypes.ORDERED {
		return sdkerrors.Wrapf(channeltypes.ErrInvalidChannelOrdering, "expected %s channel, got %s", channeltypes.ORDERED, order)
	}

	//2. check port ID (host's port ID) that controller wants to connect to is correct
	if portID != icptypes.PortID {
		return sdkerrors.Wrapf(porttypes.ErrInvalidPort, "expected %s, got %s", icptypes.PortID, portID)
	}

	//3. check correct format and get connection sequence of both host and controller party
	connSequence, err := icptypes.ParseHostConnSequence(counterparty.PortId)
	if err != nil {
		return sdkerrors.Wrapf(err, "expected format %s, got %s", icptypes.ControllerPortFormat, counterparty.PortId)
	}

	counterpartyConnSequence, err := icptypes.ParseControllerConnSequence(counterparty.PortId)
	if err != nil {
		return sdkerrors.Wrapf(err, "expected format %s, got %s", icptypes.ControllerPortFormat, counterparty.PortId)
	}

	//4. validate if connSequence and counterpartyConnSequence is valid and not mismatched
	if err := k.validateControllerPortParams(ctx, channelID, portID, connSequence, counterpartyConnSequence); err != nil {
		return sdkerrors.Wrapf(err, "failed to validate controller port %s", counterparty.PortId)
	}

	//5. check version of both host and controller party
	if err := icptypes.ValidateVersion(version); err != nil {
		return sdkerrors.Wrap(err, "version validation failed")
	}

	if counterpartyVersion != icptypes.VersionPrefix {
		return sdkerrors.Wrapf(icptypes.ErrInvalidVersion, "expected %s, got %s", icptypes.VersionPrefix, version)
	}

	//============ END CHECKING GROUP ============

	// On the host chain the capability may only be claimed during the OnChanOpenTry
	// The capability being claimed in OpenInit is for a controller chain (the port is different)
	if err := k.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
		return sdkerrors.Wrapf(err, "failed to claim capability for channel %s on port %s", channelID, portID)
	}

	return nil
}

// validateControllerPortParams asserts the provided connection sequence and counterparty connection sequence
// match that of the associated connection stored in state
func (k Keeper) validateControllerPortParams(ctx sdk.Context, channelID, portID string, connectionSeq, counterpartyConnectionSeq uint64) error {
	channel, found := k.channelKeeper.GetChannel(ctx, portID, channelID)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID %s channel ID %s", portID, channelID)
	}

	counterpartyHops, found := k.channelKeeper.CounterpartyHops(ctx, channel)
	if !found {
		return sdkerrors.Wrap(connectiontypes.ErrConnectionNotFound, channel.ConnectionHops[0])
	}

	connSeq, err := connectiontypes.ParseConnectionSequence(channel.ConnectionHops[0])
	if err != nil {
		return sdkerrors.Wrapf(err, "failed to parse connection sequence %s", channel.ConnectionHops[0])
	}

	counterpartyConnSeq, err := connectiontypes.ParseConnectionSequence(counterpartyHops[0])
	if err != nil {
		return sdkerrors.Wrapf(err, "failed to parse counterparty connection sequence %s", counterpartyHops[0])
	}

	if connSeq != connectionSeq {
		return sdkerrors.Wrapf(connectiontypes.ErrInvalidConnection, "sequence mismatch, expected %d, got %d", connSeq, connectionSeq)
	}

	if counterpartyConnSeq != counterpartyConnectionSeq {
		return sdkerrors.Wrapf(connectiontypes.ErrInvalidConnection, "counterparty sequence mismatch, expected %d, got %d", counterpartyConnSeq, counterpartyConnectionSeq)
	}

	return nil
}
