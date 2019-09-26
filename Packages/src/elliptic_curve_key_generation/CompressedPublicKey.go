// Generate compressed public key from private key as, if y co-ordinate of the public key is negative or odd, prefix is "03" else "02"

package elliptic_curve_key_generation

import ("math/big";
		"fmt";
		)
		
func CompressedPublicKey(privateKey string) string {
	publicKeyX, publicKeyY := PublicKeyGenerator(privateKey)
	
	publicKeyXInt,_ := new(big.Int).SetString(publicKeyX,10)
	
	publicKeyXHex := fmt.Sprintf("%X",publicKeyXInt)
	var lastDigit byte = publicKeyY[len(publicKeyY)-1]
	
	var compressedPublicKey string
	
	if (lastDigit-48)%2==0{
		compressedPublicKey = "02" + publicKeyXHex
	}else{
		compressedPublicKey = "03" + publicKeyXHex
	}
	
	return compressedPublicKey
}