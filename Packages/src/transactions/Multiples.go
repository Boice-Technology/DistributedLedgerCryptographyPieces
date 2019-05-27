package transactions

import ("math/big";
		"fmt";
		"eckg";
		)

func Multiples(t int16, x *big.Int, y *big.Int) [][2]string {
	multiples := make([][2]string,t+1)
	multiples[0][0] = fmt.Sprintf("%s",x)
	multiples[0][1] = fmt.Sprintf("%s",y)
	x1,_ := new(big.Int).SetString(multiples[0][0],10)
	y1,_ := new(big.Int).SetString(multiples[0][1],10)
	three,_ := new(big.Int).SetString("3",10)
	a,_ := new(big.Int).SetString(eckg.A,10)
	two,_ := new(big.Int).SetString("2",10)
	primef,_ := new(big.Int).SetString(eckg.PrimeF,10)
	var i int
	for i=1 ; i<=int(t) ; i++ {
		slope,_ := new(big.Int).SetString("0",10)
		sqr,_ := new(big.Int).SetString("0",10)
		sqr.Mul(x1,x1)
		nume,_ := new(big.Int).SetString("0",10)
		nume.Mul(three,sqr)
		nume.Add(nume,a)
		deno,_ := new(big.Int).SetString("0",10)
		deno.Mul(two,y1)
		modInv,_ := new(big.Int).SetString("0",10)
		modInv.ModInverse(deno,primef)
		slope.Mul(nume,modInv)
		slope.Mod(slope,primef)
		x2,_ := new(big.Int).SetString("0",10)
		y2,_ := new(big.Int).SetString("0",10)
		sqr.Mul(slope,slope)
		prod,_ := new(big.Int).SetString("0",10)
		prod.Mul(two,x1)
		x2.Sub(sqr,prod)
		x2.Mod(x2,primef)
		diff,_ := new(big.Int).SetString("0",10)
		diff.Sub(x1,x2)
		prod.Mul(slope,diff)
		y2.Sub(prod,y1)
		y2.Mod(y2,primef)
		multiples[i][0] = fmt.Sprintf("%s",x2)
		multiples[i][1] = fmt.Sprintf("%s",y2)
		x1 = x2
		y1 = y2
	}
	return multiples
}
