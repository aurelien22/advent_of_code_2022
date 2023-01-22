package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	total := 0
	m := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 2}
	winMap := map[string]string{"A": "B", "B": "C", "C": "A"}
	looseMap := map[string]string{"A": "C", "B": "A", "C": "B"}

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(data), "\n")

	for i := 0; i < len(lines); i++ {

		line := strings.Split(lines[i], " ")

		switch line[1] {
		case "X":
			total += m[looseMap[line[0]]]
		case "Y":
			total += m[line[0]] + 3
		case "Z":
			total += m[winMap[line[0]]] + 6
		}
	}

	fmt.Println(total)
}
