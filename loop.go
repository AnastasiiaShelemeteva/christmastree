package main

import (
	"fmt"
	"strings"
)

func main() {
	star := "*"
	heigth := 10
	space := "1"
	

 for i := heigth; i > 0; i-- {
	repSpace := strings.Repeat(space, i)
	
	fmt.Println(repSpace, star)
	}

}