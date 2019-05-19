package eckg

import ("github.com/btcsuite/btcutil/base58";
		"fmt";
		)

func CheckBoiceAddress(base58Enc string) bool {
	decodedBytes := base58.Decode(base58Enc)
	decodeString := fmt.Sprintf("%s",decodedBytes)
	stringChecksum := Checksum(decodeString[:len(decodeString)-4])
	if stringChecksum == decodeString[len(decodeString)-4:]{
		return true
	}
	return false
}