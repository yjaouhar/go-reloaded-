package op

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Applic_Flg(s []string) []string {
	for i := 0; i < len(s); i++ {
		if Is_Flg(s[i]) {

			count := 0 // Counter counts how many times a flag has applied on a word
			for k := i - 1; k >= 0; k-- {
				if count == 1 { // if a flag got applied on the first word it stops the loop searching for a valid word
					break
				}
				if s[i] == "(cap)" || s[i] == "(up)" || s[i] == "(low)" {
					if s[k] == "" || s[k] == " " || s[k] == "  " || !Check_words(s[k]) { // skips every non-valid element
						continue
					} else {
						if s[i] == "(cap)" {
							s[k] = Cap(s[k])
							count++
						} else if s[i] == "(up)" {
							s[k] = strings.ToUpper(s[k])
							count++
						} else if s[i] == "(low)" {
							s[k] = strings.ToLower(s[k])
							count++
						}
					}
				} else {
					if s[k] == "" || s[k] == " " || s[k] == "  " {
						continue
					} else {
						if s[i] == "(bin)" {
							if !Check_Bin(s[k]) {
								fmt.Println("The word is not a Binary number  ")
								os.Exit(1)
							} else {
								Bin_value, err := strconv.ParseInt(s[k], 2, 64)
								if err != nil {
									fmt.Println("The  number  you are trying to convert is out of range of int64!")
									os.Exit(1)
								}

								s[k] = strconv.FormatInt(Bin_value, 10)
								count++
							}
						} else if s[i] == "(hex)" {
							if!Check_Hex(s[k]) {
								fmt.Println("The word is not a Hexadecimal number  ")
								os.Exit(1)
							} else {
								Hex_value, err := strconv.ParseInt(s[k], 16, 64)
								if err != nil {
									fmt.Println("The  number  you are trying to convert is out of range of int64!")
									os.Exit(1)
								}
								s[k] = strconv.FormatInt(Hex_value, 10)
								count++
							}
						}
					}
				}

			}

			s[i] = "  "

		} else { // If not Flag normal Check if Flag with number
			if i != len(s)-1 && s[i] != "" && s[i+1] != "" {

				Valid, number := Is_Flg_num(s[i], s[i+1])

				if Valid {
					if s[i] == "(cap," || s[i] == "(low," || s[i] == "(up," {
						
						if number > i {
							number = i
						}
						count := 0
						if number >= 0 {
							for k := i - 1; k >= 0; k-- {

								if count == number {
									break
								}
								if s[k] == "" || s[k] == " " || s[k] == "  " || !Check_words(s[k]) {
									continue
								} else {
									if s[i] == "(cap," {
										s[k] = Cap(s[k])
										count++
									} else if s[i] == "(low," {
										s[k] = strings.ToLower(s[k])
										count++
									} else if s[i] == "(up," {

										s[k] = strings.ToUpper(s[k])
										count++
									}
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
