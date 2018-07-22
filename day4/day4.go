package day4

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func Run() int {

	phrases := []string{
		"aa bb cc dd ee",
		"aa bb cc dd aa",
		"aa bb cc dd aaa",
	}

	dat, err := ioutil.ReadFile("puzzleinput.txt")
	if err != nil {
		return -1
	}

	phrases = strings.Split(string(dat), "\n")

	valid := 0

	for _, phrase := range phrases {
		if isValid(phrase) {
			valid++
			//fmt.Println(phrase)
		} else {
			fmt.Println(phrase)
		}
	}

	return valid
}

func isValid(phrase string) bool {

	codes := strings.Split(phrase, " ")

	for index, code := range codes {

		for index2, code2 := range codes {
			if index == index2 {
				continue
			}
			if SortString(code) == SortString(code2) {
				return false
			}
		}
	}
	return true
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
