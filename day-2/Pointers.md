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