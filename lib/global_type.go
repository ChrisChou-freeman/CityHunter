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

func NewLevelData()*LevelData{
  ld := new(LevelData)
  ld.init()
  return ld
}

func(l *LevelData)init(){
  l.CollisionData = make(map[string][]int)
}

