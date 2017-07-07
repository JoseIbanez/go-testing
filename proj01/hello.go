package main

import (
	"fmt"
	"go-testing/proj01/lib"
	"log"
)

func main() {
	log.Println("test")
	fmt.Printf("Hello, world!!\n")
	c := lib.Resta(1, 1)
	d := lib.Suma(2, 2)
	fmt.Println(c)
	fmt.Println(d)

}
