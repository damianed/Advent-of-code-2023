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
  for scanner.Scan() {
    line := scanner.Text()
    data := strings.Split(line, ": ")
    numbers := strings.Split(data[1], " | ")
    winning := strings.Split(numbers[0], " ")
    has := strings.Split(numbers[1], " ")

    w := getWinningNumbers(winning, has)
    gameSum := 0

    for i := 0; i < len(w); i++ {
      if (gameSum == 0) {
        gameSum = 1;
        continue
      }

      gameSum *= 2
    }

    sum += gameSum
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
