// Generates SHA256 hash of a input string

package elliptic_curve_key_generation

import ("crypto/sha256";
		"fmt";
		)
		
func SHA256(inputString string) string {
	hashBytes := sha256.Sum256([]byte(inputString))
	hash := fmt.Sprintf("%X",hashBytes)
	return hash
}