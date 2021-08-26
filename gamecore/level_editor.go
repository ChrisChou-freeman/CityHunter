package gamecore

import(
  "io"
  "os"
  "log"
  "fmt"
  "path"
  "strconv"
  "strings"
  "image/color"
  "encoding/json"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/inpututil"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"

  "github.com/ChrisChou-freeman/CityHunter/gamecore/input"
  "github.com/ChrisChou-freeman/CityHunter/gamecore/texture"
  "github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
  "github.com/ChrisChou-freeman/CityHunter/gamecore/ui"
)

type LevelEditor struct{
  levelEdirorLayers []*texture.Sprite
  tileOpenBtn *ui.Button
  tileBtnList []*ui.Button
  keymap *input.KeyMap
  lineList []*tool.FRectangle
  menuContainer tool.FRectangle
  levelData *tool.LevelData
  levelEditorLayerNumber  int
  levelEditorLayerRepeat int
  layerScrollSpeed int
  inSelectTile int
  srufaceLayerWidth int
  tileColNumber int
  currentLevel int
  globelScroll float64
  showGrids bool
  showTilesContainer bool
}

func NewLevelEditor() *LevelEditor{
  le := new(LevelEditor)
  le.init()
  return le
}

func (l *LevelEditor)init(){
  l.levelEditorLayerNumber  = 4
  l.levelEditorLayerRepeat = 2
  l.layerScrollSpeed = 8
  l.keymap = new(input.KeyMap)
  l.showGrids = false
  l.showTilesContainer = false
  l.menuContainer = tool.FRectangle{Min:tool.FPoint{X:0, Y:0}, Max:tool.FPoint{X:float64(tool.SCRREN_ORI_WIDTH), Y:100}} 
  l.levelData = tool.NewLevelData()
  l.LoadContent()
}

func (l *LevelEditor)loadLevelEditorLayers() int {
  var lastLayerWidth int
  for i := 0; i<l.levelEditorLayerRepeat; i++{
    for j := 0; j<l.levelEditorLayerNumber; j++{
      var layerPath string = fmt.Sprintf("content/gamestart/layer%v_0.png", j) 
      var layerImage *ebiten.Image
      var err error
      layerImage, _, err = ebitenutil.NewImageFromFile(layerPath) 
      if err != nil{
        log.Fatal(err)
      }
      layWidth, _ := layerImage.Size()
      laySprite := &texture.Sprite{Position:&tool.FPoint{X: float64(i*layWidth), Y: float64(j*80)}, Texture:layerImage}
      l.levelEdirorLayers = append(l.levelEdirorLayers, laySprite)
      if j == l.levelEditorLayerNumber -1{
        lastLayerWidth = layWidth
      }
    }
  }
  return lastLayerWidth
}

func (l *LevelEditor)LoadGridLine(){
  for i:=0; i<tool.SCRREN_ORI_HEIGHT; i+=tool.TILEHEIGHT{
    newLine := &tool.FRectangle{
      Min:tool.FPoint{X: 0, Y: float64(i)},
      Max:tool.FPoint{X: float64(l.srufaceLayerWidth * l.levelEditorLayerRepeat),
      Y: float64(i)},
    }
    l.lineList = append(l.lineList, newLine)
  }
  for i:=0; i<l.srufaceLayerWidth * l.levelEditorLayerRepeat; i+=tool.TILEWIDTH{
    newLine := &tool.FRectangle{Min:tool.FPoint{X: float64(i), Y: 0}, Max:tool.FPoint{X: float64(i), Y: float64(tool.SCRREN_ORI_HEIGHT)}}
    l.lineList = append(l.lineList, newLine)
  }
}

func (l *LevelEditor)loadSurfaceButton(){
  newTexture, _, err := ebitenutil.NewImageFromFile("content/menu/gridMenu.png")
  if err != nil{
    log.Fatal(err)
  }
  newOpenTileBtn := &ui.Button{Sprite: &texture.Sprite{Texture:newTexture, Position:&tool.FPoint{X: 10, Y: 10}}} 
  l.tileOpenBtn = newOpenTileBtn
}

func(l *LevelEditor)loadTilesButton(){
  baseFileDir := "content/tiles/"
  files, err := os.ReadDir(baseFileDir)
  if err != nil{
    log.Fatal(err)
  }

  tileWidthSpace := 10
  cols := tool.SCRREN_ORI_WIDTH / (tool.TILEWIDTH+ tileWidthSpace)
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
    newPosition := &tool.FPoint{X: float64(currentCol * tool.TILEWIDTH+ currentCol* tileWidthSpace), Y: float64(currentRow * tool.TILEHEIGHT)}
    currentCol++
    if currentCol == cols{
      currentRow ++
      currentCol = 0
    }
    newButton := &ui.Button{Sprite: &texture.Sprite{Texture:newTexture, Position:newPosition, SpriteName:file.Name()}}
    l.tileBtnList = append(l.tileBtnList, newButton)
  }
}

