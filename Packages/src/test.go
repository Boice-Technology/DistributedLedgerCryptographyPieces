package main

import ("transactions";
		"math/big";
		"fmt";
		)
		
func main(){
	pri, _ := new(big.Int).SetString("1",10)
	x, y := transactions.EphemeralPublicKey(pri)
	fmt.Println(x,y)
}