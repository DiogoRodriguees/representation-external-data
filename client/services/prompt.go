package services

import (
	"fmt"
	"log"
	"protobuffer/interfaces"
	"strings"
)

func Prompt() int {
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

func executeCreate() *interfaces.Request {
	request := Create(CreateMovie(), CreateFilter())

	movie := &interfaces.Movie{}
	fmt.Println("Insert movie title: ")
	fmt.Scan(&movie.Title)

	var directors string
	fmt.Println("Insert movie directors separeteds per , (ex: diogo,gustavo,christofer): ")
	fmt.Scan(&directors)
	movie.Directors = strings.Split(directors, ",")

	var genres string
	fmt.Println("Insert movie genres separeteds per , (ex: drama,suspense): ")
	fmt.Scan(&genres)
	movie.Genres = strings.Split(directors, ",")

	var cast string
	fmt.Println("Insert movie cast separeteds per , (ex: diogo,gustavo,christofer): ")
	fmt.Scan(&cast)
	movie.Cast = strings.Split(directors, ",")

	request.Movie = movie
	return request
}

func executeUpdate() *interfaces.Request {
	request := Create(CreateMovie(), CreateFilter())

	movie := &interfaces.Movie{}
	fmt.Println("Insert movie id: ")
	fmt.Scan(&movie.Id)

	fmt.Println("Insert movie title: ")
	fmt.Scan(&movie.Title)

	var directors string
	fmt.Println("Insert movie directors separeteds per , (ex: diogo,gustavo,christofer): ")
	fmt.Scan(&directors)
	movie.Directors = strings.Split(directors, ",")

	var genres string
	fmt.Println("Insert movie genres separeteds per , (ex: drama,suspense): ")
	fmt.Scan(&genres)
	movie.Genres = strings.Split(directors, ",")

	var cast string
	fmt.Println("Insert movie cast separeteds per , (ex: diogo,gustavo,christofer): ")
	fmt.Scan(&cast)
	movie.Cast = strings.Split(directors, ",")

	request.Movie = movie
	return request
}

func executeDelete() *interfaces.Request {
	request := Create(CreateMovie(), CreateFilter())

	movie := &interfaces.Movie{}
	fmt.Println("Insert movie id: ")
	fmt.Scan(&movie.Id)
	return request
}

func ExecuteOptionSelected(option int) *interfaces.Request {
	if option == 1 {
		log.Println("[Movie] Creating movie")
		return executeCreate()
	}
	if option == 2 {
		log.Println("[Movie] Finding by ator")
		return FindByAtor()
	}
	if option == 3 {
		log.Println("[Movie] Finding by categoria")
		return FindByCategoria()
	}
	if option == 5 {
		log.Println("[Movie] Updating movie")
		return executeUpdate()
	}
	if option == 6 {
		log.Println("[Movie] Deleting movie")
		return executeDelete()
	}
	return FindByAtor()
}
