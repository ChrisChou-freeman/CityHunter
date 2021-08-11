package lib

import(
  "github.com/hajimehoshi/ebiten/v2"
)

type GameManager interface{
  Update()
  Draw(scrren *ebiten.Image)
  Dispose()
}
