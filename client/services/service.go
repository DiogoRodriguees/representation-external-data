package services

import "protobuffer/interfaces"

func Create(movie *interfaces.Movie, filters *interfaces.MovieFilters) *interfaces.Request {
	return &interfaces.Request{
		Method:  "CREATE",
		Movie:   movie,
		Filters: filters,
	}
}

func FindByAtor() *interfaces.Request {
	return &interfaces.Request{
		Method: "FIND_BY_ATOR",
		Movie:  &interfaces.Movie{},
		Filters: &interfaces.MovieFilters{
			Values: []string{"Frank Powell"},
		},
	}
}
func FindByCategoria() *interfaces.Request {
	return &interfaces.Request{
		Method: "FIND_BY_CATEGORIA",
		Movie:  &interfaces.Movie{},
		Filters: &interfaces.MovieFilters{
			Values: []string{"Drama"},
		},
	}
}

func CreateMovie() *interfaces.Movie {
	return &interfaces.Movie{
		Title:  "O Poderoso Chefão",
		Genres: []string{"Ação", "Comédia"},
	}
}

func CreateFilter() *interfaces.MovieFilters {
	return &interfaces.MovieFilters{
		Values: []string{"Ação"},
	}
}
