package sum

// Ints return the sum of a list of integers
func Ints(vs ...int) int {
	return ints(vs)
}

// Recursive function
func ints(vs []int) int {
	if len(vs) == 0 {
		return 0
	}
	return ints(vs[1:]) + vs[0]
}
