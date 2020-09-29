package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)



func main() {
	
	var input = bufio.NewScanner(os.Stdin)	
	
	fmt.Print("Would you like to have a star? ")
	input.Scan()
	star := input.Text()
	
	fmt.Print("How tall would you like to have your Christmas Tree? ")	
	input.Scan()
	height, _ := strconv.Atoi(input.Text())

switch star {
default:
	fmt.Println("It's not a command, try again")
	main()
case "Y", "y":
	printStar(height)
	printTree(height)
case "N", "n":
	printTree(height)

}

}

func printTree(x int) {


// in the first loop we are setting the heigth of our tree
 for i := 1; i <= x; i++ {
		
// in the second loop we are printing spaces for each line
// it means, we are starting to count at 1 and print a space every iteration until space becomes height-1
	for space := 1; space <= x - i; space++ {
		
		fmt.Print(" ") 
	}  
// in the third loop we are printing stars for each line
// it means, we are starting count at 1 and print a star every iteration until star becomes i*2-1
	for star := 1; star <= i*2-1; star++ {
		fmt.Print("x")
	}
// here we just print a new line
	fmt.Println()
	
	}

// here we print the bottom part of a tree 
for j := 0; j < x - 1; j++ {
	fmt.Print(" ")
} 

fmt.Println("I")
	}


func printStar(x int) {

for i := 0; i < x - 1; i++ {
	fmt.Print(" ")
	} 
	
	fmt.Println("*")
}