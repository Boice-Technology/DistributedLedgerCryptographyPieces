// Y - Co-ordinate of the resultant point, after addition of a point to itself which is given by yr = (slope * (x - xr) - y) % p

package elliptic_curve_key_generation

import ("fmt";
		"math/big";
		)
		
func YOnePointAddition(x string, y string, slope string, xr string) string {
	xInt,_ := new(big.Int).SetString(x,10)
	yInt,_ := new(big.Int).SetString(y,10)
	slopeInt,_ := new(big.Int).SetString(slope,10)
	xrInt,_ := new(big.Int).SetString(xr,10)
	
	pInt,_ := new(big.Int).SetString(P,16)
	
	multiplier,_ := new(big.Int).SetString("0",10)
	substractor,_ := new(big.Int).SetString("0",10)
	yr,_ := new(big.Int).SetString("0",10)
	
	substractor.Sub(xInt,xrInt)
	multiplier.Mul(slopeInt,substractor)
	
	yr.Sub(multiplier,yInt)
	yr.Mod(yr,pInt)
	
	yrStr := fmt.Sprintf("%s",yr)
	
	return yrStr
}