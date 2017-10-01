package data

type CharacterFrameData struct {
	Character Character
	Action    Action
	Frame     int
	Hitbox_1_Status
}

type ActionFrame struct {
	Action          Action
	Frame           int
	Hitbox_1_Status bool
	Hitbox_1_Size   float
	Hitbox_1_X      float
	Hitbox_1_Y      float
	Hitbox_2_Status bool
	Hitbox_2_Size   float
	Hitbox_2_X      float
	Hitbox_2_Y      float
	Hitbox_3_Status bool
	Hitbox_3_Size   float
	Hitbox_3_X      float
	Hitbox_3_Y      float
	Hitbox_4_Status bool
	Hitbox_4_Size   float
	Hitbox_4_X      float
	Hitbox_4_Y      float
	LocomotionX     float
	LocomotionY     float
	IASA            bool
	FacingChanged   bool
	Projectile      bool
}
