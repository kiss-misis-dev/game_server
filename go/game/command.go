package game

type Command struct {
	ID
	members map[Identifier]*Member
}

func NewCommand() Command {
	return Command{NewID(), make(map[Identifier]*Member)}
}

func (command *Command) AttachMember(member *Member) {
	command.members[member.id] = member
	member.commands[command.id] = command
}

func (command *Command) RemoveMember(memberId Identifier) {
	delete(command.members[memberId].commands, command.id)
	delete(command.members, memberId)
}
