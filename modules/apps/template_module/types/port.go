package types

import (
	"strconv"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	porttypes "github.com/cosmos/ibc-go/v2/modules/core/05-port/types"
)

const (
	// ControllerPortFormat is the expected port identifier format to which controller chains must conform
	// See (TODO: Link to spec when updated)
	ControllerPortFormat = "<app-version>.<controller-conn-seq>.<host-conn-seq>.<owner>"
)

// ParseControllerConnSequence attempts to parse the controller connection sequence from the provided port identifier
// The port identifier must match the controller chain format outlined in (TODO: link spec), otherwise an empty string is returned
func ParseControllerConnSequence(portID string) (uint64, error) {
	s := strings.Split(portID, Delimiter)
	if len(s) != 4 {
		return 0, sdkerrors.Wrap(porttypes.ErrInvalidPort, "failed to parse port identifier")
	}

	seq, err := strconv.ParseUint(s[1], 10, 64)
	if err != nil {
		return 0, sdkerrors.Wrapf(err, "failed to parse connection sequence (%s)", s[1])
	}

	return seq, nil
}

// ParseHostConnSequence attempts to parse the host connection sequence from the provided port identifier
// The port identifier must match the controller chain format outlined in (TODO: link spec), otherwise an empty string is returned
func ParseHostConnSequence(portID string) (uint64, error) {
	s := strings.Split(portID, Delimiter)
	if len(s) != 4 {
		return 0, sdkerrors.Wrap(porttypes.ErrInvalidPort, "failed to parse port identifier")
	}

	seq, err := strconv.ParseUint(s[2], 10, 64)
	if err != nil {
		return 0, sdkerrors.Wrapf(err, "failed to parse connection sequence (%s)", s[2])
	}

	return seq, nil
}
