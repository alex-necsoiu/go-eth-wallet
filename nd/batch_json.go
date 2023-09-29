package nd

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	keystorev4 "github.com/alex-necsoiu/go-eth-wallet/keystore"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type batchEntryJSON struct {
	UUID   uuid.UUID `json:"uuid"`
	Name   string    `json:"name"`
	Pubkey string    `json:"pubkey"`
}

func (b *batchEntry) MarshalJSON() ([]byte, error) {
	res, err := json.Marshal(&batchEntryJSON{
		UUID:   b.id,
		Name:   b.name,
		Pubkey: fmt.Sprintf("%#x", b.pubkey),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal JSON")
	}

	return res, nil
}

func (b *batchEntry) UnmarshalJSON(input []byte) error {
	data := batchEntryJSON{}
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}
	b.id = data.UUID
	b.name = data.Name
	var err error
	b.pubkey, err = hex.DecodeString(strings.TrimPrefix(data.Pubkey, "0x"))
	if err != nil {
		return errors.Wrap(err, "invalid pubkey")
	}

	return nil
}

type batchJSON struct {
	Entries   []*batchEntry  `json:"entries"`
	Crypto    map[string]any `json:"crypto"`
	Encryptor string         `json:"encryptor"`
	Version   int            `json:"version"`
}

func (b *batch) MarshalJSON() ([]byte, error) {
	res, err := json.Marshal(&batchJSON{
		Entries:   b.entries,
		Crypto:    b.crypto,
		Encryptor: b.encryptor.String(),
		Version:   version,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal JSON")
	}

	return res, nil
}

func (b *batch) UnmarshalJSON(input []byte) error {
	data := batchJSON{}
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}
	if data.Version != version {
		return fmt.Errorf("unsupported version %d", data.Version)
	}
	b.entries = data.Entries
	switch data.Encryptor {
	case "keystorev4":
		b.encryptor = keystorev4.New()
	default:
		return fmt.Errorf("unsupported encryptor %s", data.Encryptor)
	}
	b.crypto = data.Crypto

	return nil
}
