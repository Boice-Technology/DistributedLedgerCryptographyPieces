/* Function to compute SHA256 hash of the string */

package eckg

import ("crypto/sha256";
		"fmt";
		)

func Sha256(data string) string {
	hash := sha256.Sum256([]byte(data))
	hashstr := fmt.Sprintf("%X",hash)
	return hashstr
}