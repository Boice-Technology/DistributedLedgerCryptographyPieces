package main

import ("fmt";
		"math/big";
		"eckg";
		)
		
func main(){
	num,_ := new(big.Int).SetString("63",10)
	fmt.Println(eckg.Partition(num))
}
		