package yinyang

import "strings"

func ChangeDateLocation(s string) string {
	var res string
	var n int
	for i := 1; s[i] != '{'; i++ {
		n++
	}
	n += 2
	res += string('•')
	res += " "
	for i := n; i < len(s); i++ {
		if s[i] == ',' {
			if i-1 >= 0 && s[i-1] != '"' {
				res += "\n"
				res += string('•')
				res += " "
			} else {
				res += string(s[i])
				res += " "
			}
		} else if s[i] == '-' && i != 0 && i+1 < len(s) {
			if s[i-1] >= '0' && s[i-1] <= '9' {
				if s[i+1] >= '0' && s[i+1] <= '9' {
					res += "."
				} else {
					res += string(s[i])
				}
			} else {
				res += string(s[i])
			}
		} else if s[i] == '{' || s[i] == '}' || s[i] == '"' {
			continue
		} else if s[i] == ':' {
			res += string(s[i])
			res += " "
		} else {
			res += string(s[i])
		}
	}
	res = strings.ReplaceAll(res, "_", " ")
	res = strings.ReplaceAll(res, "-", " ")
	res = strings.ReplaceAll(res, "[", "")
	res = strings.ReplaceAll(res, "]", "")
	res = strings.ReplaceAll(res, ", ", " | ")
	return res
}
