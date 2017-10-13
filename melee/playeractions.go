package melee

type Button string

const (
	BUTTON_A       Button = "A"
	BUTTON_B              = "B"
	BUTTON_X              = "X"
	BUTTON_Y              = "Y"
	BUTTON_Z              = "Z"
	BUTTON_L              = "L"
	BUTTON_R              = "R"
	BUTTON_START          = "START"
	BUTTON_D_UP           = "D_UP"
	BUTTON_D_DOWN         = "D_DOWN"
	BUTTON_D_LEFT         = "D_LEFT"
	BUTTON_D_RIGHT        = "D_RIGHT"
	// Control sticks considered "buttons" here
	BUTTON_MAIN = "MAIN"
	BUTTON_C    = "C"
)

type PlayerAction uint32

const (
	DEAD_DOWN              PlayerAction = 0x0
	DEAD_LEFT                           = 0x1
	DEAD_RIGHT                          = 0x2
	DEAD_FLY_STAR                       = 0x4
	DEAD_FLY                            = 0x6 // When you have been hit upwards and are dead
	DEAD_FLY_SPLATTER                   = 0x7 // Hit upwards and have splattered on the camera
	DEAD_FLY_SPLATTER_FLAT              = 0x8 // Hit upwards and have splattered on the camera
	ON_HALO_DESCENT                     = 0xc
	ON_HALO_WAIT                        = 0x0d

	STANDING    = 0x0e
	WALK_SLOW   = 0x0f
	WALK_MIDDLE = 0x10
	WALK_FAST   = 0x11
	TURNING     = 0x12
	TURNING_RUN = 0x13
	DASHING     = 0x14
	RUNNING     = 0x15
	RUN_DIRECT  = 0x16
	RUN_BRAKE   = 0x17

	KNEE_BEND              = 0x18 // pre-jump animation.
	JUMPING_FORWARD        = 0x19
	JUMPING_BACKWARD       = 0x1A
	JUMPING_ARIAL_FORWARD  = 0x1b
	JUMPING_ARIAL_BACKWARD = 0x1c
	FALLING                = 0x1D // The "wait" state of the air.
	FALLING_AERIAL         = 0x20 // After double-jump
	DEAD_FALL              = 0x23 // Falling after up-b
	SPECIAL_FALL_FORWARD   = 0x24
	SPECIAL_FALL_BACK      = 0x25
	TUMBLING               = 0x26

	CROUCH_START    = 0x27 // Going from stand to crouch
	CROUCHING       = 0x28
	CROUCH_END      = 0x29 // Standing up from crouch
	LANDING         = 0x2a // Can be canceled. Not stunned
	LANDING_SPECIAL = 0x2b // Landing special like from wavedash. Stunned.

	NEUTRAL_ATTACK_1      = 0x2c
	NEUTRAL_ATTACK_2      = 0x2d
	NEUTRAL_ATTACK_3      = 0x2e
	LOOPING_ATTACK_START  = 0x2f
	LOOPING_ATTACK_MIDDLE = 0x30
	LOOPING_ATTACK_END    = 0x31
	DASH_ATTACK           = 0x32

	FTILT_HIGH     = 0x33
	FTILT_HIGH_MID = 0x34
	FTILT_MID      = 0x35
	FTILT_LOW_MID  = 0x36
	FTILT_LOW      = 0x37
	UPTILT         = 0x38
	DOWNTILT       = 0x39

	FSMASH_HIGH     = 0x3a
	FSMASH_MID_HIGH = 0x3b
	FSMASH_MID      = 0x3c
	FSMASH_MID_LOW  = 0x3d
	FSMASH_LOW      = 0x3e
	UPSMASH         = 0x3f
	DOWNSMASH       = 0x40

	NAIR         = 0x41
	FAIR         = 0x42
	BAIR         = 0x43
	UAIR         = 0x44
	DAIR         = 0x45
	NAIR_LANDING = 0x46
	FAIR_LANDING = 0x47
	BAIR_LANDING = 0x48
	UAIR_LANDING = 0x49
	DAIR_LANDING = 0x4a

	DAMAGE_HIGH_1      = 0x4b
	DAMAGE_HIGH_2      = 0x4c
	DAMAGE_HIGH_3      = 0x4d
	DAMAGE_NEUTRAL_1   = 0x4e
	DAMAGE_NEUTRAL_2   = 0x4f
	DAMAGE_NEUTRAL_3   = 0x50
	DAMAGE_LOW_1       = 0x51
	DAMAGE_LOW_2       = 0x52
	DAMAGE_LOW_3       = 0x53
	DAMAGE_AIR_1       = 0x54
	DAMAGE_AIR_2       = 0x55
	DAMAGE_AIR_3       = 0x56
	DAMAGE_FLY_HIGH    = 0x57
	DAMAGE_FLY_NEUTRAL = 0x58
	DAMAGE_FLY_LOW     = 0x59
	DAMAGE_FLY_TOP     = 0x5a
	DAMAGE_FLY_ROLL    = 0x5b

	ITEM_PICKUP_LIGHT                  = 0x5C
	ITEM_PICKUP_HEAVY                  = 0x5D
	ITEM_THROW_LIGHT_FORWARD           = 0x5E
	ITEM_THROW_LIGHT_BACK              = 0x5F
	ITEM_THROW_LIGHT_HIGH              = 0x60
	ITEM_THROW_LIGHT_LOW               = 0x61
	ITEM_THROW_LIGHT_DASH              = 0x62
	ITEM_THROW_LIGHT_DROP              = 0x63
	ITEM_THROW_LIGHT_AIR_FORWARD       = 0x64
	ITEM_THROW_LIGHT_AIR_BACK          = 0x65
	ITEM_THROW_LIGHT_AIR_HIGH          = 0x66
	ITEM_THROW_LIGHT_AIR_LOW           = 0x67
	ITEM_THROW_HEAVY_FORWARD           = 0x68
	ITEM_THROW_HEAVY_BACK              = 0x69
	ITEM_THROW_HEAVY_HIGH              = 0x6A
	ITEM_THROW_HEAVY_LOW               = 0x6B
	ITEM_THROW_LIGHT_SMASH_FORWARD     = 0x6C
	ITEM_THROW_LIGHT_SMASH_BACK        = 0x6D
	ITEM_THROW_LIGHT_SMASH_UP          = 0x6e
	ITEM_THROW_LIGHT_SMASH_DOWN        = 0x6F
	ITEM_THROW_LIGHT_AIR_SMASH_FORWARD = 0x70
	ITEM_THROW_LIGHT_AIR_SMASH_BACK    = 0x71
	ITEM_THROW_LIGHT_AIR_SMASH_HIGH    = 0x72
	ITEM_THROW_LIGHT_AIR_SMASH_LOW     = 0x73
	ITEM_THROW_HEAVY_AIR_SMASH_FORWARD = 0x74
	ITEM_THROW_HEAVY_AIR_SMASH_BACK    = 0x75
	ITEM_THROW_HEAVY_AIR_SMASH_HIGH    = 0x76
	ITEM_THROW_HEAVY_AIR_SMASH_LOW     = 0x77

	BEAM_SWORD_SWING_1 = 0x78
	BEAM_SWORD_SWING_2 = 0x79
	BEAM_SWORD_SWING_3 = 0x7A
	BEAM_SWORD_SWING_4 = 0x7B

	SHIELD_START   = 0xb2
	SHIELD         = 0xb3
	SHIELD_RELEASE = 0xb4
	SHIELD_STUN    = 0xb5
	SHIELD_REFLECT = 0xb6

	TECH_MISS_UP              = 0xb7 //  "facing" up. Not important to us
	LYING_GROUND_UP           = 0xb8
	LYING_GROUND_UP_HIT       = 0xb9
	GROUND_GETUP              = 0xba
	GROUND_ATTACK_UP          = 0xbb
	GROUND_ROLL_FORWARD_UP    = 0xbc
	GROUND_ROLL_BACKWARD_UP   = 0xbd
	TECH_MISS_DOWN            = 0xbf
	LYING_GROUND_DOWN         = 0xc0
	DAMAGE_GROUND             = 0xc1
	NEUTRAL_GETUP             = 0xc2
	GETUP_ATTACK              = 0xc3
	GROUND_ROLL_FORWARD_DOWN  = 0xc4
	GROUND_ROLL_BACKWARD_DOWN = 0xc5

	NEUTRAL_TECH   = 0xc7
	FORWARD_TECH   = 0xc8
	BACKWARD_TECH  = 0xc9
	WALL_TECH      = 0xca
	WALL_TECH_JUMP = 0xcb
	CEILING_TECH   = 0xcc

	SHIELD_BREAK_FLY     = 0xcd
	SHIELD_BREAK_FALL    = 0xce
	SHIELD_BREAK_DOWN_U  = 0xcf
	SHIELD_BREAK_DOWN_D  = 0xd0
	SHIELD_BREAK_STAND_U = 0xd1
	SHIELD_BREAK_STAND_D = 0xd2
	SHIELD_BREAK_TEETER  = 0xd3

	GRAB              = 0xd4
	GRAB_PULLING      = 0xd5
	GRAB_RUNNING      = 0xd6
	GRAB_WAIT         = 0xd8
	GRAB_PUMMEL       = 0xd9
	GRAB_BREAK        = 0xda
	THROW_FORWARD     = 0xdb
	THROW_BACK        = 0xdc
	THROW_UP          = 0xdd // yuck
	THROW_DOWN        = 0xde
	GRAB_PULLING_HIGH = 0xdf
	GRABBED_WAIT_HIGH = 0xe0 // XXX Not sure about this
	PUMMELED_HIGH     = 0xe1 // XXX Not sure about this
	GRAB_PULL         = 0xe2 // Being pulled inwards from the grab
	GRABBED           = 0xe3 // Grabbed
	GRAB_PUMMELED     = 0xe4 // Being pummeled
	GRAB_ESCAPE       = 0xe5

	ROLL_FORWARD  = 0xe9
	ROLL_BACKWARD = 0xea
	SPOTDODGE     = 0xEB
	AIRDODGE      = 0xEC

	THROWN_FORWARD = 0xEF
	THROWN_BACK    = 0xF0
	THROWN_UP      = 0xF1
	THROWN_DOWN    = 0xF2
	THROWN_DOWN_2  = 0xf3

	PLATFORM_DROP        = 0xf4
	EDGE_TEETERING_START = 0xF5 // Starting of edge teetering
	EDGE_TEETERING       = 0xF6
	BOUNCE_WALL          = 0xf7
	BOUNCE_CEILING       = 0xf8
	BUMP_WALL            = 0xf9
	BUMP_CIELING         = 0xfa
	SLIDING_OFF_EDGE     = 0xfb // When you get hit and slide off an edge
	EDGE_CATCHING        = 0xFC // Initial grabbing of edge stuck in stun here
	EDGE_HANGING         = 0xFD
	EDGE_GETUP_SLOW      = 0xFE  //  >= 100% damage
	EDGE_GETUP_QUICK     = 0xFF  //  < 100% damage
	EDGE_ATTACK_SLOW     = 0x100 //  < 100% damage
	EDGE_ATTACK_QUICK    = 0x101 //  >= 100% damage
	EDGE_ROLL_SLOW       = 0x102 //  >= 100% damage
	EDGE_ROLL_QUICK      = 0x103 //  < 100% damage
	EDGE_JUMP_1_SLOW     = 0x104
	EDGE_JUMP_2_SLOW     = 0x105
	EDGE_JUMP_1_QUICK    = 0x106
	EDGE_JUMP_2_QUICK    = 0x107

	TAUNT_RIGHT = 0x108
	TAUNT_LEFT  = 0x109

	ENTRY                     = 0x142 // Start of match. Can't move
	ENTRY_START               = 0x143 // Start of match. Can't move
	ENTRY_END                 = 0x144 // Start of match. Can't move
	LASER_GUN_PULL            = 0x155
	NEUTRAL_B_CHARGING        = 0x156
	NEUTRAL_B_ATTACKING       = 0x157
	NEUTRAL_B_FULL_CHARGE     = 0x158
	WAIT_ITEM                 = 0x159 // No idea what this is
	NEUTRAL_B_CHARGING_AIR    = 0x15A
	NEUTRAL_B_ATTACKING_AIR   = 0x15B
	NEUTRAL_B_FULL_CHARGE_AIR = 0x15C

	//SWORD_DANCE_1          = 0x15d
	//SWORD_DANCE_2_HIGH     = 0x15e
	//SWORD_DANCE_2_MID      = 0x15f
	//SWORD_DANCE_3_HIGH     = 0x160
	//SWORD_DANCE_3_MID      = 0x161
	//SWORD_DANCE_3_LOW      = 0x162
	//SWORD_DANCE_4_HIGH     = 0x163
	//SWORD_DANCE_4_MID      = 0x164
	//SWORD_DANCE_4_LOW      = 0x165
	//SWORD_DANCE_1_AIR      = 0x166
	//SWORD_DANCE_2_HIGH_AIR = 0x167
	//SWORD_DANCE_2_MID_AIR  = 0x168
	//SWORD_DANCE_3_HIGH_AIR = 0x169
	//SWORD_DANCE_3_MID_AIR  = 0x16a
	//SWORD_DANCE_3_LOW_AIR  = 0x16b
	//SWORD_DANCE_4_HIGH_AIR = 0x16c
	//SWORD_DANCE_4_MID_AIR  = 0x16d
	//SWORD_DANCE_4_LOW_AIR  = 0x16e

	FOX_ILLUSION_START     = 0x15e
	FOX_ILLUSION           = 0x15f
	FOX_ILLUSION_SHORTENED = 0x160

	FIREFOX_WAIT_GROUND = 0x161 // Firefox wait on the ground
	FIREFOX_WAIT_AIR    = 0x162 // Firefox wait in the air
	FIREFOX_GROUND      = 0x163 // Firefox on the ground
	FIREFOX_AIR         = 0x164 // Firefox in the air

	DOWN_B_GROUND_START = 0x168
	DOWN_B_GROUND       = 0x169
	SHINE_TURN          = 0x16c
	DOWN_B_STUN         = 0x16d // Fox is stunned in these frames
	DOWN_B_AIR          = 0x16e

	UP_B_GROUND       = 0x16f
	SHINE_RELEASE_AIR = 0x170
	UP_B_AIR          = 0x170 // The upswing of the UP-B. (At least for marth)

	MARTH_COUNTER         = 0x171
	PARASOL_FALLING       = 0x172
	MARTH_COUNTER_FALLING = 0x173

	//NESS_SHEILD_START   = 0x174
	NESS_SHEILD         = 0x174
	NESS_SHEILD_AIR     = 0x175
	NESS_SHEILD_AIR_END = 0x177

	UNKNOWN_ANIMATION = 0xffff
)

