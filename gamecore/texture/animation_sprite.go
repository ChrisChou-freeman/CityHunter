package texture 

import(

  "image"
  // "fmt"

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
  textWidth, _ := texture.Size()
  as.framNum = textWidth / framWidth
}

func(s *AnimationSprite)GetRec() image.Rectangle{
  // width, height := s.texture.Size()

  rec := image.Rectangle{
    Min:image.Point{int(s.position.X), int(s.position.Y)},
    Max:image.Point{int(s.position.X) + s.framWidth, int(s.position.Y) + s.framHeight},
  }
  return rec 
}

func(as *AnimationSprite)Update(position *image.Point){
  as.position = position
  as.count++
  as.currentFrame = (as.count/6) % as.framNum 
}

func(as *AnimationSprite)Draw(screen *ebiten.Image){
  iop := new(ebiten.DrawImageOptions)
  iop.GeoM.Translate(float64(as.position.X), float64(as.position.Y))
  framX := as.currentFrame * as.framWidth
  screen.DrawImage(as.texture.SubImage(image.Rect(framX, 0, framX + as.framWidth, as.framHeight)).(*ebiten.Image), iop)
}

func(as *AnimationSprite)Dispose(){
  as.texture.Dispose()
}

