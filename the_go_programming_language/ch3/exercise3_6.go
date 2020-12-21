package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y0 := float64(py)/height*(ymax-ymin) + ymin
		y1 := (float64(py)+0.5)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x0 := float64(px)/width*(xmax-xmin) + xmin
			x1 := (float64(px)+0.5)/width*(xmax-xmin) + xmin
			z0 := complex(x0, y0)
			z1 := complex(x1, y0)
			z2 := complex(x0, y1)
			z3 := complex(x1, y1)
			color := getAverageColor([]color.Color{
				mandelbort(z0),
				mandelbort(z1),
				mandelbort(z2),
				mandelbort(z3),
			})
			img.Set(px, py, color)
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbort(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func getAverageColor(colors []color.Color) color.Color {
	if len(colors) < 1 {
		return nil
	}

	var r, g, b, a float64

	for _, color := range colors {
		dr, dg, db, da := color.RGBA()
		r += float64(dr>>8) / float64(len(colors))
		g += float64(dg>>8) / float64(len(colors))
		b += float64(db>>8) / float64(len(colors))
		a += float64(da>>8) / float64(len(colors))
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}