var ACTION_NAMES = map[PlayerAction]string{
	DEAD_DOWN:                          "DEAD_DOWN",
	DEAD_LEFT:                          "DEAD_LEFT",
	DEAD_RIGHT:                         "DEAD_RIGHT",
	DEAD_FLY_STAR:                      "DEAD_FLY_STAR",
	DEAD_FLY:                           "DEAD_FLY",
	DEAD_FLY_SPLATTER:                  "DEAD_FLY_SPLATTER",
	DEAD_FLY_SPLATTER_FLAT:             "DEAD_FLY_SPLATTER_FLAT",
	ON_HALO_DESCENT:                    "ON_HALO_DESCENT",
	ON_HALO_WAIT:                       "ON_HALO_WAIT",
	STANDING:                           "STANDING",
	WALK_SLOW:                          "WALK_SLOW",
	WALK_MIDDLE:                        "WALK_MIDDLE",
	WALK_FAST:                          "WALK_FAST",
	TURNING:                            "TURNING",
	TURNING_RUN:                        "TURNING_RUN",
	DASHING:                            "DASHING",
	RUNNING:                            "RUNNING",
	RUN_DIRECT:                         "RUN_DIRECT",
	RUN_BRAKE:                          "RUN_BRAKE",
	KNEE_BEND:                          "KNEE_BEND",
	JUMPING_FORWARD:                    "JUMPING_FORWARD",
	JUMPING_BACKWARD:                   "JUMPING_BACKWARD",
	JUMPING_ARIAL_FORWARD:              "JUMPING_ARIAL_FORWARD",
	JUMPING_ARIAL_BACKWARD:             "JUMPING_ARIAL_BACKWARD",
	FALLING:                            "FALLING",
	FALLING_AERIAL:                     "FALLING_AERIAL",
	DEAD_FALL:                          "DEAD_FALL",
	SPECIAL_FALL_FORWARD:               "SPECIAL_FALL_FORWARD",
	SPECIAL_FALL_BACK:                  "SPECIAL_FALL_BACK",
	TUMBLING:                           "TUMBLING",
	CROUCH_START:                       "CROUCH_START",
	CROUCHING:                          "CROUCHING",
	CROUCH_END:                         "CROUCH_END",
	LANDING:                            "LANDING",
	LANDING_SPECIAL:                    "LANDING_SPECIAL",
	NEUTRAL_ATTACK_1:                   "NEUTRAL_ATTACK_1",
	NEUTRAL_ATTACK_2:                   "NEUTRAL_ATTACK_2",
	NEUTRAL_ATTACK_3:                   "NEUTRAL_ATTACK_3",
	LOOPING_ATTACK_START:               "LOOPING_ATTACK_START",
	LOOPING_ATTACK_MIDDLE:              "LOOPING_ATTACK_MIDDLE",
	LOOPING_ATTACK_END:                 "LOOPING_ATTACK_END",
	DASH_ATTACK:                        "DASH_ATTACK",
	FTILT_HIGH:                         "FTILT_HIGH",
	FTILT_HIGH_MID:                     "FTILT_HIGH_MID",
	FTILT_MID:                          "FTILT_MID",
	FTILT_LOW_MID:                      "FTILT_LOW_MID",
	FTILT_LOW:                          "FTILT_LOW",
	UPTILT:                             "UPTILT",
	DOWNTILT:                           "DOWNTILT",
	FSMASH_HIGH:                        "FSMASH_HIGH",
	FSMASH_MID_HIGH:                    "FSMASH_MID_HIGH",
	FSMASH_MID:                         "FSMASH_MID",
	FSMASH_MID_LOW:                     "FSMASH_MID_LOW",
	FSMASH_LOW:                         "FSMASH_LOW",
	UPSMASH:                            "UPSMASH",
	DOWNSMASH:                          "DOWNSMASH",
	NAIR:                               "NAIR",
	FAIR:                               "FAIR",
	BAIR:                               "BAIR",
	UAIR:                               "UAIR",
	DAIR:                               "DAIR",
	NAIR_LANDING:                       "NAIR_LANDING",
	FAIR_LANDING:                       "FAIR_LANDING",
	BAIR_LANDING:                       "BAIR_LANDING",
	UAIR_LANDING:                       "UAIR_LANDING",
	DAIR_LANDING:                       "DAIR_LANDING",
	DAMAGE_HIGH_1:                      "DAMAGE_HIGH_1",
	DAMAGE_HIGH_2:                      "DAMAGE_HIGH_2",
	DAMAGE_HIGH_3:                      "DAMAGE_HIGH_3",
	DAMAGE_NEUTRAL_1:                   "DAMAGE_NEUTRAL_1",
	DAMAGE_NEUTRAL_2:                   "DAMAGE_NEUTRAL_2",
	DAMAGE_NEUTRAL_3:                   "DAMAGE_NEUTRAL_3",
	DAMAGE_LOW_1:                       "DAMAGE_LOW_1",
	DAMAGE_LOW_2:                       "DAMAGE_LOW_2",
	DAMAGE_LOW_3:                       "DAMAGE_LOW_3",
	DAMAGE_AIR_1:                       "DAMAGE_AIR_1",
	DAMAGE_AIR_2:                       "DAMAGE_AIR_2",
	DAMAGE_AIR_3:                       "DAMAGE_AIR_3",
	DAMAGE_FLY_HIGH:                    "DAMAGE_FLY_HIGH",
	DAMAGE_FLY_NEUTRAL:                 "DAMAGE_FLY_NEUTRAL",
	DAMAGE_FLY_LOW:                     "DAMAGE_FLY_LOW",
	DAMAGE_FLY_TOP:                     "DAMAGE_FLY_TOP",
	DAMAGE_FLY_ROLL:                    "DAMAGE_FLY_ROLL",
	ITEM_PICKUP_LIGHT:                  "ITEM_PICKUP_LIGHT",
	ITEM_PICKUP_HEAVY:                  "ITEM_PICKUP_HEAVY",
	ITEM_THROW_LIGHT_FORWARD:           "ITEM_THROW_LIGHT_FORWARD",
	ITEM_THROW_LIGHT_BACK:              "ITEM_THROW_LIGHT_BACK",
	ITEM_THROW_LIGHT_HIGH:              "ITEM_THROW_LIGHT_HIGH",
	ITEM_THROW_LIGHT_LOW:               "ITEM_THROW_LIGHT_LOW",
	ITEM_THROW_LIGHT_DASH:              "ITEM_THROW_LIGHT_DASH",
	ITEM_THROW_LIGHT_DROP:              "ITEM_THROW_LIGHT_DROP",
	ITEM_THROW_LIGHT_AIR_FORWARD:       "ITEM_THROW_LIGHT_AIR_FORWARD",
	ITEM_THROW_LIGHT_AIR_BACK:          "ITEM_THROW_LIGHT_AIR_BACK",
	ITEM_THROW_LIGHT_AIR_HIGH:          "ITEM_THROW_LIGHT_AIR_HIGH",
	ITEM_THROW_LIGHT_AIR_LOW:           "ITEM_THROW_LIGHT_AIR_LOW",
	ITEM_THROW_HEAVY_FORWARD:           "ITEM_THROW_HEAVY_FORWARD",
	ITEM_THROW_HEAVY_BACK:              "ITEM_THROW_HEAVY_BACK",
	ITEM_THROW_HEAVY_HIGH:              "ITEM_THROW_HEAVY_HIGH",
	ITEM_THROW_HEAVY_LOW:               "ITEM_THROW_HEAVY_LOW",
	ITEM_THROW_LIGHT_SMASH_FORWARD:     "ITEM_THROW_LIGHT_SMASH_FORWARD",
	ITEM_THROW_LIGHT_SMASH_BACK:        "ITEM_THROW_LIGHT_SMASH_BACK",
	ITEM_THROW_LIGHT_SMASH_UP:          "ITEM_THROW_LIGHT_SMASH_UP",
	ITEM_THROW_LIGHT_SMASH_DOWN:        "ITEM_THROW_LIGHT_SMASH_DOWN",
	ITEM_THROW_LIGHT_AIR_SMASH_FORWARD: "ITEM_THROW_LIGHT_AIR_SMASH_FORWARD",
	ITEM_THROW_LIGHT_AIR_SMASH_BACK:    "ITEM_THROW_LIGHT_AIR_SMASH_BACK",
	ITEM_THROW_LIGHT_AIR_SMASH_HIGH:    "ITEM_THROW_LIGHT_AIR_SMASH_HIGH",
	ITEM_THROW_LIGHT_AIR_SMASH_LOW:     "ITEM_THROW_LIGHT_AIR_SMASH_LOW",
	ITEM_THROW_HEAVY_AIR_SMASH_FORWARD: "ITEM_THROW_HEAVY_AIR_SMASH_FORWARD",
	ITEM_THROW_HEAVY_AIR_SMASH_BACK:    "ITEM_THROW_HEAVY_AIR_SMASH_BACK",
	ITEM_THROW_HEAVY_AIR_SMASH_HIGH:    "ITEM_THROW_HEAVY_AIR_SMASH_HIGH",
	ITEM_THROW_HEAVY_AIR_SMASH_LOW:     "ITEM_THROW_HEAVY_AIR_SMASH_LOW",
	BEAM_SWORD_SWING_1:                 "BEAM_SWORD_SWING_1",
	BEAM_SWORD_SWING_2:                 "BEAM_SWORD_SWING_2",
	BEAM_SWORD_SWING_3:                 "BEAM_SWORD_SWING_3",
	BEAM_SWORD_SWING_4:                 "BEAM_SWORD_SWING_4",
	SHIELD_START:                       "SHIELD_START",
	SHIELD:                             "SHIELD",
	SHIELD_RELEASE:                     "SHIELD_RELEASE",
	SHIELD_STUN:                        "SHIELD_STUN",
	SHIELD_REFLECT:                     "SHIELD_REFLECT",
	TECH_MISS_UP:                       "TECH_MISS_UP",
	LYING_GROUND_UP:                    "LYING_GROUND_UP",
	LYING_GROUND_UP_HIT:                "LYING_GROUND_UP_HIT",
	GROUND_GETUP:                       "GROUND_GETUP",
	GROUND_ATTACK_UP:                   "GROUND_ATTACK_UP",
	GROUND_ROLL_FORWARD_UP:             "GROUND_ROLL_FORWARD_UP",
	GROUND_ROLL_BACKWARD_UP:            "GROUND_ROLL_BACKWARD_UP",
	TECH_MISS_DOWN:                     "TECH_MISS_DOWN",
	LYING_GROUND_DOWN:                  "LYING_GROUND_DOWN",
	DAMAGE_GROUND:                      "DAMAGE_GROUND",
	NEUTRAL_GETUP:                      "NEUTRAL_GETUP",
	GETUP_ATTACK:                       "GETUP_ATTACK",
	GROUND_ROLL_FORWARD_DOWN:           "GROUND_ROLL_FORWARD_DOWN",
	GROUND_ROLL_BACKWARD_DOWN:          "GROUND_ROLL_BACKWARD_DOWN",
	NEUTRAL_TECH:                       "NEUTRAL_TECH",
	FORWARD_TECH:                       "FORWARD_TECH",
	BACKWARD_TECH:                      "BACKWARD_TECH",
	WALL_TECH:                          "WALL_TECH",
	WALL_TECH_JUMP:                     "WALL_TECH_JUMP",
	CEILING_TECH:                       "CEILING_TECH",
	SHIELD_BREAK_FLY:                   "SHIELD_BREAK_FLY",
	SHIELD_BREAK_FALL:                  "SHIELD_BREAK_FALL",
	SHIELD_BREAK_DOWN_U:                "SHIELD_BREAK_DOWN_U",
	SHIELD_BREAK_DOWN_D:                "SHIELD_BREAK_DOWN_D",
	SHIELD_BREAK_STAND_U:               "SHIELD_BREAK_STAND_U",
	SHIELD_BREAK_STAND_D:               "SHIELD_BREAK_STAND_D",
	SHIELD_BREAK_TEETER:                "SHIELD_BREAK_TEETER",
	GRAB:                               "GRAB",
	GRAB_PULLING:                       "GRAB_PULLING",
	GRAB_RUNNING:                       "GRAB_RUNNING",
	GRAB_WAIT:                          "GRAB_WAIT",
	GRAB_PUMMEL:                        "GRAB_PUMMEL",
	GRAB_BREAK:                         "GRAB_BREAK",
	THROW_FORWARD:                      "THROW_FORWARD",
	THROW_BACK:                         "THROW_BACK",
	THROW_UP:                           "THROW_UP",
	THROW_DOWN:                         "THROW_DOWN",
	GRAB_PULLING_HIGH:                  "GRAB_PULLING_HIGH",
	GRABBED_WAIT_HIGH:                  "GRABBED_WAIT_HIGH",
	PUMMELED_HIGH:                      "PUMMELED_HIGH",
	GRAB_PULL:                          "GRAB_PULL",
	GRABBED:                            "GRABBED",
	GRAB_PUMMELED:                      "GRAB_PUMMELED",
	GRAB_ESCAPE:                        "GRAB_ESCAPE",
	ROLL_FORWARD:                       "ROLL_FORWARD",
	ROLL_BACKWARD:                      "ROLL_BACKWARD",
	SPOTDODGE:                          "SPOTDODGE",
	AIRDODGE:                           "AIRDODGE",
	THROWN_FORWARD:                     "THROWN_FORWARD",
	THROWN_BACK:                        "THROWN_BACK",
	THROWN_UP:                          "THROWN_UP",
	THROWN_DOWN:                        "THROWN_DOWN",
	THROWN_DOWN_2:                      "THROWN_DOWN_2",
	PLATFORM_DROP:                      "PLATFORM_DROP",
	EDGE_TEETERING_START:               "EDGE_TEETERING_START",
	EDGE_TEETERING:                     "EDGE_TEETERING",
	BOUNCE_WALL:                        "BOUNCE_WALL",
	BOUNCE_CEILING:                     "BOUNCE_CEILING",
	BUMP_WALL:                          "BUMP_WALL",
	BUMP_CIELING:                       "BUMP_CIELING",
	SLIDING_OFF_EDGE:                   "SLIDING_OFF_EDGE",
	EDGE_CATCHING:                      "EDGE_CATCHING",
	EDGE_HANGING:                       "EDGE_HANGING",
	EDGE_GETUP_SLOW:                    "EDGE_GETUP_SLOW",
	EDGE_GETUP_QUICK:                   "EDGE_GETUP_QUICK",
	EDGE_ATTACK_SLOW:                   "EDGE_ATTACK_SLOW",
	EDGE_ATTACK_QUICK:                  "EDGE_ATTACK_QUICK",
	EDGE_ROLL_SLOW:                     "EDGE_ROLL_SLOW",
	EDGE_ROLL_QUICK:                    "EDGE_ROLL_QUICK",
	EDGE_JUMP_1_SLOW:                   "EDGE_JUMP_1_SLOW",
	EDGE_JUMP_2_SLOW:                   "EDGE_JUMP_2_SLOW",
	EDGE_JUMP_1_QUICK:                  "EDGE_JUMP_1_QUICK",
	EDGE_JUMP_2_QUICK:                  "EDGE_JUMP_2_QUICK",
	TAUNT_RIGHT:                        "TAUNT_RIGHT",
	TAUNT_LEFT:                         "TAUNT_LEFT",
	ENTRY:                              "ENTRY",
	ENTRY_START:                        "ENTRY_START",
	ENTRY_END:                          "ENTRY_END",
	LASER_GUN_PULL:                     "LASER_GUN_PULL",
	NEUTRAL_B_CHARGING:                 "NEUTRAL_B_CHARGING",
	NEUTRAL_B_ATTACKING:                "NEUTRAL_B_ATTACKING",
	NEUTRAL_B_FULL_CHARGE:              "NEUTRAL_B_FULL_CHARGE",
	WAIT_ITEM:                          "WAIT_ITEM",
	NEUTRAL_B_CHARGING_AIR:             "NEUTRAL_B_CHARGING_AIR",
	NEUTRAL_B_ATTACKING_AIR:            "NEUTRAL_B_ATTACKING_AIR",
	NEUTRAL_B_FULL_CHARGE_AIR:          "NEUTRAL_B_FULL_CHARGE_AIR",
	//SWORD_DANCE_1:                      "SWORD_DANCE_1",
	//SWORD_DANCE_2_HIGH:                 "SWORD_DANCE_2_HIGH",
	//SWORD_DANCE_2_MID:                  "SWORD_DANCE_2_MID",
	//SWORD_DANCE_3_HIGH:                 "SWORD_DANCE_3_HIGH",
	//SWORD_DANCE_3_MID:                  "SWORD_DANCE_3_MID",
	//SWORD_DANCE_3_LOW:                  "SWORD_DANCE_3_LOW",
	//SWORD_DANCE_4_HIGH:                 "SWORD_DANCE_4_HIGH",
	//SWORD_DANCE_4_MID:                  "SWORD_DANCE_4_MID",
	//SWORD_DANCE_4_LOW:                  "SWORD_DANCE_4_LOW",
	//SWORD_DANCE_1_AIR:                  "SWORD_DANCE_1_AIR",
	//SWORD_DANCE_2_HIGH_AIR:             "SWORD_DANCE_2_HIGH_AIR",
	//SWORD_DANCE_2_MID_AIR:              "SWORD_DANCE_2_MID_AIR",
	//SWORD_DANCE_3_HIGH_AIR:             "SWORD_DANCE_3_HIGH_AIR",
	//SWORD_DANCE_3_MID_AIR:              "SWORD_DANCE_3_MID_AIR",
	//SWORD_DANCE_3_LOW_AIR:              "SWORD_DANCE_3_LOW_AIR",
	//SWORD_DANCE_4_HIGH_AIR:             "SWORD_DANCE_4_HIGH_AIR",
	//SWORD_DANCE_4_MID_AIR:              "SWORD_DANCE_4_MID_AIR",
	//SWORD_DANCE_4_LOW_AIR:              "SWORD_DANCE_4_LOW_AIR",
	FOX_ILLUSION_START:     "FOX_ILLUSION_START",
	FOX_ILLUSION:           "FOX_ILLUSION",
	FOX_ILLUSION_SHORTENED: "FOX_ILLUSION_SHORTENED",
	FIREFOX_WAIT_GROUND:    "FIREFOX_WAIT_GROUND",
	FIREFOX_WAIT_AIR:       "FIREFOX_WAIT_AIR",
	FIREFOX_GROUND:         "FIREFOX_GROUND",
	FIREFOX_AIR:            "FIREFOX_AIR",
	DOWN_B_GROUND_START:    "DOWN_B_GROUND_START",
	DOWN_B_GROUND:          "DOWN_B_GROUND",
	SHINE_TURN:             "SHINE_TURN",
	DOWN_B_STUN:            "DOWN_B_STUN",
	DOWN_B_AIR:             "DOWN_B_AIR",
	UP_B_GROUND:            "UP_B_GROUND",
	SHINE_RELEASE_AIR:      "SHINE_RELEASE_AIR OR UP_B_AIR",
	MARTH_COUNTER:          "MARTH_COUNTER",
	PARASOL_FALLING:        "PARASOL_FALLING",
	MARTH_COUNTER_FALLING:  "MARTH_COUNTER_FALLING",
	//NESS_SHEILD_START:      "NESS_SHEILD_START",
	NESS_SHEILD:         "NESS_SHEILD",
	NESS_SHEILD_AIR:     "NESS_SHEILD_AIR",
	NESS_SHEILD_AIR_END: "NESS_SHEILD_AIR_END",
	UNKNOWN_ANIMATION:   "UNKNOWN_ANIMATION",
}

