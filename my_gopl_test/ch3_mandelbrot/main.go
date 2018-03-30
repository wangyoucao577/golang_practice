// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	//"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {

	/*
		// complex test
		var c1 complex64 = complex(1, 1)
		var c2 complex64 = complex(2, 3)
		fmt.Println(c1)    // 1+1i
		fmt.Println(c2)    // 2+3i
		fmt.Println(real(c2))  //2
		fmt.Println(imag(c2))  //3
		fmt.Println(c1 + c2)   //3+4i
		fmt.Println(c1 - c2)   //-1-2i
		fmt.Println(c1 * c2)
		fmt.Println(c1 / c2)
		return
	*/

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) //NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
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
