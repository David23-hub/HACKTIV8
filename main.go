package main

import "fmt"

func main() {

	david := "david"
	kevin := "fadli"

	friends := []*string{&david, &kevin}

	function := func(teman []*string) {
		for _, t := range teman {
			fmt.Println(*t)
		}

	}

	function(friends)
}
