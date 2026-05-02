Return a function from another function and call it later.

# Hands-on 166: Returning Functions

## Purpose
Demonstrate higher-order functions by returning a function from another function.

## What This Program Shows
- `address(city, state)` returns a new function.
- The returned function closes over `city` and `state`.
- `main()` stores that function in `result` and calls it.

## Run
```bash
go run .
```

## Expected Behavior
The program prints `Toronto, Ontario, Canada`.

## Key Takeaway
Functions can return other functions, which is a common foundation for closures.
