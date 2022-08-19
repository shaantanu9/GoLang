"bool" use to define boolean values in GoLang

by default bool is false

We can use ! to change the value of bool as javascript

```go
package main
import "fmt"
func main() {
    var a bool
    fmt.Println(a) // false
    a = true
    fmt.Println(a) // true
    a = !a
    fmt.Println(a) // false
}
```
