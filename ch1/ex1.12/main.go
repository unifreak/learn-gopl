// Modify the Lissajous server to read parameter values from the URL. For
// example, you might arrange it so that a URL like
// 		http://localhost:8000/?cycles=20
// sets the number of cycles to 20 instead of the default 5. Use the strconv.Atoi
// funciton to convert the string parameter into an integer. You can see its
// documentation with go doc strconv.Atoi
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

var (
	cycles  = 5     // number of complete x oscillator revolutions
	res     = 0.001 // angular resolution
	size    = 100   // image canvas covers [-size..+size]
	nframes = 64    // number of animation frames
	delay   = 8     // delay between frames in 10ms units
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalf("handler parsing form: %v", err)
	}

	if fcycles := r.FormValue("cycles"); fcycles != "" {
		if v, err := strconv.Atoi(fcycles); err == nil {
			cycles = v
		}
	}

	if v := r.FormValue("res"); v != "" {
		if vr, err := strconv.ParseFloat(v, 64); err == nil {
			res = vr
		}
	}

	if v := r.FormValue("size"); v != "" {
		if vs, err := strconv.Atoi(v); err == nil {
			size = vs
		}
	}

	lissajous(w)
}

func lissajous(out io.Writer) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // 201*201 rect
		img := image.NewPaletted(rect, palette)      // initial set to pallete's zero value: color.White
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			// every loop generate new image by setting pixel to black
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		// then newly generated image is appended as new frame of gif
		phase += 0.1
		anim.Delay = append(anim.Delay, delay) // sturct field can be accessed by dot notation
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
