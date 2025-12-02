package Relode

import (
	"strings"
)

func Vol(myarr []string) []string {
	cleanaar := Clean(myarr)
	for i := 0; i < len(cleanaar)-1; i++ {

		if cleanaar[i] == "A" {
			count := i + 1

			for count < len(cleanaar)-1 && cleanaar[count] == " " {
				count++
			}
			if cleanaar[count][0] == 'a' ||
				cleanaar[count][0] == 'e' ||
				cleanaar[count][0] == 'i' ||
				cleanaar[count][0] == 'o' ||
				cleanaar[count][0] == 'u' ||
				cleanaar[count][0] == 'h' {
				cleanaar[i] = "An"
			}
		} else if strings.TrimSpace(cleanaar[i]) == "a" {
			count := i + 1
			for count < len(cleanaar)-1 && cleanaar[count] == " " {
				count++
			}
			if cleanaar[count][0] == 'a' ||
				cleanaar[count][0] == 'e' ||
				cleanaar[count][0] == 'i' ||
				cleanaar[count][0] == 'o' ||
				cleanaar[count][0] == 'u' ||
				cleanaar[count][0] == 'h' {
				cleanaar[i] = "an "
			}
		}

		if cleanaar[i] == "a" {
			count := i + 1
			for count < len(cleanaar)-1 && cleanaar[count] == " " {
				count++
			}
			if cleanaar[count][0] == 'a' ||
				cleanaar[count][0] == 'e' ||
				cleanaar[count][0] == 'i' ||
				cleanaar[count][0] == 'o' ||
				cleanaar[count][0] == 'u' ||
				cleanaar[count][0] == 'h' {
				cleanaar[i] = "an"
			}
		} else if strings.TrimSpace(cleanaar[i]) == "a" {
			count := i + 1
			for count < len(cleanaar)-1 && cleanaar[count] == " " {
				count++
			}
			if cleanaar[count][0] == 'a' ||
				cleanaar[count][0] == 'e' ||
				cleanaar[count][0] == 'i' ||
				cleanaar[count][0] == 'o' ||
				cleanaar[count][0] == 'u' ||
				cleanaar[count][0] == 'h' {
				cleanaar[i] = "an "
			}
		}
	}
	return cleanaar
}
