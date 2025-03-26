package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/startup"
)

type HelloWorld struct {
	classdb.Extension[HelloWorld, Node2D.Instance]
}

func (h *HelloWorld) Ready() {
	fmt.Println("Hello World from Go!")
}

func main() {
	classdb.Register[HelloWorld]()
	startup.Scene()
}
