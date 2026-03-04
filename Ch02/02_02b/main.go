package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42

	fmt.Println(str1, str2, str3)
	//When you call the print function you can get two return values such as
	// Length of the String and the Error for that Println
	stringLength, err := fmt.Println("The value is", aNumber)

	if err == nil {
		fmt.Println("String length is:", stringLength)
	}

	fmt.Printf("Value of number: %v\n", aNumber)
	fmt.Printf("Data type: %T\n", aNumber)
}
