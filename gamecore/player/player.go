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
  isJumping bool
  jumpForce int
  counter int
}

func NewPlayer(position *image.Point, levelData *tool.LevelData) *Player{
  p := new(Player)
  p.init(position, levelData)
  return p
}

func(p *Player)init(position *image.Point, levelData *tool.LevelData){
  p.position = position
  p.gravity = 6
  p.speed = 3
  p.jumpForce = 80
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
  if p.keymap.IsKeyLeftHoldPressed(){
    p.action = "moveLeft"
  }else if p.keymap.IsKeyRightHoldPressed(){
    p.action = "moveRight"
  }else{
    p.action = "idle"
  }

  if p.keymap.IsKeyJumpPressed() && p.isOnGround(){
    p.isJumping = true
  }
}

func(p *Player)isOnGround()bool{
  return p.vector.Y == 0
}

func(p *Player)motivation(){
  switch p.action {
  case "idle":
    p.vector.Y = p.gravity
    p.vector.X = 0
  case "moveRight":
    p.vector.X = p.speed
    p.vector.Y = p.gravity
  case "moveLeft":
    p.vector.X = -p.speed
    p.vector.Y = p.gravity
  }

  if p.isJumping{
    p.vector.Y = -p.speed * 2
  }

}

func(p *Player)move(){
  tool.CollisionDetect(p.idle.GetRec(), p.levelData, p.vector, p.position, p.gravity)
  p.position.X += p.vector.X
  p.position.Y += p.vector.Y
  if p.vector.Y < 0 {
    p.counter -= p.vector.Y
    if p.counter >= p.jumpForce{
      p.isJumping = false
      p.vector.Y = p.speed
      p.counter = 0
    }
  }
}

func(p *Player)Update(){
  p.idle.Update(p.position)
  p.keyEvent()
  p.motivation()
  p.move()
}

func(p *Player)Draw(screen *ebiten.Image){
  p.idle.Draw(screen)
}

func(p *Player)Dispose(){
  p.idle.Dispose()
}
