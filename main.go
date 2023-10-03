package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/text/unicode/norm"

	keystorev4 "github.com/alex-necsoiu/go-eth-wallet/keystore"
	"github.com/alex-necsoiu/go-eth-wallet/scratch"
	"github.com/alex-necsoiu/go-eth-wallet/util"
	"github.com/alex-necsoiu/go-eth-wallet/wallet"
	wtypes "github.com/alex-necsoiu/go-eth-wallet/wallet/types"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"

	"github.com/google/uuid"
)

const (
	S3_BUCKET_NAME = "your-s3-bucket-name"
	SECRET_NAME    = "your-secret-name-in-aws-secrets-manager"
	REGION         = "your-aws-region"
)
const hdKeysPath = "m/12381/3600/0/1"

func main() {
	Id := uuid.New()

	store := scratch.New()
	encryptor := keystorev4.New()

	fmt.Println("User ID: ", Id.String())

	encryption := wallet.WithEncryptor(encryptor)
	storeOpt := wallet.WithStore(store)

	mnemonic := generateMnemonic()
	seed := bip39.NewSeed(mnemonic, hdKeysPath)
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		fmt.Printf("Error Master Key: %v\n", err)
	}
	// components := strings.Split(hdKeysPath, "/")
	// path := make([]uint32, len(components)-1)
	// for i := 0; i < len(components)-1; i++ {
	// 	if strings.HasSuffix(components[i+1], "'") {
	// 		unhardened, err := strconv.ParseInt(components[i+1][:len(components[i+1])-1], 10, 0)
	// 		if err != nil {
	// 			fmt.Printf("Error Master Key: %v\n", err)
	// 		}
	// 		path[i] = 0x80000000 + uint32(unhardened)
	// 	} else {
	// 		unhardened, err := strconv.ParseInt(components[i+1], 10, 0)
	// 		if err != nil {
	// 			fmt.Printf("Error Master Key: %v\n", err)
	// 		}
	// 		path[i] = uint32(unhardened)
	// 	}
	// }
	// childKey := masterKey
	// for i := range path {
	// 	childKey, err = childKey.NewChildKey(path[i])
	// }

	// key, err := crypto.ToECDSA(childKey.Key)
	// // masterKey, err := util.DeriveMasterSK([]byte(mnemonic))
	// // if err != nil {
	// // 	fmt.Printf("Error Private Key: %v\n", err)
	// // }

	fmt.Printf("Master Key: %v\n", masterKey)
	// fmt.Printf("Private key:  %v\n", key.D)
	// fmt.Printf("Public key: %v\n", hex.EncodeToString(crypto.FromECDSAPub(&key.PublicKey)))
	// fmt.Printf("Ethereum address: %v\n", crypto.PubkeyToAddress(key.PublicKey).Hex())

	sk, err := util.PrivateKeyFromSeedAndPath(seed, hdKeysPath)
	if err != nil {
		fmt.Printf("Error Private Key: %v\n", err)
	}

	// fmt.Printf("Private Key fromSeed: %v\n", sk)

	passPhrase := wallet.WithPassphrase([]byte("secret"))
	walletType := wallet.WithType("hd")

	newWallet, err := wallet.CreateWallet(Id.String(),

		wallet.Option(encryption),
		wallet.Option(storeOpt),
		wallet.Option(passPhrase),
		wallet.Option(walletType),
		wallet.WithSeed(seed),
	)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	ctx := context.Background()
	// fmt.Printf("Wallet: %s\n", newWallet.ID().String())

	// // fmt.Printf("Wallet Name: %s\n", newWallet.Name())
	// // fmt.Printf("Wallet Acc: %s\n", newWallet.Accounts(ctx))
	// // fmt.Printf("Wallet Type: %s\n", newWallet.Type())

	locker := newWallet.(wtypes.WalletLocker)

	err = locker.Unlock(context.Background(), []byte("secret"))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	_, err = newWallet.(wtypes.WalletAccountCreator).CreateAccount(ctx, "name1", sk.Marshal())
	if err != nil {
		fmt.Printf("Error Creating the account: %s\n", err)
	}

	// // _, err = newWallet.(wtypes.WalletAccountCreator).CreateAccount(ctx, "name2", sk.Marshal())
	// // if err != nil {
	// // 	fmt.Printf("Error: %s\n", err)
	// // }
	// // fmt.Printf("Wallet Name: %s\n", acc.Name())
	// // fmt.Printf("Wallet Name: %s\n", acc2.Name())
	// fmt.Printf("Wallet: %s\n", newWallet.Accounts(ctx))

}

