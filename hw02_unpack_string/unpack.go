package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(name string) (string, error) {
	linkName := &name

	if name == "" {
		return "", nil
	} else if unicode.IsDigit(rune(name[0])) {
		return "", ErrInvalidString
	}

	for index, val := range name {
		if unicode.IsDigit(val) && unicode.IsDigit(rune(name[index+1])) {
			return "", ErrInvalidString
		} else if unicode.IsDigit(val) {
			currentNumber, _ := strconv.Atoi(string(val))
			splitedName := strings.Split(name, string(val))
			lastChar := len(splitedName[0]) - 1
			if currentNumber == 0 {
				currentNumber = 1
			}
			repeatedChar := strings.Repeat(string(splitedName[0][lastChar]), currentNumber-1)
			var b strings.Builder

			for i := 1; i >= 1; i-- {
				if currentNumber == 1 {
					b.WriteString(splitedName[0][0 : len(splitedName[0])-1])
				} else {
					b.WriteString(splitedName[0])
					b.WriteString(repeatedChar)
				}
			}

			splitedName[0] = b.String()
			*linkName = strings.Join(splitedName, "")
		}
	}

	return name, nil
}
