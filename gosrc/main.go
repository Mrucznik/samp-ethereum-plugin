package main

import (
	"C"
	"crypto/ecdsa"

	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cast"
)

type RequestResponse struct {
	Data *RequestData `json:"data"`
}

type RequestData struct {
	Base     string `json:"base"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

//export GetCurrentETHPrice
func GetCurrentETHPrice() float32 {
	resp, err := http.Get("https://api.coinbase.com/v2/prices/ETH-PLN/sell")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	respStruct := RequestResponse{
		Data: &RequestData{},
	}
	err = json.Unmarshal(body, &respStruct)
	if err != nil {
		log.Fatal(err)
	}
	return cast.ToFloat32(respStruct.Data.Amount)
}

func getKeys() (string, string, string) {
	// generating a random private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	//convert it to bytes
	privateKeyBytes := crypto.FromECDSA(privateKey)

	//convert private key to a hexadecimal string and print. We strip of the 0x
	fmt.Println("Private key:", hexutil.Encode(privateKeyBytes)[2:])

	//public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	//Print public key.
	//We strip off the 0x and the first 2 characters 04 which is always the EC prefix and is not required.
	fmt.Println("Public key:", hexutil.Encode(publicKeyBytes)[4:]) // 9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	//Public address to send eth.
	//The public address is simply the Keccak-256 hash of the public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	return hexutil.Encode(privateKeyBytes)[2:], hexutil.Encode(publicKeyBytes)[4:], address
}

func run() {
	client, err := ethclient.Dial("http://localhost:8545") //for prod: "https://mainnet.infura.io"
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("We have a connection!")

	//get random wallet and keys
	fmt.Println("Generating keys...")
	_, _, address := getKeys()
	fmt.Println("Done.")

	//get balance
	fmt.Println("Getting account balance.")
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance: ", balance)

	ethPrice := GetCurrentETHPrice()
	fmt.Printf("Current ETH price in PLN: %.2fzł\n", ethPrice)
	fmt.Printf("Balance ib PLN: %.2fzł\n", cast.ToFloat32(balance)*ethPrice)

	fmt.Println("Bye.")
}

func main() {

}
