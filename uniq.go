package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"github.com/elliotchance/orderedmap/v2"
)

type lineData struct {
	count int
	originalLine string
}

type options struct {
	countFlags 		  bool
	repeatLine 	 	  bool
	noRepeatLine 	  bool
	ignoreRegister 	  bool
	ignoreFirstFields int
	ignoreFirstSymbol int
}

func main() {
	par := options{
		countFlags : *flag.Bool("c", false, "number of line encounters"), 
		repeatLine : *flag.Bool("d", false, "number of line encounters"),
		noRepeatLine : *flag.Bool("u", false, "repeatable lines"),
		ignoreRegister : *flag.Bool("i", false, "ignore the register of letters"),
		ignoreFirstFields : *flag.Int("f", 0, "ignoring the first fields"),
		ignoreFirstSymbol : *flag.Int("s", 0, "ignoring the first symbol"),
	}

	flag.Parse()
	args := flag.Args()

	inputReader := bufio.NewReader(os.Stdin)
	outputWriter := io.Writer(os.Stdout)

	if len(args) > 0 {
		fileInput, err := os.Open(args[0])
		if err != nil {
			fmt.Println("No open file!")
			os.Exit(-1)
		}
		defer fileInput.Close()
		inputReader = bufio.NewReader(fileInput)

		if len(args) > 1 {
			fileOutput, err := os.Create(args[1])
			if err != nil {
				fmt.Println("No open or create file!")
				os.Exit(-2)
			}
			defer fileOutput.Close()
			outputWriter = io.Writer(fileOutput)
		}
	}

	mapString := readStrings(inputReader, par)

	outputResult(outputWriter, mapString, par)
}


func outputResult(outputWriter io.Writer, mapString *orderedmap.OrderedMap[string, lineData], par options) {
	for _, key := range mapString.Keys() {
		value, _ := mapString.Get(key)

		switch {
		case par.countFlags:
			if len(value.originalLine) > 0 {
				fmt.Fprint(outputWriter, value.count, " ", value.originalLine)
			}
		case par.repeatLine:
			if value.count > 1 {
				fmt.Fprint(outputWriter, value.originalLine)
			}
		case par.noRepeatLine:
			if value.count == 1 {
				fmt.Fprint(outputWriter, value.originalLine)
			}
		default:
			fmt.Fprint(outputWriter, value.originalLine)
		}
	}
}

func readStrings(reader *bufio.Reader, opt options) *orderedmap.OrderedMap[string, lineData] {
	mapStr := orderedmap.NewOrderedMap[string, lineData]()
	var err error
	for err == nil {
		var str string
		str, err = reader.ReadString('\n')

		if err == nil {
			addToMap(mapStr, str, opt)
		}
	}

	return mapStr
}

func addToMap(mapStr *orderedmap.OrderedMap[string, lineData], str string, opt options) {
	checkStr := updateStrOptions(str, opt)

	if val, ok := mapStr.Get(checkStr); !ok{
		mapStr.Set(checkStr, lineData {
			count: 		  1,
			originalLine: str,
		})
	} else {
		val.count++
		mapStr.Set(checkStr, val)
	}
}

func updateStrOptions(updateStr string, opt options) string {
	if opt.ignoreRegister {
		updateStr = strings.ToLower(updateStr)
	}	

	
	n := opt.ignoreFirstFields 
	if n > 0 {
		words := strings.FieldsFunc(updateStr, unicode.IsSpace)
		if len(words) > 1 && n < len(words) {
			updateStr = strings.Join(words[n : ], " ")
		}
	}

	n = opt.ignoreFirstSymbol
	if n > 0 && len(updateStr) != 0 {
		if n < len(updateStr) {
			updateStr = updateStr[n : ]
		}
	} 

	return updateStr
}
