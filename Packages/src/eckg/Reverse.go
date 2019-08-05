package eckg
		
func Reverse(s string) string {
	length := len(s)
	var count int = 0 
	var s1 string = ""
	for count<length{
		s1 += s[length-1-count:length-count]
		count++
	}
	return s1
}