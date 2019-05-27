package main

import ("fmt";
		"transactions";
		"math/big";
		"eckg";
		)
		
func main(){
	ephePri1,_ := new(big.Int).SetString("1",10)
	ephePri2,_ := new(big.Int).SetString("1",10)
	ephePri3,_ := new(big.Int).SetString("1",10)
	fmt.Println("ephepri",ephePri1)
	pri1,_ := new(big.Int).SetString("2",10)
	pri2,_ := new(big.Int).SetString("2",10)
	pri3,_ := new(big.Int).SetString("2",10)
	fmt.Println("pri",pri1)
	xepheStr := transactions.RSignature(ephePri2)
	r1,_ := new(big.Int).SetString(xepheStr,10)
	r2,_ := new(big.Int).SetString(xepheStr,10)
	r3,_ := new(big.Int).SetString(xepheStr,10)
	fmt.Println("r",r1)
	m := "Hello"
	sStr := transactions.SSignature(ephePri3,pri2,r2,m)
	s1,_ := new(big.Int).SetString(sStr,10)
	s2,_ := new(big.Int).SetString(sStr,10)
	fmt.Println("s",s1)
	xpub_Str, ypub_Str := eckg.PublicKey(pri3)
	x_pub1,_ := new(big.Int).SetString(xpub_Str,10)
	y_pub1,_ := new(big.Int).SetString(ypub_Str,10)
	x_pub2,_ := new(big.Int).SetString(xpub_Str,10)
	y_pub2,_ := new(big.Int).SetString(ypub_Str,10)
	fmt.Println("xpub", "ypub", x_pub1, y_pub1)
	xVer, _ := transactions.VerificationPoint(r3,s2,m,x_pub2,y_pub2)
	fmt.Println("xVer", xVer)
}