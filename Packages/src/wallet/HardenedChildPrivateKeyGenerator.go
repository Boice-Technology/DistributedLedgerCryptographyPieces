// Generate hardended child private key from parent private key

package wallet

import ("fmt";
		"elliptic_curve_key_generation";
		"math/big";
		)

func HardenedChildPrivateKeyGenerator(parentPrivateKey string, parentChainCode string, index int32) (string, string){
	indexString := fmt.Sprintf("%d",index)
	inputString := parentPrivateKey + parentChainCode + indexString
	hashString :=  SHA512(inputString)
	leftHashString := hashString[:64]
	childChainCode := hashString[64:]
	parentPrivateKeyInt,_ := new(big.Int).SetString(parentPrivateKey,16)
	leftHashInt,_ := new(big.Int).SetString(leftHashString,16)
	childPrivateKeyInt,_ := new(big.Int).SetString("0",10)
	pInt,_ := new(big.Int).SetString(elliptic_curve_key_generation.P,16)
	childPrivateKeyInt.Add(parentPrivateKeyInt, leftHashInt)
	childPrivateKeyInt.Mod(childPrivateKeyInt,pInt)
	childPrivateKey := fmt.Sprintf("%X",childPrivateKeyInt)
	return childPrivateKey, childChainCode
}