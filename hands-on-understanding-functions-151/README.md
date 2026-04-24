# Hands-on 151: Recursion

## Purpose
Demonstrate recursion by calculating a factorial.

## What This Program Shows
- `factorial(a int) int` calls itself with `a-1`.
- The base case is `a == 1`, which stops the recursive calls.
- `main()` prints `factorial(5)`.

## Run
```bash
go run .
```

## Expected Behavior
The program prints `120`.

## Key Takeaway
Recursive functions must always have a base case, otherwise they will keep calling themselves indefinitely.
