package main

import (
	"fmt"
	"strings"

	// "home/student/piscine"
	"aigerim/functions"
	"os"
)

func checkValidityOfFiles(args []string) bool {
	if len(args) != 2 {
		fmt.Println("ERROR: the number of arguments is incorrect: SHOULD BE 2")
		return false
	} else {
		if args[0] == args[1] {
			fmt.Println("ERROR: the Files must be named differently")
			return false
		}

		if (!strings.HasSuffix(args[0], ".txt")) || (!strings.HasSuffix(args[1], ".txt")) {
			fmt.Println("ERROR: the files should have .txt extension")
			return false
		}
		return true
	}
}

func main() {
	files := os.Args[1:]
	if checkValidityOfFiles(files) {
		// reading from file
		text, err := os.ReadFile(files[0])
		if err != nil {
			fmt.Printf("ERROR: the file %v coudn't be read: %v", files[0], err.Error())
		}

		//===========================================================
		array := functions.SplitByDelimiters(string(text))
		array = functions.RemoveSpaces(array)

		array = functions.MergeBracket(array)
		array = functions.ChangeArticle(array)

		array = functions.ApplyAllCommands(array)

		array = functions.ReplaceUnknownRunes(array)

		array = functions.ApplyAllCommands(array)
		array = functions.AddAllSigns(array)
		array = functions.HandleRestOfBrackets(array)

		array, balancedQuotes := functions.MergeQuotes(array)

		if balancedQuotes {
			fmt.Println("ERROR: Quotes are not balanced. If you have commands inside the quotes will not be applied")
		}

		array = functions.RemoveSpaces(array)

		// // PRINTING
		// for i := 0; i < len(array); i++ {
		// 	if array[i] != "" {
		// 		fmt.Printf("%v. $%v$ --- %v\n", i, array[i], string(array[i][0]))
		// 	} else if array[i] != " " {
		// 		fmt.Printf("%v. $%v$ +++ %v\n", i, array[i], string(array[i]))
		// 	} else if array[i] == "" {
		// 		fmt.Printf("%v. $%v$ === %v\n", i, array[i], string(array[i]))
		// 	}
		// }

		// create the string for writing
		result_text := ""
		for ind, w := range array {
			if (ind < len(array)-1 && array[ind+1] == "\n") || array[ind] == "\n" || ind == len(array)-1 {
				result_text += w
			} else {
				result_text += w + " "
			}
		}

		err1 := os.WriteFile(files[1], []byte(result_text), 0644)

		if err1 != nil {
			fmt.Printf("ERROR: %v", err.Error())
		}

	}
}
