package tool

type FPoint struct {
	X float64
	Y float64
}

func (fp *FPoint) Add(fpoint FPoint) {
	fp.X += fpoint.X
	fp.Y += fpoint.Y
}

type FRectangle struct {
	Min FPoint
	Max FPoint
}

type LevelData struct {
	TileData      []map[string]int
	CollisionData map[string][]int
	LevelInfo     map[string]int
}

func NewLevelData() *LevelData {
	ld := new(LevelData)
	ld.init()
	return ld
}

func (l *LevelData) init() {
	l.CollisionData = make(map[string][]int)
}
