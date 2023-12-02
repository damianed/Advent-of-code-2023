package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  file, _ := os.Open("input.txt")

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  var numbers []int

  digitWords := []string {"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

  for scanner.Scan() {
    line := scanner.Text()

    first := ""
    last := ""

    for i := 0; i < len(line); i++ {
      char := string(byte(line[i]))

      for j := 0; j < len(digitWords); j++ {
        substr := line[i:]
        if (strings.HasPrefix(substr, digitWords[j])) {
          char = strconv.Itoa(j + 1)
          break
        }
      }

      if _, err := strconv.Atoi(char); err == nil {
        if (first == "") {
          first = char
        } else {
          last = char
        }
      }
    }

    if (last == "") {
      last = first
    }

    number, _ := strconv.Atoi(first + last)
    numbers = append(numbers, number)
  }

  sum := 0

  for i := 0; i < len(numbers); i++ {
    sum += numbers[i]
  }

  fmt.Println(sum)

  file.Close()
}
