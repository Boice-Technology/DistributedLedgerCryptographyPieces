package main

import ("eckg";
		"math/big";
		"fmt"
		)
		
func main(){
	var pri1,_ = new(big.Int).SetString("3",10)
	var pri2,_ = new(big.Int).SetString("3",10)
	var uncomppub string = eckg.UncompPublicKey(pri1)
	var comppub string = eckg.CompPublicKey(pri2)
	fmt.Println(uncomppub)
	fmt.Println(comppub)
	fmt.Println(eckg.BoiceAddressGenerator(uncomppub))
	fmt.Println(eckg.BoiceAddressGenerator(comppub))
}