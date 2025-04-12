package lexer

import (
	"github.com/Mixturka/GoCC/internal/lexer/tables"
)

type Lexer struct {
	classifierTable map[int]int
	transitionTable [][]int
	tokenTypeTable  map[int]struct {
		Name     string
		Priority int
	}
}

func NewLexer() *Lexer {
	return &Lexer{
		classifierTable: tables.ClassifierTable,
	}
}
