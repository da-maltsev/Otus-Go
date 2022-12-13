package main

import (
	"fmt"
	"regexp"
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

func sentenceToSlice(sentence string) []string {
	splitRegexp := regexp.MustCompile(`[\n\t.,!?;: «»()—\"']+`)
	words := splitRegexp.Split(sentence, -1)
	return words
}

func countWords(words []string) map[string]int {
	counter := make(map[string]int)
	for _, word := range words {
		counter[word]++
	}
	return counter
}

func sortedWords(wordsCounted map[string]int) []string {
	pairList := make([]Pair, 0, len(wordsCounted))
	i := 0
	for key, value := range wordsCounted {
		pairList = append(pairList, Pair{key, value})
		i++
	}

	sort.Slice(pairList, func(i, j int) bool {
		return pairList[i].Value > pairList[j].Value
	})
	resultArray := make([]string, 0, 10)
	j := 0
	for _, pair := range pairList {
		if j >= 10 {
			break
		}
		resultArray = append(resultArray, pair.Key)
		j++
	}
	return resultArray
}

func doAllTheWork(s string) []string {
	slice := sentenceToSlice(s)
	mapper := countWords(slice)
	result := sortedWords(mapper)

	return result
}

func main() {
	simpleText := "Собака бывает кусачей Только от жизни собачьей. Только от жизни, от жизни собачьей Собака бывает кусачей."
	result := doAllTheWork(simpleText)
	fmt.Println(result)
}
