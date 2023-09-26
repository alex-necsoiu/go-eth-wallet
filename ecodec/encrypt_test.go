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
	"encoding/hex"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func _byteArray(input string) []byte {
	res, err := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	if err != nil {
		panic(err)
	}
	return res
}

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name          string
		data          []byte
		key           []byte
		err           error
		encryptedData []byte
	}{
		{
			name: "DataNil",
			err:  errors.New("no data"),
		},
		{
			name: "DataEmpty",
			data: _byteArray(""),
			err:  errors.New("no data"),
		},
		{
			name: "DataShort",
			data: _byteArray("0102030405060708090a0b0c0d0e0f"),
			key:  _byteArray("0102030405060708090a0b0c0d0e0f10"),
			err:  errors.New("data must be at least 16 bytes"),
		},
		{
			name: "KeyNil",
			data: _byteArray("0102030405060708"),
			err:  errors.New("no key"),
		},
		{
			name: "KeyEmpty",
			data: _byteArray("0102030405060708"),
			key:  _byteArray(""),
			err:  errors.New("no key"),
		},
		{
			name: "Good1",
			data: _byteArray("0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"),
			key:  _byteArray("0102030405060708090a0b0c0d0e0f10"),
		},
		{
			name: "Good2",
			data: _byteArray("0102030405060708090a0b0c0d0e0f10"),
			key:  _byteArray("0102030405060708090a0b0c0d0e0f10"),
		},
		{
			name: "Good3",
			data: _byteArray("0102030405060708090a0b0c0d0e0f1011"),
			key:  _byteArray("0102030405060708090a0b0c0d0e0f10"),
		},
		{
			name: "Good4",
			data: _byteArray("0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"),
			key:  _byteArray("0102030405060708090a0b0c0d0e0f100102030405060708090a0b0c0d0e0f10"),
		},
		{
			name: "Good5",
			data: _byteArray("0102030405060708090a0b0c0d0e0f10"),
			key:  _byteArray("01"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := Encrypt(test.data, test.key)
			if test.err != nil {
				require.NotNil(t, err)
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				require.Nil(t, err)
				// Can't test result due to random IV; results checked in round-trip testing.
			}
		})
	}
}
