package lib

import(
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/inpututil"
)

type KeyMap struct{}

var MouseLeftInUsing bool

// key up
func(k *KeyMap)IsKeyUpPressed() bool{
  return inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyUp)
}

// key down
func(k *KeyMap)IsKeyDownPressed()bool{
  return inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyDown)
}

// key enter
func(k *KeyMap)IsKeyEnterPressed()bool{
  return inpututil.IsKeyJustPressed(ebiten.KeyEnter)
}

// key back
func(k *KeyMap)IsKeyBackPressed()bool{
  return inpututil.IsKeyJustPressed(ebiten.KeyEscape)
}

// key left
func(k *KeyMap)IsKeyLeftPressed()bool{
  return inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyLeft)
}

func(k *KeyMap)IsKeyLeftHoldPressed()bool{
  return ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) 
}

// key right
func(k *KeyMap)IsKeyRightPressed()bool{
  return inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyRight)
}

func(k *KeyMap)IsKeyRightHoldPressed()bool{
  return ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) 
}

// mouse left
func(k *KeyMap)IsMouseLeftKeyPressed()bool{
  return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func(k *KeyMap)IsMouseRightKeyPressed()bool{
  return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight)
}
