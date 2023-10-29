package game

type ID struct {
	id int64
}

var idIota int64 = 0

func NewID() ID {
	idIota++
	return ID{idIota}
}
