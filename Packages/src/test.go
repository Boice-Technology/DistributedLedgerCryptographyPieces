package main

import ("fmt";
		"walletcrypto";
		)
		
func main(){
	str1, str2 := walletcrypto.MasterNode("hello")
	fmt.Println(str1,str2)
	fmt.Println(len(str1),len(str2))
}