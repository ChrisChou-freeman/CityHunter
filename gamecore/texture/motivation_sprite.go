package texture

import(
  "image"

  "github.com/hajimehoshi/ebiten/v2"

  "github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
)

type MotivationSprite struct{
  *Sprite
  vector *image.Point
  lifecycle int
  endLife bool
  counter int
}

func NewMotivationSprite(sprite *Sprite, lifeCircle int, vector *image.Point) *MotivationSprite {
  ms := new(MotivationSprite)
  ms.init(sprite, lifeCircle, vector)
  return ms
}

func(ms MotivationSprite)init(sprite *Sprite, lifeCircle int, vector *image.Point){
  ms.Sprite = sprite
  ms.lifecycle = lifeCircle
  ms.vector = vector
}

func(ms MotivationSprite)Update(){
  if ms.endLife{
    return
  }
  ms.Position.Add(tool.FPoint{X: float64(ms.vector.X), Y: float64(ms.vector.Y)})
  ms.counter++
  if ms.counter >= ms.lifecycle{
    ms.endLife = true
  }
}

func(ms MotivationSprite)Islife()bool{
  return ms.endLife
}

func(ms MotivationSprite)Draw(screen *ebiten.Image){
  if !ms.endLife{
    ms.Sprite.Draw(screen)
  }
}
