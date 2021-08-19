package ui 

import(
  "github.com/hajimehoshi/ebiten/v2"

  "github.com/ChrisChou-freeman/CityHunter/gamecore/texture"
)

type Button struct{
  *texture.Sprite
}

func(b *Button)mouseHoverOnButton()bool{
  x, y := ebiten.CursorPosition()
  thisRec := b.GetRec()
  return x>=thisRec.Min.X && y>=thisRec.Min.Y && x<=thisRec.Max.X && y<=thisRec.Max.Y
}

func(b *Button)onClick(clickfuc func()){
  if !b.mouseHoverOnButton(){
    return
  }
  clickfuc()
}

func(b *Button)Update(clickfuc func()){
  b.onClick(clickfuc)
}
