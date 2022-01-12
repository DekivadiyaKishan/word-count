package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"sort"
)

func main() {

	http.ListenAndServe(":8080", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		var req RequestBody
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			rw.WriteHeader(500)
			rw.Header().Add("Content-type", "application/json")
			rw.Write([]byte(`{"error": "something is wrong try again"}`))
		}

		wordsWithCount := countWords(fetchWords(req.Input))
		if len(wordsWithCount) == 0 {
			rw.WriteHeader(200)
			rw.Header().Add("Content-type", "application/json")
			rw.Write([]byte(`[]`))
			return
		} else {
			sortedList := rankByWordCount(wordsWithCount)
			sortedList = sortedList[:10]
			byteData, err := json.Marshal(sortedList)
			if err != nil {
				rw.WriteHeader(200)
				rw.Header().Add("Content-type", "application/json")
				rw.Write([]byte(`{"error": "something is wrong try again"}`))
			}
			rw.WriteHeader(200)
			rw.Header().Add("Content-type", "application/json")
			rw.Write(byteData)
			return
		}

	}))
}

type RequestBody struct {
	Input string `json:"input"`
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
