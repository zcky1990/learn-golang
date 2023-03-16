package main

import (
	"fmt"
	
	"lumachrome.com/main/mypackage"
	"lumachrome.com/main/mypackage/mysubpackage"
)

func main() {
    fmt.Println("Hello, world.")
	mypackage.Package()
	mysubpackage.Mysubpackage()
}