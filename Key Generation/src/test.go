package main

import ("eckg";
		"fmt";
		"math/big"
		)

func main(){
	var Pri,_ = new(big.Int).SetString("3",10)
	fmt.Println(eckg.PublicKey(Pri))
}

