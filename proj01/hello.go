package main

import (
	"demo/proj01/lib"
	"fmt"
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
