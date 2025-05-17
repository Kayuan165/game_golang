package game

import (
	"my-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image *ebiten.Image
}

func NewPlayer() *Player {

	image := assets.PlayerSprite
	return &Player{
		image: image,
	}
}

func (p *Player) Update() {

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//posição x e y que a  imagem vai ser desenhada
	op.GeoM.Translate(100, 100)
	//desenha a imagem na tela
	screen.DrawImage(p.image, op)
}
