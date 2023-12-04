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

      if (string(char) == "*") {
        var adjacentNums []int
        var num int
        // add number from line above
        if i > 0 {
          if j > 0 {
            num, lines = getNum(lines, i - 1, j -1)
            if num != -1 {
              adjacentNums = append(adjacentNums, num)
            }
          }

          num, lines = getNum(lines, i - 1, j)
          if num != -1 {
            adjacentNums = append(adjacentNums, num)
          }

          if j < len(lines[i]) - 1 {
            num, lines = getNum(lines, i - 1, j + 1)
            if num != -1 {
              adjacentNums = append(adjacentNums, num)
            }
          }
        }

        //add numbers from same line
        if j > 0 {
          num, lines = getNum(lines, i, j - 1)
          if num != -1 {
            adjacentNums = append(adjacentNums, num)
          }
        }

        if j < len(lines[i]) - 1 {
          num, lines = getNum(lines, i, j + 1)
          if num != -1 {
            adjacentNums = append(adjacentNums, num)
          }
        }

        //add numbers from line below
        if i < len(lines) - 1 {
          if j > 0 {
            num, lines = getNum(lines, i + 1, j -1)
            if num != -1 {
              adjacentNums = append(adjacentNums, num)
            }
          }

          num, lines = getNum(lines, i + 1, j)
          if num != -1 {
            adjacentNums = append(adjacentNums, num)
          }

          if j < len(lines[i]) - 1 {
            num, lines = getNum(lines, i + 1, j + 1)
            if num != -1 {
              adjacentNums = append(adjacentNums, num)
            }
          }
        }

        fmt.Println(adjacentNums)
        if (len(adjacentNums) == 2) {
          sum += adjacentNums[0] * adjacentNums[1]
        }
      }
    }
  }

  fmt.Println(sum)
}


func getNum(lines []string, i int, j int) (int, []string) {
    newString := lines[i]
    num := -1
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

      newString = newString[:start] + replacement + newString[end + 1:]
      num, _ = strconv.Atoi(numString)
      fmt.Println(num)
    }

    lines[i] = newString
    return num, lines
}
