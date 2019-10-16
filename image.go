package timg

import (
	"fmt"
	"image"
	"image/draw"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/math/fixed"
)

func drawText(lines []string, opt *Options) (image.Image, error) {

	lineHeight := calcLineHeight(opt.TType) + opt.Tpad
	imgHeight := lineHeight*len(lines) + opt.Padding.Top + opt.Padding.Bottom
	rect := image.Rect(0, 0, opt.Width, imgHeight)
	img := image.NewNRGBA(rect)
	draw.Draw(img, rect, image.White, image.ZP, draw.Src)
	py := lineHeight + opt.Padding.Top
	for _, l := range lines {
		addText(img, opt.Padding.Left, py, l, *opt)
		py += lineHeight
	}
	return img, nil

}

func calcLineHeight(opt truetype.Options) int {
	return int(opt.DPI / 72 * opt.Size)
}

func addText(img *image.NRGBA, x, y int, text string, opt Options) {
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(opt.FontColor),
		Face: goFontFace(&opt.TType, opt.Font),
		Dot:  point,
	}
	d.DrawString(text)
}

func goFontFace(opt *truetype.Options, ttf []byte) font.Face {
	if ttf == nil {
		ttf = gomono.TTF
	}
	f, err := truetype.Parse(ttf)
	if err != nil {
		panic(fmt.Sprint("cannot parse font:", err))
	}
	nf := truetype.NewFace(f, opt)
	return nf
}
