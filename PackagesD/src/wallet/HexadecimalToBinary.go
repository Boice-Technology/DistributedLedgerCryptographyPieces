// Generate binary representation of a hexadecimal representation of a string

package wallet

func HexadecimalToBinary(hexString string) string {
	binString := ""
	for i := 0 ; i < len(hexString) ;i++ {
		binString += BinStringFromHexCharacter(hexString[i:i+1])
	}
	return binString
}