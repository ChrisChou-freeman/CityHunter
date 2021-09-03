package ui

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ChrisChou-freeman/CityHunter/gamecore/texture"
	"github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
)

var InSelectButton string

type Button struct {
	*texture.Sprite
}

func (b *Button) mouseHoverOnButton() bool {
	x, y := ebiten.CursorPosition()
	thisRec := b.GetRec()
	return x >= thisRec.Min.X && y >= thisRec.Min.Y && x <= thisRec.Max.X && y <= thisRec.Max.Y
}

func (b *Button) DrawSelectedBox(screen *ebiten.Image) {
	if b.SpriteName == InSelectButton && b.SpriteName != "" {
		b.DrawEdge(screen, true, true, true, true, tool.COLOR_YELLOW)
	}
}

func (b *Button) onClick(clickfuc func()) {
	if !b.mouseHoverOnButton() {
		return
	}
	clickfuc()
}

func (b *Button) Update(clickfuc func()) {
	b.onClick(clickfuc)
}

func (b *Button) Draw(screen *ebiten.Image) {
	b.Sprite.Draw(screen)
	b.DrawSelectedBox(screen)
}
