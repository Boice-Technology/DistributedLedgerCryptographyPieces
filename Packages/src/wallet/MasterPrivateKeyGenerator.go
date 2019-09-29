// Genearate master private key and master chain code from the seed of the wallet

package wallet

func MasterPrivateKeyGenerator(walletSeed string) (string,string){
	sha512Hash := SHA512(walletSeed)
	masterPrivateKey := sha512Hash[0:64]
	masterChainCode := sha512Hash[64:]
	return masterPrivateKey, masterChainCode
}