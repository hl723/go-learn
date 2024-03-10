# A Tour of Go - Concurrency

## Goroutines
- These are lightweight threads kicked off by the "go" keyword

- The function and arguments are evaluated when the goroutine is called, then the execution of the function happens in the new goroutine.

- They use the same address space, access to shared memory needs to be synchronized.

- The `sync` packages provides useful primitives.

## Channels
- Type conduit which values are sent and received.

- Use the channel operator `<-`. The data flows in the direction of the arrow.

```
c := make(chan int) // Create a channel for passing ints
ch := make(chan int, 2) // Creates a buffered channel of size 2.
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.
```

- By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

- Sends to a buffered channel will block only when it is full. Receives will block when the channel is empty.

- Deadlock occurs (runtime error) when channel buffer overflows!

### Closing Channels
- (Only) A sender can close a channel to indicate no more values will be sent. `close(c)` If a receiver closes a channel, it will panic. A channel need not be closed unless the receiver explicitly needs the channel closed.

- Receivers can test if a channel is closed by `v, ok := <-ch`, ok is false if channel is closed.

### Range
- Use `range c` to receive values from a channel until is closed.

### Select
- Similar to a switch statement, a select + case statements can let a goroutine wait on multiple communication operations.

- Select blocks until one case receives. Chooses one at random if multiple are ready.

- Use a default case if no other cases are ready.
```
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```

```
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

- See [Equivalent Binary Trees Example](./exercise-equivalent-binary-trees.go)

## sync.Mutex
- For Mutual Exclusion, we can use a mutex.

- It has both `Lock` and `Unlock` methods. 

- [Example of Web Crawler with Mutex as Cache](./exercise-web-crawler.go)
