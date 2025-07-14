package op

func JoinWithSpace(s []string, spc string) string {
	rex := ""
	for i, v := range s {
		if i != len(s)-1 && s[len(s)-1] == "  " {
			s = s[:len(s)-1]
		}
		if v == "  " {
			continue
		}
		if i < len(s)-1 {
			rex += string(v) + spc
		} else {
			rex += string(v)
		}

	}
	return rex
}
