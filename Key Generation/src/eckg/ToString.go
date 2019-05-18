package eckg

/*
	Utility function for converting from int16 to a string
*/

func ToString(a int16) string {
	var str string = ""
	for a!=0 {
		k := a%10
		str+=string(k+48)
		a/=10
	}
	return Reverse(str)
}