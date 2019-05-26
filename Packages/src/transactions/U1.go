package transactions

import ("math/big";
		"eckg";
		"fmt"
		)

func U1(z *big.Int,s *big.Int) string {
	modInv,_ := new(big.Int).SetString("0",10)
	n,_ := new(big.Int).SetString(eckg.N,10)
	modInv.ModInverse(s,n)
	prod,_ := new(big.Int).SetString("0",10)
	prod.Mul(z,modInv)
	u1,_ := new(big.Int).SetString("0",10)
	u1.Mod(prod,n)
	u1Str := fmt.Sprintf("%s",u1)
	return u1str
}
