package lib

import(
  "log"
  "fmt"
  "image/color"
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/inpututil"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type LevelEditor struct{
  buttonList []*Button
  levelEdirorLayers []*Sprite
  levelEditorLayerNumber  int
  levelEditorLayerRepeat int
  layerScrrolSpeed int
  keymap *KeyMap
  tileWidth int
  tileHeight int
  srufaceLayerWidth int
  lineList []*FRectangle
  showGrids bool
}

func (l *LevelEditor)init(){
  l.levelEditorLayerNumber  = 4
  l.levelEditorLayerRepeat = 2
  l.layerScrrolSpeed = 8
  l.keymap = new(KeyMap)
  l.tileWidth = 32
  l.tileHeight = 32
  l.LoadContent()
  l.showGrids = false
}

func (l *LevelEditor)loadLevelEditorLayers() int {
  var lastLayerWidth int
  for i := 0; i<l.levelEditorLayerRepeat; i++{
    for j := 0; j<l.levelEditorLayerNumber; j++{
      var layerPath string = fmt.Sprintf("content/backgrounds/layer%v_0.png", j) 
      var layerImage *ebiten.Image
      var err error
      layerImage, _, err = ebitenutil.NewImageFromFile(layerPath) 
      if err != nil{
        log.Fatal(err)
      }
      var layWidth int
      layWidth, _ = layerImage.Size()
      var laySprite *Sprite = new(Sprite)
      newP := new(Positon)
      newP.X = float64(i *layWidth) 
      newP.Y = float64(j*80) 
      laySprite.Position = newP 
      laySprite.Texture = layerImage
      l.levelEdirorLayers = append(l.levelEdirorLayers, laySprite)
      if j == l.levelEditorLayerNumber -1{
        lastLayerWidth = layWidth
      }
    }
  }
  return lastLayerWidth
}

func (l *LevelEditor)LoadGridLine(){
  for i:=0; i<SCRREN_ORI_HEIGHT; i+=l.tileHeight{
    newLine := new(FRectangle)
    newLine.Min.X = 0
    newLine.Min.Y = float64(i)
    newLine.Max.X = float64(l.srufaceLayerWidth * l.levelEditorLayerRepeat)
    newLine.Max.Y = float64(i)
    l.lineList = append(l.lineList, newLine)
  }
  for i:=0; i<l.srufaceLayerWidth * l.levelEditorLayerRepeat; i+=l.tileWidth{
    newLine := new(FRectangle)
    newLine.Min.X = float64(i)
    newLine.Min.Y = 0
    newLine.Max.X = float64(i)
    newLine.Max.Y = float64(SCRREN_ORI_HEIGHT)
    l.lineList = append(l.lineList, newLine)
  }
}

func (l *LevelEditor)LoadContent(){
  l.srufaceLayerWidth = l.loadLevelEditorLayers()
  l.LoadGridLine()
}

func (l *LevelEditor)handleLayerScroll(mod string){
  var leayerLength int = len(l.levelEdirorLayers)
  if mod == "right"{
      if l.levelEdirorLayers[leayerLength-1].getRec().Max.X <= SCRREN_ORI_WIDTH{
        return
      }
  }else{
      if l.levelEdirorLayers[l.levelEditorLayerNumber-1].getRec().Min.X >= 0{
        return
      }
  }
  for i := 0; i<l.levelEditorLayerRepeat; i++ {
    for j := 0; j<l.levelEditorLayerNumber; j++{
      var index int = i * l.levelEditorLayerNumber + j
       var dfLayerSpeed float64 = float64(l.layerScrrolSpeed) * (float64(j + 1) / float64(l.levelEditorLayerNumber))
      if mod == "right"{
        l.levelEdirorLayers[index].Position.X -= dfLayerSpeed 
      }else{
        l.levelEdirorLayers[index].Position.X += dfLayerSpeed 
      }
    }
  }
  for _, line := range(l.lineList){
    if mod == "right"{
      line.Min.X -= float64(l.layerScrrolSpeed)
      line.Max.X -= float64(l.layerScrrolSpeed)
    }else{
      line.Min.X += float64(l.layerScrrolSpeed)
      line.Max.X += float64(l.layerScrrolSpeed)
    }
  }
}

func (l *LevelEditor)keyEvent(){
  if l.keymap.IsKeyBackPressed(){
    GAMEMODE = "GAMEMAIN"
  }

  if l.keymap.IsKeyRightHoldPressed(){
    l.handleLayerScroll("right")
  }else if l.keymap.IsKeyLeftHoldPressed(){
    l.handleLayerScroll("left")
  }

  if inpututil.IsKeyJustPressed(ebiten.KeyG){
    if l.showGrids{
      l.showGrids = false
    }else{
      l.showGrids = true
    }
  }
}

func(l *LevelEditor) Update(){
  l.keyEvent()
}

func(l *LevelEditor)DrawLayers(scrren *ebiten.Image){
  for _,layer := range(l.levelEdirorLayers){
    var iop *ebiten.DrawImageOptions = new(ebiten.DrawImageOptions)
    iop.GeoM.Translate(layer.Position.X, layer.Position.Y)
    scrren.DrawImage(layer.Texture, iop)
  }
}

func(l *LevelEditor)DrawGrids(screen *ebiten.Image){
  if !l.showGrids{
    return
  }
  for _, item := range(l.lineList){
    ebitenutil.DrawLine(screen, item.Min.X, item.Min.Y, item.Max.X, item.Max.Y, color.White)
  }
}

func(l *LevelEditor) Draw(scrren *ebiten.Image){
  l.DrawLayers(scrren)
  l.DrawGrids(scrren)
}

func(l *LevelEditor) Dispose(){
  for _, item := range(l.levelEdirorLayers){
    item.Dispose()
  }
}

func NewLevelEditor() *LevelEditor{
  var le *LevelEditor = new(LevelEditor)
  le.init()
  return le
}
