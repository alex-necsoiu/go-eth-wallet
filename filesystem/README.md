# go-eth-wallet-filesystem

Filesystem-based store for the [Ethereum 2 wallet](https://github.com/alex-necsoiu/go-eth-wallet).


## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)

## Install

`go-eth-wallet-filesystem` is a standard Go module which can be installed with:

```sh
go get github.com/alex-necsoiu/go-eth-wallet/filesystem
```

## Usage

In normal operation this module should not be used directly.  Instead, it should be configured to be used as part of [go-eth2-wallet](https://github.com/alex-necsoiu/go-eth-wallet).

The filesystem store has the following options:

  - `location`: the base directory in which to store wallets.  If this is not configured it defaults to a sensible operating system-specific values:
    - for Linux: $HOME/.config/ethereum2/wallets
    - for OSX: $HOME/Library/Application Support/ethereum2/wallets
    - for Windows: %APPDATA%\ethereum2\wallets
  - `passphrase`: a key used to encrypt all data written to the store.  If this is not configured data is written to the store unencrypted (although wallet- and account-specific private information may be protected by their own passphrases)

### Example

```go
package main

import (
	e2wallet "github.com/alex-necsoiu/go-eth-wallet"
	filesystem "github.com/alex-necsoiu/go-eth-wallet/filesystem"
)

func main() {
    // Set up and use a simple store
    store := filesystem.New()
    e2wallet.UseStore(store)

    // Set up and use an encrypted store
    store := filesystem.New(filesystem.WithPassphrase([]byte("my secret")))
    e2wallet.UseStore(store)

    // Set up and use an encrypted store at a custom location
    store := filesystem.New(filesystem.WithPassphrase([]byte("my secret")), filesystem.WithLocation("/home/user/wallets"))
    e2wallet.UseStore(store)

    // Use e2wallet operations as normal.
}
```

## Maintainers

Alex Necsoiu: [@alex-necsoiu](https://github.com/alex-necsoiu).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/alex-necsoiu/go-eth-wallet/issues).

