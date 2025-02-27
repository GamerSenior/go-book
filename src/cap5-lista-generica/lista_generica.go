package main

import "fmt"

type ListaGenerica []interface{}

func (lista *ListaGenerica) RemoverIndice(
	indice int) interface{} {

	l := *lista
	removido := l[indice]
	*lista = append(l[0:indice], l[indice+1:]...)
	return removido
}

func (lista *ListaGenerica) RemoverInicio() interface{} {
	return lista.RemoverIndice(0)
}

func (lista *ListaGenerica) RemoverFinal() interface{} {
	return lista.RemoverIndice(len(*lista) - 1)
}

func main() {
	lista := ListaGenerica{
		1, "Café", 42, true, 23, "Bola", 3.14, false,
	}

	fmt.Printf("Lista Original: \n%v\n\n", lista)

	fmt.Printf(
		"Removendo do início: %v, após remoção:\n%v\n",
		lista.RemoverInicio(), lista)
	fmt.Printf(
		"Removendo do índice 3: %v, após remoção:\n%v\n",
		lista.RemoverIndice(3), lista)
	fmt.Printf(
		"Removendo do índice 0: %v, após remoção:\n%v\n",
		lista.RemoverIndice(0), lista)
	fmt.Printf(
		"Removendo do ultimo índice: %v, após remoção:\n%v\n",
		lista.RemoverFinal(), lista)

}
