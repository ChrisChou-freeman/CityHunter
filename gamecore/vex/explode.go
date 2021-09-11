package vex

import(
  // "fmt"
  "image"
  "math/rand"
  "image/color"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/ChrisChou-freeman/CityHunter/gamecore/shape"
  "github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
)

type Explode struct{
  position *image.Point
  circles []*shape.Circle
  expandTime int
  explodeDot int
  down bool
}

func NewExplode(position *image.Point) *Explode{
  e := new(Explode)
  e.Init(position)
  return e
}

func (e *Explode) Init(position *image.Point){
  e.position = position
  e.circles = []*shape.Circle{}
  e.explodeDot = 70
  e.expandTime = 240
  e.LoadExplodeDot()
}

func (e *Explode) LoadExplodeDot(){
  for i:=0; i<e.explodeDot; i++{
    rand.Seed(int64(i))
    radius := 15.0 + float64(rand.Intn(10))
    vector := &tool.FPoint{}
    vector.X = float64(float64(rand.Intn(40))/5 - 4)
    vector.Y = float64(-rand.Intn(6))
    colorExplode := color.RGBA{162, 162, 162, 255}
    offsetX := rand.Intn(10)
    offsetY := rand.Intn(30)
    if i%2 == 0{
      offsetX *= -1
    }
    explodPosition := &tool.FPoint{X: float64(e.position.X + offsetX), Y: float64(e.position.Y + offsetY)}
    newCircle := shape.NewCircle(radius, explodPosition, vector, colorExplode)
    e.circles = append(e.circles, newCircle)
  }
}

func (e *Explode) UpdateExplode(c *shape.Circle){
  c.Postion.Y += c.Velocity.Y
  c.Postion.X += c.Velocity.X
  if e.expandTime > 0{
    c.Radius += (rand.Float64() * float64(rand.Intn(6)))
  }else{
    c.Radius -= (rand.Float64() * float64(rand.Intn(6)))
  }
  c.Velocity.Y += 0.03
  e.expandTime --
  c.CColor.A -= 2
}

func (e *Explode)IsDown()bool{
  return e.down
}

func (e *Explode) Update(){
  need_remove := []int{}
  for index, circle := range(e.circles){
    e.UpdateExplode(circle)
    if circle.Radius <= 0 {
      need_remove = append(need_remove, index)
    }
  }
  for index, cIndex := range(need_remove){
    if cIndex == len(e.circles) -1 {
      e.circles = e.circles[:cIndex]
    }else{
      e.circles = append(e.circles[:cIndex], e.circles[cIndex+1:]...)
      for i:=index+1; i<len(need_remove); i++{
        need_remove[i]--
      }
    }
  }
  if len(e.circles) == 0{
    e.down = true
  }
}

func (e *Explode) Draw(screen *ebiten.Image){
  for _, circle := range(e.circles){
    circle.Draw(screen)
  }
}
