package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func loadImageIntoTexture(renderer *sdl.Renderer, path string) *sdl.Texture {
  texture, err := img.LoadTexture(renderer, path);

	if err != nil {
    panic("Failed to create texture")
	}

  return texture
}
