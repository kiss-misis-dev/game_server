package main

import "github.com/kiss_misis/game_server/game"

func main() {
	command1 := game.NewCommand()
	command2 := game.NewCommand()

	machine := game.NewStateMachine()
	machine.Game.RegisterCommand(&command1)
	machine.Game.RegisterCommand(&command2)

	round := game.NewRound(struct{}{})
	round.NewProblem("you need todo something")
	round.NewProblem("press kachat")

	machine.Game.AddRound(round)
}
