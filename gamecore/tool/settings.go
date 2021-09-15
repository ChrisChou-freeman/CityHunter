package tool

import (
	"image/color"
)

const (
	SCRREN_ORI_WIDTH  int     = 800
	SCRREN_ORI_HEIGHT int     = 480
	SCALE             float64 = 1.5
	SCRREN_WIDTH      int     = int(float64(SCRREN_ORI_WIDTH) * SCALE)
	SCRREN_HEIGHT     int     = int(float64(SCRREN_ORI_HEIGHT) * SCALE)
	TILEWIDTH         int     = 32
	TILEHEIGHT        int     = 32
	TILES_PATH        string  = "content/tiles/%v.png"
	PLAYERTILE        int     = 15
	GAME_MAIN         string  = "Main"
	GAME_DEVELOPMENT  string  = "Development"
	GAME_START        string  = "Start"
	GAME_QUIT         string  = "Quit"
)

var (
	COLOR_YELLOW        color.RGBA = color.RGBA{255, 255, 1, 255}
	COLOR_WHITE         color.RGBA = color.RGBA{255, 255, 255, 255}
	COLOR_GREY          color.RGBA = color.RGBA{192, 192, 192, 255}
	COLOR_RED           color.RGBA = color.RGBA{240, 52, 52, 255}
	COLOR_EXPLODE       color.RGBA = color.RGBA{180, 180, 180, 255}
	COLOR_BULLET        color.RGBA = color.RGBA{255, 255, 255, 255}
	COLOR_BULLET_TRACER color.RGBA = COLOR_BULLET
	ENEMYTILES          []int      = []int{16}
	GAME_FUCTION        string     = "Main"
)
