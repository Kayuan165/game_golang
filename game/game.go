package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player *Player
}

func NewGame() *Game {
	player := NewPlayer()
	return &Game{
		player: player,
	}
}

// reponsavel por atualizar a logica do jogo
func (g *Game) Update() error {
	return nil
}

// responsavel por desenhar o jogo
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

}

// responsavel por definir o tamanho da tela
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
