package movie

import "protobuffer/interfaces"

type Movie struct{}

func (m *Movie) Create() *interfaces.Request {
	return &interfaces.Request{
		Method:  "CREATE",
		Movie:   &interfaces.Movie{},
		Filters: &interfaces.MovieFilters{},
	}
}

func (m *Movie) Update() *interfaces.Request {
	return &interfaces.Request{
		Method:  "UPDATE",
		Movie:   &interfaces.Movie{},
		Filters: &interfaces.MovieFilters{},
	}
}

func (m *Movie) Delete() *interfaces.Request {
	return &interfaces.Request{
		Method:  "DELETE",
		Movie:   &interfaces.Movie{},
		Filters: &interfaces.MovieFilters{},
	}
}

func (m *Movie) FindByAtor() *interfaces.Request {
	return &interfaces.Request{
		Method: "FIND_BY_ATOR",
		Movie:  &interfaces.Movie{},
		Filters: &interfaces.MovieFilters{
			Values: []string{"Frank Powell"},
		},
	}
}
func (m *Movie) FindByCategoria() *interfaces.Request {
	return &interfaces.Request{
		Method: "FIND_BY_CATEGORIA",
		Movie:  &interfaces.Movie{},
		Filters: &interfaces.MovieFilters{
			Values: []string{"Drama"},
		},
	}
}
