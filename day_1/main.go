package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	result := strings.Split(string(data), "\n")

	var s []int

	total := 0
	//maxNumber := 0

	for i := 0; i < len(result); i++ {
		if result[i] != "" {
			number, err := strconv.Atoi(result[i])
			if err != nil {
				fmt.Println(err)
			}
			total = total + number
		} else {
			s = append(s, total)
			total = 0
		}
	}

	sort.Ints(s)

	fmt.Println(s[len(s)-3] + s[len(s)-2] + s[len(s)-1])
}
