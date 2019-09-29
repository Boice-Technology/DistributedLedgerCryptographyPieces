package main

import ("fmt";
		//"math/big"
		"elliptic_curve_key_generation";
		//"wallet";
		//"crypto/sha512";
		"elliptic_curve_digital_signature_generation";
		)
		
func main(){
	R, S := elliptic_curve_digital_signature_generation.EllipticCurveDigitalSignatureGenerator("2","4","hello1")
	px, py := elliptic_curve_key_generation.PublicKeyGenerator("4")
	fmt.Println(R,S,px,py)
	fmt.Println(elliptic_curve_digital_signature_generation.EllipticCurveDigitalSignatureVerifier(R,S,px,py,"hello1"))
}