func (a PlayerAction) CanLCancel() bool {
	return (a == UAIR_LANDING || a == BAIR_LANDING || a == NAIR_LANDING || a ==
		DAIR_LANDING || a == FAIR_LANDING)
}

func (a PlayerAction) IsAttack() bool {
	return (a == BAIR || a == FAIR || a == DAIR || a == NAIR || a == UAIR ||
		a == FSMASH_HIGH || a == FSMASH_MID || a == FSMASH_MID_LOW ||
		a == FSMASH_MID_HIGH || a == FSMASH_LOW || a == UPSMASH ||
		a == DOWNSMASH || a == UPTILT || a == FTILT_HIGH || a == FALLING_AERIAL ||
		a == FTILT_HIGH_MID || a == FTILT_MID || a == FTILT_LOW_MID ||
		a == FTILT_LOW || a == UPTILT || a == NEUTRAL_ATTACK_1 ||
		a == NEUTRAL_ATTACK_2 || a == NEUTRAL_ATTACK_3 ||
		a == LOOPING_ATTACK_START || a == LOOPING_ATTACK_MIDDLE ||
		a == LOOPING_ATTACK_END || a == DASH_ATTACK || a == EDGE_ATTACK_SLOW ||
		a == EDGE_ATTACK_QUICK || a == NEUTRAL_B_ATTACKING ||
		a == NEUTRAL_B_ATTACKING_AIR || a == GRAB || a == GRAB_RUNNING ||
		a == DASH_ATTACK || a == THROW_FORWARD || a == THROW_BACK || a == THROW_UP ||
		a == THROW_DOWN)
}

