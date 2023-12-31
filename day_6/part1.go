package main

import (
  "log"
  "fmt"
  "strings"
  "strconv"

  "adventofcode/helper"
)

func main() {
  lines := helper.ReadFile("input")
  times := strings.Fields(lines[0])[1:]
  distances := strings.Fields(lines[1])[1:]

  ways := []int{}
  for idx, time := range times {
    raceTime, err := strconv.Atoi(time)

    if (err != nil) {
      log.Fatalf("error parsing int %v", err)
    }
    record, err := strconv.Atoi(distances[idx])

    if (err != nil) {
      log.Fatalf("error parsing int %v", err)
    }

    minPress := 1
    maxPress := raceTime - 1

    for minPress < raceTime {
      if minPress * (raceTime - minPress) <= record {
        minPress++
      } else {
        break
      }
    }

    for maxPress > 0 {
      if maxPress * (raceTime - maxPress) <= record {
        maxPress--
      } else {
        break
      }
    }

    ways = append(ways, maxPress - minPress + 1)
  }

  total := 1
  for _, count := range ways {
    total *= count
  }

  fmt.Println(total)
}
