package gamecore

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/ChrisChou-freeman/CityHunter/gamecore/input"
	"github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
	"github.com/ChrisChou-freeman/CityHunter/gamecore/ui"
)

type GameMain struct {
	cloudNumber     int
	gameStartLayers [8]*ebiten.Image
	layerPosition   [8]tool.FPoint
	menuList        [3]*ui.Menu
	keyMap          *input.KeyMap
}

func NewGameMain() *GameMain {
	ng := new(GameMain)
	ng.init()
	return ng
}

func (g *GameMain) init() {
	g.cloudNumber = 3
	g.layerPosition = [8]tool.FPoint{
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 230},
		{X: 0, Y: 250},
		{X: 0, Y: 250},
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 0}}
	g.LoadContent()
	g.keyMap = new(input.KeyMap)
}

func (g *GameMain) loadBackGround() error {
	var err error
	var gameLarnerNumber = len(g.gameStartLayers)
	var cloudImageIndex int
	for i := 0; i < gameLarnerNumber; i++ {
		var layerPath string
		if i < gameLarnerNumber-g.cloudNumber {
			layerPath = fmt.Sprintf("content/main/layer%v_1.png", i)
		} else {
			layerPath = fmt.Sprintf("content/main/cloudAnimation%v.png", cloudImageIndex)
			cloudImageIndex++
		}
		g.gameStartLayers[i], _, err = ebitenutil.NewImageFromFile(layerPath)
	}
	return err
}

func (g *GameMain) loadMenu() {
	var menuName [3]string = [3]string{tool.GAME_START, tool.GAME_DEVELOPMENT, tool.GAME_QUIT}
	for i := 0; i < len(g.menuList); i++ {
		menu := ui.NewMenu()
		menu.MenuColor = tool.COLOR_WHITE
		menu.MenuName = menuName[i]
		menu.Position = image.Point{20, tool.SCRREN_ORI_HEIGHT/2 + 40 + 40*i}
		g.menuList[i] = menu
	}
}

func (g *GameMain) LoadContent() {
	var err error = g.loadBackGround()
	if err != nil {
		log.Fatal(err)
	}
	g.loadMenu()
}

func (g *GameMain) SelectMenu(mod string) {
	var menuIndex int
	var menuListLength = len(g.menuList)
	if ui.SelectedMenu == "" {
		if mod == "next" {
			menuIndex = 0
		} else {
			menuIndex = menuListLength - 1
		}
	} else {
		for index, menu := range g.menuList {
			if menu.MenuName == ui.SelectedMenu {
				menuIndex = index
				break
			}
		}
		if mod == "next" {
			menuIndex++
			if menuIndex > menuListLength-1 {
				menuIndex = 0
			}
		} else {
			menuIndex--
			if menuIndex < 0 {
				menuIndex = menuListLength - 1
			}
		}
	}
	ui.SelectedMenu = g.menuList[menuIndex].MenuName
}

func (g *GameMain) keyEvent() {
	if g.keyMap.IsKeyUpPressed() {
		g.SelectMenu("pre")
	} else if g.keyMap.IsKeyDownPressed() {
		g.SelectMenu("next")
	}
	if g.keyMap.IsKeyEnterPressed() && ui.SelectedMenu != "" {
		tool.GAME_FUCTION = ui.SelectedMenu
	}
}

func (g *GameMain) updateCloud() {
	var count int = 1
	var layerNumber int = len(g.gameStartLayers)
	for i := layerNumber - g.cloudNumber; i < layerNumber; i++ {
		g.layerPosition[i].X += 0.1 * float64(count)
		count++
		if (g.layerPosition[i].X) >= float64(tool.SCRREN_ORI_WIDTH) {
			g.layerPosition[i].X = -float64(tool.SCRREN_ORI_WIDTH)
		}
	}
}

func (g *GameMain) Update() {
	g.updateCloud()
	for _, item := range g.menuList {
		item.Update()
	}
	g.keyEvent()
}

func (g *GameMain) Draw(screen *ebiten.Image) {
	for index, item := range g.gameStartLayers {
		var iop *ebiten.DrawImageOptions = new(ebiten.DrawImageOptions)
		iop.GeoM.Translate(g.layerPosition[index].X, g.layerPosition[index].Y)
		screen.DrawImage(item, iop)
	}
	for _, menu := range g.menuList {
		menu.Draw(screen)
	}
}

func (g *GameMain) Dispose() {
	for _, layer := range g.gameStartLayers {
		layer.Dispose()
	}
}
