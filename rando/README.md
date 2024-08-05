# Rando

Rando is a simple command line program I use to help me learn programming languages. It is fun because it does something plausibly "useful". It generates a random integer between 1 and 100 or within a range you can specify with command line arguments. 

Rando is a program I keep re-writing in new programming languages. This time I am using Go. 

We need to parse the command line arguments. That's how we know the range to use. If no command line arguments are given, we use the default 1-100. If one number is given, assume it is the upper limit and use 1 for the lower limit. If two numbers are given, assume the first is the lower limit and the second is the upper limit. 

Let's think through this argument parsing problem. We already have the list of arguments.

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

```
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
```

# TODO List

It could benefit from improvements. 

- [ ] Add ability to use -h or --help flag to print the usage instructions.
- [ ] Add verbose mode with -v or --verbose that prints more messages.
- [ ] In verbose mode, print when reverting to default range. 
- [ ] In verbose mode, print when reverting to default lower limit. 
- [ ] In verbose mode, print the range before giving the number: "Random number between 1 and 100: 24.
- [ ] Installation instructions: `go install rando`
