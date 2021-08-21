package pearl


import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"

	"os"
)


func PlaySoundFromFile(path string) {
	sfx, _ := os.Open(path)
	decoded, _ := wav.Decode(g.audioContext, sfx)
	player, _ := audio.NewPlayer(g.audioContext, decoded)

	player.Rewind()
	player.Play()
}