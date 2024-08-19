package helloDeps10

import "fmt"

func PrintPhrase(phrase string) string {
	fmt.Println(phrase, "called helloDeps10.PrintPhrase")
	return phrase
}
