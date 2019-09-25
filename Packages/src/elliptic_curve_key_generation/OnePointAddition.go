// Give result for one point addition of a point on a curve

package elliptic_curve_key_generation

func OnePointAddition(x string,y string) (string, string) {
	slope := SlopeOfTangent(x, y)
	xr := XOnePointAddition(x, slope)
	yr := YOnePointAddition(x, y, slope, xr)
	
	return xr, yr
}