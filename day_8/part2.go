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
  startingNodes := []string{}
  for idx, line := range lines {
    if idx == 0 {
      for _, c := range line {
        lf = append(lf, string(c))
      }
      continue
    }

    data := strings.Fields(line)
    left := strings.Trim(data[2], ",)(")
    right := strings.Trim(data[3], ",)(")
    nodeName := data[0]
    nodes[nodeName] = Node{left, right}
    fmt.Println(nodeName)
    if (nodeName[len(nodeName) - 1:len(nodeName)] == "A") {
      startingNodes = append(startingNodes, nodeName)
    }
  }

  currentNodes := startingNodes
  fmt.Println(startingNodes)
  steps := 0
  finished := false
  stepsTaken := []int{}

  for i := 0; i < len(currentNodes); i++ {
    stepsTaken = append(stepsTaken, 0)
  }

  for !finished {
    checkPoint, endingNode := moveTilEnd(currentNodes[0], nodes, lf, steps)
    steps = checkPoint
    currentNodes[0] = endingNode

    for idx, node := range currentNodes {
      if idx == 0 {
        continue
      }

      finishNode := moveTilSteps(node, nodes, lf, stepsTaken[idx], steps - stepsTaken[idx])
      stepsTaken[idx] = steps
      currentNodes[idx] = finishNode

      if idx == 3 {
        fmt.Println(finishNode)
      }
      if (finishNode[len(finishNode)-1:len(finishNode)] != "Z") {
        break
      }

      if idx == len(currentNodes) - 1 {
        finished = true
      }
    }

    currentNodes[0] = moveTilSteps(currentNodes[0], nodes, lf, steps, 1)
    steps++
  }

  fmt.Println(steps - 1)
}

func moveTilEnd(currentNode string, nodes map[string]Node, instructions []string, steps int) (int, string) {
    for currentNode[len(currentNode)-1:len(currentNode)] != "Z" {
      instruction := instructions[steps % len(instructions)]
      if (instruction == "R") {
        currentNode = nodes[currentNode].right
      } else {
        currentNode = nodes[currentNode].left
      }

      steps++
    }

    return steps, currentNode
}

func moveTilSteps(currentNode string, nodes map[string]Node, instructions []string, currSteps int, steps int) string {
    for i := 0; i < steps; i++ {
      instruction := instructions[currSteps % len(instructions)]
      if (instruction == "R") {
        currentNode = nodes[currentNode].right
      } else {
        currentNode = nodes[currentNode].left
      }

      currSteps++
    }

    return currentNode
}
