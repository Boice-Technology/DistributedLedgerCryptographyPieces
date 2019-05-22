package main

import ("fmt";
		"walletcrypto";
		)
		
func main(){
	str1, str2 := walletcrypto.PrivateChildKey("abc","hii",0)
	fmt.Println(str1)
	fmt.Println(str2)
}