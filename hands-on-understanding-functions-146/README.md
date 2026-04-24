# Hands-on 146: Function Expressions

## Purpose
Show how a function can be assigned to a variable and then called through that variable.

## What This Program Shows
- Calling the normal named function `foo()`.
- Assigning an anonymous function to variable `x`.
- Assigning an anonymous function with a string parameter to variable `y`.
- Executing both stored functions later with `x()` and `y("Neerav")`.

## Run
```bash
go run .
```

## Expected Behavior
The program prints output from `foo()`, then runs the two function expressions stored in `x` and `y`.

## Key Takeaway
In Go, functions are values, so they can be stored in variables and passed around like other values.
