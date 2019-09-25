// Slope for the tangent at a point on the curve which is given by slope = ((3 * x * x + a)/ (2 * y)) % p

package elliptic_curve_key_generation

import ("fmt";
		"math/big";
		)
	
func SlopeOfTangent(x string, y string) string {
	xInt,_ := new(big.Int).SetString(x,10)
	yInt,_ := new(big.Int).SetString(y,10)
	
	pInt,_ := new(big.Int).SetString(P,16)
	
	aInt,_ := new(big.Int).SetString(A,16)
	
	numerator,_ := new(big.Int).SetString("0",10)
	denominator,_ := new(big.Int).SetString("0",10)
	multiplier,_ := new(big.Int).SetString("0",10)
	xSqr,_ := new(big.Int).SetString("0",10)
	invMultiplier,_ := new(big.Int).SetString("0",10)
	slope,_ := new(big.Int).SetString("0",10)
	
	three,_ := new(big.Int).SetString("3",10)
	two,_ := new(big.Int).SetString("2",10)
	
	xSqr.Mul(xInt,xInt)
	multiplier.Mul(three,xSqr)
	numerator.Add(multiplier,aInt)
	denominator.Mul(two,yInt)
	invMultiplier.ModInverse(denominator,pInt)
	
	slope.Mul(numerator,invMultiplier)
	slope.Mod(slope,pInt)
	
	slopeStr := fmt.Sprintf("%s",slope)
	
	return slopeStr
}
		


