package Relode

import (
	"fmt"
	"os"
	"strings"
)

func normalizeSpaces(s string) string {
	lastcopy := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' || (i > 0 && s[i-1] != ' ') {
			lastcopy += string(s[i])
		}
	}

	if len(lastcopy) > 0 && lastcopy[0] == ' ' {
		lastcopy = lastcopy[1:]
	}
	return lastcopy
}

func proses(cleanarr []string) []string {
	myarr := Clean(cleanarr)
	myarr = Base(myarr)
	myarr = Clean(myarr)
	myarr = Cap(myarr)
	myarr = Up(myarr)
	myarr = Low(myarr)
	myarr = Ncap(myarr)
	myarr = Nup(myarr)
	myarr = Nlow(myarr)
	myarr = Punct(myarr)
	myarr = Quot(myarr)
	// myarr = Punct(myarr)
	myarr = Vol(myarr)
	return Clean(myarr)
}

func Relode() {
	if len(os.Args) != 3 {
		fmt.Println("wronge inputs")
		return
	}
	content, error := os.ReadFile(os.Args[1])
	fContent := string(content)
	if len(fContent) < 1 {
		fmt.Println("the " + os.Args[1] + "is emty")
		return
	} else {
		fContent = string(content) + " "
	}
	if error != nil {
		fmt.Println("name file does not exist")
		return
	}
	var lastarr [][]string
	lines := ""
	for ind, val := range fContent {
		if val != '\n' {
			lines += string(val)
		} else if val == '\n' {
			myarr := SPlit(lines)
			myarr = proses(myarr)
			lastarr = append(lastarr, myarr)
			lines = ""
		}
		if ind == len(fContent)-1 && lines != "" {
			myarr := SPlit(lines)
			myarr = proses(myarr)

			lastarr = append(lastarr, myarr)
			lines = ""
		}
	}
	text := ""
	for ind, val := range lastarr {
		i := 0
		for i = 0; i < len(val); i++ {
			j := i
			if val[i] == "()" || val[i][0] == '(' || val[i][0] == ')' {
				text += val[i]
			} else if i < len(val)-1 && i > 0 && (val[i+1] == "()" || val[i+1] == "(" || val[j+1] == ")" && val[j-1] == "()" || val[j-1] == "(" || val[j-1] == ")") {
				text += val[i]
			} else {
				text += " " + val[i] + " "
			}
			// this is if we wante not adding the sapce(up)space => sapcespace
			// else {
			// 	text +=  val[i]
			// }
		}

		if ind < len(lastarr)-1 {
			text += "\n"
		}
	}

	lastcopy := normalizeSpaces(text)

	start := 0
	for ind, val := range os.Args[2] {
		if val == '.' {
			start = ind
			break
		}
	}
	if os.Args[2][start:] == ".txt" && len(os.Args[2]) > 5 {
		myfile, err := os.Create(os.Args[2])
		if err != nil {
			fmt.Printf("we can not creat this file")
		}
		_, err = myfile.WriteString(strings.Trim(lastcopy, " "))
		if err != nil {
			fmt.Printf("we can not copyed this line")
		}
	} else {
		fmt.Printf("we can not creat this file")
	}
}
