/* This function computes y-coordinate of the point which ia addition of two pints as yr = s*(x1 - xr) - y1 */

package eckg

import ("fmt";
		"math/big";
		)
		
func EllipticAddY(slope string,xr string,x1 string,y1 string) string {
	s_int,_ := new(big.Int).SetString(slope,16)
	x1_int,_ := new(big.Int).SetString(x1,16)
	y1_int,_ := new(big.Int).SetString(y1,16)
	xr_int,_ := new(big.Int).SetString(xr,16)
	p_int,_ := new(big.Int).SetString(P,16)
	x_sub,_ := new(big.Int).SetString("0",10)
	x_sub.Sub(x1_int,xr_int)
	x_mul,_ := new(big.Int).SetString("0",10)
	x_mul.Mul(s_int,x_sub)
	resy,_ := new(big.Int).SetString("0",10)
	resy.Sub(x_mul,y1_int)
	resy.Mod(resy,p_int)
	resy_str := fmt.Sprintf("%X",resy)
	return resy_str
}