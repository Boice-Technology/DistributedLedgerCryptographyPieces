/* Function returns compressed public keys */

package eckg

import ("fmt";
		"math/big";
		)
		
func CompressedPublicKey(x string,y string) string {
	y_int,_ := new(big.Int).SetString(y,16)
	two,_ := new(big.Int).SetString("2",10)
	res,_ := new(big.Int).SetString("0",10)
	res.Mod(y_int,two)
	res_str := fmt.Sprintf("%s",res)
	result := ""
	if res_str=="0"{
		result += "02 "+x
	}else{
		result += "03 "+x
	}
	return result
}