package eckg

import ("fmt";
		"crypto/sha256";
		"golang.org/x/crypto/ripemd160";
		"github.com/btcsuite/btcutil/base58"
		)

func BoiceAddressGenerator(publicKey string) string { 
	boiceAddr := sha256.Sum256([]byte(publicKey))
	shaHash := fmt.Sprintf("%x",boiceAddr)
	digest := ripemd160.New()
	digest.Write([]byte(shaHash))
	boiceAddrStr := fmt.Sprintf("%x",digest.Sum(nil))
	boiceAddrStr = "0x00" + boiceAddrStr
	checksum := Checksum(boiceAddrStr)
	boiceAddrStr += checksum
	base58EncStr := base58.Encode([]byte(boiceAddrStr))
	base58EncStr = "1" + base58EncStr
	return base58EncStr
}