package walletcrypto

func HexToBin(hex string) string {
	var binstr string = ""
	var i int
	for i=0;i<len(hex);i++ {
		binstr += BinEqui[HexEqui(hex[i:i+1])]
	}
	return binstr
}