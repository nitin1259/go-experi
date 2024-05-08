# Day 2

## A. String:

Strings are a sequence of characters represented in GoLang as a slice of bytes.
They are immutable, meaning once created, their contents cannot be changed.
GoLang provides a rich set of built-in functions and methods for manipulating strings,
including concatenation, slicing, searching, and formatting.

Certainly! Let's cover each topic with examples:

### 1. Creating strings:

Strings in GoLang are created using string literals enclosed within double quotes (`"`). Here's an example:

```go
str := "Hello, GoLang!"
```

### 2. String literals and escaping:

String literals can contain special characters that need to be escaped using backslashes (`\`). For example:

```go
str := "This is a \"quoted\" string."
```

### 3. String manipulation methods:

GoLang provides various built-in functions for manipulating strings:

- **Length**: Get the length of a string using `len()`.
- **Substring**: Extract substrings using slicing.
- **Concatenation**: Concatenate strings using the `+` operator or `fmt.Sprintf`.

```go
str := "Hello, World!"
length := len(str)           // length = 13
substring := str[0:5]        // substring = "Hello"
concatenated := str + "123"  // concatenated = "Hello, World!123"
```

### 4. Unicode support:

GoLang has native support for Unicode characters, allowing you to work with international text seamlessly.

```go
str := "こんにちは、世界！"
```

### 5. String formatting:

String formatting allows you to create formatted strings using placeholders. GoLang provides `fmt.Sprintf` for string formatting.

```go
name := "Alice"
age := 30
formatted := fmt.Sprintf("Name: %s, Age: %d", name, age)
// formatted = "Name: Alice, Age: 30"
```

### 6. String conversions:

GoLang provides the `strconv` package for converting between strings and other data types.

```go
str := "123"
num, _ := strconv.Atoi(str)       // num = 123
numStr := strconv.Itoa(456)       // numStr = "456"
```

### 7. Working with runes:

Runes in GoLang represent Unicode characters and are aliases for `int32`. You can iterate over strings rune-by-rune using a `for` loop.

```go
str := "こんにちは"
for _, runeValue := range str {
    fmt.Println(string(runeValue))
}
```

These examples cover the basics of working with strings in GoLang, including creation, manipulation, formatting, and conversion, as well as support for Unicode characters and runes. Experimenting with these examples and exploring further will deepen your understanding of string handling in GoLang.

## B. Input/output:

Input/output operations in GoLang involve reading data from standard input, files,
or other sources, and writing data to standard output, files, or other destinations.
GoLang provides the `fmt` package for formatted I/O and the `os` package for file I/O operations.

### 1. Reading input from the console (fmt.Scan, fmt.Scanln, fmt.Scanf):

GoLang provides several functions in the `fmt` package for reading input from the console:

- **`fmt.Scan`**: Reads space-separated values from the standard input.
- **`fmt.Scanln`**: Reads values until a newline character from the standard input.
- **`fmt.Scanf`**: Reads formatted input from the standard input based on a format specifier.

```go
package main

import "fmt"

func main() {
    var name string
    var age int

    fmt.Println("Enter your name:")
    fmt.Scanln(&name) // Reads a line of input

    fmt.Println("Enter your age:")
    fmt.Scan(&age) // Reads a space-separated value

    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

### 2. Writing output to the console (fmt.Println, fmt.Printf):

GoLang provides functions in the `fmt` package for printing output to the console:

- **`fmt.Println`**: Prints values followed by a newline character to the standard output.
- **`fmt.Printf`**: Formats and prints values based on format specifiers to the standard output.

```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 30

    fmt.Println("Name:", name)
    fmt.Printf("Age: %d\n", age)
}
```

### 3. Reading and writing files (os.Open, os.Create, io/ioutil.ReadFile, io/ioutil.WriteFile):

GoLang provides the `os` package for file I/O operations and the `io/ioutil` package for reading and writing files.

- **`os.Open`**: Opens a file for reading.
- **`os.Create`**: Creates a new file or truncates an existing file for writing.
- **`io/ioutil.ReadFile`**: Reads the contents of a file into a byte slice.
- **`io/ioutil.WriteFile`**: Writes data to a file.

```go
package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
    // Writing to a file
    data := []byte("Hello, GoLang!")
    err := ioutil.WriteFile("output.txt", data, 0644)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("File written successfully!")

    // Reading from a file
    fileData, err := ioutil.ReadFile("output.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("File contents:", string(fileData))
}
```

### 4. Buffered I/O (bufio package):

The `bufio` package in GoLang provides buffered I/O operations, which can improve performance when reading from or writing to files.

- **`bufio.NewReader`**: Creates a new buffered reader.
- **`bufio.NewWriter`**: Creates a new buffered writer.

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Reading from a file using buffered I/O
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Writing to a file using buffered I/O
    outFile, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer outFile.Close()

    writer := bufio.NewWriter(outFile)
    _, err = writer.WriteString("Hello, GoLang!\n")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    writer.Flush()
    fmt.Println("File written successfully!")
}
```

### 5. Error handling in I/O operations:

Error handling is crucial in I/O operations to handle potential errors gracefully. GoLang provides built-in error handling mechanisms using `error` values returned by I/O functions.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Reading from a file
    file, err := os.Open("nonexistent.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Writing to a file
    outFile, err := os.Create("/etc/hosts") // Permission denied
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer outFile.Close()
}
```

## C. Struct, Method & Interface:

Structs are composite data types in GoLang that allow you to group together different data types under one name.
Methods are functions associated with a type, and interfaces define a set of methods that a
type must implement to satisfy the interface.

### 1. Declaring and initializing structs:

In GoLang, structs are used to create custom data types composed of fields of different data types.

```go
package main

import "fmt"

// Declare a struct type
type Person struct {
    Name string
    Age  int
}

func main() {
    // Initialize a struct instance
    person := Person{
        Name: "Alice",
        Age:  30,
    }
    fmt.Println(person)
}
```

### 2. Accessing struct fields:

You can access struct fields using dot notation.

```go
fmt.Println("Name:", person.Name)
fmt.Println("Age:", person.Age)
```

### 3. Struct embedding and composition:

Struct embedding allows you to embed one struct type within another to achieve composition.

```go
type Address struct {
    City  string
    State string
}

type Employee struct {
    Name    string
    Age     int
    Address // Embedded struct
}
```

### 4. Methods and receivers:

Methods are functions associated with a type. Receivers specify the type on which a method operates.

```go
func (e Employee) Print() {
    fmt.Printf("Name: %s, Age: %d\n", e.Name, e.Age)
}

// Call the method
employee.Print()
```

### 5. Pointer receivers vs. value receivers:

Pointer receivers allow methods to modify the receiver struct, while value receivers create a copy of the receiver struct.

```go
func (e *Employee) UpdateName(newName string) {
    e.Name = newName
}

// Call the method
employee.UpdateName("Bob")
```

### 6. Interfaces and type assertion:

Interfaces in GoLang define a set of methods that a type must implement to satisfy the interface.

```go
type Printable interface {
    Print()
}

// Implement the interface
func (e Employee) Print() {
    fmt.Printf("Name: %s, Age: %d\n", e.Name, e.Age)
}

// Type assertion
var printable Printable = employee
```

### 7. Implementing interfaces implicitly and explicitly:

GoLang allows both implicit and explicit implementation of interfaces.

```go
// Implicit implementation
type Printable interface {
    Print()
}

type Employee struct {
    Name string
    Age  int
}

func (e Employee) Print() {
    fmt.Printf("Name: %s, Age: %d\n", e.Name, e.Age)
}

// Explicit implementation
type PrintableEmployee struct {
    Employee
}

func (p PrintableEmployee) Print() {
    fmt.Printf("Printable Employee - Name: %s, Age: %d\n", p.Name, p.Age)
}
```

### 8. Empty interfaces (interface{}) and type switches:

Empty interfaces can hold values of any type, useful for working with unknown types.

```go
func PrintValue(value interface{}) {
    switch v := value.(type) {
    case int:
        fmt.Println("Integer value:", v)
    case string:
        fmt.Println("String value:", v)
    default:
        fmt.Println("Unknown type")
    }
}

PrintValue(42)      // Integer value: 42
PrintValue("hello") // String value: hello
PrintValue(true)    // Unknown type
```

## D. Errors:

Errors handling is an essential aspect of robust software development.
GoLang provides a built-in error type and idiomatic ways to handle errors using the `error` interface
and the `errors` package.

### 1. Go error handling philosophy:

GoLang's error handling philosophy emphasizes explicit error checking and handling,
making it difficult to ignore errors and encouraging developers to handle them gracefully.

### 2. Error types and the `error` interface:

Errors in GoLang are represented by the `error` interface, which has a single method `Error()`
that returns a string describing the error.

```go
package main

import (
    "errors"
    "fmt"
)

func divide(x, y int) (int, error) {
    if y == 0 {
        return 0, errors.New("division by zero")
    }
    return x / y, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

### 3. Returning errors from functions:

Functions in GoLang can return errors alongside other return values to indicate failure.

### 4. Handling errors using `if` statements:

Errors can be checked using `if` statements to handle them gracefully.

### 5. Error propagation using `defer`, `panic`, and `recover`:

GoLang provides `defer`, `panic`, and `recover` for error propagation and handling panics.

```go
func readData() error {
    file, err := os.Open("data.txt")
    if err != nil {
        return err
    }
    defer file.Close()

    // Read data from file
}

func main() {
    err := readData()
    if err != nil {
        log.Fatal(err)
    }
}
```

### 6. Custom error types:

Developers can define custom error types to provide more context or additional information about errors.

```go
type MyError struct {
    Msg  string
    Code int
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}

func somethingThatMayFail() error {
    return &MyError{"Something went wrong", 500}
}
```

### 7. Error wrapping (`fmt.Errorf`, `errors.Wrap`, `errors.Wrapf`):

GoLang provides functions for wrapping errors with additional context using the `fmt` and `errors` packages.

```go
func main() {
    err := errors.New("original error")
    wrappedErr := fmt.Errorf("wrapped error: %w", err)
    fmt.Println(wrappedErr)

    // Using errors.Wrap
    err = errors.New("original error")
    wrappedErr = errors.Wrap(err, "wrapped error")
    fmt.Println(wrappedErr)

    // Using errors.Wrapf
    err = errors.New("original error")
    wrappedErr = errors.Wrapf(err, "wrapped error: %s", "additional info")
    fmt.Println(wrappedErr)
}
```

### 8. Error handling best practices:

- Always check errors after function calls.
- Propagate errors up the call stack when appropriate.
- Use specific error types to provide context.
- Provide informative error messages.
- Handle errors gracefully without crashing the program.

## E. Logging:

Logging is crucial for monitoring and debugging applications. GoLang provides the `log` package for basic logging functionality and the `logrus`, `zap`, and other third-party packages for more advanced logging features.

Certainly! Let's delve into each topic in depth:

### 1. Basic logging using the `log` package (`log.Print`, `log.Printf`, `log.Println`):

The `log` package in GoLang provides basic logging functionality for printing messages to the standard output.

```go
package main

import "log"

func main() {
    log.Print("This is a log message")
    log.Printf("This is a formatted log message with value: %d", 42)
    log.Println("This is a log message with a newline")
}
```

### 2. Configuring log output:

The `log` package does not provide extensive configuration options. However, you can redirect log output to a file or another destination by setting `log.SetOutput()`.

```go
package main

import (
    "log"
    "os"
)

func main() {
    file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    log.SetOutput(file)
    log.Println("Logging to a file")
}
```

### 3. Logrus and Zap overview:

Logrus and Zap are popular logging libraries in GoLang known for their performance, flexibility, and extensive features. They offer structured logging, multiple log levels, and support for various output formats.

### 4. Log levels (debug, info, warning, error, fatal):

Both Logrus and Zap support log levels such as `debug`, `info`, `warning`, `error`, and `fatal`. You can set the log level to filter messages based on their severity.

### 5. Structured logging:

Structured logging involves logging data in a structured format, such as JSON or key-value pairs,
to facilitate easy parsing and analysis.

### 6. Logging to files and other destinations:

Logrus and Zap provide built-in support for logging to files, standard output, standard error,
syslog, and custom destinations.

### 7. Contextual logging and loggers inheritance:

Contextual logging involves associating additional context or metadata with log messages,
such as request IDs or user IDs. Loggers inheritance allows creating child loggers with inherited
properties from parent loggers.

### 8. Logging best practices:

- Use descriptive log messages that provide context.
- Log errors and other important events.
- Use appropriate log levels based on the severity of the message.
- Avoid excessive logging to prevent performance overhead.
- Log sensitive information securely (avoid logging passwords or personally identifiable information).
  hese topics cover the fundamentals of logging in GoLang and introduce more advanced concepts such as structured logging, log levels, and best practices. Experimenting with different logging libraries and exploring further will deepen your understanding of logging in GoLang.

## F. Testing:

Testing is an integral part of software development to ensure the correctness and reliability of code.
GoLang comes with a built-in testing framework, making it easy to write and execute tests for Go packages.

### 1. Writing unit tests (`*_test.go` files):

In GoLang, unit tests are written in files with names ending in `_test.go`.
These files contain test functions that verify the behavior of the code under test.

```go
// Example_test.go

package example

import "testing"

func Add(a, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

### 2. Test functions and test cases:

Test functions in GoLang are functions with names starting with `Test`. Each test function can contain multiple test cases.

```go
func TestAdd(t *testing.T) {
    t.Run("Add positive numbers", func(t *testing.T) {
        result := Add(2, 3)
        expected := 5
        if result != expected {
            t.Errorf("Add(2, 3) = %d; want %d", result, expected)
        }
    })
    t.Run("Add negative numbers", func(t *testing.T) {
        result := Add(-2, -3)
        expected := -5
        if result != expected {
            t.Errorf("Add(-2, -3) = %d; want %d", result, expected)
        }
    })
}
```

### 3. Table-driven tests:

Table-driven tests involve using test tables to specify input and expected output for multiple test cases.

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {2, 3, 5},
        {-2, -3, -5},
        {0, 0, 0},
    }
    for _, tt := range tests {
        result := Add(tt.a, tt.b)
        if result != tt.expected {
            t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
        }
    }
}
```

### 4. Subtests and test grouping:

Subtests allow organizing tests into groups and executing them individually.

```go
func TestAdd(t *testing.T) {
    t.Run("Add positive numbers", func(t *testing.T) {
        // Test cases for adding positive numbers
    })
    t.Run("Add negative numbers", func(t *testing.T) {
        // Test cases for adding negative numbers
    })
}
```

### 5. Testing utilities (`testing.T`, `testing.M`, `testing.B`):

The `testing` package in GoLang provides utilities for writing tests, including the `testing.T` type for testing, `testing.M` type for running multiple tests, and `testing.B` type for benchmarking.

### 6. Benchmarking (`go test -bench`):

Benchmark tests in GoLang measure the performance of code by running it repeatedly and reporting execution time.

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

### 7. Test coverage (`go test -cover`):

Test coverage in GoLang measures the percentage of code covered by tests.

```sh
go test -cover
```

### 8. Mocking and dependency injection:

Mocking involves creating fake implementations of dependencies for testing purposes. Dependency injection allows injecting dependencies into code at runtime, making it easier to test.

### 9. Testing best practices:

- Write tests before writing code (test-driven development).
- Keep tests small, focused, and independent.
- Use descriptive test names and comments.
- Test boundary conditions and edge cases.
- Use table-driven tests for repetitive test cases.
- Refactor tests regularly to maintain readability.
