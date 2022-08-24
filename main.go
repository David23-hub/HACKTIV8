package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	name    string
	address string
	job     string
	reason  string
}

func main() {

	poeple := []Person{
		{
			name:    "david",
			address: "Tanjung Duren",
			job:     "Bos Pertamina",
			reason:  "iseng",
		},
		{
			name:    "fadli",
			address: "Cempaka Putih",
			job:     "Dirut Pertamina",
			reason:  "iseng juga",
		},
	}

	args := os.Args
	indexArgs := args[1]
	index, _ := strconv.Atoi(indexArgs)

	for i, t := range poeple {
		if i == index-1 {
			fmt.Println(t)
		}
	}
}
