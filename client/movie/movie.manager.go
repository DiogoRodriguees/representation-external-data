package movie

import (
	"protobuffer/interfaces"
)

type MovieManager struct {
	Option map[int]func() *interfaces.Request
}

func Init() *MovieManager {
	options := make(map[int]func() *interfaces.Request)
	options[1] = ExecuteCreate
	options[2] = ExecuteFindByActor
	options[3] = ExecuteFindByCategorie
	options[5] = ExecuteUpdate
	options[6] = ExecuteDelete

	return &MovieManager{Option: options}
}
