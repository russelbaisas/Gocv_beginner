// Shows the picture in the new window

package main

import (
  "fmt"
  "os"
  "gocv.io/x/gocv"
  "image"
  "image/color"
)

func main() {

  if len(os.Args) < 2 {
    fmt.Println("Missing Argument")
    return
  }

  //get the file from argument 1
  filename := os.Args[1]
  window := gocv.NewWindow("Showing Image")
  img := gocv.IMRead(filename, gocv.IMReadColor)
  blue := color.RGBA{0, 0, 255, 0}
  if img.Empty() {
    fmt.Println("Unable to read image file")
    return
  }
  pt := image.Pt(60,  60)
  gocv.PutText(&img,"Lebron", pt, gocv.FontHersheyPlain, 6, blue, 4)
  for {
    window.IMShow(img)
    if window.WaitKey(1) >= 0 {
      break
    }
  }
}
