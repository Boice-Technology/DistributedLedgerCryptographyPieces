package eckg

import ("fmt";
		"math/big";
		)
		
func PublicKeyGenerator(PrivateKey string) (string,string) {
	pri_int,_ := new(big.Int).SetString(PrivateKey,16)
	pri_bin := fmt.Sprintf("%b",pri_int)
	pri_bin_rev := Reverse(pri_bin)
	length := len(pri_bin_rev)
	var count int = 0
	var Pos []int = make([]int,0,1)
	for count<length{
		if pri_bin_rev[count]==49{
			Pos = append(Pos,count)
		}
		count++
	}
	len_pos := len(Pos)
	var resx string
	var resy string
	if len_pos==1{
		return Doubles[Pos[0]][0],Doubles[Pos[0]][1]
	}else if len_pos>1 {
		var x []string = make([]string,0,1)
		var y []string = make([]string,0,1)
		count = 0
		for count<len_pos{
			x = append(x,Doubles[Pos[count]][0])
			y = append(y,Doubles[Pos[count]][1])
			count++
		}
		count = 1
		resx = x[0]
		resy = y[0]
		for count<len_pos{
			resx,resy = EllipticAdd(resx,resy,x[count],y[count])
			count++
		}
	}
	return resx,resy
}