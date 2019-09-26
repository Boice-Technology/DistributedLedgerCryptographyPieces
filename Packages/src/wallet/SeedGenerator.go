// Generate seed from mnemonic and passphrase provided by the user

package wallet

func SeedGenerator(mnemonic string, passphrase string) string {
	inputString := mnemonic + "mnemonic" + passphrase
	var intermediateString string = inputString
	for i := 1 ; i <= 2048 ; i++{
		intermediateString = SHA512(intermediateString)
	}
	var seed string = intermediateString
	return seed
}