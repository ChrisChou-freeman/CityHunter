package lib

import(
  "image"
  "github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct{
  Texture *ebiten.Image
  Position image.Point 
}

func(s *Sprite)getRec() image.Rectangle{
  var width, height int = s.Texture.Size()
  var rec *image.Rectangle = new(image.Rectangle)
  rec.Min.X = s.Position.X
  rec.Min.Y = s.Position.Y
  rec.Max.X = s.Position.X + width
  rec.Max.Y = s.Position.Y + height
  return *rec
}

func(s *Sprite)Update(){}

func(s *Sprite)Draw(screen *ebiten.Image){
  var iop *ebiten.DrawImageOptions = new(ebiten.DrawImageOptions)
  iop.GeoM.Translate(float64(s.Position.X), float64(s.Position.Y))
  screen.DrawImage(s.Texture, iop)
}

func(s *Sprite)Dispose(){
  s.Texture.Dispose()
}
