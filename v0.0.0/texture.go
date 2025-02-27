package bizmuth

import (
	"image"
	"log"
	"os"

	"image/draw"
	_ "image/jpeg"
	_ "image/png"

	"github.com/go-gl/gl/v3.3-core/gl"
	//_ "image/gif"
)

type Image struct {
	image.Image
}

func LoadImage(path string, sampling uint32) uint32 {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(PrefixErr + "Can't open image file!")
	}

	defer file.Close()

	imgInfo, format, err := image.DecodeConfig(file)
	if err != nil {
		log.Fatal(PrefixErr + "Can't decode image!")
	}

	file.Seek(0, 0)

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(PrefixErr + "Can't decode image!")
	}

	if format != "png" && format != "jpeg" {
		log.Fatal(PrefixErr + "Unsupported image format!")
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba == nil {
		log.Fatal(PrefixErr + "Can't create RGBA image!")
	}

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	flipVertically(rgba)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, int32(sampling))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, int32(sampling))

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(imgInfo.Width), int32(imgInfo.Height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	return texture
}

func flipVertically(img *image.RGBA) {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	temp := make([]byte, width*4)

	for y := 0; y < height/2; y++ {
		top := img.Pix[y*img.Stride : (y+1)*img.Stride]
		bottom := img.Pix[(height-1-y)*img.Stride : (height-y)*img.Stride]
		copy(temp, top)
		copy(top, bottom)
		copy(bottom, temp)
	}
}
