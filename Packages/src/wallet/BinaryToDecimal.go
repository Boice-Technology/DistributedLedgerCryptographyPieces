// Generate decimal representation of a binary number representation

package wallet

import ("elliptic_curve_key_generation";
		)

func BinaryToDecimal(binaryString string) int {
	multiplier := 1
	reverseString := elliptic_curve_key_generation.Reverse(binaryString)
	decimalNumber := 0
	for i := 0 ; i < len(binaryString) ; i++ {
		if reverseString[i:i+1] == "1"{
			decimalNumber += multiplier
		}
		multiplier *= 2
	}
	return decimalNumber
}
