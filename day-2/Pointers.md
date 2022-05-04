# Pointers
Values are stored in memory addresses. **To see an address in memory in Go you use the** ```&``` **operator**

    package main

    import (
        "fmt"
    )

    func main() {
        a := 42
        fmt.Println(a)
        // 42

        fmt.Println(&a)
        // 0x1040a124 or equivalent

        fmt.Printf("%T", a)
        // int

        fmt.Printf("%T", &a)
        *int

        b := &a
        fmt.Println(b)
        // 0x1040a124 

        fmt.Println(*b)

    }

The ```*``` operator specifies that it is a **pointer** to a type, so ```*int``` would be described as **a pointer to an int**.

Everything in Go is passed by value. However, pointers allow you to share addresses where a value is stored. 

However, when you use the ```*``` operator in code, it will give you the **value stored at the address.** 

    b := &a // Ampersand gives you the address
    fmt.Println(b)
    // 0x1040a124 

    fmt.Println(*b) // Asterisk gives you the value stored at an address
    // 42

## When to Use Pointers
Pointers allow you to share a value stored in some memory location. Use poniters when:

1. You don't want to pass around a lot of data
2. You want to change data at a location

### Example (a):

    package main

    import (
        "fmt"
    )

    func main() {
        x := 0
        foo(x)
        fmt.Println(x)
        // 0 - Third Print
    }

    func foo(y int) {
        fmt.Println(y)
        // 0 - First Print
        y = 43
        fmt.Println(y)
        // 43 - Second Print
    }

The values you see printed above occur becuase GO PASSES VALUES

### Example (b)

    package main

    import (
        "fmt"
    )

    func main() {
        x := 0
        fmt.Println("x before: ", x)
        fmt.Println("&x before: ", &x)

        foo(&x)

        fmt.Println("x after: ", x)
        fmt.Println("&x after: ", &x)
    }

    func foo(y *int) {
        fmt.Println("y before: ", y)
        fmt.Println("*y before: ", *y)

        *y = 43
        // y is the address of the value
        // * operator means you're getting the value at the address
        // which is then reassigned by the = operator to 43

        fmt.Println("y after: ", y)
        fmt.Println("*y after: ", *y)
    }

    //OUTPUT 

    // x before:  0
    // &x before:  0xc000130000
    // y before:  0xc000130000
    // *y before:  0
    // y after:  0xc000130000
    // *y after:  43
    // x after:  43
    // &x after:  0xc000130000

You can see that the address stays contant throughout, but in this case the value changes as we reassign the value at the address, mutating it with:

 ```*y = 43```

