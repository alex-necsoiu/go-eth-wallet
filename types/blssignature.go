package types

import (
	bls "github.com/herumi/bls-eth-go-binary/bls"
	"github.com/pkg/errors"
)

// BLSSignature is a BLS signature.
type BLSSignature struct {
	sig *bls.Sign
}

// BLSSignatureFromBytes creates a BLS signature from a byte slice.
func BLSSignatureFromBytes(data []byte) (Signature, error) {
	var sig bls.Sign
	if err := sig.Deserialize(data); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize signature")
	}
	return &BLSSignature{sig: &sig}, nil
}

// BLSSignatureFromSig creates a BLS signature from an existing signature.
func BLSSignatureFromSig(sig bls.Sign) (Signature, error) {
	return &BLSSignature{sig: &sig}, nil
}

// Verify a bls signature given a public key and a message.
func (s *BLSSignature) Verify(msg []byte, pubKey PublicKey) bool {
	return s.sig.VerifyByte(pubKey.(*BLSPublicKey).key, msg)
}

// VerifyAggregate verifies each public key against its respective message.
// Note: this is vulnerable to a rogue public-key attack.
func (s *BLSSignature) VerifyAggregate(msgs [][]byte, pubKeys []PublicKey) bool {
	if len(pubKeys) == 0 {
		return false
	}
	keys := make([]bls.PublicKey, len(pubKeys))
	for i, v := range pubKeys {
		keys[i] = *v.(*BLSPublicKey).key
	}
	return s.sig.VerifyAggregateHashes(keys, msgs)
}

// VerifyAggregateCommon verifies each public key against a single message.
// Note: this is vulnerable to a rogue public-key attack.
func (s *BLSSignature) VerifyAggregateCommon(msg []byte, pubKeys []PublicKey) bool {
	if len(pubKeys) == 0 {
		return false
	}
	keys := make([]bls.PublicKey, len(pubKeys))
	for i, v := range pubKeys {
		keys[i] = *v.(*BLSPublicKey).key
	}
	return s.sig.FastAggregateVerify(keys, msg)
}

// Marshal a signature into a byte slice.
func (s *BLSSignature) Marshal() []byte {
	return s.sig.Serialize()
}

// AggregateSignatures aggregates signatures.
func AggregateSignatures(sigs []Signature) *BLSSignature {
	if len(sigs) == 0 {
		return nil
	}
	aggSig := &bls.Sign{}
	//#nosec G104
	_ = aggSig.Deserialize(sigs[0].(*BLSSignature).Marshal())
	for _, sig := range sigs[1:] {
		aggSig.Add(sig.(*BLSSignature).sig)
	}
	return &BLSSignature{sig: aggSig}
}
