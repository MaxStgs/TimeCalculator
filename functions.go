package main

import "fmt"

func generateURI() (s string) {
	s = "http://" + site + ":" + port
	return
}

func generateIndex() (s string) {
	s = generateURI() + "/index"
	return
}

func getText(s string) (translate string) {

	list, ok := listTranslations[s]
	if !ok {
		errorHandler("functions.getText() not found Translate for "+s, 1)
		return s
	}
	translate = list[lang]
	return
}

// errorType: 1 - to console, 2 - to log file
func errorHandler(m string, errorType Error) {
	switch errorType {
	case console:
		fmt.Println(m)
	case file:
		fmt.Println("ToFile: ", m)
	case unknownError:
	default:
		fmt.Println("Unhandled errorHandler: ", m)
	}
}
