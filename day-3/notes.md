# Concurrency, Goroutines, & Channels 
## Concurrency:
It means multiple computations are happning at the same time. Concurrency is everywhere in modern programming, whether we like it or not.

Two common models:
- Shared memory
- Message passing

Pricess: A process is an instance of a running program that is isolated from other processes on the same machine. In particular, it has its own private section of the machine's memory. 

A thread is a locus of control. Threads are automatically ready for shared memory.

A process is mostly isolated, whereas threads share memory. 

#

## Actor Model:
Issues arise when two separate things try and make changes to the same variable at the same time. 

### What is concurrency?
It is simply doing work in parallel. Conccurency fails when we:

- Mutate shared state
- Shared across goroutines

As a result, the data become corrupted, and unpredictable. Example of bad code:

    // Protected Data resource - not concurrent safe
    var coffeeOrder = []byte("        ")

    func main() {
        go writeText("Choca          ")
        go writeText("Mocha with Milk")
        go writeText("Banana Shake   ")
        go writeText("Tea no sugar   ")

        // What on earth will be in coffeeOrder now?
        time.Sleep(1 * time.Second)
    }

    func writeText(newORder string) {
        // CRITICAL SECTION STARTS
        orderAsBytes := []byte(newOrder)
        for index, b := range orderAsBytes {
            coffeeOrder[index] = b
            time.Sleep(10 * time.Millisecond)
        }

        // CRITICAL SECTION ENDS
    }

    // OUTPUT (something like...)
    // Chchna s g r

## Traditional Solutions
- Mutexes, i.e. "Mutual Exclusives". Sometimes known as locks. 
    - Mutexes can be visualised as a lock on a door. When there:
        - If Unlocked
            - Enter the protected area
            - Lock tje door
            - Do the thing
            - Unlock the doot
            - Leave
        - If Locked
            - Wait until your turn and door is unlocked
    - In go: 
        - ```mutex.Lock()```
            - Wait in line to be unlocked
            - Lock
        - Run 'Critical Section' Code
            - The code that needs making safe
        - ```mutex.Lock()```
            - Open the lock for others

Example code:        
        
    var coffeeOrder = []byte("        ")

    func main() {
        go writeText("Choca          ")
        go writeText("Mocha with Milk")
        go writeText("Banana Shake   ")
        go writeText("Tea no sugar   ")

        // What on earth will be in coffeeOrder now?
        time.Sleep(1 * time.Second)
    }

    func writeText(newORder string) {
        
        mutex.Lock()

        defer mutex.Unlock()

        // CRITICAL SECTION STARTS
        orderAsBytes := []byte(newOrder)
        for index, b := range orderAsBytes {
            coffeeOrder[index] = b
            time.Sleep(10 * time.Millisecond)
        }

        // CRITICAL SECTION ENDS
    }

    // OUTPUT (something like...)
    // Chchna s g r

## Bet365 DO NOT ALLOW THE USE OF MUTEXES AS THEY HAVE POOR PERFROMANCE AT HIGH LOAD. 