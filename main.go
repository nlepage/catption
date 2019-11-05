package main

import (
	"os"
	"strings"

	"github.com/fogleman/gg"
)

const (
	maxSize  = 1024
	fontSize = 96
	margin   = 20
)

func main() {
	img, err := gg.LoadJPG("assets/knife.jpg")
	if err != nil {
		panic(err)
	}

	var origSize = img.Bounds().Size()
	var s, w, h float64

	if origSize.X > origSize.Y {
		s = maxSize / float64(origSize.X)
		w, h = maxSize, float64(origSize.Y)*s
	} else {
		s = maxSize / float64(origSize.Y)
		h, w = maxSize, float64(origSize.X)*s
	}

	ctx := gg.NewContext(int(w), int(h))

	ctx.Push()
	ctx.Scale(s, s)
	ctx.DrawImage(img, 0, 0)
	ctx.Pop()

	if err := ctx.LoadFontFace("assets/impact.ttf", fontSize); err != nil {
		panic(err)
	}

	text := strings.Join(os.Args[1:], " ")

	ctx.SetRGB(0, 0, 0)
	n := 6 // "stroke" size
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			x := w/2 + float64(dx)
			y := h - fontSize/2 - margin + float64(dy)
			ctx.DrawStringAnchored(text, x, y, 0.5, 0.5)
		}
	}
	ctx.SetRGB(1, 1, 1)
	ctx.DrawStringAnchored(text, w/2, h-fontSize/2-margin, 0.5, 0.5)

	if err := gg.SaveJPG("out.jpg", ctx.Image(), 90); err != nil {
		panic(err)
	}
}
