package eckg

import ("crypto/sha256";
		"fmt";
		)
		
func Checksum(pubHash string) string {
	sha1Bytes := sha256.Sum256([]byte(pubHash))
	sha1Str := fmt.Sprintf("%x",sha1Bytes)
	sha2Bytes := sha256.Sum256([]byte(sha1Str))
	str := fmt.Sprintf("%x",sha2Bytes)
	str = str[len(str)-4:]
	return str
}