package main

import "fmt"

func main() {
	fmt.Println("Arrays")

<<<<<<< HEAD
	var cryptoList [4]string

	cryptoList[0] = "BITCOIN"
	// cryptoList[1] = "ETHEREUM"
	// cryptoList[2] = "BINANCE"
	cryptoList[3] = "SOLANA"

	fmt.Println("CryptoList is \n", cryptoList)
	fmt.Println("Length of CryptoList Array is ", len(cryptoList)) //4
=======
	cryptoList := [5]string{"BITCOIN", "ETHEREUM", "BINANCE", "SOLANA", "TON"}

	fmt.Println("CryptoList is \n", cryptoList)
	fmt.Println("Length of CryptoList Array is ", len(cryptoList)) 

	fmt.Println("CryptoList: ")
	for i := 0; i < len(cryptoList); i++ {
		fmt.Println(cryptoList[i])
	}
>>>>>>> 8361048 (Golang practice)
}
