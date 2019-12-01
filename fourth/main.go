//Sample face detection from gocv

package main

import (
  "fmt"
  "image"
  "image/color"
  "os"
  "strconv"

  "gocv.io/x/gocv"
)

func main() {
  if len(os.Args) < 3 {
    fmt.Println("Need to run with 3 arguments eg. go run main.go 0 <cascade file>")
    return
  }

  deviceID, _ := strconv.Atoi(os.Args[1]) // parse decimal string to int
  xmlFile := os.Args[2] // get the cascade file in argument 2

  // Prepare the webcam
  webcam, err := gocv.VideoCaptureDevice(int(deviceID))
  if err != nil {
    fmt.Println(err)
    return
  }
  defer webcam.Close()

  // Open a new window
  window := gocv.NewWindow("Detect Face")
  defer window.Close()

  // Create a new image matrix
  img := gocv.NewMat()
  defer img.Close()

  blue := color.RGBA{0, 0, 255, 0}

  // load the cascade classifier
  classifier := gocv.NewCascadeClassifier()
  defer classifier.Close()

  if !classifier.Load(xmlFile) {
    fmt.Println("Error reading cascade file")
    return
  }
  fmt.Printf("start reading camera device: %v/n", deviceID)

  for {
    if ok := webcam.Read(&img); !ok {
      fmt.Printf("cannot read device")
      return
    }
    if img.Empty() {
      continue
    }
    //detecting the face
    rects := classifier.DetectMultiScale(img)

    for _, r := range rects {
      gocv.Rectangle(&img, r, blue, 3)

      size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
      pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
      gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
    }
    window.IMShow(img)
    if window.WaitKey(1) >= 0 {
      break
    }
  }
}
