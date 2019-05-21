package walletcrypto

func Mnemonic(entropyString string) string {
	checksum := Checksum(entropyString)
	entropyString += checksum
	entropyString = HexToBin(entropyString)
}