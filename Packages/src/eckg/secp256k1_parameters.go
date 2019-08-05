/* SECP256K1 curve specifications*/
package eckg
/* Elliptic curve cryptography of the form y^2 = x^3 + ax + b */

/* Finite field prime number P */
const P string = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F"

/* Coefficient A of the elliptic curve */
const A string = "0000000000000000000000000000000000000000000000000000000000000000"

/* Coefficient B of the elliptic curve */
const B string = "0000000000000000000000000000000000000000000000000000000000000007"

/* The base point or the generator point of the elliptic curve */
const GenX string = "79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798"
const GenY string = "483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8"

/* Order of the elliptic curve N */
const N string = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141"

/* Cofactor of the elliptic curve */
const H string = "01"