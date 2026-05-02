Pass one function into another as a callback.

# Hands-on 167: Callback Functions

## Purpose
Show how one function can receive another function and use it to complete part of its work.

## What This Program Shows
- `location(street, city, c)` accepts a callback function.
- `country(a string)` matches the callback signature and returns the country name.
- `main()` passes `country` into `location()` and prints the formatted result.

## Run
```bash
go run .
```

## Expected Behavior
The program prints `I am in Merivale Road, Ottawa, Canada`.

## Key Takeaway
Callbacks let you inject behavior into a function without changing the function itself.
