# First Program in Go

This is a simple "Hello, World!" program written in Go (Golang). It demonstrates the basic syntax and structure of a Go program.

## Code Explanation

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")  // Outputs: Hello, World!
}
```
## Key Points:
- `package main`: Defines the package as main, which is required if you want to create an executable program.
- `import "fmt"`: Imports the fmt package, which provides functions for formatting text, including printing to the console.
- `func main()`: The main function is the entry point of the program. When you run the program, the code inside the main function is executed.
- `fmt.Println("Hello, World!")`: This line prints the string "Hello, World!" to the console.

## Steps to Run the Program
- Initialize Go Module
> Run the following command to create a go.mod file in the current directory:

```bash
go mod init HelloWorld
```
- Run the Program
> To execute the Go program, use the go run command:

```bash
go run main.go
```
