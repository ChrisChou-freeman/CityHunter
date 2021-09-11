package motivation

import(
  "github.com/ChrisChou-freeman/CityHunter/gamecore/texture"
)

type Bullet struct{
  *texture.MotivationSprite
}

func NewBullet(ms *texture.MotivationSprite) *Bullet {
  b := new(Bullet)
  b.init(ms)
  return b
}

func(b *Bullet)init(ms *texture.MotivationSprite){
  b.MotivationSprite = ms
}

func(b *Bullet) Update(){
}

func(b *Bullet) Draw(){
}
