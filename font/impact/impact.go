package impact

import (
	"io/ioutil"
	"sync"

	"github.com/golang/freetype/truetype"
	"github.com/markbates/pkger"
	"golang.org/x/image/font"
)

var (
	ttf     *truetype.Font
	readTtf sync.Once
)

// FontFace returns impact fontface w/ size of points
func FontFace(points float64) font.Face {
	readTtf.Do(func() {
		f, err := pkger.Open("/impact.ttf")
		if err != nil {
			panic(err)
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		ttf, err = truetype.Parse(b)
		if err != nil {
			panic(err)
		}
	})

	return truetype.NewFace(ttf, &truetype.Options{
		Size: points,
	})
}
