package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

const (
	peak   = 1
	valley = 2
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	// outputting into file so I can view the result in my code editor with svg preview
	file, fileErr := os.Create("surface.svg")
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	fmt.Fprintf(file, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	fmt.Fprintf(file, "<style type=\"text/css\">")
	fmt.Fprintf(file, ".blue {fill: #0000FF; }\n .red {fill:#FF0000;}")
	fmt.Fprintf(file, "</style>")

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, type1 := corner(i+1, j)
			bx, by, type2 := corner(i, j)
			cx, cy, type3 := corner(i, j+1)
			dx, dy, type4 := corner(i+1, j+1)

			var current string

			if type1 == valley || type2 == valley || type3 == valley || type4 == valley {
				current = "blue"
			} else if type1 == peak || type2 == peak || type3 == peak || type4 == peak {
				current = "red"
			} else {
				current = "grey"
			}

			fmt.Fprintf(file, "<polygon points='%g, %g %g,%g %g,%g %g,%g'" + " class=\"%s\"/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, current)

		}
	}
	fmt.Println(file, "</svg>")
}

func corner(i, j int) (float64, float64, int) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, kind := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, kind
}

// finding the 2nd derivative of f(x) where f(x)= sin(x)/x
// math taken from https://xingdl2007.gitbooks.io/gopl-soljutions/content/chapter-3-basic-data-types.html
// because I don't remember calc too much

func f(x, y float64) (float64, int) {
	d := math.Hypot(x, y)
	r := math.Hypot(x, y) // distance from (0,0)
	kind := 0
	if math.Abs(d-math.Tan(d)) < 3 {
		kind = peak
		if 2*(math.Sin(d)-d*math.Cos(d))-d*d*math.Sin(d) > 0 {
			kind = valley
		}
	}
	return math.Sin(r) / r, kind
}
