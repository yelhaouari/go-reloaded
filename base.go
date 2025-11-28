package Relode

import (
	"strconv"
	"strings"
)

func Base(myarr []string) []string {
	cleanaarr := Clean(myarr)
	for ind := 0; ind < len(cleanaarr); ind++ {
		cleanaarr[ind] = strings.TrimSpace(cleanaarr[ind])
		if ind == 0 && cleanaarr[ind] == "(hex)" || ind == 0 && cleanaarr[ind] == "(bin)" {
			cleanaarr[ind] = ""
			continue
		}
		if cleanaarr[ind] == "(hex)" {
			j := ind - 1
			for cleanaarr[j] == "" && j >= 0 {
				j--
			}
			decimalValue, err := strconv.ParseInt(cleanaarr[j], 16, 64)
			if err != nil {
				cleanaarr[ind] = ""
				continue
			} else {
				cleanaarr[j] = strconv.Itoa(int(decimalValue))
			}
			cleanaarr[ind] = ""
		}
		if cleanaarr[ind] == "(bin)" {
			j := ind - 1
			for cleanaarr[j] == "" && j >= 0 {
				j--
			}
			decimalValue, err := strconv.ParseInt(cleanaarr[j], 2, 64)
			if err != nil {
				cleanaarr[ind] = ""
				continue
			} else {
				cleanaarr[j] = strconv.Itoa(int(decimalValue))
			}
			cleanaarr[ind] = ""
		}
	}

	return cleanaarr
}
