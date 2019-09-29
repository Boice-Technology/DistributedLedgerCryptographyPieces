// Utility funciton to reverse a string

package elliptic_curve_key_generation

func Reverse(str string) string {
	var length int = len(str)
	var byteString []byte = []byte(str)
	for i := 0 ; i < length/2 ; i++ {
		byteString[i],byteString[length-i-1] = byteString[length-1-i],byteString[i]
	}
	str = string(byteString)
	return str
}