package main

import (
	"fmt"
	"math/rand"
	"os"
)

func getRange() (int, int) {
	return 1, 100
}

func print_usage() {
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

func main() {
	low, high := getRange()
	randomInt := rand.Intn(high-low+1) + low
	fmt.Println(randomInt)

	args := os.Args
	fmt.Printf("Number of Args = %d\n", len(args))
	fmt.Printf("Type of Args = %T\n", args)
	fmt.Println(args)
	print_usage()

	/*
	   Let's think through this argument parsing problem.
	   We already have the list of arguments.
	   We want the program to fail if:
	       1. the arguments are not integers.
	       2. the second argument is larger than the first
	       3. either argument is negative
	   Case: No argument given
	       use default range 1-100
	   Case: One integer given
	       if the integer is n, the range is 1-n.
	       Case: Two integers given
	       if the integers are i and j, the range is i-j.
	   Else: Something else
	       Respond with an error message.
	       Print usage message.
	       Give examples.
	*/
}
