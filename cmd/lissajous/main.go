// Lissajous generates GIF animations of random Lissajous figures
// example usage: 'go build main.go' && './main >out.gif' (open gif locally)
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var red = color.RGBA{
	R: 255,
	G: 0,
	B: 0,
	A: 0xff,
}

var green = color.RGBA{
	R: 0,
	G: 255,
	B: 0,
	A: 0xff,
}

var blue = color.RGBA{
	R: 0,
	G: 0,
	B: 255,
	A: 0xff,
}

var palette = []color.Color{color.White, color.Black, red, green, blue}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	rand.Seed(time.Now().UnixNano())
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(rand.Intn(len(palette))+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}