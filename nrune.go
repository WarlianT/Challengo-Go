package student

func NRune(s string, n int) rune {
	var aRune []rune = []rune(s)
	if n > 0 && n <= StrLen(s) {
		return aRune[n-1]
	} else {
		return 0
	}
}