package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	m := make(map[string][]string)

	content, err := ioutil.ReadFile("c:\\dev\\sorbet\\input.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)

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

	start := rand.Intn(len(all))
	sentence := ""
	currentWord := all[start]
	//loopUntil := 0

	for {
		choose := rand.Intn(len(m[currentWord]))
		nextWord := m[currentWord][choose]
		if nextWord == "." || nextWord == "," {
			sentence += nextWord
		} else {
			sentence += " " + nextWord
		}

		//loopUntil++
		if nextWord == "." {
			break
		}
		currentWord = nextWord
	}

	fmt.Println(sentence)

}

// for i := range m["fish"] {
// 	fmt.Println(i, m["fish"][i])
// }
