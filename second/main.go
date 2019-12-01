// This program is a basic of how to open a video

package main

import (
  "fmt"
  "os"
  "gocv.io/x/gocv"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("go run main.go [camera ID]")
    return
  }
}
