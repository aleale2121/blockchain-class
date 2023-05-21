package main

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Tx struct {
	FromID string `json:"from"`
	ToID   string `json:"to"`
	Value  uint64 `json:"value"`
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {

	privateKey, err := crypto.LoadECDSA("zblock/accounts/kennedy.ecdsa")
	if err != nil {
		return fmt.Errorf("unable to load private key for node: %w", err)
	}

	tx := Tx{
		FromID: "9f332e3700d8fc2446eaf6d15034cf96e0c2745e40353deef032a5dbf1dfed93",
		ToID:   "Aaron",
		Value:  1000,
	}

	data, err := json.Marshal(tx)
	if err != nil {
		return fmt.Errorf("unable to marshal: %w", err)
	}

	v := crypto.Keccak256(data)

	sig, err := crypto.Sign(v, privateKey)

	if err != nil {
		return fmt.Errorf("unable to sign: %w", err)
	}

	fmt.Println(string(sig))
	fmt.Println(hexutil.Encode(sig))

	publicKey, err := crypto.SigToPub(v, sig)
	if err != nil {
		return fmt.Errorf("unable to pub: %w", err)
	}

	fmt.Println("PUB:", crypto.PubkeyToAddress(*publicKey).String())

	tx = Tx{
		FromID: "9f332e3700d8fc2446eaf6d15034cf96e0c2745e40353deef032a5dbf1dfed93",
		ToID:   "Frank",
		Value:  250,
	}

	data, err = json.Marshal(tx)
	if err != nil {
		return fmt.Errorf("unable to marshal: %w", err)
	}

	v2 := crypto.Keccak256(data)

	sig2, err := crypto.Sign(v2, privateKey)

	if err != nil {
		return fmt.Errorf("unable to sign: %w", err)
	}

	fmt.Println(string(sig2))
	fmt.Println(hexutil.Encode(sig2))

	publicKey, err = crypto.SigToPub(v2, sig2)
	if err != nil {
		return fmt.Errorf("unable to pub: %w", err)
	}

	fmt.Println("PUB:", crypto.PubkeyToAddress(*publicKey).String())
	// over the wire =======================================

	tx2 := Tx{
		FromID: "9f332e3700d8fc2446eaf6d15034cf96e0c2745e40353deef032a5dbf1dfed93",
		ToID:   "Frank",
		Value:  250,
	}

	data, err = json.Marshal(tx2)
	if err != nil {
		return fmt.Errorf("unable to marshal: %w", err)
	}

	v2 = crypto.Keccak256(data)

	publicKey, err = crypto.SigToPub(v2, sig2)
	if err != nil {
		return fmt.Errorf("unable to pub: %w", err)
	}

	fmt.Println("PUB:", crypto.PubkeyToAddress(*publicKey).String())
	return nil
}
