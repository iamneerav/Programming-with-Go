# Hands-on 149: Closures

## Purpose
Demonstrate how a returned function can keep access to variables from its surrounding scope.

## What This Program Shows
- `incrementor()` returns a function of type `func() int`.
- The returned function updates and remembers the variable `x` across calls.
- Closure `f` keeps its own state.
- Closure `g` is created separately and starts with its own fresh state.

## Run
```bash
go run .
```

## Expected Behavior
Repeated calls to `f()` keep increasing the same captured value, while `g()` starts over from a separate captured value.

## Key Takeaway
Closures let a function preserve state without using a global variable.
