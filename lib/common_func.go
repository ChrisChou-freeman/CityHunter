package lib

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
