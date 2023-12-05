package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func main() {
  file, _ := os.Open("input")

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  sum := 0
  cards := map[int]int{}
  index := 0
  for scanner.Scan() {
    line := scanner.Text()
    data := strings.Split(line, ": ")

    _, card := cards[index]
    if (card) {
      cards[index] += 1
    } else {
      cards[index] = 1
    }

    numbers := strings.Split(data[1], " | ")
    winning := strings.Split(numbers[0], " ")
    has := strings.Split(numbers[1], " ")

    w := getWinningNumbers(winning, has)

    for i := 1; i <= len(w); i++ {
      idx := index + i
      _, card := cards[index]
      if (card) {
        cards[idx] += 1 * cards[index]
      } else {
        cards[idx] = 1 * cards[index]
      }
    }

    sum += 1 + (len(w) * cards[index])
    index += 1
  }

  fmt.Println(sum)

  file.Close()
}

func getWinningNumbers(winning []string, has []string) []string {
  set := map[string]bool{}
  var num []string

  for i := 0; i < len(winning); i++ {
    if winning[i] == "" {
      continue
    }

    set[winning[i]] = true
  }

  for i := 0; i < len(has); i++ {
    if set[has[i]] {
      num = append(num, has[i])
      set[has[i]] = false
    }
  }

  return num
}
