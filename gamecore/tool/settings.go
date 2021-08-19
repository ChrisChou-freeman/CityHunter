package tool 

import(
  "image/color"
)

const (
  SCRREN_WIDTH = 1280 
  SCRREN_HEIGHT  = 720 
  SCRREN_ORI_WIDTH = 800
  SCRREN_ORI_HEIGHT = 480
  PLAYERTILE = 15

  SCRREN_WIDTH_SCAL = float64(SCRREN_ORI_WIDTH) / float64(SCRREN_WIDTH)
  SCRREN_HEIGHT_SCAL = float64(SCRREN_ORI_HEIGHT) / float64(SCRREN_HEIGHT) 
  TILES_PATH = "content/tiles/%v.png"
)

var (
  COLOR_YELLOW color.RGBA = color.RGBA{255, 255, 1, 255}
  COLOR_WHITE color.RGBA = color.RGBA{255, 255, 255, 255}
  COLOR_GREY color.RGBA = color.RGBA{192, 192, 192, 255}
  COLOR_RED color.RGBA = color.RGBA{240, 52, 52, 255}
  ENEMYTILES []int = []int{16} 
  GAME_FUCTION string = "GAMEMAIN"
)

