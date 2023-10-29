package game

type Command struct {
	ID
	members map[int64]*Member
}

func NewCommand() Command {
	return Command{NewID(), make(map[int64]*Member)}
}

func AddMember(command *Command, member *Member) {
	command.members[member.id] = member
	member.commands[command.id] = command
}
