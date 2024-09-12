package op

import (
	"fmt"
	"strconv"
)

func Check_words(s string) bool { // Check
	word := false
	slic := []rune(s)
	for i := 0; i < len(slic); i++ {
		if Check_Word(string(slic[i])) {
			word = true
			break

		}
	}
	return word
}

func Check_Word(s string) bool {
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || !Is_punc(string(v)) {
			return true
		}
	}
	return false
}

func Is_Flg(s string) bool {
	if len(s) == 5 || len(s) == 4 {
		if s[0] == '(' && s[len(s)-1] == ')' {
			if s[1:len(s)-1] == "cap" || s[1:len(s)-1] == "up" || s[1:len(s)-1] == "low" || s[1:len(s)-1] == "hex" || s[1:len(s)-1] == "bin" {
				return true
			}
		}
	}
	return false
}

func Is_Flg_num(flg, nb string) (bool, int) {
	if flg[0] == '(' && (flg[1:] == "cap," || flg[1:] == "low," || flg[1:] == "up,") {
		if nb[len(nb)-1] == ')' {
			num, err := strconv.Atoi(nb[0 : len(nb)-1])
			if err != nil {
				fmt.Println("The are is not number")
				return false, 0
			}
			if num < 0 {
				fmt.Println("The number is negative")
				return true, num
			} else {
				return true, num
			}

		}
	}

	return false, 0
}

func Check_Bin(s string) bool {
	sl := []rune(s)

	for i := 0; i < len(sl); i++ {
		if !(sl[i] == '1' || sl[i] == '0') {
			return false
		}
	}
	return true
}

func Check_Hex(s string) bool {
	sl := []rune(s)
	for i := 0; i < len(sl); i++ {
		if !((sl[i] >= 'a' && sl[i] <= 'f') || (sl[i] >= 'A' && sl[i] <= 'F') || (sl[i] >= '0' && sl[i] <= '9') || (i == 0 && (sl[i] == '-' || sl[i] == '+'))) {
			return false
		}
	}
	return true
}
