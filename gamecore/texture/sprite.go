package texture 

import(
  "image"
  "image/color"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"

  "github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
)

type Sprite struct{
  Texture *ebiten.Image
  Position *tool.FPoint
  SpriteName string
  CollisionInfo string
}

func(s *Sprite)GetRec() image.Rectangle{
  width, height := s.Texture.Size()
  rec := image.Rectangle{
    Min:image.Point{int(s.Position.X), int(s.Position.Y)},
    Max:image.Point{int(s.Position.X) + width, int(s.Position.Y) + height},
  }
  return rec 
}

func(s *Sprite)Update(){}

func(s *Sprite)DrawEdge(screen *ebiten.Image, top bool, left bool, bottom bool, right bool, lineColor color.RGBA){
    thisRec := s.GetRec()
    if top{
      // up line
      ebitenutil.DrawLine(screen, float64(thisRec.Min.X), float64(thisRec.Min.Y), float64(thisRec.Max.X), float64(thisRec.Min.Y), lineColor)
    }
    if left{
      // left line
      ebitenutil.DrawLine(screen, float64(thisRec.Min.X), float64(thisRec.Min.Y), float64(thisRec.Min.X), float64(thisRec.Max.Y), lineColor)
    }
    if bottom{
      // bottom line
      ebitenutil.DrawLine(screen, float64(thisRec.Min.X), float64(thisRec.Max.Y), float64(thisRec.Max.X), float64(thisRec.Max.Y), lineColor)
    }
    if right{
      // right line
      ebitenutil.DrawLine(screen, float64(thisRec.Max.X), float64(thisRec.Min.Y), float64(thisRec.Max.X), float64(thisRec.Max.Y), lineColor)
    }
}

func(s *Sprite)DrawCollisionVisual(screen *ebiten.Image){
  if s.CollisionInfo == ""{
    return
  }
  switch s.CollisionInfo{
  case "full":
    s.DrawEdge(screen, true, true, true, true, tool.COLOR_RED)
  }
}

func(s *Sprite)Draw(screen *ebiten.Image){
  var iop *ebiten.DrawImageOptions = new(ebiten.DrawImageOptions)
  iop.GeoM.Translate(s.Position.X, s.Position.Y)
  screen.DrawImage(s.Texture, iop)
}

func(s *Sprite)Dispose(){
  s.Texture.Dispose()
}
