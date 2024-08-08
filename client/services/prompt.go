package services

import (
	"fmt"
	"log"
	"protobuffer/interfaces"
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
	fmt.Println("-----------------------------------------------")
	fmt.Scan(&option)
	return option
}

func ExecuteOptionSelected(option int) *interfaces.Request {
	if option == 1 {
		log.Println("[Movie] Creating movie")
		return Create(CreateMovie(), CreateFilter())
	}
	if option == 2 {
		log.Println("[Movie] Finding by ator")
		return FindByAtor()
	}
	if option == 3 {
		log.Println("[Movie] Finding by categoria")
		return FindByCategoria()
	}
	return FindByAtor()
}
