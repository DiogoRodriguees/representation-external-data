package movie

import (
	"fmt"
	"protobuffer/interfaces"
	"strings"
)

func CreateMovie() *interfaces.Movie {
	return &interfaces.Movie{
		Title:     "O Poderoso Chefão",
		Genres:    []string{},
		Cast:      []string{},
		Directors: []string{},
	}
}

func CreateFilter() *interfaces.MovieFilters {
	return &interfaces.MovieFilters{
		Values: []string{"Ação"},
	}
}

func ShowMenu() int {
	var option int
	fmt.Println("\n-----------------------------------------------")
	fmt.Println("*              O que voce deseja?             *")
	fmt.Println("-----------------------------------------------")
	fmt.Println("1. Criar um filme")
	fmt.Println("2. Listar por ator")
	fmt.Println("3. Listar por categoria")
	fmt.Println("4. Sair")
	fmt.Println("5. Atualizar filme")
	fmt.Println("6. Excluir filme")
	fmt.Println("-----------------------------------------------")
	fmt.Scan(&option)
	return option
}

func ExecuteCreate() *interfaces.Request {
	m := &Movie{}
	request := m.Create()

	movie := &interfaces.Movie{}
	fmt.Println("Insert movie title: ")
	fmt.Scan(&movie.Title)

	var directors string
	fmt.Println("Insert movie directors separeteds per comma (ex: diogo,gustavo,christofer): ")
	fmt.Scan(&directors)
	movie.Directors = strings.Split(directors, ",")

	var genres string
	fmt.Println("Insert movie genres separeteds per comma (ex: drama,suspense): ")
	fmt.Scan(&genres)
	movie.Genres = strings.Split(directors, ",")

	var cast string
	fmt.Println("Insert movie cast separeteds per comma (ex: diogo,gustavo,christofer): ")
	fmt.Scan(&cast)
	movie.Cast = strings.Split(directors, ",")

	request.Movie = movie
	return request
}

func ExecuteUpdate() *interfaces.Request {
	m := &Movie{}
	request := m.Update()

	fmt.Println("Insert movie id: ")
	fmt.Scan(&request.Movie.Id)

	fmt.Println("Insert movie title: ")
	fmt.Scan(&request.Movie.Title)

	var directors string
	fmt.Println("Insert movie directors separeteds per comma (ex: diogo,gustavo,christofer): ")
	fmt.Scan(&directors)
	request.Movie.Directors = strings.Split(directors, ",")

	var genres string
	fmt.Println("Insert movie genres separeteds per comma (ex: drama,suspense): ")
	fmt.Scan(&genres)
	request.Movie.Genres = strings.Split(directors, ",")

	var cast string
	fmt.Println("Insert movie cast separeteds per comma (ex: diogo,gustavo,christofer): ")
	fmt.Scan(&cast)
	request.Movie.Cast = strings.Split(directors, ",")

	return request
}

func ExecuteDelete() *interfaces.Request {
	m := &Movie{}
	request := m.Delete()
	fmt.Println("Insert movie id: ")
	fmt.Scan(&request.Movie.Id)
	return request
}

func ExecuteFindByActor() *interfaces.Request {
	m := &Movie{}
	request := m.FindByAtor()

	var actors string
	fmt.Println("Insert movie cast separeteds per comma (ex: diogo,gustavo,christofer): ")
	fmt.Scan(&actors)

	request.Filters.Values = strings.Split(actors, ",")
	return request
}

func ExecuteFindByCategorie() *interfaces.Request {
	m := &Movie{}
	request := m.FindByCategoria()

	var genres string
	fmt.Println("Insert movie genres separeteds per comma (ex: drama,suspense): ")
	fmt.Scan(&genres)

	request.Filters.Values = strings.Split(genres, ",")
	return request
}
