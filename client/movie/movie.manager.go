package movie

import (
	"protobuffer/interfaces"
)

type MovieManager struct {
	Option map[int]func() *interfaces.Request
}

func Init() *MovieManager {
	movie := &MovieManager{}
	movie.Option[1] = ExecuteCreate
	movie.Option[2] = ExecuteFindByActor
	movie.Option[3] = ExecuteFindByCategorie
	movie.Option[5] = ExecuteUpdate
	movie.Option[6] = ExecuteDelete

	return movie
}
