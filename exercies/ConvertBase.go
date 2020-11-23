package piscine

func RecursivePowerr(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		return nb * RecursivePowerr(nb, power-1)
	}
}
func ConvertBaseRec(n int, runes []rune, len int) string {
	res := ""
	for n/len > 0 {
		mod := n % len
		res += string(runes[mod])
		n /= len
	}
	if n > 0 {
		mod := n % len
		res += string(runes[mod])
	}
	return res
}

func StrLen(str string) int {
	runes := []rune(str)
	var count int = 0
	for i := range runes {
		count++
		i = i
	}
	return count
}
func StrRev(s string) string {
	bytes := []byte(s)
	var len int = 0
	var tempByte byte
	for v := range bytes {
		len++
		v = v
	}
	for i := 0; i < len/2; i++ {
		tempByte = bytes[i]
		bytes[i] = bytes[len-i-1]
		bytes[len-i-1] = tempByte
	}
	return string(bytes)
}
func ConvertBase(nbr, baseFrom, baseTo string) string {
	return StrRev(ConvertBaseRec(AtoiBase(nbr, baseFrom), []rune(baseTo), StrLen(baseTo)))
}
