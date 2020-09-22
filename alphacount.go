package student

func AlphaCount(s string) int {

	var i int = 0
	for _, alpha := range s {
		if alpha > 64 && 91 > alpha || alpha > 96 && 123 > alpha {
			i++
		}
	}
	return i
}
