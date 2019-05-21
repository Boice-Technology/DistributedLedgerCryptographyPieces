package main

import ("fmt";
		"walletcrypto";
		)
	
func main(){
	fmt.Println(walletcrypto.Mnemonic("0123456789abcdef0123456789abcdea"))
}