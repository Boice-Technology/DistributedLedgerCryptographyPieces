/* Function to generate boice address */

package eckg

func BoiceAddressGenerator(publicKey string) string {
	hash := Sha256(publicKey)
	hash = Ripemd160(hash)
	hash = "0x00"+hash
	checksum := BoiceAddressChecksum(hash)
	hash = hash+checksum
	boice_addr := Base58Encode(hash)
	return boice_addr
}