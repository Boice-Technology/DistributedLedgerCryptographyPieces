package transactions

import ("math/big";
		"eckg";
		"fmt"
		)

func U2(r *big.Int,s *big.Int) string {
	modInv,_ := new(big.Int).SetString("0",10)
	n,_ := new(big.Int).SetString(eckg.N,10)
	modInv.ModInverse(s,n)
	prod,_ := new(big.Int).SetString("0",10)
	prod.Mul(r,modInv)
	u2,_ := new(big.Int).SetString("0",10)
	u2.Mod(prod,n)
	u2Str := fmt.Sprintf("%s",u1)
	return u2str
}
