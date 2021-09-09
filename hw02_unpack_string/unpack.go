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

	fullLen := len(name)
	countDigitsInWord := 0
	countBackslash := 0

	for _, val := range name {
		if unicode.IsDigit(val) {
			countDigitsInWord++
		} else if val == 92 {
			countBackslash++
		}
	}

	if name == "" {
		return "", nil
	} else if unicode.IsDigit(rune(name[0])) || fullLen == countDigitsInWord {
		return "", ErrInvalidString
	}

	if countBackslash > 0 {
		for index, val := range name {
			if index+1 <= len(name)-1 && val == 92 && name[index+1] != 92 {
				if val == 92 && unicode.IsDigit(rune(name[index+1])) {
					first := strings.Split(name, string(val))
					var b strings.Builder

					for i := 1; i >= 1; i-- {
						b.WriteString(first[0])
					}

					first[0] = b.String()
					*linkName = strings.Join(first, "")
				}
			}

		}
		return name, nil
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
