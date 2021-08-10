package lib

import(
  // "fmt"
  "log"
  "image"
  "image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/text"
  "github.com/hajimehoshi/ebiten/v2/inpututil"
  "github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
)

type Menu struct{
  MenuName string
  MenuColor color.RGBA
  Position image.Point
  fontType font.Face
}

var selectedMenu string

func (m *Menu)init(){
  m.loadContent()
}

func (m *Menu)loadContent(){
  m.loadFont()
}

func (m *Menu)loadFont(){
  var tt *opentype.Font
  var err error
  tt, err = opentype.Parse(fonts.MPlus1pRegular_ttf) 
  if err != nil{
    log.Fatal(err)
  }
  var fontOption *opentype.FaceOptions = new(opentype.FaceOptions)
  fontOption.Size = 34
  fontOption.DPI = 70
  fontOption.Hinting = font.HintingFull
  m.fontType, err = opentype.NewFace(tt, fontOption)
  if err != nil{
    log.Fatal(err)
  }
}

func (m *Menu)getRec() image.Rectangle {
  var rec image.Rectangle = text.BoundString(m.fontType, m.MenuName)
  var newRec *image.Rectangle = new(image.Rectangle)
  var rSize image.Point = rec.Size()
  newRec.Min.X = m.Position.X
  newRec.Min.Y = m.Position.Y + rec.Min.Y
  newRec.Max.X = m.Position.X + rSize.X
  newRec.Max.Y = m.Position.Y + rec.Min.Y + rSize.Y
  return *newRec 
}

func (m *Menu)onClick(){
 var x, y int = ebiten.CursorPosition()
  if m.containPoint(x, y) {
    selectedMenu = m.MenuName
    if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft){
      switch m.MenuName{
        case "EXIT":
          GAMEMODE = "EXIT"
        case "DEV":
          GAMEMODE = "DEV"
        case "START":
          GAMEMODE = "START"
      }
    }
  }
}

func (m *Menu)Update(){
  m.onClick()
}

func (m *Menu)Draw(scrren *ebiten.Image){
  var menuColor color.RGBA = m.MenuColor
  if m.MenuName == selectedMenu{
    menuColor = COLOR_YELLOW
  }
  text.Draw(scrren, m.MenuName, m.fontType, m.Position.X, m.Position.Y, menuColor)
}

func (m *Menu)containPoint(x, y int) bool{
  var rec image.Rectangle = m.getRec()
  return x >= rec.Min.X && y >= rec.Min.Y && x <= rec.Max.X && y <= rec.Max.Y
}

func NewMenu()*Menu{
  var m *Menu = new(Menu)
  m.init()
  return m
}
