package transactions

import ("eckg";
		"math/big";
		"fmt";
		)

func RSignature(ephePri *big.Int) (string) {
	x_coor,_ := EphemeralPublicKey(ephePri)
	r,_ := new(big.Int).SetString(x_coor,10)
	primef,_ := new(big.Int).SetString(eckg.PrimeF,10)
	r.Mod(r,primef)
	rStr := fmt.Sprintf("%s",r)
	return rStr
}
