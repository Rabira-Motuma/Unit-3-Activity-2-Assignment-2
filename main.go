/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date: 2025-11-25
 * Fileoverview: This program rounds off numbers to a specified number of decimal places.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variables
	var numberOfCentsAsString string
	var numberOfCentsAsNumber int
	var dollars int
	var cents int

	reader := bufio.NewReader(os.Stdin)

	// input
	fmt.Print("Input the cents please: ")
	numberOfCentsAsString, _ = reader.ReadString('\n')
	numberOfCentsAsString = strings.TrimSpace(numberOfCentsAsString)

	// process
	numberOfCentsAsNumber, _ = strconv.Atoi(numberOfCentsAsString)
	dollars = numberOfCentsAsNumber / 100
	cents = numberOfCentsAsNumber % 100

	// output
	fmt.Println()
	fmt.Printf(`That is %d dollars and %d cents`, dollars, cents)

	fmt.Println("\nDone.")
}
