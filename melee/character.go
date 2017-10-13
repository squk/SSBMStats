package melee

type Character uint32

const (
	DOC               Character = 0x00
	MARIO                       = 0x01
	LUIGI                       = 0x02
	BOWSER                      = 0x03
	PEACH                       = 0x04
	YOSHI                       = 0x05
	DK                          = 0x06
	CPTFALCON                   = 0x07
	GANONDORF                   = 0x08
	FALCO                       = 0x09
	FOX                         = 0x0a
	NESS                        = 0x0b
	ICECLIMBERS                 = 0x0c
	KIRBY                       = 0x0d
	SAMUS                       = 0x0e
	ZELDA                       = 0x0f
	LINK                        = 0x10
	YLINK                       = 0x11
	PICHU                       = 0x12
	PIKACHU                     = 0x13
	JIGGLYPUFF                  = 0x14
	MEWTWO                      = 0x15
	GAMEANDWATCH                = 0x16
	MARTH                       = 0x17
	ROY                         = 0x18
	SHEIK                       = 0x42 // not an actual value in Melee.
	UNKNOWN_CHARACTER           = 0xff
)

var CHARACTER_NAMES = map[Character]string{
	DOC:               "DOC",
	MARIO:             "MARIO",
	LUIGI:             "LUIGI",
	BOWSER:            "BOWSER",
	PEACH:             "PEACH",
	YOSHI:             "YOSHI",
	DK:                "DK",
	CPTFALCON:         "CPTFALCON",
	GANONDORF:         "GANONDORF",
	FALCO:             "FALCO",
	FOX:               "FOX",
	NESS:              "NESS",
	ICECLIMBERS:       "ICECLIMBERS",
	KIRBY:             "KIRBY",
	SAMUS:             "SAMUS",
	ZELDA:             "ZELDA",
	LINK:              "LINK",
	YLINK:             "YLINK",
	PICHU:             "PICHU",
	PIKACHU:           "PIKACHU",
	JIGGLYPUFF:        "JIGGLYPUFF",
	MEWTWO:            "MEWTWO",
	GAMEANDWATCH:      "GAMEANDWATCH",
	MARTH:             "MARTH",
	ROY:               "ROY",
	SHEIK:             "SHEIK",
	UNKNOWN_CHARACTER: "UNKNOWN_CHARACTER",
}
