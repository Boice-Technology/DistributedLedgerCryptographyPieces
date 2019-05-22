package walletcrypto

import ("fmt";
		"math/big";
		"crypto/sha512";
		)
		
func HardPrivateChildKey(priKey string, chainCode string, index int) (string,string) {
	priKeyInt,_ := new(big.Int).SetString(priKey,10)
	priKeyHex := fmt.Sprintf("%x",priKeyInt)
	priKeyHex = HexPrefixer(priKeyHex)
	hashedStr := priKeyHex + chainCode + fmt.Sprintf("%s",index)
	hash512Byte := sha512.Sum512([]byte(hashedStr))
	hash512Str := fmt.Sprintf("%x",hash512Byte)
	childPriInter,_ := new(big.Int).SetString(hash512Str[:64],16)
	childPriInter.Add(priKeyInt,childPriInter)
	priChildKey := fmt.Sprintf("%x",childPriInter)
	priChildKey = priChildKey[len(priChildKey)-64:]
	childChainCode := hash512Str[64:]
	return priChildKey, childChainCode
}
		