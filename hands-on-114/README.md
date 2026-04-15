# hands-on-114

## Purpose
Practice creating a slice with reserved capacity, appending data, and iterating over the result.

## What I Practiced
- Creating an empty `[]string` with `make` and a capacity of 50
- Appending all 50 U.S. states into the slice
- Printing slice length, capacity, and values by index

## Run
```bash
go run .
```

## Notes
- The exercise starts with a zero-length slice and preallocates enough capacity for all values.
- After appending the states, the program prints the updated length and capacity.
- A `for` loop then walks the slice by index and prints each state.