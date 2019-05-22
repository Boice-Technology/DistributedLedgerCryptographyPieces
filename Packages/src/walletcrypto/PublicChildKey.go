package walletcrypto

import ("fmt";
		"math/big";
		"crypto/sha512";
		)
		
func PublicChildKey(pubKey string, chainCode string, index int) (string,string) {
	pubKeyPre, pubKeySuf := pubKey[:2], pubKey[2:]
	hashedStr := pubKey + chainCode + fmt.Sprintf("%s",index)
	hash512Byte := sha512.Sum512([]byte(hashedStr))
	hash512Str := fmt.Sprintf("%x",hash512Byte)
	pubKeyInt,_ := new(big.Int).SetString(pubKeySuf,16)
	childPubInter,_ := new(big.Int).SetString(hash512Str[:64],16)
	childPubInter.Add(childPubInter,pubKeyInt)
	childPubKey := fmt.Sprintf("%x",childPubInter)
	childPubKey = pubKeyPre + childPubKey[len(childPubKey)-64:]
	childChainCode := hash512Str[64:]
	return childPubKey, childChainCode
}