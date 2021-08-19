package player 

import(
  "image"

  "github.com/ChrisChou-freeman/CityHunter/gamecore/texture"
)

type Player struct{
  idle *texture.AnimationSprite
  position *image.Point
}

func NewPlayer(position *image.Point) *Player{
  p := new(Player)
  p.init(position)
  return p
}

func(p *Player)init(position *image.Point){
  p.position = position
}

func(p *Player)Update(){

}

func(p *Player)Draw(){

}
