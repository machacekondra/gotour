package main

import (
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  words := make(map[string]int)

  for _, w := range strings.Fields(s) {
    if _, t := words[w]; t == false {
      words[w] = 0
    }
    words[w] += 1
  }

  return words
}

func main() {
  wc.Test(WordCount)
}
