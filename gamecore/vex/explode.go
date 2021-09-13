package vex

import (
	// "fmt"
	"image"
	"math/rand"

	"github.com/ChrisChou-freeman/CityHunter/gamecore/shape"
	"github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
	"github.com/hajimehoshi/ebiten/v2"
)

type Explode struct {
	position   *image.Point
	circles    []*shape.Circle
	expandTime int
	explodeDot int
	randSize   int
	down       bool
	radius     int
}

func NewExplode(position *image.Point) *Explode {
	e := new(Explode)
	e.Init(position)
	return e
}

func (e *Explode) Init(position *image.Point) {
	e.position = position
	e.circles = []*shape.Circle{}
	e.explodeDot = 70
	e.expandTime = 240
	e.randSize = 10
	e.radius = 15
	e.LoadExplodeDot()
}

func (e *Explode) ExplodeRound() int {
	return (e.radius + e.randSize) * 4
}

func (e *Explode) LoadExplodeDot() {
	for i := 0; i < e.explodeDot; i++ {
		rand.Seed(int64(i))
		radius := float64(e.radius) + float64(rand.Intn(e.randSize))
		vector := &tool.FPoint{}
		vector.X = float64(float64(rand.Intn(e.randSize*4))/5 - 4)
		vector.Y = float64(-rand.Intn(6))
		offsetX := rand.Intn(e.randSize)
		offsetY := rand.Intn(e.randSize * 3)
		if i%2 == 0 {
			offsetX *= -1
		}
		explodPosition := &tool.FPoint{
			X: float64(e.position.X + offsetX),
			Y: float64(e.position.Y + offsetY),
		}
		newCircle := shape.NewCircle(radius, explodPosition, vector, tool.COLOR_EXPLODE)
		e.circles = append(e.circles, newCircle)
	}
}

func (e *Explode) UpdateExplode(c *shape.Circle) {
	c.Postion.Y += c.Velocity.Y
	c.Postion.X += c.Velocity.X
	randValue := rand.Float64() * float64(rand.Intn(e.randSize))
	if e.expandTime > 0 {
		c.Radius += randValue
	} else {
		c.Radius -= randValue
	}
	c.Velocity.Y += 0.03
	e.expandTime--
	c.CColor.A -= 2
}

func (e *Explode) IsDown() bool {
	return e.down
}

func (e *Explode) Update() {
	need_remove := []int{}
	for index, circle := range e.circles {
		e.UpdateExplode(circle)
		if circle.Radius <= 0 {
			need_remove = append(need_remove, index)
		}
	}
	for index, cIndex := range need_remove {
		if cIndex == len(e.circles)-1 {
			e.circles = e.circles[:cIndex]
		} else {
			e.circles = append(e.circles[:cIndex], e.circles[cIndex+1:]...)
			for i := index + 1; i < len(need_remove); i++ {
				need_remove[i]--
			}
		}
	}
	if len(e.circles) == 0 {
		e.down = true
	}
}

func (e *Explode) Draw(screen *ebiten.Image) {
	for _, circle := range e.circles {
		circle.Draw(screen)
	}
}
