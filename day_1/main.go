package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	result := strings.Split(string(data), "\n")

	total := 0
	maxNumber := 0

	for i := 0; i < len(result); i++ {
		if result[i] != "" {
			number, err := strconv.Atoi(result[i])
			if err != nil {
				fmt.Println(err)
			}
			total = total + number
		} else {
			if total > maxNumber {
				maxNumber = total
			}
			total = 0
		}
	}
	fmt.Println(maxNumber)
}
