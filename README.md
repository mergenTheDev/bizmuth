# Bizmuth
A game engine written in Go.

## Basic Usage
```go
package main

import "github.com/mergenTheDev/bizmuth"

func main() {
   bizmuth.Init()
   bizmuth.CreateWindow(800, 600, "Hello", glfw.False)
   bizmuth.GameLoop()
}
```

## To-Do List

- [ ] Mobile support
- [ ] Changing OpenGL version
- [ ] Vulkan and OpenGLES support
- [ ] 3D support