package main

import (
	"fmt"
)

type value int
type slice []value

func genSlice(s int) slice {
	vault := slice{}

	for i := 1; i <= s; i++ {
		vault = append(vault, value(i))
	}

	return vault
}

func (s slice) print() {
	for _, number := range s {
		fmt.Println(number)
	}
}

func (v value) isOddOrEven() string {
	if v%2 != 0 {
		return "Odd"
	} else {
		return "Even"
	}
}

func printOddOrEven(s slice) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%02v\t ===> \t %02v \n", s[i], s[i].isOddOrEven())
	}
}
