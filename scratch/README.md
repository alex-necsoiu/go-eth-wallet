# go-eth-wallet-store-scratch

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)

## Install

`go-eth-wallet-scratch` is a standard Go module which can be installed with:

```sh
go get github.com/alex-necsoiu/go-eth-wallet/scratch
```

## Usage

In normal operation this module should not be used directly.  Instead, it should be configured to be used as part of [go-eth2-wallet](https://github.com/alex-necsoiu/go-eth-wallet).

The scratch store has no options.

Note that because this store is ephemeral there should be no funds left in the wallet at the end of the runtime, or else they will become inaccessible due to the lack of private key
### Example

```go
package main

import (
	e2wallet "github.com/alex-necsoiu/go-eth-wallet"
	scratch "github.com/alex-necsoiu/go-eth-wallet/scratch"
)

func main() {
    // Set up and use an ephemeral store
    store := scratch.New()
    e2wallet.UseStore(store)

    // Use e2wallet operations as normal.
}
```

## Maintainers

Alex Necsoiu: [@alex-necsoiu](https://github.com/alex-necsoiu).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/alex-necsoiu/go-eth-wallet/issues).

