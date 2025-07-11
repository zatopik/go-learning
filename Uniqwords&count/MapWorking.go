package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "input.txt"
	content, err := os.ReadFile(filename)

	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
		return
	}
	fmt.Println(strings.Join(getUniqueWords(content), ""))
	printWordCounts(content)
}
func getUniqueWords(word []byte) []string {
	content := string(word)
	var text = strings.Fields(strings.ToLower(content))
	mape := make(map[string]struct{})
	for _, text := range text {
		cleaned := strings.TrimFunc(text, func(r rune) bool {
			return !((r >= 'a' && r <= 'z') || (r >= 'а' && r <= 'я') || r == '\'')
		})
		if cleaned != "" {
			mape[text] = struct{}{}
		}
	}
	result := make([]string, 0, len(mape))
	for word := range mape {
		result = append(result, word)
	}
	return result
}
func printWordCounts(wordses []byte) {
	result := showshowmuchthatwordbeingusedintext(wordses)
	for text, count := range result {
		fmt.Printf("%s: %d\n", text, count)
	}
}
func showshowmuchthatwordbeingusedintext(words []byte) map[string]int {
	content := string(words)
	var text = strings.Fields(strings.ToLower(content))
	mape := make(map[string]int)
	for _, text := range text {
		cleaned := strings.TrimFunc(text, func(r rune) bool {
			return !((r >= 'a' && r <= 'z') || (r >= 'а' && r <= 'я') || r == '\'')
		})
		if cleaned != "" {
			mape[text]++
		}
	}
	return mape
}
