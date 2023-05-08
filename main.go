package main

import "fmt"

func main() {
	// STRINGS
	// ---------------------------------------------------------
	var nameOne string = "mario"
	// Type inference
	var nameTwo = "luigi"
	// Undefined variable
	var nameThree string

	// Reassign variables
	nameOne = "peach"
	nameThree = "bowser"

	// Short hand for declaring and assigning variables
	// Only works inside functions
	nameFour := "yoshi"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)


	// INTEGERS
	// ---------------------------------------------------------
	// All numeric types https://go.dev/ref/spec#Numeric_types
	var ageOne int = 20
	// Type inference
	var ageTwo = 30
	// Shorthand
	ageThree := 40

	// Bits and memory
	// 8 bit integer range -128 to 127
	var numOne int8 = 127
	// 16 bit integer range -32768 to 32767
	var numTwo int16 = 32767
	// Unsigned integers === only positive numbers
	var numThree uint = 65535

	// FLOATS
	// ---------------------------------------------------------
	// 32 bit float === 6 decimal places
	var scoreOne float32 = 26.123456
	// 64 bit float === 15 decimal places
	var scoreTwo float64 = 356789123456789654.4567891234567891011121314151
	// Type inference === 64 bit float
	scoreThree := 1.5


	// PRINTING VARIABLES
	// ---------------------------------------------------------
	// Print doe not add a new line
	fmt.Print("Hello, " + nameOne + "!")

	// Println adds a new line
	fmt.Println("Hello, " + nameTwo + "!")
	// \n also adds new line
	fmt.Println("Hello, " + nameThree + "!\n")
	// Add variables without + sign adds space
	fmt.Println("Hello,", nameFour, "!")

	// Printf (formatted string) %_ = format specifier
	// %v = value in default format
	fmt.Printf("My age is %v and my score is %v \n", ageOne, scoreOne)
	// %q = double-quoted string safely escaped with Go syntax
	fmt.Printf("My name is %q \n", nameOne)
	// %T = type of the value
	fmt.Printf("Age is of type %T, and name is of type %T \n", ageOne, nameOne)
	// %f = floating point number
	fmt.Printf("You scored %f points! \n", scoreOne)
	// %0.2f = floating point number with 2 decimal places
	fmt.Printf("You scored %.2f points! \n", 25.13)

	// Sprintf (save formatted string in a variable)
	var str = fmt.Sprintf("My age is %v and my score is %v \n", ageOne, scoreOne)
	fmt.Println("The saved string is:", str)


	// ARRAYS
	// ---------------------------------------------------------
	// Arrays are fixed length
	// Arrays must be of the same type
	// Arrays are rarely used in Go
	var ages [3]int = [3]int{20, 25, 30}
	ages[1] = 21
	// Shorthand
	names := [4]string{"yoshi", "mario", "peach", "bowser"}

	// Slices
	// ---------------------------------------------------------
	// Slices are dynamic
	// Slices can contain different types
	// Slices are used 99% of the time
	var scores = []int{100, 50, 60}
	// Shorthand
	scoresTwo := []int{1, 2, 3, 4, 5}
	// Append to slice
	scoresTwo = append(scoresTwo, 6)
	fmt.Println(scoresTwo, len(scoresTwo))

	// Slice ranges
	// name two and three
	rangeOne := names[1:3]
	// name two until the end
	rangeTwo := names[2:]
	// start until name three (not included)
	rangeThree := names[:3]
	// all names
	rangeFour := names[:]
	fmt.Println(rangeOne, rangeTwo, rangeThree, rangeFour)













	fmt.Println(ageOne, ageTwo, ageThree, numOne, numTwo, numThree, scoreOne, scoreTwo, scoreThree, ages, names, scores)
}