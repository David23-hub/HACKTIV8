package main

import "fmt"

type person struct {
	name string
}

func main() {

	names := []string{"david", "fadli", "rafi", "fisa"}

	var friends []*person

	for i := 0; i < len(names); i++ {

		person := person{
			name: names[i],
		}

		friends = append(friends, &person)
	}

	function := func(teman []*person) {
		for _, t := range teman {
			fmt.Println(t.name)
		}

	}

	function(friends)
}
