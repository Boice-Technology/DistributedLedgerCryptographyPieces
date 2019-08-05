/* Function to encode the data to base58 */

package eckg

import ("github.com/btcsuite/btcutil/base58";
		)

func Base58Encode(data string) string {
	str := base58.Encode([]byte(data))
	return str
}