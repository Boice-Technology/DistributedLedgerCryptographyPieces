package eckg

import ("math/big";
		"fmt";
		)
		
func UncompPublicKey(pri *big.Int) string{
	x,y := PublicKey(pri)
	xBig,_ := new(big.Int).SetString(x,10)
	yBig,_ := new(big.Int).SetString(y,10)
	xHex := fmt.Sprintf("%x",xBig)
	yHex := fmt.Sprintf("%x",yBig)
	UncompPubKey := "04" + xHex + yHex
	return UncompPubKey
}