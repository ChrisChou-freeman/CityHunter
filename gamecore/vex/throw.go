package vex 

import(
  "math"
  "log"
  "image"
  "image/color"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Throw struct{
  ball *ebiten.Image
  postion *image.Point
  vector *image.Point
  ground image.Point
  gravity int
  force int
  counter int
  rCounter int
  speed int
  hitGround bool
  rotateDegree int
  rebound int
}

func NewThrow() *Throw {
  ts := new(Throw)
  ts.init()
  return ts
}

func (ts *Throw) init() {
  var err error
  ts.ball, _, err = ebitenutil.NewImageFromFile("content/grenade.png")
  if err != nil{
    log.Fatal(err)
  }
  ts.counter = 0
  ts.rCounter = 0
  ts.hitGround = false
  ts.postion = &image.Point{80, 380}
  ts.vector = &image.Point{}
  ts.force = 20
  ts.gravity = 10
  ts.speed = 4
  ts.ground = image.Point{0, 420}
  ts.rotateDegree = 90
  ts.rebound = 1
}

func (ts *Throw) ballBottom() int {
  return ts.postion.Y + ts.ball.Bounds().Dy()
}

func (ts *Throw) throw(){
  ts.counter ++
  ts.vector.Y = -ts.force + ts.gravity
  ts.vector.X = 5
  ts.counter ++
  if ts.force > 0 && ts.counter%ts.speed == 0{
    ts.force --
  }
  newPostion := ts.postion.Add(*ts.vector)
  ts.postion = &newPostion
}

func (ts *Throw) Update() {
  if ts.ballBottom() >= ts.ground.Y {
    ts.postion.Y = ts.ground.Y - ts.ball.Bounds().Dy()
    ts.hitGround = true
    if ts.rebound > 0 {
      ts.postion.Y -= 1
      ts.force = 13
      ts.rebound--
    }
    return
  }
  ts.throw()
}

func (ts *Throw) rotateImage(iopt *ebiten.DrawImageOptions) float64 {
  if ts.rCounter < ts.rotateDegree{
    ts.rCounter += 3
  }
  iopt.GeoM.Rotate(float64(ts.rCounter%360) * 2 * math.Pi / 360)
  return float64(ts.rCounter) / float64(ts.rotateDegree) * float64(ts.ball.Bounds().Dy()) 
}

func (ts *Throw) Draw(screen *ebiten.Image) {
  iopt := new(ebiten.DrawImageOptions) 
  var xoffsite float64
  xoffsite = ts.rotateImage(iopt)
  iopt.GeoM.Translate(float64(ts.postion.X) + xoffsite, float64(ts.postion.Y))
  screen.DrawImage(ts.ball, iopt)
  ebitenutil.DrawLine(screen, 0, float64(ts.ground.Y), 800, float64(ts.ground.Y), color.White)
}

