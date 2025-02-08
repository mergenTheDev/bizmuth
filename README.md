![image](bizmuth.png "Bizmuth") # Bizmuth
A game engine written in Go.

## Basic Usage
```go
package main

import "github.com/mergenTheDev/bizmuth"

func main() {
   bizmuth.Init()
   window := bizmuth.CreateWindow(800, 600, "Hello", bizmuth.FALSE)
   bizmuth.Render(window)
}
```

## To-Do List

- [ ] GUI
- [ ] Mobile support
- [ ] Changing OpenGL version
- [ ] Vulkan and OpenGLES support
- [ ] 3D support
