package eckg

func EllipticAdd(x1 string,y1 string,x2 string,y2 string) (string,string) {
	slope := EllipticAddSlope(x1,y1,x2,y2)
	xr := EllipticAddX(slope,x1,x2)
	yr := EllipticAddY(slope,xr,x1,y1)
	return xr,yr
}