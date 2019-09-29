package main

import ("fmt";
		//"math/big"
		"elliptic_curve_key_generation";
		//"wallet";
		//"crypto/sha512";
		"elliptic_curve_digital_signature_generation";
		)
		
func main(){
	/*
	//fmt.Println(elliptic_curve_key_generation.DecimalToBinary("3"))
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
	//fmt.Println(elliptic_curve_key_generation.RIPEMD160("Hello World"))
	fmt.Println("+"+wallet.MnemonicGenerator("0C1E24E5917779D297E14D45F14E1A1A")+"+	")
	bytes := sha512.Sum512([]byte("hello world"))
	fmt.Printf("%X",bytes)
	Integer,_ := new(big.Int).SetString("A",16)
	fmt.Printf("%d",Integer)
	
	fmt.Println(elliptic_curve_key_generation.CompressedPublicKey("2"))
	fmt.Println(wallet.ChildPublicKeyGenerator("02C6047F9441ED7D6D3045406E95C07CD85C778E4B8CEF3CA7ABAC09B95C709EE5","2",0))
	fmt.Println(wallet.ChildPrivateKeyGenerator("2","2",0))
	fmt.Println(elliptic_curve_key_generation.CompressedPublicKey("47581640817870642330691603300449775897285483678155889044195286196953467712712"))
	*/
	R,S := elliptic_curve_digital_signature_generation.EllipticCurveDigitalSignatureGenerator("2","3","Hello")
	pX, pY := elliptic_curve_key_generation.PublicKeyGenerator("3")
	fmt.Println(elliptic_curve_digital_signature_generation.EllipticCurveDigitalSignatureVerifier("89565891926547004231252920425935692360644145829622209833684329913297188986597","24068917058926912194676773038160490119087835964610282300377278172857912065757","112711660439710606056748659173929673102114977341539408544630613555209775888121","25583027980570883691656905877401976406448868254816295069919888960541586679410","Hello"))
	fmt.Println(pX, pY)
	fmt.Println(R,S)
	var k []int = make([]int,1)
	k = append(k,1,2)
	fmt.Println(k)
}