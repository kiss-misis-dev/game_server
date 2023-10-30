package game

type Problem struct {
	Situation ProblemSituation
}

type ProblemSituation interface {
	// TODO:
}

func (round *Round) NewProblem(situation ProblemSituation) {
	round.problems = append(round.problems, Problem{situation})
}
