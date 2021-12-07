package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	printAnagram(groupAnagramWord("kita", "atik", "tika", "aku", "kia", "makan", "kua"))
}

func printAnagram(listAnagram map[string][]string) {
	for _, words := range listAnagram {
		for _, w := range words {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}
}

func groupAnagramWord(words ...string) map[string][]string {
	list := make(map[string][]string)

	for _, word := range words {
		s := strings.Split(word, "")
		sort.Strings(s)
		key := strings.Join(s, "")
		list[key] = append(list[key], word)
	}

	return list
}
