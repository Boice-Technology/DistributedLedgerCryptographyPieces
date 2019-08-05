/* To compute y-coordiante of the double point as yr = s*(x - xr) - y */

package eckg

import ("math/big";
		"fmt";
		)

func DoubleOperationY(s string, x2 string, x string, y string) string {
	s_int,_ := new(big.Int).SetString(s,16)
	x2_int,_ := new(big.Int).SetString(x2,16)
	x_int,_ := new(big.Int).SetString(x,16)
	y_int,_ := new(big.Int).SetString(y,16)
	x_sub,_ := new(big.Int).SetString("0",10)
	x_sub.Sub(x_int,x2_int)
	x_mul,_ := new(big.Int).SetString("0",10)
	x_mul.Mul(s_int,x_sub)
	res,_ := new(big.Int).SetString("0",10)
	res.Sub(x_mul,y_int)
	p_int,_ := new(big.Int).SetString(P,16)
	res.Mod(res,p_int)
	y_res := fmt.Sprintf("%X",res)
	return y_res
}