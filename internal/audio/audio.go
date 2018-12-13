package audio

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/eihigh/love-and-hate/internal/assets"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

var (
	bgmPlayers = map[string]*audio.Player{}
	sePlayers  = map[string]*audio.Player{}

	audioContext *audio.Context

	BgmVolume = 0.4
	SeVolume  = 0.1
)

func init() {
	var err error
	audioContext, err = audio.NewContext(44100)
	if err != nil {
		log.Fatal(err)
	}
}

func Load() {

	// BGMs
	for _, name := range []string{
		"Retrospect",
		"Scent_of_flowers",
	} {
		f, err := assets.Assets.Open(fmt.Sprintf("audio/%s.mp3", name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
		PushBgm(name, b)
	}

	// SEs
	for _, name := range []string{
		"buzzer",
		"drop",
	} {
		f, err := assets.Assets.Open(fmt.Sprintf("audio/%s.wav", name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
		PushSe(name, b)
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

func PushBgm(name string, b []byte) {
	f := audio.BytesReadSeekCloser(b)
	stream, err := mp3.Decode(audioContext, f)
	if err != nil {
		panic(err)
	}

	p, err := audio.NewPlayer(audioContext, stream)
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
	p.SetVolume(BgmVolume)
	if err := p.Rewind(); err == nil {
		p.Play()
	}
}

// SE
func PushSe(name string, b []byte) {
	f := audio.BytesReadSeekCloser(b)
	stream, err := wav.Decode(audioContext, f)
	if err != nil {
		panic(err)
	}
	p, err := audio.NewPlayer(audioContext, stream)
	if err != nil {
		panic(err)
	}
	sePlayers[name] = p
}

func SetSeVolume(volume float64) {
	for _, p := range sePlayers {
		p.SetVolume(volume)
	}
}

func PlaySe(name string) {
	p := sePlayers[name]
	p.SetVolume(SeVolume)
	p.Rewind()
	p.Play()
}
