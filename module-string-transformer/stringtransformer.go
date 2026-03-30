package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var (
	text string
	err  error
)

func toUpper() {
	fmt.Println("===> ToUpper ===>")
	fmt.Print("Enter Text: ")
	text, _ = reader.ReadString('\n')
	text = strings.ToUpper(text)
	text = strings.TrimSpace(text)
	if text == "" {
		fmt.Print("Text Block Can Not Be Empty")

	}
	fmt.Println(text)
}

func toLower() {
	fmt.Println("===> ToLower ===>")
	fmt.Print("Enter Text: ")
	text, _ = reader.ReadString('\n')
	text = strings.ToLower(text)
	fmt.Println(text)

}

func toCap() {
textblock:
	fmt.Println("===> ToCap ===>")
	fmt.Print("Enter Text: ")
	text, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error Reading Text")
	}
	text = strings.TrimSpace(text)
	if text == "" {
		fmt.Println("Text Block Can Not Be Empty")
		goto textblock
	}

	word := strings.Fields(text)
	for i := range word {
		word[i] = strings.ToUpper(string(word[i][:1])) + strings.ToLower(word[i][1:])

	}
	fmt.Println(word)
}
func toTitle() {

	fmt.Println("===> ToTitle ===>")
	fmt.Print("Enter Text: ")
	text, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error Reading Text")
	}
	text = strings.TrimSpace(text)
	if text == "" {
		fmt.Println("Text Block Can Not Be Empty")
	}
	words := map[string]bool{
		"for": true, "nor": true, "on": true, "at": true, "to": true, "by": true, "in": true, "of": true, "up": true, "as": true, "is": true, "it": true}

	if strings.ContainsAny(text, words[0]) {
		text = strings.ToLower(text[:1])
	}

	text = strings.ToUpper(text)
	fmt.Println(text)
}
func toSnake() {
	text = strings.ReplaceAll(text, " ", "_")
	fmt.Println(text)
}

func main() {

	for {
		//menu:
		fmt.Println("====> SENTINEL REPORT PANEL =====> ")
		fmt.Println("> 1. ToUpper")
		fmt.Println("> 2. ToLower")
		fmt.Println("> 3. ToCap")
		fmt.Println("> 4. ToTitle")
		fmt.Println("> 5. ToSnake")
		fmt.Println("> 6. Exit")
		fmt.Print("Select Action: ")
		actionSelector, _ := reader.ReadString('\n')
		actionSelector = strings.TrimSpace(actionSelector)
		actionSelectorInput, actioerr := strconv.Atoi(actionSelector)

		if actioerr != nil {
			fmt.Println("Pick An Action To Format Your Text!!")
			continue
		}

		switch actionSelectorInput {
		case 1:
			toUpper()
		case 2:
			toLower()
		case 3:
			toCap()
		case 4:
			toTitle()
		}

	}

}
