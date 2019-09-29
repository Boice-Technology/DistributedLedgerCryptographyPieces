// Slope of a two points line on the curve can be given as s = ((y2 - y1) / (x2 - x1)) % p

package elliptic_curve_key_generation

import ("math/big";
		"fmt";
		)

func SlopeOfTwoPointLine(x1 string, y1 string, x2 string, y2 string) string {
	x1Int,_ := new(big.Int).SetString(x1,10)
	y1Int,_ := new(big.Int).SetString(y1,10)
	x2Int,_ := new(big.Int).SetString(x2,10)
	y2Int,_ := new(big.Int).SetString(y2,10)
	
	numerator,_ := new(big.Int).SetString("0",10)
	denominator,_ := new(big.Int).SetString("0",10)
	
	multiplier,_ := new(big.Int).SetString("0",10)
	
	pInt,_ := new(big.Int).SetString(P,16)
	
	slope,_ := new(big.Int).SetString("0",10)
	
	numerator.Sub(y2Int, y1Int)
	denominator.Sub(x2Int, x1Int)
	
	multiplier.ModInverse(denominator,pInt)
	
	numerator.Mul(numerator,multiplier)
	
	slope.Mod(numerator,pInt)
	
	slopeStr := fmt.Sprintf("%s",slope)
	
	return slopeStr
}