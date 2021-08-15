package lib

import(
  "io"
  "os"
  "log"
  "fmt"

  "path"
  "image/color"
  "encoding/json"
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
  levelData *LevelData
  levelEditorLayerNumber  int
  levelEditorLayerRepeat int
  layerScrollSpeed int
  tileWidth int
  tileHeight int
  inSelectTile int
  srufaceLayerWidth int
  tileColNumber int
  currentLevel int
  globelScroll float64
  showGrids bool
  showTilesContainer bool
}

func NewLevelEditor() *LevelEditor{
  var le *LevelEditor = new(LevelEditor)
  le.init()
  return le
}

func (l *LevelEditor)init(){
  l.levelEditorLayerNumber  = 4
  l.levelEditorLayerRepeat = 2
  l.layerScrollSpeed = 8
  l.keymap = new(KeyMap)
  l.tileWidth = 32
  l.tileHeight = 32
  l.showGrids = false
  l.showTilesContainer = false
  l.menuContainer = FRectangle{Min:FPoint{0, 0}, Max:FPoint{float64(SCRREN_ORI_WIDTH), 100}} 
  l.levelData = NewLevelData()
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
    newLine := &FRectangle{
      Min:FPoint{0, float64(i)},
      Max:FPoint{float64(l.srufaceLayerWidth * l.levelEditorLayerRepeat),
      float64(i)},
    }
    l.lineList = append(l.lineList, newLine)
  }
  for i:=0; i<l.srufaceLayerWidth * l.levelEditorLayerRepeat; i+=l.tileWidth{
    newLine := &FRectangle{Min:FPoint{float64(i), 0}, Max:FPoint{float64(i), float64(SCRREN_ORI_HEIGHT)}}
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

func(l *LevelEditor)initialLevelData(){
  cols := l.srufaceLayerWidth * l.levelEditorLayerRepeat / l.tileWidth
  rows := SCRREN_ORI_HEIGHT / l.tileHeight
  for row := 0; row < rows; row++{
    for col := 0; col < cols; col ++{
      newMap := map[string]int{"X": col * l.tileWidth, "Y": row * l.tileHeight, "tile": -1}
      l.levelData.TileData = append(l.levelData.TileData, newMap)
    }
  }
}

func(l *LevelEditor)loadLevelData() error{ 
  levelDataPath := fmt.Sprintf("content/leveldata/%v.json", l.currentLevel)
  var isNewFile bool
  var fileObj *os.File
  var err error
  var fileState os.FileInfo
  defer fileObj.Close()
  if fileState, err = os.Stat(levelDataPath); os.IsNotExist(err){
    isNewFile = true
    fileObj, err = os.Create(levelDataPath)
    if err != nil{
      return err
    }
  }else{
    fileObj, err = os.Open(levelDataPath)
    if err != nil{
      return err
    }
  }
  if isNewFile || fileState.Size()==0{
    l.initialLevelData()
  }else{
    jsonByte, err := io.ReadAll(fileObj)
    if err != nil{
      return err
    }
    err = json.Unmarshal(jsonByte, l.levelData)
    if err != nil{
      return err
    }
  }
  return nil
}

func(l *LevelEditor)saveLevelData(){
  levelDataPath := fmt.Sprintf("content/leveldata/%v.json", l.currentLevel)
  jsonByte, err := json.Marshal(l.levelData)
  if err != nil{
    log.Fatal(err)
  }
  err = os.WriteFile(levelDataPath, jsonByte, 0644)
  if err != nil{
    log.Fatal(err)
  }
}

func (l *LevelEditor)LoadContent(){
  l.srufaceLayerWidth = l.loadLevelEditorLayers()
  l.tileColNumber = l.srufaceLayerWidth * l.levelEditorLayerRepeat / l.tileWidth
  l.LoadGridLine()
  l.loadSurfaceButton()
  l.loadTilesButton()
  err := l.loadLevelData()
  if err != nil{
    log.Fatal(err)
  }
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
      // dfLayerSpeed is deffrent layer speed
      var dfLayerSpeed float64 = float64(l.layerScrollSpeed) * (float64(j + 1) / float64(l.levelEditorLayerNumber))
      if mod == "right"{
        l.levelEdirorLayers[index].Position.X -= dfLayerSpeed 
      }else{
        l.levelEdirorLayers[index].Position.X += dfLayerSpeed 
      }
    }
  }
  for _, line := range(l.lineList){
    if mod == "right"{
      line.Min.X -= float64(l.layerScrollSpeed)
      line.Max.X -= float64(l.layerScrollSpeed)
    }else{
      line.Min.X += float64(l.layerScrollSpeed)
      line.Max.X += float64(l.layerScrollSpeed)
    }
  }
  if mod == "right"{
    l.globelScroll -= float64(l.layerScrollSpeed)
  }else{
    l.globelScroll += float64(l.layerScrollSpeed)
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

  if inpututil.IsKeyJustPressed(ebiten.KeyS){
    l.saveLevelData()
  }
}

