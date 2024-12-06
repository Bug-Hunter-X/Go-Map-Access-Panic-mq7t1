```go
func main() {
    var m map[string]int
    key := "myKey"
    if m == nil {
        fmt.Println("Map is nil, cannot access.")
    } else {
        fmt.Println(m[key])
    }
}
```