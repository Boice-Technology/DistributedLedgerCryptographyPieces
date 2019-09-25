package main

import ("fmt";
		"math/big"
		"elliptic_curve_key_generation";
		)
		
func main(){
	//fmt.Println(elliptic_curve_key_generation.DecimalToBinary("3"))
	x, y := elliptic_curve_key_generation.PublicKeyGenerator("13840170145645816737842251482747434280357113762558403558088249138233286766301")
		//num1,_ := new(big.Int).SetString("65341020041517633956166170261014086368942546761318486551877808671514674964848",10)
		//pInt,_ := new(big.Int).SetString(elliptic_curve_key_generation.P,16)
		//num1.ModInverse(num1,pInt)
		//fmt.Println(num1)
	fmt.Println(x,y)
	num1,_ := new(big.Int).SetString("108626704259373488493324494832963472198167093861790438979212875284973843395610",10)
	fmt.Printf("%X",num1)
}