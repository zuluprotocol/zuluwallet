package wallet

import (
	"errors"
	"fmt"
)

var (
	ErrIsolatedWalletCantGenerateKeyPairs = errors.New("isolated wallet can't generate key pairs")
	ErrIsolatedWalletCantGetMasterKeyPair = errors.New("isolated wallet can't get master key pair")
	ErrInvalidMnemonic                    = errors.New("mnemonic is not valid")
	ErrPubKeyAlreadyTainted               = errors.New("public key is already tainted")
	ErrPubKeyIsTainted                    = errors.New("public key is tainted")
	ErrPubKeyNotTainted                   = errors.New("public key is not tainted")
	ErrPubKeyDoesNotExist                 = errors.New("public key does not exist")
	ErrWalletAlreadyExists                = errors.New("a wallet with the same name already exists")
	ErrWalletDoesNotExists                = errors.New("wallet does not exist")
	ErrWalletNotLoggedIn                  = errors.New("wallet is not logged in")
	ErrWrongPassphrase                    = errors.New("wrong passphrase")
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
