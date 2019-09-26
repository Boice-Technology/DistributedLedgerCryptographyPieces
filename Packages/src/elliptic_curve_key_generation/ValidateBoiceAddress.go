// Validates the boice address by using checksum of the string

package elliptic_curve_key_generation

func ValidateBoiceAddress(boiceAddress string) bool {
	base58Decoded := Base58Decode(boiceAddress)
	doubleSHA256Hash := SHA256(SHA256(base58Decoded[:len(base58Decoded)-8]))
	
	if doubleSHA256Hash[0:8] == base58Decoded[len(base58Decoded)-8:] {
		return true
	}else{
		return false
	}
}
