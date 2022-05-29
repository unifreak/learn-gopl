// Color each polygon based on its height, so that the peaks are colored red
// (#ff0000) and the valleys blue (#0000ff)
//
// @todo
// - how to decide color from blue to red according to height?
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x,y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, c1, aok := corner(i+1, j)
			bx, by, c2, bok := corner(i, j)
			cx, cy, c3, cok := corner(i, j+1)
			dx, dy, c4, dok := corner(i+1, j+1)
			if !aok || !bok || !cok || !dok {
				continue
			}

			color := (c1 + c2 + c3 + c4) / 4
			fmt.Printf("<polygon stroke=\"#%6X\" points='%g,%g %g,%g %g,%g %g,%g' />\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (sx, sy float64, c int, ok bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compoute surface height z.
	z := f(x, y)
	signs := [3]int{-1, 1, 0}
	for s := range signs {
		if math.IsInf(z, s) {
			return 0, 0, 0, false
		}
	}

	// Compute color string

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, paint(z), true
}

const colors = []string{
	"#0000ff"
	"#0000cc"
	"#000099"
	"#000066"
	"#000033"
	"#330000"
	"#660000"
	"#990000"
	"#ff0000"
}
func paint(z float64) int {
	return int(z) * stride
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
