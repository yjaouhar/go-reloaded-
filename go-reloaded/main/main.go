package main

import (
	"log"
	"os"
	"strings"

	op "Reloaded/operations"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("The argumant you entered is not correct. Please entre argumant contains: name of program | flie for witing text (.txt) | file for writ text edited(.txt ) ")
	}
	arg := os.Args[1:3]
	resultFileNam := arg[1]
	if !strings.HasSuffix(arg[0], ".txt") {
		log.Fatal("The is file (", arg[0], ") is not (.txt) file !!")
	}
	inputFile, err := os.ReadFile(arg[0])
	if err != nil {
		log.Fatal("ERRORE : File not found!!! Give an existing file name or create a new file")
	}
	op.Handeltext([]rune(string(inputFile)), resultFileNam)
}
