package lib

type GameStart struct{
  layers []*Sprite
  tilesList []*Sprite
  levelData *LevelData
  keymap *KeyMap
  levelNumber int
}

func NewGameStart() *GameStart{
  ngs := new(GameStart)
  ngs.init()
  return ngs
}

func(gs *GameStart)init(){
}
