package parser

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Setting struct {
	Line    uint   //the line where the comment is (add 1 to get the variable's location)
	Name    string //the setting name to display
	Type    string //the type of input for the pannel to handle
	Comment string //comment for the setting
	Value   string //the value on the next line
}

//This function will parse a settings file from go-evil, to be used with the panel
func Parse(path string) []Setting {
	var ret = make([]Setting, 0, 50)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var i uint = 0

	for scanner.Scan() {
		inp := scanner.Text()
		if strings.HasPrefix(inp, "///") {
			i += 1
			continue
		}
		if strings.HasPrefix(inp, "//") {
			splitted := strings.Split(inp, " ")

			scanner.Scan()
			inp = scanner.Text()
			value := strings.Split(inp, "= ")

			ret = append(ret, Setting{
				Line:    i,
				Name:    splitted[1],
				Type:    splitted[2],
				Comment: strSliceToStr(splitted[3:], " "),
				Value:   strSliceToStr(value[1:], " "),
			})

			i += 1
		}
		i += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ret
}

//TODO: Make a function that translates a Setting to the apropriate html render

func strSliceToStr(slice []string, sep string) string {
	var ret strings.Builder
	sliceLen := len(slice)
	for i := 0; i < sliceLen; i++ {
		ret.WriteString(slice[i])

		if i != sliceLen-1 {
			ret.WriteString(sep)
		}
	}
	return ret.String()
}
