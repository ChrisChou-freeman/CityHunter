package texture 

import(
  // "fmt"
  "image"

  "github.com/hajimehoshi/ebiten/v2"
)

type AnimationSprite struct{
  position *image.Point
  texture *ebiten.Image
  framNum int
  framWidth int
  framHeight int
  currentFrame int
  count int
  loop bool
  AnimationPlayEnd bool
  Flip bool
}

func NewAnimationSprite(texture *ebiten.Image, framWidth int, framHeight int) *AnimationSprite{
  as := new(AnimationSprite)
  as.init(texture, framWidth, framHeight)
  return as
}

func(as *AnimationSprite)init(texture *ebiten.Image, framWidth int, framHeight int){
  as.texture = texture
  as.framWidth = framWidth
  as.framHeight = framHeight
  textureWidth, _ := texture.Size()
  as.framNum = textureWidth / framWidth
}

func(as *AnimationSprite)GetRec() image.Rectangle{
  rec := image.Rectangle{
    Min:image.Point{int(as.position.X), int(as.position.Y)},
    Max:image.Point{int(as.position.X) + as.framWidth, int(as.position.Y) + as.framHeight},
  }
  return rec 
}

func(as *AnimationSprite)Play(){
  as.loop = true 
}

func(as *AnimationSprite)PlayOnce(){
  as.loop = false
  as.currentFrame = 0
  as.count = 0
}

func(as *AnimationSprite)Update(position *image.Point){
  as.position = position
  if as.loop{
    as.count++
  }else{
    if as.currentFrame < as.framNum - 1{
      as.count++
    }else{
      as.AnimationPlayEnd = true
    }
  }
  as.currentFrame = (as.count/6) % as.framNum 
}

func(as *AnimationSprite)Draw(screen *ebiten.Image){
  iop := new(ebiten.DrawImageOptions)
  drawPositonX := as.position.X
  if as.Flip{
    iop.GeoM.Scale(-1, 1)
    drawPositonX += as.framWidth
  }
  iop.GeoM.Translate(float64(drawPositonX), float64(as.position.Y))
  framX := as.currentFrame * as.framWidth
  screen.DrawImage(as.texture.SubImage(image.Rect(framX, 0, framX + as.framWidth, as.framHeight)).(*ebiten.Image), iop)
}

func(as *AnimationSprite)Dispose(){
  as.texture.Dispose()
}

