package player 

import(
  // "fmt"
  "log"
  "image"

  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "github.com/hajimehoshi/ebiten/v2"

  "github.com/ChrisChou-freeman/CityHunter/gamecore/texture"
  "github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
  "github.com/ChrisChou-freeman/CityHunter/gamecore/input"
)

type Player struct{
  idle *texture.AnimationSprite
  run *texture.AnimationSprite
  death *texture.AnimationSprite
  jump *texture.AnimationSprite
  position *image.Point
  vector *image.Point 
  levelData *tool.LevelData
  keymap *input.KeyMap
  action string
  gravity int
  speed int
}

func NewPlayer(position *image.Point, levelData *tool.LevelData) *Player{
  p := new(Player)
  p.init(position, levelData)
  return p
}

func(p *Player)init(position *image.Point, levelData *tool.LevelData){
  p.position = position
  p.gravity = 3
  p.speed = 2
  p.vector = &image.Point{}
  p.action = "idle"
  p.levelData = levelData
  p.keymap = new(input.KeyMap)
  p.loadContent()
}

func(p *Player)loadAnimation() error {
    idleT, _, err := ebitenutil.NewImageFromFile("content/player/Idle/sheet.png")
    if err != nil{
      return err
    }
    p.idle = texture.NewAnimationSprite(idleT, 28,36)
    return nil
}

func(p *Player)loadContent(){
  err := p.loadAnimation()
  if err != nil{
    log.Fatal(err)
  }
}

func(p *Player)keyEvent(){ 
}

func(p *Player)gravitySim(){
  tool.CollisionDetect(p.idle.GetRec(), p.levelData, p.vector, p.position)
}

func(p *Player)motivation(){
  switch p.action {
  case "idle":
    p.vector.Y = p.gravity
  case "moveRight":
    p.vector.X = p.speed
  case "moveLeft":
    p.vector.X = -p.speed
  }
}

func(p *Player)Update(){
  p.idle.Update(p.position)
  p.motivation()
  p.gravitySim()
}

func(p *Player)Draw(screen *ebiten.Image){
  p.idle.Draw(screen)
}

func(p *Player)Dispose(){
  p.idle.Dispose()
}
