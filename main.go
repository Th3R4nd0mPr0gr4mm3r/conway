package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const SCREEN_WIDTH = 1000
const SCREEN_HEIGHT = 1000

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Conway's game of life", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		SCREEN_WIDTH, SCREEN_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	running := true

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	conway := conway{}
	menu := NewMenu(renderer)
  defer menu.Unload()

	var oldTimeStamp uint32 = 0

	for running {
		menu.draw(surface, renderer)
		conway.draw(renderer)

		renderer.Present()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyboardEvent:
				keyCode := t.Keysym.Sym

        if t.Type == sdl.KEYUP {
          if keyCode == sdl.K_c {
            conway.clear()
          } else if keyCode == sdl.K_r {
            conway.random()
          }
        }

				break
			case *sdl.MouseMotionEvent:
				menu.update(t.X, t.Y)
				break
			case *sdl.MouseButtonEvent:
				if t.State == sdl.PRESSED {
					conway.click(t.X, t.Y)
					menu.click(t.X, t.Y)
				} else if t.State == sdl.RELEASED {
          menu.clickStop()
        }
				break
			}
		}

		if menu.isRunning {
			currentTimeStamp := uint32(time.Now().UnixMilli())
			var fpsTime uint32 = 1000 / menu.fps

			if currentTimeStamp-oldTimeStamp > fpsTime {
				conway.nextFrame()
				oldTimeStamp = currentTimeStamp
			}
		}

		sdl.Delay(16)
	}
}
