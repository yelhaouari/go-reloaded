package Relode

func Clean(myarr []string) []string {
	var cleanarr []string
	for _, val := range myarr {
		if val != "" && val != " " {
			cleanarr = append(cleanarr, val)
		}
	}
	return cleanarr
}
