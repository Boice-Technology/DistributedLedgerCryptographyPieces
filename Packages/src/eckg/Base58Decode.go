/* Function to decode the base58 string to normal string */

package eckg

import ("github.com/btcsuite/btcutil/base58";
		)
		
func Base58Decode(data string) string {
	str := string(base58.Decode(data))
	return str
}