package main

import eb "github.com/hajimehoshi/ebiten"

func onKeys(keys []eb.Key) bool {
	for _, k := range keys {
		if eb.IsKeyPressed(k) {
			return true
		}
	}
	return false
}

func onUp() bool {
	keys := []eb.Key{
		eb.KeyUp,
		eb.KeyR,
	}
	return onKeys(keys)
}

func onDown() bool {
	keys := []eb.Key{
		eb.KeyDown,
		eb.KeyT,
	}
	return onKeys(keys)
}

func onLeft() bool {
	keys := []eb.Key{
		eb.KeyLeft,
		eb.KeyH,
	}
	return onKeys(keys)
}

func onRight() bool {
	keys := []eb.Key{
		eb.KeyRight,
		eb.KeyS,
	}
	return onKeys(keys)
}

func onQuit() bool {
	keys := []eb.Key{eb.KeyQ}
	return onKeys(keys)
}
