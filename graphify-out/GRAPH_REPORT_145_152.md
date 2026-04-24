# Graph Report - .  (2026-04-24)

## Corpus Check
- Corpus is ~2,869 words - fits in a single context window. You may not need a graph.

## Summary
- 51 nodes · 71 edges · 10 communities detected
- Extraction: 92% EXTRACTED · 8% INFERRED · 0% AMBIGUOUS · INFERRED: 6 edges (avg confidence: 0.5)
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- [[_COMMUNITY_Community 0|Community 0]]
- [[_COMMUNITY_Community 1|Community 1]]
- [[_COMMUNITY_Community 2|Community 2]]
- [[_COMMUNITY_Community 3|Community 3]]
- [[_COMMUNITY_Community 4|Community 4]]
- [[_COMMUNITY_Community 5|Community 5]]
- [[_COMMUNITY_Community 6|Community 6]]
- [[_COMMUNITY_Community 7|Community 7]]
- [[_COMMUNITY_Community 8|Community 8]]
- [[_COMMUNITY_Community 9|Community 9]]

## God Nodes (most connected - your core abstractions)
1. `main()` - 16 edges
2. `foo()` - 5 edges
3. `writeLog()` - 3 edges
4. `bar()` - 2 edges
5. `calculator()` - 2 edges
6. `addnumber()` - 2 edges
7. `subnumber()` - 2 edges
8. `incrementor()` - 2 edges
9. `factorial()` - 2 edges
10. `book` - 2 edges

## Surprising Connections (you probably didn't know these)
- `main()` --calls--> `bar()`  [EXTRACTED]
  graphify-out\review-145-152-input\hands-on-understanding-functions-152\main.go → graphify-out\review-145-152-input\hands-on-understanding-functions-147\main.go
- `main()` --calls--> `incrementor()`  [EXTRACTED]
  graphify-out\review-145-152-input\hands-on-understanding-functions-152\main.go → graphify-out\review-145-152-input\hands-on-understanding-functions-149\main.go
- `main()` --calls--> `factorial()`  [EXTRACTED]
  graphify-out\review-145-152-input\hands-on-understanding-functions-152\main.go → graphify-out\review-145-152-input\hands-on-understanding-functions-151\main.go
- `main()` --calls--> `foo()`  [EXTRACTED]
  graphify-out\review-145-152-input\hands-on-understanding-functions-152\main.go → graphify-out\review-145-152-input\hands-on-understanding-functions-150\main.go
- `main()` --calls--> `addnumber()`  [EXTRACTED]
  graphify-out\review-145-152-input\hands-on-understanding-functions-152\main.go → graphify-out\review-145-152-input\hands-on-understanding-functions-148\main.go

## Hyperedges (group relationships)
- **Core Function Concepts** —  [INFERRED]
- **Advanced Function Patterns** —  [INFERRED]
- **Callback Implementation** —  [INFERRED]
- **Closure with State** —  [INFERRED]
- **Wrapper with Interface** —  [INFERRED]
- **Recursion Pattern** —  [INFERRED]
- **Hands-on 145-152 Series** —  [INFERRED]

## Communities

### Community 0 - "Community 0"
Cohesion: 0.0
Nodes (9): Anonymous Function, foo(), Function Expression, Function Fundamentals, hands-on-145: Anonymous Functions, hands-on-146: Function Expressions, hands-on-147: Returning Functions, hands-on-150: Function Fundamentals (+1 more)

### Community 1 - "Community 1"
Cohesion: 0.0
Nodes (8): addnumber(), calculator(), Callback Pattern, Closure, hands-on-148: Callbacks, hands-on-149: Closures, incrementor(), subnumber()

### Community 2 - "Community 2"
Cohesion: 0.0
Nodes (8): book.String(), count.String(), Error Wrapping, fmt.Stringer Interface, hands-on-152: Wrapper Functions, readFile(), Wrapper Function, writeLog()

### Community 3 - "Community 3"
Cohesion: 0.0
Nodes (3): book, count, writeLog()

### Community 4 - "Community 4"
Cohesion: 0.0
Nodes (4): addnumber(), calculator(), main(), subnumber()

### Community 5 - "Community 5"
Cohesion: 0.0
Nodes (1): foo()

### Community 6 - "Community 6"
Cohesion: 0.0
Nodes (4): Base Case (Recursion), factorial(), hands-on-151: Recursion, Recursion

### Community 7 - "Community 7"
Cohesion: 0.0
Nodes (1): incrementor()

### Community 8 - "Community 8"
Cohesion: 0.0
Nodes (1): bar()

### Community 9 - "Community 9"
Cohesion: 0.0
Nodes (1): factorial()

## Knowledge Gaps
- **Thin community `Community 7`** (2 nodes): `main.go`, `incrementor()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 8`** (2 nodes): `main.go`, `bar()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 9`** (2 nodes): `main.go`, `factorial()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.