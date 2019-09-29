// X - Co-ordinate of the resultant point, after addition of point to itself which is given by xr = slope ^ 2 - 2 * x

package elliptic_curve_key_generation

import ("fmt";
		"math/big";
		)
		
func XOnePointAddition(x string, slope string) string {
	xInt,_ := new(big.Int).SetString(x,10)
	slopeInt,_ := new(big.Int).SetString(slope,10)
	
	two,_ := new(big.Int).SetString("2",10)
	multiplier1,_ := new(big.Int).SetString("0",10)
	multiplier2,_ := new(big.Int).SetString("0",10)
	xr,_ := new(big.Int).SetString("0",10)
	pInt,_ := new(big.Int).SetString(P,16)
	
	multiplier1.Mul(slopeInt,slopeInt)
	multiplier2.Mul(two,xInt)
	xr.Sub(multiplier1,multiplier2)
	xr.Mod(xr,pInt)
	
	xrStr := fmt.Sprintf("%s",xr)
	
	return xrStr
}
		
		