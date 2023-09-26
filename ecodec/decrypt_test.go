// Copyright © 2019 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ecodec

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name          string
		encryptedData []byte
		key           []byte
		data          []byte
		err           error
	}{
		{
			name:          "DataShort",
			encryptedData: _byteArray("01"),
			key:           _byteArray("0102030405060708090a0b0c0d0e0f100102030405060708090a0b0c0d0e0f10"),
			err:           errors.New("encrypted data must be at least 81 bytes"),
		},
		{
			name:          "VersionIncorrect",
			encryptedData: _byteArray("0x02361a82b78fc468b2cc49ecabed25a96b028a013c92f053e507d180ae8b2258da5e2b4e7d90c413a9326a5a6f89e2b6af879a34a397dfd4148e08ac6fdaa07051c3e1a70dfa367f0811734a0a77fde4d19d5c0ec8c1d9879103526468c991a49ecacf9ea2f2a3e7e1939d1b3dae79828f"),
			key:           _byteArray("0102030405060708090a0b0c0d0e0f100102030405060708090a0b0c0d0e0f10"),
			err:           errors.New("unhandled version 0x02"),
		},
		{
			name:          "KeyInvalid",
			encryptedData: _byteArray("0x01361a82b78fc468b2cc49ecabed25a96b028a013c92f053e507d180ae8b2258da5e2b4e7d90c413a9326a5a6f89e2b6af879a34a397dfd4148e08ac6fdaa07051c3e1a70dfa367f0811734a0a77fde4d19d5c0ec8c1d9879103526468c991a49ecacf9ea2f2a3e7e1939d1b3dae79828f"),
			key:           _byteArray("0102030405060708090a0b0c0d0e0f100102030405060708090a0b0c0d0e0f11"),
			err:           errors.New("invalid key"),
		},
		{
			name:          "Good1",
			encryptedData: _byteArray("0x01361a82b78fc468b2cc49ecabed25a96b028a013c92f053e507d180ae8b2258da5e2b4e7d90c413a9326a5a6f89e2b6af879a34a397dfd4148e08ac6fdaa07051c3e1a70dfa367f0811734a0a77fde4d19d5c0ec8c1d9879103526468c991a49ecacf9ea2f2a3e7e1939d1b3dae79828f"),
			key:           _byteArray("0102030405060708090a0b0c0d0e0f100102030405060708090a0b0c0d0e0f10"),
			data:          _byteArray("0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, err := Decrypt(test.encryptedData, test.key)
			if test.err != nil {
				require.NotNil(t, err)
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				require.Nil(t, err)
				assert.Equal(t, test.data, data)
			}
		})
	}
}
