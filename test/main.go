package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/mergenTheDev/bizmuth"
)

func main() {
	bizmuth.Init()
	window := bizmuth.createWindow(800, 600, "Deneme", glfw.False)
}
