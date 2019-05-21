package main

import ("fmt";
		"crypto/sha256")

func main(){
	byteHash := sha256.Sum256([]byte("dshjbdsjd"))
	fmt.Println(byteHash)
	strHex := fmt.Sprintf("%x",byteHash)
	fmt.Println(len(strHex))
}