package stringUnpack

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func StringUnpack(s string) (string, error) {
	var builtString strings.Builder
	var multiplier strings.Builder
	var previousChar rune
	var isEscaped bool
	for _, r := range s {
		if r == '\\' && !isEscaped {
			isEscaped = true
			continue
		}
		if isEscaped {
			builtString.WriteRune(r)
			previousChar = r
			isEscaped = false
			continue
		}
		if unicode.IsDigit(r) {
			if previousChar > 0 {
				multiplier.WriteRune(r)
			} else {
				return "", fmt.Errorf("invalid input string format")
			}
		} else {
			if m, err := strconv.Atoi(multiplier.String()); err == nil && m > 0 && previousChar > 0 {
				builtString.WriteString(strings.Repeat(string(previousChar), m-1))
				multiplier.Reset()
			}
			builtString.WriteRune(r)
			previousChar = r
		}
	}
	if m, err := strconv.Atoi(multiplier.String()); err == nil && m > 0 && previousChar > 0 {
		builtString.WriteString(strings.Repeat(string(previousChar), m-1))
	}
	if res := builtString.String(); len(res) > 0 {
		return builtString.String(), nil
	} else {
		return "", fmt.Errorf("invalid input string format")
	}
}
