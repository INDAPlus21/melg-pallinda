// Stefan Nilsson 2013-02-27

// This program creates pictures of Julia sets (en.wikipedia.org/wiki/Julia_set).
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
	"strconv"
	"sync"
	"time"
)

type ComplexFunc func(complex128) complex128

var Funcs []ComplexFunc = []ComplexFunc{
	func(z complex128) complex128 { return z*z - 0.61803398875 },
	func(z complex128) complex128 { return z*z + complex(0, 1) },
	func(z complex128) complex128 { return z*z + complex(-0.835, -0.2321) },
	func(z complex128) complex128 { return z*z + complex(0.45, 0.1428) },
	func(z complex128) complex128 { return z*z*z + 0.400 },
	func(z complex128) complex128 { return cmplx.Exp(z*z*z) - 0.621 },
	func(z complex128) complex128 { return (z*z+z)/cmplx.Log(z) + complex(0.268, 0.060) },
	func(z complex128) complex128 { return cmplx.Sqrt(cmplx.Sinh(z*z)) + complex(0.065, 0.122) },
}

func main() {
	var wg sync.WaitGroup
	before := time.Now()

	for n, fn := range Funcs {
		wg.Add(1)
		go func(n int, fn ComplexFunc) {
			defer wg.Done()
			err := CreatePng("picture-"+strconv.Itoa(n)+".png", fn, 1024)
			if err != nil {
				log.Fatal(err)
			}
		}(n, fn)
	}

	wg.Wait()

	fmt.Println("time:", time.Now().Sub(before))
}

// CreatePng creates a PNG picture file with a Julia image of size n x n.
func CreatePng(filename string, f ComplexFunc, n int) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()
	err = png.Encode(file, Julia(f, n))
	return
}

// Julia returns an image of size n x n of the Julia set for f.
func Julia(f ComplexFunc, n int) image.Image {
	bounds := image.Rect(-n/2, -n/2, n/2, n/2)
	img := image.NewRGBA(bounds)
	s := float64(n / 4)
	var wg sync.WaitGroup

	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		worksize := 300
		threads := int(math.Max(0, float64((bounds.Max.Y-bounds.Min.X)/worksize))) + 1 // + 1 to not miss the last partial row
		for t := 0; t < threads; t++ {

			wg.Add(1)
			go func(min int, max int, i int) {
				defer wg.Done()
				for j := min; j < max; j++ {
					n := Iterate(f, complex(float64(i)/s, float64(j)/s), 256)
					r := uint8(0)
					g := uint8(0)
					b := uint8(n % 32 * 8)
					img.Set(i, j, color.RGBA{r, g, b, 255})
				}
			}(bounds.Min.Y+t*worksize, int(math.Min(float64(bounds.Min.Y+(t+1)*worksize), float64(bounds.Max.Y))), i)
		}
	}

	wg.Wait() // Wait for everything to finish

	return img
}

// Iterate sets z_0 = z, and repeatedly computes z_n = f(z_{n-1}), n â‰¥ 1,
// until |z_n| > 2  or n = max and returns this n.
func Iterate(f ComplexFunc, z complex128, max int) (n int) {
	for ; n < max; n++ {
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			break
		}
		z = f(z)
	}
	return
}
