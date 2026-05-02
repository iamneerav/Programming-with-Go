Compare variadic arguments with passing a slice directly.

# Hands-on 155: Variadic Functions

## Purpose
Show the difference between passing values to a variadic function and passing a slice to a regular function.

## What This Program Shows
- `foo(x ...int)` accepts a variadic list of integers.
- `bar(a []int)` accepts a slice directly.
- Both functions sum the same set of numbers.
- `main()` demonstrates that both approaches can produce the same result.

## Run
```bash
go run .
```

## Expected Behavior
The program prints `150`, a separator line, and then `150` again.

## Key Takeaway
Variadic parameters are convenient when callers want to pass individual values, while slices are useful when the data is already grouped.
