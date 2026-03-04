package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Simple read to get the user input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your input ...")
	str, _ := reader.ReadString('\n') // in Go if you want to ignore the error place an "_" for its name.
	fmt.Println(str)

	// Simple read to ge the user input int

	fmt.Print("Please enter your number ...")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", f)
	}
}
