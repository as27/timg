/*
Package timg allows you to generate an image. Image from a
given string.
*/
package timg

import (
	"image"
	"image/color"

	"github.com/golang/freetype/truetype"
)

// Box defines int values for Top, Right, Bottom, Left
// a little bit similar like in CSS.
type Box struct {
	Top, Right, Bottom, Left int
}

// Options which let you configure Draw()
type Options struct {
	Padding   Box // absolute padding of the image
	Width     int // width of the image
	TType     truetype.Options
	Tpad      int // padding for the line height
	FontColor color.NRGBA
	Font      []byte // TTF font
}

// Draw takes a string and generates an image.Image from the
// text. If you are ok with the default Options you can pass
// nil for *Options.
//
//   // call with default options
//   img, err := timg.Draw(txt, nil)
func Draw(lines []string, opt *Options) (image.Image, error) {
	if opt == nil {
		opt = Default()
	}
	return drawText(lines, nil)
}

// Default returns the default options for Draw(). If you want
// to change some of the default values, you call this function
// and change just that values.
//
//   opt := timg.Default()
//   opt.Width = 5000
//   img, err := timg.Draw(txt, opt)
//
// It is recomended to get always the default options, because
// that will decouple your code from that options. When some
// other options are added in later versions you will get them
// automatically.
func Default() *Options {
	return &Options{
		Padding:   Box{Top: 50, Bottom: 50, Left: 20, Right: 20},
		Width:     2000,
		TType:     truetype.Options{Size: 12, DPI: 300},
		Tpad:      10,
		FontColor: color.NRGBA{0, 0, 0, 255},
		Font:      nil,
	}
}
