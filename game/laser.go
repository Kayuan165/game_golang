package game

import (
	"my-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	image    *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {

	image := assets.LaserSprite
	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2 //metade da largura da imagem do laser
	halfH := float64(bounds.Dy()) / 2 //metade da altura da imagem do laser

	position.x -= halfW //ajusta a posição x do laser para o centro
	position.y -= halfH //ajusta a posição y do laser para o centro

	return &Laser{
		image:    image,
		position: position,
	}
}

func (l *Laser) Update() {
	speed := 7.0

	l.position.y -= speed //move o laser para cima
}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//posição x e y que a  imagem vai ser desenhada
	op.GeoM.Translate(l.position.x, l.position.y)
	//desenha a imagem na tela
	screen.DrawImage(l.image, op)
}
