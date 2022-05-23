// Modify the Lissajous program to produce images in multiple colors by adding
// more values to palette and then displaying them by changing the third
// argument of SetColoroIndex in some interesting way.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var green = color.RGBA{0x00, 0xff, 0x00, 0xff}
var red = color.RGBA{0xff, 0x00, 0x00, 0xff}
var grey = color.RGBA{0x77, 0x77, 0x77, 0xff}
var palette = []color.Color{color.Black, green, red, grey}
const (
	blackIndex = 0
	greenIndex = 1
	redIndex = 2
	greyIndex = 3
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles 	= 5
		res    	= 0.001
		size 	= 100
		nframes = 64
		delay 	= 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		paintIndex := 0
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			paintIndex = (paintIndex+1) % len(palette)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(paintIndex))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}