func (a PlayerAction) IsAttackStrict() bool {
	return (a.IsAttack() || a == GROUND_ATTACK_UP || a == DASH_ATTACK ||
		a == NAIR_LANDING || a == FAIR_LANDING || a == BAIR_LANDING ||
		a == UAIR_LANDING || a == DAIR_LANDING || a == ITEM_THROW_LIGHT_FORWARD ||
		a == ITEM_THROW_LIGHT_BACK || a == ITEM_THROW_LIGHT_HIGH ||
		a == ITEM_THROW_LIGHT_LOW || a == ITEM_THROW_LIGHT_DASH ||
		a == ITEM_THROW_LIGHT_DROP || a == ITEM_THROW_LIGHT_AIR_FORWARD ||
		a == ITEM_THROW_LIGHT_AIR_BACK || a == ITEM_THROW_LIGHT_AIR_HIGH ||
		a == ITEM_THROW_LIGHT_AIR_LOW || a == ITEM_THROW_HEAVY_FORWARD ||
		a == ITEM_THROW_HEAVY_BACK || a == ITEM_THROW_HEAVY_HIGH ||
		a == ITEM_THROW_HEAVY_LOW || a == ITEM_THROW_LIGHT_SMASH_FORWARD ||
		a == ITEM_THROW_LIGHT_SMASH_BACK || a == ITEM_THROW_LIGHT_SMASH_UP ||
		a == ITEM_THROW_LIGHT_SMASH_DOWN || a == ITEM_THROW_LIGHT_AIR_SMASH_FORWARD ||
		a == ITEM_THROW_LIGHT_AIR_SMASH_BACK || a == ITEM_THROW_LIGHT_AIR_SMASH_HIGH ||
		a == ITEM_THROW_LIGHT_AIR_SMASH_LOW || a == ITEM_THROW_HEAVY_AIR_SMASH_FORWARD ||
		a == ITEM_THROW_HEAVY_AIR_SMASH_BACK || a == ITEM_THROW_HEAVY_AIR_SMASH_HIGH ||
		a == ITEM_THROW_HEAVY_AIR_SMASH_LOW)
}

