package core

func Exponent (a,n uint64) uint64  {
	result := uint64(1)
	for i := n ; i > 0; i >>= 1 {
		if i&1 != 0 {
			result *= a
		}
		a *= a
	}
	return result
}