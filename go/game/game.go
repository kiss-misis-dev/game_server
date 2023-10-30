package game

type Game struct {
	commands map[Identifier]*Command
	rounds   []Round
	// Points
}

func NewGame() Game {
	return Game{}
}

func (game *Game) RegisterCommand(command *Command) {
	game.commands[command.id] = command
}

func (game *Game) UnregisterCommand(commandId Identifier) {
	delete(game.commands, commandId)
}

func (game *Game) AddRound(round Round) {
	game.rounds = append(game.rounds, round)
}
