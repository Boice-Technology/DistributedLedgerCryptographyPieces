/* This function computes slope of the line passing through two points slope = (y2 - y1)/(x2 - x1) */

package eckg

import ("math/big";
		"fmt";
		)
		
func EllipticAddSlope(x1 string,y1 string,x2 string,y2 string)(string){
	y1_int,_ := new(big.Int).SetString(y1,16)
	y2_int,_ := new(big.Int).SetString(y2,16)
	x1_int,_ := new(big.Int).SetString(x1,16)
	x2_int,_ := new(big.Int).SetString(x2,16)
	p_int,_ := new(big.Int).SetString(P,16)
	y_sub,_ := new(big.Int).SetString("0",10)
	y_sub.Sub(y1_int,y2_int)
	x_sub,_ := new(big.Int).SetString("0",10)
	x_sub.Sub(x1_int,x2_int)
	x_inv,_ := new(big.Int).SetString("0",10)
	x_inv.ModInverse(x_sub,p_int)
	slope,_ := new(big.Int).SetString("0",10)
	slope.Mul(y_sub,x_inv)
	slope.Mod(slope,p_int)
	res_s := fmt.Sprintf("%X",slope)
	return res_s
}