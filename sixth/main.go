// View images and put some text or color

package main

import (
  "fmt"
  "os"
  "gocv.io/x/gocv"
  "image"
  "image/color"
)

func main() {
  // open the file
  // view the images
  //add an overlay of text
  if len(os.Args) < 2 {
    fmt.Println("Missing argument")
    return
  }
  //get the file in argument
  filepath := os.Args[1]
  // create a new window
  window := gocv.NewWindow("Put a sample text")
  blue := color.RGBA{0, 0, 255, 0}
  img := gocv.IMRead(filepath, gocv.IMReadGrayScale)
  pt := image.Pt(60, 60)
  gocv.PutText(&img, "Lebron", pt, gocv.FontHersheyPlain, 6, blue, 4)
  for {
    window.IMShow(img)
    if gocv.WaitKey(1) >= 0 {
        break
    }
  }

}
