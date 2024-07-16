package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var cryptoList [4]string

	cryptoList[0] = "BITCOIN"
	// cryptoList[1] = "ETHEREUM"
	// cryptoList[2] = "BINANCE"
	cryptoList[3] = "SOLANA"

	fmt.Println("CryptoList is \n", cryptoList)
	fmt.Println("Length of CryptoList Array is ", len(cryptoList)) //4
}
