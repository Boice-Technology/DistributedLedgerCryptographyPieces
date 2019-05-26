package transactions

import ("eckg";
		"math/big";
		)
		
func EphemeralPublicKey(pri *big.Int) (string, string) {
	x_coor, y_coor := eckg.PublicKey(pri)
	return x_coor, y_coor
}
		