package eckg

func DoubleOperation(x string,y string)(string ,string){
	slope := DoubleOperationSlope(x,y)
	x2 := DoubleOperationX(slope,x)
	y2 := DoubleOperationY(slope,x2,x,y)
	return x2, y2
}