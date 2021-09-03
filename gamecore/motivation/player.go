package motivation

import (
	// "fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/ChrisChou-freeman/CityHunter/gamecore/input"
	"github.com/ChrisChou-freeman/CityHunter/gamecore/texture"
	"github.com/ChrisChou-freeman/CityHunter/gamecore/tool"
)

type Player struct {
	idle             *texture.AnimationSprite
	run              *texture.AnimationSprite
	death            *texture.AnimationSprite
	jump             *texture.AnimationSprite
	currentAnimation *texture.AnimationSprite
	position         *image.Point
	vector           *image.Point
	levelData        *tool.LevelData
	keymap           *input.KeyMap
	action           string
	gravity          int
	speed            int
	isJumping        bool
	isFulling        bool
	flip             bool
	jumpForce        int
	counter          int
	bulletSpeed      int
	aotuBulletSpeed  int
	bulletList       []*texture.MotivationSprite
	bulletTexture    *ebiten.Image
}

func NewPlayer(position *image.Point, levelData *tool.LevelData) *Player {
	p := new(Player)
	p.init(position, levelData)
	return p
}

func (p *Player) init(position *image.Point, levelData *tool.LevelData) {
	p.position = position
	p.gravity = 6
	p.speed = 2
	p.jumpForce = 60
	p.bulletSpeed = 5
	p.vector = &image.Point{}
	p.action = "idle"
	p.levelData = levelData
	p.keymap = new(input.KeyMap)
	p.loadContent()
}

func (p *Player) loadAnimation() error {
	idleT, _, err := ebitenutil.NewImageFromFile("content/player/Idle/sheet.png")
	if err != nil {
		return err
	}
	p.idle = texture.NewAnimationSprite(idleT, 28, 36)

	deathT, _, err := ebitenutil.NewImageFromFile("content/player/death/sheet.png")
	if err != nil {
		return err
	}
	p.death = texture.NewAnimationSprite(deathT, 35, 36)

	jumpT, _, err := ebitenutil.NewImageFromFile("content/player/jump/0.png")
	if err != nil {
		return err
	}
	p.jump = texture.NewAnimationSprite(jumpT, 28, 36)

	runT, _, err := ebitenutil.NewImageFromFile("content/player/run/sheet.png")
	if err != nil {
		return err
	}
	p.run = texture.NewAnimationSprite(runT, 28, 36)

	bulletT, _, err := ebitenutil.NewImageFromFile("content/items/bullet.png")
	if err != nil {
		return nil
	}
	p.bulletTexture = bulletT

	return nil
}

func (p *Player) loadContent() {
	err := p.loadAnimation()
	if err != nil {
		log.Fatal(err)
	}
}

func (p *Player) keyEvent() {
	p.action = "idle"
	if p.keymap.IsKeyLeftHoldPressed() {
		p.action = "moveLeft"
		p.flip = true
	}

	if p.keymap.IsKeyRightHoldPressed() {
		p.action = "moveRight"
		p.flip = false
	}

	if p.keymap.IsKeyJumpPressed() && p.isOnGround() {
		p.isJumping = true
	}

	if p.keymap.IsKeyAttackPressed() {
		p.shotBullet()
	}
}

func (p *Player) isOnGround() bool {
	return p.vector.Y == 0
}

func (p *Player) motivation() {
	switch p.action {
	case "idle":
		p.vector.Y = p.gravity
		p.vector.X = 0
	case "moveRight":
		p.vector.X = p.speed
		p.vector.Y = p.gravity
	case "moveLeft":
		p.vector.X = -p.speed
		p.vector.Y = p.gravity
	}
	if p.isJumping && !p.isFulling {
		p.vector.Y = -p.speed * 2
	}
}

func (p *Player) vectorHandle() {
	if p.currentAnimation == nil {
		return
	}
	tool.CollisionDetect(p.currentAnimation.GetRec(), p.levelData, p.vector, p.position)
	p.position.X += p.vector.X
	p.position.Y += p.vector.Y
	if p.isJumping {
		p.counter -= p.vector.Y
		if p.counter >= p.jumpForce {
			p.vector.Y = p.speed
			p.isJumping = false
			p.isFulling = true
			p.counter = 0
		}
	}
	if p.vector.Y == 0 {
		p.isFulling = false
	}
}

func (p *Player) animationControl() {
	if p.isJumping || p.isFulling {
		if p.currentAnimation == p.jump {
			return
		}
		p.currentAnimation = p.jump
		p.currentAnimation.PlayOnce()
	} else {
		switch p.action {
		case "idle":
			p.currentAnimation = p.idle
			p.currentAnimation.Play()
		case "moveRight":
			p.currentAnimation = p.run
			p.currentAnimation.Play()
		case "moveLeft":
			p.currentAnimation = p.run
			p.currentAnimation.Play()
		}
	}
	p.currentAnimation.Flip = p.flip
}

func (p *Player) shotBullet() {
	var bulletVectory *image.Point
	var bulletPosition *tool.FPoint
	if p.flip {
		bulletVectory = &image.Point{X: -p.bulletSpeed}
		bulletPosition = &tool.FPoint{X: float64(p.position.X) - float64(p.bulletTexture.Bounds().Dx()), Y: float64(p.position.Y + p.idle.GetRec().Dy()/2 - 5)}
	} else {
		bulletVectory = &image.Point{X: p.bulletSpeed}
		bulletPosition = &tool.FPoint{X: float64(p.position.X + p.idle.GetRec().Dx()), Y: float64(p.position.Y + p.idle.GetRec().Dy()/2 - 5)}
	}
	newSprite := &texture.Sprite{
		Texture:  p.bulletTexture,
		Position: bulletPosition,
	}
	newBullet := texture.NewMotivationSprite(newSprite, 120, bulletVectory)
	p.bulletList = append(p.bulletList, newBullet)
}

func (p *Player) updateBullet() {
	deadBulletIndex := []int{}
	for index, bullet := range p.bulletList {
		tool.CollisionDetect(bullet.GetRec(), p.levelData, bullet.Vector, &image.Point{int(bullet.Position.X), int(bullet.Position.Y)})
		bullet.Update()
		if !bullet.Islife() {
			deadBulletIndex = append(deadBulletIndex, index)
		}
	}

	// remove endlife bullet
	for index, bindex := range deadBulletIndex {
		if bindex == (len(p.bulletList) - 1) {
			p.bulletList = p.bulletList[:bindex]
		} else {
			p.bulletList = append(p.bulletList[:bindex], p.bulletList[bindex+1:]...)
			for i := index + 1; i < len(deadBulletIndex); i++ {
				deadBulletIndex[i] -= 1
			}
		}
	}
}

func (p *Player) Update() {
	p.keyEvent()
	p.motivation()
	p.vectorHandle()
	p.animationControl()
	p.currentAnimation.Update(p.position)
	p.updateBullet()
}

func (p *Player) DrawBullet(screen *ebiten.Image) {
	for _, bullet := range p.bulletList {
		bullet.Draw(screen)
		if bullet.Vector.X == 0 {
			bullet.Kill()
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.currentAnimation.Draw(screen)
	p.DrawBullet(screen)
}

func (p *Player) Dispose() {
	p.idle.Dispose()
	p.death.Dispose()
	p.run.Dispose()
	p.jump.Dispose()
	p.bulletTexture.Dispose()
}
