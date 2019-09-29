// Generate HMAC SHA512 hash of a input string

package wallet

import ("fmt";
		"crypto/sha512";
		)

func SHA512(inputString string) string {
	sha512HashBytes := sha512.Sum512([]byte(inputString))
	sha512Hash := fmt.Sprintf("%X",sha512HashBytes)
	return sha512Hash
}