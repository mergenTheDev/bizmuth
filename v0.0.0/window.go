package bizmuth

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	_ "image/png"
)

func CreateWindow(width int, height int, title string, resizable int) *glfw.Window {
	glfw.WindowHint(glfw.Resizable, resizable)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)

	if err != nil {
		log.Fatal(PrefixErr + "Can't create window!")
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatal(PrefixErr + "Can't initialize OpenGL!")
	}

	//something useless

	/*SetIcon := func() {
		f, err := os.Open("bizmuth.png")

		if err != nil {
			panic(err)
		}

		var fi []image.Image
		img, format, err := image.Decode(f)

		if format != "png" {
			panic(PrefixErr + "Image is not in PNG format It's a " + format)
		}

		if err != nil {
			panic(err)
		}
		fmt.Println(format)
		fi = append(fi, img)

		window.SetIcon(fi)
	}*/
	//useless thing end here

	gl.Viewport(0, 0, int32(width), int32(height))

	return window
}

func SetIcon() {
	f, err := os.Open("bizmuth.png")

	if err != nil {
		panic(err)
	}

	var fi []image.Image
	img, format, err := image.Decode(f)

	if format != "png" {
		panic(PrefixErr + "Image is not in PNG format It's a " + format)
	}

	if err != nil {
		panic(err)
	}
	fmt.Println(format)
	fi = append(fi, img)

	//window.SetIcon(fi)
}
