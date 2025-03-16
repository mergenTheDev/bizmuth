# Bizmuth
A game engine written in Go.

[![Go Reference](https://pkg.go.dev/badge/github.com/mergenTheDev/bizmuth.svg)](https://pkg.go.dev/github.com/mergenTheDev/bizmuth)

Easy to use and lightweight.

`Based on go-gl/gl and go-gl/glfw.`

# Still Under Development! Currently Does Not Work!

## Basic Usage
```go get -u github.com/mergenTheDev/bizmuth```

```go
package main

import "github.com/mergenTheDev/bizmuth"

func main() {
	bizmuth.Init()

	window := bizmuth.CreateWindow(800, 600, "Deneme", bizmuth.FALSE)
	window.SetIcon("bizmuth.png")

	bizmuth.BackgroundColor(0.5, 0.1, 0.6, 1.0)

	camera := bizmuth.CreateCamera(0, 0)

	bizmuth.GameLoop(func(){
		camera.Update()
  		//Do Something.
        })

	defer bizmuth.End()
}
```

## To-Do List

- [ ] Physics
- [ ] Lights and Shadows
- [ ] Normal Textures
- [ ] Controller Support

## Future To-Do

- [ ] GUI (Maybe)
- [ ] Mobile support
- [ ] Vulkan and OpenGLES support
- [ ] 3D
