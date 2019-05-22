package main

import ("walletcrypto";
		"fmt";
		)
		
func main(){
	str1, str2 := walletcrypto.HardPrivateChildKey("31","abcd",0)
	fmt.Println(str1)
	fmt.Println(str2)
}
