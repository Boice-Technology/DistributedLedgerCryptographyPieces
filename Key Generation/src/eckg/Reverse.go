package eckg

import "math"

func Reverse(s string) string{
	str := []rune(s)
	i := 0
	j := len(str)-1
	for i < int(math.Ceil(float64(len(str))/2)) {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
	return string(str)
}