func(l *LevelEditor)initialLevelData(){
  cols := l.srufaceLayerWidth * l.levelEditorLayerRepeat / tool.TILEWIDTH 
  rows := tool.SCRREN_ORI_HEIGHT / tool.TILEHEIGHT 
  for row := 0; row < rows; row++{
    for col := 0; col < cols; col ++{
      newMap := map[string]int{"X": col * tool.TILEWIDTH, "Y": row * tool.TILEHEIGHT, "tile": -1}
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
  l.tileColNumber = l.srufaceLayerWidth * l.levelEditorLayerRepeat / tool.TILEWIDTH
  l.LoadGridLine()
  l.loadSurfaceButton()
  l.loadTilesButton()
  err := l.loadLevelData()
  if err != nil{
    log.Fatal(err)
  }
  l.levelData.LevelInfo["tileColNumber"] = l.tileColNumber
  l.levelData.LevelInfo["tileRowNumber"] = int(tool.SCRREN_ORI_HEIGHT) / tool.TILEHEIGHT
}

func (l *LevelEditor)handleLayerScroll(mod string){
  var leayerLength int = len(l.levelEdirorLayers)
  if mod == "right"{
      if l.levelEdirorLayers[leayerLength-1].GetRec().Max.X <= tool.SCRREN_ORI_WIDTH{
        return
      }
  }else{
      if l.levelEdirorLayers[l.levelEditorLayerNumber-1].GetRec().Min.X >= 0{
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
    tool.GAME_FUCTION = "GAMEMAIN"
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

func(l *LevelEditor)getGridArea()bool{
  x, y := ebiten.CursorPosition()
  inMenuContainer := float64(x) >= l.menuContainer.Min.X && float64(y) >= l.menuContainer.Min.Y && float64(x) <= l.menuContainer.Max.X && float64(y) <= l.menuContainer.Max.Y
  tileOpenBtnRec := l.tileOpenBtn.GetRec()
  inTileOpenBtn := x >= tileOpenBtnRec.Min.X && y >= tileOpenBtnRec.Min.Y && x <= tileOpenBtnRec.Max.X && y <= tileOpenBtnRec.Max.Y
  return (!inMenuContainer || !l.showTilesContainer) && !inTileOpenBtn
}

func(b *LevelEditor)tileOpenBtnClick(){
  if  b.keymap.IsMouseLeftKeyPressed() && !input.MouseLeftInUsing{
    input.MouseLeftInUsing = true
    if b.showTilesContainer{
      b.showTilesContainer = false
      b.tileOpenBtn.Position.Y -= b.menuContainer.Max.Y
    }else{
      b.showTilesContainer = true 
      b.tileOpenBtn.Position.Y += b.menuContainer.Max.Y
    }
  }
}

func(l *LevelEditor)tileSelectClick(button *ui.Button){
  texture.InSelectSprite = button.SpriteName
  n := strings.Split(button.SpriteName, ".")[0]
  var err error
  l.inSelectTile, err = strconv.Atoi(n)
  if err != nil{
    log.Fatal(err)
  }
}

func(l *LevelEditor)tileCollisionSetClick(button *ui.Button){
  l.tileSelectClick(button)
  if fullCollisionList, ok := l.levelData.CollisionData["full"]; ok{
    if index := tool.SliceIndexOf(fullCollisionList, l.inSelectTile); index != -1{
      fullCollisionList = tool.SliceRemove(fullCollisionList, index)
      l.levelData.CollisionData["full"] = fullCollisionList
    }else{
      fullCollisionList = append(fullCollisionList, l.inSelectTile)
      l.levelData.CollisionData["full"] = fullCollisionList
    }
  }else{
      l.levelData.CollisionData["full"] = []int{l.inSelectTile} 
  }
}

func (l *LevelEditor)mouseEvent(){
  // grid area click
  x, y := ebiten.CursorPosition()
  if x >= tool.SCRREN_ORI_WIDTH || y >= tool.SCRREN_ORI_HEIGHT{
    return
  }
  if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !input.MouseLeftInUsing && l.getGridArea(){
    x -= int(l.globelScroll)
    x/=tool.TILEWIDTH
    y/=tool.TILEHEIGHT
    l.levelData.TileData[y*l.tileColNumber+x]["tile"] = l.inSelectTile
  }else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && !input.MouseLeftInUsing && l.getGridArea(){
    x -= int(l.globelScroll)
    x/=tool.TILEWIDTH
    y/=tool.TILEHEIGHT
    l.levelData.TileData[y*l.tileColNumber+x]["tile"] = -1
  }
}

func (l *LevelEditor)tileButtonUpdate(){
  // handle tile button mouse click
  for _, tile := range(l.tileBtnList){
    var clickEvent func()
    if l.keymap.IsMouseLeftKeyPressed() && !input.MouseLeftInUsing && l.showTilesContainer{
      clickEvent = func(){l.tileSelectClick(tile)} 
      tile.Update(clickEvent)
    }else if l.keymap.IsMouseRightKeyPressed() && !input.MouseLeftInUsing && l.showTilesContainer{
      clickEvent = func(){l.tileCollisionSetClick(tile)} 
      tile.Update(clickEvent)
    }
  }
}

func(l *LevelEditor) Update(){
  if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft){
    input.MouseLeftInUsing = false
  }
  l.keyEvent()
  l.tileButtonUpdate()
  l.tileOpenBtn.Update(func(){l.tileOpenBtnClick()})
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
  ebitenutil.DrawRect(screen, l.menuContainer.Min.X, l.menuContainer.Min.Y, l.menuContainer.Max.X, l.menuContainer.Max.Y, tool.COLOR_GREY)
  for _, btn := range(l.tileBtnList){
    btn.Draw(screen)
  }
}

func(l *LevelEditor)DrawButton(screen *ebiten.Image){
  l.tileOpenBtn.Draw(screen)
}

func(l *LevelEditor)getTileCollisionInfo(tile int)string{
  if fullCollisionList, ok := l.levelData.CollisionData["full"]; ok{
    if tool.SliceIndexOf(fullCollisionList, tile) != -1{
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
    newSprite := &texture.Sprite{
      Texture: textTure,
      Position: &tool.FPoint{X:float64(tilex), Y:float64(tile["Y"])},
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

