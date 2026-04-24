# Hands-on 147: Returning a Function

## Purpose
Demonstrate that a Go function can return another function as a value.

## What This Program Shows
- `foo()` returns an `int`.
- `bar()` returns a function of type `func() int`.
- The returned function is stored in variable `y` and called later with `y()`.
- The program also prints the types of `foo`, `bar`, and `y`.

## Run
```bash
go run .
```

## Expected Behavior
The program prints the result from `foo()`, shows the returned function value from `bar()`, and then prints the result of calling that returned function.

## Key Takeaway
Returning functions is the basis for patterns such as callbacks, factories, and closures.
