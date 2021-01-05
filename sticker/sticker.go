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

func fillImgSize(img image.Point, r rectangle) dimensions {
	if img.X == r.Width && img.Y == r.Height {
		return dimensions{
			Width:  img.X,
			Height: img.Y,
		}
	}
	rr := float64(r.Width) / float64(r.Height)
	ir := float64(img.X) / float64(img.Y)
	if rr > ir {
		k := float64(r.Height) / float64(img.Y)
		return dimensions{
			Width:  int(float64(img.X) * k),
			Height: r.Height,
		}
	}
	k := float64(r.Width) / float64(img.X)
	return dimensions{
		Width:  r.Width,
		Height: int(float64(img.Y) * k),
	}
}
