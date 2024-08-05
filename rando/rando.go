package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func processArgs() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "-h" || args[i] == "--help" {
			printUsage()
		}
	}
}

func getRange() (int, int) {
	args := os.Args
	low := 1
	high := 100

	if len(args[1:]) == 1 {
		if h, err := strconv.Atoi(args[1:][0]); err == nil {
			high = h
		} else {
			printInvalidArg()
			printUsage()
			os.Exit(1)
		}
	}

	if len(args[1:]) == 2 {
		if l, err := strconv.Atoi(args[1:][0]); err == nil {
			low = l
		} else {
			printInvalidArg()
			printUsage()
			os.Exit(1)
		}
		if h, err := strconv.Atoi(args[1:][1]); err == nil {
			high = h
		} else {
			printInvalidArg()
			printUsage()
			os.Exit(1)
		}
	}

	return low, high
}

func printInvalidArg() {
	fmt.Println("Invalid argument: " + strings.Join(os.Args[1:], " "))
}

func printUsage() {
	fmt.Println(`
NAME
    rando - generate random integers

SYNOPSIS
    rando [[LOW] HIGH]

DESCRIPTION
    rando generates a number from LOW to HIGH. Default is 1-100.

EXAMPLES
    $ rando
    Generates a number from 1-100.

    $ rando 42
    Generates a number from 1-42.
    
    $ rando 75-90
    Generates a number from 75-90.

OPTIONS
    HIGH
        Set the high end of the range. Optional. Default is 100.
    LOW
        Set the low end of the range. Optional. Default is 1. 
    `)
}

func getNumber() int {
	low, high := getRange()
	return rand.Intn(high-low+1) + low
}

func main() {
	processArgs()
	fmt.Println(getNumber())
}
