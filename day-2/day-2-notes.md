# Day 2 Notes
## Functions
### Basics
You define a function with parameters (if any). You call a function with arguments (if any).
    
    func (r receiver) identifier(parameter(s)) (return(s)) {
        // code
    }

    func foo(s string) string {
        fmt.Println("This function will print a string and return a string 'bar'")
        return s
    }

Everything in Go is PASS BY VALUE

### Multiple Returns
    func mouse(fn string, ln string) (string, bool) {
        a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
        b := false
        return a, b
    }

    x, y := mouse("Ian", "Fleming")
    
    // OUTPUT = IanFleming, says "Hello"

### Variadic Parameter
Variadic parameters return a slice of what you pass in. You can pass in 0 -> inf arguments of the type specified.  

    func foo(x ...int) {
        fmt.Prinln(x)
        fmt("T\n", x)
    }

    foo(1, 2, 3)
    // OUTPUT
    // [1, 2, 3]
    // []int

    func bar(x ...int) int {
        sum := 0
        for _, v := range x {
            sum += v
        }
        return sum
    }

    total := bar(1, 2, 3)
    fmt.Println("Total is equal to: ", total)
    
    // OUTPUT
    // Total is equal to:  6

Variadic parameters must be the last parameters you pass in

    foo(s string, x ...int) {
        // code
    }

If no arguments are passed for the variadic parameter, the value of the parameters is ```nil``` NOT an empty array. 

#### Unfurling
    func bar(x ...int) int {
        sum := 0
        for _, v := range x {
            sum += v
        }
        return sum
    }
    
    xi := []int{1, 2, 3}
    total := bar(xi...) \\ <--- This is the unfurling

    fmt.Println("Total is equal to: ", total)
    // OUTPUT
    // Total is equal to:  6




## Glossary 
- Lexical Element: Catch all term for elements in Go. For example, comments, operators, tokens, identifiers etc. 
- Variadic: Passing in zero or more values
