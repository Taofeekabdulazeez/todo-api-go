package util

import "strconv"

func ParseUInt(s string) uint {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return uint(i)
}
