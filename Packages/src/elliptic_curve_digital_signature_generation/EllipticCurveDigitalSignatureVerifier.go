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
	nInt,_ := new(big.Int).SetString(elliptic_curve_key_generation.N,16)
	inverseModuloMultiplier.ModInverse(SSignatureComponentInt,nInt)
	multiplier1,_ := new(big.Int).SetString("0",10)
	multiplier2,_ := new(big.Int).SetString("0",10)
	multiplier1.Mul(inverseModuloMultiplier,messageHashInt)
	multiplier1.Mod(multiplier1,nInt)
	multiplier2.Mul(inverseModuloMultiplier,RSignatureComponentInt)
	multiplier2.Mod(multiplier2,nInt)
	multiplier1String := fmt.Sprintf("%d",multiplier1)
	multiplier2String := fmt.Sprintf("%d",multiplier2)
	hashPublicKeyX, hashPublicKeyY := elliptic_curve_key_generation.PublicKeyGenerator(multiplier1String)
	hashPublicKeyXInt,_ := new(big.Int).SetString(hashPublicKeyX,10)
	hashPublicKeyYInt,_ := new(big.Int).SetString(hashPublicKeyY,10)
	hashPublicKeyXInt.Mod(hashPublicKeyXInt,nInt)
	hashPublicKeyYInt.Mod(hashPublicKeyYInt,nInt)
	hashPublicKeyX = fmt.Sprintf("%d",hashPublicKeyXInt)
	hashPublicKeyY = fmt.Sprintf("%d",hashPublicKeyYInt)
	publicKeyXInt,_ := new(big.Int).SetString(publicKeyX,10)
	publicKeyYInt,_ := new(big.Int).SetString(publicKeyY,10)
	publicKeyXInt.Mod(publicKeyXInt,nInt)
	publicKeyYInt.Mod(publicKeyYInt,nInt)
	publicKeyX = fmt.Sprintf("%d",publicKeyXInt)
	publicKeyY = fmt.Sprintf("%d",publicKeyYInt)
	rPublicKeyX, rPublicKeyY := MultipliedPoint(multiplier2String, publicKeyX, publicKeyY)
	rPublicKeyXInt,_ := new(big.Int).SetString(rPublicKeyX,10)
	rPublicKeyYInt,_ := new(big.Int).SetString(rPublicKeyY,10)
	rPublicKeyXInt.Mod(rPublicKeyXInt,nInt)
	rPublicKeyYInt.Mod(rPublicKeyYInt,nInt)
	rPublicKeyX = fmt.Sprintf("%d",rPublicKeyXInt)
	rPublicKeyY = fmt.Sprintf("%d",rPublicKeyYInt)
	RCalculatedComponent,_ := elliptic_curve_key_generation.TwoPointAddition( rPublicKeyX, rPublicKeyY,hashPublicKeyX, hashPublicKeyY)
	RCalculatedComponentInt,_ := new(big.Int).SetString(RCalculatedComponent,10)
	RCalculatedComponentInt.Mod(RCalculatedComponentInt,nInt)
	RCalculatedComponent = fmt.Sprintf("%d",RCalculatedComponentInt)
	var result bool
	if RCalculatedComponent==RSignatureComponent{
		result = true
	}else{
		result = false
	}
	return result
}