// Generate a mnemonic from a 128 bit entropy

package wallet

import ("elliptic_curve_key_generation";
		)

func MnemonicGenerator(entropy string) string {
	sha256Hash := elliptic_curve_key_generation.SHA256(entropy)
	randomSequence := entropy + sha256Hash[:1]
	binaryString := HexadecimalToBinary(randomSequence)
	mnemonic := ""
	for i := 1 ; i <= 12 ; i++ {
		binaryNumber := binaryString[11*(i-1):11*i]
		decimalNumber := BinaryToDecimal(binaryNumber)
		mnemonic += " " + MnemonicWords[decimalNumber-1]
	}
	mnemonic = mnemonic[1:]
	return mnemonic
}