func (a PlayerAction) IsNeutral() bool {
	return (a == STANDING || a == WALK_SLOW || a == WALK_MIDDLE || a == WALK_FAST ||
		a == TURNING || a == TURNING_RUN || a == DASHING || a == RUNNING ||
		a == RUN_DIRECT || a == RUN_BRAKE || a == FALLING || a == CROUCH_START ||
		a == CROUCHING || a == CROUCH_END || a == LANDING || a == LANDING_SPECIAL ||
		a == ITEM_PICKUP_LIGHT || a == ITEM_PICKUP_HEAVY || a == ROLL_FORWARD ||
		a == ROLL_BACKWARD || a == SPOTDODGE || a == AIRDODGE || a == PLATFORM_DROP ||
		a == EDGE_TEETERING_START || a == EDGE_TEETERING || a == TAUNT_RIGHT ||
		a == TAUNT_LEFT || a == NESS_SHEILD || a == NESS_SHEILD_AIR ||
		a == NESS_SHEILD_AIR_END || a == ON_HALO_DESCENT || a == ON_HALO_WAIT ||
		a.IsShield() || a.IsJump())
}

// Includes lasers, neutral B charging etc.
func (a PlayerAction) IsNeutralStrict() bool {
	return (a.IsNeutral() || a == LASER_GUN_PULL || a == NEUTRAL_B_CHARGING ||
		a == NEUTRAL_B_ATTACKING || a == NEUTRAL_B_FULL_CHARGE || a == NEUTRAL_B_CHARGING_AIR ||
		a == NEUTRAL_B_ATTACKING_AIR || a == NEUTRAL_B_FULL_CHARGE_AIR)
}

