package arrays

func reverse(str string) string {
	if str == "" {
		return ""
	}

	runeArray := []rune(str)

	for i := 0; i < len(runeArray)/2; i++ {
		j := len(runeArray) - 1 - i
		runeArray[i], runeArray[j] = runeArray[j], runeArray[i]
	}
	return string(runeArray)
}
