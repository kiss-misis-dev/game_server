package game

type StateMachine struct {
	Game  Game
	state State
}

func NewStateMachine() StateMachine {
	return StateMachine{NewGame(), StateScheduled}
}

func (machine *StateMachine) ApplyEvent(event Event) {
	err := machine.state.HandleEvent(event)
	if err != nil {
		// TODO:
		panic(err.Error())
	}
}
