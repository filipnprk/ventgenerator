package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var ventText = "Hello. Please enter the number of times you would like to vent: "
var ventPhrase = "THIS IS AN ABOMINATION AND YOU CAN FORGET ABOUT IT"

func main() {
	fmt.Print(ventText)
	ventInput, _ := reader.ReadString('\n')

	ventInput = strings.Replace(ventInput, "\n", "", -1)

	vent, _ := strconv.ParseFloat(ventInput, 64)
	var integerVent = int(vent)

	// for loop to print out the vent number of venting the user requested
	for i := 1; i <= integerVent; i++ {
		fmt.Println(ventPhrase)
	}
}
