package op

import (
	"strings"
)

func Quotes(str string) string {
	Sl_Rune := []rune(str)
	Sl_String := []string{}
	B_Q := ""  // store word befor punctuation mark
	word := "" // stor word that contradict punctuation mark
	Text := "" // text finale for return
	Open := 0  // check the punctuation mark  is open  or close
	for i := 0; i < len(Sl_Rune); i++ {
		// check if punctuation mark  it is not found between the words
		if Sl_Rune[i] != '\'' || (Sl_Rune[i] == '\'' && (i != 0 && (Sl_Rune[i-1] != ' ' && Sl_Rune[i-1] != '\'') && (i != len(Sl_Rune)-1 && (Sl_Rune[i+1] != ' ' && Sl_Rune[i+1] != '\'')))) {
			word += string(Sl_Rune[i])
		} else { // search for punctuation mark is available or not
			if Sl_Rune[i] == '\'' && Open == 0 {
				Open = 1
				B_Q = word
				word = ""
				Sl_String = append(Sl_String, B_Q)
				if i != 0 && Sl_Rune[i-1] != ' ' && Sl_Rune[i-1] != '\'' { // chcke befor punctuation mark is have space
					Sl_String = append(Sl_String, " ")
				}
				Sl_String = append(Sl_String, string(Sl_Rune[i]))

			} else if Sl_Rune[i] == '\'' && Open == 1 { //
				if !Check_Space(word) {
					Sl_String = append(Sl_String, Remove_Space(word))
					Sl_String = append(Sl_String, "'")
					if i != len(Sl_Rune)-1 && Sl_Rune[i+1] != ' ' && Sl_Rune[i+1] != '\'' && !Is_punc(string(Sl_Rune[i+1])) { // chcke after punctuation mark is have space
						Sl_String = append(Sl_String, " ")
					}
					B_Q = ""
					word = ""
					Open = 0
				} else {
					Sl_String = append(Sl_String, word)
					Sl_String = append(Sl_String, "'")
					word = ""
				}
			}
		}
	}
	Sl_String = append(Sl_String, word)

	for _, v := range Sl_String {
		Text += v
	}

	return Text
}

// Check insid of punctuation mark if have word
func Check_Space(word string) bool {
	sl := []rune(word)
	ver := true
	for i := 0; i < len(sl); i++ {
		if sl[i] != ' ' {
			ver = false
			break
		}
	}

	return ver
}

// Remove space after and befor word

func Remove_Space(ss string) string {
	ss = strings.TrimRight(ss, " ")
	ss = strings.TrimLeft(ss, " ")
	return ss
}
