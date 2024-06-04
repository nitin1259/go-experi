# questions

Question: What is the time complexity of len(...) for each data type

```go
  package main

  import (
  "fmt"
  )

  func main() {
  _ = len("string")
  _ = len([]int{1, 2, 3})
  \_ = len(map[string]int{"one": 1, "two": 2, "three": 3})
  }

  The answer to this question is very easy:
  The time complexity of the len(...) for all cases is O(1).
```
