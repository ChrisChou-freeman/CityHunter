package lib

import(
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/inpututil"
)

type KeyMap struct{}

func(k *KeyMap)IsKeyUpPressed() bool{
  return inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyUp)
}

func(k *KeyMap)IsKeyDownPressed()bool{
  return inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyDown)
}

func(k *KeyMap)IsKeyEnterPressed()bool{
  return inpututil.IsKeyJustPressed(ebiten.KeyEnter)
}

func(k *KeyMap)IsKeyBackPressed()bool{
  return inpututil.IsKeyJustPressed(ebiten.KeyEscape)
}
