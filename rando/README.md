# Rando

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
