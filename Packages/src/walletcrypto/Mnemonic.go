package walletcrypto

func Mnemonic(entropyString string) string {
	checksum := Checksum(entropyString)
	entropyString += checksum
	entropyString = HexToBin(entropyString)
	var mnemonic string = MnemonicWords[BinToDec(entropyString[0:11])]
	var i int
	for i=1; i<=11 ;i++ {
		mnemonic += " " + MnemonicWords[BinToDec(entropyString[11*i:11*(i+1)])]
	}
	return mnemonic
}