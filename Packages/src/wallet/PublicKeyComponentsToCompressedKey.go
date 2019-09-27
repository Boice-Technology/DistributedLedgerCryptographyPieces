// Generate compressed public key from the public key components

package wallet

import ("math/big";
		"fmt";
		)
		
func PublicKeyComponentsToCompressedKey(publicKeyX string, publicKeyY string) string {
	var lastDigit string = publicKeyY[len(publicKeyY)-1:]
	publicKeyXInt,_ := new(big.Int).SetString(publicKeyX,10)
	hexPublicKeyX := fmt.Sprintf("%X",publicKeyXInt)
	var prefix string
	if (lastDigit[0]-48)%2==0{
		prefix = "02"
	}else{
		prefix = "03"
	}
	compressedPublicKey := prefix + hexPublicKeyX
	return compressedPublicKey
}