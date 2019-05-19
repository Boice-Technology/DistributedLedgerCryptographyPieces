package main

import ("crypto/sha256";
		"golang.org/x/crypto/ripemd160";
		"fmt";
		)

func main(){
	sum := sha256.Sum256([]byte("Hello World"))
	fmt.Printf("%x\n",sum)
	str := fmt.Sprintf("%x",sum)
	fmt.Println(str)
	digest := ripemd160.New()
	digest.Write([]byte(str))
	sum2 := digest.Sum(nil)
	fmt.Printf("%x\n",sum2)
}