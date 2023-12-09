package main

import (
  "fmt"
  "strings"
  "strconv"

  "adventofcode/helper"
)

func main() {
  lines := helper.ReadFile("input")

  nextNums := []int{}
  for _, line := range lines {
    numberslines := strings.Fields(line)
    numbers := []int{}
    for _, element := range numberslines {
        num, _ := strconv.Atoi(element)
        numbers = append(numbers, num)
    }
    nextNums = append(nextNums, getNextNum(numbers))
  }

  sum := 0
  for _, num := range nextNums {
    sum += num
  }

  fmt.Println(sum)
}

func getNextNum(numbers []int) int {
  if (allZeros(numbers)) {
    return 0
  }

  newl := []int{}
  prev := numbers[0]
  for _, num := range numbers[1:] {
    newl = append(newl, num - prev)
    prev = num
  }

  return numbers[len(numbers) - 1] + getNextNum(newl)
}

func allZeros(numbers []int) bool {
  if (len(numbers) == 0)  {
    return true
  }

  for idx, num := range numbers {
    if num != 0 {
      break
    }

    if idx == len(numbers) - 1 {
      return true
    }
  }

  return false
}
