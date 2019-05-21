package main

import ("walletcrypto";
		"fmt";
		)

func main(){
	str := walletcrypto.Seed("hello","hey")
	fmt.Println(str)
	fmt.Println(len(str))
}		