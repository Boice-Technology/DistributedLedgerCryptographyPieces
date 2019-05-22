package walletcrypto

import ("eckg";
		"math/big";
		"fmt";
		"crypto/sha512"
		)
		
func PrivateChildKey(priKey string,chainCode string, index int) (string, string) {
	priInt,_ := new(big.Int).SetString(priKey,16)
	publicKey := eckg.CompPublicKey(priInt)
	hashedStr := publicKey + chainCode + fmt.Sprintf("%s",index)
	hash512Bytes := sha512.Sum512([]byte(hashedStr))
	hash512Str := fmt.Sprintf("%x",hash512Bytes)
	priInt,_ = new(big.Int).SetString(priKey,16)
	childInterInt,_ := new(big.Int).SetString(hash512Str[:64],16)
	childChainCode := hash512Str[64:]
	childInterInt.Add(childInterInt,priInt)
	childPriStr := fmt.Sprintf("%x",childInterInt)
	return childPriStr[len(childPriStr)-64:], childChainCode
}
