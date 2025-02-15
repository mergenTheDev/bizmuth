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

type Window struct {
	*glfw.Window
}

func CreateWindow(width int, height int, title string, resizable int) *Window {
	glfw.WindowHint(glfw.Resizable, resizable)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)

	if err != nil {
		log.Fatal(PrefixErr + "Can't create window!")
	}

	window.MakeContextCurrent()

	fmt.Println(PrefixInfo + "OpenGL version: " + gl.GoStr(gl.GetString(gl.VERSION)))

	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	return &Window{window}
}

func (window *Window) SetIcon(path string) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(PrefixErr + "Can't open file!")
	}

	defer file.Close()

	img, format, err := image.Decode(file)

	if format != "png" {
		fmt.Println(PrefixErr + "Image is not in PNG format It's a " + format)
	}

	if err != nil {
		fmt.Println(PrefixErr + "Can't decode image!")
	}

	var icon []image.Image
	icon = append(icon, img)

	window.Window.SetIcon(icon)
}
