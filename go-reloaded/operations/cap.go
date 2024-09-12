package op

import (
	"strings"
	"unicode"
)

func Cap(s string) string {
	s = strings.ToLower(s)
	sl := []rune(s)
	for i := 0; i < len(sl); i++ {
		if unicode.IsLetter(sl[i]) {
			sl[i] = unicode.ToUpper(sl[i])
			break
		} else {
			continue
		}
	}

	return string(sl)
}
