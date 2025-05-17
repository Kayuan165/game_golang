package game

import (
	"my-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image             *ebiten.Image
	position          Vector
	game              *Game
	laserLoadingTimer *Timer
}

func NewPlayer(game *Game) *Player {

	image := assets.PlayerSprite
	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := Vector{
		x: (screenWidth / 2) - halfW,
		y: 500,
	}

	return &Player{
		image:             image,
		position:          position,
		game:              game,
		laserLoadingTimer: NewTimer(12),
	}
}

func (p *Player) Update() {

	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.x -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.x += speed
	}

	p.laserLoadingTimer.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTimer.IsReady() {
		p.laserLoadingTimer.Reset()
		bounds := p.image.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		spawnPos := Vector{
			p.position.x + halfW,
			p.position.y - halfH/2,
		}

		laser := NewLaser(spawnPos)
		p.game.Addlaser(laser)

	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//posição x e y que a  imagem vai ser desenhada
	op.GeoM.Translate(p.position.x, p.position.y)
	//desenha a imagem na tela
	screen.DrawImage(p.image, op)
}
