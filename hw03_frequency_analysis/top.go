package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordform struct {
	Word    string
	Counter int
}

var wordforms []wordform

func Top10(text string) []string {
	if len(text) == 0 {
		return nil
	}

	m := make(map[string]int)

	newWithoutComma := strings.ReplaceAll(text, ",", " ")
	newWithoutDot := strings.ReplaceAll(newWithoutComma, ".", " ")
	newWithoutExclamationMark := strings.ReplaceAll(newWithoutDot, "!", " ")
	newWithoutQuote := strings.ReplaceAll(newWithoutExclamationMark, ":", " ")
	newWithoutColon := strings.ReplaceAll(newWithoutQuote, "\"", " ")
	newClear := strings.ReplaceAll(newWithoutColon, " - ", " ")

	newToLower := strings.ToLower(newClear)
	newSplited := strings.Fields(newToLower)

	for _, val := range newSplited {
		_, ok := m[val]
		if ok {
			m[val]++
		} else {
			m[val] = 1
		}
	}

	for k, v := range m {
		wordforms = append(wordforms, wordform{k, v})
	}

	sort.Slice(wordforms, func(i, j int) bool {
		return wordforms[i].Counter > wordforms[j].Counter
	})

	s := make([]string, 0, 10)

	for _, kv := range wordforms[:10] {
		s = append(s, kv.Word)
	}

	return s
}
