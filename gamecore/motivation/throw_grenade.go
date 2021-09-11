package motivation 

import(
  // "fmt"
  "math"
  "log"
  "image"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
  "github.com/ChrisChou-freeman/CityHunter/gamecore/vex"
)

type ThrowGrenade struct{
  ball *ebiten.Image
  postion *image.Point
  vector *image.Point
  ground image.Point
  levelData *tool.LevelData 
  gravity int
  force int
  counter int
  rCounter int
  speed int
  hitGround bool
  rotateDegree int
  rotateSpeed int
  rebound int
  flip bool
  exploding bool
  explode bool
  explotion *vex.Explode
}

func NewThrowGrenade(position *image.Point, levelData *tool.LevelData, flip bool) *ThrowGrenade {
  tg := new(ThrowGrenade)
  tg.init(position, levelData, flip)
  return tg
}

func (tg *ThrowGrenade) init(position *image.Point, levelData *tool.LevelData, flip bool) {
  var err error
  tg.ball, _, err = ebitenutil.NewImageFromFile("content/items/grenade.png")
  if err != nil{
    log.Fatal(err)
  }
  tg.flip = flip
  tg.postion = position 
  tg.levelData = levelData
  tg.vector = &image.Point{X: 4, Y: 0}
  if tg.flip{
    tg.vector.X *= -1
  }
  tg.force = 15
  tg.gravity = 10
  tg.speed = 3
  tg.rotateDegree = 90
  tg.rotateSpeed = 3
  tg.rebound = 1
}

func (tg *ThrowGrenade) ballBottom() int {
  return tg.postion.Y + tg.ball.Bounds().Dy()
}

func (tg *ThrowGrenade) getRec() image.Rectangle {
  return image.Rect(
    tg.postion.X,
    tg.postion.Y,
    tg.postion.X + tg.ball.Bounds().Dx(),
    tg.postion.Y + tg.ball.Bounds().Dy(),
  )
}

func (tg *ThrowGrenade) grenadeExplode(){
  if tg.exploding || tg.explode{
    return
  }
  tg.exploding = true
  tg.explotion = vex.NewExplode(tg.postion)
}

func (tg *ThrowGrenade) throw(){
  tg.counter ++
  tg.vector.Y = -tg.force + tg.gravity
  tg.hitGround = tool.CollisionDetect(
    tg.getRec(),
    tg.levelData,
    tg.vector,
    tg.postion,
    true,
  )
  if tg.hitGround{
      tg.ground = *tg.postion
      if tg.rebound > 0 {
        tg.force = tg.gravity + (int(tg.gravity/3)*tg.rebound) 
        tg.rebound--
      }
  }
  if tg.force > 0 && tg.counter%tg.speed == 0{
    tg.force --
  }
  if tg.hitGround && tg.postion.Y + tg.vector.Y >= tg.ground.Y {
    if tg.force == 0 && tg.postion.Y == tg.ground.Y{
      tg.grenadeExplode()
    }
    return
  }
  newPostion := tg.postion.Add(*tg.vector)
  tg.postion = &newPostion
}

func (tg *ThrowGrenade) Update() {
  if tg.exploding{
    if tg.explotion != nil && !tg.explotion.IsDown(){
      tg.explotion.Update()
    }else{
      tg.explode = true
    }
    return
  }
  tg.throw()
}

func (tg *ThrowGrenade) rotateImage(iopt *ebiten.DrawImageOptions) float64 {
  if tg.rCounter < tg.rotateDegree{
    tg.rCounter += tg.rotateSpeed
  }
  iopt.GeoM.Rotate(float64(tg.rCounter%360) * 2 * math.Pi / 360)
  return float64(tg.rCounter) / float64(tg.rotateDegree) * float64(tg.ball.Bounds().Dy()) 
}

func (tg *ThrowGrenade) Draw(screen *ebiten.Image) {
  if tg.exploding {
    if tg.explotion != nil && !tg.explotion.IsDown(){
      tg.explotion.Draw(screen)
    }
    return
  }
  iopt := new(ebiten.DrawImageOptions) 
  var xoffsite float64
  xoffsite = tg.rotateImage(iopt)
  iopt.GeoM.Translate(float64(tg.postion.X) + xoffsite, float64(tg.postion.Y))
  screen.DrawImage(tg.ball, iopt)
}

