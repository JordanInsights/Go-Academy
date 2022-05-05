# Udemy Course Concurrency Notes
Go is the first language to natively take advantage of multiple cores. This makes a significant difference to how you handle concurrency with Go. 

## Concurrency vs. Parallelism
Concurrency is a design pattern, concurrent code can itself run in parallel. Having more than one CPU allows you to run code in parallel (i.e. more than one thread running at once)

## Wait Group

    package main

    import (
        "fmt"
        "runtime"
        "sync"
    )

    var wg sync.WaitGroup

    func main() {
        fmt.Println("OS\t\t", runtime.GOOS )
        fmt.Println("ARCH\t\t", runtime.GOARCH)
        fmt.Println("CPU\t\t", runtime.NumCPU())
        fmt.Println("Goroutines\t", runtime.NumGoroutine())

        wg.Add(1)
        go foo()
        bar()
        
        fmt.Println("CPU\t\t", runtime.NumCPU())
        fmt.Println("Goroutines\t", runtime.NumGoroutine())

        wg.Wait()
    }
        
    func foo() {
        for i := 0; i <= 10; i++ {
            fmt.Println("foo: ", i)
        }
        wg.Done()
    }

    func bar() {
        for i := 0; i < 10; i++ {
            fmt.Println("bar: ", i)
        }
    }


## Race Conditions

    package main

    import (
        "fmt"
        "runtime"
        "sync"
    )

    func main() {
        fmt.Println("CPUs: ", runtime.NumCPU())
        fmt.Println("Goroutines: ", runtime.NumGoroutine())

        counter := 0
        const gs = 100
        
        var wg sync.WaitGroup
        wg.Add(gs)


        for i := 0; i < gs; i++ {
            go func() {
                v := counter
                runtime.Gosched()
                v++
                counter = v
                wg.Done()
            }()
        }

        wg.Wait()
        fmt.Println("Goroutines: ", runtime.NumGoroutine())
        fmt.Println("count: ", counter)
    }

    // IN THIS EXAMPLE COUNT DOESNT REACH 100 VAST MAJORITY OF TIME. WHICH IS OBVIOUSLY INCORRECT
    // ILLUSTRATES THE LIMITATIONS OF WAIT