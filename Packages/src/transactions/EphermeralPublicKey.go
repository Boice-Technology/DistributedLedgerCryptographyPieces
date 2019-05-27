package transactions

import ("eckg";
		"math/big";
		)
		
func EphemeralPublicKey(ephePri *big.Int) (string, string) {
	genx,_ := new(big.Int).SetString(eckg.GenX,10)
	geny,_ := new(big.Int).SetString(eckg.GenY,10)
	x_coor, y_coor := ProductPoint(ephePri,genx,geny)
	return x_coor, y_coor
}