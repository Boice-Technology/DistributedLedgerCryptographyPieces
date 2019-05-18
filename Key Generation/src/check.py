x = int(input())
y = int(input())
primef = 115792089237316195423570985008687907853269984665640564039457584007908834671663

def Check():
	k1 = (x**3 + 7)%primef
	k2 = (y**2)%primef
	return k1==k2

print(Check())