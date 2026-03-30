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
	//var text string

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
		actionSelectorInput, err := strconv.Atoi(actionSelector)
		if err != nil {
			fmt.Println("Pick A Text Format!!")
			continue
		}

		switch actionSelectorInput {
		case 1:
			fmt.Println("===> ToUpper ===>")
			fmt.Print("Enter Text: ")
			text, _ := reader.ReadString('\n')
			text = strings.ToUpper(text)
			// if text == "" {
			// 	fmt.Println("Text Block Can Not Be Empty")
			// }
			fmt.Println(text)
		case 2:
			fmt.Println("===> ToLower ===>")
			fmt.Print("Enter Text: ")
			text, _ := reader.ReadString('\n')
			text = strings.ToLower(text)
			fmt.Println(text)
		case 3:
			var captext string
			fmt.Println("===> ToCap ===>")
			fmt.Print("Enter Text: ")
			captext, _ = reader.ReadString('\n')
			word := strings.Split(captext, " ")
			for i := range word {
				word[i] = strings.ToUpper(word[i][:1]) + strings.ToLower(word[i][1:])
				words := strings.Join(word, " ")
				fmt.Println(words)
				break
			}

		}

	}

}
