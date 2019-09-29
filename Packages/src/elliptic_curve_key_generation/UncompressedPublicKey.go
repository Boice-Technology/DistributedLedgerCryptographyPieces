// Generate uncompressed public key from the private key.

package elliptic_curve_key_generation

import ("math/big";
		"fmt";
		)
		
func UncompressedPublicKey(privateKey string) string {
	publicKeyX, publicKeyY := PublicKeyGenerator(privateKey)
	
	publicKeyXInt,_ := new(big.Int).SetString(publicKeyX,10)
	publicKeyYInt,_ := new(big.Int).SetString(publicKeyY,10)
	
	publicKeyXHex := fmt.Sprintf("%X",publicKeyXInt)
	publicKeyYHex := fmt.Sprintf("%X",publicKeyYInt)
	
	uncompressedPublicKey := "04" + publicKeyXHex + publicKeyYHex
	
	return uncompressedPublicKey
}