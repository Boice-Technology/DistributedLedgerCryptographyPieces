// Decode the input string from base58 to standard unicode string

package elliptic_curve_key_generation

import ("github.com/btcsuite/btcutil/base58";
		)
		
func Base58Decode(inputString string) string {
	base58Decoded := base58.Decode(inputString)
	base58DecodedString := string(base58Decoded)
	return base58DecodedString
}