// Give result of two point addition of the points on the curve

package elliptic_curve_key_generation

func TwoPointAddition(x1 string, y1 string, x2 string, y2 string) (string, string){
	slope := SlopeOfTwoPointLine(x1, y1, x2, y2)
	xr := XTwoPointAddition(slope, x1, x2)
	yr := YTwoPointAddition(slope, x2, y2, xr)
	
	return xr, yr
}