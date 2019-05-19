package eckg

import ("github.com/btcsuite/btcutil/base58";
		)

func CheckBoiceAddress(base58Enc string){
	decodedBytes := base58.Decode(string)
	fmt.Printf("%s",decodedBytes)
}