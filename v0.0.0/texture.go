package bizmuth

import (
	"image"
	"log"
	"os"

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

	//Optimize here later
	imgInfo, format, err := image.DecodeConfig(file)
	img, _, _ := image.Decode(file)
	if err != nil {
		log.Fatal(PrefixErr + "Can't decode image!")
	}

	if format != "png" && format != "jpeg" {
		log.Fatal(PrefixErr + "Unsupported image format!")
	}

	rgba := image.NewRGBA(img.Bounds())

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, int32(sampling))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, int32(sampling))

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(imgInfo.Width), int32(imgInfo.Height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	gl.BindTexture(gl.TEXTURE_2D, 0)

	return texture
}

/*func Preload(path string, assetType string) {
	switch assetType {

	case "texture":

	}

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)

}*/
