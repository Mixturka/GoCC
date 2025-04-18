package tables

var ClassifierTable = map[rune]int{
	119: 0,
	52:  1,
	57:  1,
	88:  0,
	50:  1,
	103: 0,
	113: 0,
	53:  1,
	95:  0,
	98:  0,
	48:  1,
	77:  0,
	100: 0,
	101: 0,
	107: 0,
	109: 0,
	117: 0,
	56:  1,
	69:  0,
	72:  0,
	106: 0,
	116: 0,
	118: 0,
	121: 0,
	73:  0,
	83:  0,
	84:  0,
	86:  0,
	51:  1,
	55:  1,
	76:  0,
	99:  0,
	108: 0,
	74:  0,
	87:  0,
	89:  0,
	97:  0,
	104: 0,
	68:  0,
	75:  0,
	78:  0,
	81:  0,
	49:  1,
	54:  1,
	82:  0,
	85:  0,
	115: 0,
	102: 0,
	111: 0,
	65:  0,
	67:  0,
	71:  0,
	112: 0,
	70:  0,
	90:  0,
	105: 0,
	122: 0,
	66:  0,
	79:  0,
	80:  0,
	110: 0,
	120: 0,
	114: 0,
}

var TransitionTable = [][]int{
	[]int{0, 0},
	[]int{-1, 1},
	[]int{0, 1},
}

var TokenTypeTable = map[int]struct {
	Name     string
	Priority int
}{
	0: {Name: "T_IDENTIFIER", Priority: 2},
	1: {Name: "T_NUM", Priority: 1},
}
