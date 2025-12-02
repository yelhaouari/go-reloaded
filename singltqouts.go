package Relode

import (
	"strings"
)

func Quot(myarr []string) []string {
	cleanarr := Clean(myarr)
	cleantext := ""
	Ftext := ""
	for _, val := range cleanarr {
		cleantext += val 
	}
	
	text := []rune(cleantext)
	i := 0
	for i = 0; i < len(text)-1; i++ {
		if text[i] == '\'' {
			j := 0
			for j = i + 1; j < len(text); j++ {
				if j >  0 &&  text[j] == '\''  &&  text[j-1] == ' ' && text[j+1] == ' ' {
					Ftext +=  string(text[i]) + strings.TrimSpace(string(text[i+1:j])) + string(text[j]) 
					// Ftext += " " + string(text[i]) + strings.TrimSpace(string(text[i+1:j])) + string(text[j]) + " "
					i = j
					break
				} else if j == len(text)-1 {
					Ftext += string(text[i])
				}
			}
		} else {
			Ftext += string(text[i])
		}
	}
	Ftext += string(text[i])
	return SPlit(Ftext)
}
