package main

import "fmt"

func getResult(letter string, document string) []byte {
	var result []byte
	firstElem := false
	for index := range document {
		if !firstElem {
			if document[index] != '0' && document[index] != letter[0] {
				result = append(result, document[index])
				firstElem = true
			}
		} else if document[index] != letter[0] {
			result = append(result, document[index])
		}
	}

	if !firstElem {
		return []byte("0")
	} else {
		return result
	}
}

func main() {
	var letter, document string

	for {
		fmt.Scanf("%s %s", &letter, &document)
		if letter == "0" && document == "0" {
			break
		}
		fmt.Println(string(getResult(letter, document)))
	}
}
