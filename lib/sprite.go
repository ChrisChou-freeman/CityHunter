package lib

import(
  "image"
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var InSelectSprite string

type Sprite struct{
  Texture *ebiten.Image
  SpriteName string
  Position *FPoint
}

func(s *Sprite)getRec() image.Rectangle{
  width, height := s.Texture.Size()
  var rec *image.Rectangle = new(image.Rectangle)
  rec.Min.X = int(s.Position.X) 
  rec.Min.Y = int(s.Position.Y) 
  rec.Max.X = int(s.Position.X) + width
  rec.Max.Y = int(s.Position.Y) + height
  return *rec
}

func(s *Sprite)Update(){}

func(s *Sprite)Draw(screen *ebiten.Image){
  var iop *ebiten.DrawImageOptions = new(ebiten.DrawImageOptions)
  iop.GeoM.Translate(s.Position.X, s.Position.Y)
  screen.DrawImage(s.Texture, iop)
  
  if s.SpriteName == InSelectSprite && s.SpriteName != ""{
    thisRec := s.getRec()
    ebitenutil.DrawLine(screen, float64(thisRec.Min.X), float64(thisRec.Min.Y), float64(thisRec.Max.X), float64(thisRec.Min.Y), COLOR_YELLOW)
    ebitenutil.DrawLine(screen, float64(thisRec.Min.X), float64(thisRec.Min.Y), float64(thisRec.Min.X), float64(thisRec.Max.Y), COLOR_YELLOW)
    ebitenutil.DrawLine(screen, float64(thisRec.Max.X), float64(thisRec.Min.Y), float64(thisRec.Max.X), float64(thisRec.Max.Y), COLOR_YELLOW)
    ebitenutil.DrawLine(screen, float64(thisRec.Min.X), float64(thisRec.Max.Y), float64(thisRec.Max.X), float64(thisRec.Max.Y), COLOR_YELLOW)
  }
}

func(s *Sprite)Dispose(){
  s.Texture.Dispose()
}
