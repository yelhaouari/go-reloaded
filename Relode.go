package Relode

import (
	"fmt"
	"os"
	"strings"
)

func normalizeSpaces(str string) string {
	fmt.Println(str)
	s := []rune(str)
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
	fmt.Println(myarr[0] + "<")
	myarr = Up(myarr)
	myarr = Low(myarr)
	myarr = Cap(myarr)
	myarr = Nup(myarr)
	myarr = Nlow(myarr)
	myarr = Ncap(myarr)
	myarr = Clean(myarr)
	myarr = Punct(myarr)
	myarr = Quot(myarr)
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
			lines = strings.TrimSpace(lines)
			myarr := SPlit(lines)
			myarr = proses(myarr)
			lastarr = append(lastarr, myarr)
			lines = ""
		}
		if ind == len(fContent)-1 && lines != "" {
			lines = strings.TrimSpace(lines)
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
			text += val[i]
		}
		if ind < len(lastarr)-1 {
			text = strings.TrimSpace(text)
			text += "\n"

		}
	}

	lastcopy := normalizeSpaces(text)

	if strings.HasSuffix(os.Args[2], ".txt") && len(os.Args[2]) > 5 {
		myfile, err := os.Create(os.Args[2])
		if err != nil {
			fmt.Printf("we can not creat this file")
		}
		_, err = myfile.WriteString(strings.TrimSpace(lastcopy))
		if err != nil {
			fmt.Printf("we can not copyed this line")
		}
	} else {
		fmt.Printf("we can not creat this file")
	}
}
