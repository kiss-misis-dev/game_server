package game

type Identifier int64

type ID struct {
	id Identifier
}

var idIota Identifier = 0

func NewID() ID {
	idIota++
	return ID{idIota}
}

func (id *ID) GetID() Identifier {
	return id.id
}
