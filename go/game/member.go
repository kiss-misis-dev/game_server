package game

type Member struct {
	ID
	commands map[Identifier]*Command
}

func NewMember() Member {
	return Member{NewID(), make(map[Identifier]*Command)}
}
