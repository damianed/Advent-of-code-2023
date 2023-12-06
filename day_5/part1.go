package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "unicode"
  "strconv"
  "math"
)

func main() {
  lines := readFile("input")
  seeds, maps := getSeedsAndMaps(lines)

  minLocation := math.MaxInt32
  for _, seed := range seeds {
    soil := getValueFromMap(seed, maps["seed-to-soil"])
    fertilizer := getValueFromMap(soil, maps["soil-to-fertilizer"])
    water := getValueFromMap(fertilizer, maps["fertilizer-to-water"])
    light := getValueFromMap(water, maps["water-to-light"])
    temperature := getValueFromMap(light, maps["light-to-temperature"])
    humidity := getValueFromMap(temperature, maps["temperature-to-humidity"])
    location := getValueFromMap(humidity, maps["humidity-to-location"])

    if (minLocation > location) {
      minLocation = location
    }
  }
  fmt.Println(minLocation)
}

func getValueFromMap(needle int, haystack [][]int) int {
  for _, element := range haystack {
    if (element[0] <= needle && element[1] >= needle) {
      return element[2] + (needle - element[0] )
    }
  }

  return needle
}

func getSeedsAndMaps(lines []string) ([]int, map[string][][]int) {
  seeds := []int{}
  maps := make(map[string][][]int)
  for i := 0; i < len(lines) - 1; i++ {
    if (i == 0) {
      seedStrings := strings.Fields(lines[i])[1:]
      for _, element := range seedStrings {
        value, _ := strconv.Atoi(element)
        seeds = append(seeds, value)
      }
      continue
    }

    mapName := strings.Split(lines[i], " ")[0]
    i++
    for i < len(lines) && unicode.IsDigit(rune(lines[i][0])) {
      if _, ok := maps[mapName]; !ok {
        maps[mapName] = [][]int{}
      }

      fields := strings.Fields(lines[i])
      destination, _ := strconv.Atoi(fields[0])
      source, _ := strconv.Atoi(fields[1])
      length, _ := strconv.Atoi(fields[2])

      maps[mapName] = append(maps[mapName], []int{source, source + length - 1, destination})
      i++
    }
    i--
  }

  return seeds, maps
}

func readFile(fileLocation string) []string {
  var lines []string
  file, _ := os.Open(fileLocation)
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  for scanner.Scan() {
    text := scanner.Text()
    if (len(text) == 0)  {
      continue
    }

    lines = append(lines, text)
  }

  return lines
}
