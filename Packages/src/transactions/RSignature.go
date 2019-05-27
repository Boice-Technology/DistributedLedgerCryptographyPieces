package transactions

import ("math/big";
		"fmt";
		"eckg";
		)

func RSignature(ephePri *big.Int) (string) {
	x_coor,_ := EphemeralPublicKey(ephePri)
	r,_ := new(big.Int).SetString(x_coor,10)
	n,_ := new(big.Int).SetString(eckg.N,10)
	r.Mod(r,n)
	x_coor = fmt.Sprintf("%s",r)
	return x_coor
}
