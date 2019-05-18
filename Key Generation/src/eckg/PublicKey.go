package eckg

import ("math/big";
		)
		
func PrimaryKey(pri *big.Int) (string,string) {
	var partition = Partition(pri)
	length := len(partition)
	i := 1
	var x,_ = new(big.Int).SetString(GenMultiples[partition[0]][0],10)
	var y,_ = new(big.Int).SetString(GenMultiples[partition[0]][1],10)
	var slope,_ = new(big.Int).SetString("0",10)	
	var primef,_ = new(big.Int).SetString(PrimeF,10)
	var nume,_ = new(big.Int).SetString("0",10)
	var deno,_ = new(big.Int).SetString("0",10)
	var mul,_ = new(big.Int).SetString("0",10)
	var add,_ = new(big.Int).SetString("0",10)
	var sub,_ = new(big.Int).SetString("0",10)
	var modInv,_ = new(big.Int).SetString("0",10)
	for i<length {	
		var x2,_ = new(big.Int).SetString(GenMultiples[partition[i]][0],10)
		var y2,_ = new(big.Int).SetString(GenMultiples[partition[i]][1],10)
		nume.Sub(y2,y)
		deno.Sub(x2,x)
		modInv.ModInverse(deno,primef)
		slope.Mul(nume,modInv)
		slope.Mod(slope,primef)
		mul.Mul(slope,slope)
		add.Add(x,x2)
		x.Sub(mul,add)
		x.Mod(x,primef)
		sub.Sub(x2,x)
		mul.Mul(slope,sub)
		y.Sub(mul,y2)
		y.Mod(y,primef)
		i++
	}
	return x.String(),y.String()
}