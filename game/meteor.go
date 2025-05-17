package game

import (
	"math/rand"
	"my-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Meteor struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func NewMeteor() *Meteor {

	image := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]
	speed := (rand.Float64() * 13)

	position := Vector{
		x: rand.Float64() * screenWidth,
		y: -100,
	}

	return &Meteor{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (m *Meteor) Update() {
	m.position.y += m.speed
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//posição x e y que a  imagem vai ser desenhada
	op.GeoM.Translate(m.position.x, m.position.y)
	//desenha a imagem na tela
	screen.DrawImage(m.image, op)
}
