package transactions

import ("math/big";
		"eckg";
		"crypto/sha256";
		"fmt";
		)

func SSignature(ephePri *big.Int, pri *big.Int, r *big.Int, m string) string {
	epheModInv,_ := new(big.Int).SetString("0",10)
	n,_ := new(big.Int).SetString(eckg.N,10)
	epheModInv.ModInverse(ephePri,n)
	sizedMessByte := sha256.Sum256([]byte(m))
	messHashStr := fmt.Sprintf("%x",sizedMessByte)
	messHash,_ := new(big.Int).SetString(messHashStr,16)
	prod,_ := new(big.Int).SetString("0",10)
	prod.Mul(r,pri)
	add,_ := new(big.Int).SetString("0",10)
	add.Add(messHash, prod)
	s,_ := new(big.Int).SetString("0",10)
	s.Mul(add,epheModInv)
	s.Mod(s,n)
	sStr := fmt.Sprintf("%s",s)
	return sStr
}