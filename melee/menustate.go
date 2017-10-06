package melee

type MenuState byte

const (
	CHARACTER_SELECT MenuState = 0
	STAGE_SELECT               = 1
	IN_GAME                    = 2
	POSTGAME_SCORES            = 4
)

func GetMenuStateName(state MenuState) string {
	switch state {
	case CHARACTER_SELECT:
		return "CHARACTER_SELECT"
	case STAGE_SELECT:
		return "STAGE_SELECT"
	case IN_GAME:
		return "IN_GAME"
	case POSTGAME_SCORES:
		return "POSTGAME_SCORES"
	default:
		return ""
	}
}
