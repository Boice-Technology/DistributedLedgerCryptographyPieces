package walletcrypto

func HexPrefixer(hex string) string {
	length := len(hex)
	rem := 64 - length
	var pre string = ""
	var i int
	for i=1 ; i<=rem ; i++ {
		pre += "0"
	}
	return pre + hex
}