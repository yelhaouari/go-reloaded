package Relode

import (
	"strings"
)

func Punct(myarr []string) []string {
	copy := ""

	var cleanarr []string

	if len(Clean(myarr)) == 1 {
		if strings.Contains(Clean(myarr)[0], ".") || strings.Contains(Clean(myarr)[0], ",") || strings.Contains(Clean(myarr)[0], "?") || strings.Contains(Clean(myarr)[0], ";") || strings.Contains(Clean(myarr)[0], ":") || strings.Contains(Clean(myarr)[0], "!") {
			if len(myarr[0]) > 1 {
				cleanarr = append(cleanarr, string(Clean(myarr)[0][0:]))
				cleanarr = Clean(cleanarr)
			} else {
				return Clean(myarr)
			}
		} else {
			if len(Clean(myarr)) > 1 {
				cleanarr = strings.Split(myarr[0], " ")
			} else {
				return myarr
			}
		}
	} else {
		cleanarr = Clean(myarr)
	}

	if len(cleanarr) >= 1 {
		ind := 0
		for ind = 0; ind < len(cleanarr)-1; ind++ {
			i := 0
			cleanarr[ind] = strings.TrimSpace(cleanarr[ind])
			if strings.Contains(cleanarr[ind], ".") || strings.Contains(cleanarr[ind], ",") || strings.Contains(cleanarr[ind], "?") || strings.Contains(cleanarr[ind], ";") || strings.Contains(cleanarr[ind], ":") || strings.Contains(cleanarr[ind], "!") {

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
				if strings.TrimSpace(cleanarr[ind+1])[0] == '.' || strings.TrimSpace(cleanarr[ind+1])[0] == ',' || strings.TrimSpace(cleanarr[ind+1])[0] == '?' || strings.TrimSpace(cleanarr[ind+1])[0] == ';' || strings.TrimSpace(cleanarr[ind+1])[0] == ':' || strings.TrimSpace(cleanarr[ind+1])[0] == '!' {
					copy += string(mytext[i])
				} else {
					copy += string(mytext[i]) + " "
				}

			} else {
				if strings.TrimSpace(cleanarr[ind+1])[0] == '.' || strings.TrimSpace(cleanarr[ind+1])[0] == ',' || strings.TrimSpace(cleanarr[ind+1])[0] == '?' || strings.TrimSpace(cleanarr[ind+1])[0] == ';' || strings.TrimSpace(cleanarr[ind+1])[0] == ':' || strings.TrimSpace(cleanarr[ind+1])[0] == '!' {
					copy += cleanarr[ind]
				} else {
					copy += " " + cleanarr[ind] + " "
				}
			}

		}
		if cleanarr[ind][0] == '.' || cleanarr[ind][0] == ',' || cleanarr[ind][0] == '?' || cleanarr[ind][0] == ';' || cleanarr[ind][0] == ':' || cleanarr[ind][0] == '!' {

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

	return strings.Split(copy, " ")
}
