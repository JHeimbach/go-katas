package kata

func EquableTriangle(a, b, c int) bool {
	p := a + b + c
	return 16*p == (p-2*a)*(p-2*b)*(p-2*c)
}
