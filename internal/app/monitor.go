package app

import "fmt"

func ExecuteCommand(command int) {
	switch command {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa...")
	default:
		fmt.Println("Não conheço este comando")
	}
}
