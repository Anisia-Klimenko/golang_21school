package main

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{R: 250, G: 250, B: 250, A: 255}
	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func EachPixel(img *image.RGBA, w, h int, colorFunction func(int, int) color.RGBA) {
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, colorFunction(x, y))
		}
	}
}

func main() {
	width := 300
	height := 300

	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	//green := color.RGBA{R: 0, G: 225, B: 179, A: 0xff}
	//violet := color.RGBA{R: 134, G: 27, B: 226, A: 0xff}
	//black := color.RGBA{R: 22, G: 29, B: 39, A: 0xff}
	white := color.RGBA{R: 235, G: 236, B: 244, A: 0xff}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < 2*width/3 && y < 2*height/3:
				img.Set(x, y, color.RGBA{
					G: 255, B: 179, A: uint8(math.Sqrt(math.Pow(float64(x), 2)+math.Pow(float64(y), 2)) * 100 / math.Sqrt(math.Pow(float64(width), 2)+math.Pow(float64(height), 2))),
				})
				//EachPixel(img, 2*width/3, 2*height/3, func(x, y int) color.RGBA {
				//	return color.RGBA{
				//		G: 255, B: 179, A: uint8(math.Sqrt(math.Pow(float64(x), 2)+math.Pow(float64(y), 2)) * 100/math.Sqrt(math.Pow(float64(width), 2)+math.Pow(float64(height), 2))),
				//	}
				//})
			case x > width/3 && y < 2*height/3:
				img.Set(x, y, color.RGBA{
					R: 235, G: 236, B: 244, A: uint8(math.Sqrt(math.Pow(float64(x), 2)+math.Pow(float64(y), 2)) * 100 / math.Sqrt(math.Pow(float64(width), 2)+math.Pow(float64(height), 2))),
				})
				//EachPixel(img, width/3, 2*height/3, func(x, y int) color.RGBA {
				//	return color.RGBA{
				//		R: 235, G: 236, B: 244, A: uint8(math.Sqrt(math.Pow(float64(x), 2)+math.Pow(float64(y), 2)) * 100/math.Sqrt(math.Pow(float64(width), 2)+math.Pow(float64(height), 2))),
				//	}
				//})
			case x < 2*width/3 && y > height/3:
				img.Set(x, y, color.RGBA{
					R: 22, G: 29, B: 39, A: uint8(math.Sqrt(math.Pow(float64(x), 2)+math.Pow(float64(y), 2)) * 100 / math.Sqrt(math.Pow(float64(width), 2)+math.Pow(float64(height), 2))),
				})
				//EachPixel(img, 2*width/3, height/3, func(x, y int) color.RGBA {
				//	return color.RGBA{
				//		R: 22, G: 29, B: 39, A: uint8(math.Sqrt(math.Pow(float64(x), 2)+math.Pow(float64(y), 2)) * 100/math.Sqrt(math.Pow(float64(width), 2)+math.Pow(float64(height), 2))),
				//	}
				//})
			case x > width/3 && y > height/3:
				img.Set(x, y, color.RGBA{
					R: 134, G: 27, B: 226, A: uint8(math.Sqrt(math.Pow(float64(x), 2)+math.Pow(float64(y), 2)) * 100 / math.Sqrt(math.Pow(float64(width), 2)+math.Pow(float64(height), 2))),
				})
				//EachPixel(img, width/3, height/3, func(x, y int) color.RGBA {
				//	return color.RGBA{
				//		R: 134, G: 27, B: 226, A: uint8(math.Sqrt(math.Pow(float64(x), 2)+math.Pow(float64(y), 2)) * 100/math.Sqrt(math.Pow(float64(width), 2)+math.Pow(float64(height), 2))),
				//	}
				//})
			default:
				img.Set(x, y, white)

				//case x < width/4 && y >= 3*height/4:
				//	img.Set(x, y, violet)
				//case x >= 3*width/4 && y < height/4:
				//	img.Set(x, y, black)
				//case x < width/4 && y < height/4:
				//	img.Set(x, y, green)
				//case x >= width/4 && x < width/2 && y >= height/4 && y < height/2:
				//	img.Set(x, y, green)
				//case x >= width/2 && x < 3*width/4 && y >= height/2 && y < 3*height/4:
				//	img.Set(x, y, green)
				//case x >= 3*width/4 && y >= 3*height/4:
				//	img.Set(x, y, green)
				//default:
				//	img.Set(x, y, white)
			}
		}
	}

	addLabel(img, 78*width/100, height/10, "acristin")
	addLabel(img, 78*width/100, height/6, "21 school")

	f, _ := os.Create("amazing_logo.png")
	err := png.Encode(f, img)
	if err != nil {
		return
	}
}
