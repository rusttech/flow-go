package unittest

import (
	"encoding/hex"
	"fmt"

	"github.com/onflow/flow-go/crypto"
	"github.com/onflow/flow-go/crypto/hash"

	"github.com/onflow/cadence"

	"github.com/onflow/flow-go/model/flow"
)

// Used below with random service key
// privateKey := flow.AccountPrivateKey{
//	 PrivateKey: rootKey,
//	 SignAlgo:   crypto.ECDSAP256,
//	 HashAlgo:   hash.SHA2_256,
// }

const ServiceAccountPrivateKeyHex = "8ae3d0461cfed6d6f49bfc25fa899351c39d1bd21fdba8c87595b6c49bb4cc43"
const ServiceAccountPrivateKeySignAlgo = crypto.ECDSAP256
const ServiceAccountPrivateKeyHashAlgo = hash.SHA2_256

// Pre-calculated state commitment with root account with the above private key
const GenesisStateCommitmentHex = "1fa4f0ccd3b991627d2c95a7aca3294fbe7407f82711b55f6863dc03c970ce08"

var GenesisStateCommitment flow.StateCommitment

var GenesisTokenSupply = func() cadence.UFix64 {
	// value, err := cadence.NewUFix64("10000000000.0") // 10 billion
	value, err := cadence.NewUFix64("1000000000.0") // 1 billion
	if err != nil {
		panic(fmt.Errorf("invalid genesis token supply: %w", err))
	}
	return value
}()

var ServiceAccountPrivateKey flow.AccountPrivateKey
var ServiceAccountPublicKey flow.AccountPublicKey

func init() {
	var err error
	GenesisStateCommitmentBytes, err := hex.DecodeString(GenesisStateCommitmentHex)
	if err != nil {
		panic("error while hex decoding hardcoded state commitment")
	}
	GenesisStateCommitment, err = flow.ToStateCommitment(GenesisStateCommitmentBytes)
	if err != nil {
		panic("genesis state commitment size is invalid")
	}

	serviceAccountPrivateKeyBytes, err := hex.DecodeString(ServiceAccountPrivateKeyHex)
	if err != nil {
		panic("error while hex decoding hardcoded root key")
	}

	ServiceAccountPrivateKey.SignAlgo = ServiceAccountPrivateKeySignAlgo
	ServiceAccountPrivateKey.HashAlgo = ServiceAccountPrivateKeyHashAlgo
	ServiceAccountPrivateKey.PrivateKey, err = crypto.DecodePrivateKey(
		ServiceAccountPrivateKey.SignAlgo, serviceAccountPrivateKeyBytes)
	if err != nil {
		panic("error while decoding hardcoded root key bytes")
	}

	// Cannot import virtual machine, due to circular dependency. Just use the value of
	// fvm.AccountKeyWeightThreshold here
	ServiceAccountPublicKey = ServiceAccountPrivateKey.PublicKey(1000)
}

// this is done by printing the state commitment in TestBootstrapLedger test with different chain ID
func GenesisStateCommitmentByChainID(chainID flow.ChainID) flow.StateCommitment {
	commitString := genesisCommitHexByChainID(chainID)
	bytes, err := hex.DecodeString(commitString)
	if err != nil {
		panic("error while hex decoding hardcoded state commitment")
	}
	commit, err := flow.ToStateCommitment(bytes)
	if err != nil {
		panic("genesis state commitment size is invalid")
	}
	return commit
}

func genesisCommitHexByChainID(chainID flow.ChainID) string {
	if chainID == flow.Mainnet {
		return GenesisStateCommitmentHex
	}
	if chainID == flow.Testnet {
		return "a7afac1d7ee4de01d5ef7accd373579ecc59595fb475903fda1cdc98d8350a8a"
	}
	if chainID == flow.Sandboxnet {
		return "75e1ec7d2af323043aa9751ec2ae3fdf5ed9f445e1cb4349d14008d5edc892e2"
	}
	return "1b7736f113f0860f0df1ae5394ea1e926f1677c5944ca6a5a1680cd4f97420ec"
}
