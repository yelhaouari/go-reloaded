package Relode

import (
	"fmt"
	"os"
	"strings"
)

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
	myarr = Punct(myarr)
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

	lastcopy := ""
	for ind, val := range lastarr {

		for _, sval := range val {
			if sval == "()" || sval == "(" || sval == ")" {
				lastcopy += sval
			} else {
				lastcopy += sval + " "
			}
		}
		if ind < len(lastarr)-1 {
			lastcopy += "\n"
		}
	}

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
