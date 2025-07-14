package op

import "strings"

func SplitWhiteSpaces(sl []rune) [][]string {
	var (
		Matrtix      [][]string
		line         string
		LineSplitted []string
	)

	for i, v := range sl {
		if v == '\n' {
			if line == "" {
				Matrtix = append(Matrtix, []string{})
			} else {
				LineSplitted = strings.Split(line, " ")
				Matrtix = append(Matrtix, LineSplitted)
				line = ""
			}
		} else {
			line += string(v)
		}
		if i == len(sl)-1 && line != "" {
			LineSplitted = strings.Split(line, " ")
			Matrtix = append(Matrtix, LineSplitted)
		}
	}

	return Matrtix
}
