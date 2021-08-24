package tool 

import(
  // "fmt"
  "image"
  "math"
)

func SliceContainItem(l []int, i int)int{
  for index, item := range(l){
    if item == i{
      return index 
    } 
  }
  return -1
}

func SliceContainItemsOr(l []int, items []int)bool{
  for _, item := range(l){
    for _, item2 := range(items){
      if item == item2{
        return true
      }
    }
  }
  return false
}

func SliceRemove(slice []int, index int) []int {
    return append(slice[:index], slice[index+1:]...)
}

func Clamp(n, a, b float64) float64 {
	return math.Max(math.Min(n, math.Max(a, b)), math.Min(a, b))
}

func getLevalDataIndex(position image.Point, levelData *LevelData) int {
  return position.Y * levelData.LevelInfo["tileColNumber"] + position.X
}

func CollisionDetect(spritRec image.Rectangle, levelData *LevelData, vector *image.Point, position *image.Point, gravity int) {
  borderLessSpritRec := spritRec
  borderLessSpritRec.Min.X += 1
  borderLessSpritRec.Max.X -= 1
  borderLessSpritRec.Min.Y += 1
  borderLessSpritRec.Max.Y -= 1
  topLeftTilePosition := image.Point{X: borderLessSpritRec.Min.X / TILEWIDTH, Y: borderLessSpritRec.Min.Y / TILEHEIGHT -1} 
  topRightTilePosition := image.Point{X: borderLessSpritRec.Max.X / TILEWIDTH, Y: borderLessSpritRec.Min.Y / TILEHEIGHT -1} 
  bottomLeftTilePosition := image.Point{X: borderLessSpritRec.Min.X / TILEWIDTH, Y: borderLessSpritRec.Max.Y / TILEHEIGHT + 1}
  bottomRightTilePosition := image.Point{X: borderLessSpritRec.Max.X / TILEWIDTH, Y: borderLessSpritRec.Max.Y / TILEHEIGHT + 1} 

  leftUpTilePostion := image.Point{X: borderLessSpritRec.Min.X / TILEWIDTH -1, Y: borderLessSpritRec.Min.Y / TILEHEIGHT}
  leftDownTilePostion := image.Point{X: borderLessSpritRec.Min.X / TILEWIDTH -1, Y: borderLessSpritRec.Max.Y / TILEHEIGHT}
  rightUpTilePostion := image.Point{X: borderLessSpritRec.Max.X / TILEWIDTH + 1, Y: borderLessSpritRec.Min.Y / TILEHEIGHT}
  rightDownTilePostion := image.Point{X: borderLessSpritRec.Max.X / TILEWIDTH + 1, Y: borderLessSpritRec.Max.Y / TILEHEIGHT}

  topleftTile := -2 
  topRightTile := -2
  bottomLeftTile := -2
  leftUpTile  := -2
  leftDownTile := -2
  rightUpTile := -2
  rightDownTile := -2

  if leftUpTilePostion.X < 0{
    if spritRec.Min.X + vector.X < 0{
      position.X = 0
      vector.X = 0
    }
  }else{
    leftUpTile = levelData.TileData[getLevalDataIndex(leftUpTilePostion, levelData)]["tile"]
    leftDownTile = levelData.TileData[getLevalDataIndex(leftDownTilePostion, levelData)]["tile"]
  }

  if rightUpTilePostion.X > levelData.LevelInfo["tileColNumber"] - 1{
    if spritRec.Max.X + vector.X > levelData.LevelInfo["tileColNumber"] * TILEWIDTH{
      position.X = levelData.LevelInfo["tileColNumber"] * TILEWIDTH - spritRec.Dx()
      vector.X = 0
    }
  }else{
    rightUpTile = levelData.TileData[getLevalDataIndex(rightUpTilePostion, levelData)]["tile"]
    rightDownTile = levelData.TileData[getLevalDataIndex(rightDownTilePostion, levelData)]["tile"]
  }

  topleftTile = levelData.TileData[getLevalDataIndex(topLeftTilePosition, levelData)]["tile"]
  topRightTile = levelData.TileData[getLevalDataIndex(topRightTilePosition, levelData)]["tile"]
  bottomLeftTile = levelData.TileData[getLevalDataIndex(bottomLeftTilePosition, levelData)]["tile"]
  bottomRightTile := levelData.TileData[getLevalDataIndex(bottomRightTilePosition, levelData)]["tile"]

  // up
  if vector.Y < 0{
    if SliceContainItemsOr(levelData.CollisionData["full"], []int{topleftTile, topRightTile}){
      if spritRec.Min.Y - 1 + vector.Y < topLeftTilePosition.Y * TILEHEIGHT +TILEHEIGHT{
        position.Y = topLeftTilePosition.Y * TILEHEIGHT + TILEHEIGHT
        vector.Y = 0
      }
    }
  }

  // down
  if vector.Y > 0{
    if SliceContainItemsOr(levelData.CollisionData["full"], []int{bottomLeftTile, bottomRightTile}){ 
      if spritRec.Max.Y + 1 + vector.Y > bottomLeftTilePosition.Y * TILEHEIGHT {
        position.Y = bottomLeftTilePosition.Y * TILEHEIGHT - spritRec.Dy()
        vector.Y = 0
      }
    }
  }

  //left
  if vector.X < 0{
    if SliceContainItemsOr(levelData.CollisionData["full"], []int{leftUpTile, leftDownTile}){
      if spritRec.Min.X - 1 + vector.X < leftUpTilePostion.X *TILEWIDTH + TILEWIDTH{
        position.X = leftUpTilePostion.X * TILEWIDTH + TILEWIDTH
        vector.X = 0
      }
    }
  }

  //right
  if vector.X > 0{
    if SliceContainItemsOr(levelData.CollisionData["full"], []int{rightUpTile, rightDownTile}){
      if spritRec.Max.X + 1 + vector.X >= rightUpTilePostion.X * TILEWIDTH{
        // fmt.Println(spritRec)
        // fmt.Println(rightUpTilePostion)
        position.X = rightUpTilePostion.X * TILEWIDTH - spritRec.Dx()
        vector.X = 0
      }
    }
  }
}
