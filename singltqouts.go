package Relode

import (
	"strings"
)

func Quot(myarr []string) []string {
	cleanarr := Clean(myarr)
	cleantext := ""
	Ftext := ""
	for _, val := range cleanarr {
		cleantext += val + " "
	}
	text := []rune(cleantext)
	i := 0
	for i = 0; i < len(text)-1; i++ {
		if text[i] == '\'' {
			j := 0
			for j = i + 1; j < len(text); j++ {
				if text[j] == '\'' {
					Ftext += " " + string(text[i]) + strings.TrimSpace(string(text[i+1:j])) + string(text[j]) + " "
					i = j
					break
				} else if j == len(text)-1 {
					Ftext += string(text[j-1])
				}
			}
		} else {
			Ftext += string(text[i])
		}
	}
	return strings.Split(Ftext, " ")
}
