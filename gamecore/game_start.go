package gamecore 

import(
  "os"
  "io"
  "log"
  "fmt"
  "image"
  "errors"
  "encoding/json"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"

  "github.com/ChrisChou-freeman/CityHunter/gamecore/input"
  "github.com/ChrisChou-freeman/CityHunter/gamecore/texture"
  "github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
  "github.com/ChrisChou-freeman/CityHunter/gamecore/player"
)

type GameStart struct{
  layers []*texture.Sprite
  tilesList []*texture.Sprite
  levelData *tool.LevelData
  keymap *input.KeyMap
  player *player.Player
  enemys []*texture.Sprite
  levelNumber int
  layerRepeat int
  layerNumber int
  currentLevel int
  surfaceLayerWidth int
}

func NewGameStart() *GameStart{
  ngs := new(GameStart)
  ngs.init()
  return ngs
}

func(gs *GameStart)init(){
  gs.layerRepeat = 2
  gs.layerNumber = 4
  gs.levelData = tool.NewLevelData() 
  gs.loadContent()
}

func (gs *GameStart)loadLayers() {
  // var lastLayerWidth int
  for i := 0; i<gs.layerRepeat; i++{
    for j := 0; j<gs.layerNumber; j++{
      var layerPath string = fmt.Sprintf("content/gamestart/layer%v_0.png", j) 
      var layerImage *ebiten.Image
      var err error
      layerImage, _, err = ebitenutil.NewImageFromFile(layerPath) 
      if err != nil{
        log.Fatal(err)
      }
      layWidth, _ := layerImage.Size()
      laySprite := &texture.Sprite{Position: &tool.FPoint{X: float64(i*layWidth), Y: float64(j*80)}, Texture:layerImage}
      gs.layers= append(gs.layers, laySprite)
      if j == gs.layerNumber-1{
        gs.surfaceLayerWidth = layWidth
      }
    }
  }
}

func(gs *GameStart)loadLevelData() error{ 
  levelDataPath := fmt.Sprintf("content/leveldata/%v.json", gs.currentLevel)
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
    return errors.New("ERROR: empty level error")
  }else{
    jsonByte, err := io.ReadAll(fileObj)
    if err != nil{
      return err
    }
    err = json.Unmarshal(jsonByte, gs.levelData)
    if err != nil{
      return err
    }
  }
  return nil
}

func(gs *GameStart)initialTilesByLevelData(){
  for _, tileInfo := range(gs.levelData.TileData){
    if tileInfo["tile"] == -1{
      continue
    }
    tilePath := fmt.Sprintf(tool.TILES_PATH, tileInfo["tile"])
    newTexture, _, err := ebitenutil.NewImageFromFile(tilePath)
    if err != nil{
      log.Fatal(err)
    }
    newSprite := &texture.Sprite{Texture: newTexture, Position: &tool.FPoint{X: float64(tileInfo["X"]), Y: float64(tileInfo["Y"])}}
    switch {
    case tileInfo["tile"] == tool.PLAYERTILE:
      gs.player = player.NewPlayer(&image.Point{X: tileInfo["X"], Y: tileInfo["Y"]})
    case tool.SliceContainItem(tool.ENEMYTILES, tileInfo["tile"]) != -1:
      gs.enemys = nil 
    default:
      gs.tilesList = append(gs.tilesList, newSprite)
    }
  }
}

func(gs *GameStart)loadContent(){
  gs.loadLayers()
  err := gs.loadLevelData()
  if err != nil{
    log.Fatal(err)
  }
  gs.initialTilesByLevelData()
}

func(gs *GameStart)Update(){
}

func(gs *GameStart)Draw(screen *ebiten.Image){
  for _, s := range(gs.layers){
    s.Draw(screen)
  }
  for _, s := range(gs.tilesList){
    s.Draw(screen)
  }
}

func(gs *GameStart)Dispose(){
  for _, s := range(gs.layers){
    s.Dispose()
  }
  for _, s := range(gs.tilesList){
    s.Dispose()
  }
  // gs.player.Dispose()
}

