package op

import (
	"log"
	"os"
	"strings"
)

func Handeltext(text []rune, resultFileNam string) {
	treatLines := []string{}
	resultFinel := ""
	matrtix := SplitWhiteSpaces(text)
	for _, line := range matrtix {
		line = HandelFlags(line)
		text := JoinWithSpace(line, " ")
		text = Punc(text)
		text = Quotes((text))
		text = Vowel(text)
		treatLines = append(treatLines, text)
		resultFinel = strings.Join(treatLines, "\n")

	}

	if !strings.HasSuffix(resultFileNam, ".txt") {
		log.Fatal("The file is not (.txt) file !!")
	}
	Result, err := os.Create(resultFileNam)
	if err != nil {
		log.Fatal("ERRORE : couldn't creat  the given  file")
	}
	_, err = Result.WriteString(resultFinel)
	if err != nil {
		log.Fatal("ERRORE : couldn't write in the given resulte file")
	}
}
