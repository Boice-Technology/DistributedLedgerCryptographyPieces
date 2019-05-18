package eckg

import ("math/big";
		"os";
		)

func MultiplesFileGenerator(){
	var prevx,_ = new(big.Int).SetString(GenX,10)
	var prevy,_ = new(big.Int).SetString(GenY,10)
	var x,_ = new(big.Int).SetString("0",10)
	var y,_ = new(big.Int).SetString("0",10)
	var j int16 = int16(1)
	var filestr string = ""
	filestr += "[2]string{\"" + GenX + "\" , "
	filestr += "\"" + GenY + "\" }"
	for j<=256 {
		slope := SlopeOfTangent(prevx,prevy)
		x = XCoordinate(slope,prevx)
		y = YCoordinate(slope,prevx,prevy,x)
		filestr += ", [2]string{ " + "\"" + x.String() + "\" , "
		filestr += "\"" + y.String() + "\" }"
		prevx = x
		prevy = y
		j++
	}
	file,err := os.Create("Multiples.go")
	if err!=nil{
		return
	}
	defer file.Close()
	file.WriteString(filestr)
}