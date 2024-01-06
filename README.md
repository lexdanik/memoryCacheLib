In Memory Cache
================================

In Memory Cache create/get/delete in key-value cache

## Methods

```go
func main() {
    cache := cache.New() // Init cache
	cache.Set("userId", 42) // Set key-value in cache
    userId := cache.Get("userId") // Get value by key
    
    fmt.Println(userId)
    
    cache.Delete("userId") // Delete value from cache
    userId = cache.Get("userId")
    
    fmt.Println(userId)
}
```