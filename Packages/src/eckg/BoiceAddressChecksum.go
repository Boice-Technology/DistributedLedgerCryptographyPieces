/* To compute the checksum of the boice address */

package eckg

func BoiceAddressChecksum(boiceAddress string) string {
	hash1 := Sha256(boiceAddress)
	hash2 := Sha256(hash1)
	return hash2[0:8]
}