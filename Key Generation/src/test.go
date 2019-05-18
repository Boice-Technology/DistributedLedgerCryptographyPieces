package main

import ("eckg";
		"fmt";
		"math/big"
		)

func main(){
	//var Pri = new(big.Int).SetString("3",10)
	//fmt.Println(PublicKey(Pri))
	var x,_ = new(big.Int).SetString("34499628904269660561674201530767158034393542375844615658184142552908072257357",10)
	var primef,_ = new(big.Int).SetString(eckg.PrimeF,10)
	fmt.Println(x.ModInverse(x,primef))
}

