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

	for key, element := range m {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	rand.Seed(time.Now().UnixNano())

	currentWord := ""

	// looking for start
	for _, sent := range doc.Tokens() {
		if sent.Text == "I" || sent.Tag == "NNP" {
			currentWord = sent.Text
		}
		// I can see Mt. Fuji from here.
		// St. Michael's Church is on 5th st. near the light.
	}

	sentence := currentWord

	loopUntil := 0

	for {
		log.Print(">>>>" + currentWord)
		l := len(m[currentWord])
		if l == 0 {
			break
		}
		choose := rand.Intn(l)
		nextWord := m[currentWord][choose]
		sentence += " " + nextWord

		loopUntil++
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

// for i := range m["fish"] {
// 	fmt.Println(i, m["fish"][i])
// }
