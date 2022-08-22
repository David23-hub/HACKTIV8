package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Print(i)
		if i%2 == 0 {
			fmt.Println(" Genap")
		} else {
			fmt.Println(" Ganjil")
		}
	}

}
