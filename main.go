package main

import (
  "os"
	"log"
  "fmt"
  _ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "github.com/ChrisChou-freeman/CityHunter/gamecore"
  "github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
)

var gameManage gamecore.GameManager 

type Game struct {}

func (g *Game) Update() error {
  switch tool.GAME_FUCTION{
  case "GAMEMAIN":
    if gameManage != nil{
      gameManage.Dispose()
    }
    gameManage = gamecore.NewGameMain() 
    tool.GAME_FUCTION = ""

  case "START":
    if gameManage != nil{
      gameManage.Dispose()
    }
    gameManage = gamecore.NewGameStart()
    tool.GAME_FUCTION = ""

  case "DEV":
    if gameManage != nil{
      gameManage.Dispose()
    }
    gameManage = gamecore.NewLevelEditor() 
    tool.GAME_FUCTION = ""

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
	return tool.SCRREN_ORI_WIDTH, tool.SCRREN_ORI_HEIGHT 
}

func main() {
	ebiten.SetWindowSize(tool.SCRREN_WIDTH, tool.SCRREN_HEIGHT)
	ebiten.SetWindowTitle("City Hunter")
  var game *Game = new(Game) 
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
