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

  instructions := []string{}
  nodes := map[string]Node{}
  startingNodes := []string{}
  for idx, line := range lines {
    if idx == 0 {
      for _, c := range line {
        instructions = append(instructions, string(c))
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

  steps := []int{}
  for _, node := range startingNodes {
    steps = append(steps, moveTilEnd(node, nodes, instructions))
  }

  fmt.Println(lcmAll(steps[0], steps[1:]...))
}

func moveTilEnd(currentNode string, nodes map[string]Node, instructions []string) int {
    steps := 0
    for currentNode[len(currentNode)-1:len(currentNode)] != "Z" {
      instruction := instructions[steps % len(instructions)]
      if (instruction == "R") {
        currentNode = nodes[currentNode].right
      } else {
        currentNode = nodes[currentNode].left
      }

      steps++
    }

    return steps
}

// gcd,lcm,lcmAll copied from:
// https://github.com/torbensky/advent-of-code-2023/blob/main/day08/main.go#L70C1-L88C2
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func lcmAll(a int, bs ...int) int {
	result := a
	for _, b := range bs {
		result = lcm(result, b)
	}

	return result
}
