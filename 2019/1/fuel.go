package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func ReadFreq(filePath string) (int, error) {
  fileHandle,_ := os.Open(filePath)
  defer fileHandle.Close()
  fileScanner := bufio.NewScanner(fileHandle)

  var result int = 0
  
  for fileScanner.Scan() {
    // fmt.Println(fileScanner.Text())
    f, err := strconv.Atoi(fileScanner.Text())
    if err != nil {
      return result, err
    }
    result += f/3 - 2
  }
  //fmt.Println(result)
  return result, fileScanner.Err()
}

func main() {
  freq,err := ReadFreq(os.Args[1])
  fmt.Println(freq, err)
}
