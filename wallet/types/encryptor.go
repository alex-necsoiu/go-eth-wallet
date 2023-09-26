// Copyright 2019 - 2023 Weald Technology Trading.
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

package types

// Encryptor is the interface for encrypting and decrypting sensitive information in wallets.
type Encryptor interface {
	// Name() provides the name of the encryptor.
	Name() string

	// Version() provides the version of the encryptor.
	Version() uint

	// String provides a string value for the encryptor.
	String() string

	// Encrypt encrypts a byte array with its encryption mechanism and key.
	Encrypt(data []byte, key string) (map[string]any, error)

	// Decrypt encrypts a byte array with its encryption mechanism and key.
	Decrypt(data map[string]any, key string) ([]byte, error)
}
