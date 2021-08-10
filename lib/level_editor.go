package lib

import(
  "image"
  "log"
  "fmt"
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type LevelEditor struct{
  buttonList []*Button
  levelEdirorLayers []*Sprite
  levelEditorLayerNumber  int
  levelEditorLayerRepeat int
  layerScrrolSpeed int
  keymap *KeyMap
}

func (l *LevelEditor)Init(){
  l.levelEditorLayerNumber  = 4
  l.levelEditorLayerRepeat = 3
  l.layerScrrolSpeed = 2
  l.LoadContent()
  l.keymap = new(KeyMap)
}

func handleLayerScroll(){
}

func (l *LevelEditor)keyEvent(){
  if l.keymap.IsKeyBackPressed(){
    GAMEMODE = "GAMEMAIN"
  }
}

func(l *LevelEditor) Update(){
  l.keyEvent()
}

func(l *LevelEditor) Draw(scrren *ebiten.Image){
  for _,layer := range(l.levelEdirorLayers){
    var iop *ebiten.DrawImageOptions = new(ebiten.DrawImageOptions)
    iop.GeoM.Translate(float64(layer.Position.X), float64(layer.Position.Y))
    scrren.DrawImage(layer.Texture, iop)
  }
}

func(l *LevelEditor) Dispose(){
  for _, item := range(l.levelEdirorLayers){
    item.Dispose()
  }
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
      laySprite.Position = image.Point{i*layWidth, j*80}
      laySprite.Texture = layerImage
      l.levelEdirorLayers = append(l.levelEdirorLayers, laySprite)
      if j == l.levelEditorLayerNumber -1{
        lastLayerWidth = layWidth
      }
    }
  }
  return lastLayerWidth
}

func (l *LevelEditor)LoadContent(){
  l.loadLevelEditorLayers()
}

func NewLevelEditor() *LevelEditor{
  var le *LevelEditor = new(LevelEditor)
  le.Init()
  return le
}
