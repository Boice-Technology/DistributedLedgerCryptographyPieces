package eckg

import ("math/big";
		"fmt")

func IntToBin(i *big.Int) string{
	var revbin string = ""
	var zero,_ = new(big.Int).SetString("0",10)
	var two,_ = new(big.Int).SetString("2",10)
	var mod,_ = new(big.Int).SetString("0",10)
	for zero.Cmp(i)!=0 {
		mod.Mod(i,two)
		i.Div(i,two)
		revbin += mod.String()
		fmt.Println(mod)
	}
	bin := Reverse(revbin)
	return bin
}