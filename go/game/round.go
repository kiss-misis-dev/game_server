package game

type Round struct {
	problems       []Problem
	Representation RoundRepresentation
}

type RoundRepresentation interface {
	// TODO:
}

func NewRound(representation RoundRepresentation) Round {
	return Round{[]Problem{}, representation}
}
