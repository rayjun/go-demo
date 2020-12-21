package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			showSvg(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	showSvg(os.Stdout)
}

func showSvg(out io.Writer) {

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	maxHeight, minHeight := getMaxMinHeight()

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			color := getColor(getHeight(i, j), maxHeight, minHeight)
			if isFinite(ax) && isFinite(ay) && isFinite(bx) && isFinite(by) && isFinite(cx) && isFinite(cy) && isFinite(dx) && isFinite(dy) {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func isFinite(f float64) bool {
	if math.IsInf(f, 0) {
		return false
	}
	if math.IsNaN(f) {
		return false
	}
	return true
}

func getHeight(i, j int) float64 {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	return f(x, y)
}

func getMaxMinHeight() (float64, float64) {
	maxHeight := math.NaN()
	minHeight := math.NaN()

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			z := getHeight(i, j)

			if isFinite(z) {
				if math.IsNaN(maxHeight) || maxHeight < z {
					maxHeight = z
				}
				if math.IsNaN(minHeight) || minHeight > z {
					minHeight = z
				}
			}
		}
	}

	return maxHeight, minHeight
}

func getColor(height, maxHeight, minHeight float64) string {
	if !isFinite(height) || !isFinite(maxHeight) || !isFinite(minHeight) {
		return "#0000FF"
	}

	n := int((height - minHeight) / (maxHeight - minHeight) * 255)
	rr := fmt.Sprintf("%02x", n)
	gg := "00"
	bb := fmt.Sprintf("%02x", 255-n)

	return fmt.Sprintf("#%s%s%s", rr, gg, bb)
}
