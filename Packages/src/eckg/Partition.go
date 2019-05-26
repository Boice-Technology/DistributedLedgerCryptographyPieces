package eckg

import ("math/big";
		)

func Partition(i *big.Int) []int16 {
	 bin := IntToBin(i)
	 length := int16(len(bin))
	 var PartitionSlice []int16 
	 var j int16
	 for j = length-1; j>=0 ; j--{
		if string(bin[length-1-j]) == "1" {
			PartitionSlice = append(PartitionSlice,j)
		}
	 }
	 return PartitionSlice
}