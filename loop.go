package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)



func main() {

fmt.Print("How tall would you like to have your Christmas Tree?")	
input := bufio.NewScanner(os.Stdin)
input.Scan()
height, _ := strconv.Atoi(input.Text())

// in the first loop we are setting the heigth of our tree
 for i := 1; i <= height; i++ {
		
// in the second loop we are printing spaces for each line
// it means, we are starting to count at 1 and print a space every iteration until space becomes height-1
	for space := 1; space <= height - i; space++ {
		
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
for j := 0; j < height - 1; j++ {
	fmt.Print(" ")
} 

fmt.Println("I")
	}

