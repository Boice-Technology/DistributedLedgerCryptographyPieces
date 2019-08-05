/* This fucntion computes x-coordinate of the point which is the addition of two points as xr = s^2 - (x1 + x2) */

package eckg

import ("fmt";
		"math/big";
		)
		
func EllipticAddX(slope string,x1 string,x2 string) string {
	s_int,_ := new(big.Int).SetString(slope,16)
	x1_int,_ := new(big.Int).SetString(x1,16)
	x2_int,_ := new(big.Int).SetString(x2,16)
	p_int,_ := new(big.Int).SetString(P,16)
	s_square,_ := new(big.Int).SetString("0",10)
	s_square.Mul(s_int,s_int)
	x_add,_ := new(big.Int).SetString("0",10)
	x_add.Add(x1_int,x2_int)
	resx,_ := new(big.Int).SetString("0",10)
	resx.Sub(s_square,x_add)
	resx.Mod(resx,p_int)
	resx_str := fmt.Sprintf("%X",resx)
	return resx_str
}