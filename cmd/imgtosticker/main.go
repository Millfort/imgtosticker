package main

import (
	"image/png"
	"log"
	"os"

	"github.com/millfort/imgfit/sticker"
)

func main() {
	srcFile, err := os.Open("in.png")
	if err != nil {
		log.Fatal(err)
	}
	srcImg, err := png.Decode(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	dstImg := sticker.New(srcImg, sticker.HalfSize)
	dstFile, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(dstFile, dstImg)
	if err != nil {
		log.Fatal(err)
	}
}
