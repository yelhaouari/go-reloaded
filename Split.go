package Relode

func cheak(str string, ind int) bool {
	for i := ind; i < len(str); i++ {
		if str[i] == '(' {
			return true
		}
		if str[i] == ')' {
			return false
		}
	}
	return false
}

func SPlit(str string) []string {
	textcopy := ""
	var myarr []string
	nextext := []rune(str)
	i := 0
	flage := false
	for i = 0; i < len(nextext)-1; i++ {

		if nextext[i] == '(' {
			myarr = append(myarr, textcopy)
			textcopy = ""
			flage = true
		}

		if nextext[i] == ')' {
			textcopy += ")"
			myarr = append(myarr, textcopy)
			textcopy = ""
			flage = false
			continue

		}

		if flage {
			if cheak(str, i) && nextext[i] == ' ' {
				if textcopy != "" {
					myarr = append(myarr, textcopy)
					textcopy = ""
				}
			}
		}

		if flage {
			textcopy += string(nextext[i])
			continue
		}
		if nextext[i] != '\n' && nextext[i] != ' ' {
			textcopy += string(nextext[i])
		}
		if nextext[i] == '\n' {
			textcopy += "\n"
		}

		if nextext[i] == ' ' && !flage {
			if textcopy != "" {
				myarr = append(myarr, textcopy)
				textcopy = ""
			}
		}
	}
	if i < len(nextext) {

		textcopy += string(nextext[i])
		if textcopy != "" {
			myarr = append(myarr, textcopy)
			textcopy = ""

		}
	}

	return myarr
}
