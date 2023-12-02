package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  file, _ := os.Open("input")

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  max := map[string]int{"red": 12, "green": 13, "blue": 14}

  sum := 0
  for scanner.Scan() {
    line := scanner.Text()
    split := strings.Split(line, ": ")
    gameId, _ := strconv.Atoi(split[0][5:])

    picks := strings.Split(split[1], "; ")
    posible := true

    for i := 0; i < len(picks); i++ {
      pick := picks[i]
      byColor := strings.Split(pick, ", ")

      for j := 0; j < len(byColor); j++ {
        colorData := strings.Split(byColor[j], " ")
        number, _ := strconv.Atoi(colorData[0])
        color := colorData[1]

        if (max[color] < number) {
          posible = false
          break
        }

        if (!posible) {
          break
        }
      }
    }

    if (posible) {
      sum += gameId
    }
  }
    fmt.Println(sum)
}
