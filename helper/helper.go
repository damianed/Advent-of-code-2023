package helper

import(
  "os"
  "bufio"
)

func ReadFile(fileLocation string) []string {
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
