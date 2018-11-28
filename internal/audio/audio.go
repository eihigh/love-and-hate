package audio

import (
	"log"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
)

var (
	bgmPlayers = map[string]*audio.Player{}
	sePlayers  = map[string]*audio.Player{}

	audioContext *audio.Context
)

func init() {
	var err error
	audioContext, err = audio.NewContext(44100)
	if err != nil {
		log.Fatal(err)
	}
}

func Finalize() {
	for _, p := range bgmPlayers {
		p.Close()
	}
	for _, p := range sePlayers {
		p.Close()
	}
}

func AddBgm(name string, b []byte) {
	f := audio.BytesReadSeekCloser(b)
	stream, err := mp3.Decode(audioContext, f)
	if err != nil {
		panic(err)
	}

	src := audio.NewInfiniteLoop(stream, stream.Length())
	p, err := audio.NewPlayer(audioContext, src)
	if err != nil {
		panic(err)
	}
	bgmPlayers[name] = p
}

func SetBgmVolume(volume float64) {
	for _, p := range bgmPlayers {
		p.SetVolume(volume)
	}
}

func PauseBgm() {
	for _, p := range bgmPlayers {
		p.Pause()
	}
}

func PlayBgm(name string) {
	p := bgmPlayers[name]
	if err := p.Rewind(); err == nil {
		p.Play()
	}
}
