package melee

type ActionData struct {
	Character   Character
	Action      Action
	ZeroIndexed bool
}

func GetActionData() []ActionData {
	characterData := []ActionData{
		{9, 322, false},
		{9, 323, true},
		{9, 324, false},
		{9, 29, true},
		{9, 42, true},
		{9, 14, true},
		{9, 18, false},
		{9, 39, false},
		{9, 15, false},
		{9, 16, false},
		{9, 44, false},
		{9, 45, true},
		{9, 47, false},
		{9, 48, true},
		{9, 49, true},
		{9, 17, false},
		{9, 56, false},
		{9, 40, true},
		{9, 57, false},
		{9, 41, true},
		{9, 53, false},
		{9, 51, false},
		{9, 55, false},
		{9, 20, false},
		{9, 21, true},
		{9, 23, true},
		{9, 63, false},
		{9, 64, false},
		{9, 60, false},
		{9, 360, false},
		{9, 361, true},
		{9, 363, true},
		{9, 24, true},
		{9, 25, true},
		{9, 347, false},
		{9, 348, true},
		{9, 349, true},
		{9, 341, false},
		{9, 342, true},
		{9, 343, true},
		{9, 353, false},
		{9, 356, true},
		{9, 358, true},
		{9, 35, true},
		{9, 43, true},
		{9, 27, true},
		{9, 354, false},
		{9, 365, false},
		{9, 366, true},
		{9, 368, true},
		{9, 32, true},
		{9, 28, true},
		{9, 350, false},
		{9, 351, true},
		{9, 352, true},
		{9, 344, false},
		{9, 345, true},
		{9, 346, true},
		{9, 252, false},
		{9, 253, true},
		{9, 257, false},
		{9, 0, false},
		{9, 12, true},
		{9, 13, true},
		{9, 67, false},
		{9, 255, false},
		{9, 68, false},
		{9, 69, false},
		{9, 66, false},
		{9, 65, false},
		{9, 70, true},
		{9, 73, true},
		{9, 72, true},
		{9, 74, true},
		{9, 71, true},
		{9, 50, false},
		{9, 212, true},
		{9, 214, true},
		{9, 245, true},
		{9, 246, true},
		{9, 259, false},
		{9, 26, true},
		{9, 262, false},
		{9, 263, false},
		{9, 235, false},
		{9, 178, false},
		{9, 182, false},
		{9, 233, false},
		{9, 234, false},
		{9, 80, false},
		{9, 90, false},
		{9, 183, true},
		{9, 184, false},
		{9, 189, false},
		{9, 197, false},
		{9, 38, true},
		{9, 191, true},
		{9, 192, false},
		{9, 194, true},
		{9, 196, false},
		{9, 187, false},
		{9, 188, false},
		{9, 186, true},
		{9, 254, false},
		{9, 258, false},
		{9, 260, false},
		{9, 261, false},
		{9, 256, false},
		{9, 4, true},
		{9, 79, false},
		{9, 199, true},
		{9, 201, true},
		{9, 200, true},
		{9, 77, false},
		{9, 264, true},
		{9, 179, false},
		{9, 205, false},
		{9, 207, true},
		{9, 209, true},
		{9, 19, true},
		{7, 322, false},
		{7, 323, true},
		{7, 324, false},
		{7, 29, true},
		{7, 42, true},
		{7, 14, true},
		{7, 39, false},
		{7, 44, false},
		{7, 45, true},
		{7, 46, true},
		{7, 47, false},
		{7, 48, true},
		{7, 49, true},
		{7, 56, false},
		{7, 40, true},
		{7, 57, false},
		{7, 41, true},
		{7, 15, false},
		{7, 51, false},
		{7, 16, false},
		{7, 17, false},
		{7, 52, false},
		{7, 18, false},
		{7, 53, false},
		{7, 55, false},
		{7, 20, false},
		{7, 21, true},
		{7, 23, true},
		{7, 245, true},
		{7, 246, true},
		{7, 19, false},
		{7, 24, true},
		{7, 25, true},
		{7, 236, false},
		{7, 43, true},
		{7, 26, true},
		{7, 63, false},
		{7, 60, false},
		{7, 64, false},
		{7, 27, true},
		{7, 353, false},
		{7, 35, true},
		{7, 349, false},
		{7, 347, false},
		{7, 357, false},
		{7, 358, true},
		{7, 32, true},
		{7, 354, false},
		{7, 359, false},
		{7, 361, true},
		{7, 28, true},
		{7, 348, false},
		{7, 252, false},
		{7, 253, true},
		{7, 351, false},
		{7, 68, false},
		{7, 69, false},
		{7, 67, false},
		{7, 262, false},
		{7, 263, false},
		{7, 66, false},
		{7, 0, false},
		{7, 12, true},
		{7, 65, false},
		{7, 13, true},
		{7, 74, true},
		{7, 71, true},
		{7, 72, true},
		{7, 73, true},
		{7, 70, true},
		{7, 50, false},
		{7, 212, true},
		{7, 214, true},
		{7, 255, false},
		{7, 259, false},
		{7, 257, false},
		{7, 178, false},
		{7, 182, false},
		{7, 235, false},
		{7, 233, false},
		{7, 234, false},
		{7, 79, false},
		{7, 80, false},
		{7, 90, false},
		{7, 183, true},
		{7, 184, false},
		{7, 38, true},
		{7, 186, true},
		{7, 191, true},
		{7, 192, false},
		{7, 197, false},
		{7, 195, false},
		{7, 196, false},
		{7, 187, false},
		{7, 188, false},
		{7, 194, true},
		{7, 4, true},
		{7, 77, false},
		{7, 199, true},
		{7, 201, true},
		{7, 200, true},
		{7, 254, false},
		{7, 258, false},
		{7, 260, false},
		{7, 261, false},
		{7, 256, false},
		{7, 264, true},
		{7, 265, true},
		{7, 179, false},
		{7, 205, false},
		{7, 208, true},
		{7, 210, true},
		{10, 322, false},
		{10, 323, true},
		{10, 324, false},
		{10, 29, true},
		{10, 42, false},
		{10, 14, true},
		{10, 39, false},
		{10, 20, false},
		{10, 18, false},
		{10, 21, true},
		{10, 19, true},
		{10, 23, true},
		{10, 28, true},
		{10, 354, false},
		{10, 356, true},
		{10, 358, true},
		{10, 35, true},
		{10, 43, true},
		{10, 15, false},
		{10, 16, false},
		{10, 17, false},
		{10, 245, true},
		{10, 246, true},
		{10, 44, false},
		{10, 45, true},
		{10, 47, false},
		{10, 48, true},
		{10, 49, true},
		{10, 56, false},
		{10, 40, true},
		{10, 57, false},
		{10, 41, true},
		{10, 53, false},
		{10, 51, false},
		{10, 55, false},
		{10, 63, false},
		{10, 64, false},
		{10, 60, false},
		{10, 341, false},
		{10, 342, true},
		{10, 343, true},
		{10, 353, false},
		{10, 347, false},
		{10, 348, true},
		{10, 349, true},
		{10, 360, false},
		{10, 361, true},
		{10, 363, true},
		{10, 24, true},
		{10, 25, true},
		{10, 27, true},
		{10, 365, false},
		{10, 366, true},
		{10, 368, true},
		{10, 252, false},
		{10, 253, true},
		{10, 350, false},
		{10, 351, true},
		{10, 352, true},
		{10, 344, false},
		{10, 345, true},
		{10, 346, true},
		{10, 68, false},
		{10, 255, false},
		{10, 69, false},
		{10, 67, false},
		{10, 262, false},
		{10, 263, false},
		{10, 66, false},
		{10, 65, false},
		{10, 73, true},
		{10, 70, true},
		{10, 72, true},
		{10, 74, true},
		{10, 71, true},
		{10, 50, false},
		{10, 212, true},
		{10, 214, true},
		{10, 257, false},
		{10, 259, false},
		{10, 26, true},
		{10, 236, false},
		{10, 178, false},
		{10, 182, false},
		{10, 235, false},
		{10, 233, false},
		{10, 234, false},
		{10, 80, false},
		{10, 90, false},
		{10, 183, true},
		{10, 184, false},
		{10, 189, false},
		{10, 38, true},
		{10, 188, false},
		{10, 191, true},
		{10, 192, false},
		{10, 196, false},
		{10, 197, false},
		{10, 187, false},
		{10, 186, true},
		{10, 195, false},
		{10, 199, true},
		{10, 200, true},
		{10, 201, true},
		{10, 260, false},
		{10, 261, false},
		{10, 254, false},
		{10, 258, false},
		{10, 256, false},
		{10, 179, false},
		{10, 205, false},
		{10, 207, true},
		{10, 209, true},
		{10, 264, true},
		{10, 180, true},
		{10, 4, true},
		{10, 12, true},
		{10, 13, true},
		{20, 322, false},
		{20, 323, true},
		{20, 324, false},
		{20, 29, true},
		{20, 42, true},
		{20, 14, true},
		{20, 39, false},
		{20, 44, false},
		{20, 45, true},
		{20, 56, false},
		{20, 40, true},
		{20, 57, false},
		{20, 41, true},
		{20, 15, false},
		{20, 53, false},
		{20, 51, false},
		{20, 16, false},
		{20, 17, false},
		{20, 55, false},
		{20, 18, false},
		{20, 24, true},
		{20, 26, true},
		{20, 63, false},
		{20, 64, false},
		{20, 60, false},
		{20, 367, false},
		{20, 20, false},
		{20, 363, false},
		{20, 346, false},
		{20, 348, true},
		{20, 349, false},
		{20, 350, false},
		{20, 351, false},
		{20, 353, false},
		{20, 21, true},
		{20, 23, true},
		{20, 371, false},
		{20, 25, true},
		{20, 341, true},
		{20, 342, true},
		{20, 343, true},
		{20, 368, false},
		{20, 236, false},
		{20, 43, true},
		{20, 252, false},
		{20, 253, true},
		{20, 255, false},
		{20, 364, false},
		{20, 0, false},
		{20, 12, true},
		{20, 13, true},
		{20, 354, false},
		{20, 356, true},
		{20, 357, false},
		{20, 358, false},
		{20, 2, false},
		{20, 68, false},
		{20, 69, false},
		{20, 67, false},
		{20, 65, false},
		{20, 344, true},
		{20, 66, false},
		{20, 19, true},
		{20, 245, true},
		{20, 246, true},
		{20, 178, false},
		{20, 182, false},
		{20, 179, false},
		{20, 180, true},
		{20, 235, false},
		{20, 233, false},
		{20, 234, false},
		{20, 212, true},
		{20, 214, true},
		{20, 262, false},
		{20, 263, false},
		{20, 257, false},
		{20, 259, false},
		{20, 77, false},
		{20, 90, false},
		{20, 38, true},
		{20, 191, true},
		{20, 192, false},
		{20, 196, false},
		{20, 183, true},
		{20, 184, false},
		{20, 188, false},
		{20, 197, false},
		{20, 194, true},
		{20, 195, false},
		{20, 260, false},
		{20, 261, false},
		{20, 254, false},
		{20, 256, false},
		{20, 258, false},
		{20, 6, true},
		{20, 7, true},
		{20, 76, false},
		{20, 189, false},
		{20, 187, false},
		{20, 4, true},
		{20, 86, false},
		{20, 199, true},
		{20, 201, true},
		{20, 200, true},
		{20, 264, true},
		{20, 265, true},
		{20, 205, false},
		{23, 322, false},
		{23, 323, true},
		{23, 324, false},
		{23, 29, true},
		{23, 42, true},
		{23, 14, true},
		{23, 39, false},
		{23, 18, false},
		{23, 20, false},
		{23, 21, true},
		{23, 19, true},
		{23, 15, false},
		{23, 16, false},
		{23, 23, true},
		{23, 24, true},
		{23, 26, true},
		{23, 236, false},
		{23, 43, true},
		{23, 41, true},
		{23, 63, false},
		{23, 60, false},
		{23, 64, false},
		{23, 56, false},
		{23, 53, false},
		{23, 40, true},
		{23, 57, false},
		{23, 44, false},
		{23, 45, true},
		{23, 25, true},
		{23, 27, true},
		{23, 367, false},
		{23, 35, true},
		{23, 369, false},
		{23, 349, false},
		{23, 351, true},
		{23, 353, true},
		{23, 356, true},
		{23, 350, true},
		{23, 352, true},
		{23, 355, true},
		{23, 354, true},
		{23, 357, true},
		{23, 371, false},
		{23, 368, false},
		{23, 252, false},
		{23, 253, true},
		{23, 255, false},
		{23, 341, false},
		{23, 342, true},
		{23, 344, false},
		{23, 343, false},
		{23, 28, true},
		{23, 345, false},
		{23, 347, false},
		{23, 346, true},
		{23, 0, false},
		{23, 12, true},
		{23, 13, true},
		{23, 178, false},
		{23, 180, true},
		{23, 358, false},
		{23, 360, true},
		{23, 362, true},
		{23, 365, true},
		{23, 179, false},
		{23, 363, true},
		{23, 366, true},
		{23, 359, true},
		{23, 361, true},
		{23, 364, true},
		{23, 65, false},
		{23, 245, true},
		{23, 67, false},
		{23, 68, false},
		{23, 69, false},
		{23, 66, false},
		{23, 70, true},
		{23, 73, true},
		{23, 74, true},
		{23, 71, true},
		{23, 212, true},
		{23, 214, true},
		{23, 50, false},
		{23, 262, false},
		{23, 263, false},
		{23, 257, false},
		{23, 259, false},
		{23, 246, true},
		{23, 32, true},
		{23, 182, false},
		{23, 235, false},
		{23, 233, false},
		{23, 234, false},
		{23, 17, false},
		{23, 80, false},
		{23, 90, false},
		{23, 183, true},
		{23, 184, false},
		{23, 38, true},
		{23, 186, true},
		{23, 191, true},
		{23, 192, false},
		{23, 194, true},
		{23, 189, false},
		{23, 196, false},
		{23, 188, false},
		{23, 4, true},
		{23, 79, false},
		{23, 77, false},
		{23, 197, false},
		{23, 187, false},
		{23, 199, true},
		{23, 195, false},
		{23, 201, true},
		{23, 200, true},
		{23, 205, false},
		{23, 206, true},
		{23, 207, true},
		{23, 209, true},
		{23, 254, false},
		{23, 260, false},
		{23, 261, false},
		{23, 258, false},
		{23, 264, true},
		{19, 322, false},
		{19, 323, true},
		{19, 324, false},
		{19, 29, true},
		{19, 42, true},
		{19, 14, true},
		{19, 39, false},
		{19, 20, false},
		{19, 18, false},
		{19, 21, true},
		{19, 23, true},
		{19, 19, true},
		{19, 15, false},
		{19, 44, false},
		{19, 56, false},
		{19, 40, true},
		{19, 57, false},
		{19, 41, true},
		{19, 51, false},
		{19, 16, false},
		{19, 53, false},
		{19, 55, false},
		{19, 63, false},
		{19, 60, false},
		{19, 64, false},
		{19, 24, true},
		{19, 25, true},
		{19, 27, true},
		{19, 356, false},
		{19, 357, false},
		{19, 358, true},
		{19, 35, true},
		{19, 43, true},
		{19, 353, false},
		{19, 0, false},
		{19, 12, true},
		{19, 13, true},
		{19, 178, false},
		{19, 182, false},
		{19, 179, false},
		{19, 181, true},
		{19, 235, false},
		{19, 17, false},
		{19, 343, false},
		{19, 344, true},
		{19, 347, true},
		{19, 350, false},
		{19, 351, true},
		{19, 346, true},
		{19, 245, true},
		{19, 246, true},
		{19, 355, true},
		{19, 359, false},
		{19, 360, true},
		{19, 362, true},
		{19, 361, true},
		{19, 341, false},
		{19, 342, false},
		{19, 363, false},
		{19, 364, true},
		{19, 365, true},
		{19, 366, true},
		{19, 252, false},
		{19, 253, true},
		{19, 255, false},
		{19, 236, false},
		{19, 26, true},
		{19, 28, true},
		{19, 348, false},
		{19, 352, true},
		{19, 349, true},
		{19, 68, false},
		{19, 67, false},
		{19, 66, false},
		{19, 65, false},
		{19, 70, true},
		{19, 69, false},
		{19, 74, true},
		{19, 72, true},
		{19, 73, true},
		{19, 71, true},
		{19, 50, false},
		{19, 212, true},
		{19, 214, true},
		{19, 262, false},
		{19, 263, false},
		{19, 257, false},
		{19, 259, false},
		{19, 233, false},
		{19, 234, false},
		{19, 80, false},
		{19, 90, false},
		{19, 183, true},
		{19, 184, false},
		{19, 38, true},
		{19, 191, true},
		{19, 192, false},
		{19, 196, false},
		{19, 188, false},
		{19, 189, false},
		{19, 4, true},
		{19, 76, false},
		{19, 77, false},
		{19, 199, true},
		{19, 197, false},
		{19, 201, true},
		{19, 79, false},
		{19, 200, true},
		{19, 195, false},
		{19, 187, false},
		{19, 254, false},
		{19, 258, false},
		{19, 256, false},
		{19, 260, false},
		{19, 261, false},
		{19, 205, false},
		{19, 207, true},
		{19, 209, true},
		{19, 265, true},
		{19, 264, true},
		{19, 354, false},
		{4, 322, false},
		{4, 323, false},
		{4, 324, false},
		{4, 29, false},
		{4, 42, false},
		{4, 14, false},
		{4, 44, false},
		{4, 45, true},
		{4, 56, false},
		{4, 15, false},
		{4, 53, false},
		{4, 16, false},
		{4, 39, false},
		{4, 40, false},
		{4, 57, false},
		{4, 41, false},
		{4, 63, false},
		{4, 64, false},
		{4, 349, false},
		{4, 350, false},
		{4, 351, false},
		{4, 24, false},
		{4, 361, false},
		{4, 369, false},
		{4, 370, true},
		{4, 43, false},
		{4, 365, false},
		{4, 354, false},
		{4, 360, false},
		{4, 355, false},
		{4, 352, false},
		{4, 25, false},
		{4, 110, false},
		{4, 27, false},
		{4, 363, false},
		{4, 35, false},
		{4, 357, false},
		{4, 358, false},
		{4, 32, false},
		{4, 367, false},
		{4, 18, false},
		{4, 17, false},
		{4, 94, false},
		{4, 97, false},
		{4, 108, false},
		{4, 96, false},
		{4, 112, false},
		{4, 100, false},
		{4, 115, false},
		{4, 103, false},
		{4, 114, false},
		{4, 341, false},
		{4, 342, false},
		{4, 102, false},
		{4, 26, false},
		{4, 113, false},
		{4, 343, false},
		{4, 109, false},
		{4, 111, false},
		{4, 20, false},
		{4, 68, false},
		{4, 66, false},
		{4, 71, false},
		{4, 67, false},
		{4, 69, false},
		{4, 347, false},
		{4, 344, false},
		{4, 348, false},
		{4, 346, false},
		{4, 345, false},
		{4, 65, false},
		{4, 50, false},
		{4, 212, false},
		{4, 214, false},
		{4, 21, false},
		{4, 23, false},
		{4, 245, false},
		{4, 246, false},
		{4, 252, false},
		{4, 253, false},
		{4, 255, false},
		{4, 259, false},
		{4, 262, false},
		{4, 263, false},
		{4, 257, false},
		{4, 80, false},
		{4, 90, false},
		{4, 183, false},
		{4, 184, false},
		{4, 38, false},
		{4, 199, false},
		{4, 201, false},
		{4, 200, false},
		{4, 4, true},
		{4, 12, false},
		{4, 79, false},
		{4, 187, false},
		{4, 191, false},
		{4, 192, false},
		{4, 195, false},
		{4, 188, false},
		{4, 196, false},
		{4, 189, false},
		{4, 197, false},
		{4, 194, true},
		{4, 186, false},
		{4, 178, false},
		{4, 182, false},
		{4, 179, false},
		{4, 235, false},
		{4, 233, false},
		{4, 234, false},
		{4, 205, false},
		{4, 206, false},
		{4, 208, false},
		{4, 210, false},
		{4, 211, true},
		{4, 254, false},
		{4, 258, false},
		{4, 256, false},
		{4, 260, false},
		{4, 261, false},
		{66, 322, false},
		{66, 323, true},
		{66, 324, false},
		{66, 29, true},
		{66, 42, true},
		{66, 14, true},
		{66, 39, false},
		{66, 244, true},
		{66, 44, false},
		{66, 45, true},
		{66, 47, false},
		{66, 48, true},
		{66, 49, true},
		{66, 56, false},
		{66, 40, true},
		{66, 57, false},
		{66, 41, true},
		{66, 15, false},
		{66, 53, false},
		{66, 16, false},
		{66, 24, true},
		{66, 25, true},
		{66, 18, false},
		{66, 20, false},
		{66, 236, false},
		{66, 43, true},
		{66, 21, true},
		{66, 19, true},
		{66, 23, true},
		{66, 245, true},
		{66, 246, true},
		{66, 63, false},
		{66, 60, false},
		{66, 64, false},
		{66, 17, false},
		{66, 355, false},
		{66, 359, false},
		{66, 360, false},
		{66, 341, false},
		{66, 342, true},
		{66, 344, true},
		{66, 349, false},
		{66, 350, true},
		{66, 351, true},
		{66, 26, true},
		{66, 27, true},
		{66, 358, false},
		{66, 35, true},
		{66, 352, false},
		{66, 353, true},
		{66, 354, true},
		{66, 345, false},
		{66, 346, true},
		{66, 348, true},
		{66, 32, true},
		{66, 28, true},
		{66, 65, false},
		{66, 68, false},
		{66, 252, false},
		{66, 253, true},
		{66, 66, false},
		{66, 67, false},
		{66, 69, false},
		{66, 50, false},
		{66, 212, true},
		{66, 214, true},
		{66, 86, false},
		{66, 90, false},
		{66, 183, true},
		{66, 184, false},
		{66, 255, false},
		{66, 259, false},
		{66, 257, false},
		{66, 262, false},
		{66, 263, false},
		{66, 80, false},
		{66, 38, true},
		{66, 199, true},
		{66, 201, true},
		{66, 186, true},
		{66, 200, true},
		{66, 191, true},
		{66, 192, false},
		{66, 194, true},
		{66, 188, false},
		{66, 187, false},
		{66, 6, true},
		{66, 7, true},
		{66, 12, true},
		{66, 79, false},
		{66, 196, false},
		{66, 197, false},
		{66, 195, false},
		{66, 189, false},
		{66, 256, false},
		{66, 254, false},
		{66, 258, false},
		{66, 260, false},
		{66, 261, false},
		{66, 178, false},
		{66, 182, false},
		{66, 235, false},
		{66, 179, false},
		{66, 233, false},
		{66, 234, false},
		{66, 361, false},
		{15, 322, false},
		{15, 323, false},
		{15, 324, false},
		{15, 29, false},
		{15, 42, false},
		{15, 14, false},
		{15, 20, false},
		{15, 18, false},
		{15, 24, false},
		{15, 26, false},
		{15, 15, false},
		{15, 25, false},
		{15, 236, false},
		{15, 43, false},
		{15, 44, false},
		{15, 56, false},
		{15, 39, false},
		{15, 40, false},
		{15, 57, false},
		{15, 41, false},
		{15, 16, false},
		{15, 53, false},
		{15, 51, false},
		{15, 55, false},
		{15, 63, false},
		{15, 64, false},
		{15, 60, false},
		{15, 349, false},
		{15, 353, false},
		{15, 354, false},
		{15, 35, false},
		{15, 343, false},
		{15, 344, false},
		{15, 345, true},
		{15, 27, false},
		{15, 352, false},
		{15, 346, false},
		{15, 347, true},
		{15, 348, false},
		{15, 341, false},
		{15, 342, false},
		{15, 68, false},
		{15, 66, false},
		{15, 67, false},
		{15, 69, false},
		{15, 65, false},
		{15, 50, false},
		{15, 212, false},
		{15, 214, false},
		{15, 252, false},
		{15, 253, false},
		{15, 255, false},
		{15, 259, false},
		{15, 257, false},
		{15, 262, false},
		{15, 263, false},
		{15, 21, false},
		{15, 23, false},
		{15, 77, false},
		{15, 90, false},
		{15, 183, false},
		{15, 184, false},
		{15, 38, false},
		{15, 199, false},
		{15, 17, false},
		{15, 201, false},
		{15, 200, false},
		{15, 186, false},
		{15, 191, false},
		{15, 192, false},
		{15, 194, false},
		{15, 188, false},
		{15, 196, false},
		{15, 254, false},
		{15, 258, false},
		{15, 256, false},
		{15, 260, false},
		{15, 261, false},
		{15, 4, true},
		{15, 12, false},
		{15, 79, false},
		{15, 80, false},
		{15, 189, false},
		{15, 197, false},
		{15, 195, false},
		{15, 178, false},
		{15, 182, false},
		{15, 235, false},
		{15, 233, false},
		{15, 234, false},
		{15, 13, true},
		{15, 187, false},
		{15, 355, false},
		{14, 322, false},
		{14, 323, true},
		{14, 324, false},
		{14, 29, true},
		{14, 42, true},
		{14, 14, true},
		{14, 44, false},
		{14, 45, true},
		{14, 56, false},
		{14, 15, false},
		{14, 39, false},
		{14, 40, true},
		{14, 57, false},
		{14, 41, true},
		{14, 53, false},
		{14, 51, false},
		{14, 16, false},
		{14, 52, false},
		{14, 55, false},
		{14, 18, false},
		{14, 20, false},
		{14, 21, true},
		{14, 23, true},
		{14, 63, false},
		{14, 60, false},
		{14, 64, false},
		{14, 353, false},
		{14, 35, true},
		{14, 43, true},
		{14, 355, false},
		{14, 356, false},
		{14, 342, false},
		{14, 343, false},
		{14, 344, true},
		{14, 345, true},
		{14, 346, true},
		{14, 350, false},
		{14, 349, false},
		{14, 24, true},
		{14, 25, true},
		{14, 27, true},
		{14, 347, false},
		{14, 348, true},
		{14, 354, false},
		{14, 352, false},
		{14, 351, false},
		{14, 68, false},
		{14, 66, false},
		{14, 69, false},
		{14, 67, false},
		{14, 65, false},
		{14, 50, false},
		{14, 212, true},
		{14, 214, true},
		{14, 252, false},
		{14, 253, true},
		{14, 255, false},
		{14, 259, false},
		{14, 83, false},
		{14, 28, true},
		{14, 32, true},
		{14, 257, false},
		{14, 262, false},
		{14, 263, false},
		{14, 235, false},
		{14, 17, false},
		{14, 178, false},
		{14, 182, false},
		{14, 233, false},
		{14, 234, false},
		{14, 80, false},
		{14, 90, false},
		{14, 183, true},
		{14, 197, false},
		{14, 184, false},
		{14, 189, false},
		{14, 38, true},
		{14, 187, false},
		{14, 188, false},
		{14, 191, false},
		{14, 192, false},
		{14, 196, false},
		{14, 195, false},
		{14, 199, true},
		{14, 26, true},
		{14, 254, false},
		{14, 258, false},
		{14, 256, false},
		{14, 260, false},
		{14, 261, false},
		{14, 4, true},
		{14, 12, true},
		{14, 13, true},
		{14, 79, false},
		{14, 179, false},
		{14, 180, true},
		{14, 186, true},
		{14, 201, true},
		{14, 200, true},
		{14, 194, true},
		{14, 357, true},
		{14, 19, false},
		{14, 203, true},
		{14, 236, false},
		{14, 213, false},
		{14, 216, true},
	}
	return characterData
}