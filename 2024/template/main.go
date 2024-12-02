package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input, err := transformInput(input, toIntArray)
	if err != nil {
		fmt.Printf("oops, there was an error: %s\n", err.Error())
	}
	fmt.Println("input:", input)
	count := 0
	for _, v := range input {
		count += v
	}
	fmt.Println("result:", count)
}

/** Premade functions **/
func toByteArray(acc []byte, input string) ([]byte, error) {
	return append(acc, []byte(input)...), nil
}

func toIntArray(acc []int, input string) ([]int, error) {
	v, err := strconv.Atoi(input)
	if err != nil {
		return nil, err
	}
	return append(acc, v), nil
}

func transformInput[T any](input string, transformFunc func(T, string) (T, error)) (result T, err error) {
	var rows []string = strings.Split(input, "\n")
	for _, line := range rows[:len(rows)-1] {
		result, err = transformFunc(result, line)
		if err != nil {
			return result, err
		}
	}
	return
}
