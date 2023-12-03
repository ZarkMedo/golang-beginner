### CompositionType is a struct type
#### Collection/Aggregation or Non-Reference Types
1. Array
2. Struct
#### Reference types
1. Slice
2. Channel
3. Map
4. Pointer
5. Function/Method
#### Interface types
1. Interface


#### What the difference between Slice and Array?
| Aspect              | Array                   | Slice                                      |
| ------------------- | ----------------------- | ------------------------------------------ |
| Declaration         | var a [3]int            | var s []int                                |
| Size                | Fixed                   | Dynamic                                    |
| Usage               | Less common             | More common                                |
| Length Access       | len(arr)                | len(slice)                                 |
| Capacity Access     | N/A                     | cap(slice)                                 |
| Type                | Value                   | Reference                                  |
| Passing to Function | Passed by value         | Passed by reference                        |
| Reflection Type     | reflect.Array           | reflect.Slice                              |
| Zero Value          | [0 0 0]                 | nil                                        |
| Copying             | Copies by value         | References underlying array                |
| Creating            | var a = [3]int{1, 2, 3} | s := make([]int, 3) or s := []int{1, 2, 3} |
| Example             | [3]int{1, 2, 3}         | []int{1, 2, 3}                             |

#### What are Channel usages?
[effective_go-channels](https://go.dev/doc/effective_go#channels)
1. Communicate between goroutines
```go
ch := make(chan int)
go func() {
    ch <- 1
}()
fmt.Println(<-ch)
```
2. Implement a pipeline
```go
func main() {
    naturals := make(chan int)
    squares := make(chan int)

    // Counter
    go func() {
        for x := 0; x < 100; x++ {
            naturals <- x
        }
        close(naturals)
    }()

    // Squarer
    go func() {
        for x := range naturals {
            squares <- x * x
        }
        close(squares)
    }()

    // Printer (in main goroutine)
    for x := range squares {
        fmt.Println(x)
    }
}
```
3. Implement a worker pool
5. Implement a pub/sub pattern
```go
func main() {
    pubsub := make(chan interface{}, 3)
    done := make(chan bool)

    go func() {
        for msg := range pubsub {
            fmt.Println(msg)
        }
        done <- true
    }()

    pubsub <- "hello"
    pubsub <- "world"
    pubsub <- "goodbye"

    close(pubsub)
    <-done
}
```
6. Implement a semaphore pattern
```go
func main() {
    sem := make(chan bool, 10)

    for i := 0; i < 100; i++ {
        sem <- true
        go func(i int) {
            fmt.Println(i)
            <-sem
        }(i)
    }

}

```
7. Implement a fan-in/fan-out pattern
8. Implement a rate limiter
9.  Implement a timeout pattern
```go
func main() {
    ch := make(chan interface{})
    go func() {
        time.Sleep(5 * time.Second)
        ch <- "result"
    }()

    select {
    case res := <-ch:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout")
    }
}
```
10. Implement a circuit breaker pattern
11. Implement a retry pattern
```go
func main() {
    retry := func(
        count int,
        sleep time.Duration,
        fn func() error,
    ) error {
        var err error
        for i := 0; i < count; i++ {
            
             = fn()
            if err == nil {
                return nil
            }
            time.Sleep(sleep)
        }
        return err
    }

    err := retry(
        3,
        2*time.Second,
        func() error {
            fmt.Println("try")
            return errors.New("failed")
        },
    )
    fmt.Println(err)
}
```

#### Channel's Buffer Type
| Aspect               | Unbuffered Channel   | Buffered Channel                    |
| -------------------- | -------------------- | ----------------------------------- |
| Declaration          | var ch chan int      | var ch chan int = make(chan int, 3) |
| Size                 | 0                    | 3                                   |
| Usage                | Less common          | More common                         |
| Need to be closed?   | Yes                  | No                                  |
| Must use goroutine?  | Yes                  | No                                  |
| Send to channel      | ch <- 1              | ch <- 1                             |
| Receive from channel | <-ch                 | <-ch                                |
| Example              | ch := make(chan int) | ch := make(chan int, 3)             |

#### Pointer Vs Value 
- [Pointer vs Value](https://go.dev/doc/effective_go#pointers_vs_values)

| Aspect          | Pointer       | Value              |
| --------------- | ------------- | ------------------ |
| Declaration     | var p *int    | var v int          |
| Dereference     | *p = 1        | v = 1              |
| Address of      | p = &v        | N/A                |
| Reflection Type | reflect.Ptr   | default value type |
| Creating        | p := new(int) | v := 1             |
| unqueness       | Yes           | No                 |

- write the append function by value type
```go
type ByteSlice []byte
// must return the result slice.
func (slice ByteSlice) Append(data []byte) []byte {
    // Body exactly the same as the Append function defined above.
}
```
- write the append function by pointer type
```go
type ByteSlice []byte
// pointer is unqueness, so we don't need to return the result slice.
func (p *ByteSlice) Append(data []byte) {
    slice := *p
    // Body as above, without the return.
    *p = slice
}
// also can write like this， it's like the append function in the standard library.
func (p * ByteSlice) Append(data []byte) (int, error) {
    slice := *p
    // Body as above, without the return.
    *p = slice
    return len(data), nil
}
```

#### how to judge the struct is implemented the interface?
- struct's method set must contain all of the interface's method set.

#### all of the Composition Type in Go, and their zero value
| Type     | Zero Value |
| -------- | ---------- |
| Array    | [0 0 0]    |
| Struct   | {0 0}      |
| Slice    | nil        |
| Map      | nil        |
| Channel  | nil        |
| Pointer  | nil        |
| Function | nil        |
| Interface| nil        |
