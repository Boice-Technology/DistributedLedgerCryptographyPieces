// Generate public key hash of a public key

package elliptic_curve_key_generation

func PublicKeyHash(publicKey string) string {
	publicKeyHash := RIPEMD160(SHA256(publicKey))
	return publicKeyHash
}