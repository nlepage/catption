package catption

import (
	"image"

	"github.com/Zenika/catption/font/impact"
	"github.com/fogleman/gg"
)

const (
	DefaultSize     = 1024
	DefaultMargin   = 20
	DefaultFontSize = 96
)

type Catption struct {
	Top      string
	Bottom   string
	Size     float64
	Margin   float64
	FontSize float64

	img image.Image
}

func LoadJPG(name string) (*Catption, error) {
	img, err := gg.LoadJPG(name)
	if err != nil {
		return nil, err
	}

	return &Catption{
		img:      img,
		Size:     DefaultSize,
		Margin:   DefaultMargin,
		FontSize: DefaultFontSize,
	}, nil
}

func (c *Catption) Image() (image.Image, error) {
	var origSize = c.img.Bounds().Size()
	var s, w, h float64

	if origSize.X > origSize.Y {
		s = c.Size / float64(origSize.X)
		w, h = c.Size, float64(origSize.Y)*s
	} else {
		s = c.Size / float64(origSize.Y)
		h, w = c.Size, float64(origSize.X)*s
	}

	ctx := gg.NewContext(int(w), int(h))

	ctx.Push()
	ctx.Scale(s, s)
	ctx.DrawImage(c.img, 0, 0)
	ctx.Pop()

	ctx.SetFontFace(impact.FontFace(c.FontSize))

	if c.Top != "" {
		c.drawText(ctx, c.Top, w/2, c.FontSize/2+c.Margin)
	}

	if c.Bottom != "" {
		c.drawText(ctx, c.Bottom, w/2, h-c.FontSize/2-c.Margin)
	}

	return ctx.Image(), nil
}

func (c *Catption) drawText(ctx *gg.Context, text string, ax, ay float64) {
	ctx.SetRGB(0, 0, 0)
	n := 6 // "stroke" size
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			x := ax + float64(dx)
			y := ay + float64(dy)
			ctx.DrawStringAnchored(text, x, y, 0.5, 0.5)
		}
	}
	ctx.SetRGB(1, 1, 1)
	ctx.DrawStringAnchored(text, ax, ay, 0.5, 0.5)
}

func (c *Catption) SaveJPG(path string) error {
	img, err := c.Image()
	if err != nil {
		return err
	}

	return gg.SaveJPG(path, img, 90)
}
