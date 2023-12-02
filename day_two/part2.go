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

  minGames := []map[string]int{}

  for scanner.Scan() {
    min := map[string]int{"red": 0, "green": 0, "blue": 0}
    line := scanner.Text()
    split := strings.Split(line, ": ")

    picks := strings.Split(split[1], "; ")

    for i := 0; i < len(picks); i++ {
      pick := picks[i]
      byColor := strings.Split(pick, ", ")

      for j := 0; j < len(byColor); j++ {
        colorData := strings.Split(byColor[j], " ")
        number, _ := strconv.Atoi(colorData[0])
        color := colorData[1]

        if (min[color] < number) {
          min[color] = number
        }
      }
    }

    minGames = append(minGames, min)
  }

  sum := 0
  for i := 0; i < len(minGames); i++ {
    game := minGames[i]
    sum += game["red"] * game["blue"] * game["green"]
  }

  fmt.Println(sum)
}
