package types

import (
	"context"

	e2types "github.com/alex-necsoiu/go-eth-wallet/types"
	"github.com/google/uuid"
)

// AccountIDProvider is the interface for accounts that can provide an ID.
type AccountIDProvider interface {
	// ID provides the ID for the account.
	ID() uuid.UUID
}

// AccountNameProvider is the interface for accounts that can provide a name.
type AccountNameProvider interface {
	// Name provides the name for the account.
	Name() string
}

// AccountPublicKeyProvider is the interface for accounts that can provide a public key.
type AccountPublicKeyProvider interface {
	// PublicKey provides the public key for the account.
	PublicKey() e2types.PublicKey
}

// AccountPathProvider is the interface for accounts that can provide a path.
type AccountPathProvider interface {
	// Path provides the path for the account.
	Path() string
}

// AccountWalletProvider is the interface for accounts that can provide their containing wallet.
type AccountWalletProvider interface {
	// Wallet provides the wallet for this account.
	Wallet() Wallet
}

// AccountLocker is the interface for accounts that can be locked and unlocked.
type AccountLocker interface {
	// Lock locks the account.  A locked account cannot sign.
	Lock(ctx context.Context) error

	// Unlock unlocks the account.  An unlocked account can sign.
	Unlock(ctx context.Context, passphrase []byte) error

	// IsUnlocked returns true if the account is unlocked.
	IsUnlocked(ctx context.Context) (bool, error)
}

// AccountSigner is the interface for accounts that can sign generic data.
type AccountSigner interface {
	// Sign signs data with the account.
	Sign(ctx context.Context, data []byte) (e2types.Signature, error)
}

// AccountProtectingSigner is the interface for accounts that sign with protection.
type AccountProtectingSigner interface {
	// SignGeneric signs a generic root with protection.
	SignGeneric(ctx context.Context, data []byte, domain []byte) (e2types.Signature, error)

	// SignBeaconProposal signs a beacon proposal with protection.
	SignBeaconProposal(ctx context.Context,
		slot uint64,
		proposerIndex uint64,
		parentRoot []byte,
		stateRoot []byte,
		bodyRoot []byte,
		domain []byte) (e2types.Signature, error)

	// SignBeaconAttestation signs a beacon attestation with protection.
	SignBeaconAttestation(ctx context.Context,
		slot uint64,
		committeeIndex uint64,
		blockRoot []byte,
		sourceEpoch uint64,
		sourceRoot []byte,
		targetEpoch uint64,
		targetRoot []byte,
		domain []byte) (e2types.Signature, error)
}

// AccountProtectingMultiSigner is the interface for accounts that sign multiple requests with protection.
type AccountProtectingMultiSigner interface {
	// SignBeaconAttestations signs multiple beacon attestations with protection.
	SignBeaconAttestations(ctx context.Context,
		slot uint64,
		accounts []Account,
		committeeIndices []uint64,
		blockRoot []byte,
		sourceEpoch uint64,
		sourceRoot []byte,
		targetEpoch uint64,
		targetRoot []byte,
		domain []byte) ([]e2types.Signature, error)
}

// AccountCompositePublicKeyProvider is the interface for accounts that can provide a composite public key.
type AccountCompositePublicKeyProvider interface {
	// CompositePublicKey provides the composite public key for the account.
	CompositePublicKey() e2types.PublicKey
}

// AccountSigningThresholdProvider is the interface for accounts that can provide a signing threshold.
type AccountSigningThresholdProvider interface {
	// SigningThreshold provides the threshold to make a valid composite signature.
	SigningThreshold() uint32
}

// AccountVerificationVectorProvider is the interface for accounts that can provide a verification vector.
type AccountVerificationVectorProvider interface {
	// VerificationVector provides the composite verification vector for regeneration.
	VerificationVector() []e2types.PublicKey
}

// AccountParticipantsProvider is the interface for accounts that can participate in distributed operations.
type AccountParticipantsProvider interface {
	// Participants provides the participants that hold the composite key.
	Participants() map[uint64]string
}

// AccountPrivateKeyProvider is the interface for accounts that can provide a private key.
type AccountPrivateKeyProvider interface {
	// PrivateKey provides the private key for the account.
	PrivateKey(ctx context.Context) (e2types.PrivateKey, error)
}

// Account is a generic interface for accounts, providing minimal required functionality.
type Account interface {
	AccountIDProvider
	AccountNameProvider
	AccountPublicKeyProvider
}

// DistributedAccount is generic interface for distributed accounts, providing minimal required functionality.
type DistributedAccount interface {
	AccountIDProvider
	AccountNameProvider
	AccountCompositePublicKeyProvider
	AccountSigningThresholdProvider
	AccountParticipantsProvider
}

// ShardedAccount is generic interface for sharded accounts, providing minimal required functionality.
type ShardedAccount interface {
	AccountIDProvider
	AccountNameProvider
	AccountCompositePublicKeyProvider
	AccountSigningThresholdProvider
	AccountParticipantsProvider
}

// AccountMetadataProvider provides metadata for an account.  It is used for various accounting purposes,
// for example to ensure that no two accounts with the same name exist in a single wallet.
type AccountMetadataProvider interface {
	// WalletID provides the ID for the wallet.
	WalletID() uuid.UUID

	// ID provides the ID for the account.
	ID() uuid.UUID

	// Name provides the name for the account.
	Name() string
}
