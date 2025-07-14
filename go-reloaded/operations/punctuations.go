package op

func Punc(sl string) string {
	s := []rune(sl)
	slc := []rune{}

	for i := 0; i < len(s); i++ {
		if Ponctuation(string(s[i])) {

			for _, v := range slc {
				if v != ' ' {
					for len(slc) > 0 && slc[len(slc)-1] == ' ' {
						slc = slc[:len(slc)-1]
					}
				} else {
					continue
				}
			}

			slc = append(slc, s[i])
			if i != len(s)-1 && s[i+1] == ' ' {
				for k := i + 1; k < len(s)-1; k++ {
					if s[k] == ' ' {
						continue
					} else {
						i = k
						i--
						break
					}
				}
			}
			if i < len(s)-1 && !Ponctuation(string(s[i+1])) && (s[i+1]) != ' ' {
				slc = append(slc, ' ')
			}

		} else {
			slc = append(slc, s[i])
		}
	}
	return (string(slc))
}

func Ponctuation(p string) bool {
	return p == "." || p == "," || p == "!" || p == "?" || p == ":" || p == ";"
}
