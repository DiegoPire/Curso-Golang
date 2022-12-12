package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de ponteiros

	var p *int = nil

	p = &i // Pegando o endereço da variável
	*p++   // desferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, i, &i)
}
