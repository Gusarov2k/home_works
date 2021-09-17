package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

func isDigit(char uint8) bool {
	if char >= 48 && char <= 57 {
		return true
	}
	return false
}

var ErrInvalidString = errors.New("invalid string")

func Unpack(name string) (string, error) {
	if name == "" {
		return "", nil
	} else if isDigit(name[0]) {
		return "", ErrInvalidString
	}

	var b strings.Builder

	for i := 0; i < len(name); i++ {
		if isDigit(name[i]) && isDigit(name[i+1]) {
			return "", ErrInvalidString
		} else if isDigit(name[i]) {
			currentNumber, _ := strconv.Atoi(string(name[i]))

			if currentNumber > 1 {
				repeatedChar := strings.Repeat(string(name[i-1]), currentNumber-1)
				b.WriteString(repeatedChar)
			} else if currentNumber == 0 {
				cutString := b.String()[:i-1]
				b.Reset()
				b.WriteString(cutString)
			}
		} else {
			b.WriteString(string(name[i]))
		}
	}

	return b.String(), nil
}
