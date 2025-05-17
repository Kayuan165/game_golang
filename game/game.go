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

//roda em 60 fps
// reponsavel por atualizar a logica do jogo
// 60 vezes por segundo
func (g *Game) Update() error {
	g.player.Update()
	return nil
}

// responsavel por desenhar o jogo
//roda 60 vezes por segundo
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

}

// responsavel por definir o tamanho da tela
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
