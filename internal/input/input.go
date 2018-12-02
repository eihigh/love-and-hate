package input

import "github.com/hajimehoshi/ebiten"

func OnKeys(keys []ebiten.Key) bool {
	for _, k := range keys {
		if ebiten.IsKeyPressed(k) {
			return true
		}
	}
	return false
}

func OnUp() bool {
	keys := []ebiten.Key{
		ebiten.KeyUp,
		ebiten.KeyR,
	}
	return OnKeys(keys)
}

func OnDown() bool {
	keys := []ebiten.Key{
		ebiten.KeyDown,
		ebiten.KeyT,
	}
	return OnKeys(keys)
}

func OnLeft() bool {
	keys := []ebiten.Key{
		ebiten.KeyLeft,
		ebiten.KeyH,
	}
	return OnKeys(keys)
}

func OnRight() bool {
	keys := []ebiten.Key{
		ebiten.KeyRight,
		ebiten.KeyS,
	}
	return OnKeys(keys)
}

func OnDecide() bool {
	keys := []ebiten.Key{
		ebiten.KeyEnter,
		ebiten.KeyZ,
		ebiten.KeySpace,
	}
	return OnKeys(keys)
}

func OnQuit() bool {
	keys := []ebiten.Key{ebiten.KeyQ}
	return OnKeys(keys)
}

func OnReset() bool {
	keys := []ebiten.Key{ebiten.KeyEscape}
	return OnKeys(keys)
}
