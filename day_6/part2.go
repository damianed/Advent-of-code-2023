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

  timeStr := ""
  recordStr := ""
  for idx, time := range times {
    timeStr += time
    recordStr += distances[idx]
  }

  raceTime, err := strconv.Atoi(timeStr)

  if (err != nil) {
    log.Fatalf("error parsing int %v", err)
  }
  record, err := strconv.Atoi(recordStr)

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

  fmt.Println(maxPress - minPress + 1)
}
