package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/jdkato/prose/v2"
)

func main() {
	m := make(map[string][]string)

	content, err := ioutil.ReadFile("c:\\Users\\Fra\\dev\\sorbet\\input.txt")

	if err != nil {
		log.Fatal(err)
	}

	text := string(content)

	doc, err := prose.NewDocument(text)

	for _, tok := range doc.Tokens() {
		fmt.Println(tok.Text, tok.Tag, tok.Label)
		// Go NNP B-GPE
		// is VBZ O
		// an DT O
		// ...
	}

	if err != nil {
		log.Fatal(err)
	}

	// second

	all := strings.Split(text, " ")

	for j := 0; j < len(all); j++ {
		if j == len(all)-1 {
			continue
		}
		if val, ok := m[all[j]]; ok {
			m[all[j]] = append(val, all[j+1])
		} else {
			m[all[j]] = []string{all[j+1]}
		}
	}

	fmt.Println("Frequencies")
	for key, element := range m {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	rand.Seed(time.Now().UnixNano())

	currentWord := ""

	// looking for start
	starters := []string{}

	for _, sent := range doc.Tokens() {
		if sent.Tag == "NNP" || sent.Text == "I" {
			starters = append(starters, sent.Text)
		}
	}

	currentWord = starters[rand.Intn(len(starters))]
	sentence := currentWord

	for {
		l := len(m[currentWord])
		if l == 0 {
			break
		}
		choose := rand.Intn(l)
		nextWord := m[currentWord][choose]
		sentence += " " + nextWord

		if isEnd(nextWord) {
			break
		}
		currentWord = nextWord
	}

	fmt.Println(sentence)
}

func isEnd(word string) bool {
	log.Print(strings.Index(word, "."))
	return strings.Index(word, ".") == len(word)
}
