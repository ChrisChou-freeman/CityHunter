package lib

import(
  "image/color"
)

const (
  SCRREN_WIDTH int = 1280 
  SCRREN_HEIGHT  int = 720 

  SCRREN_ORI_WIDTH int = 800
  SCRREN_ORI_HEIGHT int  = 480

  SCRREN_WIDTH_SCAL = float64(SCRREN_ORI_WIDTH) / float64(SCRREN_WIDTH)
  SCRREN_HEIGHT_SCAL = float64(SCRREN_ORI_HEIGHT) / float64(SCRREN_HEIGHT) 
)

var (
  COLOR_YELLOW color.RGBA = color.RGBA{255, 255, 1, 255}
  COLOR_WHITE color.RGBA = color.RGBA{255, 255, 255, 255}
  GAMEMODE string = "GAMEMAIN"
)