func (a PlayerAction) IsEdge() bool {
	return (a == EDGE_CATCHING || a == EDGE_HANGING || a == EDGE_GETUP_SLOW ||
		a == EDGE_GETUP_QUICK || a == EDGE_ATTACK_SLOW || a == EDGE_ATTACK_QUICK ||
		a == EDGE_ROLL_SLOW || a == EDGE_ROLL_QUICK || a == EDGE_JUMP_1_SLOW ||
		a == EDGE_JUMP_2_SLOW || a == EDGE_JUMP_1_QUICK || a == EDGE_JUMP_2_QUICK)
}

func (a PlayerAction) IsTech() bool {
	return (a == NEUTRAL_TECH || a == FORWARD_TECH || a == BACKWARD_TECH ||
		a == WALL_TECH || a == WALL_TECH_JUMP || a == CEILING_TECH)
}

// Includes any action you could perform from a techable impact.
func (a PlayerAction) IsTechStrict() bool {
	return (a.IsTech() || a == TECH_MISS_UP || a == TECH_MISS_DOWN ||
		a == LYING_GROUND_UP || a == LYING_GROUND_UP_HIT || a == GROUND_GETUP ||
		a == GROUND_ATTACK_UP || a == GROUND_ROLL_FORWARD_UP || a == GROUND_ROLL_BACKWARD_UP ||
		a == LYING_GROUND_DOWN || a == DAMAGE_GROUND || a == NEUTRAL_GETUP ||
		a == GETUP_ATTACK || a == GROUND_ROLL_FORWARD_DOWN || a == GROUND_ROLL_BACKWARD_DOWN)
}

