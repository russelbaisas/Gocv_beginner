//Sample for opening videos and showing videos

package main

import (
  "gocv.io/x/gocv"
)

func main() {
  webcam, _ := gocv.VideoCaptureDevice(0)
  window := gocv.NewWindow("Show Video")
  img := gocv.NewMat()

    for {
      webcam.Read(&img)
      window.IMShow(img)
      window.WaitKey(1)
    }
}
