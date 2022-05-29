// Following the approach of the Lissajous example, construct a web server that
// computes surfaces and writes SVG data to the client. The server must set the
// Content-Type header like this:
//      w.Header().Set("Content-Type", "image/svg+xml")
// Allow the client to specify values like height, width, and color as HTTP
// request parameters.
//
// @todo
// - how to change color by HTTP request parameter?
package main

import (
	"fmt"
	"math"
    "net/http"
    "log"
)

const (
	width, height = 600, 320 			// canvas size in pixels
	cells 		  = 100 				// number of grid cells
	xyrange 	  = 30.0 				// axis ranges (-xyrange..+xyrange)
	xyscale 	  = width / 2 / xyrange // pixels per x or y unit
	zscale 		  = height * 0.4 		// pixels per z unit
	angle 		  = math.Pi / 6 		// angle of x,y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
    http.HandleFunc("/", surface)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "image/svg+xml")

    fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' " +
        "style='stroke: grey; fill: white; stroke-width: 0.7' " +
        "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay, aok := corner(i+1, j)
            bx, by, bok := corner(i, j)
            cx, cy, cok := corner(i, j+1)
            dx, dy, dok := corner(i+1, j+1)
            if !aok || !bok || !cok || !dok {
                continue
            }

            fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy)
        }
    }
    fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (sx, sy float64, ok bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compoute surface height z.
	z := f(x, y)
    signs := [3]int{-1, 1, 0}
    for s := range signs {
        if math.IsInf(z, s) {
            return 0, 0, false
        }
    }

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}