package main

import "fmt"

func main() {
	var tate, yoko int
	fmt.Scan(&tate, &yoko)
	var m, s int
	m = menseki(tate, yoko)
	s = syuu(tate, yoko)

	fmt.Printf("%v %v\n", m, s)
}
func menseki(tate int, yoko int) int {
	var result int
	result = tate * yoko

	return result
}
func syuu(tate int, yoko int) int {
	var result int
	result = (tate * 2) + (yoko * 2)
	return result
}
