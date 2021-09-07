package shape 

import(
  "image"
  "image/color"
  "math"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
)

type Circle struct{
  Radius float64 
  verticesNumber int
  Postion *tool.FPoint
  Velocity *tool.FPoint
  emptyImage *ebiten.Image
  CColor *color.RGBA
}

func NewCircle(radius float64, postion *tool.FPoint, velocity *tool.FPoint, cColor color.RGBA) *Circle{
  c := new(Circle)
  c.init(radius, postion, velocity, cColor)
  return c
}

func (c *Circle) init(radius float64, postion *tool.FPoint, velocity *tool.FPoint, cColor color.RGBA){
  c.emptyImage = ebiten.NewImage(3, 3)
  c.emptyImage.Fill(color.White)
  c.Radius = radius
  c.Postion = postion
  c.verticesNumber = 100
  c.Velocity = velocity
  c.CColor = &cColor
}

func (c *Circle) genVertices() ([]ebiten.Vertex, []uint16) {
  vs := []ebiten.Vertex{}
  for i := 0; i < c.verticesNumber+1; i++ {
    rate := float64(i) / float64(c.verticesNumber)
      vs = append(vs, ebiten.Vertex{
        DstX:   float32(float64(c.Radius)*math.Cos(2*math.Pi*rate)) +  float32(c.Postion.X),
        DstY:   float32(float64(c.Radius)*math.Sin(2*math.Pi*rate)) + float32(c.Postion.Y),
        SrcX:   0,
        SrcY:   0,
        ColorR: float32(c.CColor.R)/255.0,
        ColorG: float32(c.CColor.G)/255.0,
        ColorB: float32(c.CColor.B)/255.0,
        ColorA: float32(c.CColor.A)/255.0,
      })
  }
  indices := []uint16{}
  for i := 0; i < c.verticesNumber; i++ {
		indices = append(indices, uint16(i), uint16(i+1)%uint16(c.verticesNumber), uint16(c.verticesNumber))
	}
  return vs, indices
}

func (c *Circle) Draw(screen *ebiten.Image){
  op := &ebiten.DrawTrianglesOptions{}
	op.Address = ebiten.AddressUnsafe
  vertices, indices := c.genVertices()
  pix2Point := c.emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
  screen.DrawTriangles(vertices, indices, pix2Point, op)
}

