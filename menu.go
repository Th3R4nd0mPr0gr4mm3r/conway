package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const PLAY_SELECTED_PATH = "assets/play_selected.png"
const PLAY_UNSELECTED_PATH = "assets/play_unselected.png"
const STOP_SELECTED_PATH = "assets/stop_selected.png"
const STOP_UNSELECTED_PATH = "assets/stop_unselected.png"

const MENU_HEIGHT = 50
const SLIDER_LENGTH = 400

const SLIDER_WIDTH = 10
const SLIDER_HEIGHT = 30

const MAX_FPS = 100.0

type Menu struct {
	isRunning           bool
	fps                 uint32
	stopSelected        *sdl.Texture
	stopUnselected      *sdl.Texture
	playSelected        *sdl.Texture
	playUnselected      *sdl.Texture
	isPlayButtonHovered bool
	playButtonPosition  sdl.Rect

	sliderPosition sdl.Rect

	leftSliderPosition   sdl.Rect
	rightSliderPosition  sdl.Rect
	middleSliderPosition sdl.Rect

	isSliderMoving bool
}

func NewMenu(renderer *sdl.Renderer) *Menu {
	m := new(Menu)
	m.fps = 1

	m.stopSelected = loadImageIntoTexture(renderer, STOP_SELECTED_PATH)
	m.stopUnselected = loadImageIntoTexture(renderer, STOP_UNSELECTED_PATH)
	m.playSelected = loadImageIntoTexture(renderer, PLAY_SELECTED_PATH)
	m.playUnselected = loadImageIntoTexture(renderer, PLAY_UNSELECTED_PATH)
	m.playButtonPosition = sdl.Rect{X: 10, Y: 10, W: 25, H: 25}

	sliderHeightPos := int32(MENU_HEIGHT/2 - SLIDER_HEIGHT/2)

	m.leftSliderPosition = sdl.Rect{X: 100,
		Y: sliderHeightPos, W: SLIDER_WIDTH, H: SLIDER_HEIGHT}
	m.rightSliderPosition = sdl.Rect{X: 100 + SLIDER_LENGTH, Y: sliderHeightPos,
		W: SLIDER_WIDTH, H: SLIDER_HEIGHT}
	m.middleSliderPosition = sdl.Rect{X: 100, Y: sliderHeightPos + SLIDER_HEIGHT/2 - 2, W: 5 + SLIDER_LENGTH, H: 5}

	m.sliderPosition = sdl.Rect{X: m.leftSliderPosition.X + m.leftSliderPosition.W,
		Y: m.leftSliderPosition.Y, W: SLIDER_WIDTH, H: SLIDER_HEIGHT}

	return m
}

func (m *Menu) Unload() {
	defer m.playSelected.Destroy()
	defer m.stopUnselected.Destroy()
	defer m.playSelected.Destroy()
	defer m.playUnselected.Destroy()
}

func (m *Menu) draw(sur *sdl.Surface, renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: SCREEN_WIDTH, H: SCREEN_HEIGHT})

	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.FillRect(&sdl.Rect{X: 5, Y: 5, W: 35, H: 35})

	// Draw slider
	// Left side
	renderer.FillRect(&m.leftSliderPosition)
	// Right Side
	renderer.FillRect(&m.rightSliderPosition)
	// Middle
	renderer.FillRect(&m.middleSliderPosition)

	// Moving part
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.FillRect(&m.sliderPosition)

	// Buttons logic
	src := sdl.Rect{X: 0, Y: 0, W: 25, H: 25}
	if m.isPlayButtonHovered {
		if m.isRunning {
			renderer.Copy(m.stopSelected, &src, &m.playButtonPosition)
		} else {
			renderer.Copy(m.playSelected, &src, &m.playButtonPosition)
		}
	} else {
		if m.isRunning {
			renderer.Copy(m.stopUnselected, &src, &m.playButtonPosition)
		} else {
			renderer.Copy(m.playUnselected, &src, &m.playButtonPosition)
		}
	}
}

func (m *Menu) update(x int32, y int32) {
	m.isPlayButtonHovered = isInRect(x, y, m.playButtonPosition)

	if m.isSliderMoving {
		if x < m.leftSliderPosition.X+m.leftSliderPosition.W {
			m.sliderPosition.X = m.leftSliderPosition.X + m.leftSliderPosition.W
		} else if x > m.rightSliderPosition.X-m.rightSliderPosition.W {
			m.sliderPosition.X = m.rightSliderPosition.X - m.rightSliderPosition.W
		} else {
			m.sliderPosition.X = x
		}
	}
}

func (m *Menu) click(x int32, y int32) {
	if isInRect(x, y, m.playButtonPosition) {
		m.isRunning = !m.isRunning
	}

	if isInRect(x, y, m.sliderPosition) {
		m.isSliderMoving = true
	}
}

func (m *Menu) clickStop() {
	m.isSliderMoving = false

	sliderPosRelativeToZero := float64(m.sliderPosition.X - m.leftSliderPosition.X)
	rightSliderPosRelativeToZero := float64(m.rightSliderPosition.X - m.leftSliderPosition.X)

	m.fps = uint32((sliderPosRelativeToZero / rightSliderPosRelativeToZero) * MAX_FPS)

	if m.fps == 0 {
		m.fps = 1
	}
}

func isInRect(x int32, y int32, buttonRect sdl.Rect) bool {
	return x > buttonRect.X &&
		x < buttonRect.X+buttonRect.W &&
		y > buttonRect.Y &&
		y < buttonRect.Y+buttonRect.H
}
