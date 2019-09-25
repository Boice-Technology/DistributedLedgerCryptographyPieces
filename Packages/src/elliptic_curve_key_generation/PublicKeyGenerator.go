// Public key generator from private key of the (publicX, publicY) = privateKey * (Gx,Gy)

package elliptic_curve_key_generation

import ("fmt";
		"math/big";
		)

func PublicKeyGenerator(privateKey string) (string, string) {
	privateKeyInt,_ := new(big.Int).SetString(privateKey,10)
	binaryPrivateKey := fmt.Sprintf("%b",privateKeyInt)
	revBinaryPrivateKey := Reverse(binaryPrivateKey)
	length := len(revBinaryPrivateKey)
	var listBits []int = make([]int,1)
	for i := 0 ; i < length ; i++ {
		if revBinaryPrivateKey[i]==49{
			listBits = append(listBits,i)
		}
	}
	lengthList := len(listBits)
	prevx := Power2Generator[listBits[1]][0]
	prevy := Power2Generator[listBits[1]][1]
	
	if lengthList == 2 {
		return prevx, prevy
	}else{
		for i := 2 ; i < lengthList ; i++ {
			x := Power2Generator[listBits[i]][0]
			y := Power2Generator[listBits[i]][1]
			prevx, prevy = TwoPointAddition(prevx, prevy, x, y)
		}
	}
	
	return prevx, prevy
}