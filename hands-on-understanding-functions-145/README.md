# Hands-on 145: Anonymous Functions

## Purpose
Demonstrate how anonymous functions work in Go, including immediate invocation and passing arguments directly to an unnamed function.

## What This Program Shows
- Calling a regular named function with `foo()`.
- Defining and immediately executing an anonymous function with no parameters.
- Defining and immediately executing an anonymous function that accepts a string argument.

## Run
```bash
go run .
```

## Expected Behavior
The program prints output from the named `foo()` function, then prints messages from two anonymous functions.

## Key Takeaway
Anonymous functions are useful when you want small one-off behavior without creating a separately named function.
