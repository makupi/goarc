package decoder

import (
	"errors"
	"fmt"
	"github.com/multiformats/go-multicodec"
	"regexp"
	"strings"

	"github.com/algorand/go-algorand-sdk/types"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

var (
	ErrUnknownSpec      = errors.New("unsupported template-ipfs spec")
	ErrUnsupportedField = errors.New("unsupported ipfscid field, only reserve is currently supported")
	ErrUnsupportedCodec = errors.New("unknown multicodec type in ipfscid spec")
	ErrUnsupportedHash  = errors.New("unknown hash type in ipfscid spec")
	ErrInvalidV0        = errors.New("cid v0 must always be dag-pb and sha2-256 codec/hash type")
	ErrHashEncoding     = errors.New("error encoding new hash")
	templateIPFSRegexp  = regexp.MustCompile(`template-ipfs://{ipfscid:(?P<version>[01]):(?P<codec>[a-z0-9\-]+):(?P<field>[a-z0-9\-]+):(?P<hash>[a-z0-9\-]+)}`)
)

func ParseASAUrl(asaUrl string, addr string) (string, error) {
	reserveAddress, err := types.DecodeAddress(addr)
	if err != nil {
		return "", err
	}
	matches := templateIPFSRegexp.FindStringSubmatch(asaUrl)
	if matches == nil {
		if strings.HasPrefix(asaUrl, "template-ipfs://") {
			return "", ErrUnknownSpec
		}
		return asaUrl, nil
	}
	if matches[templateIPFSRegexp.SubexpIndex("field")] != "reserve" {
		return "", ErrUnsupportedField
	}
	var (
		codec         multicodec.Code
		multihashType uint64
		hash          []byte
		cidResult     cid.Cid
	)
	if err = codec.Set(matches[templateIPFSRegexp.SubexpIndex("codec")]); err != nil {
		return "", ErrUnsupportedCodec
	}
	multihashType = multihash.Names[matches[templateIPFSRegexp.SubexpIndex("hash")]]
	if multihashType == 0 {
		return "", ErrUnsupportedHash
	}

	hash, err = multihash.Encode(reserveAddress[:], multihashType)
	if err != nil {
		return "", ErrHashEncoding
	}
	if matches[templateIPFSRegexp.SubexpIndex("version")] == "0" {
		if codec != multicodec.DagPb {
			return "", ErrInvalidV0
		}
		if multihashType != multihash.SHA2_256 {
			return "", ErrInvalidV0
		}
		cidResult = cid.NewCidV0(hash)
	} else {
		cidResult = cid.NewCidV1(uint64(codec), hash)
	}
	return fmt.Sprintf("ipfs://%s", strings.ReplaceAll(asaUrl, matches[0], cidResult.String())), nil
}
