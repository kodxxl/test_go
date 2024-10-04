package main

import (
	"fmt" 
	"math/rand"
	"github.com/kodxxl/modforimport/pack1"
)

func main() {
	pack1.Add(1, 2)
	fmt.Println("My favorite number is", pack1.Add(rand.Intn(10),rand.Intn(10)) )
}
