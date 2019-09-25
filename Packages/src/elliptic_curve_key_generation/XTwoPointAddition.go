// X - Co-ordinate of the resultant point, after adding two points on the curve and is given by xr = (slope^2 - (x1+x2)) % p

package elliptic_curve_key_generation

import ("fmt";
		"math/big";
		)
		
func XTwoPointAddition(slope string, x1 string, x2 string) string {
	slopeInt,_ := new(big.Int).SetString(slope,10)
	
	x1Int,_ := new(big.Int).SetString(x1,10)
	x2Int,_ := new(big.Int).SetString(x2,10)
	
	slopeSq,_ := new(big.Int).SetString("0",10)
	
	xAddition,_ := new(big.Int).SetString("0",10)
	
	xResultant,_ := new(big.Int).SetString("0",10)
	
	pInt,_ := new(big.Int).SetString(P,16)
	
	slopeSq.Mul(slopeInt,slopeInt)
	xAddition.Add(x1Int,x2Int)
	
	xResultant.Sub(slopeSq,xAddition)
	xResultant.Mod(xResultant,pInt)
	
	xResultantStr := fmt.Sprintf("%s",xResultant)
	
	return xResultantStr
}