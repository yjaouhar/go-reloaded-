package op

import (
	"log"
	"strconv"
	"strings"
)

func HandelFlags(s []string) []string {
	for i := 0; i < len(s); i++ {
		if Flag(s[i]) {
			count := 0 // Counter counts how many times a flag has applied on a word
			for k := i - 1; count != 1; k-- {
				if s[i] != "(bin)" && s[i] != "(hex)" {
					if !Empty(s[k]) || !ValidLetter(s[k]) { // skips every non-valid element
						if s[i] == "(cap)" {
							s[k] = Cap(s[k])
						} else if s[i] == "(up)" {
							s[k] = strings.ToUpper(s[k])
						} else if s[i] == "(low)" {
							s[k] = strings.ToLower(s[k])
						}
						count++
					}
				} else {
					if !Empty(s[k]) {
						if s[i] == "(bin)" {
							if Bin(s[k]) {
								binValue, err := strconv.ParseInt(s[k], 2, 64)
								if err != nil {
									log.Fatal("The  number  you are trying to convert is out of range of int64!")
								}
								s[k] = strconv.FormatInt(binValue, 10)
								count++
							} else {
								log.Fatal("The word is not a Bin number  ")
							}
						} else if s[i] == "(hex)" {
							if Hex(s[k]) {
								hexValue, err := strconv.ParseInt(s[k], 16, 64)
								if err != nil {
									log.Fatal("The  number  you are trying to convert is out of range of int64!")
								}
								s[k] = strconv.FormatInt(hexValue, 10)
								count++
							} else {
								log.Fatal("The word is not a Hexadecimal number  ")
							}
						}
					}
				}
			}

			s[i] = "  "

		} else {
			if i != len(s)-1 && s[i] != "" && s[i+1] != "" {
				Valid, number := FlagWithNum(s[i], s[i+1])
				if Valid {
					if s[i] == "(cap," || s[i] == "(low," || s[i] == "(up," {
						if number > i {
							number = i
						}
						count := 0
						if number >= 0 {
							for k := i - 1; count != number; k-- {
								if !Empty(s[k]) || !ValidLetter(s[k]) {
									if s[i] == "(cap," {
										s[k] = Cap(s[k])
									} else if s[i] == "(low," {
										s[k] = strings.ToLower(s[k])
									} else if s[i] == "(up," {
										s[k] = strings.ToUpper(s[k])
									}
									count++
								}
							}
						}
						s[i] = "  "
						s[i+1] = "  "

					}
				}
			}
		}
	}
	return s
}
