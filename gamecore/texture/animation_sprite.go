package texture 

import(
  "image"
  "github.com/hajimehoshi/ebiten/v2"
)

type AnimationSprite struct{
  Texture *ebiten.Image
  Position *image.Point
  framNum int
  framWidth int
  framHeight int
}

func(as *AnimationSprite)Update(){
}

func(as *AnimationSprite)Draw(){
}

func(as *AnimationSprite)Dispose(){
  as.Texture.Dispose()
}

