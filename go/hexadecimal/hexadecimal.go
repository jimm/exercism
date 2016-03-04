package hexadecimal

import (
	"errors"
	"strings"
)

func ParseHex(s string) (int64, error) {
	if s == "" {
		return 0, errors.New("huh?")
	}
	var val int64
	for _, r := range strings.ToUpper(s) {
		val *= 16
		if r >= '0' && r <= '9' {
			val += int64(r - '0')
		} else if r >= 'A' && r <= 'F' {
			val += int64(r - 'A' + 10)
		} else {
			return 0, errors.New("huh?")
		}
		if val < 0 {
			return 0, errors.New("overflow")
		}
	}
	return val, nil
}

//
// HandleErrors takes a list of inputs for ParseHex and returns a matching list
// of error cases.  It must call ParseHex on each input, handle the error result,
// and put one of three strings, "none", "syntax", or "range" in the result list
// according to the error.
func HandleErrors(hexen []string) []string {
	errors := []string{}
	for _, hex := range hexen {
		_, err := ParseHex(hex)
		if err == nil {
			errors = append(errors, "none")
		} else {
			switch err.Error() {
			case "huh?":
				errors = append(errors, "syntax")
			case "overflow":
				errors = append(errors, "range")
			}
		}
	}
	return errors
}
