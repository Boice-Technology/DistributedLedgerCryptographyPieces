// Utility function to find the binary representation of a 256 bit number

package elliptic_curve_key_generation

import "math/big"

func DecimalToBinary(decimalNumber string) string{
	var binaryNumber string = ""
	decimalNumberInt,_ := new(big.Int).SetString(decimalNumber,10)
	zero,_ := new(big.Int).SetString("0",10)
	two,_ := new(big.Int).SetString("2",10)
	
	for zero.Cmp(decimalNumberInt) != 0{
		remainder,_ := new(big.Int).SetString("0",10)
		remainder.Mod(decimalNumberInt, two)
		if remainder.Cmp(zero)==0{
			binaryNumber += "0"
		}else{
			binaryNumber += "1"
		}
		decimalNumberInt.Div(decimalNumberInt, two)
	}
	binaryNumber = Reverse(binaryNumber)
	return binaryNumber
}