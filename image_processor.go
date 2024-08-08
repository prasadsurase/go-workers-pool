package main

import (
	"fmt"
	"time"
)

// ImageProcessor holds the image url.
type ImageProcessor struct {
	imageUrl string
}

// Process method mimicks retrieving an image.
func (ip *ImageProcessor) Process() {
	fmt.Printf("Retrieving image %s\n", ip.imageUrl)
	time.Sleep(5 * time.Second)
}
