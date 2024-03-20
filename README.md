In Memory Cache
================================

In Memory Cache create/get/delete in key-value cache with TTL

## Methods

```go
func main() {
    cache := cache.New(time.Second*5, time.Second*5) // Init cache
    cache.Set("userId", 42, time.Second*5)           // Set key-value in cache
    userId, err := cache.Get("userId")               // Get value by key
    if err != nil {
    log.Fatal(err)
    }
    fmt.Println(userId)
    
    time.Sleep(time.Second * 6)
    
    userId2, err := cache.Get("userId") // Get value by key
    if err != nil {
    log.Fatal(err)
    }
    
    fmt.Println(userId2)
}
```