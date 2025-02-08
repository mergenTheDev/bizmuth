# Bizmuth
A game engine written in Go.

## Basic Usage
```go
package main

import "github.com/mergenTheDev/bizmuth"

func main() {
	bizmuth.Init()
	window := bizmuth.CreateWindow(800, 600, "Deneme", bizmuth.FALSE)
	window.SetIcon("bizmuth.png")
	bizmuth.BackgroundColor(0.5, 0.1, 0.6, 1.0)
	window.Render()
}
```

## To-Do List

- [ ] GUI
- [ ] Mobile support
- [ ] Changing OpenGL version
- [ ] Vulkan and OpenGLES support
- [ ] 3D support
