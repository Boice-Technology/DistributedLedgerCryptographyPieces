// Generate child public key from parent public key and parent chain code

package wallet

import ("fmt";
		"elliptic_curve_key_generation";
		"math/big";
		)

func ChildPublicKeyGenerator(parentPublicKey string, parentChainCode string, index int32) (string, string){
	indexString := fmt.Sprintf("%d",index)
	inputString := parentPublicKey + parentChainCode + indexString
	hashString := SHA512(inputString)
	leftHash := hashString[0:64]
	childChainCode := hashString[64:]
	leftHashInt,_ := new(big.Int).SetString(leftHash,16)
	leftHashString := fmt.Sprintf("%d",leftHashInt)
	leftPublicKeyX, leftPublicKeyY := elliptic_curve_key_generation.PublicKeyGenerator(leftHashString)
	parentPublicKeyX, parentPublicKeyY := PublicKeyComponents(parentPublicKey)
	childPublicKeyX, childPublicKeyY := elliptic_curve_key_generation.TwoPointAddition(leftPublicKeyX,leftPublicKeyY,parentPublicKeyX,parentPublicKeyY)
	childPublicKey := PublicKeyComponentsToCompressedKey(childPublicKeyX,childPublicKeyY)
	return childPublicKey, childChainCode
}
