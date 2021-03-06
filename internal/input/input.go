package input

import "github.com/hajimehoshi/ebiten"

type strset map[string]struct{}

var (
	curs  = strset{}
	lasts = strset{}
)

func Reset() {
	lasts = curs
	curs = strset{}
}

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
	if on := OnKeys(keys); on {
		curs["decide"] = struct{}{}
		return true
	}
	return false
}

func JustDecided() bool {
	if on := OnDecide(); on {
		curs["decide"] = struct{}{}
		if _, last := lasts["decide"]; !last {
			return true
		}
	}
	return false
}

func OnCancel() bool {
	keys := []ebiten.Key{
		ebiten.KeyEscape,
		ebiten.KeyX,
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
