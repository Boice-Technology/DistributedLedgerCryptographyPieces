package transactions

import ("math/big";
		"eckg";
		)
		
func ProductPoint(u *big.Int, x1 *big.Int, y1 *big.Int) (string, string) {
	partition := eckg.Partition(u)
	largest := partition[0]
	multiples := Multiples(largest,x1,y1)
	x, y := multiples[largest][0], multiples[largest][1]
	var i int 
	for i=1; i<len(partition) ; i++ {
		x, y = Add(x,y,multiples[partition[i]][0],multiples[partition[i]][1])
	}
	return x,y
}