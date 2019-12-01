// This is my program to find the opencv version

package main

import (
  "fmt"
  "gocv.io/x/gocv"
)

func main() {
  fmt.Printf("gocv Version: %s\n", gocv.Version())
  fmt.Printf("OpenCV Version: %s\n", gocv.OpenCVVersion())
}
