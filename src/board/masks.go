package board

var Win_masks = [...]Bitboard{
	// rows
	Bitboard(4123168604160),
	Bitboard(515396075520),
	Bitboard(32212254720),
	Bitboard(4026531840),
	Bitboard(251658240),
	Bitboard(31457280),
	Bitboard(1966080),
	Bitboard(245760),
	Bitboard(15360),
	Bitboard(1920),
	Bitboard(120),
	Bitboard(15),
	Bitboard(1030792151040),
	Bitboard(2061584302080),
	Bitboard(16106127360),
	Bitboard(8053063680),
	Bitboard(62914560),
	Bitboard(125829120),
	Bitboard(983040),
	Bitboard(491520),
	Bitboard(3840),
	Bitboard(7680),
	Bitboard(60),
	Bitboard(30),
	// cols
	Bitboard(135274560),
	Bitboard(67637280),
	Bitboard(33818640),
	Bitboard(16909320),
	Bitboard(8454660),
	Bitboard(4227330),
	Bitboard(2113665),
	Bitboard(2216338391040),
	Bitboard(1108169195520),
	Bitboard(554084597760),
	Bitboard(277042298880),
	Bitboard(138521149440),
	Bitboard(69260574720),
	Bitboard(34630287360),
	Bitboard(270549120),
	Bitboard(541098240),
	Bitboard(1082196480),
	Bitboard(2164392960),
	Bitboard(4328785920),
	Bitboard(8657571840),
	Bitboard(17315143680),
	// Diags
	Bitboard(279241031680),
	Bitboard(139620515840),
	Bitboard(2181570560),
	Bitboard(69810257920),
	Bitboard(17043520),
	Bitboard(34905128960),
	Bitboard(8521760),
	Bitboard(272696320),
	Bitboard(4260880),
	Bitboard(2130440),
	Bitboard(134744072),
	Bitboard(17247241216),
	Bitboard(67372036),
	Bitboard(2207646875648),
	Bitboard(33686018),
	Bitboard(1103823437824),
	Bitboard(16843009),
	Bitboard(551911718912),
	Bitboard(2155905152),
	Bitboard(275955859456),
	Bitboard(8623620608),
	Bitboard(4311810304),
	Bitboard(545392640),
	Bitboard(1090785280),
}

func CheckDraw(b Board) bool {
	if CheckAlign(b.Bitboards[0]) || CheckAlign(b.Bitboards[1]) {
		return false
	}
	return len(GetMoves(b)) == 0
}

func CheckAlign(bb Bitboard) bool {
	for _, mask := range Win_masks {
		if mask&bb == mask {
			return true
		}
	}
	return false
}

// count the number of ways the player can still win
func WinsRemaining(playerbb, oppbb Bitboard) int8 {
	var remaining int8 = 69
	for _, mask := range Win_masks {
		if mask&oppbb != 0 {
			remaining--
		}
	}
	return remaining
}
