package walletcrypto

import ("crypto/sha512";
		"fmt";
		)

func Seed(mnemonic string, passphrase string) string {
	hashed := mnemonic + "mnemonic" + passphrase
	var i int
	hashedBytes := []byte(hashed)
	for i=1 ; i<=2048 ; i++ {
		sizedBytes := sha512.Sum512(hashedBytes)
		hashedBytes = sizedBytes[:]
	}
	seed := fmt.Sprintf("%x",hashedBytes)
	return seed
}