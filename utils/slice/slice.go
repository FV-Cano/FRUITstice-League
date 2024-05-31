package slice

import (
	"utils/globals"
)

/**
* @desc: Filtra un slice según una condición dada

* @param slice: slice a filtrar
* @param condition: booleano
* @return: slice filtrado
 */
func Filter[T any](slice []T, condition func(T) bool) []T {
	var filteredSlice []T
	for _, item := range slice {
		if condition(item) {
			filteredSlice = append(filteredSlice, item)
		}
	}
	return filteredSlice
}

/**
* @desc: Ordena un slice de frutas por precio de mayor a menor
 */
type ByPrice []globals.Sale

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrice) Less(i, j int) bool { return a[i].Total > a[j].Total }

/**
* @desc: Determina si un slice contiene un elemento dado

* @param slice: slice a evaluar
* @param item: elemento a buscar
* @return: booleano
 */
func Contains(slice []string, item string) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}