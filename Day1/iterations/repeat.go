package iterations

func Repeat(char string, timeS int) string {
	var res string
	for i := 0; i < timeS; i++ {
		res += char
	}
	return res
}
