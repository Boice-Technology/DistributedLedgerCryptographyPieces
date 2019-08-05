/* To check whether a point is on elliptic curve or not which satisfies x^3 + 7 = y^2  */

package eckg

import ("math/big";	
		)

func IsEllipticCurvePoint(x string, y string) bool {
	y_int,_ := new(big.Int).SetString(y,16)
	x_int,_ := new(big.Int).SetString(x,16)
	seven,_ := new(big.Int).SetString("7",10)
	x_square,_ := new(big.Int).SetString("0",10)
	y_square,_ := new(big.Int).SetString("0",10)
	x_cube,_ := new(big.Int).SetString("0",10)
	x_square.Mul(x_int,x_int)
	x_cube.Mul(x_square,x_int)
	y_square.Mul(y_int,y_int)
	x_cube.Add(x_cube,seven)
	p_int,_ := new(big.Int).SetString(P,16)
	y_square.Mod(y_square,p_int)
	x_cube.Mod(x_cube,p_int)
	if x_cube.Cmp(y_square)!=0{
		return false
	}
	return true
}