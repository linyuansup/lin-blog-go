package utils

import "strconv"

func StrToInt(s string) int {
	data, _ := strconv.Atoi(s)
	return data
}
