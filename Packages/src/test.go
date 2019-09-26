package main

import ("fmt";
		//"math/big"
		//"elliptic_curve_key_generation";
		"wallet"
		)
		
func main(){
	/*//fmt.Println(elliptic_curve_key_generation.DecimalToBinary("3"))
	//x, y := elliptic_curve_key_generation.PublicKeyGenerator("13840170145645816737842251482747434280357113762558403558088249138233286766301")
		//num1,_ := new(big.Int).SetString("65341020041517633956166170261014086368942546761318486551877808671514674964848",10)
		//pInt,_ := new(big.Int).SetString(elliptic_curve_key_generation.P,16)
		//num1.ModInverse(num1,pInt)
		//fmt.Println(num1)
	//fmt.Println(x,y)
	//num1,_ := new(big.Int).SetString("108626704259373488493324494832963472198167093861790438979212875284973843395610",10)
	//fmt.Printf("%X",num1)
	fmt.Println(elliptic_curve_key_generation.UncompressedPublicKey("6762722"))
	fmt.Println(elliptic_curve_key_generation.CompressedPublicKey("6762722"))
	*/
	/*	//sum := sha256.Sum256([]byte("hello world"))
	//fmt.Println(sum)
	//fmt.Printf("%X",sum)
	boiceAddress := elliptic_curve_key_generation.BoiceAddressGenerator("04FA066EACB1F03731C9C904F9D2C2B2D828C35C8654D4C8FFA2B7CD3B2F64B95FD7182EEF0AC14D9DD51D4356365D099827B91B4099358EB9B72F892060F89BBB")
	fmt.Println(elliptic_curve_key_generation.ValidateBoiceAddress(boiceAddress))
	//fmt.Println(elliptic_curve_key_generation.SHA256("Hello World"))
	//fmt.Println(elliptic_curve_key_generation.RIPEMD160("Hello World"))*/
	fmt.Println("+"+wallet.MnemonicGenerator("0C1E24E5917779D297E14D45F14E1A1A")+"+	")
}