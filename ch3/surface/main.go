// Surface computes an SVG rendering of a 3-D surface function.
//
// The essence of the program is mapping between three different coordinate systems.
//
// The first is a 2-D grid of 100*100 cells identified by integer coordinates (i, j).
//
// The second coordinate system is a mesh of 3-D floating-point coordinates (x,y,z),
// where x and y are linear functions of i and j, translated so that the origin is in
// the center, and scaled by the constant xyrange. The height z is the value of
// the surface function f(x,y).
//
// The third coordinate system is the 2-D image canvas, with (0,0) in the top left corner.
// Points in this plane are denoted (sx, sy). We use an isometric projection to
// map each 3-D point (x,y,z) onto the 2-D canvas. A point appears farther to the
// right on the canvas the greater its x value or the smaller its y value. And a
// point appears farther down the canvas the greater its x value or y value, and
// the smaller its z value. The vertical and horizontal scale factors for x and y
// are derived from the sine and cosine of a 30° angle. The scale factor for z, 0.4,
// is an arbitrary parameter.
//
// isometric projection, see
// - http://www.gandraxa.com/isometric_projection.xml
// - https://en.wikipedia.org/wiki/Isometric_projection
package main

import (
	"fmt"
	"math"
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
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' " +
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}