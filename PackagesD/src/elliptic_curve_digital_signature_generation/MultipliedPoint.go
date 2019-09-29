// Computes P' which is given by P' = k * P, where k is a multiplier and P is a point

package elliptic_curve_digital_signature_generation

import ("fmt";
		"math/big";
		"elliptic_curve_key_generation";
		)

func MultipliedPoint(multiplier string, PointX string, PointY string) (string, string){
	multiplierInt,_ := new(big.Int).SetString(multiplier,10)
	binaryMultiplier := fmt.Sprintf("%b",multiplierInt)
	reverseBinaryMultiplier := elliptic_curve_key_generation.Reverse(binaryMultiplier)
	var powers []int = make([]int,1)
	for i := 0 ; i < len(binaryMultiplier) ; i++{
		if reverseBinaryMultiplier[i:i+1]=="1"{
			powers = append(powers,i)
		}
	}
	prevX := PointX
	prevY := PointY
	k := 1
	var points []string = make([]string, 1)
	if powers[k]==0{
		points = append(points, prevX, prevY)
		k++
	}
	for i := 1 ; i < 257 && k<len(powers); i++ {
		newX, newY := elliptic_curve_key_generation.OnePointAddition(prevX, prevY)
		if powers[k]==i{
			points = append(points, newX, newY)
			k++
		}
		prevX = newX
		prevY = newY
	}
	finalPointX := points[1]
	finalPointY := points[2]
	for i := 2 ; i < len(powers) ; i++{
		newX := points[2*i-1]
		newY := points[(2*i)]
		finalPointX, finalPointY = elliptic_curve_key_generation.TwoPointAddition(finalPointX, finalPointY, newX, newY)
	}
	return finalPointX, finalPointY
}