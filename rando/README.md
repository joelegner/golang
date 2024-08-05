# Rando

Rando is a simple command line program I use to help me learn programming languages. It is fun because it does something plausibly "useful". It generates a random integer between 1 and 100 or within a range you can specify with command line arguments. 

Rando is a program I keep re-writing in new programming languages. This time I am using Go. 

We need to parse the command line arguments. That's how we know the range to use. If no command line arguments are given, we use the default 1-100. If one number is given, assume it is the upper limit and use 1 for the lower limit. If two numbers are given, assume the first is the lower limit and the second is the upper limit. 

Let's think through this argument parsing problem. We already have the list of arguments.

We want the program to fail if:
1. the arguments are not integers
2. the second argument is larger than the first, or
3. either argument is negative

- CN: Generate a random integer between 1 and 100 or within a range you can specify with command line arguments. FR: Generates a random integer between 1 and 100 or within a range you can specify with command line arguments. DP: Rando.
- CN: To use a command line program named `rando` for this purpose. FR: Install the program. DP: `go install rando`.
  - FR: Enable `go install rando` to work. DP: Development environment.
    - FR: Provide an integrated development environment. DP: Microsoft Visual Studio Code (or "VSCode").
    - FR: Manage all development environment files. DP: Repository `https://github.com/joelegner/golang/tree/main/rando`
    - FR: Manage source code. DP: `git` client applications.
  - FR: Host source code in an online repository. DP: `https://github.com` web service. 
  - FR: Make it possible to install the program from the source code. DP: Golang installation. 
  - FR: Install the program from the source code. DP: `go install rando` command line.  
- CN: Tell me how to use the program. FR: Print usage instructions to the screen. DP: Code to handle `-h` and `--help` flags.
  - FR: Process arguments to check for flags and act upon them. DP: `processFlags()`. 
    - FR: Cycle through the arguments. Loop over `os.Args` in `processFlags()` function.
    - FR: Check arguments to see if `-h` or `--help` given. `if` statements in `processFlags()` function.
    - FR: Print the instructions if flags given. DP: Call to `printUsage()` function in `processFlags()` function.
  - FR: Print usage instructions to the screen. DP: `printUsage()` function.
- CN: Treat command line arguments like it says on the tin. FR: Parse command line arguments. DP: Parse command line arguments code. 
  - FR: If zero arguments are given, use a default range of 1-100. DP: Assign default values code. 
  - FR: If one argument is given, assume it is the high end of the range. DP: Assign default low value code. 
  - FR: If two argument is given, assume they are the low and high ends of the range, respectively. DP: Code that interprets two arguments. 
- FR: Accept zero, one, or two integer arguments. DP: Parse arguments code.
- FR: Interpret zero, one, or two arguments according to rules. DP: Argument handling rules. 
  
  
- FR: Convert argument strings to integers. DP: `strconv.Atoi()`. 

- FR: Do the main work of the program. DP: `main()` function. 
- FR: Follow a logical sequence of events. DP: `main()` function.
- FR: Calculate the number. DP: `getNumber()` function.
  - FR: Get the range to use. DP: `getRange()` function.
  - FR: Generate a random number within the specified range. DP: `rand.Intn(high-low+1) + low` statement. 
- FR: Print the number. `main()` function, `fmt` print statement. 
- FR: Print a message if there is an invalid argument given by the user. DP: `printInvalidArg()` function. 
- FR: Print a usage message, giving command, options, and examples. DP: `printUsage()` function. 
- FR: Generate and print a number. DP: `fmt.Println(getNumber())` statement in `main()` function.

```
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
