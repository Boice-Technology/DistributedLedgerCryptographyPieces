package eckg

import ("math/big";
		//"fmt"
		)

func SlopeOfTangent(x *big.Int, y *big.Int) *big.Int {
	/*
		slope = (3*x*x + a)/(2*y) (mod primef)
	*/
	var two,_ = new(big.Int).SetString("2",10)
	var three,_ = new(big.Int).SetString("3",10)
	var slope,_ = new(big.Int).SetString("0",10)
	var nume,_ = new(big.Int).SetString("0",10)
	var deno,_ = new(big.Int).SetString("0",10)
	var modInv,_ = new(big.Int).SetString("0",10)
	var a,_ = new(big.Int).SetString(A,10)
	var primef,_ = new(big.Int).SetString(PrimeF,10)
	nume.Mul(x,x) 
	nume.Mul(three,nume)
	nume.Add(nume,a)
	deno.Mul(two,y)
	modInv.ModInverse(deno,primef) 
	nume.Mul(nume,modInv)
	slope.Mod(nume,primef)
	return slope
}

func XCoordinate(slope *big.Int, x *big.Int) *big.Int {
	/*
		X = slope*slope - 2*x (mod primef)
	*/
	var X,_ = new(big.Int).SetString("0",10)
	var sub,_ = new(big.Int).SetString("0",10)
	var two,_ = new(big.Int).SetString("2",10)
	var primef,_ = new(big.Int).SetString(PrimeF,10)
	X.Mul(slope,slope)
	sub.Mul(two,x)
	X.Sub(X,sub)
	X.Mod(X,primef)
	return X
}

func YCoordinate(slope *big.Int,x *big.Int,y *big.Int,x2 *big.Int) *big.Int {
	/*
		Y = s*(x - x2) - y (mod primef)
	*/
	var Y,_ = new(big.Int).SetString("0",10)
	var primef,_ = new(big.Int).SetString(PrimeF,10)
	var sub,_ = new(big.Int).SetString("0",10)
	sub.Sub(x,x2)
	sub.Mul(slope,sub)
	Y.Sub(sub,y)
	Y.Mod(Y,primef)
	return Y
}

func TwoPowerMulOfGen(i int16) (string,string) {
	var prevx,_ = new(big.Int).SetString(GenX,10)
	var prevy,_ = new(big.Int).SetString(GenY,10)
	var x,_ = new(big.Int).SetString("0",10)
	var y,_ = new(big.Int).SetString("0",10)
	var j int16
	for j<i {
		slope := SlopeOfTangent(prevx,prevy)
		x = XCoordinate(slope,prevx)
		y = YCoordinate(slope,prevx,prevy,x)
		prevx = x
		prevy = y
		j++
	}
	X := prevx.String()
	Y := prevy.String()
	return X,Y
}