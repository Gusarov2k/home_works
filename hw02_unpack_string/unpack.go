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
			valAsInt, _ := strconv.Atoi(string(val))
			first := strings.Split(name, string(val))
			lastChar := len(first[0]) - 1
			if valAsInt == 0 {
				valAsInt = 1
			}
			repeatedChar := strings.Repeat(string(first[0][lastChar]), valAsInt-1)
			var b strings.Builder

			for i := 1; i >= 1; i-- {
				if valAsInt == 1 {
					b.WriteString(first[0][0 : len(first[0])-1])
				} else {
					b.WriteString(first[0])
					b.WriteString(repeatedChar)
				}
			}

			first[0] = b.String()
			*linkName = strings.Join(first, "")
		}
	}

	return name, nil
}
