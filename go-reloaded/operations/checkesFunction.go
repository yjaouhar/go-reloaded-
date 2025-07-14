package op

import (
	"strconv"
)

func ValidLetter(s string) bool { // Check
	for _, v := range []rune(s) {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || !Ponctuation(string(v)) {
			return true
		}
	}
	return false
}

func Flag(s string) bool {
	return s == "(cap)" || s == "(up)" || s == "(low)" || s == "(hex)" || s == "(bin)"
}

func FlagWithNum(flg, nb string) (bool, int) {
	if flg[0] == '(' && (flg[1:] == "cap," || flg[1:] == "low," || flg[1:] == "up,") {
		if nb[len(nb)-1] == ')' {
			num, err := strconv.Atoi(nb[0 : len(nb)-1])
			if err != nil {
				return false, 0
			}
			if num < 0 {
				return true, num
			} else {
				return true, num
			}
		}
	}

	return false, 0
}

func Bin(s string) bool {
	sl := []rune(s)
	for i := 0; i < len(sl); i++ {
		if !(sl[i] == '1' || sl[i] == '0') {
			return false
		}
	}
	return true
}

func Hex(s string) bool {
	sl := []rune(s)
	for i := 0; i < len(sl); i++ {
		if !((sl[i] >= 'a' && sl[i] <= 'f') || (sl[i] >= 'A' && sl[i] <= 'F') || (sl[i] >= '0' && sl[i] <= '9') || (i == 0 && (sl[i] == '-' || sl[i] == '+'))) {
			return false
		}
	}
	return true
}

func Empty(s string) bool {
	return s == "" || s == " " || s == "  "
}
