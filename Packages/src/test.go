package main

import ("fmt";
		"eckg"
		)

func main(){
	var i string
	fmt.Scanf("%s",&i)
	x,y := eckg.PublicKeyGenerator(i)
	fmt.Println(x,y)
	p1 := eckg.UncompressedPublicKey(x,y)
	p2 := eckg.CompressedPublicKey(x,y)
	fmt.Println(p1,p2)
	b1 := eckg.BoiceAddressGenerator(p1)
	b2 := eckg.BoiceAddressGenerator(p2)
	fmt.Println(b1,b2)
}