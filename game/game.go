package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player           *Player
	lasers           []*Laser
	meteors          []*Meteor
	meteorSpawnTimer *Timer
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(24),
	}
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

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
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

	for _, m := range g.meteors {
		m.Draw(screen)
	}

}

// responsavel por definir o tamanho da tela
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Addlaser(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}
