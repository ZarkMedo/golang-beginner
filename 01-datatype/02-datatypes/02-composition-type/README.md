### CompositionType is a struct type
#### Collection/Aggregation or Non-Reference Types
1. Array
2. Struct
#### Reference types
// 1. Slice
// 2. Map
// 3. Function/Method
// 4. Channel
// 5. Pointer
#### Interface types
// 1. Interface


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