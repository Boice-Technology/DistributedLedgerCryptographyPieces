package transactions

import ("eckg";
		"math/big";
		)
		
func EphemeralPublicKey(ephePri *big.Int) (string, string) {
	x_coor, y_coor := eckg.PublicKey(ephePri)
	return x_coor, y_coor
}
		