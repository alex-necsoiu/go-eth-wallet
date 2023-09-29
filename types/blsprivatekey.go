package types

import (
	bls "github.com/herumi/bls-eth-go-binary/bls"
	"github.com/pkg/errors"
)

// BLSPrivateKey is a private key in Ethereum 2.
// It is a point on the BLS12-381 curve.
type BLSPrivateKey struct {
	key bls.SecretKey
}

// BLSPrivateKeyFromBytes creates a BLS private key from a byte slice.
func BLSPrivateKeyFromBytes(priv []byte) (*BLSPrivateKey, error) {
	if len(priv) != 32 {
		return nil, errors.New("private key must be 32 bytes")
	}
	var sec bls.SecretKey
	if err := sec.Deserialize(priv); err != nil {
		return nil, errors.Wrap(err, "invalid private key")
	}
	return &BLSPrivateKey{key: sec}, nil
}

// GenerateBLSPrivateKey generates a random BLS private key.
func GenerateBLSPrivateKey() (*BLSPrivateKey, error) {
	var sec bls.SecretKey
	sec.SetByCSPRNG()
	return &BLSPrivateKey{key: sec}, nil
}

// Marshal a secret key into a byte slice.
func (p *BLSPrivateKey) Marshal() []byte {
	return p.key.Serialize()
}

// PublicKey obtains the public key corresponding to the BLS secret key.
func (p *BLSPrivateKey) PublicKey() PublicKey {
	return &BLSPublicKey{key: p.key.GetPublicKey()}
}

// Sign a message using a secret key.
func (p *BLSPrivateKey) Sign(msg []byte) Signature {
	sig := p.key.SignHash(msg)
	return &BLSSignature{sig}
}
