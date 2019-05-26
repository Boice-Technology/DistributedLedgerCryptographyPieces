package transactions

import ("math/big";
		"eckg";
		"crypto/sha256";
		"fmt";
		)

func SSignature(ephePri *big.Int, pri *big.Int, r *big.Int, m string) string {
	epheModInv,_ := new(big.Int).SetString("0",10)
	n,_ := new(big.Int).SetString(eckg.N,10)
	fmt.Println("n",n)
	fmt.Println("ephePri",ephePri)
	epheModInv.ModInverse(ephePri,n)
	fmt.Println("ephePri",ephePri)
	fmt.Println("epheInv",epheModInv)
	sizedMessByte := sha256.Sum256([]byte(m))
	messHashStr := fmt.Sprintf("%x",sizedMessByte)
	fmt.Println("messHashStr",messHashStr)
	messHash,_ := new(big.Int).SetString(messHashStr,16)
	prod,_ := new(big.Int).SetString("0",10)
	prod.Mul(r,pri)
	fmt.Println("prod",prod)
	add,_ := new(big.Int).SetString("0",10)
	add.Add(messHash, prod)
	fmt.Println("add",add)
	s,_ := new(big.Int).SetString("0",10)
	s.Mul(add,epheModInv)
	fmt.Println("s",s)
	s.Mod(s,n)
	fmt.Println("s",s)
	sStr := fmt.Sprintf("%s",s)
	return sStr
}