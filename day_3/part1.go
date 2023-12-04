package main

import (
  "bufio"
  "fmt"
  "os"
  "unicode"
  "strconv"
)


func main() {
  file, _ := os.Open("input")

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  var lines []string

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  file.Close()

  sum := 0
  for i := 0; i < len(lines); i++ {
    for j := 0; j < len(lines[i]); j++ {
      char := rune(lines[i][j])

      if (string(char) != "." && !unicode.IsDigit(char)) {
        // add number from line above
        if i > 0 {
          if j > 0 {
            lines = addNum(&sum, lines, i - 1, j -1)
          }

          lines = addNum(&sum, lines, i - 1, j)

          if j < len(lines[i]) - 1 {
            lines = addNum(&sum, lines, i - 1, j + 1)
          }
        }

        //add numbers from same line
        if j > 0 {
          lines = addNum(&sum, lines, i, j - 1)
        }

        if j < len(lines[i]) - 1 {
          lines = addNum(&sum, lines, i, j + 1)
        }

        //add numbers from line below
        if i < len(lines) - 1 {
          if j > 0 {
            lines = addNum(&sum, lines, i + 1, j -1)
          }

          lines = addNum(&sum, lines, i + 1, j)

          if j < len(lines[i]) - 1 {
            lines = addNum(&sum, lines, i + 1, j + 1)
          }
        }
      }
    }
  }

  fmt.Println(sum)
}


func addNum(sum *int, lines []string, i int, j int) []string {
    newline := lines[i]
    if unicode.IsDigit(rune(lines[i][j])) {
      start := j
      end := j

      for start >= 0 && end <= len(lines[i]) - 1 {
        prevStart := start
        if (start > 0 && unicode.IsDigit(rune(lines[i][start - 1]))) {
          start -= 1
        }

        prevEnd := end
        if (end < len(lines[i]) - 1 && unicode.IsDigit(rune(lines[i][end + 1]))) {
          end += 1
        }

        if (prevEnd == end && prevStart == start) {
          break
        }
      }

      numString := ""
      replacement := ""
      for k := start; k <= end; k++ {
        replacement += "."
        numString += string(byte(lines[i][k]))
      }

      num, _ := strconv.Atoi(numString)

      newline = newline[:start] + replacement + newline[end + 1:]
      *sum += num
    }

    lines[i] = string(newline)
    return lines
}
