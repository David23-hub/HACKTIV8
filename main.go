package main

import "fmt"

func main() {
	listFriend()

}

func listFriend() {
	name := []string{"david", "fadli", "fisa", "rafi", "kadek", "kevin", "iqbal", "irfan", "edwin", "jaka"}

	for i, t := range name {
		fmt.Println("Name", i+1, "=", t)
	}
}
