// Encode a input string into base58 format of representation

package elliptic_curve_key_generation

import ("github.com/btcsuite/btcutil/base58";
		)
		
func Base58Encode(inputString string) string {
	base58Encoded := base58.Encode([]byte(inputString))
	return base58Encoded
}