package melee

type Stage byte

const (
	NO_STAGE           Stage = 0
	FINAL_DESTINATION        = 0x19
	BATTLEFIELD              = 0x18
	POKEMON_STADIUM          = 0x12
	DREAMLAND                = 0x1A
	FOUNTAIN_OF_DREAMS       = 0x8
	YOSHIS_STORY             = 0x6
	RANDOM_STAGE             = 0x1D
)

func GetStageName(stage Stage) string {
	switch stage {
	case NO_STAGE:
		return "NO_STAGE"
	case FINAL_DESTINATION:
		return "FINAL_DESTINATION"
	case BATTLEFIELD:
		return "BATTLEFIELD"
	case POKEMON_STADIUM:
		return "POKEMON_STADIUM"
	case DREAMLAND:
		return "DREAMLAND"
	case FOUNTAIN_OF_DREAMS:
		return "FOUNTAIN_OF_DREAMS"
	case YOSHIS_STORY:
		return "YOSHIS_STORY"
	case RANDOM_STAGE:
		return "RANDOM_STAGE"

	// TODO: Other stages? Do we care?
	default:
		return ""
	}
}

// Get the X coordinate of the edge of the stage, approaching from
// off_stage
// IE: This is your X coordinate when hanging on the edge
// NOTE: The left edge is always the same, but negative"""
func EdgePosition(stage Stage) float32 {
	if stage == BATTLEFIELD {
		return 71.3078536987
	}
	if stage == FINAL_DESTINATION {
		return 88.4735488892
	}
	if stage == DREAMLAND {
		return 80.1791534424
	}
	if stage == FOUNTAIN_OF_DREAMS {
		return 66.2554016113
	}
	if stage == POKEMON_STADIUM {
		return 90.657852
	}
	if stage == YOSHIS_STORY {
		return 58.907848
	}

	return 1000
}

// Get the X coordinate of the edge of the stage, while standing on the
// stage
// IE: This is your X coordinate when teetering on the edge
// NOTE: The left edge is always the same, but negative"""
func EdgeGroundPosition(stage Stage) float32 {
	if stage == BATTLEFIELD {
		return 68.4000015259
	}
	if stage == FINAL_DESTINATION {
		return 85.5656967163
	}
	if stage == DREAMLAND {
		return 77.2713012695
	}
	if stage == FOUNTAIN_OF_DREAMS {
		return 63.3475494385
	}
	if stage == POKEMON_STADIUM {
		return 87.75
	}
	if stage == YOSHIS_STORY {
		return 56
	}

	return 1000
}
