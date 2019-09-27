// Genearate public key components from the compressed public key

package wallet

import ("math/big"
		"elliptic_curve_key_generation";
		"fmt";
		)

func PublicKeyComponents(compressedKey string) (string,string){
	publicKeyX := compressedKey[2:]
	publicKeyXInt,_ := new(big.Int).SetString(publicKeyX,16)
	three,_ := new(big.Int).SetString("3",10)
	seven,_ := new(big.Int).SetString("7",10)
	pInt,_ := new(big.Int).SetString(elliptic_curve_key_generation.P,16)
	LHS,_ := new(big.Int).SetString("0",10)
	pow,_ := new(big.Int).SetString("0",10)
	pow.Exp(publicKeyXInt,three,nil)
	LHS.Add(pow,seven)
	exponent,_ := new(big.Int).SetString("0",10)
	one,_ := new(big.Int).SetString("1",10)
	four,_ := new(big.Int).SetString("4",10)
	sum,_ := new(big.Int).SetString("0",10)
	sum.Add(pInt,one)
	exponent.Div(sum,four)
	publicKeyYInt,_ := new(big.Int).SetString("0",10)
	publicKeyYInt.Exp(LHS,exponent,pInt)
	publicKeyXString := fmt.Sprintf("%d",publicKeyXInt)
	publicKeyYString := fmt.Sprintf("%d",publicKeyYInt)
	return publicKeyXString, publicKeyYString
}