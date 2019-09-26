// Genearte Boice address from the public key

package elliptic_curve_key_generation

func BoiceAddressGenerator(publicKey string) string {
	publicKeyHash := PublicKeyHash(publicKey)
	payload := "0x00" + publicKeyHash
	doubleSHA256Hash := SHA256(SHA256(payload))
	checksum := doubleSHA256Hash[:8]
	inputString := payload + checksum
	boiceAddress := Base58Encode(inputString)
	return boiceAddress
}