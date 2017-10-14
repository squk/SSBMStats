package melee

/*
* Any function that needs to be performed each frame should exist here. The only
* function call resulting in a write to disk that should ever happen here is
* FWriter.LogFrame(). See framewriter.go and writecondition.go for more details.
 */

func (g *MeleeState) OnFrame() {
	frame := NewFrame(g.Players)

	if g.MenuState == IN_GAME {
		go FWriter.LogFrame(frame)
	}

	// Below should  always be called after frame watching funcs
	g.LastMenuState = g.MenuState
}

func (g *MeleeState) OnCharacter() {
	g.AssignSelfPort()
	g.InferOpponentPorts()
}

func (g *MeleeState) AssignSelfPort() {
	FWriter.Match.SelfCharacter = g.Players[g.SelfPort].GetCharacterString()
}

// Loop through all Players that aren't us and mark them as opponent if they've
// chosen a character
func (g *MeleeState) InferOpponentPorts() {
	FWriter.Match.OpponentCharacters = []string{}
	FWriter.Match.OpponentPorts = PortList{}

	for i := 1; i < 5; i++ {
		if i != g.SelfPort {
			player := g.Players[i]

			if char, ok := player.GetCharacter(); ok == nil {
				if char != UNKNOWN_CHARACTER {
					g.OpponentPorts = append(g.OpponentPorts, i)

					chars := &FWriter.Match.OpponentCharacters
					*chars = append(*chars, CHARACTER_NAMES[char])

					ports := &FWriter.Match.OpponentPorts
					*ports = append(*ports, Port(i))
				}
			}
		}
	}
}
