package walletcrypto

import ("crypto/sha512";
		"fmt";
		)
		
func MasterNode(seed string) (string,string) {
	sizedBytesHash := sha512.Sum512([]byte(seed))
	master := fmt.Sprintf("%x",sizedBytesHash)
	return master[:64], master[64:]
}