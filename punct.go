package Relode

import (
	"strings"
)

func Punct(myarr []string) []string {
	copy := ""
	cleanarr := myarr
	if len(cleanarr) >= 1 {

		ind := 0
		for ind = 0; ind < len(cleanarr)-1; ind++ {
			i := 0
			if strings.Contains(cleanarr[ind], ".") || strings.Contains(cleanarr[ind], ",") || strings.Contains(cleanarr[ind], "?") || strings.Contains(cleanarr[ind], ";") || strings.Contains(cleanarr[ind], ":") || strings.Contains(cleanarr[ind], "!") {
				cleanarr[ind] = strings.TrimSpace(cleanarr[ind])
				mytext := []rune(strings.TrimSpace(cleanarr[ind]))

				for i = 0; i < len(mytext)-1; i++ {
					if mytext[i] == '.' || mytext[i] == ',' || mytext[i] == '?' || mytext[i] == ';' || mytext[i] == ':' || mytext[i] == '!' {
						if mytext[i+1] == '.' || mytext[i+1] == ',' || mytext[i+1] == '?' || mytext[i+1] == ';' || mytext[i+1] == ':' || mytext[i+1] == '!' {
							copy += string(mytext[i])
						} else {
							copy += string(mytext[i]) + " "
						}
					} else {
						copy += string(mytext[i])
					}
				}
				count := 1
				for ind+count < len(cleanarr)-1 && cleanarr[ind+count] == " " {

					count++
				}
				if cleanarr[ind+count][0] == '.' || cleanarr[ind+count][0] == ',' || cleanarr[ind+count][0] == '?' || cleanarr[ind+count][0] == ';' || cleanarr[ind+count][0] == ':' || cleanarr[ind+count][0] == '!' {
					copy += string(mytext[i])
				} else {

					copy += string(mytext[i]) + " "
				}

			} else {

				count := 1
				for ind+count < len(cleanarr)-1 && cleanarr[ind+count] == " " {
					count++
				}

				if cleanarr[ind+count][0] == '.' || cleanarr[ind+count][0] == ',' || cleanarr[ind+count][0] == '?' || cleanarr[ind+count][0] == ';' || cleanarr[ind+count][0] == ':' || cleanarr[ind+count][0] == '!' {
					copy += strings.TrimSpace(cleanarr[ind])
				} else {
					copy += cleanarr[ind] 
				}
			}
		}

		if strings.Contains(cleanarr[ind], ".") || strings.Contains(cleanarr[ind], ",") || strings.Contains(cleanarr[ind], "?") || strings.Contains(cleanarr[ind], ";") || strings.Contains(cleanarr[ind], ":") || strings.Contains(cleanarr[ind], "!") {

			i := 0
			mytext := []rune(cleanarr[ind])
			for i = 0; i < len(mytext)-1; i++ {
				if mytext[i] == '.' || mytext[i] == ',' || mytext[i] == '?' || mytext[i] == ';' || mytext[i] == ':' || mytext[i+1] == ':' || mytext[i] == '!' {
					if mytext[i+1] == '.' || mytext[i+1] == ',' || mytext[i+1] == '?' || mytext[i+1] == ';' || mytext[i+1] == ':' || mytext[i+1] == '!' {
						copy += string(mytext[i])
					} else {
						copy += string(mytext[i]) + " "
					}
				} else {
					copy += string(mytext[i])
				}
			}
			copy += string(mytext[i])

		} else {
			if strings.Contains(cleanarr[ind], ".") || strings.Contains(cleanarr[ind], ",") || strings.Contains(cleanarr[ind], "?") || strings.Contains(cleanarr[ind], ";") || strings.Contains(cleanarr[ind], ":") || strings.Contains(cleanarr[ind], "!") {

				cleandestest := strings.Join(strings.Fields(cleanarr[ind]), " ")
				mytext := []rune(cleandestest)
				i := 0

				for i = 0; i < len(mytext)-1; i++ {
					if mytext[i] == '.' || mytext[i] == ',' || mytext[i] == '?' || mytext[i] == ';' || mytext[i] == ':' || mytext[i] == '!' {
						if mytext[i+1] == '.' || mytext[i+1] == ',' || mytext[i+1] == '?' || mytext[i+1] == ';' || mytext[i+1] == ':' || mytext[i+1] == '!' || mytext[i+1] == ' ' {
							copy += string(mytext[i])
						} else {
							copy += string(mytext[i]) + " "
						}
					} else {
						if mytext[i] == ' ' && mytext[i+1] != ' ' {
							copy += string(mytext[i])
						} else if mytext[i] == ' ' {
							continue
						} else {
							copy += string(mytext[i])
						}
					}
				}

				copy += string(mytext[i])

			} else {
				copy += cleanarr[ind]
			}
		}

	}

	return SPlit(copy)
}
