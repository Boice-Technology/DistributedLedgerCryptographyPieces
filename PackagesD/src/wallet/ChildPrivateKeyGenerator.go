// Generate child private key from parent private key and parent chain code

package wallet

import ("elliptic_curve_key_generation";
		"math/big";
		"fmt";
		)
		
func ChildPrivateKeyGenerator(parentPrivateKey string, parentChainCode string,index int32) (string,string){
	parentPrivateKeyInt,_ := new(big.Int).SetString(parentPrivateKey,16)
	parentPrivateKeyString := fmt.Sprintf("%d",parentPrivateKeyInt)
	compressedParentPublicKey := elliptic_curve_key_generation.CompressedPublicKey(parentPrivateKeyString)
	indexString := fmt.Sprintf("%d",index)
	inputString := compressedParentPublicKey + parentChainCode + indexString
	hashString := SHA512(inputString)
	leftHashString := hashString[0:64]
	childChainCode := hashString[64:]
	leftHashStringInt,_ := new(big.Int).SetString(leftHashString,16)
	childPrivateKeyInt,_ := new(big.Int).SetString("0",10)
	pInt,_ := new(big.Int).SetString(elliptic_curve_key_generation.P,16)
	childPrivateKeyInt.Add(parentPrivateKeyInt,leftHashStringInt)
	childPrivateKeyInt.Mod(childPrivateKeyInt,pInt)
	childPrivateKey := fmt.Sprintf("%X",childPrivateKeyInt)
	return childPrivateKey, childChainCode
}