func(l *LevelEditor)gridArea()bool{
  x, y := ebiten.CursorPosition()
  inMenuContainer := float64(x) >= l.menuContainer.Min.X && float64(y) >= l.menuContainer.Min.Y && float64(x) <= l.menuContainer.Max.X && float64(y) <= l.menuContainer.Max.Y
  tileOpenBtnRec := l.tileOpenBtn.getRec()
  inTileOpenBtn := x >= tileOpenBtnRec.Min.X && y >= tileOpenBtnRec.Min.Y && x <= tileOpenBtnRec.Max.X && y <= tileOpenBtnRec.Max.Y
  return (!inMenuContainer || !l.showTilesContainer) && !inTileOpenBtn
}

func (l *LevelEditor)mouseEvent(){
  // grid area click
  x, y := ebiten.CursorPosition()
  if x >= SCRREN_ORI_WIDTH || y >= SCRREN_ORI_HEIGHT{
    return
  }
  if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !MouseLeftInUsing && l.gridArea(){
    x -= int(l.globelScroll)
    x/=l.tileWidth
    y/=l.tileHeight
    l.levelData.TileData[y*l.tileColNumber+x]["tile"] = l.inSelectTile
  }else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && !MouseLeftInUsing && l.gridArea(){
    x -= int(l.globelScroll)
    x/=l.tileWidth
    y/=l.tileHeight
    l.levelData.TileData[y*l.tileColNumber+x]["tile"] = -1
  }
}

func(l *LevelEditor) Update(){
  if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft){
    MouseLeftInUsing = false
  }
  l.keyEvent()
  l.tileOpenBtn.Update()
  for _, tile := range(l.tileBtnList){
    tile.Update()
  }
  l.mouseEvent()
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

func(l *LevelEditor)getTileCollisionInfo(tile int)string{
  if fullCollisionList, ok := l.levelData.CollisionData["full"]; ok{
    if SliceContainItem(fullCollisionList, tile) != -1{
      return "full"
    }
  }
  return ""
}

func(l *LevelEditor)DrawTiles(screen *ebiten.Image){
  for _, tile := range(l.levelData.TileData){
    if tile["tile"] == -1{
      continue
    }
    var textTure *ebiten.Image
    for _, item := range(l.tileBtnList){
      if item.SpriteName == fmt.Sprintf("%v.png", tile["tile"]){
        textTure = item.Texture
        break
      }
    }
    tilex := float64(tile["X"]) + l.globelScroll
    newSprite := &Sprite{
      Texture: textTure,
      Position: &FPoint{float64(tilex), float64(tile["Y"])},
      CollisionInfo: l.getTileCollisionInfo(tile["tile"])}
    newSprite.Draw(screen)
    if l.showGrids{
      newSprite.DrawCollisionVisual(screen)
    }
  }
}

func(l *LevelEditor) Draw(scrren *ebiten.Image){
  l.DrawLayers(scrren)
  l.DrawGrids(scrren)
  l.DrawTilesContainer(scrren)
  l.DrawButton(scrren)
  l.DrawTiles(scrren)
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

