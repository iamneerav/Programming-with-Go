# Hands-on 148: Callbacks

## Purpose
Show how one function can be passed into another function to decide behavior at runtime.

## What This Program Shows
- `addnumber()` adds two integers.
- `subnumber()` subtracts two integers.
- `calculator()` accepts a function parameter of type `func(int, int) int`.
- The same `calculator()` function is reused with different callback functions.

## Run
```bash
go run .
```

## Expected Behavior
The program prints direct addition and subtraction results, then prints the same operations again through the callback-based `calculator()` function.

## Key Takeaway
Callbacks make code flexible because the caller can choose what logic the receiving function should execute.
