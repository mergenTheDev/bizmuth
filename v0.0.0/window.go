package bizmuth

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"image/draw"
	"image/png"
)

type Window struct {
	*glfw.Window
}

var (
	windowWidth  int
	windowHeight int
)

func CreateWindow(width int32, height int32, title string, resizable int) *Window {
	glfw.WindowHint(glfw.Resizable, resizable)

	window, err := glfw.CreateWindow(int(width), int(height), title, nil, nil)

	if err != nil {
		log.Fatal(PrefixErr + "Can't create window!")
	}

	window.MakeContextCurrent()

	windowWidth, windowHeight = window.GetSize()

	fmt.Println(PrefixInfo + "OpenGL version: " + gl.GoStr(gl.GetString(gl.VERSION)))

	gl.Viewport(0, 0, width, height)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	return &Window{window}
}

func (window *Window) SetIcon(path string) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(PrefixErr + "Can't open file! But it is not fatal.")
		return
	}

	defer file.Close()

	img, err := png.Decode(file)

	if err != nil {
		fmt.Println(PrefixErr + "Can't decode image! But it is not fatal.")
		return
	}

	rgbaImg := image.NewNRGBA(img.Bounds())
	draw.Draw(rgbaImg, img.Bounds(), img, image.Point{}, draw.Src)

	window.Window.SetIcon([]image.Image{rgbaImg})
}
