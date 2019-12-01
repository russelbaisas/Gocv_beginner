//blur the background for Lebron's image

package main

import (
  "fmt"
  "gocv.io/x/gocv"
  "os"
  "image/color"
)

func main(){
  if len(os.Args) < 2 {
    fmt.Println("Missing argument")
    return
  }

  //view lebron picture

  filename := os.Args[1]
  window := gocv.NewWindow("HSV test")

  img := gocv.IMRead(filename, gocv.IMReadGrayScale)
  resultImage := gocv.NewMatWithSize(500, 900, gocv.MatTypeCV8U)
  // RetrievalTree retrieves all of the contours and reconstructs a full
  // ChainApproxSimple compresses horizontal, vertical, and diagonal segments
  getcontours := gocv.FindContours(img, gocv.RetrievalTree, gocv.ChainApproxSimple)
  showimage := gocv.NewMatWithSize(resultImage.Rows(), resultImage.Cols(), gocv.MatChannels4)
  // apply HSV
  //find the threshold
  for index, element := range getcontours {
    fmt.Println(index)
    fmt.Println(element)
    gocv.DrawContours(&showimage, getcontours, index, color.RGBA{R: 0, G: 0, B: 255, A: 255}, 1)
    gocv.Rectangle(&showimage,
    gocv.BoundingRect(element),
    color.RGBA{R: 0, G: 255, B: 0, A: 100}, 1)
  }
  for {
    window.IMShow(showimage)
    window.IMShow(img)
    if gocv.WaitKey(1) >= 0 {
      break
    }
  }
}
