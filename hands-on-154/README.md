Practice functions that return one value and multiple values.

# Hands-on 154: Function Return Values

## Purpose
Demonstrate how Go functions can return either a single value or multiple values.

## What This Program Shows
- `foo(x int) int` returns one integer.
- `bar(x int, a string) (int, string)` returns both an integer and a string.
- `main()` calls both functions and prints their results.
- The sample also prints separators so the outputs are easier to see.

## Run
```bash
go run .
```

## Expected Behavior
The program prints labels from `foo()` and `bar()`, followed by `10` and `15 Hello`.

## Key Takeaway
Go makes multiple return values a normal part of function design.
