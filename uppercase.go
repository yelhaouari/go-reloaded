package Relode

import (
	"strconv"
	"strings"
)

func word(str string, start int, end int) string {
	starttext := 0
	endtext := 0

	for i := 0; i <= len(str[start:end])-1; i++ {
		if str[start:end][i] != ' ' {
			starttext = i
		}
	}
	for j := len(str[start:end]) - 1; j >= 0; j-- {
		if str[start:end][j] != ' ' {
			endtext = j
		}
	}
	return str[start:end][endtext : starttext+1]
}

func Nup(myarr []string) []string {
	cleanarr := Clean(myarr)
	for ind, val := range cleanarr {
		if strings.Contains((val), "up, ") {

			mynum, err := strconv.Atoi(strings.TrimSpace((val[strings.Index(val, ",")+1 : strings.Index(val, ")")])))
			if err != nil {
				continue
			} else {
				if strings.TrimSpace(val) == "(up, "+word(val, strings.Index(val, ",")+1, strings.Index(val, ")"))+")" {

					if mynum < 0 {
						cleanarr[ind] = ""
						continue
					}
					count := 0
					for i := ind; i >= 0; i-- {
						count++
						if cleanarr[i] != " " {
							cleanarr[i] = strings.ToUpper(cleanarr[i])
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
	return cleanarr
}

func Up(myarr []string) []string {
	cleanarr := Clean(myarr)

	for ind, val := range cleanarr {
		if (strings.TrimSpace(val)) == "(up)" {

			for i := ind - 1; i >= 0; i-- {
				if cleanarr[i] != "\n" && cleanarr[i] != " " && cleanarr[i] != "" {
					cleanarr[i] = strings.ToUpper(cleanarr[i])
					cleanarr[ind] = " "
					break
				}
			}
			cleanarr[ind] = ""
		}
	}

	return cleanarr
}
