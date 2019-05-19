package main

import ("fmt";
		"math/big";
		"eckg";
		)

func main(){
	var pri1,_ = new(big.Int).SetString("3",10)
	var pri2,_ = new(big.Int).SetString("3",10)
	fmt.Println(eckg.UncompPublicKey(pri1))
	fmt.Println(eckg.CompPublicKey(pri2))
}