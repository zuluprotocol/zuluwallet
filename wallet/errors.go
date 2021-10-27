package wallet

import (
	"errors"
	"fmt"
)

var (
	ErrIsolatedWalletCantGenerateKeyPairs = errors.New("isolated wallet can't generate key pairs")
	ErrInvalidMnemonic                    = errors.New("mnemonic is not valid")
	ErrPubKeyAlreadyTainted               = errors.New("public key is already tainted")
	ErrPubKeyIsTainted                    = errors.New("public key is tainted")
	ErrPubKeyNotTainted                   = errors.New("public key is not tainted")
	ErrPubKeyDoesNotExist                 = errors.New("public key does not exist")
	ErrWalletAlreadyExists                = errors.New("a wallet with the same name already exists")
	ErrWalletNotLoggedIn                  = errors.New("wallet is not logged in")
)

type UnsupportedWalletVersionError struct {
	UnsupportedVersion uint32
}

func NewUnsupportedWalletVersionError(v uint32) UnsupportedWalletVersionError {
	return UnsupportedWalletVersionError{
		UnsupportedVersion: v,
	}
}

func (e UnsupportedWalletVersionError) Error() string {
	return fmt.Sprintf("wallet with version %d isn't supported", e.UnsupportedVersion)
}
