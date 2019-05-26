package main

import ("fmt";
		"transactions";
		"math/big";
		)
		
func main(){
	pri,_ := new(big.Int).SetString("3",10)
	ephePri,_ := new(big.Int).SetString("1",10)
	ephePri1,_ := new(big.Int).SetString("1",10)
	r,_ := new(big.Int).SetString(transactions.RSignature(ephePri1),10)
	m := "Hello"
	fmt.Println(transactions.SSignature(ephePri,pri,r,m))
}