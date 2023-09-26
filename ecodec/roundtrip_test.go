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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		key  []byte
		err  error
	}{
		{
			name: "Good1",
			data: _byteArray("0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"),
			key:  _byteArray("0102030405060708090a0b0c0d0e0f10"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			encryptedData, err := Encrypt(test.data, test.key)
			if test.err != nil {
				require.NotNil(t, err)
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				require.Nil(t, err)
				data, err := Decrypt(encryptedData, test.key)
				require.Nil(t, err)
				assert.Equal(t, test.data, data)
			}
		})
	}
}