// func createUserWallet(userID string, mnemonic string) (string, error) {
// 	// Generate a new Ethereum wallet
// 	newWallet, err := wallet.CreateWallet(userID)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Fetch the admin secret from AWS Secrets Manager
// 	adminSecret, err := getSecretFromAWS()
// 	if err != nil {
// 		return "", err
// 	}

// 	// Encrypt the mnemonic using the admin secret
// 	encryptedMnemonic, err := encryptMnemonic(mnemonic, adminSecret)
// 	if err != nil {
// 		return "", err
// 	}

// 	fmt.Println("Encrypted mnemonic: ", encryptedMnemonic)
// 	// // Store the encrypted mnemonic in S3
// 	// err = storeInS3(encryptedMnemonic, wallet.Address.Hex())
// 	// if err != nil {
// 	// 	return "", err
// 	// }

// 	// // Generate a UUID for the user
// 	// userID := uuid.New().String()

// 	return userID, nil
// }

// func getSecretFromAWS() (string, error) {
// 	sess := session.Must(session.NewSession(&aws.Config{
// 		Region: aws.String(REGION),
// 	}))

// 	svc := secretsmanager.New(sess)
// 	input := &secretsmanager.GetSecretValueInput{
// 		SecretId: aws.String(SECRET_NAME),
// 	}

// 	result, err := svc.GetSecretValue(input)
// 	if err != nil {
// 		return "", err
// 	}

// 	return *result.SecretString, nil
// }

// func encryptMnemonic(mnemonic, secret string) (string, error) {
// 	key := crypto.Keccak256([]byte(secret))
// 	encryptedData, err := crypto.AES256EncryptWithECB([]byte(mnemonic), key)
// 	if err != nil {
// 		return "", err
// 	}

// 	return hex.EncodeToString(encryptedData), nil
// }

// func storeInS3(data, filename string) error {
// 	sess := session.Must(session.NewSession(&aws.Config{
// 		Region: aws.String(REGION),
// 	}))

// 	svc := s3.New(sess)

// 	_, err := svc.PutObject(&s3.PutObjectInput{
// 		Bucket: aws.String(S3_BUCKET_NAME),
// 		Key:    aws.String(filename),
// 		Body:   aws.ReadSeekCloser(strings.NewReader(data)),
// 	})

// 	return err
// }

// generateMnemonic creates a new mnemonic.
func generateMnemonic() string {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return mnemonic
}
func _byte(input string) []byte {
	res, err := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	if err != nil {
		panic(err)
	}
	return res
}

// expandMnmenonic expands mnemonics from their 4-letter versions.
func expandMnemonic(input string) string {
	wordList := bip39.GetWordList()
	truncatedWords := make(map[string]string, len(wordList))
	for _, word := range wordList {
		if len(word) > 4 {
			truncatedWords[firstFour(word)] = word
		}
	}
	mnemonicWords := strings.Split(input, " ")
	for i := range mnemonicWords {
		if fullWord, exists := truncatedWords[norm.NFKC.String(mnemonicWords[i])]; exists {
			mnemonicWords[i] = fullWord
		}
	}
	return strings.Join(mnemonicWords, " ")
}

// firstFour provides the first four letters for a potentially longer word.
func firstFour(s string) string {
	// Use NFKC here for composition, to avoid accents counting as their own characters.
	s = norm.NFKC.String(s)
	r := []rune(s)
	if len(r) > 4 {
		return string(r[:4])
	}
	return s
}
