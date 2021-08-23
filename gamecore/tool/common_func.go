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

func SliceRemove(slice []int, index int) []int {
    return append(slice[:index], slice[index+1:]...)
}

func Clamp(n, a, b float64) float64 {
	return math.Max(math.Min(n, math.Max(a, b)), math.Min(a, b))
}

func CollisionDetect(spritRec image.Rectangle, levelData *LevelData, vector *image.Point, position *image.Point) {
  spritCent := image.Point{X: (spritRec.Min.X + spritRec.Max.X) / 2, Y: (spritRec.Min.Y + spritRec.Max.Y) / 2}
  if spritCent.X == 0 && vector.X < 0{
    return
  }
  if spritCent.X == 13 && vector.X > 0 {
    return
  }

  var leftTilePosition image.Point
  var rightTilePosition  image.Point
  var topTilePosition image.Point
  var bottomTilePosition  image.Point
  leftTilePosition = image.Point{X: (spritCent.X / TILEWIDTH - 1) * TILEWIDTH, Y: (spritCent.Y / TILEHEIGHT) * TILEHEIGHT} 
  rightTilePosition = image.Point{X: (spritCent.X / TILEWIDTH + 1) * TILEWIDTH, Y: (spritCent.Y / TILEHEIGHT) * TILEHEIGHT} 
  topTilePosition = image.Point{X: (spritCent.X / TILEWIDTH) * TILEWIDTH, Y: (spritCent.Y / TILEHEIGHT - 1) * TILEHEIGHT} 
  bottomTilePosition = image.Point{X: (spritCent.X / TILEWIDTH) * TILEHEIGHT, Y: (spritCent.Y / TILEHEIGHT + 1) * TILEHEIGHT} 

  var leftTile int = -2
  var rightTile int = -2
  var topTile int = -2
  var bottomTile int = -2

  for _, item := range(levelData.TileData){
    if leftTilePosition.X == item["X"] && leftTilePosition.Y == item["Y"]{
      leftTile = item["tile"]
    }
    if rightTilePosition.X == item["X"] && rightTilePosition.Y == item["Y"]{
      rightTile = item["tile"]
    }
    if topTilePosition.X == item["X"] && topTilePosition.Y == item["Y"]{
     topTile = item["tile"]
    }
    if bottomTilePosition.X == item["X"] && bottomTilePosition.Y == item["Y"]{
     bottomTile = item["tile"]
    }
    if leftTile != -2 && rightTile != -2 && topTile != -2 && bottomTile != -2{
      break
    }
  }
  
  // // up
  // if vector.Y < 0{
  //   if SliceContainItem(levelData.CollisionData["full"], topTile) != -1{
  //     topTileBottom := topTilePosition.Y * TILEHEIGHT + TILEHEIGHT
  //     Y := position.Y + vector.Y
  //     position.Y = int(Clamp(float64(Y), float64(topTileBottom), float64(topTileBottom))) 
  //   }
  // }

  // down
  if vector.Y > 0{
    if SliceContainItem(levelData.CollisionData["full"], bottomTile) != -1{ 
      if spritRec.Max.Y + 1 + vector.Y > bottomTilePosition.Y {
        position.Y = bottomTilePosition.Y - TILEHEIGHT
        vector.Y = 0
      }
    }
  }

  position.X += vector.X
  position.Y += vector.Y
}
