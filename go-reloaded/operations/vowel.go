package op

func Is_vowel(s rune) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'h' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' || s == 'H' {
		return true
	}
	return false
}

func Vowel(s string) string {
	sl := []rune(s)
	n := len(sl)
	for i := 0; i < n; i++ {
		if (sl[i] == 'a' || sl[i] == 'A') && ((i != 0 && ((sl[i-1] == ' ') || sl[i-1] == '\'') || (i == 0)) || ((i != n-1 && sl[i+1] == ' ') && (i != 0 && !Check_words(string(sl[i-1]))))) {

			j := i + 1
			for j != n && sl[j] == ' ' {
				j++
			}
			if j < n && Is_vowel(sl[j]) && !Is_vowel(sl[j-1]) {
				sl = append(sl[:i+1], sl[i:]...)

				sl[i+1] = 'n'
				n++
			}

		}
	}
	return (string(sl))
}
