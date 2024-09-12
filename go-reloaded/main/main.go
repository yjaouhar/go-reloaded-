package main

import (
	"fmt"
	"os"
	"strings"

	op "Reloaded/operations"
)

var (
	Text_Final   string
	Text_Splited []string
	Text         string
	FileName     string
	ResultFile   string
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("The argumant you entered is not correct. Please entre argumant contains: name of program | flie for witing text (.txt) | file for writ text edited(.txt ) ")
		os.Exit(1)
	} else {
		arg := os.Args[1:3]
		ResultFile = arg[1]
		if !strings.HasSuffix(arg[0], ".txt") {
			fmt.Println("The is file (", arg[0], ") is not (.txt) file !!")
			os.Exit(1)
		}
		InputFile, err := os.ReadFile(arg[0])
		if err != nil {
			fmt.Println("ERRORE : File not found!!! Give an existing file name or create a new file")
			os.Exit(1)
		}
		contentStr := string(InputFile)
		content_in_rune := []rune(contentStr)
		Matrtix := op.SplitWhiteSpaces((content_in_rune))
		for l := 0; l < len(Matrtix); l++ { // loop for line of matrix
			Matrtix[l] = op.Applic_Flg(Matrtix[l]) // check line of matrix if have Flag and aplied
			Text = op.Join_with_space(Matrtix[l], " ")
			Text = op.Punc(Text)
			Text = op.Quotes((Text))
			Text = op.Vowel(Text)
			Text_Splited = append(Text_Splited, Text)
			Text_Final = strings.Join(Text_Splited, "\n")
		}
	}
	if !strings.HasSuffix(ResultFile, ".txt") {
		fmt.Println("The file is not (.txt) file !!")
		os.Exit(1)
	}
	Result, err := os.Create(ResultFile)
	if err != nil {
		fmt.Println("ERRORE : couldn't creat  the given  file")
		os.Exit(1)
	}
	_, err = Result.WriteString(Text_Final)
	if err != nil {
		fmt.Println("ERRORE : couldn't write in the given resulte file")
		os.Exit(1)
	}
}
