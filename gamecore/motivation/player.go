package motivation

import (
	"fmt"
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
	fulling          *texture.AnimationSprite
	currentAnimation *texture.AnimationSprite
	shootFire        *texture.AnimationSprite
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
	characterColor   string
	bulletList       []*texture.MotivationSprite
	bulletTexture    *ebiten.Image
	throwGrenade     *ThrowGrenade
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
	p.bulletSpeed = 6
	p.vector = &image.Point{}
	p.action = "idle"
	p.levelData = levelData
	p.characterColor = "black"
	p.keymap = new(input.KeyMap)
	p.loadContent()
}

func (p *Player) loadAnimation() error {
	idleT, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("content/player/%v/idle.png", p.characterColor))
	if err != nil {
		return err
	}
	p.idle = texture.NewAnimationSprite(idleT, 28, 36, 6)

	deathT, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("content/player/%v/death.png", p.characterColor))
	if err != nil {
		return err
	}
	p.death = texture.NewAnimationSprite(deathT, 35, 36, 6)

	jumpT, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("content/player/%v/jump.png", p.characterColor))
	if err != nil {
		return err
	}
	p.jump = texture.NewAnimationSprite(jumpT, 28, 36, 6)

	runT, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("content/player/%v/run.png", p.characterColor))
	if err != nil {
		return err
	}
	p.run = texture.NewAnimationSprite(runT, 28, 36, 6)

	fullingT, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("content/player/%v/full.png", p.characterColor))
	if err != nil {
		return err
	}
	p.fulling = texture.NewAnimationSprite(fullingT, 28, 36, 6)

	shootFireT, _, err := ebitenutil.NewImageFromFile("content/vex/shootfire.png")
	if err != nil {
		return err
	}
	p.shootFire = texture.NewAnimationSprite(shootFireT, 7, 7, 1)

	bulletT := ebiten.NewImage(3, 3)
	bulletT.Fill(tool.COLOR_GREY)
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
	if !p.isJumping && !p.isFulling {
		p.action = "idle"
		p.vector.X = 0
		p.vector.Y = p.gravity
	}

	if p.keymap.IsKeyLeftHoldPressed() {
		if p.isJumping {
			p.action = "jumpLeft"
		} else if p.isFulling {
			p.action = "fullLeft"
		} else {
			p.action = "moveLeft"
			p.vector.Y = p.gravity
		}
		p.vector.X = -p.speed
		p.flip = true
	}

	if p.keymap.IsKeyRightHoldPressed() {
		if p.isJumping {
			p.action = "jumpRight"
		} else if p.isFulling {
			p.action = "fullRight"
		} else {
			p.action = "moveRight"
			p.vector.Y = p.gravity
		}
		p.vector.X = p.speed
		p.flip = false
	}

	if p.keymap.IsKeyJumpPressed() && !p.isJumping && !p.isFulling {
		p.action = "jump"
		p.vector.Y = -p.gravity
		p.isJumping = true
	}

	if p.keymap.IsKeyAttackPressed() {
		p.shootBullet()
	}

	if p.keymap.IsKeyThrowPressed() {
		p.throwGrenade = NewThrowGrenade(p.position, p.levelData, p.flip)
	}
}

func (p *Player) vectorHandle() {
	if p.currentAnimation == nil {
		return
	}
	tool.CollisionDetect(
		p.currentAnimation.GetRec(),
		p.levelData,
		p.vector,
		p.position,
		false,
	)
	if p.vector.Y == 0 {
		p.isFulling = false
	}
	p.position.X += p.vector.X
	p.position.Y += p.vector.Y
	p.counter++
	if p.isJumping || p.isFulling {
		if p.counter%3 == 0 && p.vector.Y < p.gravity {
			p.vector.Y += 1
		}
		if p.vector.Y > 0 {
			p.isJumping = false
			p.isFulling = true
		}
	}
}

func (p *Player) animationControl() {
	switch p.action {
	case "idle":
		p.currentAnimation = p.idle
		p.currentAnimation.Play()
	case "moveRight":
		p.currentAnimation = p.run
		p.currentAnimation.Play()
	case "fullRight":
		p.currentAnimation = p.fulling
		p.currentAnimation.PlayOnce()
	case "moveLeft":
		p.currentAnimation = p.run
		p.currentAnimation.Play()
	case "fullLeft":
		p.currentAnimation = p.fulling
		p.currentAnimation.PlayOnce()
	case "jump":
		p.currentAnimation = p.jump
		p.currentAnimation.PlayOnce()
	}
	p.currentAnimation.Flip = p.flip
}

func (p *Player) shootBullet() {
	p.shootFire.PlayOnce()
	var bulletVectory *image.Point
	var bulletPosition *tool.FPoint
	if p.flip {
		bulletVectory = &image.Point{X: -p.bulletSpeed}
		bulletPosition = &tool.FPoint{
			X: float64(p.position.X) - float64(p.bulletTexture.Bounds().Dx()),
			Y: float64(p.position.Y + p.idle.GetRec().Dy()/2 - 2),
		}
	} else {
		bulletVectory = &image.Point{X: p.bulletSpeed}
		bulletPosition = &tool.FPoint{
			X: float64(p.position.X + p.idle.GetRec().Dx()),
			Y: float64(p.position.Y + p.idle.GetRec().Dy()/2 - 2),
		}
	}
	newSprite := &texture.Sprite{
		Texture:  p.bulletTexture,
		Position: bulletPosition,
	}
	newBullet := texture.NewMotivationSprite(newSprite, 120, bulletVectory)
	p.bulletList = append(p.bulletList, newBullet)
}

func (p *Player) updateShootFire() {
	p.shootFire.Flip = p.flip
	var firePx int
	var firePy int
	firePx = p.idle.GetRec().Max.X
	firePy = p.position.Y + p.idle.GetRec().Dy()/2 - 3
	if p.flip {
		firePx = p.idle.GetRec().Min.X - p.shootFire.GetRec().Dx()
		firePy = p.position.Y + p.idle.GetRec().Dy()/2 - 3
	}
	// fix fire effect position
	if p.currentAnimation == p.run || p.currentAnimation == p.jump {
		firePy -= 2
	}
	p.shootFire.Update(&image.Point{X: firePx, Y: firePy})
}

func (p *Player) updateBullet() {
	deadBulletIndex := []int{}
	for index, bullet := range p.bulletList {
		tool.CollisionDetect(
			bullet.GetRec(),
			p.levelData,
			bullet.Vector,
			&image.Point{
				int(bullet.Position.X),
				int(bullet.Position.Y),
			},
			false,
		)
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
	p.vectorHandle()
	p.animationControl()
	p.currentAnimation.Update(p.position)
	p.updateBullet()
	p.updateShootFire()
	if p.throwGrenade != nil {
		p.throwGrenade.Update()
	}
}

func (p *Player) DrawBullet(screen *ebiten.Image) {
	for _, bullet := range p.bulletList {
		bullet.Draw(screen)
		if bullet.Vector.X == 0 {
			bullet.Kill()
		}
	}
}

func (p *Player) DrawShootFire(screen *ebiten.Image) {
	if len(p.bulletList) > 0 {
		p.shootFire.Draw(screen)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.currentAnimation.Draw(screen)
	p.DrawBullet(screen)
	p.DrawShootFire(screen)
	if p.throwGrenade != nil {
		p.throwGrenade.Draw(screen)
	}
}

func (p *Player) Dispose() {
	p.idle.Dispose()
	p.death.Dispose()
	p.run.Dispose()
	p.jump.Dispose()
	p.bulletTexture.Dispose()
}
