// Generate RIPEMD160 hash of input string

package elliptic_curve_key_generation

import ("fmt";
		"golang.org/x/crypto/ripemd160";
		)
		
func RIPEMD160(inputString string) string {
	stream := ripemd160.New()
	stream.Write([]byte(inputString))
	hash := fmt.Sprintf("%X",stream.Sum(nil))
	
	return hash
}