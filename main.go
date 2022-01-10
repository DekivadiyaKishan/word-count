package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
)

func main() {
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	rawData, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	wordsWithCount := countWords(fetchWords(string(rawData)))
	pairList := rankByWordCount(wordsWithCount)
	for i := 0; i <= 9; i++ {
		fmt.Printf("%v %v\n", pairList[i].Key, pairList[i].Value)
	}

}

func fetchWords(text string) []string {
	words := regexp.MustCompile("\\w+")
	return words.FindAllString(text, -1)
}

func countWords(words []string) map[string]int {
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}
	return wordCounts
}

func rankByWordCount(wordFrequencies map[string]int) WordList {
	pl := make(WordList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = WordCount{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type WordCount struct {
	Key   string
	Value int
}

type WordList []WordCount

func (p WordList) Len() int           { return len(p) }
func (p WordList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p WordList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
