package main

import (
	"fmt"
	"os"
)

func main() {
    var numArgs int = len(os.Args)
    var upperLimit int = 100    // Default

    if numArgs > 2 {
        fmt.Println("Usage:")
    }
    switch numArgs {
    case: 1

    }
	fmt.Println(len(os.Args), os.Args)
}
