package main

import ("fmt";
		"eckg"
		)

func main(){
	var count int = 0
	for count<257{
		fmt.Println(eckg.IsEllipticCurvePoint(eckg.Doubles[count][0],eckg.Doubles[count][1]))	
		count++
	}
}