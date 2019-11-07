package main

import (
	"os"
	"os/exec"

	"github.com/fogleman/gg"
	"github.com/spf13/cobra"
)

const (
	maxSize  = 1024
	fontSize = 96
	margin   = 20
)

var (
	cmd         *cobra.Command
	top, bottom string
)

func init() {
	cmd = &cobra.Command{
		Use:   "catption",
		Short: "Cat caption generator CLI",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			img, err := gg.LoadJPG("cats/" + args[0] + ".jpg")
			if err != nil {
				return err
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
				return err
			}

			if top != "" {
				drawText(ctx, top, w/2, fontSize/2+margin)
			}

			if bottom != "" {
				drawText(ctx, bottom, w/2, h-fontSize/2-margin)
			}

			if err := gg.SaveJPG("out.jpg", ctx.Image(), 90); err != nil {
				return err
			}

			return exec.Command("xdg-open", "out.jpg").Run()
		},
	}

	cmd.Flags().StringVar(&top, "top", "", "Top text")
	cmd.Flags().StringVar(&bottom, "bottom", "", "Bottom text")
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func drawText(ctx *gg.Context, text string, ax, ay float64) {
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
