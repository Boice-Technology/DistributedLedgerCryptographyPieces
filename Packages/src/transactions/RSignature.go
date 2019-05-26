package transactions

import ("eckg";
		"math/big";
		"fmt";
		)

func RSignature(ephePri *big.Int) (string) {
	x_coor,_ := EphemeralPublicKey(ephePri)
	r,_ := new(big.Int).SetString(x_coor,10)
	fmt.Println(r)
	n,_ := new(big.Int).SetString(eckg.N,10)
	fmt.Println(n.Cmp(r))
	r.Mod(r,n)
	fmt.Println(r)
	rStr := fmt.Sprintf("%s",r)
	fmt.Println(rStr)
	return rStr
}
