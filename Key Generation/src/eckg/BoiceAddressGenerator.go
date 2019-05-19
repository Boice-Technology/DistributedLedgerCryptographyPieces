package eckg

import (//"math/big";
		"fmt";
		"crypto/sha256";
		"golang.org/x/crypto/ripemd160"
		)

func BoiceAddressGenerator(publicKey string) string { 
	boiceAddr := sha256.Sum256([]byte(publicKey))
	shaHash := fmt.Sprintf("%x",boiceAddr)
	digest := ripemd160.New()
	digest.Write([]byte(shaHash))
	boiceAddrStr := fmt.Sprintf("%x",digest.Sum(nil))
	return boiceAddrStr
}