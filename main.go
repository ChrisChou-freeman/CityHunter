package main

import (
	"fmt"
	_ "image/png"
	"log"
	"os"

	"github.com/ChrisChou-freeman/CityHunter/gamecore"
	"github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var gameManage gamecore.GameManager

type Game struct{}

func (g *Game) Update() error {
	switch tool.GAME_FUCTION {
	case tool.GAME_MAIN:
		if gameManage != nil {
			gameManage.Dispose()
		}
		gameManage = gamecore.NewGameMain()
		tool.GAME_FUCTION = ""

	case tool.GAME_START:
		if gameManage != nil {
			gameManage.Dispose()
		}
		gameManage = gamecore.NewGameStart()
		tool.GAME_FUCTION = ""

	case tool.GAME_DEVELOPMENT:
		if gameManage != nil {
			gameManage.Dispose()
		}
		gameManage = gamecore.NewLevelEditor()
		tool.GAME_FUCTION = ""

	case tool.GAME_QUIT:
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
	game := new(Game)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
