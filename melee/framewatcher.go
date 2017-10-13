package melee

func (g *MeleeState) OnFrame() {
	player_container := g.Players

	frame := NewFrame(player_container)
	if g.MenuState == IN_GAME {
		go FWriter.LogFrame(frame)
	}

	g.InferOpponentPorts()
}

func (g *MeleeState) InferOpponentPorts() {

}
