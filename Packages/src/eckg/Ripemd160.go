/* Fucntion to compute RIPEMD160 hash of the string */

package eckg

import ("golang.org/x/crypto/ripemd160";
		"fmt";
		)

func Ripemd160(data string) string {
	h := ripemd160.New()
    h.Write([]byte(data))
    hashstr := fmt.Sprintf("%X", h.Sum(nil))
	return hashstr
}