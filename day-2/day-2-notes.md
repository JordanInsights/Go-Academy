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

## Defer
A defer statement invokes a function whose execution is deferred to the moment the surrouding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponsing goroutine is panicking. 

```DeferStmt = "defer" Expression .```

    package main
    
    import (
        "fmt
    )

    func main() {
        defer foo()
        bar()
    }

    func foo() {
        fmt.Println("foo)
    }

    func bar() {
        fmt.Println("bar)
    }

    // OUTPUT
    // bar
    // foo

## Methods
    package main

    import (
        "fmt
    )

    type person struct {
        first string
        last string
    }

    type secretAgent struct {
        person
        ltk bool
    }

    // any VALUE of TYPE secretAgent has access to the speak method
    func (s secretAgent) speak() {
        fmt.Println("I am", s.first, s.last)
    }

    func main() {
        sa1 := secretAgent{
            person: person{
                "James",
                "Bond",
            },
            ltk: true,
        }

        fmt.Println(sa1)
        sa1.speak()
    }

## Interfaces
Interfaces allow you to define behaviour. They also allow you to implement polymorphism

    package main

    import (
        "fmt
    )

    type person struct {
        first string
        last string
    }

    type secretAgent struct {
        person
        ltk bool
    }

    func (s secretAgent) speak() {
        fmt.Println("I am", s.first, s.last, " - the secret agent speak")
    }

    func (p person) speak() {
        fmt.Println("I am", p.first, p.last, " - the person speak")
    }


    // ANY STRUCT WITH A METHOD of SPEAK IS OF TYPE HUMAN
    type human interface {
        speak()
    }

    func bar (h human) {
        fmt.Println("I was passed into bar", h)
    }

    func main() {
        sa1 := secretAgent{
            person: person{
                "James",
                "Bond",
            },
            ltk: true,
        } 
        
        p1 := person{
            "Jordan",
            "Shaw",
        }
        

        fmt.Println(sa1)
        sa1.speak()
        bar(sa1)

        fmt.Println(p1)
        p1.speak()
        bar(p1)

    }

**A VALUE CAN BE OF MORE THAN ONE TYPE.** Both value of type secret agent and or person ARE ALSO OF TYPE human due to the fact that they both have 'speak()' methods. 

## Anonymous functions

    package main
    
    import (
        "fmt"
    )

    main() {
        func() {
            fmt.Println("This anonymous function is called by the suffixed parentheses)
        }()

        func(x int) {
            fmt.Println("This anonymous function is called with the meaning of life: ", x)
        }(42)
    }


## Func Expression
An expreesion is something that resolves to a value. The value here is a function we will assign to a variable

    package main
    
    import (
        "fmt"
    )

    f := func() {
        fmt.Println("my first func expression")
    }

    f()

    // OUTPUT
    // my first func expression

## Returning a function
    
    package main

    import (
        "fmt
    )

    func main() {
        // foo() returns a string
        s1 := foo()
        fmt.Println(s1)

        // bar() returns a function
        s2 := bar()
        fmt.Printf("%T", s2)
        // OUTPUT
        // func() int

        x := s2()
        fmt.Println("\n", x)
        // OUTPUT 
        // 1984
    }

    func foo() string {
        s := "Hello world!"
        return s
    }

    func bar() func() int {
        return func() int {
            return 1984
        }
    }

## Closure
Closure is a concept where you enclose the scope of a variable and contain it to one area, thereby limiting its scope. In general, you want to keep the scope of a variable as narrow as possible. 

    package main

    import (
        "fmt"
    )

    func main() {
        a := incrementer()
        b := incrementer()

        fmt.Println("First call of a: ", a()) 
        fmt.Println("Second call of a: ", a()) 
        
        // first return value of b is 1 as calling incrementer returned functions with their own scope
        fmt.Println("First call of b: ", b()) 
    }

    func incrementer() func () int {
        var x int
        return func() int {
            x++
            return x
        }
    }

## Glossary 
- Lexical Element: Catch all term for elements in Go. For example, comments, operators, tokens, identifiers etc. 
- Variadic: Passing in zero or more values
