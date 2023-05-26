package encoder

import (
	"fmt"

	"github.com/algorand/go-algorand-sdk/types"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

func ReserveAddressFromCID(toEncode string) (string, error) {
	cidToEncode := cid.MustParse(toEncode)
	decodedMultiHash, err := multihash.Decode(cidToEncode.Hash())
	if err != nil {
		return "", fmt.Errorf("failed to decode ipfs cid: %w", err)
	}
	return types.EncodeAddress(decodedMultiHash.Digest)
}
