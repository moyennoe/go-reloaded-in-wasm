package piscine

//Atoi transofrms gives an integer from a number written as a string
func Atoi(s string) int {
	result := 0
	chaine := s
	signe := 1

	if s == "" {
		return 0
	}
	if s[0] == '-' {

		chaine = s[1:]
		signe = -1

	} else if s[0] == '+' {

		chaine = s[1:]

	}
	for _, cha := range chaine {

		if cha >= '0' && cha <= '9' {

			chiffre := int(cha) - 48
			result = result*10 + chiffre

		} else {

			return 0
		}
	}

	return signe * result
}
