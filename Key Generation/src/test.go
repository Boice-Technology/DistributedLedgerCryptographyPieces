package main

import ("eckg";
		"math/big";
		"fmt")

func main(){
	var num,_ = new(big.Int).SetString("43",10)
	fmt.Println(eckg.IntToBin(num))
}