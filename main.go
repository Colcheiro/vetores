package main

import "fmt"

func main() {

	nomes := []string{"Alice", "Bruno", "Carlos", "Diana", "Eduardo"}

	fmt.Println("Dois primeiros:", nomes[:2])

	fmt.Println("Dois últimos:", nomes[len(nomes)-2:])

	fmt.Println("Nome do meio:", nomes[2:3])
}