func (a PlayerAction) IsDamage() bool {
	return (a == DAMAGE_HIGH_1 || a == DAMAGE_HIGH_2 || a == DAMAGE_HIGH_3 ||
		a == DAMAGE_NEUTRAL_1 || a == DAMAGE_NEUTRAL_2 || a == DAMAGE_NEUTRAL_3 ||
		a == DAMAGE_LOW_1 || a == DAMAGE_LOW_2 || a == DAMAGE_LOW_3 ||
		a == DAMAGE_AIR_1 || a == DAMAGE_AIR_2 || a == DAMAGE_AIR_3 ||
		a == DAMAGE_FLY_HIGH || a == DAMAGE_FLY_NEUTRAL || a == DAMAGE_FLY_LOW ||
		a == DAMAGE_FLY_TOP || a == DAMAGE_FLY_ROLL)
}

func (a PlayerAction) IsStun() bool {
	return (a == EDGE_CATCHING || a == LANDING_SPECIAL || a == DOWN_B_STUN ||
		a == SHIELD_STUN)
}

func (a PlayerAction) IsShield() bool {
	return (a == SHIELD_START || a == SHIELD || a == SHIELD_RELEASE ||
		a == SHIELD_STUN || a == SHIELD_REFLECT)
}

func (a PlayerAction) IsIdle() bool {
	return (a == STANDING || a == CROUCH_START || a == CROUCHING ||
		a == CROUCH_END || a == LANDING ||
		a.IsThrown() || a.IsShieldStrict() || a.IsDead())
}

