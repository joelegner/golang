# numbering

This is an experimental Go program that is intended to take a hierarchy like this:

```
Design metal building foundations | Determine program
    Get input | readFile function
        Modeling Module
            Read input file
            Create model instance
    Process input
        Engineering Module
            Manage design criteria
                Design Criteria Class
                    Return edge_dist
                    Return concrete strength
                    Return rebar yield
                    Return q_a
                    Return T_inc
                    Return T_min
            Obtain model instance
```

And assign FR and DP numbers like this:

```
FR0: Design metal building foundations
DP0: Determine program
FR1: Get input
DP1: readFile function
FR11: Modeling Module
FR111: Read input file
FR112: Create model instance
FR2: Process input
FR21: Engineering Module
FR211: Manage design criteria
FR2111: Design Criteria Class
FR21111: Return edge_dist
FR21112: Return concrete strength
FR21113: Return rebar yield
FR21114: Return q_a
FR21115: Return T_inc
FR21116: Return T_min
FR212: Obtain model instance
```

The format of each line is:

`<functional requirement>[ | <design parameter>]`


