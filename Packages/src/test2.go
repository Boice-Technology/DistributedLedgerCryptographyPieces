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
	str1 := fmt.Sprint(eckg.BoiceAddressGenerator(uncomppub))
	str2 := fmt.Sprint(eckg.BoiceAddressGenerator(comppub))
	fmt.Println(eckg.CheckBoiceAddress(str1))
	fmt.Println(eckg.CheckBoiceAddress(str2))
}