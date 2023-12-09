package main

import (
  "fmt"

  "adventofcode/helper"
  "strings"
)

type Node struct {
  left string
  right string
}

func main() {
  lines := helper.ReadFile("input")

  lf := []string{}
  nodes := map[string]Node{}
  for idx, line := range lines {
    if idx == 0 {
      for _, c := range line {
        lf = append(lf, string(c))
      }
      continue
    }

    data := strings.Fields(line)
    //fmt.Println(strings.Strip(data[2], ))
    left := strings.Trim(data[2], ",)(")
    right := strings.Trim(data[3], ",)(")
    nodes[data[0]] = Node{left, right}
  }

  currNode := "AAA"
  steps := 0
  for currNode != "ZZZ" {
    instruction := lf[steps % len(lf)]
    if (instruction == "R") {
      currNode = nodes[currNode].right
    } else {
      currNode = nodes[currNode].left
    }

    steps++
  }

  fmt.Println(steps)
}
