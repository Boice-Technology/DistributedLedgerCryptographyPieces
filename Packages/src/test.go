package main

import ("fmt";
		"walletcrypto";
		"eckg";
		"math/big";
		)
	
func main() {
	priKey,_ := new(big.Int).SetString("31",10)
	pubKey := eckg.UncompPublicKey(priKey)
	str1, str2 := walletcrypto.PublicChildKey(pubKey,"89",0)
	fmt.Println(str1)
	fmt.Println(str2)
}