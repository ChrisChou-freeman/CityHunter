package lib

type FPoint struct{
  X float64
  Y float64
}

type FRectangle struct{
  Min FPoint
  Max FPoint
}

type LevelData struct{
  TileData []map[string]int
  CollisionData map[string][]int
}

