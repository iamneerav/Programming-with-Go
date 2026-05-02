Use interfaces to calculate areas for different shapes.

# Hands-on 158: Interfaces and Polymorphism

## Purpose
Demonstrate how different types can satisfy the same interface and be processed through one shared function.

## What This Program Shows
- `SQUARE` and `CIRCLE` each implement `Area() float64`.
- `SHAPE` is an interface that requires `Area()`.
- `INFO()` accepts any `SHAPE` and returns its area.
- `main()` prints the area of both a square and a circle.

## Run
```bash
go run .
```

## Expected Behavior
The program prints the area of the square as `20` and the area of the circle as approximately `28.274333882308138`.

## Key Takeaway
Interfaces enable polymorphism by allowing different concrete types to be used through one shared contract.
