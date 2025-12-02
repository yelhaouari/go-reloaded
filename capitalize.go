package Relode

import (
	"strconv"
	"strings"
	"unicode"
)

func Ncap(myarr []string) []string {
	cleanarr := Clean(myarr)
	for ind, val := range cleanarr {
		if strings.Contains(val, "(cap, ") {
			mynum, err := strconv.Atoi(strings.TrimSpace((val[strings.Index(val, ",")+1 : strings.Index(val, ")")])))
			if err != nil {
				continue
			} else {
				if strings.TrimSpace(val) == "(cap, "+word(val, strings.Index(val, ",")+1, strings.Index(val, ")"))+")" {
					if mynum < 0 {
						cleanarr[ind] = ""
						continue
					} else {
						count := 0
						for i := ind; i >= 0; i-- {
							count++
							if cleanarr[i] != "\n" && cleanarr[i] != " " && cleanarr[i] != "" {

								stringe := []rune(strings.TrimSpace(cleanarr[i]))
								stringe[0] = unicode.ToTitle(stringe[0])
								cleanarr[i] = string(stringe[0]) + strings.ToLower(string(stringe[1:]))
							}
							if count-1 == ind || count-1 == int(mynum) {
								break
							}
						}
						cleanarr[ind] = ""
					}
				}
			}
		}
	}
	return cleanarr
}

func Cap(myarr []string) []string {
	cleanarr := Clean(myarr)
	for ind, val := range cleanarr {

		if strings.TrimSpace(strings.TrimSpace(val)) == "(cap)" {
			for i := ind - 1; i >= 0; i-- {
				if cleanarr[i] != "\n" && cleanarr[i] != " " && cleanarr[i] != "" {
					stringe := []rune(strings.TrimSpace(cleanarr[i]))
					stringe[0] = unicode.ToTitle(stringe[0])
					cleanarr[i] = string(stringe[0]) + strings.ToLower(string(stringe[1:]))
					cleanarr[i] = cleanarr[i] + cleanarr[ind][strings.Index(val, ")")+1:]
					cleanarr[ind] = ""
					break
				}
			}
			cleanarr[ind] = ""
		}
	}

	return cleanarr
}
