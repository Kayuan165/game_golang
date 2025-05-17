package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
}

//reponsavel por atualizar a logica do jogo
func (g *Game) Update() error {
	return nil
}

//responsavel por desenhar o jogo
func (g *Game) Draw(screen *ebiten.Image) {

}

//responsavel por definir o tamanho da tela
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	g := &Game{}

	err := ebiten.RunGame(g)

	if err != nil {
		panic(err)
	}
}
