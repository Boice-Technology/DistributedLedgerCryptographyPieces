package eckg

import ("math/big";
		"fmt";
		"strconv";
		)

func CompPublicKey(pri *big.Int) string {
	x,y := PublicKey(pri)
	xBig,_ := new(big.Int).SetString(x,10)
	xHex := fmt.Sprintf("%x",xBig)
	length := len(y)
	lastDig := string( y[length-1] )
	lastDigInt,_ := strconv.Atoi(lastDig)
	compPubKey := ""
	if lastDigInt % 2 == 0 {
		compPubKey += "02" + xHex
	}else {
		compPubKey += "03" + xHex
	}
	return compPubKey
}