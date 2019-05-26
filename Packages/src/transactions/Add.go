package transactions

import ("math/big";
		"fmt";
		"eckg";
		)

func Add(x1Str string, y1Str string, x2Str string, y2Str string) (string, string) {
	x1,_ := new(big.Int).SetString(x1Str,10)
	y1,_ := new(big.Int).SetString(y1Str,10)
	x2,_ := new(big.Int).SetString(x2Str,10)
	y2,_ := new(big.Int).SetString(y2Str,10)
	primef,_ := new(big.Int).SetString(eckg.PrimeF,10)
	diff1,_ := new(big.Int).SetString("0",10)
	diff2,_ := new(big.Int).SetString("0",10)
	diff1.Sub(y2,y1)
	diff2.Sub(x2,x1)
	modInv,_ := new(big.Int).SetString("0",10)
	modInv.ModInverse(diff2,primef)
	slope,_ := new(big.Int).SetString("0",10)
	slope.Mul(diff1,modInv)
	slope.Mod(slope,primef)
	sqr,_ := new(big.Int).SetString("0",10)
	sqr.Mul(slope,slope)
	add,_ := new(big.Int).SetString("0",10)
	add.Add(x1,x2)
	x,_ := new(big.Int).SetString("0",10)
	x.Sub(sqr,add)
	x.Mod(x,primef)
	diff3,_ := new(big.Int).SetString("0",10)
	diff3.Sub(x2,x1)
	prod,_ := new(big.Int).SetString("0",10)
	prod.Mul(slope,diff3)
	y,_ := new(big.Int).SetString("0",10)
	y.Sub(prod,y2)
	y.Mod(y,primef)
	xStr := fmt.Sprintf("%s",x)
	yStr := fmt.Sprintf("%s",y)
	return xStr, yStr
}