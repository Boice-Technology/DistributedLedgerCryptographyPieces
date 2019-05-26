package main

import ("transactions";
		"math/big";
		"fmt";
		)
		
func main(){	
	ephePri,_ := new(big.Int).SetString("1",10)
	fmt.Println(transactions.RSignature(ephePri))
}