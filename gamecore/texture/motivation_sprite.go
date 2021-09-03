package texture

import (
	"image"
	// "fmt"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
)

type MotivationSprite struct {
	*Sprite
	Vector    *image.Point
	lifecycle int
	counter   int
	endLife   bool
}

func NewMotivationSprite(sprite *Sprite, lifeCircle int, vector *image.Point) *MotivationSprite {
	ms := new(MotivationSprite)
	ms.init(sprite, lifeCircle, vector)
	return ms
}

func (ms *MotivationSprite) init(sprite *Sprite, lifeCircle int, vector *image.Point) {
	ms.Sprite = sprite
	ms.lifecycle = lifeCircle
	ms.Vector = vector
}

func (ms *MotivationSprite) Update() {
	if ms.endLife {
		return
	}
	ms.Position.Add(tool.FPoint{X: float64(ms.Vector.X), Y: float64(ms.Vector.Y)})
	ms.counter++
	if ms.counter >= ms.lifecycle {
		ms.endLife = true
	}
}

func (ms *MotivationSprite) Islife() bool {
	return !ms.endLife
}

func (ms *MotivationSprite) Kill() {
	ms.endLife = true
}

func (ms *MotivationSprite) Draw(screen *ebiten.Image) {
	if !ms.endLife {
		ms.Sprite.Draw(screen)
	}
}
