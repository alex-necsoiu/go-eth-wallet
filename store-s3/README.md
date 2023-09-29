# go-eth2-wallet-store-s3


Amazon S3-based store for the [Ethereum 2 wallet](https://github.com/alex-necsoiu/go-eth-wallet).


## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)


## Install

`go-eth2-wallet-store-s3` is a standard Go module which can be installed with:

```sh
go get github.com/alex-necsoiu/go-eth-wallet/store-s3
```

## Usage

In normal operation this module should not be used directly.  Instead, it should be configured to be used as part of [go-eth2-wallet](https://github.com/alex-necsoiu/go-eth-wallet).

The S3 store has the following options:

  - `region`: the Amazon S3 region in which the wallet is to be stored.  This can be any valid region string as per [the Amazon list](https://docs.aws.amazon.com/general/latest/gr/rande.html#apigateway_region), for example `ap-northeast-2` or `eu-north-1`
  - `id`: an ID that is used to differentiate multiple stores created by the same account.  If this is not configured an empty ID is used
  - `passphrase`: a key used to encrypt all data written to the store.  If this is not configured data is written to the store unencrypted (although wallet- and account-specific private information may be protected by their own passphrases)
  - `bucket`: the name of a bucket in which the store will place wallets.  If this is not configured it generates one based on the AWS credentials and ID
  - `path`: a path inside the bucket in which to place wallets.  If this is not configured it uses the root directory of the bucket
  - `endpoint`: a URL for an S3-compatible service, for example 'https://storage.googleapis.com` for Google Cloud Storage

When initiating a connection to Amazon S3 the Amazon credentials are required.  Details on how to make the credentials available to the store are available at [the Amazon S3 documentation](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#shared-credentials-file)

### Example

```go
package main

import (
	e2wallet "github.com/alex-necsoiu/go-eth-wallet"
	s3 "github.com/alex-necsoiu/go-eth-wallet/store-s3"
)

func main() {
    // Set up and use an encrypted store
    store, err := s3.New(s3.WithPassphrase([]byte("my secret")))
    if err != nil {
        panic(err)
    }
    e2wallet.UseStore(store)

    // Set up and use an encrypted store in the central Canada region
    store, err = s3.New(s3.WithPassphrase([]byte("my secret")), s3.WithRegion("ca-central-1"))
    if err != nil {
        panic(err)
    }
    e2wallet.UseStore(store)

    // Set up and use an encrypted store with a custom ID
    store, err = s3.New(s3.WithPassphrase([]byte("my secret")), s3.WithID([]byte("store 2")))
    if err != nil {
        panic(err)
    }
    e2wallet.UseStore(store)

    // Set up and use a store with a custom bucket and path
    store, err = s3.New(s3.WithBucket("my-store"), s3.WithPath("data/keystore"))
    if err != nil {
        panic(err)
    }
    e2wallet.UseStore(store)

    // Set up and use a store with non-dfeault credentials.
    store, err = s3.New(s3.WithCredentialsID("ABCDEF"), s3.WithCredentialsSecret("XXXXXXXXXXXX"))
    if err != nil {
        panic(err)
    }
    e2wallet.UseStore(store)
}
```

## Maintainers

Alex Necsoiu: [@alex-necsoiu](https://github.com/alex-necsoiu).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/alex-necsoiu/go-eth-wallet/tore-s3/issues).