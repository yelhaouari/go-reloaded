package Relode

func Vol(myarr []string) []string {
	cleanaar := Clean(myarr)

	for i := 0; i < len(cleanaar)-1; i++ {

		if cleanaar[i] == "A" {
			if cleanaar[i+1][0] == 'a' ||
				cleanaar[i+1][0] == 'e' ||
				cleanaar[i+1][0] == 'i' ||
				cleanaar[i+1][0] == 'o' ||
				cleanaar[i+1][0] == 'u' ||
				cleanaar[i+1][0] == 'h' {
				cleanaar[i] = "An"
			}
		}
		if cleanaar[i] == "a" {
			if cleanaar[i+1][0] == 'a' ||
				cleanaar[i+1][0] == 'e' ||
				cleanaar[i+1][0] == 'i' ||
				cleanaar[i+1][0] == 'o' ||
				cleanaar[i+1][0] == 'u' ||
				cleanaar[i+1][0] == 'h' {
				cleanaar[i] = "an"
			}
		}
	}
	return cleanaar
}
