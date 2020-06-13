package sticker

import (
	"image"

	"github.com/disintegration/imaging"
)

const (
	FullSize = 512
	HalfSize = 256
)

func New(img image.Image, height int) image.Image {
	fitImg := imaging.Fit(img, 512, height, imaging.Lanczos)
	bgImg := image.NewAlpha(image.Rect(0, 0, 512, fitImg.Bounds().Size().Y))
	dstImg := imaging.PasteCenter(bgImg, fitImg)
	return dstImg
}
