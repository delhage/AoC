package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

var m = make(map[int])

func ReadFreq(filePath string, init int) (int, error) {
  fileHandle,_ := os.Open(filePath)
  defer fileHandle.Close()
  fileScanner := bufio.NewScanner(fileHandle)

  var result int = init
   
  for fileScanner.Scan() {
    // fmt.Println(fileScanner.Text())
    i, err := strconv.Atoi(fileScanner.Text())
    if err != nil {
      return result, err
    }
    result += i
    m[result] += 1
    if m[result] == 2 {
       return result, fileScanner.Err()
  }
  //fmt.Println(result)
  return result, fileScanner.Err()
}

func main() {
  freq,err := ReadFreq(os.Args[1])
  fmt.Println(freq, err)
}
