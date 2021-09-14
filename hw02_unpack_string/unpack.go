package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

func isDigit(char int32) bool {
	if char >= 48 && char <= 57 {
		return true
	}

	return false
}

var ErrInvalidString = errors.New("invalid string")

func Unpack(name string) (string, error) {
	if name == "" {
		return "", nil
	} else if isDigit(int32(name[0])) {
		return "", ErrInvalidString
	}

	for index, val := range name {
		if isDigit(val) && (isDigit(int32(name[index+1])) && index+1 != len(name)) {
			return "", ErrInvalidString
		} else if isDigit(val) {
			currentNumber, _ := strconv.Atoi(string(val))
			splitedName := strings.SplitN(name, string(val), 2)

			lastChar := len(splitedName[0]) - 1

			repeatedChar := strings.Repeat(string(splitedName[0][lastChar]), currentNumber)

			var b strings.Builder

			b.WriteString(splitedName[0][0 : len(splitedName[0])-1])
			if currentNumber > 1 {
				b.WriteString(repeatedChar)
			}

			splitedName[0] = b.String()
			name = strings.Join(splitedName, "")
		}
	}

	return name, nil
}
