package transactions

import ("math/big";
		"crypto/sha256";
		"fmt";
		"eckg";
		)

func VerificationPoint(r *big.Int, s * big.Int, m string, x_pub *big.Int, y_pub *big.Int) (string, string) {
	sizedMessByte := sha256.Sum256([]byte(m))
	messStr := fmt.Sprintf("%x",sizedMessByte)
	z,_ := new(big.Int).SetString(messStr,16)
	u1,_ := new(big.Int).SetString(U1(z,s),10)
	u2,_ := new(big.Int).SetString(U2(r,s),10)
	genx,_ := new(big.Int).SetString(eckg.GenX,10)
	geny,_ := new(big.Int).SetString(eckg.GenY,10)
	x1Str, y1Str := ProductPoint(u1,genx,geny)
	x2Str, y2Str := ProductPoint(u2,x_pub,y_pub)
	x, y := Add(x1Str,y1Str,x2Str,y2Str)
	r1,_ := new(big.Int).SetString(x,10)
	n,_ := new(big.Int).SetString(eckg.N,10)
	r1.Mod(r1,n)
	x = fmt.Sprintf("%s",r1)
	return x,y
}