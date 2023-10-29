package game

type Member struct {
	ID
	commands map[int64]*Command
}

func NewMember() Member {
	return Member{NewID(), make(map[int64]*Command)}
}
