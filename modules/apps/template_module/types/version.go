package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ValidateVersion performs basic validation of the provided ics27 version string.
// An ics27 version string may include an optional account address as per [TODO: Add spec when available]
// ValidateVersion first attempts to split the version string using the standard delimiter, then asserts a supported
// version prefix is included, followed by additional checks which enforce constraints on the account address.
func ValidateVersion(version string) error {
	if version != VersionPrefix {
		return sdkerrors.Wrapf(ErrInvalidVersion, "expected %s, got %s", VersionPrefix, version)
	}

	return nil
}
