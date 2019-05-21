package walletcrypto

import ("crypto/sha256";
		"fmt";
		)
			
func Checksum(str string) string {
	shaHashBytes := sha256.Sum256([]byte(str))
	hashString := fmt.Sprintf("%x",shaHashBytes)
	return hashString[0]
}