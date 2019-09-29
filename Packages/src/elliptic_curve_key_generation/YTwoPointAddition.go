// Y - Co-ordinate of the resultant point, after addition of two points on the curve and is given by yr = (slope * (x2 - xr) - y2) % p

package elliptic_curve_key_generation

import ("fmt";
		"math/big";
		)
	
func YTwoPointAddition(slope string, x2 string, y2 string, xr string) string {
	slopeInt,_ := new(big.Int).SetString(slope,10)
	x2Int,_ := new(big.Int).SetString(x2,10)
	y2Int,_ := new(big.Int).SetString(y2,10)
	xrInt,_ := new(big.Int).SetString(xr,10)
	
	xSubstractor,_ := new(big.Int).SetString("0",10)
	multiplier,_ := new(big.Int).SetString("0",10)
	resSubstractor,_ := new(big.Int).SetString("0",10)
	
	pInt,_ := new(big.Int).SetString(P,16)
	
	yr,_ := new(big.Int).SetString("0",10)
	
	xSubstractor.Sub(x2Int,xrInt)
	multiplier.Mul(slopeInt,xSubstractor)
	resSubstractor.Sub(multiplier,y2Int)
	
	yr.Mod(resSubstractor,pInt)
	
	yrStr := fmt.Sprintf("%s",yr)
	
	return yrStr
}