// Compute power-modulus of extremely big numbers

package wallet

import ("math/big";
		"fmt";
		)
		
func PowerModulus(base string, exponent string, modulus string) string {
	result,_ := new(big.Int).SetString("1",10)
	baseInt,_ := new(big.Int).SetString(base,10)
	exponentInt,_ := new(big.Int).SetString(exponent,10)
	modulusInt,_ := new(big.Int).SetString(modulus,10)
	zero,_ := new(big.Int).SetString("0",10)
	two,_ := new(big.Int).SetString("2",10)
	one,_ := new(big.Int).SetString("1",10)
	mod,_ := new(big.Int).SetString("0",10)
	for zero.Cmp(exponentInt) != 0{
		mod.Mod(exponentInt,two)
		if mod.Cmp(one) == 0{
			result.Mul(result,baseInt)
			result.Mod(result,modulusInt)
		}
		baseInt.Mul(baseInt,baseInt)
		baseInt.Mod(baseInt,modulusInt)
		exponentInt.Div(exponentInt,two)
	}
	resultString := fmt.Sprintf("%d",result)
	return resultString
}