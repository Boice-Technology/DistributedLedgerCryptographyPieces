/* To remove the slope of the tangent passing through a point as s = (3*x^2 + a)/(2*y) */

package eckg

import ("math/big";
		"fmt";
		)

func DoubleOperationSlope(x string,y string) string {
	x_int,_ := new(big.Int).SetString(x,16)
	y_int,_ := new(big.Int).SetString(y,16)
	three,_ := new(big.Int).SetString("3",10)
	two,_ := new(big.Int).SetString("2",10)
	p_int,_ := new(big.Int).SetString(P,16)
	a_int,_ := new(big.Int).SetString(A,16)
	x_square,_ :=  new(big.Int).SetString("0",10)
	x_square.Mul(x_int,x_int)
	x_mul,_ := new(big.Int).SetString("0",10)
	x_mul.Mul(three,x_square)
	nume,_ := new(big.Int).SetString("0",10)
	nume.Add(x_mul,a_int)
	nume.Mod(nume,p_int)
	y_mul,_ := new(big.Int).SetString("0",10)
	y_mul.Mul(two,y_int)
	deno,_ := new(big.Int).SetString("0",10)
	deno.ModInverse(y_mul,p_int)
	res,_ := new(big.Int).SetString("0",10)
	res.Mul(deno,nume)
	res.Mod(res,p_int)
	slope := fmt.Sprintf("%X",res)
	return slope
}