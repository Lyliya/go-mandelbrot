package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

var palette = []color.RGBA{
	{255,0,0,255},
	{0,255,0,255},
	{0,0,255,255},
	{128, 0, 0, 255},
	{128, 0, 32, 255},
	{128, 0, 64, 255},
	{128, 0, 96, 255},
	{128, 0, 128, 255},
	{128, 0, 160, 255},
	{128, 0, 192, 255},
	{128, 0, 224, 255},
	{128, 0, 255, 255},
	{128, 32, 0, 255},
	{128, 32, 32, 255},
	{128, 32, 64, 255},
	{128, 32, 96, 255},
	{128, 32, 128, 255},
	{128, 32, 160, 255},
	{128, 32, 192, 255},
	{128, 32, 224, 255},
	{128, 32, 255, 255},
	{128, 64, 0, 255},
	{128, 64, 32, 255},
	{128, 64, 64, 255},
	{128, 64, 96, 255},
	{128, 64, 128, 255},
	{128, 64, 160, 255},
	{128, 64, 192, 255},
	{128, 64, 224, 255},
	{128, 64, 255, 255},
	{128, 96, 0, 255},
	{128, 96, 32, 255},
	{128, 96, 64, 255},
	{128, 96, 96, 255},
	{128, 96, 128, 255},
	{128, 96, 160, 255},
	{128, 96, 192, 255},
	{128, 96, 224, 255},
	{128, 96, 255, 255},
	{128, 128, 0, 255},
	{128, 128, 32, 255},
	{128, 128, 64, 255},
	{128, 128, 96, 255},
	{128, 128, 128, 255},
	{128, 128, 160, 255},
	{128, 128, 192, 255},
	{128, 128, 224, 255},
	{128, 128, 255, 255},
	{128, 160, 0, 255},
	{128, 160, 32, 255},
	{128, 160, 64, 255},
	{128, 160, 96, 255},
	{128, 160, 128, 255},
	{128, 160, 160, 255},
	{128, 160, 192, 255},
	{128, 160, 224, 255},
	{128, 160, 255, 255},
	{128, 192, 0, 255},
	{128, 192, 32, 255},
	{128, 192, 64, 255},
	{128, 192, 96, 255},
	{128, 192, 128, 255},
	{128, 192, 160, 255},
	{128, 192, 192, 255},
	{128, 192, 224, 255},
	{128, 192, 255, 255},
	{128, 224, 0, 255},
	{128, 224, 32, 255},
	{128, 224, 64, 255},
	{128, 224, 96, 255},
	{128, 224, 128, 255},
	{128, 224, 160, 255},
	{128, 224, 192, 255},
	{128, 224, 224, 255},
	{128, 224, 255, 255},
	{128, 255, 0, 255},
	{128, 255, 32, 255},
	{128, 255, 64, 255},
	{128, 255, 96, 255},
	{128, 255, 128, 255},
	{128, 255, 160, 255},
	{128, 255, 192, 255},
	{128, 255, 224, 255},
	{128, 255, 255, 255},
}

func clamp(x float64, a int, b int) float64 {
	return float64(math.Max(float64(a), math.Min(float64(x), float64(b))))
}

func linear_interpolate(color1 color.RGBA, color2 color.RGBA, ratio float64) color.RGBA {
	r := math.Floor((float64(color2.R - color1.R))) * ratio + float64(color1.R)
	g := math.Floor((float64(color2.G - color1.G))) * ratio + float64(color1.G)
	b := math.Floor((float64(color2.B - color1.B))) * ratio + float64(color1.B)

	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

func interpolation(i float64) color.RGBA {
	color1 := palette[int(math.Floor(float64(i)))]
	color2 := palette[int(math.Floor(float64(i))) + 1]
	return linear_interpolate(color1, color2, math.Mod(i, 1))
}

func Mandelbrot(width int, height int, x1 float64, x2 float64, y1 float64, y2 float64, iteration_max int) []color.RGBA {
	pixels := make([]color.RGBA, width * height)
	
	zoom_x := float64(width) / (x2 - x1)
	zoom_y := float64(height) / (y2 - y1)

	for Px := 0; Px < width; Px++ {
		for Py := 0; Py < height; Py++ {
			x0 := float64(Px) / zoom_x + x1
			y0 := float64(Py) / zoom_y + y1;
			x := 0.0;
			y := 0.0;
			i := 0;

			for x*x + y*y <= (1 << 16) && i < iteration_max {
				xtemp := x*x - y*y + x0
				y = 2 * x * y + y0
				x = xtemp
				i += 1
			}

			float_i := float64(i)
			if i < iteration_max {
				log_zn := math.Log(x * x + y * y) / 2
				nu := math.Log(log_zn / math.Log(2)) / math.Log(2)
				float_i = float_i + 1 - nu
			}
			inter_color := interpolation(float_i / float64(iteration_max) * float64(len(palette) - 2))
			pixels[Py * width + Px] = inter_color;
		}
	}

	return pixels
}

func main() {
	// Set IMG Size
	width := 1000
	height := 1000

	// Create Image
	upLeft := image.Point{0 ,0}
	bottomRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, bottomRight})

	pixels := Mandelbrot(width, height, -2.1, 0.6, -1.2, 1.2, 1000)

	for i := 0; i < len(pixels); i++ {
		x := i % width
		y := i / width
		img.Set(x, y, pixels[i])
	}

	// Write Image to filesystem
	f, _ := os.Create("render.png")
	png.Encode(f, img)
}