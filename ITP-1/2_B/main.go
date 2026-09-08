package main

import "fmt"

func main() {

	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if (a < c) && (b < c) {
		if a < b {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else {
		fmt.Println("No")
	}
}
