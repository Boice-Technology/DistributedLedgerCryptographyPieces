// Verify the signature is legit or not

package elliptic_curve_digital_signature_generation

import ("elliptic_curve_key_generation";
		"math/big";
		"fmt";
		)
		
func EllipticCurveDigitalSignatureVerifier(RSignatureComponent string, SSignatureComponent string, publicKeyX string, publicKeyY string, message string) bool {
	messageHash := elliptic_curve_key_generation.SHA256(message)
	messageHashInt,_ := new(big.Int).SetString(messageHash,16)
	inverseModuloMultiplier,_ := new(big.Int).SetString("0",10)
	SSignatureComponentInt,_ := new(big.Int).SetString(SSignatureComponent,10)
	RSignatureComponentInt,_ := new(big.Int).SetString(RSignatureComponent,10)
	pInt,_ := new(big.Int).SetString(elliptic_curve_key_generation.P,16)
	inverseModuloMultiplier.ModInverse(SSignatureComponentInt,pInt)
	multiplier1,_ := new(big.Int).SetString("0",10)
	multiplier2,_ := new(big.Int).SetString("0",10)
	multiplier1.Mul(inverseModuloMultiplier,messageHashInt)
	multiplier1.Mod(multiplier1,pInt)
	multiplier2.Mul(inverseModuloMultiplier,RSignatureComponentInt)
	multiplier2.Mod(multiplier2,pInt)
	multiplier1String := fmt.Sprintf("%d",multiplier1)
	multiplier2String := fmt.Sprintf("%d",multiplier2)
	hashPublicKeyX, hashPublicKeyY := elliptic_curve_key_generation.PublicKeyGenerator(multiplier1String)
	rPublicKeyX, rPublicKeyY := MultipliedPoint(multiplier2String, publicKeyX, publicKeyY)
	fmt.Println(multiplier2String)
	RCalculatedComponent,_ := elliptic_curve_key_generation.TwoPointAddition(hashPublicKeyX, hashPublicKeyY, rPublicKeyX, rPublicKeyY)
	fmt.Println(RCalculatedComponent)
	var result bool
	if RCalculatedComponent==RSignatureComponent{
		result = true
	}else{
		result = false
	}
	return result
}