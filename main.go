package main

import (
  "os"
	"log"
  "fmt"
  _ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "github.com/ChrisChou-freeman/CityHunter/lib"
)

var gameManage lib.GameManager 

func init(){
  var err error
  if err != nil {
      log.Fatal(err)
  }
}

type Game struct {}

func (g *Game) Update() error {
  switch lib.GAMEMODE{
    case "GAMEMAIN":
      gameManage = lib.NewGameMain() 
      lib.GAMEMODE = ""
    case "DEV":
      gameManage = lib.NewLevelEditor() 
      lib.GAMEMODE = ""
    case "EXIT":
      os.Exit(0)
  }
  gameManage.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  gameManage.Draw(screen)
  ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return lib.SCRREN_ORI_WIDTH, lib.SCRREN_ORI_HEIGHT 
}

func main() {
	ebiten.SetWindowSize(lib.SCRREN_WIDTH, lib.SCRREN_HEIGHT)
	ebiten.SetWindowTitle("City Hunter")
  var game *Game = new(Game) 
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
