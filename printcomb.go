package challenge_go

import "github.com/01-edu/z01"

func PrintComb() {
	min := '0'
	max := '9'
	for i := min; i <= max; i++ {
		for j := min; j <= max; j++ {
			for k := min; k <= max; k++ {
				if i < j && j < k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					if i < '7' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
			}
		}
	}
	z01.PrintRune('\n')
}
