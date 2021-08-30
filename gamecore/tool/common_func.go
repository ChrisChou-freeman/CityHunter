package tool 

import(
  "image"
  "math"
)

func SliceIndexOf(l []int, i int)int{
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
