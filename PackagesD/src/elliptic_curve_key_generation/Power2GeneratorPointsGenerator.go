// Power 2 Generator(G) points generator, like 2^n*G

package elliptic_curve_key_generation

import ("os";
		"fmt";
		"math/big"
		)

func Power2GeneratorPointsGenerator(){
	file,error := os.Create("Power2GeneratorPoints.go")
	if error!=nil{
		return
	}
	 
	defer file.Close()
	 
	var fileString string = ""
	gxInt,_ := new(big.Int).SetString(Gx,16)
	gyInt,_ := new(big.Int).SetString(Gy,16)
	gxStr := fmt.Sprintf("%s",gxInt)
	gyStr := fmt.Sprintf("%s",gyInt)
	
	fileString += "[2]string {\"" + gxStr + "\", \"" + gyStr + "\"}, "
	prevX := gxStr
	prevY := gyStr
	
	for i := 1 ; i < 257 ; i++ {
		gxStr, gyStr = OnePointAddition(prevX, prevY)
		fileString += "[2]string {\"" + gxStr + "\", \"" + gyStr + "\"}, "
		prevX = gxStr
		prevY = gyStr
	}
	
	file.WriteString(fileString)
}