package main

func replaceSpaces(str string) string {
	res := ""
	n := len(str)
	for i := 0; i < n; i++ {
		c := str[i]
		if c == ' ' {
			res += "%20"
		} else {
			res += string(c)
		}
	}
	return res
}
