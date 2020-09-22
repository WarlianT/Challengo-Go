
package student

func LastRune(s string) rune {
	var aRune []rune = []rune(s)
	return aRune[StrLen(s)-1]
}
