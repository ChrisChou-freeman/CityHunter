package lib

import(
  "log"
  "fmt"
  "path"
  "image/color"
  "io/ioutil"
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/inpututil"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type LevelEditor struct{
  levelEdirorLayers []*Sprite
  tileOpenBtn *Button
  tileBtnList []*Button
  keymap *KeyMap
  lineList []*FRectangle
  menuContainer FRectangle
  levelEditorLayerNumber  int
  levelEditorLayerRepeat int
  layerScrrolSpeed int
  tileWidth int
  tileHeight int
  inSelectTile int
  srufaceLayerWidth int
  showGrids bool
  showTilesContainer bool
  mouseLeftInUse bool
}

func (l *LevelEditor)init(){
  l.levelEditorLayerNumber  = 4
  l.levelEditorLayerRepeat = 2
  l.layerScrrolSpeed = 8
  l.keymap = new(KeyMap)
  l.tileWidth = 32
  l.tileHeight = 32
  l.showGrids = false
  l.showTilesContainer = false
  l.menuContainer = FRectangle{Min:FPoint{0, 0}, Max:FPoint{float64(SCRREN_ORI_WIDTH), 100}} 
  l.LoadContent()
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
      layWidth, _ := layerImage.Size()
      laySprite := &Sprite{Position:&FPoint{float64(i*layWidth), float64(j*80)}, Texture:layerImage}
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
    newLine := &FRectangle{Min:FPoint{0, float64(i)}, Max:FPoint{float64(l.srufaceLayerWidth * l.levelEditorLayerRepeat), float64(i)}}
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

func (l *LevelEditor)loadSurfaceButton(){
  newTexture, _, err := ebitenutil.NewImageFromFile("content/menu/gridMenu.png")
  if err != nil{
    log.Fatal(err)
  }
  newSprite := &Sprite{Texture:newTexture, Position:&FPoint{10, 10}}
  newOpenTileBtn := &Button{buttonType:"tileOpenBtn", levelEditor:l, Sprite:newSprite, keyMap:l.keymap} 
  l.tileOpenBtn = newOpenTileBtn
}

func(l *LevelEditor)loadTilesButton(){
  baseFileDir := "content/tiles/"
  files, err := ioutil.ReadDir(baseFileDir)
  if err != nil{
    log.Fatal(err)
  }

  tileWidthSpace := 10
  cols := SCRREN_ORI_WIDTH / (l.tileWidth + tileWidthSpace)
  row := len(files) / cols
  if len(files) % cols > 0 {
    row ++
  }

  var currentRow int
  var currentCol int
  for _, file := range(files){
    filePath := path.Join(baseFileDir, file.Name())
    newTexture, _, err := ebitenutil.NewImageFromFile(filePath)
    if err != nil {
      log.Fatal(err)
    }
    newPosition := &FPoint{float64(currentCol * l.tileWidth + currentCol* tileWidthSpace), float64(currentRow * l.tileHeight)}
    newSprite := &Sprite{Texture:newTexture, Position:newPosition, SpriteName:file.Name()}
    currentCol++
    if currentCol == cols{
      currentRow ++
      currentCol = 0
    }
    newButton := &Button{Sprite:newSprite, levelEditor:l, buttonType:"tile", keyMap:l.keymap}
    l.tileBtnList = append(l.tileBtnList, newButton)
  }
}

func (l *LevelEditor)LoadContent(){
  l.srufaceLayerWidth = l.loadLevelEditorLayers()
  l.LoadGridLine()
  l.loadSurfaceButton()
  l.loadTilesButton()
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
  if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft){
    l.mouseLeftInUse = false
  }
  l.keyEvent()
  l.tileOpenBtn.Update()
  for _, tile := range(l.tileBtnList){
    tile.Update()
  }
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

func(l *LevelEditor)DrawTilesContainer(screen *ebiten.Image){
  if !l.showTilesContainer {
    return
  }
  ebitenutil.DrawRect(screen, l.menuContainer.Min.X, l.menuContainer.Min.Y, l.menuContainer.Max.X, l.menuContainer.Max.Y, COLOR_GREY)
  for _, btn := range(l.tileBtnList){
    btn.Draw(screen)
  }
}

func(l *LevelEditor)DrawButton(screen *ebiten.Image){
  l.tileOpenBtn.Draw(screen)
}

func(l *LevelEditor) Draw(scrren *ebiten.Image){
  l.DrawLayers(scrren)
  l.DrawGrids(scrren)
  l.DrawTilesContainer(scrren)
  l.DrawButton(scrren)
}

func(l *LevelEditor) Dispose(){
  for _, item := range(l.levelEdirorLayers){
    item.Dispose()
  }
  for _, tile := range(l.tileBtnList){
    tile.Dispose()
  }
  l.tileOpenBtn.Dispose()
}

func NewLevelEditor() *LevelEditor{
  var le *LevelEditor = new(LevelEditor)
  le.init()
  return le
}
