# Bizmuth
A game engine written in Go.

Easy to use and lightweight.

`Based on go-gl/gl and go-gl/glfw.`

# Still Under Development! Currently Does Not Work!

## Basic Usage
```go
package main

import "github.com/mergenTheDev/bizmuth"

func main() {
	bizmuth.Init()

	window := bizmuth.CreateWindow(800, 600, "Deneme", bizmuth.FALSE)
	window.SetIcon("bizmuth.png")

	bizmuth.BackgroundColor(0.5, 0.1, 0.6, 1.0)

	bizmuth.GameLoop(func(){
  
        })

	defer bizmuth.End()
}
```

## I'm Currently Work On

- [ ] Render Engine
- [ ] Physics
- [ ] Things that makes a game engine

## To-Do List

- [ ] GUI
- [ ] Mobile support
- [ ] Be able to change OpenGL version
- [ ] Vulkan and OpenGLES support
- [ ] 3D support
