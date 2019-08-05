/* Double.go file generator function */

package eckg

import ("os";
		"fmt";
		)

func DoublesFileGenerator(){
	file,err := os.Create("Doubles.go")
	if err!=nil{
		fmt.Println("Couldn't open the file.")
		fmt.Println(err)
	}
	defer file.Close()
	var fileString string
	fileString += "[2]string { " + "\"" + GenX +"\", \"" + GenY +"\" }, "
	var count int = 0
	var prevx string = GenX
	var prevy string = GenY
	for count<256{
		newx,newy := DoubleOperation(prevx, prevy)
		prevx = newx
		prevy = newy
		fileString += "[2]string { " + "\"" + newx +"\", \"" + newy +"\" }, "
		count++
	}
	file.WriteString(fileString)
}