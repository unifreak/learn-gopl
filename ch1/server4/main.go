// Server4 print a lissajous amime.
// Visit http://localhost:8000 to see the GIF image.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"log"
	"net/http"
)

// `composite literal` to init any Go's composite type
var palette = []color.Color{color.White, color.Black} // [] a slice
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

// This example demoenstrate the power of interface (here io.Writer)
// See how simple to change the output stream of lissajous gif to ResponseWriter
func main() {
	// `function literal`: a anonymous function defined at its point of use.
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer) {
	const (
		cycles 	= 5 	// number of complete x oscillator revolutions
		res    	= 0.001 // angular resolution
		size 	= 100 	// image canvas covers [-size..+size]
		nframes = 64 	// number of animation frames
		delay 	= 8 	// delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} 	// {} a struct GIF
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // 201*201 rect
		img := image.NewPaletted(rect, palette) // initial set to pallete's zero value: color.White
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			// every loop generate new image by setting pixel to black
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		// then newly generated image is appended as new frame of gif
		phase += 0.1
		anim.Delay = append(anim.Delay, delay) // sturct field can be accessed by dot notation
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}