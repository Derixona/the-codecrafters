package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	var (
		text string
		err  error
	)
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
			fmt.Println("Pick A Text Format!!")
			continue
		}

		switch actionSelectorInput {
		case 1:
			fmt.Println("===> ToUpper ===>")
			fmt.Print("Enter Text: ")
			text, _ = reader.ReadString('\n')
			text = strings.ToUpper(text)
			// if text == "" {
			// 	fmt.Println("Text Block Can Not Be Empty")
			// }
			fmt.Println(text)
		case 2:
			fmt.Println("===> ToLower ===>")
			fmt.Print("Enter Text: ")
			text, _ = reader.ReadString('\n')
			text = strings.ToLower(text)
			fmt.Println(text)
		case 3:

			fmt.Println("===> ToCap ===>")
			fmt.Print("Enter Text: ")
			text, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error Reading Text")
			}
			text = strings.TrimSpace(text)
			if text == "" {
				fmt.Println("Text Block Can Not Be Empty")
				continue
			}

			word := strings.Fields(text)
			for i := range word {
				word[i] = strings.ToUpper(string(word[i][0])) + strings.ToLower(word[i][1:])
				fmt.Println(word)
				break
			}

		}

	}

}
