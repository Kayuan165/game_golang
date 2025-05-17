package game

import (
	"my-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image    *ebiten.Image
	position Vector
}

func NewPlayer() *Player {

	image := assets.PlayerSprite
	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := Vector{
		x: (screenWidth / 2) - halfW,
		y: 500,
	}

	return &Player{
		image:    image,
		position: position,
	}
}

func (p *Player) Update() {

	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.x -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.x += speed
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//posição x e y que a  imagem vai ser desenhada
	op.GeoM.Translate(p.position.x, p.position.y)
	//desenha a imagem na tela
	screen.DrawImage(p.image, op)
}
