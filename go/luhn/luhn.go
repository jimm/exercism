package luhn

import (
	"strconv"
	"strings"
)

func Valid(s string) bool {
	if strings.TrimSpace(s) == "" {
		return false
	}
	return checksum(toUint64(s)) == 0
}

func AddCheck(s string) string {
	return ""
}

func toUint64(s string) uint64 {
	i, _ := strconv.ParseUint(strings.Replace(s, " ", "", -1), 10, 64)
	return i
}

func checksum(i uint64) int {
	csum := 0
	for i != 0 {
		first := int(i % 10)
		second := int((i / 10) % 10)
		i /= 100
		csum += wrap(first) + wrap(second*2)
	}
	return csum % 10
}

func wrap(i int) int {
	if i >= 10 {
		return i - 9
	}
	return i
}
