package tx

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"

	ojoapp "github.com/ojo-network/ojo/app"
)

const (
	keyringPassphrase = "testpassphrase"
	keyringAppName    = "testnet"
)

// CreateAccountFromMnemonic creates a new account from a mnemonic
func CreateAccountFromMnemonic(name, mnemonic string) (*keyring.Record, keyring.Keyring, error) {
	encodingConfig := ojoapp.MakeEncodingConfig()
	cdc := encodingConfig.Codec

	kb, err := keyring.New(keyringAppName, keyring.BackendMemory, "", nil, cdc)
	if err != nil {
		return nil, nil, err
	}

	keyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), keyringAlgos)
	if err != nil {
		return nil, nil, err
	}

	account, err := kb.NewAccount(name, mnemonic, "", sdk.FullFundraiserPath, algo)
	if err != nil {
		return nil, nil, err
	}

	return account, kb, nil
}

func CreateMnemonic() (string, error) {
	entropySeed, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}

	mnemonic, err := bip39.NewMnemonic(entropySeed)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}