func (a PlayerAction) IsMoving() bool {
	return !a.IsIdle()
}

func (a PlayerAction) IsJump() bool {
	return (a == KNEE_BEND || a == JUMPING_FORWARD || a == JUMPING_BACKWARD ||
		a == JUMPING_ARIAL_FORWARD || a == JUMPING_ARIAL_BACKWARD)
}

// includes shield breaking(letting go of shield?)
func (a PlayerAction) IsShieldStrict() bool {
	return (a.IsShield() || a == SHIELD_BREAK_FLY || a == SHIELD_BREAK_FALL ||
		a == SHIELD_BREAK_DOWN_U || a == SHIELD_BREAK_DOWN_D ||
		a == SHIELD_BREAK_STAND_U || a == SHIELD_BREAK_STAND_D || a == SHIELD_BREAK_TEETER)
}

func (a PlayerAction) IsDead() bool {
	return (a == DEAD_DOWN || a == DEAD_FALL || a == DEAD_LEFT ||
		a == DEAD_RIGHT || a == DEAD_FLY_STAR || a == DEAD_FLY ||
		a == DEAD_FLY_SPLATTER || a == DEAD_FLY_SPLATTER_FLAT)
}

func (a PlayerAction) IsGrabbed() bool {
	return (a == GRABBED || a == GRABBED_WAIT_HIGH || a == PUMMELED_HIGH ||
		a == GRAB_PUMMELED || a == GRAB_ESCAPE || a == GRAB_PULL)
}

func (a PlayerAction) IsThrown() bool {
	return (a == THROWN_FORWARD || a == THROWN_BACK || a == THROWN_UP ||
		a == THROWN_DOWN || a == THROWN_DOWN_2)
}
