// To generate a binary string corresponding to a hexadecimal character

package wallet

func BinStringFromHexCharacter(hexCharacter string) string {
	var binString string
	switch hexCharacter{
		case "0":
			binString = "0000"
		case "1":
			binString = "0001"
		case "2":
			binString = "0010"
		case "3":
			binString = "0011"
		case "4":
			binString = "0100"
		case "5":
			binString = "0101"
		case "6":
			binString = "0110"
		case "7":
			binString = "0111"
		case "8":
			binString = "1000"
		case "9":
			binString = "1001"
		case "A":
			binString = "1010"
		case "B":
			binString = "1011"
		case "C":
			binString = "1100"
		case "D":
			binString = "1101"
		case "E":
			binString = "1110"
		case "F":
			binString = "1111"
	}
	return binString
}