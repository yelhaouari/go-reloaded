package Relode

import (
	"strconv"
	"strings"
)

func Nlow(myarr []string) []string {
	cleanarr := Clean(myarr)
	for ind, val := range cleanarr {
		if strings.Contains((val), "low, ") {

			mynum, err := strconv.Atoi(strings.TrimSpace((val[strings.Index(val, ",")+1 : strings.Index(val, ")")])))
			if err != nil {
				continue
			} else {
				if strings.TrimSpace(val) == "(low, "+word(val, strings.Index(val, ",")+1, strings.Index(val, ")"))+")" {

					if mynum < 0 {
						cleanarr[ind] = ""
						continue
					}

					count := 0
					for i := ind; i >= 0; i-- {
						count++
						if cleanarr[i] != " " {
							cleanarr[i] = strings.ToLower(cleanarr[i])
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

func Low(myarr []string) []string {
	cleanarr := Clean(myarr)

	for ind, val := range cleanarr {
		if (strings.TrimSpace(val)) == "(low)" {
			for i := ind - 1; i >= 0; i-- {
				if cleanarr[i] != "\n" && cleanarr[i] != " " && cleanarr[i] != "" {
					cleanarr[i] = strings.ToLower(cleanarr[i])
					cleanarr[ind] = ""
					break
				}
			}
			cleanarr[ind] = ""
		}
	}
	return cleanarr
}
