/* Function returns uncompressed public keys */

package eckg

func UncompressedPublicKey(x string,y string) string {
	key := "04 "+x+" "+y
	return key
}