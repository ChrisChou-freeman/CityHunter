package lib 

import(
  "fmt"
  "log"
  "strconv"
  "strings"
  "github.com/hajimehoshi/ebiten/v2"
)

type Button struct{
  *Sprite
  levelEditor *LevelEditor
  buttonType string
  keyMap *KeyMap
}

func(b *Button)mouseHoverOnButton()bool{
  x, y := ebiten.CursorPosition()
  thisRec := b.getRec()
  return x>=thisRec.Min.X && y>=thisRec.Min.Y && x<=thisRec.Max.X && y<=thisRec.Max.Y
}

func(b *Button)tileOpenBtnClick(){
  MouseLeftInUsing = true
  if b.levelEditor.showTilesContainer{
    b.levelEditor.showTilesContainer = false
    b.levelEditor.tileOpenBtn.Position.Y -= b.levelEditor.menuContainer.Max.Y
  }else{
    b.levelEditor.showTilesContainer = true 
    b.levelEditor.tileOpenBtn.Position.Y += b.levelEditor.menuContainer.Max.Y
  }
}

func(b *Button)tileSelectClick(){
  if !b.levelEditor.showTilesContainer{
    return
  }
  InSelectSprite = b.SpriteName
  n := strings.Split(b.SpriteName, ".")[0]
  var err error
  b.levelEditor.inSelectTile, err = strconv.Atoi(n)
  if err != nil{
    log.Fatal(err)
  }
}

func(b *Button)tileCollisionSetClick(){
  if(!b.levelEditor.showTilesContainer){
    return
  }
  b.tileSelectClick()
  if fullCollisionList, ok := b.levelEditor.levelData.CollisionData["full"]; ok{
    if index := SliceContainItem(fullCollisionList, b.levelEditor.inSelectTile); index != -1{
      fullCollisionList = SliceRemove(fullCollisionList, index)
      b.levelEditor.levelData.CollisionData["full"] = fullCollisionList
    }else{
      fullCollisionList = append(fullCollisionList, b.levelEditor.inSelectTile)
      b.levelEditor.levelData.CollisionData["full"] = fullCollisionList
    }
  }else{
      b.levelEditor.levelData.CollisionData["full"] = []int{b.levelEditor.inSelectTile} 
  }
  fmt.Println(b.levelEditor.levelData.CollisionData)
}

func(b *Button)onClick(){
  if b.mouseHoverOnButton(){
    if b.keyMap.IsMouseLeftKeyPressed() && !MouseLeftInUsing{
      switch b.buttonType{
        case "tileOpenBtn":
          b.tileOpenBtnClick()
        case "tile":
          b.tileSelectClick()
      }
    }
    if b.keyMap.IsMouseRightKeyPressed() && !MouseLeftInUsing{
      switch b.buttonType{
      case "tile":
        b.tileCollisionSetClick()
      }
    }
  }
}

func(b *Button)Update(){
  b.onClick()
}
