package lib

import (
  "image"
  "log"
  "fmt"
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type FPoint struct{
  X float64
  Y float64
}

type GameStart struct{
  cloudNumber int
  gameStartLayers [8]*ebiten.Image
  layerPosition [8]FPoint
  menuList [3]*Menu
  keyMap *KeyMap
}

func (g *GameStart) Init(){
  g.cloudNumber = 3
  g.layerPosition  = [8]FPoint{{0, 0}, {0, 0}, {0, 230}, {0, 250}, {0, 250}, {0, 0}, {0, 0}, {0, 0}}
  g.LoadContent()
  g.keyMap = new(KeyMap) 
}

func (g *GameStart)loadBackGround() error{
  var err error
  var gameLarnerNumber = len(g.gameStartLayers)
  var cloudImageIndex int
  for i := 0; i < gameLarnerNumber; i++{
    var layerPath string 
    if i < gameLarnerNumber - g.cloudNumber{
      layerPath = fmt.Sprintf("content/backGrounds/layer%v_1.png", i)
    }else{
      layerPath = fmt.Sprintf("content/backgrounds/cloudAnimation%v.png", cloudImageIndex);
      cloudImageIndex++
    }
    g.gameStartLayers[i], _, err = ebitenutil.NewImageFromFile(layerPath)
  }
  return err
}

func (g *GameStart)loadMenu() {
  var menuName [3]string = [3]string{"START", "DEV", "EXIT"}
  for i:=0; i<len(g.menuList); i++{
    var  menu *Menu = NewMenu() 
    menu.MenuColor = COLOR_WHITE 
    menu.MenuName = menuName[i]
    menu.Position = image.Point{20, SCRREN_ORI_HEIGHT/2 + 40 + 40 * i}
    g.menuList[i] = menu
  }
}

func (g *GameStart)updateCloud(){
  var count int = 1
  var layerNumber int  = len(g.gameStartLayers)
  for i := layerNumber - g.cloudNumber; i< layerNumber; i++{
    g.layerPosition[i].X += 0.05 * float64(count) 
    count ++
    if(g.layerPosition[i].X) >= float64(SCRREN_ORI_WIDTH){
      g.layerPosition[i].X = -float64(SCRREN_ORI_WIDTH)
    }
  }
}

func(g *GameStart)LoadContent(){
  var err error = g.loadBackGround()
  if err != nil {
    log.Fatal(err)
  }
  g.loadMenu()
}

func (g *GameStart)SelectMenu(mod string){
  var menuIndex int
  var menuListLength = len(g.menuList)
  if selectedMenu == ""{
    if(mod == "next"){
      menuIndex = 0
    }else{
      menuIndex = menuListLength -1
    }
  }else{
    for index, menu := range(g.menuList){
      if menu.MenuName == selectedMenu {
        menuIndex = index
        break
      }
    }
    if(mod == "next"){
      menuIndex ++
      if menuIndex > menuListLength -1 {
        menuIndex = 0
      }
    }else{
      menuIndex -- 
      if menuIndex < 0 {
        menuIndex = menuListLength - 1
      }
    }
  }
  selectedMenu = g.menuList[menuIndex].MenuName
}

func (g *GameStart)keyEvent(){
  if g.keyMap.IsKeyUpPressed(){
    g.SelectMenu("pre")
  }else if g.keyMap.IsKeyDownPressed(){
    g.SelectMenu("next")
  }
  if g.keyMap.IsKeyEnterPressed() && selectedMenu != ""{
    GAMEMODE = selectedMenu
  }
}

func (g *GameStart ) Update(){
  g.updateCloud()
  for _, item := range(g.menuList){
    item.Update()
  }
  g.keyEvent()
}

func (g *GameStart) Draw(screen *ebiten.Image){
  for index, item := range(g.gameStartLayers){
    var iop *ebiten.DrawImageOptions = new(ebiten.DrawImageOptions)
    iop.GeoM.Translate(g.layerPosition[index].X, g.layerPosition[index].Y)
    screen.DrawImage(item, iop)
  }
  for _, menu := range(g.menuList){
    menu.Draw(screen)
  }
}

func (g *GameStart) Dispose(){
  for _, layer := range(g.gameStartLayers){
    layer.Dispose()
  }
}

func NewGameMain()*GameStart{
  var ng *GameStart = new(GameStart)
  ng.Init()
  return ng 
}
