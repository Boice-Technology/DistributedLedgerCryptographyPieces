package eckg

/*
	The SECP256K1 with N as order, curve is y^2 = x^3 + 7 (mod PrimeF) 
	where A and B are curve coefficients.
	PrimeF is Prime field of a curve.
	N is the order of curve, ie. Private key can range between 0 to N-1.
	(GenX,GenY) is a generator for the curve which lies on a curve having relation, Pub = Pri * Gen
	where Pub is a public key and Pri is private key
*/


const PrimeF string = "115792089237316195423570985008687907853269984665640564039457584007908834671663"

const A string = "0"

const B string = "7"

const GenX string = "55066263022277343669578718895168534326250603453777594175500187360389116729240"

const GenY string = "32670510020758816978083085130507043184471273380659243275938904335757337482424"

const N string = "115792089237316195423570985008687907852837564279074904382605163141518161494337"
