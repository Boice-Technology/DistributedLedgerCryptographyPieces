// Parameters of the standard SECP256K1 curve of the form y^2 = x^3 + a*x + b

package elliptic_curve_key_generation

// Prime order(p) 
const P string = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F"

// Coefficient A(a) of curve
const A string = "0000000000000000000000000000000000000000000000000000000000000000"

// Coefficient B(b) of curve
const B string = "0000000000000000000000000000000000000000000000000000000000000007"

// X - Co-ordinate of the Generator point(G) for y^2 = x^3 + 7
const Gx string = "79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798"

// Y - Co-ordinate of the Generator point(G) for y^2 = x^3 + 7
const Gy string = "483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8"

// Order(n) of the Generator point(G)
const N string = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141"

// Cofactor(h) of the Generator point(G)
const H string = "01"