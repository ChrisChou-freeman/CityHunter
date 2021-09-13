package motivation

import (
	"github.com/ChrisChou-freeman/CityHunter/gamecore/texture"
	"github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

type Bullet struct {
	*texture.MotivationSprite
	istance  int
	color    color.RGBA
	trakeLen int
}

func NewBullet(position *tool.FPoint, color *color.RGBA, vector *image.Point) *Bullet {
	b := new(Bullet)
	b.init(position, *color, vector)
	return b
}

func (b *Bullet) init(position *tool.FPoint, color color.RGBA, vector *image.Point) {
	b.trakeLen = 100
	b.color = color
	newBuletT := ebiten.NewImage(7, 7)
	newBuletT.Fill(color)
	newBulet := &texture.Sprite{
		Texture:  newBuletT,
		Position: position,
	}
	b.MotivationSprite = texture.NewMotivationSprite(newBulet, 120, vector)
}

func (b *Bullet) Update() {
}

func (b *Bullet) Draw() {
}
