/* To compute the x-coordiante of the double point as xr = s^2 - 2*x */

package eckg

import ("math/big";
		"fmt";
		)
		
func DoubleOperationX(s string,x string) string {
	s_int,_ := new(big.Int).SetString(s,16)
	s_square,_ := new(big.Int).SetString("0",10)
	s_square.Mul(s_int,s_int)
	x_int,_ := new(big.Int).SetString(x,16)
	x_mul,_ := new(big.Int).SetString("0",10)
	res,_ := new(big.Int).SetString("0",10)
	two,_ := new(big.Int).SetString("2",10)
	x_mul.Mul(two,x_int)
	res.Sub(s_square,x_mul)
	p_int,_ := new(big.Int).SetString(P,16)
	res.Mod(res,p_int)
	x_res := fmt.Sprintf("%X",res)
	return x_res
}