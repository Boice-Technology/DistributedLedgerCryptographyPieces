package walletcrypto

import ("math";
		)
		
func BinToDec(binStr string) int {
	length := len(binStr)
	var i int
	var acc int
	for i=length-1;i>=0;i-- {
		if binStr[length-i-1] == 49{
			acc += int(math.Pow(2,float64(i)))
		}
	}
	return acc
}