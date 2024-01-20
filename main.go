package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

func CountWords(text string) (int64, int64) {
	// Count the number of Chinese characters
	hanRe := regexp.MustCompile("[\u4e00-\u9fa5]")
	hanRes := hanRe.FindAllString(text, -1)
	newHanText := strings.Join(hanRes, "")
	//fmt.Println(utf8.RuneCountInString(newHanText))

	// Count the number of Chinese characters
	hanCharRe := regexp.MustCompile(`[^\x00-\xff]+`)
	hanCharRes := hanCharRe.FindAllString(text, -1)
	newHanCharText := strings.Join(hanCharRes, "")
	//fmt.Println(utf8.RuneCountInString(newHanCharText))

	// Count the number of English words
	engRe := regexp.MustCompile("[a-zA-Z]+")
	engRes := engRe.FindAllString(text, -1)
	newEngText := strings.Join(engRes, " ")
	newEngTextSlice := strings.Fields(newEngText)
	//fmt.Println(len(newEngTextSlice))

	// Count the number of Chinese characters
	engCharRe := regexp.MustCompile("[a-zA-Z]+|[[:punct:]]")
	engCharRes := engCharRe.FindAllString(text, -1)
	newEngCharText := strings.Join(engCharRes, " ")
	newEngCharTextSlice := strings.Fields(newEngCharText)
	//fmt.Println(len(newEngCharTextSlice))

	// Number of Chinese characters and words
	pureWordsCnt := int64(utf8.RuneCountInString(newHanText)) + int64(len(newEngTextSlice))

	// All character count
	allWordsCnt := int64(utf8.RuneCountInString(newHanCharText)) + int64(len(newEngCharTextSlice))

	return pureWordsCnt, allWordsCnt
}

func main() {
	text := "This is a sample sentence. 这是一个示例。"
	pureWordsCnt, allWordsCnt := CountWords(text)
	fmt.Println("Count of PureWords: ", pureWordsCnt)
	fmt.Println("Count of All Characters: ", allWordsCnt)
}
