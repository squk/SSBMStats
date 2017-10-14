package melee

type StateCheck func(Player) bool

type Flag int

const (
	NEUTRAL_FLAG Flag = iota
)

// More than 2 players are currently in neutral
func IN_IMMEDIATE_NEUTRAL(p Player) bool {
	if a, ok := p.GetPlayerAction(); ok == nil {
		if a.IsNeutral() {
			return true
		}
	}

	return false
}

func DEAD(p Player) bool {
	if a, ok := p.GetPlayerAction(); ok == nil {
		if a.IsDead() {
			return true
		}
	}

	return false
}

func HIT(p Player) bool {
	if a, ok := p.GetPlayerAction(); ok == nil {
		if a.IsDamage() {
			return true
		}
	}

	return false
}

// Approaching is analogous to attacking for now.
func ATTACKING(p Player) bool {
	if a, ok := p.GetPlayerAction(); ok == nil {
		if a.IsAttack() {
			return true
		}
	}

	return false
}

func (f Frame) PortsInState(CHECK StateCheck) PortList {
	ports := PortList{}

	// TODO: Make a proper iterator for the ports we care about...
	// For now looping 1-4
	for p := 1; p < 5; p++ {
		player := f.Players[p]

		if CHECK(player) {
			ports = append(ports, Port(p))
		}
	}

	return ports
}

func (f Frame) SelfIs(CHECK StateCheck) bool {
	return CHECK(f.Players[FWriter.GetSelfPort()])
}

func (f Frame) OpponentIs(CHECK StateCheck) bool {
	for _, o := range FWriter.GetOpponentPorts() {
		if CHECK(f.Players[o]) {
			return true
		}
	}

	return false
}
