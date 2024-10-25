package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type Pair struct {
	word string
	freq int
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].freq > p[j].freq
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	// reading from book.txt file
	f, err := os.Open("book.txt")
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	// defer close when main method is exited
	defer f.Close()

	// counting words and adding it to map to count its occurence
	words, err := freq(f)
	if err != nil {
		log.Fatalf("err from freq in main: %s", err)
	}

	// sort the word frequencies
	pairs := sortWords(words)

	var count int
	for _, word := range pairs {
		if word.freq > 1 {
			fmt.Printf("%v \t %v\n", word.freq, word.word)
		} else {
			count++
		}
	}
	fmt.Printf("There are %v words with count=1\n", count)

	k, c, err := wordWithMaxFreq(pairs)
	if err != nil {
		log.Fatalf("err in finding word with max freq in main: %s", err)
	}
	fmt.Printf("`%v` word with count=%v\n", k, c)

}

func wordWithMaxFreq(pl PairList) (string, int, error) {
	return pl[0].word, pl[0].freq, nil
}

// this is 1st way to sort by create Len(), Swap(), Less() methods for PairList which is nothing but a slice type
func sortWords(words map[string]int) PairList {
	ps := make(PairList, 0, len(words))
	for k, v := range words {
		ps = append(ps, Pair{k, v})
	}
	sort.Sort(ps)
	return ps
}

// this is 2nd way to sort
func sortWords1(words map[string]int) []Pair {
	ps := make([]Pair, 0, len(words))
	for k, v := range words {
		ps = append(ps, Pair{k, v})
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].freq > ps[j].freq
	})
	return ps
}

func freq(f io.Reader) (map[string]int, error) {
	wordFreq := make(map[string]int)
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)

	for s.Scan() {
		word := strings.ToLower(s.Text())
		wordFreq[word]++
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return wordFreq, nil
}
