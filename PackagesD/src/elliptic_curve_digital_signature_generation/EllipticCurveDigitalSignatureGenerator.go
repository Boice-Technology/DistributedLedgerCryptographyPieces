// Generate elliptic curve digital signature from ephemeral private key, signing private key and the message on which signature has to be made

package elliptic_curve_digital_signature_generation

import ("math/big";
		"elliptic_curve_key_generation";
		"fmt";
		)

func EllipticCurveDigitalSignatureGenerator(ephemeralPrivateKey string, signingPrivateKey string, message string) (string, string) {
	nInt,_ := new(big.Int).SetString(elliptic_curve_key_generation.N,16)
	messageHash := elliptic_curve_key_generation.SHA256(message)
	messageHashInt,_ := new(big.Int).SetString(messageHash,16)
	RSignatureComponent, _ := elliptic_curve_key_generation.PublicKeyGenerator(ephemeralPrivateKey)
	RSignatureComponentInt,_ := new(big.Int).SetString(RSignatureComponent,10)
	RSignatureComponentInt.Mod(RSignatureComponentInt,nInt)
	RSignatureComponent = fmt.Sprintf("%d",RSignatureComponentInt)
	ephemeralPrivateKeyInt,_ := new(big.Int).SetString(ephemeralPrivateKey,10)
	signingPrivateKeyInt,_ := new(big.Int).SetString(signingPrivateKey,10)
	multiplier,_ := new(big.Int).SetString("0",10)
	multiplier.Mul(RSignatureComponentInt,signingPrivateKeyInt)
	inverseModuloMultiplier,_ := new(big.Int).SetString("0",10)
	inverseModuloMultiplier.ModInverse(ephemeralPrivateKeyInt,nInt)
	sum,_ := new(big.Int).SetString("0",10)
	sum.Add(messageHashInt,multiplier)
	SSignatureComponentInt,_ := new(big.Int).SetString("0",10)
	SSignatureComponentInt.Mul(sum,inverseModuloMultiplier)
	SSignatureComponentInt.Mod(SSignatureComponentInt,nInt)
	SSignatureComponent := fmt.Sprintf("%d",SSignatureComponentInt)
	return RSignatureComponent, SSignatureComponent
}