package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player *Player
	lasers []*Laser
}

func NewGame() *Game {
	g := &Game{}
	player := NewPlayer(g)
	g.player = player
	return g
}

//roda em 60 fps
// reponsavel por atualizar a logica do jogo
// 60 vezes por segundo
//1 tik = 1 frame
func (g *Game) Update() error {
	g.player.Update()

	for _, l := range g.lasers {
		l.Update()
	}
	return nil
}

// responsavel por desenhar o jogo
//roda 60 vezes por segundo
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	for _, l := range g.lasers {
		l.Draw(screen)
	}

}

// responsavel por definir o tamanho da tela
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Addlaser(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}
