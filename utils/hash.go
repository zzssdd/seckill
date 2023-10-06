package utils

func StringHash(s string) int {
	bytes := []byte(s)
	ret := 0
	for _, v := range bytes {
		ret += int(v)
	}
	return ret
}
