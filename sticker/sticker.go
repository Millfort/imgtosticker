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
	r := rectangle{
		Width:  FullSize,
		Height: height,
	}
	size := fillImgSize(img.Bounds().Size(), r)
	fillImg := imaging.Resize(img, size.Width, size.Height, imaging.Lanczos)
	bgImg := image.NewAlpha(image.Rect(0, 0, 512, fillImg.Bounds().Size().Y))
	dstImg := imaging.PasteCenter(bgImg, fillImg)
	return dstImg
}

type dimensions struct {
	Width  int
	Height int
}

type rectangle struct {
	Width  int
	Height int
}

// find finds x in x/bx=a/b
func find(bx, a, b int) int {
	return (bx * a) / b
}

func fillImgSize(img image.Point, r rectangle) dimensions {
	if img.X == r.Width || img.Y == r.Height {
		return dimensions{
			Width:  img.X,
			Height: img.Y,
		}
	}
	verticalDiff := img.Y - r.Height
	horizontalDiff := img.X - r.Width
	if horizontalDiff >= verticalDiff {
		width := r.Width
		height := find(width, img.Y, img.X)
		return dimensions{
			Width:  width,
			Height: height,
		}
	}
	height := r.Height
	width := find(height, img.X, img.Y)
	return dimensions{
		Width:  width,
		Height: height,
	}
}
